package action

import (
	"errors"
	"fmt"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"github.com/oracle/bosh-oracle-cpi/bmc/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"

	fakeuuid "github.com/cloudfoundry/bosh-utils/uuid/fakes"
	clientfakes "github.com/oracle/bosh-oracle-cpi/bmc/client/fakes"
	vmfakes "github.com/oracle/bosh-oracle-cpi/bmc/vm/fakes"
	registryfakes "github.com/oracle/bosh-oracle-cpi/registry/fakes"
)

var _ = Describe("CreateVM", func() {
	var (
		connector      *clientfakes.FakeConnector
		logger         boshlog.Logger
		registryClient *registryfakes.FakeClient
		uuidGen        *fakeuuid.FakeGenerator
		creator        *vmfakes.FakeVMCreator

		err             error
		networks        Networks
		cloudProps      VMCloudProperties
		env             Environment
		registryOptions registry.ClientOptions
		agentOptions    registry.AgentOptions

		expectedAgentSettings registry.AgentSettings
		vmCID                 VMCID

		createVM CreateVM
	)

	BeforeEach(func() {
		installVMCreatorFactory(func(c client.Connector, l boshlog.Logger, vcnName string,
			subnetName string, availabilityDomain string) vm.Creator {
			return creator
		})

		connector = &clientfakes.FakeConnector{}
		uuidGen = &fakeuuid.FakeGenerator{}
		logger = boshlog.NewLogger(boshlog.LevelNone)
		registryClient = &registryfakes.FakeClient{}
		registryOptions = registry.ClientOptions{
			Protocol: "http",
			Host:     "fake-registry-host",
			Port:     25777,
			Username: "fake-registry-username",
			Password: "fake-registry-password",
		}
		agentOptions = registry.AgentOptions{
			Mbus: "http://fake-mbus",
			Blobstore: registry.BlobstoreOptions{
				Provider: "fake-blobstore-type",
			},
		}
		connector.AgentOptionsResult = agentOptions
		createVM = NewCreateVM(connector, logger, registryClient, uuidGen)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		BeforeEach(func() {

			cloudProps = VMCloudProperties{
				Shape:              "fake-BM-32",
				AvailabilityDomain: "fake-availabilityDomain",
			}

			networks = Networks{
				"fake-network-name": &Network{
					Type:    "manual",
					IP:      "10.0.0.X",
					Gateway: "fake-network-gateway",
					Netmask: "fake-network-netmask",
					DNS:     []string{"fake-network-dns"},
					DHCP:    true,
					Default: []string{"fake-network-default"},
					CloudProperties: NetworkCloudProperties{
						VcnName:    "fake-vcn",
						SubnetName: "fake-subnet1",
					},
				},
			}
			uuidGen.GeneratedUUID = "fake-uuid"
			expectedAgentSettings = registry.AgentSettings{
				AgentID: "fake-agent-id",
				Blobstore: registry.BlobstoreSettings{
					Provider: "fake-blobstore-type",
				},
				Disks: registry.DisksSettings{
					System:     "/dev/sda",
					Ephemeral:  "/dev/sda",
					Persistent: map[string]registry.PersistentSettings{},
				},
				Mbus: "http://fake-mbus",
				Networks: registry.NetworksSettings{
					"fake-network-name": registry.NetworkSetting{
						Type:          "manual",
						IP:            "10.0.0.X",
						Gateway:       "fake-network-gateway",
						Netmask:       "fake-network-netmask",
						DNS:           []string{"fake-network-dns"},
						UseDHCP:       true,
						Default:       []string{"fake-network-default"},
						Preconfigured: true,
					},
				},
				VM: registry.VMSettings{
					Name: fmt.Sprintf("bosh-%s", uuidGen.GeneratedUUID),
				},
			}
		})

		It("creates the vm", func() {
			creator = &vmfakes.FakeVMCreator{
				CreateInstanceResult: resource.NewInstance("fake-ocid", resource.Location{}),
				CreateInstanceError:  nil,
			}
			vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
			Expect(err).NotTo(HaveOccurred())
			Expect(creator.CreateInstanceCalled).To(BeTrue())
			Expect(registryClient.UpdateCalled).To(BeTrue())
			Expect(registryClient.UpdateSettings).To(Equal(expectedAgentSettings))
			Expect(vmCID).To(Equal(VMCID("fake-ocid")))
		})

		It("doesn't update registry in case of vm creation error", func() {
			creator = &vmfakes.FakeVMCreator{
				CreateInstanceResult: nil,
				CreateInstanceError:  errors.New("Error launching Instance"),
			}
			vmCID, err = createVM.Run("fake-agent-id", "fake-stemcell-id", cloudProps, networks, nil, env)
			Expect(err).To(HaveOccurred())
			Expect(creator.CreateInstanceCalled).To(BeTrue())
			Expect(registryClient.UpdateCalled).To(BeFalse())
			Expect(vmCID).To(Equal(VMCID("")))
		})
	})
})

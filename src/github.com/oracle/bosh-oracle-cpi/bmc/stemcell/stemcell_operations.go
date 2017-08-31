package stemcell

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"oracle/baremetal/core/client/compute"
)

type stemcellOperations struct {
	connector client.Connector
	logger    boshlog.Logger
}

func (so stemcellOperations) DeleteStemcell(stemcellId string) (err error) {

	so.logger.Debug(stemCellLogTag,
		"Stemcells can't be deleted. Ignoring request to delete stemmcell %s",
		stemcellId)
	return nil
}

func (so stemcellOperations) CreateStemcell(imageOCID string) (stemcellId string, err error) {

	// Verify image exists
	cs := so.connector.CoreSevice()
	p := compute.NewGetImageParams().WithImageID(imageOCID)
	if _, err = cs.Compute.GetImage(p); err != nil {
		return "", fmt.Errorf("Unable to find image %s. Reason:%s",
			imageOCID, bmc.CoreModelErrorMsg(err))
	}
	return imageOCID, nil
}

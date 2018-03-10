package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new compute API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddImageShapeCompatibilityEntry adds image shape compatibility entry

Adds a shape to the compatible shapes list for the image.
*/
func (a *Client) AddImageShapeCompatibilityEntry(params *AddImageShapeCompatibilityEntryParams) (*AddImageShapeCompatibilityEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddImageShapeCompatibilityEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddImageShapeCompatibilityEntry",
		Method:             "PUT",
		PathPattern:        "/images/{imageId}/shapes/{shapeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddImageShapeCompatibilityEntryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddImageShapeCompatibilityEntryOK), nil

}

/*
AttachVnic attaches vnic

Creates a secondary VNIC and attaches it to the specified instance.
For more information about secondary VNICs, see
[Managing Virtual Network Interface Cards (VNICs)](/Content/Network/Tasks/managingVNICs.htm).

*/
func (a *Client) AttachVnic(params *AttachVnicParams) (*AttachVnicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachVnicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AttachVnic",
		Method:             "POST",
		PathPattern:        "/vnicAttachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachVnicReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachVnicOK), nil

}

/*
AttachVolume attaches volume

Attaches the specified storage volume to the specified instance.

*/
func (a *Client) AttachVolume(params *AttachVolumeParams) (*AttachVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AttachVolume",
		Method:             "POST",
		PathPattern:        "/volumeAttachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachVolumeOK), nil

}

/*
CaptureConsoleHistory captures console history

Captures the most recent serial console data (up to a megabyte) for the
specified instance.

The `CaptureConsoleHistory` operation works with the other console history operations
as described below.

1. Use `CaptureConsoleHistory` to request the capture of up to a megabyte of the
most recent console history. This call returns a `ConsoleHistory`
object. The object will have a state of REQUESTED.
2. Wait for the capture operation to succeed by polling `GetConsoleHistory` with
the identifier of the console history metadata. The state of the
`ConsoleHistory` object will go from REQUESTED to GETTING-HISTORY and
then SUCCEEDED (or FAILED).
3. Use `GetConsoleHistoryContent` to get the actual console history data (not the
metadata).
4. Optionally, use `DeleteConsoleHistory` to delete the console history metadata
and the console history data.

*/
func (a *Client) CaptureConsoleHistory(params *CaptureConsoleHistoryParams) (*CaptureConsoleHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCaptureConsoleHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CaptureConsoleHistory",
		Method:             "POST",
		PathPattern:        "/instanceConsoleHistories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CaptureConsoleHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CaptureConsoleHistoryOK), nil

}

/*
CreateImage creates image

Creates a boot disk image for the specified instance or imports an exported image from the Oracle Bare Metal Cloud Object Storage Service.

When creating a new image, you must provide the OCID of the instance you want to use as the basis for the image, and
the OCID of the compartment containing that instance. For more information about images,
see [Managing Custom Images](/Content/Compute/Tasks/managingcustomimages.htm).

When importing an exported image from the Object Storage Service, you specify the source information
in [ImageSourceDetails](#/en/iaas/latest/requests/ImageSourceDetails).

When importing an image based on the namespace, bucket name, and object name,
use [ImageSourceViaObjectStorageTupleDetails](#/en/iaas/latest/requests/ImageSourceViaObjectStorageTupleDetails).

When importing an image based on the Object Storage Service URL, use
[ImageSourceViaObjectStorageUriDetails](#/en/iaas/latest/requests/ImageSourceViaObjectStorageUriDetails).
See [Object Storage URLs](/Content/Compute/Tasks/imageimportexport.htm#URLs) and [pre-authenticated requests](/Content/Object/Tasks/managingaccess.htm#pre-auth)
for constructing URLs for image import/export.

For more information about importing exported images, see
[Image Import/Export](/Content/Compute/Tasks/imageimportexport.htm).

You may optionally specify a *display name* for the image, which is simply a friendly name or description.
It does not have to be unique, and you can change it. See [UpdateImage](#/en/iaas/20160918/Image/UpdateImage).
Avoid entering confidential information.

*/
func (a *Client) CreateImage(params *CreateImageParams) (*CreateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateImage",
		Method:             "POST",
		PathPattern:        "/images/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateImageOK), nil

}

/*
DeleteConsoleHistory deletes console history

Deletes the specified console history metadata and the console history data.
*/
func (a *Client) DeleteConsoleHistory(params *DeleteConsoleHistoryParams) (*DeleteConsoleHistoryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConsoleHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteConsoleHistory",
		Method:             "DELETE",
		PathPattern:        "/instanceConsoleHistories/{instanceConsoleHistoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteConsoleHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteConsoleHistoryNoContent), nil

}

/*
DeleteImage deletes image

Deletes an image.
*/
func (a *Client) DeleteImage(params *DeleteImageParams) (*DeleteImageNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteImage",
		Method:             "DELETE",
		PathPattern:        "/images/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteImageNoContent), nil

}

/*
DetachVnic detaches vnic

Detaches and deletes the specified secondary VNIC.
This operation cannot be used on the instance's primary VNIC.
When you terminate an instance, all attached VNICs (primary
and secondary) are automatically detached and deleted.

*/
func (a *Client) DetachVnic(params *DetachVnicParams) (*DetachVnicNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachVnicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DetachVnic",
		Method:             "DELETE",
		PathPattern:        "/vnicAttachments/{vnicAttachmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DetachVnicReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachVnicNoContent), nil

}

/*
DetachVolume detaches volume

Detaches a storage volume from an instance. You must specify the OCID of the volume attachment.

This is an asynchronous operation; the attachment's `lifecycleState` will change to DETACHING temporarily
until the attachment is completely removed.

*/
func (a *Client) DetachVolume(params *DetachVolumeParams) (*DetachVolumeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DetachVolume",
		Method:             "DELETE",
		PathPattern:        "/volumeAttachments/{volumeAttachmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DetachVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachVolumeNoContent), nil

}

/*
ExportImage exports image

Exports the specified image to the Oracle Bare Metal Cloud Object Storage Service. You can use the Object Storage Service URL,
or the namespace, bucket name, and object name when specifying the location to export to.

For more information about exporting images, see [Image Import/Export](/Content/Compute/Tasks/imageimportexport.htm).

To perform an image export, you need write access to the Object Storage Service bucket for the image,
see [Let Users Write Objects to Object Storage Buckets](/Content/Identity/Concepts/commonpolicies.htm#Let4).

See [Object Storage URLs](/Content/Compute/Tasks/imageimportexport.htm#URLs) and [pre-authenticated requests](/Content/Object/Tasks/managingaccess.htm#pre-auth)
for constructing URLs for image import/export.

*/
func (a *Client) ExportImage(params *ExportImageParams) (*ExportImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExportImage",
		Method:             "POST",
		PathPattern:        "/images/{imageId}/actions/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExportImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExportImageOK), nil

}

/*
GetConsoleHistory gets console history

Shows the metadata for the specified console history.
See [CaptureConsoleHistory](#/en/iaas/20160918/ConsoleHistory/CaptureConsoleHistory)
for details about using the console history operations.

*/
func (a *Client) GetConsoleHistory(params *GetConsoleHistoryParams) (*GetConsoleHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsoleHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConsoleHistory",
		Method:             "GET",
		PathPattern:        "/instanceConsoleHistories/{instanceConsoleHistoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConsoleHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConsoleHistoryOK), nil

}

/*
GetConsoleHistoryContent gets console history content

Gets the actual console history data (not the metadata).
See [CaptureConsoleHistory](#/en/iaas/20160918/ConsoleHistory/CaptureConsoleHistory)
for details about using the console history operations.

*/
func (a *Client) GetConsoleHistoryContent(params *GetConsoleHistoryContentParams) (*GetConsoleHistoryContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsoleHistoryContentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConsoleHistoryContent",
		Method:             "GET",
		PathPattern:        "/instanceConsoleHistories/{instanceConsoleHistoryId}/data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConsoleHistoryContentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConsoleHistoryContentOK), nil

}

/*
GetImage gets image

Gets the specified image.
*/
func (a *Client) GetImage(params *GetImageParams) (*GetImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetImage",
		Method:             "GET",
		PathPattern:        "/images/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImageOK), nil

}

/*
GetInstance gets instance

Gets information about the specified instance.
*/
func (a *Client) GetInstance(params *GetInstanceParams) (*GetInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInstance",
		Method:             "GET",
		PathPattern:        "/instances/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInstanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInstanceOK), nil

}

/*
GetVnicAttachment gets vnic attachment

Gets the information for the specified VNIC attachment.

*/
func (a *Client) GetVnicAttachment(params *GetVnicAttachmentParams) (*GetVnicAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVnicAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVnicAttachment",
		Method:             "GET",
		PathPattern:        "/vnicAttachments/{vnicAttachmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVnicAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVnicAttachmentOK), nil

}

/*
GetVolumeAttachment gets volume attachment

Gets information about the specified volume attachment.
*/
func (a *Client) GetVolumeAttachment(params *GetVolumeAttachmentParams) (*GetVolumeAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVolumeAttachment",
		Method:             "GET",
		PathPattern:        "/volumeAttachments/{volumeAttachmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVolumeAttachmentOK), nil

}

/*
GetWindowsInstanceInitialCredentials gets windows instance initial credentials

Gets the generated credentials for the instance. Only works for Windows instances. The returned credentials
are only valid for the initial login.

*/
func (a *Client) GetWindowsInstanceInitialCredentials(params *GetWindowsInstanceInitialCredentialsParams) (*GetWindowsInstanceInitialCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWindowsInstanceInitialCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWindowsInstanceInitialCredentials",
		Method:             "GET",
		PathPattern:        "/instances/{instanceId}/initialCredentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWindowsInstanceInitialCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWindowsInstanceInitialCredentialsOK), nil

}

/*
InstanceAction instances action

Performs one of the power actions (start, stop, softreset, or reset)
on the specified instance.

**start** - power on

**stop** - power off

**softreset** - ACPI shutdown and power on

**reset** - power off and power on

Note that the **stop** state has no effect on the resources you consume.
Billing continues for instances that you stop, and related resources continue
to apply against any relevant quotas. You must terminate an instance
([TerminateInstance](#/en/iaas/20160918/Instance/TerminateInstance))
to remove its resources from billing and quotas.

*/
func (a *Client) InstanceAction(params *InstanceActionParams) (*InstanceActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInstanceActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InstanceAction",
		Method:             "POST",
		PathPattern:        "/instances/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InstanceActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InstanceActionOK), nil

}

/*
LaunchInstance launches instance

Creates a new instance in the specified compartment and the specified Availability Domain.
For general information about instances, see
[Overview of the Compute Service](/Content/Compute/Concepts/computeoverview.htm).

For information about access control and compartments, see
[Overview of the IAM Service](/Content/Identity/Concepts/overview.htm).

For information about Availability Domains, see
[Regions and Availability Domains](/Content/General/Concepts/regions.htm).
To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

All Oracle Bare Metal Cloud Services resources, including instances, get an Oracle-assigned,
unique ID called an Oracle Cloud Identifier (OCID).
When you create a resource, you can find its OCID in the response. You can
also retrieve a resource's OCID by using a List API operation
on that resource type, or by viewing the resource in the Console.

When you launch an instance, it is automatically attached to a virtual
network interface card (VNIC), called the *primary VNIC*. The VNIC
has a private IP address from the subnet's CIDR. You can either assign a
private IP address of your choice or let Oracle automatically assign one.
You can choose whether the instance has a public IP address. To retrieve the
addresses, use the [ListVnicAttachments](#/en/iaas/20160918/VnicAttachment/ListVnicAttachments)
operation to get the VNIC ID for the instance, and then call
[GetVnic](#/en/iaas/20160918/Vnic/GetVnic) with the VNIC ID.

You can later add secondary VNICs to an instance. For more information, see
[Managing Virtual Network Interface Cards (VNICs)](/Content/Network/Tasks/managingVNICs.htm).

*/
func (a *Client) LaunchInstance(params *LaunchInstanceParams) (*LaunchInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLaunchInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LaunchInstance",
		Method:             "POST",
		PathPattern:        "/instances/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LaunchInstanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LaunchInstanceOK), nil

}

/*
ListConsoleHistories lists console histories

Lists the console history metadata for the specified compartment or instance.

*/
func (a *Client) ListConsoleHistories(params *ListConsoleHistoriesParams) (*ListConsoleHistoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConsoleHistoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListConsoleHistories",
		Method:             "GET",
		PathPattern:        "/instanceConsoleHistories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListConsoleHistoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListConsoleHistoriesOK), nil

}

/*
ListImages lists images

Lists the available images in the specified compartment. For more
information about images, see
[Managing Custom Images](/Content/Compute/Tasks/managingcustomimages.htm).

*/
func (a *Client) ListImages(params *ListImagesParams) (*ListImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListImages",
		Method:             "GET",
		PathPattern:        "/images/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListImagesOK), nil

}

/*
ListInstances lists instances

Lists the instances in the specified compartment and the specified Availability Domain.
You can filter the results by specifying an instance name (the list will include all the identically-named
instances in the compartment).

*/
func (a *Client) ListInstances(params *ListInstancesParams) (*ListInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances/",
		ProducesMediaTypes: []string{"application/json", "application/x-json-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListInstancesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInstancesOK), nil

}

/*
ListShapes lists shapes

Lists the shapes that can be used to launch an instance within the specified compartment. You can
filter the list by compatibility with a specific image.

*/
func (a *Client) ListShapes(params *ListShapesParams) (*ListShapesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListShapesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListShapes",
		Method:             "GET",
		PathPattern:        "/shapes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListShapesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListShapesOK), nil

}

/*
ListVnicAttachments lists vnic attachments

Lists the VNIC attachments in the specified compartment. A VNIC attachment
resides in the same compartment as the attached instance. The list can be
filtered by instance, VNIC, or Availability Domain.

*/
func (a *Client) ListVnicAttachments(params *ListVnicAttachmentsParams) (*ListVnicAttachmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVnicAttachmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVnicAttachments",
		Method:             "GET",
		PathPattern:        "/vnicAttachments/",
		ProducesMediaTypes: []string{"application/json", "application/x-json-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVnicAttachmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVnicAttachmentsOK), nil

}

/*
ListVolumeAttachments lists volume attachments

Lists the volume attachments in the specified compartment. You can filter the
list by specifying an instance OCID, volume OCID, or both.

Currently, the only supported volume attachment type is [IScsiVolumeAttachment](#/en/iaas/20160918/IScsiVolumeAttachment/).

*/
func (a *Client) ListVolumeAttachments(params *ListVolumeAttachmentsParams) (*ListVolumeAttachmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVolumeAttachmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVolumeAttachments",
		Method:             "GET",
		PathPattern:        "/volumeAttachments/",
		ProducesMediaTypes: []string{"application/json", "application/x-json-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVolumeAttachmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVolumeAttachmentsOK), nil

}

/*
RemoveImageShapeCompatibilityEntry removes image shape compatibility entry

Removes a shape from the compatible shapes list for the image.
*/
func (a *Client) RemoveImageShapeCompatibilityEntry(params *RemoveImageShapeCompatibilityEntryParams) (*RemoveImageShapeCompatibilityEntryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveImageShapeCompatibilityEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveImageShapeCompatibilityEntry",
		Method:             "DELETE",
		PathPattern:        "/images/{imageId}/shapes/{shapeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveImageShapeCompatibilityEntryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveImageShapeCompatibilityEntryNoContent), nil

}

/*
TerminateInstance terminates instance

Terminates the specified instance. Any attached VNICs and volumes are automatically detached
when the instance terminates.

This is an asynchronous operation; the instance's `lifecycleState` will change to TERMINATING temporarily
until the instance is completely removed.

*/
func (a *Client) TerminateInstance(params *TerminateInstanceParams) (*TerminateInstanceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TerminateInstance",
		Method:             "DELETE",
		PathPattern:        "/instances/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TerminateInstanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TerminateInstanceNoContent), nil

}

/*
UpdateImage updates image

Updates the display name of the image. Avoid entering confidential information.
*/
func (a *Client) UpdateImage(params *UpdateImageParams) (*UpdateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateImage",
		Method:             "PUT",
		PathPattern:        "/images/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateImageOK), nil

}

/*
UpdateInstance updates instance

Updates the display name of the specified instance. Avoid entering confidential information.
The OCID of the instance remains the same.

*/
func (a *Client) UpdateInstance(params *UpdateInstanceParams) (*UpdateInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateInstance",
		Method:             "PUT",
		PathPattern:        "/instances/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateInstanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateInstanceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

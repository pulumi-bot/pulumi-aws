// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ID of a registered AMI for use in other
// resources.
//
// ## Example Usage
func GetAmi(ctx *pulumi.Context, args *GetAmiArgs, opts ...pulumi.InvokeOption) (*GetAmiResult, error) {
	var rv GetAmiResult
	err := ctx.Invoke("aws:index/getAmi:getAmi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAmi.
type GetAmiArgs struct {
	// Limit search to users with *explicit* launch permission on
	// the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers []string `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters []GetAmiFilter `pulumi:"filters"`
	// If more than one result is returned, use the most
	// recent AMI.
	MostRecent *bool `pulumi:"mostRecent"`
	// A regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API. This
	// filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. It is recommended to combine this with other
	// options to narrow down the list AWS returns.
	NameRegex *string `pulumi:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
	Owners []string `pulumi:"owners"`
	// Any tags assigned to the image.
	// * `tags.#.key` - The key name of the tag.
	// * `tags.#.value` - The value of the tag.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getAmi.
type GetAmiResult struct {
	// The OS architecture of the AMI (ie: `i386` or `x8664`).
	Architecture string `pulumi:"architecture"`
	// The block device mappings of the AMI.
	// * `block_device_mappings.#.device_name` - The physical name of the device.
	// * `block_device_mappings.#.ebs.delete_on_termination` - `true` if the EBS volume
	//   will be deleted on termination.
	// * `block_device_mappings.#.ebs.encrypted` - `true` if the EBS volume
	//   is encrypted.
	// * `block_device_mappings.#.ebs.iops` - `0` if the EBS volume is
	//   not a provisioned IOPS image, otherwise the supported IOPS count.
	// * `block_device_mappings.#.ebs.snapshot_id` - The ID of the snapshot.
	// * `block_device_mappings.#.ebs.volume_size` - The size of the volume, in GiB.
	// * `block_device_mappings.#.ebs.volume_type` - The volume type.
	// * `block_device_mappings.#.no_device` - Suppresses the specified device
	//   included in the block device mapping of the AMI.
	// * `block_device_mappings.#.virtual_name` - The virtual device name (for
	//   instance stores).
	BlockDeviceMappings []GetAmiBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// The date and time the image was created.
	CreationDate string `pulumi:"creationDate"`
	// The description of the AMI that was provided during image
	// creation.
	Description     string         `pulumi:"description"`
	ExecutableUsers []string       `pulumi:"executableUsers"`
	Filters         []GetAmiFilter `pulumi:"filters"`
	// The hypervisor type of the image.
	Hypervisor string `pulumi:"hypervisor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the AMI. Should be the same as the resource `id`.
	ImageId string `pulumi:"imageId"`
	// The location of the AMI.
	ImageLocation string `pulumi:"imageLocation"`
	// The AWS account alias (for example, `amazon`, `self`) or
	// the AWS account ID of the AMI owner.
	ImageOwnerAlias string `pulumi:"imageOwnerAlias"`
	// The type of image.
	ImageType string `pulumi:"imageType"`
	// The kernel associated with the image, if any. Only applicable
	// for machine images.
	KernelId   string `pulumi:"kernelId"`
	MostRecent *bool  `pulumi:"mostRecent"`
	// The name of the AMI that was provided during image creation.
	Name      string  `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	// The AWS account ID of the image owner.
	OwnerId string   `pulumi:"ownerId"`
	Owners  []string `pulumi:"owners"`
	// The value is Windows for `Windows` AMIs; otherwise blank.
	Platform string `pulumi:"platform"`
	// Any product codes associated with the AMI.
	// * `product_codes.#.product_code_id` - The product code.
	// * `product_codes.#.product_code_type` - The type of product code.
	ProductCodes []GetAmiProductCode `pulumi:"productCodes"`
	// `true` if the image has public launch permissions.
	Public bool `pulumi:"public"`
	// The RAM disk associated with the image, if any. Only applicable
	// for machine images.
	RamdiskId string `pulumi:"ramdiskId"`
	// The device name of the root device.
	RootDeviceName string `pulumi:"rootDeviceName"`
	// The type of root device (ie: `ebs` or `instance-store`).
	RootDeviceType string `pulumi:"rootDeviceType"`
	// The snapshot id associated with the root device, if any
	// (only applies to `ebs` root devices).
	RootSnapshotId string `pulumi:"rootSnapshotId"`
	// Specifies whether enhanced networking is enabled.
	SriovNetSupport string `pulumi:"sriovNetSupport"`
	// The current state of the AMI. If the state is `available`, the image
	// is successfully registered and can be used to launch an instance.
	State string `pulumi:"state"`
	// Describes a state change. Fields are `UNSET` if not available.
	// * `state_reason.code` - The reason code for the state change.
	// * `state_reason.message` - The message for the state change.
	StateReason map[string]interface{} `pulumi:"stateReason"`
	// Any tags assigned to the image.
	// * `tags.#.key` - The key name of the tag.
	// * `tags.#.value` - The value of the tag.
	Tags map[string]interface{} `pulumi:"tags"`
	// The type of virtualization of the AMI (ie: `hvm` or
	// `paravirtual`).
	VirtualizationType string `pulumi:"virtualizationType"`
}

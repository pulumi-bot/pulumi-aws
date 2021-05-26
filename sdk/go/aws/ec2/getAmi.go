// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of a registered AMI for use in other
// resources.
func LookupAmi(ctx *pulumi.Context, args *LookupAmiArgs, opts ...pulumi.InvokeOption) (*LookupAmiResult, error) {
	var rv LookupAmiResult
	err := ctx.Invoke("aws:ec2/getAmi:getAmi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAmi.
type LookupAmiArgs struct {
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
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getAmi.
type LookupAmiResult struct {
	// The OS architecture of the AMI (ie: `i386` or `x8664`).
	Architecture string `pulumi:"architecture"`
	// The ARN of the AMI.
	Arn string `pulumi:"arn"`
	// Set of objects with block device mappings of the AMI.
	BlockDeviceMappings []GetAmiBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// The date and time the image was created.
	CreationDate string `pulumi:"creationDate"`
	// The description of the AMI that was provided during image
	// creation.
	Description string `pulumi:"description"`
	// Specifies whether enhanced networking with ENA is enabled.
	EnaSupport      bool           `pulumi:"enaSupport"`
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
	// The platform details associated with the billing code of the AMI.
	PlatformDetails string `pulumi:"platformDetails"`
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
	StateReason map[string]string `pulumi:"stateReason"`
	// Any tags assigned to the image.
	// * `tags.#.key` - The key name of the tag.
	// * `tags.#.value` - The value of the tag.
	Tags map[string]string `pulumi:"tags"`
	// The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
	UsageOperation string `pulumi:"usageOperation"`
	// The type of virtualization of the AMI (ie: `hvm` or
	// `paravirtual`).
	VirtualizationType string `pulumi:"virtualizationType"`
}

func LookupAmiApply(ctx *pulumi.Context, args LookupAmiApplyInput, opts ...pulumi.InvokeOption) LookupAmiResultOutput {
	return args.ToLookupAmiApplyOutput().ApplyT(func(v LookupAmiArgs) (LookupAmiResult, error) {
		r, err := LookupAmi(ctx, &v, opts...)
		return *r, err

	}).(LookupAmiResultOutput)
}

// LookupAmiApplyInput is an input type that accepts LookupAmiApplyArgs and LookupAmiApplyOutput values.
// You can construct a concrete instance of `LookupAmiApplyInput` via:
//
//          LookupAmiApplyArgs{...}
type LookupAmiApplyInput interface {
	pulumi.Input

	ToLookupAmiApplyOutput() LookupAmiApplyOutput
	ToLookupAmiApplyOutputWithContext(context.Context) LookupAmiApplyOutput
}

// A collection of arguments for invoking getAmi.
type LookupAmiApplyArgs struct {
	// Limit search to users with *explicit* launch permission on
	// the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers pulumi.StringArrayInput `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters GetAmiFilterArrayInput `pulumi:"filters"`
	// If more than one result is returned, use the most
	// recent AMI.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// A regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API. This
	// filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. It is recommended to combine this with other
	// options to narrow down the list AWS returns.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
	Owners pulumi.StringArrayInput `pulumi:"owners"`
	// Any tags assigned to the image.
	// * `tags.#.key` - The key name of the tag.
	// * `tags.#.value` - The value of the tag.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupAmiApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAmiArgs)(nil)).Elem()
}

func (i LookupAmiApplyArgs) ToLookupAmiApplyOutput() LookupAmiApplyOutput {
	return i.ToLookupAmiApplyOutputWithContext(context.Background())
}

func (i LookupAmiApplyArgs) ToLookupAmiApplyOutputWithContext(ctx context.Context) LookupAmiApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupAmiApplyOutput)
}

// A collection of arguments for invoking getAmi.
type LookupAmiApplyOutput struct{ *pulumi.OutputState }

func (LookupAmiApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAmiArgs)(nil)).Elem()
}

func (o LookupAmiApplyOutput) ToLookupAmiApplyOutput() LookupAmiApplyOutput {
	return o
}

func (o LookupAmiApplyOutput) ToLookupAmiApplyOutputWithContext(ctx context.Context) LookupAmiApplyOutput {
	return o
}

// Limit search to users with *explicit* launch permission on
// the image. Valid items are the numeric account ID or `self`.
func (o LookupAmiApplyOutput) ExecutableUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAmiArgs) []string { return v.ExecutableUsers }).(pulumi.StringArrayOutput)
}

// One or more name/value pairs to filter off of. There are
// several valid keys, for a full reference, check out
// [describe-images in the AWS CLI reference][1].
func (o LookupAmiApplyOutput) Filters() GetAmiFilterArrayOutput {
	return o.ApplyT(func(v LookupAmiArgs) []GetAmiFilter { return v.Filters }).(GetAmiFilterArrayOutput)
}

// If more than one result is returned, use the most
// recent AMI.
func (o LookupAmiApplyOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAmiArgs) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// A regex string to apply to the AMI list returned
// by AWS. This allows more advanced filtering not supported from the AWS API. This
// filtering is done locally on what AWS returns, and could have a performance
// impact if the result is large. It is recommended to combine this with other
// options to narrow down the list AWS returns.
func (o LookupAmiApplyOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAmiArgs) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
func (o LookupAmiApplyOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAmiArgs) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

// Any tags assigned to the image.
// * `tags.#.key` - The key name of the tag.
// * `tags.#.value` - The value of the tag.
func (o LookupAmiApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAmiArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getAmi.
type LookupAmiResultOutput struct{ *pulumi.OutputState }

func (LookupAmiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAmiResult)(nil)).Elem()
}

func (o LookupAmiResultOutput) ToLookupAmiResultOutput() LookupAmiResultOutput {
	return o
}

func (o LookupAmiResultOutput) ToLookupAmiResultOutputWithContext(ctx context.Context) LookupAmiResultOutput {
	return o
}

// The OS architecture of the AMI (ie: `i386` or `x8664`).
func (o LookupAmiResultOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Architecture }).(pulumi.StringOutput)
}

// The ARN of the AMI.
func (o LookupAmiResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Set of objects with block device mappings of the AMI.
func (o LookupAmiResultOutput) BlockDeviceMappings() GetAmiBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v LookupAmiResult) []GetAmiBlockDeviceMapping { return v.BlockDeviceMappings }).(GetAmiBlockDeviceMappingArrayOutput)
}

// The date and time the image was created.
func (o LookupAmiResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The description of the AMI that was provided during image
// creation.
func (o LookupAmiResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Description }).(pulumi.StringOutput)
}

// Specifies whether enhanced networking with ENA is enabled.
func (o LookupAmiResultOutput) EnaSupport() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAmiResult) bool { return v.EnaSupport }).(pulumi.BoolOutput)
}

func (o LookupAmiResultOutput) ExecutableUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAmiResult) []string { return v.ExecutableUsers }).(pulumi.StringArrayOutput)
}

func (o LookupAmiResultOutput) Filters() GetAmiFilterArrayOutput {
	return o.ApplyT(func(v LookupAmiResult) []GetAmiFilter { return v.Filters }).(GetAmiFilterArrayOutput)
}

// The hypervisor type of the image.
func (o LookupAmiResultOutput) Hypervisor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Hypervisor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAmiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the AMI. Should be the same as the resource `id`.
func (o LookupAmiResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.ImageId }).(pulumi.StringOutput)
}

// The location of the AMI.
func (o LookupAmiResultOutput) ImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.ImageLocation }).(pulumi.StringOutput)
}

// The AWS account alias (for example, `amazon`, `self`) or
// the AWS account ID of the AMI owner.
func (o LookupAmiResultOutput) ImageOwnerAlias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.ImageOwnerAlias }).(pulumi.StringOutput)
}

// The type of image.
func (o LookupAmiResultOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.ImageType }).(pulumi.StringOutput)
}

// The kernel associated with the image, if any. Only applicable
// for machine images.
func (o LookupAmiResultOutput) KernelId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.KernelId }).(pulumi.StringOutput)
}

func (o LookupAmiResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAmiResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// The name of the AMI that was provided during image creation.
func (o LookupAmiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAmiResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAmiResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The AWS account ID of the image owner.
func (o LookupAmiResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupAmiResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAmiResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

// The value is Windows for `Windows` AMIs; otherwise blank.
func (o LookupAmiResultOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.Platform }).(pulumi.StringOutput)
}

// The platform details associated with the billing code of the AMI.
func (o LookupAmiResultOutput) PlatformDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.PlatformDetails }).(pulumi.StringOutput)
}

// Any product codes associated with the AMI.
// * `product_codes.#.product_code_id` - The product code.
// * `product_codes.#.product_code_type` - The type of product code.
func (o LookupAmiResultOutput) ProductCodes() GetAmiProductCodeArrayOutput {
	return o.ApplyT(func(v LookupAmiResult) []GetAmiProductCode { return v.ProductCodes }).(GetAmiProductCodeArrayOutput)
}

// `true` if the image has public launch permissions.
func (o LookupAmiResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAmiResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// The RAM disk associated with the image, if any. Only applicable
// for machine images.
func (o LookupAmiResultOutput) RamdiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.RamdiskId }).(pulumi.StringOutput)
}

// The device name of the root device.
func (o LookupAmiResultOutput) RootDeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.RootDeviceName }).(pulumi.StringOutput)
}

// The type of root device (ie: `ebs` or `instance-store`).
func (o LookupAmiResultOutput) RootDeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.RootDeviceType }).(pulumi.StringOutput)
}

// The snapshot id associated with the root device, if any
// (only applies to `ebs` root devices).
func (o LookupAmiResultOutput) RootSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.RootSnapshotId }).(pulumi.StringOutput)
}

// Specifies whether enhanced networking is enabled.
func (o LookupAmiResultOutput) SriovNetSupport() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.SriovNetSupport }).(pulumi.StringOutput)
}

// The current state of the AMI. If the state is `available`, the image
// is successfully registered and can be used to launch an instance.
func (o LookupAmiResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.State }).(pulumi.StringOutput)
}

// Describes a state change. Fields are `UNSET` if not available.
// * `state_reason.code` - The reason code for the state change.
// * `state_reason.message` - The message for the state change.
func (o LookupAmiResultOutput) StateReason() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAmiResult) map[string]string { return v.StateReason }).(pulumi.StringMapOutput)
}

// Any tags assigned to the image.
// * `tags.#.key` - The key name of the tag.
// * `tags.#.value` - The value of the tag.
func (o LookupAmiResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAmiResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
func (o LookupAmiResultOutput) UsageOperation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.UsageOperation }).(pulumi.StringOutput)
}

// The type of virtualization of the AMI (ie: `hvm` or
// `paravirtual`).
func (o LookupAmiResultOutput) VirtualizationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAmiResult) string { return v.VirtualizationType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAmiApplyOutput{})
	pulumi.RegisterOutputType(LookupAmiResultOutput{})
}

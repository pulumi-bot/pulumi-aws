// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networkInterfaceAttachment

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attach an Elastic network interface (ENI) resource with EC2 instance.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/network_interface_attachment.html.markdown.
type NetworkInterfaceAttachment struct {
	pulumi.CustomResourceState

	// The ENI Attachment ID.
	AttachmentId pulumi.StringOutput `pulumi:"attachmentId"`
	// Network interface index (int).
	DeviceIndex pulumi.IntOutput `pulumi:"deviceIndex"`
	// Instance ID to attach.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// ENI ID to attach.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The status of the Network Interface Attachment.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewNetworkInterfaceAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceAttachment(ctx *pulumi.Context,
	name string, args *NetworkInterfaceAttachmentArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceAttachment, error) {
	if args == nil || args.DeviceIndex == nil {
		return nil, errors.New("missing required argument 'DeviceIndex'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil {
		args = &NetworkInterfaceAttachmentArgs{}
	}
	var resource NetworkInterfaceAttachment
	err := ctx.RegisterResource("aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceAttachment gets an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceAttachmentState, opts ...pulumi.ResourceOption) (*NetworkInterfaceAttachment, error) {
	var resource NetworkInterfaceAttachment
	err := ctx.ReadResource("aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceAttachment resources.
type networkInterfaceAttachmentState struct {
	// The ENI Attachment ID.
	AttachmentId *string `pulumi:"attachmentId"`
	// Network interface index (int).
	DeviceIndex *int `pulumi:"deviceIndex"`
	// Instance ID to attach.
	InstanceId *string `pulumi:"instanceId"`
	// ENI ID to attach.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The status of the Network Interface Attachment.
	Status *string `pulumi:"status"`
}

type NetworkInterfaceAttachmentState struct {
	// The ENI Attachment ID.
	AttachmentId pulumi.StringPtrInput
	// Network interface index (int).
	DeviceIndex pulumi.IntPtrInput
	// Instance ID to attach.
	InstanceId pulumi.StringPtrInput
	// ENI ID to attach.
	NetworkInterfaceId pulumi.StringPtrInput
	// The status of the Network Interface Attachment.
	Status pulumi.StringPtrInput
}

func (NetworkInterfaceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceAttachmentState)(nil)).Elem()
}

type networkInterfaceAttachmentArgs struct {
	// Network interface index (int).
	DeviceIndex int `pulumi:"deviceIndex"`
	// Instance ID to attach.
	InstanceId string `pulumi:"instanceId"`
	// ENI ID to attach.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceAttachment resource.
type NetworkInterfaceAttachmentArgs struct {
	// Network interface index (int).
	DeviceIndex pulumi.IntInput
	// Instance ID to attach.
	InstanceId pulumi.StringInput
	// ENI ID to attach.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceAttachmentArgs)(nil)).Elem()
}


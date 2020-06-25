// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource attaches a security group to an Elastic Network Interface (ENI).
// It can be used to attach a security group to any existing ENI, be it a
// secondary ENI or one attached as the primary interface on an instance.
//
// > **NOTE on instances, interfaces, and security groups:** This provider currently
// provides the capability to assign security groups via the [`ec2.Instance`][1]
// and the [`ec2.NetworkInterface`][2] resources. Using this resource in
// conjunction with security groups provided in-line in those resources will cause
// conflicts, and will lead to spurious diffs and undefined behavior - please use
// one or the other.
//
// [1]: https://www.terraform.io/docs/providers/aws/d/instance.html
// [2]: https://www.terraform.io/docs/providers/aws/r/network_interface.html
//
// ## Output Reference
//
// There are no outputs for this resource.
type NetworkInterfaceSecurityGroupAttachment struct {
	pulumi.CustomResourceState

	// The ID of the network interface to attach to.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The ID of the security group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
}

// NewNetworkInterfaceSecurityGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceSecurityGroupAttachment(ctx *pulumi.Context,
	name string, args *NetworkInterfaceSecurityGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAttachment, error) {
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil || args.SecurityGroupId == nil {
		return nil, errors.New("missing required argument 'SecurityGroupId'")
	}
	if args == nil {
		args = &NetworkInterfaceSecurityGroupAttachmentArgs{}
	}
	var resource NetworkInterfaceSecurityGroupAttachment
	err := ctx.RegisterResource("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceSecurityGroupAttachment gets an existing NetworkInterfaceSecurityGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceSecurityGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceSecurityGroupAttachmentState, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAttachment, error) {
	var resource NetworkInterfaceSecurityGroupAttachment
	err := ctx.ReadResource("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAttachment resources.
type networkInterfaceSecurityGroupAttachmentState struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ID of the security group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

type NetworkInterfaceSecurityGroupAttachmentState struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringPtrInput
}

func (NetworkInterfaceSecurityGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAttachmentState)(nil)).Elem()
}

type networkInterfaceSecurityGroupAttachmentArgs struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the security group.
	SecurityGroupId string `pulumi:"securityGroupId"`
}

// The set of arguments for constructing a NetworkInterfaceSecurityGroupAttachment resource.
type NetworkInterfaceSecurityGroupAttachmentArgs struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId pulumi.StringInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringInput
}

func (NetworkInterfaceSecurityGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAttachmentArgs)(nil)).Elem()
}

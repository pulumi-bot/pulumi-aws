// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eipAssociation

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS EIP Association as a top level resource, to associate and
// disassociate Elastic IPs from AWS Instances and Network Interfaces.
// 
// > **NOTE:** Do not use this resource to associate an EIP to `lb.LoadBalancer` or `ec2.NatGateway` resources. Instead use the `allocationId` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
// 
// > **NOTE:** `ec2.EipAssociation` is useful in scenarios where EIPs are either
// pre-existing or distributed to customers or users and therefore cannot be changed.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/eip_association.html.markdown.
type EipAssociation struct {
	pulumi.CustomResourceState

	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolPtrOutput `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
}

// NewEipAssociation registers a new resource with the given unique name, arguments, and options.
func NewEipAssociation(ctx *pulumi.Context,
	name string, args *EipAssociationArgs, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	if args == nil {
		args = &EipAssociationArgs{}
	}
	var resource EipAssociation
	err := ctx.RegisterResource("aws:ec2/eipAssociation:EipAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEipAssociation gets an existing EipAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEipAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipAssociationState, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	var resource EipAssociation
	err := ctx.ReadResource("aws:ec2/eipAssociation:EipAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EipAssociation resources.
type eipAssociationState struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId *string `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation *bool `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp *string `pulumi:"publicIp"`
}

type EipAssociationState struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringPtrInput
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolPtrInput
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringPtrInput
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringPtrInput
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringPtrInput
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringPtrInput
}

func (EipAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAssociationState)(nil)).Elem()
}

type eipAssociationArgs struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId *string `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation *bool `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp *string `pulumi:"publicIp"`
}

// The set of arguments for constructing a EipAssociation resource.
type EipAssociationArgs struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringPtrInput
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolPtrInput
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringPtrInput
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringPtrInput
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringPtrInput
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringPtrInput
}

func (EipAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAssociationArgs)(nil)).Elem()
}


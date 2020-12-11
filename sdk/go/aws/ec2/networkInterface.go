// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic network interface (ENI) resource.
//
// ## Import
//
// Network Interfaces can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/networkInterface:NetworkInterface test eni-e5aa89a3
// ```
type NetworkInterface struct {
	pulumi.CustomResourceState

	// Block to define the attachment of the ENI. Documented below.
	Attachments NetworkInterfaceAttachmentTypeArrayOutput `pulumi:"attachments"`
	// A description for the network interface.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
	Ipv6AddressCount pulumi.IntOutput `pulumi:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
	Ipv6Addresses pulumi.StringArrayOutput `pulumi:"ipv6Addresses"`
	// The MAC address of the network interface.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	OutpostArn pulumi.StringOutput `pulumi:"outpostArn"`
	// The private DNS name of the network interface (IPv4).
	PrivateDnsName pulumi.StringOutput `pulumi:"privateDnsName"`
	PrivateIp      pulumi.StringOutput `pulumi:"privateIp"`
	// List of private IPs to assign to the ENI.
	PrivateIps pulumi.StringArrayOutput `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
	PrivateIpsCount pulumi.IntOutput `pulumi:"privateIpsCount"`
	// List of security group IDs to assign to the ENI.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Whether to enable source destination checking for the ENI. Default true.
	SourceDestCheck pulumi.BoolPtrOutput `pulumi:"sourceDestCheck"`
	// Subnet ID to create the ENI in.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource NetworkInterface
	err := ctx.RegisterResource("aws:ec2/networkInterface:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("aws:ec2/networkInterface:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
	// Block to define the attachment of the ENI. Documented below.
	Attachments []NetworkInterfaceAttachmentType `pulumi:"attachments"`
	// A description for the network interface.
	Description *string `pulumi:"description"`
	// The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
	Ipv6AddressCount *int `pulumi:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// The MAC address of the network interface.
	MacAddress *string `pulumi:"macAddress"`
	OutpostArn *string `pulumi:"outpostArn"`
	// The private DNS name of the network interface (IPv4).
	PrivateDnsName *string `pulumi:"privateDnsName"`
	PrivateIp      *string `pulumi:"privateIp"`
	// List of private IPs to assign to the ENI.
	PrivateIps []string `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	// List of security group IDs to assign to the ENI.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Whether to enable source destination checking for the ENI. Default true.
	SourceDestCheck *bool `pulumi:"sourceDestCheck"`
	// Subnet ID to create the ENI in.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type NetworkInterfaceState struct {
	// Block to define the attachment of the ENI. Documented below.
	Attachments NetworkInterfaceAttachmentTypeArrayInput
	// A description for the network interface.
	Description pulumi.StringPtrInput
	// The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
	Ipv6AddressCount pulumi.IntPtrInput
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
	Ipv6Addresses pulumi.StringArrayInput
	// The MAC address of the network interface.
	MacAddress pulumi.StringPtrInput
	OutpostArn pulumi.StringPtrInput
	// The private DNS name of the network interface (IPv4).
	PrivateDnsName pulumi.StringPtrInput
	PrivateIp      pulumi.StringPtrInput
	// List of private IPs to assign to the ENI.
	PrivateIps pulumi.StringArrayInput
	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
	PrivateIpsCount pulumi.IntPtrInput
	// List of security group IDs to assign to the ENI.
	SecurityGroups pulumi.StringArrayInput
	// Whether to enable source destination checking for the ENI. Default true.
	SourceDestCheck pulumi.BoolPtrInput
	// Subnet ID to create the ENI in.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// Block to define the attachment of the ENI. Documented below.
	Attachments []NetworkInterfaceAttachmentType `pulumi:"attachments"`
	// A description for the network interface.
	Description *string `pulumi:"description"`
	// The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
	Ipv6AddressCount *int `pulumi:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	PrivateIp     *string  `pulumi:"privateIp"`
	// List of private IPs to assign to the ENI.
	PrivateIps []string `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	// List of security group IDs to assign to the ENI.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Whether to enable source destination checking for the ENI. Default true.
	SourceDestCheck *bool `pulumi:"sourceDestCheck"`
	// Subnet ID to create the ENI in.
	SubnetId string `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Block to define the attachment of the ENI. Documented below.
	Attachments NetworkInterfaceAttachmentTypeArrayInput
	// A description for the network interface.
	Description pulumi.StringPtrInput
	// The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
	Ipv6AddressCount pulumi.IntPtrInput
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
	Ipv6Addresses pulumi.StringArrayInput
	PrivateIp     pulumi.StringPtrInput
	// List of private IPs to assign to the ENI.
	PrivateIps pulumi.StringArrayInput
	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
	PrivateIpsCount pulumi.IntPtrInput
	// List of security group IDs to assign to the ENI.
	SecurityGroups pulumi.StringArrayInput
	// Whether to enable source destination checking for the ENI. Default true.
	SourceDestCheck pulumi.BoolPtrInput
	// Subnet ID to create the ENI in.
	SubnetId pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

type NetworkInterfacePtrInput interface {
	pulumi.Input

	ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput
	ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput
}

func (NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

func (i NetworkInterface) ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput {
	return i.ToNetworkInterfacePtrOutputWithContext(context.Background())
}

func (i NetworkInterface) ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePtrOutput)
}

type NetworkInterfaceOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceOutput)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

type NetworkInterfacePtrOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfacePtrOutput) ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput {
	return o
}

func (o NetworkInterfacePtrOutput) ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfacePtrOutput{})
}

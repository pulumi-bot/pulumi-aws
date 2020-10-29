// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC DHCP Options resource.
//
// ## Remarks
//
// * Notice that all arguments are optional but you have to specify at least one argument.
// * `domainNameServers`, `netbiosNameServers`, `ntpServers` are limited by AWS to maximum four servers only.
// * To actually use the DHCP Options Set you need to associate it to a VPC using `ec2.VpcDhcpOptionsAssociation`.
// * If you delete a DHCP Options Set, all VPCs using it will be associated to AWS's `default` DHCP Option Set.
// * In most cases unless you're configuring your own DNS you'll want to set `domainNameServers` to `AmazonProvidedDNS`.
type VpcDhcpOptions struct {
	pulumi.CustomResourceState

	// The ARN of the DHCP Options Set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers pulumi.StringArrayOutput `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayOutput `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrOutput `pulumi:"netbiosNodeType"`
	// List of NTP servers to configure.
	NtpServers pulumi.StringArrayOutput `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVpcDhcpOptions registers a new resource with the given unique name, arguments, and options.
func NewVpcDhcpOptions(ctx *pulumi.Context,
	name string, args *VpcDhcpOptionsArgs, opts ...pulumi.ResourceOption) (*VpcDhcpOptions, error) {
	if args == nil {
		args = &VpcDhcpOptionsArgs{}
	}
	var resource VpcDhcpOptions
	err := ctx.RegisterResource("aws:ec2/vpcDhcpOptions:VpcDhcpOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcDhcpOptions gets an existing VpcDhcpOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcDhcpOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcDhcpOptionsState, opts ...pulumi.ResourceOption) (*VpcDhcpOptions, error) {
	var resource VpcDhcpOptions
	err := ctx.ReadResource("aws:ec2/vpcDhcpOptions:VpcDhcpOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcDhcpOptions resources.
type vpcDhcpOptionsState struct {
	// The ARN of the DHCP Options Set.
	Arn *string `pulumi:"arn"`
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName *string `pulumi:"domainName"`
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers []string `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers []string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	// List of NTP servers to configure.
	NtpServers []string `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VpcDhcpOptionsState struct {
	// The ARN of the DHCP Options Set.
	Arn pulumi.StringPtrInput
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName pulumi.StringPtrInput
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers pulumi.StringArrayInput
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	// List of NTP servers to configure.
	NtpServers pulumi.StringArrayInput
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VpcDhcpOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsState)(nil)).Elem()
}

type vpcDhcpOptionsArgs struct {
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName *string `pulumi:"domainName"`
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers []string `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers []string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	// List of NTP servers to configure.
	NtpServers []string `pulumi:"ntpServers"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcDhcpOptions resource.
type VpcDhcpOptionsArgs struct {
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName pulumi.StringPtrInput
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers pulumi.StringArrayInput
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	// List of NTP servers to configure.
	NtpServers pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VpcDhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsArgs)(nil)).Elem()
}

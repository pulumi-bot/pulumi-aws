// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the [default AWS DHCP Options Set](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html#AmazonDNS)
// in the current region.
// 
// Each AWS region comes with a default set of DHCP options.
// **This is an advanced resource**, and has special caveats to be aware of when
// using it. Please read this document in its entirety before using this resource.
// 
// The `ec2.DefaultVpcDhcpOptions` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_vpc_dhcp_options.html.markdown.
type DefaultVpcDhcpOptions struct {
	pulumi.CustomResourceState

	DomainName pulumi.StringOutput `pulumi:"domainName"`
	DomainNameServers pulumi.StringOutput `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayOutput `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrOutput `pulumi:"netbiosNodeType"`
	NtpServers pulumi.StringOutput `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDefaultVpcDhcpOptions registers a new resource with the given unique name, arguments, and options.
func NewDefaultVpcDhcpOptions(ctx *pulumi.Context,
	name string, args *DefaultVpcDhcpOptionsArgs, opts ...pulumi.ResourceOption) (*DefaultVpcDhcpOptions, error) {
	if args == nil {
		args = &DefaultVpcDhcpOptionsArgs{}
	}
	var resource DefaultVpcDhcpOptions
	err := ctx.RegisterResource("aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultVpcDhcpOptions gets an existing DefaultVpcDhcpOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultVpcDhcpOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultVpcDhcpOptionsState, opts ...pulumi.ResourceOption) (*DefaultVpcDhcpOptions, error) {
	var resource DefaultVpcDhcpOptions
	err := ctx.ReadResource("aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultVpcDhcpOptions resources.
type defaultVpcDhcpOptionsState struct {
	DomainName *string `pulumi:"domainName"`
	DomainNameServers *string `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers []string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	NtpServers *string `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId *string `pulumi:"ownerId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type DefaultVpcDhcpOptionsState struct {
	DomainName pulumi.StringPtrInput
	DomainNameServers pulumi.StringPtrInput
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	NtpServers pulumi.StringPtrInput
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (DefaultVpcDhcpOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcDhcpOptionsState)(nil)).Elem()
}

type defaultVpcDhcpOptionsArgs struct {
	// List of NETBIOS name servers.
	NetbiosNameServers []string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultVpcDhcpOptions resource.
type DefaultVpcDhcpOptionsArgs struct {
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (DefaultVpcDhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcDhcpOptionsArgs)(nil)).Elem()
}


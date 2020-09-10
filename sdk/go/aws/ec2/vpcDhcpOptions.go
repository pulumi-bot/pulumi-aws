// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VpcDhcpOptions struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput      `pulumi:"arn"`
	DomainName         pulumi.StringPtrOutput   `pulumi:"domainName"`
	DomainNameServers  pulumi.StringArrayOutput `pulumi:"domainNameServers"`
	NetbiosNameServers pulumi.StringArrayOutput `pulumi:"netbiosNameServers"`
	NetbiosNodeType    pulumi.StringPtrOutput   `pulumi:"netbiosNodeType"`
	NtpServers         pulumi.StringArrayOutput `pulumi:"ntpServers"`
	OwnerId            pulumi.StringOutput      `pulumi:"ownerId"`
	Tags               pulumi.StringMapOutput   `pulumi:"tags"`
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
	Arn                *string           `pulumi:"arn"`
	DomainName         *string           `pulumi:"domainName"`
	DomainNameServers  []string          `pulumi:"domainNameServers"`
	NetbiosNameServers []string          `pulumi:"netbiosNameServers"`
	NetbiosNodeType    *string           `pulumi:"netbiosNodeType"`
	NtpServers         []string          `pulumi:"ntpServers"`
	OwnerId            *string           `pulumi:"ownerId"`
	Tags               map[string]string `pulumi:"tags"`
}

type VpcDhcpOptionsState struct {
	Arn                pulumi.StringPtrInput
	DomainName         pulumi.StringPtrInput
	DomainNameServers  pulumi.StringArrayInput
	NetbiosNameServers pulumi.StringArrayInput
	NetbiosNodeType    pulumi.StringPtrInput
	NtpServers         pulumi.StringArrayInput
	OwnerId            pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (VpcDhcpOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsState)(nil)).Elem()
}

type vpcDhcpOptionsArgs struct {
	DomainName         *string           `pulumi:"domainName"`
	DomainNameServers  []string          `pulumi:"domainNameServers"`
	NetbiosNameServers []string          `pulumi:"netbiosNameServers"`
	NetbiosNodeType    *string           `pulumi:"netbiosNodeType"`
	NtpServers         []string          `pulumi:"ntpServers"`
	Tags               map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcDhcpOptions resource.
type VpcDhcpOptionsArgs struct {
	DomainName         pulumi.StringPtrInput
	DomainNameServers  pulumi.StringArrayInput
	NetbiosNameServers pulumi.StringArrayInput
	NetbiosNodeType    pulumi.StringPtrInput
	NtpServers         pulumi.StringArrayInput
	Tags               pulumi.StringMapInput
}

func (VpcDhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsArgs)(nil)).Elem()
}

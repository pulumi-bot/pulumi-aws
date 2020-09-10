// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type DefaultNetworkAcl struct {
	pulumi.CustomResourceState

	Arn                 pulumi.StringOutput                 `pulumi:"arn"`
	DefaultNetworkAclId pulumi.StringOutput                 `pulumi:"defaultNetworkAclId"`
	Egress              DefaultNetworkAclEgressArrayOutput  `pulumi:"egress"`
	Ingress             DefaultNetworkAclIngressArrayOutput `pulumi:"ingress"`
	OwnerId             pulumi.StringOutput                 `pulumi:"ownerId"`
	SubnetIds           pulumi.StringArrayOutput            `pulumi:"subnetIds"`
	Tags                pulumi.StringMapOutput              `pulumi:"tags"`
	VpcId               pulumi.StringOutput                 `pulumi:"vpcId"`
}

// NewDefaultNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewDefaultNetworkAcl(ctx *pulumi.Context,
	name string, args *DefaultNetworkAclArgs, opts ...pulumi.ResourceOption) (*DefaultNetworkAcl, error) {
	if args == nil || args.DefaultNetworkAclId == nil {
		return nil, errors.New("missing required argument 'DefaultNetworkAclId'")
	}
	if args == nil {
		args = &DefaultNetworkAclArgs{}
	}
	var resource DefaultNetworkAcl
	err := ctx.RegisterResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultNetworkAcl gets an existing DefaultNetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultNetworkAclState, opts ...pulumi.ResourceOption) (*DefaultNetworkAcl, error) {
	var resource DefaultNetworkAcl
	err := ctx.ReadResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultNetworkAcl resources.
type defaultNetworkAclState struct {
	Arn                 *string                    `pulumi:"arn"`
	DefaultNetworkAclId *string                    `pulumi:"defaultNetworkAclId"`
	Egress              []DefaultNetworkAclEgress  `pulumi:"egress"`
	Ingress             []DefaultNetworkAclIngress `pulumi:"ingress"`
	OwnerId             *string                    `pulumi:"ownerId"`
	SubnetIds           []string                   `pulumi:"subnetIds"`
	Tags                map[string]string          `pulumi:"tags"`
	VpcId               *string                    `pulumi:"vpcId"`
}

type DefaultNetworkAclState struct {
	Arn                 pulumi.StringPtrInput
	DefaultNetworkAclId pulumi.StringPtrInput
	Egress              DefaultNetworkAclEgressArrayInput
	Ingress             DefaultNetworkAclIngressArrayInput
	OwnerId             pulumi.StringPtrInput
	SubnetIds           pulumi.StringArrayInput
	Tags                pulumi.StringMapInput
	VpcId               pulumi.StringPtrInput
}

func (DefaultNetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNetworkAclState)(nil)).Elem()
}

type defaultNetworkAclArgs struct {
	DefaultNetworkAclId string                     `pulumi:"defaultNetworkAclId"`
	Egress              []DefaultNetworkAclEgress  `pulumi:"egress"`
	Ingress             []DefaultNetworkAclIngress `pulumi:"ingress"`
	SubnetIds           []string                   `pulumi:"subnetIds"`
	Tags                map[string]string          `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultNetworkAcl resource.
type DefaultNetworkAclArgs struct {
	DefaultNetworkAclId pulumi.StringInput
	Egress              DefaultNetworkAclEgressArrayInput
	Ingress             DefaultNetworkAclIngressArrayInput
	SubnetIds           pulumi.StringArrayInput
	Tags                pulumi.StringMapInput
}

func (DefaultNetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNetworkAclArgs)(nil)).Elem()
}

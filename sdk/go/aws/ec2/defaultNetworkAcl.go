// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the default AWS Network ACL. VPC Only.
//
// Each VPC created in AWS comes with a Default Network ACL that can be managed, but not
// destroyed. **This is an advanced resource**, and has special caveats to be aware
// of when using it. Please read this document in its entirety before using this
// resource.
//
// The `ec2.DefaultNetworkAcl` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead attempts to "adopt" it
// into management. We can do this because each VPC created has a Default Network
// ACL that cannot be destroyed, and is created with a known set of default rules.
//
// When this provider first adopts the Default Network ACL, it **immediately removes all
// rules in the ACL**. It then proceeds to create any rules specified in the
// configuration. This step is required so that only the rules specified in the
// configuration are created.
//
// This resource treats its inline rules as absolute; only the rules defined
// inline are created, and any additions/removals external to this resource will
// result in diffs being shown. For these reasons, this resource is incompatible with the
// `ec2.NetworkAclRule` resource.
//
// For more information about Network ACLs, see the AWS Documentation on
// [Network ACLs][aws-network-acls].
type DefaultNetworkAcl struct {
	pulumi.CustomResourceState

	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`
	// Specifies an egress rule. Parameters defined below.
	Egress DefaultNetworkAclEgressArrayOutput `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress DefaultNetworkAclIngressArrayOutput `pulumi:"ingress"`
	// The ID of the AWS account that owns the Default Network ACL
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the associated VPC
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
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
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId *string `pulumi:"defaultNetworkAclId"`
	// Specifies an egress rule. Parameters defined below.
	Egress []DefaultNetworkAclEgress `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress []DefaultNetworkAclIngress `pulumi:"ingress"`
	// The ID of the AWS account that owns the Default Network ACL
	OwnerId *string `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the associated VPC
	VpcId *string `pulumi:"vpcId"`
}

type DefaultNetworkAclState struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringPtrInput
	// Specifies an egress rule. Parameters defined below.
	Egress DefaultNetworkAclEgressArrayInput
	// Specifies an ingress rule. Parameters defined below.
	Ingress DefaultNetworkAclIngressArrayInput
	// The ID of the AWS account that owns the Default Network ACL
	OwnerId pulumi.StringPtrInput
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the associated VPC
	VpcId pulumi.StringPtrInput
}

func (DefaultNetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNetworkAclState)(nil)).Elem()
}

type defaultNetworkAclArgs struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId string `pulumi:"defaultNetworkAclId"`
	// Specifies an egress rule. Parameters defined below.
	Egress []DefaultNetworkAclEgressArgs `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress []DefaultNetworkAclIngressArgs `pulumi:"ingress"`
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultNetworkAcl resource.
type DefaultNetworkAclArgs struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringInput
	// Specifies an egress rule. Parameters defined below.
	Egress DefaultNetworkAclEgressArgsArrayInput
	// Specifies an ingress rule. Parameters defined below.
	Ingress DefaultNetworkAclIngressArgsArrayInput
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (DefaultNetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNetworkAclArgs)(nil)).Elem()
}

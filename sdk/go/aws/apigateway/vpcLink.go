// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VpcLink struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput    `pulumi:"arn"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	TargetArn   pulumi.StringOutput    `pulumi:"targetArn"`
}

// NewVpcLink registers a new resource with the given unique name, arguments, and options.
func NewVpcLink(ctx *pulumi.Context,
	name string, args *VpcLinkArgs, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	if args == nil || args.TargetArn == nil {
		return nil, errors.New("missing required argument 'TargetArn'")
	}
	if args == nil {
		args = &VpcLinkArgs{}
	}
	var resource VpcLink
	err := ctx.RegisterResource("aws:apigateway/vpcLink:VpcLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcLink gets an existing VpcLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcLinkState, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	var resource VpcLink
	err := ctx.ReadResource("aws:apigateway/vpcLink:VpcLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcLink resources.
type vpcLinkState struct {
	Arn         *string           `pulumi:"arn"`
	Description *string           `pulumi:"description"`
	Name        *string           `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
	TargetArn   *string           `pulumi:"targetArn"`
}

type VpcLinkState struct {
	Arn         pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Tags        pulumi.StringMapInput
	TargetArn   pulumi.StringPtrInput
}

func (VpcLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkState)(nil)).Elem()
}

type vpcLinkArgs struct {
	Description *string           `pulumi:"description"`
	Name        *string           `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
	TargetArn   string            `pulumi:"targetArn"`
}

// The set of arguments for constructing a VpcLink resource.
type VpcLinkArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Tags        pulumi.StringMapInput
	TargetArn   pulumi.StringInput
}

func (VpcLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkArgs)(nil)).Elem()
}

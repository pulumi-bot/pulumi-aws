// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type WebAclAssociation struct {
	pulumi.CustomResourceState

	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	WebAclArn   pulumi.StringOutput `pulumi:"webAclArn"`
}

// NewWebAclAssociation registers a new resource with the given unique name, arguments, and options.
func NewWebAclAssociation(ctx *pulumi.Context,
	name string, args *WebAclAssociationArgs, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	if args == nil || args.ResourceArn == nil {
		return nil, errors.New("missing required argument 'ResourceArn'")
	}
	if args == nil || args.WebAclArn == nil {
		return nil, errors.New("missing required argument 'WebAclArn'")
	}
	if args == nil {
		args = &WebAclAssociationArgs{}
	}
	var resource WebAclAssociation
	err := ctx.RegisterResource("aws:wafv2/webAclAssociation:WebAclAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAclAssociation gets an existing WebAclAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAclAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclAssociationState, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	var resource WebAclAssociation
	err := ctx.ReadResource("aws:wafv2/webAclAssociation:WebAclAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAclAssociation resources.
type webAclAssociationState struct {
	ResourceArn *string `pulumi:"resourceArn"`
	WebAclArn   *string `pulumi:"webAclArn"`
}

type WebAclAssociationState struct {
	ResourceArn pulumi.StringPtrInput
	WebAclArn   pulumi.StringPtrInput
}

func (WebAclAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationState)(nil)).Elem()
}

type webAclAssociationArgs struct {
	ResourceArn string `pulumi:"resourceArn"`
	WebAclArn   string `pulumi:"webAclArn"`
}

// The set of arguments for constructing a WebAclAssociation resource.
type WebAclAssociationArgs struct {
	ResourceArn pulumi.StringInput
	WebAclArn   pulumi.StringInput
}

func (WebAclAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationArgs)(nil)).Elem()
}

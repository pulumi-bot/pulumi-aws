// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Group struct {
	pulumi.CustomResourceState

	Arn           pulumi.StringOutput      `pulumi:"arn"`
	Description   pulumi.StringPtrOutput   `pulumi:"description"`
	Name          pulumi.StringOutput      `pulumi:"name"`
	ResourceQuery GroupResourceQueryOutput `pulumi:"resourceQuery"`
	Tags          pulumi.StringMapOutput   `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.ResourceQuery == nil {
		return nil, errors.New("missing required argument 'ResourceQuery'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("aws:resourcegroups/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws:resourcegroups/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	Arn           *string             `pulumi:"arn"`
	Description   *string             `pulumi:"description"`
	Name          *string             `pulumi:"name"`
	ResourceQuery *GroupResourceQuery `pulumi:"resourceQuery"`
	Tags          map[string]string   `pulumi:"tags"`
}

type GroupState struct {
	Arn           pulumi.StringPtrInput
	Description   pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	ResourceQuery GroupResourceQueryPtrInput
	Tags          pulumi.StringMapInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	Description   *string            `pulumi:"description"`
	Name          *string            `pulumi:"name"`
	ResourceQuery GroupResourceQuery `pulumi:"resourceQuery"`
	Tags          map[string]string  `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	Description   pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	ResourceQuery GroupResourceQueryInput
	Tags          pulumi.StringMapInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

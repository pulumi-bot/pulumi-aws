// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// `aws_ec2_tag` can be imported by using the EC2 resource identifier and key, separated by a comma (`,`), e.g.
//
// ```sh
//  $ pulumi import aws:ec2/tag:Tag example tgw-attach-1234567890abcdef,Name
// ```
type Tag struct {
	pulumi.CustomResourceState

	// The tag name.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the EC2 resource to manage the tag for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The value of the tag.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource Tag
	err := ctx.RegisterResource("aws:ec2/tag:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("aws:ec2/tag:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
	// The tag name.
	Key *string `pulumi:"key"`
	// The ID of the EC2 resource to manage the tag for.
	ResourceId *string `pulumi:"resourceId"`
	// The value of the tag.
	Value *string `pulumi:"value"`
}

type TagState struct {
	// The tag name.
	Key pulumi.StringPtrInput
	// The ID of the EC2 resource to manage the tag for.
	ResourceId pulumi.StringPtrInput
	// The value of the tag.
	Value pulumi.StringPtrInput
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// The tag name.
	Key string `pulumi:"key"`
	// The ID of the EC2 resource to manage the tag for.
	ResourceId string `pulumi:"resourceId"`
	// The value of the tag.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// The tag name.
	Key pulumi.StringInput
	// The ID of the EC2 resource to manage the tag for.
	ResourceId pulumi.StringInput
	// The value of the tag.
	Value pulumi.StringInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(ctx context.Context) TagOutput
}

type TagPtrInput interface {
	pulumi.Input

	ToTagPtrOutput() TagPtrOutput
	ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput
}

func (Tag) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (i Tag) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i Tag) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

func (i Tag) ToTagPtrOutput() TagPtrOutput {
	return i.ToTagPtrOutputWithContext(context.Background())
}

func (i Tag) ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPtrOutput)
}

type TagOutput struct {
	*pulumi.OutputState
}

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOutput)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

type TagPtrOutput struct {
	*pulumi.OutputState
}

func (TagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (o TagPtrOutput) ToTagPtrOutput() TagPtrOutput {
	return o
}

func (o TagPtrOutput) ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagPtrOutput{})
}

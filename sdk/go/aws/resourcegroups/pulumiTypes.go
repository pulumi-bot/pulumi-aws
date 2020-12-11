// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GroupResourceQuery struct {
	// The resource query as a JSON string.
	Query string `pulumi:"query"`
	// The type of the resource query. Defaults to `TAG_FILTERS_1_0`.
	Type *string `pulumi:"type"`
}

// GroupResourceQueryInput is an input type that accepts GroupResourceQueryArgs and GroupResourceQueryOutput values.
// You can construct a concrete instance of `GroupResourceQueryInput` via:
//
//          GroupResourceQueryArgs{...}
type GroupResourceQueryInput interface {
	pulumi.Input

	ToGroupResourceQueryOutput() GroupResourceQueryOutput
	ToGroupResourceQueryOutputWithContext(context.Context) GroupResourceQueryOutput
}

type GroupResourceQueryArgs struct {
	// The resource query as a JSON string.
	Query pulumi.StringInput `pulumi:"query"`
	// The type of the resource query. Defaults to `TAG_FILTERS_1_0`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GroupResourceQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupResourceQuery)(nil)).Elem()
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryOutput() GroupResourceQueryOutput {
	return i.ToGroupResourceQueryOutputWithContext(context.Background())
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryOutputWithContext(ctx context.Context) GroupResourceQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupResourceQueryOutput)
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return i.ToGroupResourceQueryPtrOutputWithContext(context.Background())
}

func (i GroupResourceQueryArgs) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupResourceQueryOutput).ToGroupResourceQueryPtrOutput()
}

// GroupResourceQueryPtrInput is an input type that accepts GroupResourceQueryArgs, GroupResourceQueryPtr and GroupResourceQueryPtrOutput values.
// You can construct a concrete instance of `GroupResourceQueryPtrInput` via:
//
//          GroupResourceQueryArgs{...}
//
//  or:
//
//          nil
type GroupResourceQueryPtrInput interface {
	pulumi.Input

	ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput
	ToGroupResourceQueryPtrOutputWithContext(context.Context) GroupResourceQueryPtrOutput
}

type groupResourceQueryPtrType GroupResourceQueryArgs

func GroupResourceQueryPtr(v *GroupResourceQueryArgs) GroupResourceQueryPtrInput {
	return (*groupResourceQueryPtrType)(v)
}

func (*groupResourceQueryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupResourceQuery)(nil)).Elem()
}

func (i *groupResourceQueryPtrType) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return i.ToGroupResourceQueryPtrOutputWithContext(context.Background())
}

func (i *groupResourceQueryPtrType) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupResourceQueryOutput).ToGroupResourceQueryPtrOutput()
}

type GroupResourceQueryOutput struct{ *pulumi.OutputState }

func (GroupResourceQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupResourceQuery)(nil)).Elem()
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryOutput() GroupResourceQueryOutput {
	return o
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryOutputWithContext(ctx context.Context) GroupResourceQueryOutput {
	return o
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return o.ToGroupResourceQueryPtrOutputWithContext(context.Background())
}

func (o GroupResourceQueryOutput) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return o.ApplyT(func(v GroupResourceQuery) *GroupResourceQuery {
		return &v
	}).(GroupResourceQueryPtrOutput)
}

// The resource query as a JSON string.
func (o GroupResourceQueryOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v GroupResourceQuery) string { return v.Query }).(pulumi.StringOutput)
}

// The type of the resource query. Defaults to `TAG_FILTERS_1_0`.
func (o GroupResourceQueryOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupResourceQuery) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GroupResourceQueryPtrOutput struct{ *pulumi.OutputState }

func (GroupResourceQueryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupResourceQuery)(nil)).Elem()
}

func (o GroupResourceQueryPtrOutput) ToGroupResourceQueryPtrOutput() GroupResourceQueryPtrOutput {
	return o
}

func (o GroupResourceQueryPtrOutput) ToGroupResourceQueryPtrOutputWithContext(ctx context.Context) GroupResourceQueryPtrOutput {
	return o
}

func (o GroupResourceQueryPtrOutput) Elem() GroupResourceQueryOutput {
	return o.ApplyT(func(v *GroupResourceQuery) GroupResourceQuery { return *v }).(GroupResourceQueryOutput)
}

// The resource query as a JSON string.
func (o GroupResourceQueryPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupResourceQuery) *string {
		if v == nil {
			return nil
		}
		return &v.Query
	}).(pulumi.StringPtrOutput)
}

// The type of the resource query. Defaults to `TAG_FILTERS_1_0`.
func (o GroupResourceQueryPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupResourceQuery) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupResourceQueryOutput{})
	pulumi.RegisterOutputType(GroupResourceQueryPtrOutput{})
}

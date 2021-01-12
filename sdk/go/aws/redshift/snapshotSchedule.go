// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redshift.NewSnapshotSchedule(ctx, "_default", &redshift.SnapshotScheduleArgs{
// 			Definitions: pulumi.StringArray{
// 				pulumi.String("rate(12 hours)"),
// 			},
// 			Identifier: pulumi.String("tf-redshift-snapshot-schedule"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Redshift Snapshot Schedule can be imported using the `identifier`, e.g.
//
// ```sh
//  $ pulumi import aws:redshift/snapshotSchedule:SnapshotSchedule default tf-redshift-snapshot-schedule
// ```
type SnapshotSchedule struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions pulumi.StringArrayOutput `pulumi:"definitions"`
	// The description of the snapshot schedule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringOutput `pulumi:"identifierPrefix"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSnapshotSchedule registers a new resource with the given unique name, arguments, and options.
func NewSnapshotSchedule(ctx *pulumi.Context,
	name string, args *SnapshotScheduleArgs, opts ...pulumi.ResourceOption) (*SnapshotSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definitions == nil {
		return nil, errors.New("invalid value for required argument 'Definitions'")
	}
	var resource SnapshotSchedule
	err := ctx.RegisterResource("aws:redshift/snapshotSchedule:SnapshotSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotSchedule gets an existing SnapshotSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotScheduleState, opts ...pulumi.ResourceOption) (*SnapshotSchedule, error) {
	var resource SnapshotSchedule
	err := ctx.ReadResource("aws:redshift/snapshotSchedule:SnapshotSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotSchedule resources.
type snapshotScheduleState struct {
	Arn *string `pulumi:"arn"`
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions []string `pulumi:"definitions"`
	// The description of the snapshot schedule.
	Description *string `pulumi:"description"`
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SnapshotScheduleState struct {
	Arn pulumi.StringPtrInput
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions pulumi.StringArrayInput
	// The description of the snapshot schedule.
	Description pulumi.StringPtrInput
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy pulumi.BoolPtrInput
	// The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SnapshotScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotScheduleState)(nil)).Elem()
}

type snapshotScheduleArgs struct {
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions []string `pulumi:"definitions"`
	// The description of the snapshot schedule.
	Description *string `pulumi:"description"`
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SnapshotSchedule resource.
type SnapshotScheduleArgs struct {
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions pulumi.StringArrayInput
	// The description of the snapshot schedule.
	Description pulumi.StringPtrInput
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy pulumi.BoolPtrInput
	// The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SnapshotScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotScheduleArgs)(nil)).Elem()
}

type SnapshotScheduleInput interface {
	pulumi.Input

	ToSnapshotScheduleOutput() SnapshotScheduleOutput
	ToSnapshotScheduleOutputWithContext(ctx context.Context) SnapshotScheduleOutput
}

func (*SnapshotSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSchedule)(nil))
}

func (i *SnapshotSchedule) ToSnapshotScheduleOutput() SnapshotScheduleOutput {
	return i.ToSnapshotScheduleOutputWithContext(context.Background())
}

func (i *SnapshotSchedule) ToSnapshotScheduleOutputWithContext(ctx context.Context) SnapshotScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotScheduleOutput)
}

func (i *SnapshotSchedule) ToSnapshotSchedulePtrOutput() SnapshotSchedulePtrOutput {
	return i.ToSnapshotSchedulePtrOutputWithContext(context.Background())
}

func (i *SnapshotSchedule) ToSnapshotSchedulePtrOutputWithContext(ctx context.Context) SnapshotSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSchedulePtrOutput)
}

type SnapshotSchedulePtrInput interface {
	pulumi.Input

	ToSnapshotSchedulePtrOutput() SnapshotSchedulePtrOutput
	ToSnapshotSchedulePtrOutputWithContext(ctx context.Context) SnapshotSchedulePtrOutput
}

type snapshotSchedulePtrType SnapshotScheduleArgs

func (*snapshotSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSchedule)(nil))
}

func (i *snapshotSchedulePtrType) ToSnapshotSchedulePtrOutput() SnapshotSchedulePtrOutput {
	return i.ToSnapshotSchedulePtrOutputWithContext(context.Background())
}

func (i *snapshotSchedulePtrType) ToSnapshotSchedulePtrOutputWithContext(ctx context.Context) SnapshotSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotScheduleOutput).ToSnapshotSchedulePtrOutput()
}

// SnapshotScheduleArrayInput is an input type that accepts SnapshotScheduleArray and SnapshotScheduleArrayOutput values.
// You can construct a concrete instance of `SnapshotScheduleArrayInput` via:
//
//          SnapshotScheduleArray{ SnapshotScheduleArgs{...} }
type SnapshotScheduleArrayInput interface {
	pulumi.Input

	ToSnapshotScheduleArrayOutput() SnapshotScheduleArrayOutput
	ToSnapshotScheduleArrayOutputWithContext(context.Context) SnapshotScheduleArrayOutput
}

type SnapshotScheduleArray []SnapshotScheduleInput

func (SnapshotScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SnapshotSchedule)(nil))
}

func (i SnapshotScheduleArray) ToSnapshotScheduleArrayOutput() SnapshotScheduleArrayOutput {
	return i.ToSnapshotScheduleArrayOutputWithContext(context.Background())
}

func (i SnapshotScheduleArray) ToSnapshotScheduleArrayOutputWithContext(ctx context.Context) SnapshotScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotScheduleArrayOutput)
}

// SnapshotScheduleMapInput is an input type that accepts SnapshotScheduleMap and SnapshotScheduleMapOutput values.
// You can construct a concrete instance of `SnapshotScheduleMapInput` via:
//
//          SnapshotScheduleMap{ "key": SnapshotScheduleArgs{...} }
type SnapshotScheduleMapInput interface {
	pulumi.Input

	ToSnapshotScheduleMapOutput() SnapshotScheduleMapOutput
	ToSnapshotScheduleMapOutputWithContext(context.Context) SnapshotScheduleMapOutput
}

type SnapshotScheduleMap map[string]SnapshotScheduleInput

func (SnapshotScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SnapshotSchedule)(nil))
}

func (i SnapshotScheduleMap) ToSnapshotScheduleMapOutput() SnapshotScheduleMapOutput {
	return i.ToSnapshotScheduleMapOutputWithContext(context.Background())
}

func (i SnapshotScheduleMap) ToSnapshotScheduleMapOutputWithContext(ctx context.Context) SnapshotScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotScheduleMapOutput)
}

type SnapshotScheduleOutput struct {
	*pulumi.OutputState
}

func (SnapshotScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSchedule)(nil))
}

func (o SnapshotScheduleOutput) ToSnapshotScheduleOutput() SnapshotScheduleOutput {
	return o
}

func (o SnapshotScheduleOutput) ToSnapshotScheduleOutputWithContext(ctx context.Context) SnapshotScheduleOutput {
	return o
}

func (o SnapshotScheduleOutput) ToSnapshotSchedulePtrOutput() SnapshotSchedulePtrOutput {
	return o.ToSnapshotSchedulePtrOutputWithContext(context.Background())
}

func (o SnapshotScheduleOutput) ToSnapshotSchedulePtrOutputWithContext(ctx context.Context) SnapshotSchedulePtrOutput {
	return o.ApplyT(func(v SnapshotSchedule) *SnapshotSchedule {
		return &v
	}).(SnapshotSchedulePtrOutput)
}

type SnapshotSchedulePtrOutput struct {
	*pulumi.OutputState
}

func (SnapshotSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSchedule)(nil))
}

func (o SnapshotSchedulePtrOutput) ToSnapshotSchedulePtrOutput() SnapshotSchedulePtrOutput {
	return o
}

func (o SnapshotSchedulePtrOutput) ToSnapshotSchedulePtrOutputWithContext(ctx context.Context) SnapshotSchedulePtrOutput {
	return o
}

type SnapshotScheduleArrayOutput struct{ *pulumi.OutputState }

func (SnapshotScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SnapshotSchedule)(nil))
}

func (o SnapshotScheduleArrayOutput) ToSnapshotScheduleArrayOutput() SnapshotScheduleArrayOutput {
	return o
}

func (o SnapshotScheduleArrayOutput) ToSnapshotScheduleArrayOutputWithContext(ctx context.Context) SnapshotScheduleArrayOutput {
	return o
}

func (o SnapshotScheduleArrayOutput) Index(i pulumi.IntInput) SnapshotScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SnapshotSchedule {
		return vs[0].([]SnapshotSchedule)[vs[1].(int)]
	}).(SnapshotScheduleOutput)
}

type SnapshotScheduleMapOutput struct{ *pulumi.OutputState }

func (SnapshotScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SnapshotSchedule)(nil))
}

func (o SnapshotScheduleMapOutput) ToSnapshotScheduleMapOutput() SnapshotScheduleMapOutput {
	return o
}

func (o SnapshotScheduleMapOutput) ToSnapshotScheduleMapOutputWithContext(ctx context.Context) SnapshotScheduleMapOutput {
	return o
}

func (o SnapshotScheduleMapOutput) MapIndex(k pulumi.StringInput) SnapshotScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SnapshotSchedule {
		return vs[0].(map[string]SnapshotSchedule)[vs[1].(string)]
	}).(SnapshotScheduleOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotScheduleOutput{})
	pulumi.RegisterOutputType(SnapshotSchedulePtrOutput{})
	pulumi.RegisterOutputType(SnapshotScheduleArrayOutput{})
	pulumi.RegisterOutputType(SnapshotScheduleMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type StateMachineLoggingConfiguration struct {
	// Determines whether execution data is included in your log. When set to FALSE, data is excluded.
	IncludeExecutionData *bool `pulumi:"includeExecutionData"`
	// Defines which category of execution history events are logged. Valid Values: ALL | ERROR | FATAL | OFF
	Level *string `pulumi:"level"`
	// Amazon Resource Name (ARN) of CloudWatch log group. Make sure the State Machine does have the right IAM Policies for Logging.
	LogDestination *string `pulumi:"logDestination"`
}

// StateMachineLoggingConfigurationInput is an input type that accepts StateMachineLoggingConfigurationArgs and StateMachineLoggingConfigurationOutput values.
// You can construct a concrete instance of `StateMachineLoggingConfigurationInput` via:
//
//          StateMachineLoggingConfigurationArgs{...}
type StateMachineLoggingConfigurationInput interface {
	pulumi.Input

	ToStateMachineLoggingConfigurationOutput() StateMachineLoggingConfigurationOutput
	ToStateMachineLoggingConfigurationOutputWithContext(context.Context) StateMachineLoggingConfigurationOutput
}

type StateMachineLoggingConfigurationArgs struct {
	// Determines whether execution data is included in your log. When set to FALSE, data is excluded.
	IncludeExecutionData pulumi.BoolPtrInput `pulumi:"includeExecutionData"`
	// Defines which category of execution history events are logged. Valid Values: ALL | ERROR | FATAL | OFF
	Level pulumi.StringPtrInput `pulumi:"level"`
	// Amazon Resource Name (ARN) of CloudWatch log group. Make sure the State Machine does have the right IAM Policies for Logging.
	LogDestination pulumi.StringPtrInput `pulumi:"logDestination"`
}

func (StateMachineLoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineLoggingConfiguration)(nil)).Elem()
}

func (i StateMachineLoggingConfigurationArgs) ToStateMachineLoggingConfigurationOutput() StateMachineLoggingConfigurationOutput {
	return i.ToStateMachineLoggingConfigurationOutputWithContext(context.Background())
}

func (i StateMachineLoggingConfigurationArgs) ToStateMachineLoggingConfigurationOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineLoggingConfigurationOutput)
}

func (i StateMachineLoggingConfigurationArgs) ToStateMachineLoggingConfigurationPtrOutput() StateMachineLoggingConfigurationPtrOutput {
	return i.ToStateMachineLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i StateMachineLoggingConfigurationArgs) ToStateMachineLoggingConfigurationPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineLoggingConfigurationOutput).ToStateMachineLoggingConfigurationPtrOutputWithContext(ctx)
}

// StateMachineLoggingConfigurationPtrInput is an input type that accepts StateMachineLoggingConfigurationArgs, StateMachineLoggingConfigurationPtr and StateMachineLoggingConfigurationPtrOutput values.
// You can construct a concrete instance of `StateMachineLoggingConfigurationPtrInput` via:
//
//          StateMachineLoggingConfigurationArgs{...}
//
//  or:
//
//          nil
type StateMachineLoggingConfigurationPtrInput interface {
	pulumi.Input

	ToStateMachineLoggingConfigurationPtrOutput() StateMachineLoggingConfigurationPtrOutput
	ToStateMachineLoggingConfigurationPtrOutputWithContext(context.Context) StateMachineLoggingConfigurationPtrOutput
}

type stateMachineLoggingConfigurationPtrType StateMachineLoggingConfigurationArgs

func StateMachineLoggingConfigurationPtr(v *StateMachineLoggingConfigurationArgs) StateMachineLoggingConfigurationPtrInput {
	return (*stateMachineLoggingConfigurationPtrType)(v)
}

func (*stateMachineLoggingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StateMachineLoggingConfiguration)(nil)).Elem()
}

func (i *stateMachineLoggingConfigurationPtrType) ToStateMachineLoggingConfigurationPtrOutput() StateMachineLoggingConfigurationPtrOutput {
	return i.ToStateMachineLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i *stateMachineLoggingConfigurationPtrType) ToStateMachineLoggingConfigurationPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineLoggingConfigurationPtrOutput)
}

type StateMachineLoggingConfigurationOutput struct{ *pulumi.OutputState }

func (StateMachineLoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachineLoggingConfiguration)(nil)).Elem()
}

func (o StateMachineLoggingConfigurationOutput) ToStateMachineLoggingConfigurationOutput() StateMachineLoggingConfigurationOutput {
	return o
}

func (o StateMachineLoggingConfigurationOutput) ToStateMachineLoggingConfigurationOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationOutput {
	return o
}

func (o StateMachineLoggingConfigurationOutput) ToStateMachineLoggingConfigurationPtrOutput() StateMachineLoggingConfigurationPtrOutput {
	return o.ToStateMachineLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (o StateMachineLoggingConfigurationOutput) ToStateMachineLoggingConfigurationPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v StateMachineLoggingConfiguration) *StateMachineLoggingConfiguration {
		return &v
	}).(StateMachineLoggingConfigurationPtrOutput)
}

// Determines whether execution data is included in your log. When set to FALSE, data is excluded.
func (o StateMachineLoggingConfigurationOutput) IncludeExecutionData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StateMachineLoggingConfiguration) *bool { return v.IncludeExecutionData }).(pulumi.BoolPtrOutput)
}

// Defines which category of execution history events are logged. Valid Values: ALL | ERROR | FATAL | OFF
func (o StateMachineLoggingConfigurationOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StateMachineLoggingConfiguration) *string { return v.Level }).(pulumi.StringPtrOutput)
}

// Amazon Resource Name (ARN) of CloudWatch log group. Make sure the State Machine does have the right IAM Policies for Logging.
func (o StateMachineLoggingConfigurationOutput) LogDestination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StateMachineLoggingConfiguration) *string { return v.LogDestination }).(pulumi.StringPtrOutput)
}

type StateMachineLoggingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StateMachineLoggingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StateMachineLoggingConfiguration)(nil)).Elem()
}

func (o StateMachineLoggingConfigurationPtrOutput) ToStateMachineLoggingConfigurationPtrOutput() StateMachineLoggingConfigurationPtrOutput {
	return o
}

func (o StateMachineLoggingConfigurationPtrOutput) ToStateMachineLoggingConfigurationPtrOutputWithContext(ctx context.Context) StateMachineLoggingConfigurationPtrOutput {
	return o
}

func (o StateMachineLoggingConfigurationPtrOutput) Elem() StateMachineLoggingConfigurationOutput {
	return o.ApplyT(func(v *StateMachineLoggingConfiguration) StateMachineLoggingConfiguration { return *v }).(StateMachineLoggingConfigurationOutput)
}

// Determines whether execution data is included in your log. When set to FALSE, data is excluded.
func (o StateMachineLoggingConfigurationPtrOutput) IncludeExecutionData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StateMachineLoggingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.IncludeExecutionData
	}).(pulumi.BoolPtrOutput)
}

// Defines which category of execution history events are logged. Valid Values: ALL | ERROR | FATAL | OFF
func (o StateMachineLoggingConfigurationPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StateMachineLoggingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

// Amazon Resource Name (ARN) of CloudWatch log group. Make sure the State Machine does have the right IAM Policies for Logging.
func (o StateMachineLoggingConfigurationPtrOutput) LogDestination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StateMachineLoggingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.LogDestination
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(StateMachineLoggingConfigurationOutput{})
	pulumi.RegisterOutputType(StateMachineLoggingConfigurationPtrOutput{})
}

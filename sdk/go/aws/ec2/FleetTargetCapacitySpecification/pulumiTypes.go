// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FleetTargetCapacitySpecification

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type FleetTargetCapacitySpecification struct {
	// Default target capacity type. Valid values: `on-demand`, `spot`.
	DefaultTargetCapacityType string `pulumi:"defaultTargetCapacityType"`
	// The number of On-Demand units to request.
	OnDemandTargetCapacity *int `pulumi:"onDemandTargetCapacity"`
	// The number of Spot units to request.
	SpotTargetCapacity *int `pulumi:"spotTargetCapacity"`
	// The number of units to request, filled using `defaultTargetCapacityType`.
	TotalTargetCapacity int `pulumi:"totalTargetCapacity"`
}

type FleetTargetCapacitySpecificationInput interface {
	pulumi.Input

	ToFleetTargetCapacitySpecificationOutput() FleetTargetCapacitySpecificationOutput
	ToFleetTargetCapacitySpecificationOutputWithContext(context.Context) FleetTargetCapacitySpecificationOutput
}

type FleetTargetCapacitySpecificationArgs struct {
	// Default target capacity type. Valid values: `on-demand`, `spot`.
	DefaultTargetCapacityType pulumi.StringInput `pulumi:"defaultTargetCapacityType"`
	// The number of On-Demand units to request.
	OnDemandTargetCapacity pulumi.IntPtrInput `pulumi:"onDemandTargetCapacity"`
	// The number of Spot units to request.
	SpotTargetCapacity pulumi.IntPtrInput `pulumi:"spotTargetCapacity"`
	// The number of units to request, filled using `defaultTargetCapacityType`.
	TotalTargetCapacity pulumi.IntInput `pulumi:"totalTargetCapacity"`
}

func (FleetTargetCapacitySpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FleetTargetCapacitySpecification)(nil)).Elem()
}

func (i FleetTargetCapacitySpecificationArgs) ToFleetTargetCapacitySpecificationOutput() FleetTargetCapacitySpecificationOutput {
	return i.ToFleetTargetCapacitySpecificationOutputWithContext(context.Background())
}

func (i FleetTargetCapacitySpecificationArgs) ToFleetTargetCapacitySpecificationOutputWithContext(ctx context.Context) FleetTargetCapacitySpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetTargetCapacitySpecificationOutput)
}

func (i FleetTargetCapacitySpecificationArgs) ToFleetTargetCapacitySpecificationPtrOutput() FleetTargetCapacitySpecificationPtrOutput {
	return i.ToFleetTargetCapacitySpecificationPtrOutputWithContext(context.Background())
}

func (i FleetTargetCapacitySpecificationArgs) ToFleetTargetCapacitySpecificationPtrOutputWithContext(ctx context.Context) FleetTargetCapacitySpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetTargetCapacitySpecificationOutput).ToFleetTargetCapacitySpecificationPtrOutputWithContext(ctx)
}

type FleetTargetCapacitySpecificationPtrInput interface {
	pulumi.Input

	ToFleetTargetCapacitySpecificationPtrOutput() FleetTargetCapacitySpecificationPtrOutput
	ToFleetTargetCapacitySpecificationPtrOutputWithContext(context.Context) FleetTargetCapacitySpecificationPtrOutput
}

type fleetTargetCapacitySpecificationPtrType FleetTargetCapacitySpecificationArgs

func FleetTargetCapacitySpecificationPtr(v *FleetTargetCapacitySpecificationArgs) FleetTargetCapacitySpecificationPtrInput {	return (*fleetTargetCapacitySpecificationPtrType)(v)
}

func (*fleetTargetCapacitySpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetTargetCapacitySpecification)(nil)).Elem()
}

func (i *fleetTargetCapacitySpecificationPtrType) ToFleetTargetCapacitySpecificationPtrOutput() FleetTargetCapacitySpecificationPtrOutput {
	return i.ToFleetTargetCapacitySpecificationPtrOutputWithContext(context.Background())
}

func (i *fleetTargetCapacitySpecificationPtrType) ToFleetTargetCapacitySpecificationPtrOutputWithContext(ctx context.Context) FleetTargetCapacitySpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetTargetCapacitySpecificationPtrOutput)
}

type FleetTargetCapacitySpecificationOutput struct { *pulumi.OutputState }

func (FleetTargetCapacitySpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FleetTargetCapacitySpecification)(nil)).Elem()
}

func (o FleetTargetCapacitySpecificationOutput) ToFleetTargetCapacitySpecificationOutput() FleetTargetCapacitySpecificationOutput {
	return o
}

func (o FleetTargetCapacitySpecificationOutput) ToFleetTargetCapacitySpecificationOutputWithContext(ctx context.Context) FleetTargetCapacitySpecificationOutput {
	return o
}

func (o FleetTargetCapacitySpecificationOutput) ToFleetTargetCapacitySpecificationPtrOutput() FleetTargetCapacitySpecificationPtrOutput {
	return o.ToFleetTargetCapacitySpecificationPtrOutputWithContext(context.Background())
}

func (o FleetTargetCapacitySpecificationOutput) ToFleetTargetCapacitySpecificationPtrOutputWithContext(ctx context.Context) FleetTargetCapacitySpecificationPtrOutput {
	return o.ApplyT(func(v FleetTargetCapacitySpecification) *FleetTargetCapacitySpecification {
		return &v
	}).(FleetTargetCapacitySpecificationPtrOutput)
}
// Default target capacity type. Valid values: `on-demand`, `spot`.
func (o FleetTargetCapacitySpecificationOutput) DefaultTargetCapacityType() pulumi.StringOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) string { return v.DefaultTargetCapacityType }).(pulumi.StringOutput)
}

// The number of On-Demand units to request.
func (o FleetTargetCapacitySpecificationOutput) OnDemandTargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) *int { return v.OnDemandTargetCapacity }).(pulumi.IntPtrOutput)
}

// The number of Spot units to request.
func (o FleetTargetCapacitySpecificationOutput) SpotTargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) *int { return v.SpotTargetCapacity }).(pulumi.IntPtrOutput)
}

// The number of units to request, filled using `defaultTargetCapacityType`.
func (o FleetTargetCapacitySpecificationOutput) TotalTargetCapacity() pulumi.IntOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) int { return v.TotalTargetCapacity }).(pulumi.IntOutput)
}

type FleetTargetCapacitySpecificationPtrOutput struct { *pulumi.OutputState}

func (FleetTargetCapacitySpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetTargetCapacitySpecification)(nil)).Elem()
}

func (o FleetTargetCapacitySpecificationPtrOutput) ToFleetTargetCapacitySpecificationPtrOutput() FleetTargetCapacitySpecificationPtrOutput {
	return o
}

func (o FleetTargetCapacitySpecificationPtrOutput) ToFleetTargetCapacitySpecificationPtrOutputWithContext(ctx context.Context) FleetTargetCapacitySpecificationPtrOutput {
	return o
}

func (o FleetTargetCapacitySpecificationPtrOutput) Elem() FleetTargetCapacitySpecificationOutput {
	return o.ApplyT(func (v *FleetTargetCapacitySpecification) FleetTargetCapacitySpecification { return *v }).(FleetTargetCapacitySpecificationOutput)
}

// Default target capacity type. Valid values: `on-demand`, `spot`.
func (o FleetTargetCapacitySpecificationPtrOutput) DefaultTargetCapacityType() pulumi.StringOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) string { return v.DefaultTargetCapacityType }).(pulumi.StringOutput)
}

// The number of On-Demand units to request.
func (o FleetTargetCapacitySpecificationPtrOutput) OnDemandTargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) *int { return v.OnDemandTargetCapacity }).(pulumi.IntPtrOutput)
}

// The number of Spot units to request.
func (o FleetTargetCapacitySpecificationPtrOutput) SpotTargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) *int { return v.SpotTargetCapacity }).(pulumi.IntPtrOutput)
}

// The number of units to request, filled using `defaultTargetCapacityType`.
func (o FleetTargetCapacitySpecificationPtrOutput) TotalTargetCapacity() pulumi.IntOutput {
	return o.ApplyT(func (v FleetTargetCapacitySpecification) int { return v.TotalTargetCapacity }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetTargetCapacitySpecificationOutput{})
	pulumi.RegisterOutputType(FleetTargetCapacitySpecificationPtrOutput{})
}

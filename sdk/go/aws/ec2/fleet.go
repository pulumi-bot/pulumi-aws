// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage EC2 Fleets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = ec2.NewFleet(ctx, "example", &ec2.FleetArgs{
// 			LaunchTemplateConfig: &ec2.FleetLaunchTemplateConfigArgs{
// 				LaunchTemplateSpecification: &ec2.FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs{
// 					LaunchTemplateId: pulumi.String(aws_launch_template.Example.Id),
// 					Version:          pulumi.String(aws_launch_template.Example.Latest_version),
// 				},
// 			},
// 			TargetCapacitySpecification: &ec2.FleetTargetCapacitySpecificationArgs{
// 				DefaultTargetCapacityType: pulumi.String("spot"),
// 				TotalTargetCapacity:       pulumi.Int(5),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Fleet struct {
	pulumi.CustomResourceState

	// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
	ExcessCapacityTerminationPolicy pulumi.StringPtrOutput `pulumi:"excessCapacityTerminationPolicy"`
	// Nested argument containing EC2 Launch Template configurations. Defined below.
	LaunchTemplateConfig FleetLaunchTemplateConfigOutput `pulumi:"launchTemplateConfig"`
	// Nested argument containing On-Demand configurations. Defined below.
	OnDemandOptions FleetOnDemandOptionsPtrOutput `pulumi:"onDemandOptions"`
	// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
	ReplaceUnhealthyInstances pulumi.BoolPtrOutput `pulumi:"replaceUnhealthyInstances"`
	// Nested argument containing Spot configurations. Defined below.
	SpotOptions FleetSpotOptionsPtrOutput `pulumi:"spotOptions"`
	// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Nested argument containing target capacity configurations. Defined below.
	TargetCapacitySpecification FleetTargetCapacitySpecificationOutput `pulumi:"targetCapacitySpecification"`
	// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
	TerminateInstances pulumi.BoolPtrOutput `pulumi:"terminateInstances"`
	// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
	TerminateInstancesWithExpiration pulumi.BoolPtrOutput `pulumi:"terminateInstancesWithExpiration"`
	// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil || args.LaunchTemplateConfig == nil {
		return nil, errors.New("missing required argument 'LaunchTemplateConfig'")
	}
	if args == nil || args.TargetCapacitySpecification == nil {
		return nil, errors.New("missing required argument 'TargetCapacitySpecification'")
	}
	if args == nil {
		args = &FleetArgs{}
	}
	var resource Fleet
	err := ctx.RegisterResource("aws:ec2/fleet:Fleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetState, opts ...pulumi.ResourceOption) (*Fleet, error) {
	var resource Fleet
	err := ctx.ReadResource("aws:ec2/fleet:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fleet resources.
type fleetState struct {
	// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
	ExcessCapacityTerminationPolicy *string `pulumi:"excessCapacityTerminationPolicy"`
	// Nested argument containing EC2 Launch Template configurations. Defined below.
	LaunchTemplateConfig *FleetLaunchTemplateConfig `pulumi:"launchTemplateConfig"`
	// Nested argument containing On-Demand configurations. Defined below.
	OnDemandOptions *FleetOnDemandOptions `pulumi:"onDemandOptions"`
	// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
	ReplaceUnhealthyInstances *bool `pulumi:"replaceUnhealthyInstances"`
	// Nested argument containing Spot configurations. Defined below.
	SpotOptions *FleetSpotOptions `pulumi:"spotOptions"`
	// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
	Tags map[string]interface{} `pulumi:"tags"`
	// Nested argument containing target capacity configurations. Defined below.
	TargetCapacitySpecification *FleetTargetCapacitySpecification `pulumi:"targetCapacitySpecification"`
	// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
	TerminateInstances *bool `pulumi:"terminateInstances"`
	// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
	TerminateInstancesWithExpiration *bool `pulumi:"terminateInstancesWithExpiration"`
	// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
	Type *string `pulumi:"type"`
}

type FleetState struct {
	// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
	ExcessCapacityTerminationPolicy pulumi.StringPtrInput
	// Nested argument containing EC2 Launch Template configurations. Defined below.
	LaunchTemplateConfig FleetLaunchTemplateConfigPtrInput
	// Nested argument containing On-Demand configurations. Defined below.
	OnDemandOptions FleetOnDemandOptionsPtrInput
	// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
	ReplaceUnhealthyInstances pulumi.BoolPtrInput
	// Nested argument containing Spot configurations. Defined below.
	SpotOptions FleetSpotOptionsPtrInput
	// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
	Tags pulumi.MapInput
	// Nested argument containing target capacity configurations. Defined below.
	TargetCapacitySpecification FleetTargetCapacitySpecificationPtrInput
	// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
	TerminateInstances pulumi.BoolPtrInput
	// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
	TerminateInstancesWithExpiration pulumi.BoolPtrInput
	// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
	Type pulumi.StringPtrInput
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
	ExcessCapacityTerminationPolicy *string `pulumi:"excessCapacityTerminationPolicy"`
	// Nested argument containing EC2 Launch Template configurations. Defined below.
	LaunchTemplateConfig FleetLaunchTemplateConfig `pulumi:"launchTemplateConfig"`
	// Nested argument containing On-Demand configurations. Defined below.
	OnDemandOptions *FleetOnDemandOptions `pulumi:"onDemandOptions"`
	// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
	ReplaceUnhealthyInstances *bool `pulumi:"replaceUnhealthyInstances"`
	// Nested argument containing Spot configurations. Defined below.
	SpotOptions *FleetSpotOptions `pulumi:"spotOptions"`
	// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
	Tags map[string]interface{} `pulumi:"tags"`
	// Nested argument containing target capacity configurations. Defined below.
	TargetCapacitySpecification FleetTargetCapacitySpecification `pulumi:"targetCapacitySpecification"`
	// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
	TerminateInstances *bool `pulumi:"terminateInstances"`
	// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
	TerminateInstancesWithExpiration *bool `pulumi:"terminateInstancesWithExpiration"`
	// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
	ExcessCapacityTerminationPolicy pulumi.StringPtrInput
	// Nested argument containing EC2 Launch Template configurations. Defined below.
	LaunchTemplateConfig FleetLaunchTemplateConfigInput
	// Nested argument containing On-Demand configurations. Defined below.
	OnDemandOptions FleetOnDemandOptionsPtrInput
	// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
	ReplaceUnhealthyInstances pulumi.BoolPtrInput
	// Nested argument containing Spot configurations. Defined below.
	SpotOptions FleetSpotOptionsPtrInput
	// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
	Tags pulumi.MapInput
	// Nested argument containing target capacity configurations. Defined below.
	TargetCapacitySpecification FleetTargetCapacitySpecificationInput
	// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
	TerminateInstances pulumi.BoolPtrInput
	// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
	TerminateInstancesWithExpiration pulumi.BoolPtrInput
	// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
	Type pulumi.StringPtrInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

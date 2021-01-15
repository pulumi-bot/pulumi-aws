// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Gamelift Fleet resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/gamelift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gamelift.NewFleet(ctx, "example", &gamelift.FleetArgs{
// 			BuildId:         pulumi.Any(aws_gamelift_build.Example.Id),
// 			Ec2InstanceType: pulumi.String("t2.micro"),
// 			FleetType:       pulumi.String("ON_DEMAND"),
// 			RuntimeConfiguration: &gamelift.FleetRuntimeConfigurationArgs{
// 				ServerProcesses: gamelift.FleetRuntimeConfigurationServerProcessArray{
// 					&gamelift.FleetRuntimeConfigurationServerProcessArgs{
// 						ConcurrentExecutions: pulumi.Int(1),
// 						LaunchPath:           pulumi.String("C:\\game\\GomokuServer.exe"),
// 					},
// 				},
// 			},
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
// Gamelift Fleets cannot be imported at this time.
type Fleet struct {
	pulumi.CustomResourceState

	// Fleet ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId pulumi.StringOutput `pulumi:"buildId"`
	// Human-readable description of the fleet.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions FleetEc2InboundPermissionArrayOutput `pulumi:"ec2InboundPermissions"`
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType pulumi.StringOutput `pulumi:"ec2InstanceType"`
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType pulumi.StringPtrOutput `pulumi:"fleetType"`
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn pulumi.StringPtrOutput   `pulumi:"instanceRoleArn"`
	LogPaths        pulumi.StringArrayOutput `pulumi:"logPaths"`
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups pulumi.StringArrayOutput `pulumi:"metricGroups"`
	// The name of the fleet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy pulumi.StringPtrOutput `pulumi:"newGameSessionProtectionPolicy"`
	// Operating system of the fleet's computing resources.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy FleetResourceCreationLimitPolicyPtrOutput `pulumi:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration FleetRuntimeConfigurationPtrOutput `pulumi:"runtimeConfiguration"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildId == nil {
		return nil, errors.New("invalid value for required argument 'BuildId'")
	}
	if args.Ec2InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'Ec2InstanceType'")
	}
	var resource Fleet
	err := ctx.RegisterResource("aws:gamelift/fleet:Fleet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:gamelift/fleet:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fleet resources.
type fleetState struct {
	// Fleet ARN.
	Arn *string `pulumi:"arn"`
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId *string `pulumi:"buildId"`
	// Human-readable description of the fleet.
	Description *string `pulumi:"description"`
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions []FleetEc2InboundPermission `pulumi:"ec2InboundPermissions"`
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType *string `pulumi:"ec2InstanceType"`
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType *string `pulumi:"fleetType"`
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn *string  `pulumi:"instanceRoleArn"`
	LogPaths        []string `pulumi:"logPaths"`
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups []string `pulumi:"metricGroups"`
	// The name of the fleet.
	Name *string `pulumi:"name"`
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy *string `pulumi:"newGameSessionProtectionPolicy"`
	// Operating system of the fleet's computing resources.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy *FleetResourceCreationLimitPolicy `pulumi:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration *FleetRuntimeConfiguration `pulumi:"runtimeConfiguration"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type FleetState struct {
	// Fleet ARN.
	Arn pulumi.StringPtrInput
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId pulumi.StringPtrInput
	// Human-readable description of the fleet.
	Description pulumi.StringPtrInput
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions FleetEc2InboundPermissionArrayInput
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType pulumi.StringPtrInput
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType pulumi.StringPtrInput
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn pulumi.StringPtrInput
	LogPaths        pulumi.StringArrayInput
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups pulumi.StringArrayInput
	// The name of the fleet.
	Name pulumi.StringPtrInput
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy pulumi.StringPtrInput
	// Operating system of the fleet's computing resources.
	OperatingSystem pulumi.StringPtrInput
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy FleetResourceCreationLimitPolicyPtrInput
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration FleetRuntimeConfigurationPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId string `pulumi:"buildId"`
	// Human-readable description of the fleet.
	Description *string `pulumi:"description"`
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions []FleetEc2InboundPermission `pulumi:"ec2InboundPermissions"`
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType string `pulumi:"ec2InstanceType"`
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType *string `pulumi:"fleetType"`
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn *string `pulumi:"instanceRoleArn"`
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups []string `pulumi:"metricGroups"`
	// The name of the fleet.
	Name *string `pulumi:"name"`
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy *string `pulumi:"newGameSessionProtectionPolicy"`
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy *FleetResourceCreationLimitPolicy `pulumi:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration *FleetRuntimeConfiguration `pulumi:"runtimeConfiguration"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId pulumi.StringInput
	// Human-readable description of the fleet.
	Description pulumi.StringPtrInput
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions FleetEc2InboundPermissionArrayInput
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType pulumi.StringInput
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType pulumi.StringPtrInput
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn pulumi.StringPtrInput
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups pulumi.StringArrayInput
	// The name of the fleet.
	Name pulumi.StringPtrInput
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy pulumi.StringPtrInput
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy FleetResourceCreationLimitPolicyPtrInput
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration FleetRuntimeConfigurationPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

type FleetInput interface {
	pulumi.Input

	ToFleetOutput() FleetOutput
	ToFleetOutputWithContext(ctx context.Context) FleetOutput
}

func (*Fleet) ElementType() reflect.Type {
	return reflect.TypeOf((*Fleet)(nil))
}

func (i *Fleet) ToFleetOutput() FleetOutput {
	return i.ToFleetOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetOutput)
}

func (i *Fleet) ToFleetPtrOutput() FleetPtrOutput {
	return i.ToFleetPtrOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetPtrOutput)
}

type FleetPtrInput interface {
	pulumi.Input

	ToFleetPtrOutput() FleetPtrOutput
	ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput
}

type fleetPtrType FleetArgs

func (*fleetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil))
}

func (i *fleetPtrType) ToFleetPtrOutput() FleetPtrOutput {
	return i.ToFleetPtrOutputWithContext(context.Background())
}

func (i *fleetPtrType) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetPtrOutput)
}

// FleetArrayInput is an input type that accepts FleetArray and FleetArrayOutput values.
// You can construct a concrete instance of `FleetArrayInput` via:
//
//          FleetArray{ FleetArgs{...} }
type FleetArrayInput interface {
	pulumi.Input

	ToFleetArrayOutput() FleetArrayOutput
	ToFleetArrayOutputWithContext(context.Context) FleetArrayOutput
}

type FleetArray []FleetInput

func (FleetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Fleet)(nil))
}

func (i FleetArray) ToFleetArrayOutput() FleetArrayOutput {
	return i.ToFleetArrayOutputWithContext(context.Background())
}

func (i FleetArray) ToFleetArrayOutputWithContext(ctx context.Context) FleetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetArrayOutput)
}

// FleetMapInput is an input type that accepts FleetMap and FleetMapOutput values.
// You can construct a concrete instance of `FleetMapInput` via:
//
//          FleetMap{ "key": FleetArgs{...} }
type FleetMapInput interface {
	pulumi.Input

	ToFleetMapOutput() FleetMapOutput
	ToFleetMapOutputWithContext(context.Context) FleetMapOutput
}

type FleetMap map[string]FleetInput

func (FleetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Fleet)(nil))
}

func (i FleetMap) ToFleetMapOutput() FleetMapOutput {
	return i.ToFleetMapOutputWithContext(context.Background())
}

func (i FleetMap) ToFleetMapOutputWithContext(ctx context.Context) FleetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetMapOutput)
}

type FleetOutput struct {
	*pulumi.OutputState
}

func (FleetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Fleet)(nil))
}

func (o FleetOutput) ToFleetOutput() FleetOutput {
	return o
}

func (o FleetOutput) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return o
}

func (o FleetOutput) ToFleetPtrOutput() FleetPtrOutput {
	return o.ToFleetPtrOutputWithContext(context.Background())
}

func (o FleetOutput) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return o.ApplyT(func(v Fleet) *Fleet {
		return &v
	}).(FleetPtrOutput)
}

type FleetPtrOutput struct {
	*pulumi.OutputState
}

func (FleetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil))
}

func (o FleetPtrOutput) ToFleetPtrOutput() FleetPtrOutput {
	return o
}

func (o FleetPtrOutput) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return o
}

type FleetArrayOutput struct{ *pulumi.OutputState }

func (FleetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Fleet)(nil))
}

func (o FleetArrayOutput) ToFleetArrayOutput() FleetArrayOutput {
	return o
}

func (o FleetArrayOutput) ToFleetArrayOutputWithContext(ctx context.Context) FleetArrayOutput {
	return o
}

func (o FleetArrayOutput) Index(i pulumi.IntInput) FleetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Fleet {
		return vs[0].([]Fleet)[vs[1].(int)]
	}).(FleetOutput)
}

type FleetMapOutput struct{ *pulumi.OutputState }

func (FleetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Fleet)(nil))
}

func (o FleetMapOutput) ToFleetMapOutput() FleetMapOutput {
	return o
}

func (o FleetMapOutput) ToFleetMapOutputWithContext(ctx context.Context) FleetMapOutput {
	return o
}

func (o FleetMapOutput) MapIndex(k pulumi.StringInput) FleetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Fleet {
		return vs[0].(map[string]Fleet)[vs[1].(string)]
	}).(FleetOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetOutput{})
	pulumi.RegisterOutputType(FleetPtrOutput{})
	pulumi.RegisterOutputType(FleetArrayOutput{})
	pulumi.RegisterOutputType(FleetMapOutput{})
}

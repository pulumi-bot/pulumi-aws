// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).
//
// > **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `minSize` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testGroup, err := autoscaling.NewGroup(ctx, "testGroup", &autoscaling.GroupArgs{
// 			Tags: autoscaling.GroupTagArray{
// 				&autoscaling.GroupTagArgs{
// 					Key:               pulumi.String("AmazonECSManaged"),
// 					PropagateAtLaunch: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecs.NewCapacityProvider(ctx, "testCapacityProvider", &ecs.CapacityProviderArgs{
// 			AutoScalingGroupProvider: &ecs.CapacityProviderAutoScalingGroupProviderArgs{
// 				AutoScalingGroupArn:          testGroup.Arn,
// 				ManagedTerminationProtection: pulumi.String("ENABLED"),
// 				ManagedScaling: &ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs{
// 					MaximumScalingStepSize: pulumi.Int(1000),
// 					MinimumScalingStepSize: pulumi.Int(1),
// 					Status:                 pulumi.String("ENABLED"),
// 					TargetCapacity:         pulumi.Int(10),
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
type CapacityProvider struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the capacity provider.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Nested argument defining the provider for the ECS auto scaling group. Defined below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProviderOutput `pulumi:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewCapacityProvider registers a new resource with the given unique name, arguments, and options.
func NewCapacityProvider(ctx *pulumi.Context,
	name string, args *CapacityProviderArgs, opts ...pulumi.ResourceOption) (*CapacityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutoScalingGroupProvider == nil {
		return nil, errors.New("invalid value for required argument 'AutoScalingGroupProvider'")
	}
	var resource CapacityProvider
	err := ctx.RegisterResource("aws:ecs/capacityProvider:CapacityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityProvider gets an existing CapacityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityProviderState, opts ...pulumi.ResourceOption) (*CapacityProvider, error) {
	var resource CapacityProvider
	err := ctx.ReadResource("aws:ecs/capacityProvider:CapacityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityProvider resources.
type capacityProviderState struct {
	// The Amazon Resource Name (ARN) that identifies the capacity provider.
	Arn *string `pulumi:"arn"`
	// Nested argument defining the provider for the ECS auto scaling group. Defined below.
	AutoScalingGroupProvider *CapacityProviderAutoScalingGroupProvider `pulumi:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type CapacityProviderState struct {
	// The Amazon Resource Name (ARN) that identifies the capacity provider.
	Arn pulumi.StringPtrInput
	// Nested argument defining the provider for the ECS auto scaling group. Defined below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProviderPtrInput
	// The name of the capacity provider.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (CapacityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityProviderState)(nil)).Elem()
}

type capacityProviderArgs struct {
	// Nested argument defining the provider for the ECS auto scaling group. Defined below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProvider `pulumi:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CapacityProvider resource.
type CapacityProviderArgs struct {
	// Nested argument defining the provider for the ECS auto scaling group. Defined below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProviderInput
	// The name of the capacity provider.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (CapacityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityProviderArgs)(nil)).Elem()
}

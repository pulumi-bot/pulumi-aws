// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Managed Scaling policy for EMR Cluster. With Amazon EMR versions 5.30.0 and later (except for Amazon EMR 6.0.0), you can enable EMR managed scaling to automatically increase or decrease the number of instances or units in your cluster based on workload. See [Using EMR Managed Scaling in Amazon EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-managed-scaling.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/emr"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sample, err := emr.NewCluster(ctx, "sample", &emr.ClusterArgs{
// 			ReleaseLabel: pulumi.String("emr-5.30.0"),
// 			MasterInstanceGroup: &emr.ClusterMasterInstanceGroupArgs{
// 				InstanceType: pulumi.String("m4.large"),
// 			},
// 			CoreInstanceGroup: &emr.ClusterCoreInstanceGroupArgs{
// 				InstanceType: pulumi.String("c4.large"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = emr.NewManagedScalingPolicy(ctx, "samplepolicy", &emr.ManagedScalingPolicyArgs{
// 			ClusterId: sample.ID(),
// 			ComputeLimits: emr.ManagedScalingPolicyComputeLimitArray{
// 				&emr.ManagedScalingPolicyComputeLimitArgs{
// 					UnitType:                     pulumi.String("Instances"),
// 					MinimumCapacityUnits:         pulumi.Int(2),
// 					MaximumCapacityUnits:         pulumi.Int(10),
// 					MaximumOndemandCapacityUnits: pulumi.Int(2),
// 					MaximumCoreCapacityUnits:     pulumi.Int(10),
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
// EMR Managed Scaling Policies can be imported via the EMR Cluster identifier, e.g. console
//
// ```sh
//  $ pulumi import aws:emr/managedScalingPolicy:ManagedScalingPolicy example j-123456ABCDEF
// ```
type ManagedScalingPolicy struct {
	pulumi.CustomResourceState

	// The id of the EMR cluster
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Configuration block with compute limit settings. Described below.
	ComputeLimits ManagedScalingPolicyComputeLimitArrayOutput `pulumi:"computeLimits"`
}

// NewManagedScalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedScalingPolicy(ctx *pulumi.Context,
	name string, args *ManagedScalingPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedScalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ComputeLimits == nil {
		return nil, errors.New("invalid value for required argument 'ComputeLimits'")
	}
	var resource ManagedScalingPolicy
	err := ctx.RegisterResource("aws:emr/managedScalingPolicy:ManagedScalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedScalingPolicy gets an existing ManagedScalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedScalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedScalingPolicyState, opts ...pulumi.ResourceOption) (*ManagedScalingPolicy, error) {
	var resource ManagedScalingPolicy
	err := ctx.ReadResource("aws:emr/managedScalingPolicy:ManagedScalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedScalingPolicy resources.
type managedScalingPolicyState struct {
	// The id of the EMR cluster
	ClusterId *string `pulumi:"clusterId"`
	// Configuration block with compute limit settings. Described below.
	ComputeLimits []ManagedScalingPolicyComputeLimit `pulumi:"computeLimits"`
}

type ManagedScalingPolicyState struct {
	// The id of the EMR cluster
	ClusterId pulumi.StringPtrInput
	// Configuration block with compute limit settings. Described below.
	ComputeLimits ManagedScalingPolicyComputeLimitArrayInput
}

func (ManagedScalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedScalingPolicyState)(nil)).Elem()
}

type managedScalingPolicyArgs struct {
	// The id of the EMR cluster
	ClusterId string `pulumi:"clusterId"`
	// Configuration block with compute limit settings. Described below.
	ComputeLimits []ManagedScalingPolicyComputeLimit `pulumi:"computeLimits"`
}

// The set of arguments for constructing a ManagedScalingPolicy resource.
type ManagedScalingPolicyArgs struct {
	// The id of the EMR cluster
	ClusterId pulumi.StringInput
	// Configuration block with compute limit settings. Described below.
	ComputeLimits ManagedScalingPolicyComputeLimitArrayInput
}

func (ManagedScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedScalingPolicyArgs)(nil)).Elem()
}

type ManagedScalingPolicyInput interface {
	pulumi.Input

	ToManagedScalingPolicyOutput() ManagedScalingPolicyOutput
	ToManagedScalingPolicyOutputWithContext(ctx context.Context) ManagedScalingPolicyOutput
}

type ManagedScalingPolicyPtrInput interface {
	pulumi.Input

	ToManagedScalingPolicyPtrOutput() ManagedScalingPolicyPtrOutput
	ToManagedScalingPolicyPtrOutputWithContext(ctx context.Context) ManagedScalingPolicyPtrOutput
}

func (ManagedScalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedScalingPolicy)(nil)).Elem()
}

func (i ManagedScalingPolicy) ToManagedScalingPolicyOutput() ManagedScalingPolicyOutput {
	return i.ToManagedScalingPolicyOutputWithContext(context.Background())
}

func (i ManagedScalingPolicy) ToManagedScalingPolicyOutputWithContext(ctx context.Context) ManagedScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedScalingPolicyOutput)
}

func (i ManagedScalingPolicy) ToManagedScalingPolicyPtrOutput() ManagedScalingPolicyPtrOutput {
	return i.ToManagedScalingPolicyPtrOutputWithContext(context.Background())
}

func (i ManagedScalingPolicy) ToManagedScalingPolicyPtrOutputWithContext(ctx context.Context) ManagedScalingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedScalingPolicyPtrOutput)
}

type ManagedScalingPolicyOutput struct {
	*pulumi.OutputState
}

func (ManagedScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedScalingPolicyOutput)(nil)).Elem()
}

func (o ManagedScalingPolicyOutput) ToManagedScalingPolicyOutput() ManagedScalingPolicyOutput {
	return o
}

func (o ManagedScalingPolicyOutput) ToManagedScalingPolicyOutputWithContext(ctx context.Context) ManagedScalingPolicyOutput {
	return o
}

type ManagedScalingPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ManagedScalingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedScalingPolicy)(nil)).Elem()
}

func (o ManagedScalingPolicyPtrOutput) ToManagedScalingPolicyPtrOutput() ManagedScalingPolicyPtrOutput {
	return o
}

func (o ManagedScalingPolicyPtrOutput) ToManagedScalingPolicyPtrOutputWithContext(ctx context.Context) ManagedScalingPolicyPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedScalingPolicyOutput{})
	pulumi.RegisterOutputType(ManagedScalingPolicyPtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Managed Scaling policy for EMR Cluster. With Amazon EMR versions 5.30.0 and later (except for Amazon EMR 6.0.0), you can enable EMR managed scaling to automatically increase or decrease the number of instances or units in your cluster based on workload. See [Using EMR Managed Scaling in Amazon EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-managed-scaling.html) for more information.
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
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.ComputeLimits == nil {
		return nil, errors.New("missing required argument 'ComputeLimits'")
	}
	if args == nil {
		args = &ManagedScalingPolicyArgs{}
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

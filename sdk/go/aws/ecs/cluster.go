// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ECS cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewCluster(ctx, "foo", nil)
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
// ECS clusters can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:ecs/cluster:Cluster stateless stateless-app
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the cluster
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders pulumi.StringArrayOutput `pulumi:"capacityProviders"`
	// The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
	DefaultCapacityProviderStrategies ClusterDefaultCapacityProviderStrategyArrayOutput `pulumi:"defaultCapacityProviderStrategies"`
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings ClusterSettingArrayOutput `pulumi:"settings"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}

	var resource Cluster
	err := ctx.RegisterResource("aws:ecs/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws:ecs/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The Amazon Resource Name (ARN) that identifies the cluster
	Arn *string `pulumi:"arn"`
	// List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders []string `pulumi:"capacityProviders"`
	// The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
	DefaultCapacityProviderStrategies []ClusterDefaultCapacityProviderStrategy `pulumi:"defaultCapacityProviderStrategies"`
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name *string `pulumi:"name"`
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings []ClusterSetting `pulumi:"settings"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type ClusterState struct {
	// The Amazon Resource Name (ARN) that identifies the cluster
	Arn pulumi.StringPtrInput
	// List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders pulumi.StringArrayInput
	// The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
	DefaultCapacityProviderStrategies ClusterDefaultCapacityProviderStrategyArrayInput
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringPtrInput
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings ClusterSettingArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders []string `pulumi:"capacityProviders"`
	// The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
	DefaultCapacityProviderStrategies []ClusterDefaultCapacityProviderStrategy `pulumi:"defaultCapacityProviderStrategies"`
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name *string `pulumi:"name"`
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings []ClusterSetting `pulumi:"settings"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
	CapacityProviders pulumi.StringArrayInput
	// The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
	DefaultCapacityProviderStrategies ClusterDefaultCapacityProviderStrategyArrayInput
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringPtrInput
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings ClusterSettingArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}

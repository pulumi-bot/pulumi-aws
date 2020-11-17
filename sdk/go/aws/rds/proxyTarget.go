// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS DB proxy target resource.
//
// ## Import
//
// RDS DB Proxy Targets can be imported using the `db_proxy_name`, `target_group_name`, target type (e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`), and resource identifier separated by forward slashes (`/`), e.g. Instances
//
// ```sh
//  $ pulumi import aws:rds/proxyTarget:ProxyTarget example example-proxy/default/RDS_INSTANCE/example-instance
// ```
//
//  Provisioned Clusters
//
// ```sh
//  $ pulumi import aws:rds/proxyTarget:ProxyTarget example example-proxy/default/TRACKED_CLUSTER/example-cluster
// ```
type ProxyTarget struct {
	pulumi.CustomResourceState

	// DB cluster identifier.
	DbClusterIdentifier pulumi.StringPtrOutput `pulumi:"dbClusterIdentifier"`
	// DB instance identifier.
	DbInstanceIdentifier pulumi.StringPtrOutput `pulumi:"dbInstanceIdentifier"`
	// The name of the DB proxy.
	DbProxyName pulumi.StringOutput `pulumi:"dbProxyName"`
	// Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Port for the target RDS DB Instance or Aurora DB Cluster.
	Port pulumi.IntOutput `pulumi:"port"`
	// Identifier representing the DB Instance or DB Cluster target.
	RdsResourceId pulumi.StringOutput `pulumi:"rdsResourceId"`
	// Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
	// The name of the target group.
	TargetGroupName pulumi.StringOutput `pulumi:"targetGroupName"`
	// DB Cluster identifier for the DB Instance target. Not returned unless manually importing an `RDS_INSTANCE` target that is part of a DB Cluster.
	TrackedClusterId pulumi.StringOutput `pulumi:"trackedClusterId"`
	// Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProxyTarget registers a new resource with the given unique name, arguments, and options.
func NewProxyTarget(ctx *pulumi.Context,
	name string, args *ProxyTargetArgs, opts ...pulumi.ResourceOption) (*ProxyTarget, error) {
	if args == nil || args.DbProxyName == nil {
		return nil, errors.New("missing required argument 'DbProxyName'")
	}
	if args == nil || args.TargetGroupName == nil {
		return nil, errors.New("missing required argument 'TargetGroupName'")
	}
	if args == nil {
		args = &ProxyTargetArgs{}
	}
	var resource ProxyTarget
	err := ctx.RegisterResource("aws:rds/proxyTarget:ProxyTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxyTarget gets an existing ProxyTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxyTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyTargetState, opts ...pulumi.ResourceOption) (*ProxyTarget, error) {
	var resource ProxyTarget
	err := ctx.ReadResource("aws:rds/proxyTarget:ProxyTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProxyTarget resources.
type proxyTargetState struct {
	// DB cluster identifier.
	DbClusterIdentifier *string `pulumi:"dbClusterIdentifier"`
	// DB instance identifier.
	DbInstanceIdentifier *string `pulumi:"dbInstanceIdentifier"`
	// The name of the DB proxy.
	DbProxyName *string `pulumi:"dbProxyName"`
	// Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type.
	Endpoint *string `pulumi:"endpoint"`
	// Port for the target RDS DB Instance or Aurora DB Cluster.
	Port *int `pulumi:"port"`
	// Identifier representing the DB Instance or DB Cluster target.
	RdsResourceId *string `pulumi:"rdsResourceId"`
	// Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API.
	TargetArn *string `pulumi:"targetArn"`
	// The name of the target group.
	TargetGroupName *string `pulumi:"targetGroupName"`
	// DB Cluster identifier for the DB Instance target. Not returned unless manually importing an `RDS_INSTANCE` target that is part of a DB Cluster.
	TrackedClusterId *string `pulumi:"trackedClusterId"`
	// Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`
	Type *string `pulumi:"type"`
}

type ProxyTargetState struct {
	// DB cluster identifier.
	DbClusterIdentifier pulumi.StringPtrInput
	// DB instance identifier.
	DbInstanceIdentifier pulumi.StringPtrInput
	// The name of the DB proxy.
	DbProxyName pulumi.StringPtrInput
	// Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type.
	Endpoint pulumi.StringPtrInput
	// Port for the target RDS DB Instance or Aurora DB Cluster.
	Port pulumi.IntPtrInput
	// Identifier representing the DB Instance or DB Cluster target.
	RdsResourceId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API.
	TargetArn pulumi.StringPtrInput
	// The name of the target group.
	TargetGroupName pulumi.StringPtrInput
	// DB Cluster identifier for the DB Instance target. Not returned unless manually importing an `RDS_INSTANCE` target that is part of a DB Cluster.
	TrackedClusterId pulumi.StringPtrInput
	// Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`
	Type pulumi.StringPtrInput
}

func (ProxyTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyTargetState)(nil)).Elem()
}

type proxyTargetArgs struct {
	// DB cluster identifier.
	DbClusterIdentifier *string `pulumi:"dbClusterIdentifier"`
	// DB instance identifier.
	DbInstanceIdentifier *string `pulumi:"dbInstanceIdentifier"`
	// The name of the DB proxy.
	DbProxyName string `pulumi:"dbProxyName"`
	// The name of the target group.
	TargetGroupName string `pulumi:"targetGroupName"`
}

// The set of arguments for constructing a ProxyTarget resource.
type ProxyTargetArgs struct {
	// DB cluster identifier.
	DbClusterIdentifier pulumi.StringPtrInput
	// DB instance identifier.
	DbInstanceIdentifier pulumi.StringPtrInput
	// The name of the DB proxy.
	DbProxyName pulumi.StringInput
	// The name of the target group.
	TargetGroupName pulumi.StringInput
}

func (ProxyTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyTargetArgs)(nil)).Elem()
}

type ProxyTargetInput interface {
	pulumi.Input

	ToProxyTargetOutput() ProxyTargetOutput
	ToProxyTargetOutputWithContext(ctx context.Context) ProxyTargetOutput
}

func (ProxyTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyTarget)(nil)).Elem()
}

func (i ProxyTarget) ToProxyTargetOutput() ProxyTargetOutput {
	return i.ToProxyTargetOutputWithContext(context.Background())
}

func (i ProxyTarget) ToProxyTargetOutputWithContext(ctx context.Context) ProxyTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyTargetOutput)
}

type ProxyTargetOutput struct {
	*pulumi.OutputState
}

func (ProxyTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyTargetOutput)(nil)).Elem()
}

func (o ProxyTargetOutput) ToProxyTargetOutput() ProxyTargetOutput {
	return o
}

func (o ProxyTargetOutput) ToProxyTargetOutputWithContext(ctx context.Context) ProxyTargetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProxyTargetOutput{})
}

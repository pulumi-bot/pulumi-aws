// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeDeploy deployment config for an application
type DeploymentConfig struct {
	pulumi.CustomResourceState

	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringPtrOutput `pulumi:"computePlatform"`
	// The AWS Assigned deployment config id
	DeploymentConfigId pulumi.StringOutput `pulumi:"deploymentConfigId"`
	// The name of the deployment config.
	DeploymentConfigName pulumi.StringOutput `pulumi:"deploymentConfigName"`
	// A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsPtrOutput `pulumi:"minimumHealthyHosts"`
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigPtrOutput `pulumi:"trafficRoutingConfig"`
}

// NewDeploymentConfig registers a new resource with the given unique name, arguments, and options.
func NewDeploymentConfig(ctx *pulumi.Context,
	name string, args *DeploymentConfigArgs, opts ...pulumi.ResourceOption) (*DeploymentConfig, error) {
	if args == nil || args.DeploymentConfigName == nil {
		return nil, errors.New("missing required argument 'DeploymentConfigName'")
	}
	if args == nil {
		args = &DeploymentConfigArgs{}
	}
	var resource DeploymentConfig
	err := ctx.RegisterResource("aws:codedeploy/deploymentConfig:DeploymentConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentConfig gets an existing DeploymentConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentConfigState, opts ...pulumi.ResourceOption) (*DeploymentConfig, error) {
	var resource DeploymentConfig
	err := ctx.ReadResource("aws:codedeploy/deploymentConfig:DeploymentConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentConfig resources.
type deploymentConfigState struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform *string `pulumi:"computePlatform"`
	// The AWS Assigned deployment config id
	DeploymentConfigId *string `pulumi:"deploymentConfigId"`
	// The name of the deployment config.
	DeploymentConfigName *string `pulumi:"deploymentConfigName"`
	// A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts *DeploymentConfigMinimumHealthyHosts `pulumi:"minimumHealthyHosts"`
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig *DeploymentConfigTrafficRoutingConfig `pulumi:"trafficRoutingConfig"`
}

type DeploymentConfigState struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringPtrInput
	// The AWS Assigned deployment config id
	DeploymentConfigId pulumi.StringPtrInput
	// The name of the deployment config.
	DeploymentConfigName pulumi.StringPtrInput
	// A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsPtrInput
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigPtrInput
}

func (DeploymentConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentConfigState)(nil)).Elem()
}

type deploymentConfigArgs struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform *string `pulumi:"computePlatform"`
	// The name of the deployment config.
	DeploymentConfigName string `pulumi:"deploymentConfigName"`
	// A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts *DeploymentConfigMinimumHealthyHostsArgs `pulumi:"minimumHealthyHosts"`
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig *DeploymentConfigTrafficRoutingConfigArgs `pulumi:"trafficRoutingConfig"`
}

// The set of arguments for constructing a DeploymentConfig resource.
type DeploymentConfigArgs struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringPtrInput
	// The name of the deployment config.
	DeploymentConfigName pulumi.StringInput
	// A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsArgsPtrInput
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigArgsPtrInput
}

func (DeploymentConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentConfigArgs)(nil)).Elem()
}

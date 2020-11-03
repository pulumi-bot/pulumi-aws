// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeDeploy deployment config for an application
//
// ## Example Usage
// ### Server Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooDeploymentConfig, err := codedeploy.NewDeploymentConfig(ctx, "fooDeploymentConfig", &codedeploy.DeploymentConfigArgs{
// 			DeploymentConfigName: pulumi.String("test-deployment-config"),
// 			MinimumHealthyHosts: &codedeploy.DeploymentConfigMinimumHealthyHostsArgs{
// 				Type:  pulumi.String("HOST_COUNT"),
// 				Value: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codedeploy.NewDeploymentGroup(ctx, "fooDeploymentGroup", &codedeploy.DeploymentGroupArgs{
// 			AppName:              pulumi.Any(aws_codedeploy_app.Foo_app.Name),
// 			DeploymentGroupName:  pulumi.String("bar"),
// 			ServiceRoleArn:       pulumi.Any(aws_iam_role.Foo_role.Arn),
// 			DeploymentConfigName: fooDeploymentConfig.ID(),
// 			Ec2TagFilters: codedeploy.DeploymentGroupEc2TagFilterArray{
// 				&codedeploy.DeploymentGroupEc2TagFilterArgs{
// 					Key:   pulumi.String("filterkey"),
// 					Type:  pulumi.String("KEY_AND_VALUE"),
// 					Value: pulumi.String("filtervalue"),
// 				},
// 			},
// 			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
// 				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
// 					TriggerEvents: pulumi.StringArray{
// 						pulumi.String("DeploymentFailure"),
// 					},
// 					TriggerName:      pulumi.String("foo-trigger"),
// 					TriggerTargetArn: pulumi.String("foo-topic-arn"),
// 				},
// 			},
// 			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
// 				Enabled: pulumi.Bool(true),
// 				Events: pulumi.StringArray{
// 					pulumi.String("DEPLOYMENT_FAILURE"),
// 				},
// 			},
// 			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
// 				Alarms: pulumi.StringArray{
// 					pulumi.String("my-alarm-name"),
// 				},
// 				Enabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Lambda Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooDeploymentConfig, err := codedeploy.NewDeploymentConfig(ctx, "fooDeploymentConfig", &codedeploy.DeploymentConfigArgs{
// 			DeploymentConfigName: pulumi.String("test-deployment-config"),
// 			ComputePlatform:      pulumi.String("Lambda"),
// 			TrafficRoutingConfig: &codedeploy.DeploymentConfigTrafficRoutingConfigArgs{
// 				Type: pulumi.String("TimeBasedLinear"),
// 				TimeBasedLinear: &codedeploy.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs{
// 					Interval:   pulumi.Int(10),
// 					Percentage: pulumi.Int(10),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codedeploy.NewDeploymentGroup(ctx, "fooDeploymentGroup", &codedeploy.DeploymentGroupArgs{
// 			AppName:              pulumi.Any(aws_codedeploy_app.Foo_app.Name),
// 			DeploymentGroupName:  pulumi.String("bar"),
// 			ServiceRoleArn:       pulumi.Any(aws_iam_role.Foo_role.Arn),
// 			DeploymentConfigName: fooDeploymentConfig.ID(),
// 			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
// 				Enabled: pulumi.Bool(true),
// 				Events: pulumi.StringArray{
// 					pulumi.String("DEPLOYMENT_STOP_ON_ALARM"),
// 				},
// 			},
// 			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
// 				Alarms: pulumi.StringArray{
// 					pulumi.String("my-alarm-name"),
// 				},
// 				Enabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DeploymentConfigName == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentConfigName'")
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
	MinimumHealthyHosts *DeploymentConfigMinimumHealthyHosts `pulumi:"minimumHealthyHosts"`
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig *DeploymentConfigTrafficRoutingConfig `pulumi:"trafficRoutingConfig"`
}

// The set of arguments for constructing a DeploymentConfig resource.
type DeploymentConfigArgs struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringPtrInput
	// The name of the deployment config.
	DeploymentConfigName pulumi.StringInput
	// A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsPtrInput
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigPtrInput
}

func (DeploymentConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentConfigArgs)(nil)).Elem()
}

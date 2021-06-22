// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeDeploy Deployment Group for a CodeDeploy Application
//
// > **NOTE on blue/green deployments:** When using `greenFleetProvisioningOption` with the `COPY_AUTO_SCALING_GROUP` action, CodeDeploy will create a new ASG with a different name. This ASG is _not_ managed by this provider and will conflict with existing configuration and state. You may want to use a different approach to managing deployments that involve multiple ASG, such as `DISCOVER_EXISTING` with separate blue and green ASG.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"codedeploy.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "aWSCodeDeployRole", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"),
// 			Role:      exampleRole.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := sns.NewTopic(ctx, "exampleTopic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
// 			AppName:             exampleApplication.Name,
// 			DeploymentGroupName: pulumi.String("example-group"),
// 			ServiceRoleArn:      exampleRole.Arn,
// 			Ec2TagSets: codedeploy.DeploymentGroupEc2TagSetArray{
// 				&codedeploy.DeploymentGroupEc2TagSetArgs{
// 					Ec2TagFilters: codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArray{
// 						&codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArgs{
// 							Key:   pulumi.String("filterkey1"),
// 							Type:  pulumi.String("KEY_AND_VALUE"),
// 							Value: pulumi.String("filtervalue"),
// 						},
// 						&codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArgs{
// 							Key:   pulumi.String("filterkey2"),
// 							Type:  pulumi.String("KEY_AND_VALUE"),
// 							Value: pulumi.String("filtervalue"),
// 						},
// 					},
// 				},
// 			},
// 			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
// 				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
// 					TriggerEvents: pulumi.StringArray{
// 						pulumi.String("DeploymentFailure"),
// 					},
// 					TriggerName:      pulumi.String("example-trigger"),
// 					TriggerTargetArn: exampleTopic.Arn,
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
// ### Blue Green Deployments with ECS
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", &codedeploy.ApplicationArgs{
// 			ComputePlatform: pulumi.String("ECS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
// 			AppName:              exampleApplication.Name,
// 			DeploymentConfigName: pulumi.String("CodeDeployDefault.ECSAllAtOnce"),
// 			DeploymentGroupName:  pulumi.String("example"),
// 			ServiceRoleArn:       pulumi.Any(aws_iam_role.Example.Arn),
// 			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
// 				Enabled: pulumi.Bool(true),
// 				Events: pulumi.StringArray{
// 					pulumi.String("DEPLOYMENT_FAILURE"),
// 				},
// 			},
// 			BlueGreenDeploymentConfig: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs{
// 				DeploymentReadyOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs{
// 					ActionOnTimeout: pulumi.String("CONTINUE_DEPLOYMENT"),
// 				},
// 				TerminateBlueInstancesOnDeploymentSuccess: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs{
// 					Action:                       pulumi.String("TERMINATE"),
// 					TerminationWaitTimeInMinutes: pulumi.Int(5),
// 				},
// 			},
// 			DeploymentStyle: &codedeploy.DeploymentGroupDeploymentStyleArgs{
// 				DeploymentOption: pulumi.String("WITH_TRAFFIC_CONTROL"),
// 				DeploymentType:   pulumi.String("BLUE_GREEN"),
// 			},
// 			EcsService: &codedeploy.DeploymentGroupEcsServiceArgs{
// 				ClusterName: pulumi.Any(aws_ecs_cluster.Example.Name),
// 				ServiceName: pulumi.Any(aws_ecs_service.Example.Name),
// 			},
// 			LoadBalancerInfo: &codedeploy.DeploymentGroupLoadBalancerInfoArgs{
// 				TargetGroupPairInfo: &codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs{
// 					ProdTrafficRoute: &codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs{
// 						ListenerArns: pulumi.StringArray{
// 							pulumi.Any(aws_lb_listener.Example.Arn),
// 						},
// 					},
// 					TargetGroups: codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArray{
// 						&codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs{
// 							Name: pulumi.Any(aws_lb_target_group.Blue.Name),
// 						},
// 						&codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs{
// 							Name: pulumi.Any(aws_lb_target_group.Green.Name),
// 						},
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
// ### Blue Green Deployments with Servers and Classic ELB
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
// 			AppName:             exampleApplication.Name,
// 			DeploymentGroupName: pulumi.String("example-group"),
// 			ServiceRoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
// 			DeploymentStyle: &codedeploy.DeploymentGroupDeploymentStyleArgs{
// 				DeploymentOption: pulumi.String("WITH_TRAFFIC_CONTROL"),
// 				DeploymentType:   pulumi.String("BLUE_GREEN"),
// 			},
// 			LoadBalancerInfo: &codedeploy.DeploymentGroupLoadBalancerInfoArgs{
// 				ElbInfos: codedeploy.DeploymentGroupLoadBalancerInfoElbInfoArray{
// 					&codedeploy.DeploymentGroupLoadBalancerInfoElbInfoArgs{
// 						Name: pulumi.Any(aws_elb.Example.Name),
// 					},
// 				},
// 			},
// 			BlueGreenDeploymentConfig: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs{
// 				DeploymentReadyOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs{
// 					ActionOnTimeout:   pulumi.String("STOP_DEPLOYMENT"),
// 					WaitTimeInMinutes: pulumi.Int(60),
// 				},
// 				GreenFleetProvisioningOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs{
// 					Action: pulumi.String("DISCOVER_EXISTING"),
// 				},
// 				TerminateBlueInstancesOnDeploymentSuccess: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs{
// 					Action: pulumi.String("KEEP_ALIVE"),
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
// CodeDeploy Deployment Groups can be imported by their `app_name`, a colon, and `deployment_group_name`, e.g.
//
// ```sh
//  $ pulumi import aws:codedeploy/deploymentGroup:DeploymentGroup example my-application:my-deployment-group
// ```
//
//  [1]http://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-sns-event-notifications-create-trigger.html
type DeploymentGroup struct {
	pulumi.CustomResourceState

	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration DeploymentGroupAlarmConfigurationPtrOutput `pulumi:"alarmConfiguration"`
	// The name of the application.
	AppName pulumi.StringOutput `pulumi:"appName"`
	// The ARN of the CodeDeploy deployment group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration DeploymentGroupAutoRollbackConfigurationPtrOutput `pulumi:"autoRollbackConfiguration"`
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups pulumi.StringArrayOutput `pulumi:"autoscalingGroups"`
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig DeploymentGroupBlueGreenDeploymentConfigOutput `pulumi:"blueGreenDeploymentConfig"`
	// The destination platform type for the deployment.
	ComputePlatform pulumi.StringOutput `pulumi:"computePlatform"`
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName pulumi.StringPtrOutput `pulumi:"deploymentConfigName"`
	// The ID of the CodeDeploy deployment group.
	DeploymentGroupId pulumi.StringOutput `pulumi:"deploymentGroupId"`
	// The name of the deployment group.
	DeploymentGroupName pulumi.StringOutput `pulumi:"deploymentGroupName"`
	// Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
	DeploymentStyle DeploymentGroupDeploymentStylePtrOutput `pulumi:"deploymentStyle"`
	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters DeploymentGroupEc2TagFilterArrayOutput `pulumi:"ec2TagFilters"`
	// Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
	Ec2TagSets DeploymentGroupEc2TagSetArrayOutput `pulumi:"ec2TagSets"`
	// Configuration block(s) of the ECS services for a deployment group (documented below).
	EcsService DeploymentGroupEcsServicePtrOutput `pulumi:"ecsService"`
	// Single configuration block of the load balancer to use in a blue/green deployment (documented below).
	LoadBalancerInfo DeploymentGroupLoadBalancerInfoPtrOutput `pulumi:"loadBalancerInfo"`
	// On premise tag filters associated with the group. See the AWS docs for details.
	OnPremisesInstanceTagFilters DeploymentGroupOnPremisesInstanceTagFilterArrayOutput `pulumi:"onPremisesInstanceTagFilters"`
	// The service role ARN that allows deployments.
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations DeploymentGroupTriggerConfigurationArrayOutput `pulumi:"triggerConfigurations"`
}

// NewDeploymentGroup registers a new resource with the given unique name, arguments, and options.
func NewDeploymentGroup(ctx *pulumi.Context,
	name string, args *DeploymentGroupArgs, opts ...pulumi.ResourceOption) (*DeploymentGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.DeploymentGroupName == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentGroupName'")
	}
	if args.ServiceRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ServiceRoleArn'")
	}
	var resource DeploymentGroup
	err := ctx.RegisterResource("aws:codedeploy/deploymentGroup:DeploymentGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentGroup gets an existing DeploymentGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentGroupState, opts ...pulumi.ResourceOption) (*DeploymentGroup, error) {
	var resource DeploymentGroup
	err := ctx.ReadResource("aws:codedeploy/deploymentGroup:DeploymentGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentGroup resources.
type deploymentGroupState struct {
	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration *DeploymentGroupAlarmConfiguration `pulumi:"alarmConfiguration"`
	// The name of the application.
	AppName *string `pulumi:"appName"`
	// The ARN of the CodeDeploy deployment group.
	Arn *string `pulumi:"arn"`
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration *DeploymentGroupAutoRollbackConfiguration `pulumi:"autoRollbackConfiguration"`
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups []string `pulumi:"autoscalingGroups"`
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig *DeploymentGroupBlueGreenDeploymentConfig `pulumi:"blueGreenDeploymentConfig"`
	// The destination platform type for the deployment.
	ComputePlatform *string `pulumi:"computePlatform"`
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName *string `pulumi:"deploymentConfigName"`
	// The ID of the CodeDeploy deployment group.
	DeploymentGroupId *string `pulumi:"deploymentGroupId"`
	// The name of the deployment group.
	DeploymentGroupName *string `pulumi:"deploymentGroupName"`
	// Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
	DeploymentStyle *DeploymentGroupDeploymentStyle `pulumi:"deploymentStyle"`
	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters []DeploymentGroupEc2TagFilter `pulumi:"ec2TagFilters"`
	// Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
	Ec2TagSets []DeploymentGroupEc2TagSet `pulumi:"ec2TagSets"`
	// Configuration block(s) of the ECS services for a deployment group (documented below).
	EcsService *DeploymentGroupEcsService `pulumi:"ecsService"`
	// Single configuration block of the load balancer to use in a blue/green deployment (documented below).
	LoadBalancerInfo *DeploymentGroupLoadBalancerInfo `pulumi:"loadBalancerInfo"`
	// On premise tag filters associated with the group. See the AWS docs for details.
	OnPremisesInstanceTagFilters []DeploymentGroupOnPremisesInstanceTagFilter `pulumi:"onPremisesInstanceTagFilters"`
	// The service role ARN that allows deployments.
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations []DeploymentGroupTriggerConfiguration `pulumi:"triggerConfigurations"`
}

type DeploymentGroupState struct {
	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration DeploymentGroupAlarmConfigurationPtrInput
	// The name of the application.
	AppName pulumi.StringPtrInput
	// The ARN of the CodeDeploy deployment group.
	Arn pulumi.StringPtrInput
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration DeploymentGroupAutoRollbackConfigurationPtrInput
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups pulumi.StringArrayInput
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig DeploymentGroupBlueGreenDeploymentConfigPtrInput
	// The destination platform type for the deployment.
	ComputePlatform pulumi.StringPtrInput
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName pulumi.StringPtrInput
	// The ID of the CodeDeploy deployment group.
	DeploymentGroupId pulumi.StringPtrInput
	// The name of the deployment group.
	DeploymentGroupName pulumi.StringPtrInput
	// Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
	DeploymentStyle DeploymentGroupDeploymentStylePtrInput
	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters DeploymentGroupEc2TagFilterArrayInput
	// Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
	Ec2TagSets DeploymentGroupEc2TagSetArrayInput
	// Configuration block(s) of the ECS services for a deployment group (documented below).
	EcsService DeploymentGroupEcsServicePtrInput
	// Single configuration block of the load balancer to use in a blue/green deployment (documented below).
	LoadBalancerInfo DeploymentGroupLoadBalancerInfoPtrInput
	// On premise tag filters associated with the group. See the AWS docs for details.
	OnPremisesInstanceTagFilters DeploymentGroupOnPremisesInstanceTagFilterArrayInput
	// The service role ARN that allows deployments.
	ServiceRoleArn pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations DeploymentGroupTriggerConfigurationArrayInput
}

func (DeploymentGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentGroupState)(nil)).Elem()
}

type deploymentGroupArgs struct {
	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration *DeploymentGroupAlarmConfiguration `pulumi:"alarmConfiguration"`
	// The name of the application.
	AppName string `pulumi:"appName"`
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration *DeploymentGroupAutoRollbackConfiguration `pulumi:"autoRollbackConfiguration"`
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups []string `pulumi:"autoscalingGroups"`
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig *DeploymentGroupBlueGreenDeploymentConfig `pulumi:"blueGreenDeploymentConfig"`
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName *string `pulumi:"deploymentConfigName"`
	// The name of the deployment group.
	DeploymentGroupName string `pulumi:"deploymentGroupName"`
	// Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
	DeploymentStyle *DeploymentGroupDeploymentStyle `pulumi:"deploymentStyle"`
	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters []DeploymentGroupEc2TagFilter `pulumi:"ec2TagFilters"`
	// Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
	Ec2TagSets []DeploymentGroupEc2TagSet `pulumi:"ec2TagSets"`
	// Configuration block(s) of the ECS services for a deployment group (documented below).
	EcsService *DeploymentGroupEcsService `pulumi:"ecsService"`
	// Single configuration block of the load balancer to use in a blue/green deployment (documented below).
	LoadBalancerInfo *DeploymentGroupLoadBalancerInfo `pulumi:"loadBalancerInfo"`
	// On premise tag filters associated with the group. See the AWS docs for details.
	OnPremisesInstanceTagFilters []DeploymentGroupOnPremisesInstanceTagFilter `pulumi:"onPremisesInstanceTagFilters"`
	// The service role ARN that allows deployments.
	ServiceRoleArn string `pulumi:"serviceRoleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations []DeploymentGroupTriggerConfiguration `pulumi:"triggerConfigurations"`
}

// The set of arguments for constructing a DeploymentGroup resource.
type DeploymentGroupArgs struct {
	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration DeploymentGroupAlarmConfigurationPtrInput
	// The name of the application.
	AppName pulumi.StringInput
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration DeploymentGroupAutoRollbackConfigurationPtrInput
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups pulumi.StringArrayInput
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig DeploymentGroupBlueGreenDeploymentConfigPtrInput
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName pulumi.StringPtrInput
	// The name of the deployment group.
	DeploymentGroupName pulumi.StringInput
	// Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
	DeploymentStyle DeploymentGroupDeploymentStylePtrInput
	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters DeploymentGroupEc2TagFilterArrayInput
	// Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
	Ec2TagSets DeploymentGroupEc2TagSetArrayInput
	// Configuration block(s) of the ECS services for a deployment group (documented below).
	EcsService DeploymentGroupEcsServicePtrInput
	// Single configuration block of the load balancer to use in a blue/green deployment (documented below).
	LoadBalancerInfo DeploymentGroupLoadBalancerInfoPtrInput
	// On premise tag filters associated with the group. See the AWS docs for details.
	OnPremisesInstanceTagFilters DeploymentGroupOnPremisesInstanceTagFilterArrayInput
	// The service role ARN that allows deployments.
	ServiceRoleArn pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations DeploymentGroupTriggerConfigurationArrayInput
}

func (DeploymentGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentGroupArgs)(nil)).Elem()
}

type DeploymentGroupInput interface {
	pulumi.Input

	ToDeploymentGroupOutput() DeploymentGroupOutput
	ToDeploymentGroupOutputWithContext(ctx context.Context) DeploymentGroupOutput
}

func (*DeploymentGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroup)(nil))
}

func (i *DeploymentGroup) ToDeploymentGroupOutput() DeploymentGroupOutput {
	return i.ToDeploymentGroupOutputWithContext(context.Background())
}

func (i *DeploymentGroup) ToDeploymentGroupOutputWithContext(ctx context.Context) DeploymentGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupOutput)
}

func (i *DeploymentGroup) ToDeploymentGroupPtrOutput() DeploymentGroupPtrOutput {
	return i.ToDeploymentGroupPtrOutputWithContext(context.Background())
}

func (i *DeploymentGroup) ToDeploymentGroupPtrOutputWithContext(ctx context.Context) DeploymentGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupPtrOutput)
}

type DeploymentGroupPtrInput interface {
	pulumi.Input

	ToDeploymentGroupPtrOutput() DeploymentGroupPtrOutput
	ToDeploymentGroupPtrOutputWithContext(ctx context.Context) DeploymentGroupPtrOutput
}

type deploymentGroupPtrType DeploymentGroupArgs

func (*deploymentGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentGroup)(nil))
}

func (i *deploymentGroupPtrType) ToDeploymentGroupPtrOutput() DeploymentGroupPtrOutput {
	return i.ToDeploymentGroupPtrOutputWithContext(context.Background())
}

func (i *deploymentGroupPtrType) ToDeploymentGroupPtrOutputWithContext(ctx context.Context) DeploymentGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupPtrOutput)
}

// DeploymentGroupArrayInput is an input type that accepts DeploymentGroupArray and DeploymentGroupArrayOutput values.
// You can construct a concrete instance of `DeploymentGroupArrayInput` via:
//
//          DeploymentGroupArray{ DeploymentGroupArgs{...} }
type DeploymentGroupArrayInput interface {
	pulumi.Input

	ToDeploymentGroupArrayOutput() DeploymentGroupArrayOutput
	ToDeploymentGroupArrayOutputWithContext(context.Context) DeploymentGroupArrayOutput
}

type DeploymentGroupArray []DeploymentGroupInput

func (DeploymentGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DeploymentGroup)(nil))
}

func (i DeploymentGroupArray) ToDeploymentGroupArrayOutput() DeploymentGroupArrayOutput {
	return i.ToDeploymentGroupArrayOutputWithContext(context.Background())
}

func (i DeploymentGroupArray) ToDeploymentGroupArrayOutputWithContext(ctx context.Context) DeploymentGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupArrayOutput)
}

// DeploymentGroupMapInput is an input type that accepts DeploymentGroupMap and DeploymentGroupMapOutput values.
// You can construct a concrete instance of `DeploymentGroupMapInput` via:
//
//          DeploymentGroupMap{ "key": DeploymentGroupArgs{...} }
type DeploymentGroupMapInput interface {
	pulumi.Input

	ToDeploymentGroupMapOutput() DeploymentGroupMapOutput
	ToDeploymentGroupMapOutputWithContext(context.Context) DeploymentGroupMapOutput
}

type DeploymentGroupMap map[string]DeploymentGroupInput

func (DeploymentGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DeploymentGroup)(nil))
}

func (i DeploymentGroupMap) ToDeploymentGroupMapOutput() DeploymentGroupMapOutput {
	return i.ToDeploymentGroupMapOutputWithContext(context.Background())
}

func (i DeploymentGroupMap) ToDeploymentGroupMapOutputWithContext(ctx context.Context) DeploymentGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupMapOutput)
}

type DeploymentGroupOutput struct {
	*pulumi.OutputState
}

func (DeploymentGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroup)(nil))
}

func (o DeploymentGroupOutput) ToDeploymentGroupOutput() DeploymentGroupOutput {
	return o
}

func (o DeploymentGroupOutput) ToDeploymentGroupOutputWithContext(ctx context.Context) DeploymentGroupOutput {
	return o
}

func (o DeploymentGroupOutput) ToDeploymentGroupPtrOutput() DeploymentGroupPtrOutput {
	return o.ToDeploymentGroupPtrOutputWithContext(context.Background())
}

func (o DeploymentGroupOutput) ToDeploymentGroupPtrOutputWithContext(ctx context.Context) DeploymentGroupPtrOutput {
	return o.ApplyT(func(v DeploymentGroup) *DeploymentGroup {
		return &v
	}).(DeploymentGroupPtrOutput)
}

type DeploymentGroupPtrOutput struct {
	*pulumi.OutputState
}

func (DeploymentGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentGroup)(nil))
}

func (o DeploymentGroupPtrOutput) ToDeploymentGroupPtrOutput() DeploymentGroupPtrOutput {
	return o
}

func (o DeploymentGroupPtrOutput) ToDeploymentGroupPtrOutputWithContext(ctx context.Context) DeploymentGroupPtrOutput {
	return o
}

type DeploymentGroupArrayOutput struct{ *pulumi.OutputState }

func (DeploymentGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentGroup)(nil))
}

func (o DeploymentGroupArrayOutput) ToDeploymentGroupArrayOutput() DeploymentGroupArrayOutput {
	return o
}

func (o DeploymentGroupArrayOutput) ToDeploymentGroupArrayOutputWithContext(ctx context.Context) DeploymentGroupArrayOutput {
	return o
}

func (o DeploymentGroupArrayOutput) Index(i pulumi.IntInput) DeploymentGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentGroup {
		return vs[0].([]DeploymentGroup)[vs[1].(int)]
	}).(DeploymentGroupOutput)
}

type DeploymentGroupMapOutput struct{ *pulumi.OutputState }

func (DeploymentGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeploymentGroup)(nil))
}

func (o DeploymentGroupMapOutput) ToDeploymentGroupMapOutput() DeploymentGroupMapOutput {
	return o
}

func (o DeploymentGroupMapOutput) ToDeploymentGroupMapOutputWithContext(ctx context.Context) DeploymentGroupMapOutput {
	return o
}

func (o DeploymentGroupMapOutput) MapIndex(k pulumi.StringInput) DeploymentGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeploymentGroup {
		return vs[0].(map[string]DeploymentGroup)[vs[1].(string)]
	}).(DeploymentGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentGroupOutput{})
	pulumi.RegisterOutputType(DeploymentGroupPtrOutput{})
	pulumi.RegisterOutputType(DeploymentGroupArrayOutput{})
	pulumi.RegisterOutputType(DeploymentGroupMapOutput{})
}

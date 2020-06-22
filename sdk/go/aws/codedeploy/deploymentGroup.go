// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"codedeploy.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
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
// 			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
// 				Alarms: pulumi.StringArray{
// 					pulumi.String("my-alarm-name"),
// 				},
// 				Enabled: pulumi.Bool(true),
// 			},
// 			AppName: exampleApplication.Name,
// 			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
// 				Enabled: pulumi.Bool(true),
// 				Events: pulumi.StringArray{
// 					pulumi.String("DEPLOYMENT_FAILURE"),
// 				},
// 			},
// 			DeploymentGroupName: pulumi.String("example-group"),
// 			Ec2TagSets: codedeploy.DeploymentGroupEc2TagSetArray{
// 				&codedeploy.DeploymentGroupEc2TagSetArgs{
// 					Ec2TagFilter: pulumi.MapArray{
// 						pulumi.Map{
// 							"key":   pulumi.String("filterkey1"),
// 							"type":  pulumi.String("KEY_AND_VALUE"),
// 							"value": pulumi.String("filtervalue"),
// 						},
// 						pulumi.Map{
// 							"key":   pulumi.String("filterkey2"),
// 							"type":  pulumi.String("KEY_AND_VALUE"),
// 							"value": pulumi.String("filtervalue"),
// 						},
// 					},
// 				},
// 			},
// 			ServiceRoleArn: exampleRole.Arn,
// 			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
// 				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
// 					TriggerEvents: pulumi.StringArray{
// 						pulumi.String("DeploymentFailure"),
// 					},
// 					TriggerName:      pulumi.String("example-trigger"),
// 					TriggerTargetArn: exampleTopic.Arn,
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
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
// 			AppName: exampleApplication.Name,
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
// 			DeploymentGroupName: pulumi.String("example-group"),
// 			DeploymentStyle: &codedeploy.DeploymentGroupDeploymentStyleArgs{
// 				DeploymentOption: pulumi.String("WITH_TRAFFIC_CONTROL"),
// 				DeploymentType:   pulumi.String("BLUE_GREEN"),
// 			},
// 			LoadBalancerInfo: &codedeploy.DeploymentGroupLoadBalancerInfoArgs{
// 				ElbInfo: pulumi.MapArray{
// 					pulumi.Map{
// 						"name": pulumi.String(aws_elb.Example.Name),
// 					},
// 				},
// 			},
// 			ServiceRoleArn: pulumi.String(aws_iam_role.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DeploymentGroup struct {
	pulumi.CustomResourceState

	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration DeploymentGroupAlarmConfigurationPtrOutput `pulumi:"alarmConfiguration"`
	// The name of the application.
	AppName pulumi.StringOutput `pulumi:"appName"`
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration DeploymentGroupAutoRollbackConfigurationPtrOutput `pulumi:"autoRollbackConfiguration"`
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups pulumi.StringArrayOutput `pulumi:"autoscalingGroups"`
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig DeploymentGroupBlueGreenDeploymentConfigOutput `pulumi:"blueGreenDeploymentConfig"`
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName pulumi.StringPtrOutput `pulumi:"deploymentConfigName"`
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
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations DeploymentGroupTriggerConfigurationArrayOutput `pulumi:"triggerConfigurations"`
}

// NewDeploymentGroup registers a new resource with the given unique name, arguments, and options.
func NewDeploymentGroup(ctx *pulumi.Context,
	name string, args *DeploymentGroupArgs, opts ...pulumi.ResourceOption) (*DeploymentGroup, error) {
	if args == nil || args.AppName == nil {
		return nil, errors.New("missing required argument 'AppName'")
	}
	if args == nil || args.DeploymentGroupName == nil {
		return nil, errors.New("missing required argument 'DeploymentGroupName'")
	}
	if args == nil || args.ServiceRoleArn == nil {
		return nil, errors.New("missing required argument 'ServiceRoleArn'")
	}
	if args == nil {
		args = &DeploymentGroupArgs{}
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
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration *DeploymentGroupAutoRollbackConfiguration `pulumi:"autoRollbackConfiguration"`
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups []string `pulumi:"autoscalingGroups"`
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig *DeploymentGroupBlueGreenDeploymentConfig `pulumi:"blueGreenDeploymentConfig"`
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName *string `pulumi:"deploymentConfigName"`
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
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations []DeploymentGroupTriggerConfiguration `pulumi:"triggerConfigurations"`
}

type DeploymentGroupState struct {
	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration DeploymentGroupAlarmConfigurationPtrInput
	// The name of the application.
	AppName pulumi.StringPtrInput
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration DeploymentGroupAutoRollbackConfigurationPtrInput
	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups pulumi.StringArrayInput
	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig DeploymentGroupBlueGreenDeploymentConfigPtrInput
	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName pulumi.StringPtrInput
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
	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations DeploymentGroupTriggerConfigurationArrayInput
}

func (DeploymentGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentGroupArgs)(nil)).Elem()
}

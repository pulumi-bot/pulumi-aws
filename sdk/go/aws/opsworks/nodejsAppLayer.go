// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type NodejsAppLayer struct {
	pulumi.CustomResourceState

	Arn                      pulumi.StringOutput                `pulumi:"arn"`
	AutoAssignElasticIps     pulumi.BoolPtrOutput               `pulumi:"autoAssignElasticIps"`
	AutoAssignPublicIps      pulumi.BoolPtrOutput               `pulumi:"autoAssignPublicIps"`
	AutoHealing              pulumi.BoolPtrOutput               `pulumi:"autoHealing"`
	CustomConfigureRecipes   pulumi.StringArrayOutput           `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes      pulumi.StringArrayOutput           `pulumi:"customDeployRecipes"`
	CustomInstanceProfileArn pulumi.StringPtrOutput             `pulumi:"customInstanceProfileArn"`
	CustomJson               pulumi.StringPtrOutput             `pulumi:"customJson"`
	CustomSecurityGroupIds   pulumi.StringArrayOutput           `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes       pulumi.StringArrayOutput           `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes    pulumi.StringArrayOutput           `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes    pulumi.StringArrayOutput           `pulumi:"customUndeployRecipes"`
	DrainElbOnShutdown       pulumi.BoolPtrOutput               `pulumi:"drainElbOnShutdown"`
	EbsVolumes               NodejsAppLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	ElasticLoadBalancer      pulumi.StringPtrOutput             `pulumi:"elasticLoadBalancer"`
	InstallUpdatesOnBoot     pulumi.BoolPtrOutput               `pulumi:"installUpdatesOnBoot"`
	InstanceShutdownTimeout  pulumi.IntPtrOutput                `pulumi:"instanceShutdownTimeout"`
	Name                     pulumi.StringOutput                `pulumi:"name"`
	NodejsVersion            pulumi.StringPtrOutput             `pulumi:"nodejsVersion"`
	StackId                  pulumi.StringOutput                `pulumi:"stackId"`
	SystemPackages           pulumi.StringArrayOutput           `pulumi:"systemPackages"`
	Tags                     pulumi.StringMapOutput             `pulumi:"tags"`
	UseEbsOptimizedInstances pulumi.BoolPtrOutput               `pulumi:"useEbsOptimizedInstances"`
}

// NewNodejsAppLayer registers a new resource with the given unique name, arguments, and options.
func NewNodejsAppLayer(ctx *pulumi.Context,
	name string, args *NodejsAppLayerArgs, opts ...pulumi.ResourceOption) (*NodejsAppLayer, error) {
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	if args == nil {
		args = &NodejsAppLayerArgs{}
	}
	var resource NodejsAppLayer
	err := ctx.RegisterResource("aws:opsworks/nodejsAppLayer:NodejsAppLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodejsAppLayer gets an existing NodejsAppLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodejsAppLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodejsAppLayerState, opts ...pulumi.ResourceOption) (*NodejsAppLayer, error) {
	var resource NodejsAppLayer
	err := ctx.ReadResource("aws:opsworks/nodejsAppLayer:NodejsAppLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodejsAppLayer resources.
type nodejsAppLayerState struct {
	Arn                      *string                   `pulumi:"arn"`
	AutoAssignElasticIps     *bool                     `pulumi:"autoAssignElasticIps"`
	AutoAssignPublicIps      *bool                     `pulumi:"autoAssignPublicIps"`
	AutoHealing              *bool                     `pulumi:"autoHealing"`
	CustomConfigureRecipes   []string                  `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes      []string                  `pulumi:"customDeployRecipes"`
	CustomInstanceProfileArn *string                   `pulumi:"customInstanceProfileArn"`
	CustomJson               *string                   `pulumi:"customJson"`
	CustomSecurityGroupIds   []string                  `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes       []string                  `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes    []string                  `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes    []string                  `pulumi:"customUndeployRecipes"`
	DrainElbOnShutdown       *bool                     `pulumi:"drainElbOnShutdown"`
	EbsVolumes               []NodejsAppLayerEbsVolume `pulumi:"ebsVolumes"`
	ElasticLoadBalancer      *string                   `pulumi:"elasticLoadBalancer"`
	InstallUpdatesOnBoot     *bool                     `pulumi:"installUpdatesOnBoot"`
	InstanceShutdownTimeout  *int                      `pulumi:"instanceShutdownTimeout"`
	Name                     *string                   `pulumi:"name"`
	NodejsVersion            *string                   `pulumi:"nodejsVersion"`
	StackId                  *string                   `pulumi:"stackId"`
	SystemPackages           []string                  `pulumi:"systemPackages"`
	Tags                     map[string]string         `pulumi:"tags"`
	UseEbsOptimizedInstances *bool                     `pulumi:"useEbsOptimizedInstances"`
}

type NodejsAppLayerState struct {
	Arn                      pulumi.StringPtrInput
	AutoAssignElasticIps     pulumi.BoolPtrInput
	AutoAssignPublicIps      pulumi.BoolPtrInput
	AutoHealing              pulumi.BoolPtrInput
	CustomConfigureRecipes   pulumi.StringArrayInput
	CustomDeployRecipes      pulumi.StringArrayInput
	CustomInstanceProfileArn pulumi.StringPtrInput
	CustomJson               pulumi.StringPtrInput
	CustomSecurityGroupIds   pulumi.StringArrayInput
	CustomSetupRecipes       pulumi.StringArrayInput
	CustomShutdownRecipes    pulumi.StringArrayInput
	CustomUndeployRecipes    pulumi.StringArrayInput
	DrainElbOnShutdown       pulumi.BoolPtrInput
	EbsVolumes               NodejsAppLayerEbsVolumeArrayInput
	ElasticLoadBalancer      pulumi.StringPtrInput
	InstallUpdatesOnBoot     pulumi.BoolPtrInput
	InstanceShutdownTimeout  pulumi.IntPtrInput
	Name                     pulumi.StringPtrInput
	NodejsVersion            pulumi.StringPtrInput
	StackId                  pulumi.StringPtrInput
	SystemPackages           pulumi.StringArrayInput
	Tags                     pulumi.StringMapInput
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (NodejsAppLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodejsAppLayerState)(nil)).Elem()
}

type nodejsAppLayerArgs struct {
	AutoAssignElasticIps     *bool                     `pulumi:"autoAssignElasticIps"`
	AutoAssignPublicIps      *bool                     `pulumi:"autoAssignPublicIps"`
	AutoHealing              *bool                     `pulumi:"autoHealing"`
	CustomConfigureRecipes   []string                  `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes      []string                  `pulumi:"customDeployRecipes"`
	CustomInstanceProfileArn *string                   `pulumi:"customInstanceProfileArn"`
	CustomJson               *string                   `pulumi:"customJson"`
	CustomSecurityGroupIds   []string                  `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes       []string                  `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes    []string                  `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes    []string                  `pulumi:"customUndeployRecipes"`
	DrainElbOnShutdown       *bool                     `pulumi:"drainElbOnShutdown"`
	EbsVolumes               []NodejsAppLayerEbsVolume `pulumi:"ebsVolumes"`
	ElasticLoadBalancer      *string                   `pulumi:"elasticLoadBalancer"`
	InstallUpdatesOnBoot     *bool                     `pulumi:"installUpdatesOnBoot"`
	InstanceShutdownTimeout  *int                      `pulumi:"instanceShutdownTimeout"`
	Name                     *string                   `pulumi:"name"`
	NodejsVersion            *string                   `pulumi:"nodejsVersion"`
	StackId                  string                    `pulumi:"stackId"`
	SystemPackages           []string                  `pulumi:"systemPackages"`
	Tags                     map[string]string         `pulumi:"tags"`
	UseEbsOptimizedInstances *bool                     `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a NodejsAppLayer resource.
type NodejsAppLayerArgs struct {
	AutoAssignElasticIps     pulumi.BoolPtrInput
	AutoAssignPublicIps      pulumi.BoolPtrInput
	AutoHealing              pulumi.BoolPtrInput
	CustomConfigureRecipes   pulumi.StringArrayInput
	CustomDeployRecipes      pulumi.StringArrayInput
	CustomInstanceProfileArn pulumi.StringPtrInput
	CustomJson               pulumi.StringPtrInput
	CustomSecurityGroupIds   pulumi.StringArrayInput
	CustomSetupRecipes       pulumi.StringArrayInput
	CustomShutdownRecipes    pulumi.StringArrayInput
	CustomUndeployRecipes    pulumi.StringArrayInput
	DrainElbOnShutdown       pulumi.BoolPtrInput
	EbsVolumes               NodejsAppLayerEbsVolumeArrayInput
	ElasticLoadBalancer      pulumi.StringPtrInput
	InstallUpdatesOnBoot     pulumi.BoolPtrInput
	InstanceShutdownTimeout  pulumi.IntPtrInput
	Name                     pulumi.StringPtrInput
	NodejsVersion            pulumi.StringPtrInput
	StackId                  pulumi.StringInput
	SystemPackages           pulumi.StringArrayInput
	Tags                     pulumi.StringMapInput
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (NodejsAppLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodejsAppLayerArgs)(nil)).Elem()
}

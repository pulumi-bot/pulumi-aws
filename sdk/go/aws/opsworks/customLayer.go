// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type CustomLayer struct {
	pulumi.CustomResourceState

	Arn                      pulumi.StringOutput             `pulumi:"arn"`
	AutoAssignElasticIps     pulumi.BoolPtrOutput            `pulumi:"autoAssignElasticIps"`
	AutoAssignPublicIps      pulumi.BoolPtrOutput            `pulumi:"autoAssignPublicIps"`
	AutoHealing              pulumi.BoolPtrOutput            `pulumi:"autoHealing"`
	CustomConfigureRecipes   pulumi.StringArrayOutput        `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes      pulumi.StringArrayOutput        `pulumi:"customDeployRecipes"`
	CustomInstanceProfileArn pulumi.StringPtrOutput          `pulumi:"customInstanceProfileArn"`
	CustomJson               pulumi.StringPtrOutput          `pulumi:"customJson"`
	CustomSecurityGroupIds   pulumi.StringArrayOutput        `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes       pulumi.StringArrayOutput        `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes    pulumi.StringArrayOutput        `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes    pulumi.StringArrayOutput        `pulumi:"customUndeployRecipes"`
	DrainElbOnShutdown       pulumi.BoolPtrOutput            `pulumi:"drainElbOnShutdown"`
	EbsVolumes               CustomLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	ElasticLoadBalancer      pulumi.StringPtrOutput          `pulumi:"elasticLoadBalancer"`
	InstallUpdatesOnBoot     pulumi.BoolPtrOutput            `pulumi:"installUpdatesOnBoot"`
	InstanceShutdownTimeout  pulumi.IntPtrOutput             `pulumi:"instanceShutdownTimeout"`
	Name                     pulumi.StringOutput             `pulumi:"name"`
	ShortName                pulumi.StringOutput             `pulumi:"shortName"`
	StackId                  pulumi.StringOutput             `pulumi:"stackId"`
	SystemPackages           pulumi.StringArrayOutput        `pulumi:"systemPackages"`
	Tags                     pulumi.StringMapOutput          `pulumi:"tags"`
	UseEbsOptimizedInstances pulumi.BoolPtrOutput            `pulumi:"useEbsOptimizedInstances"`
}

// NewCustomLayer registers a new resource with the given unique name, arguments, and options.
func NewCustomLayer(ctx *pulumi.Context,
	name string, args *CustomLayerArgs, opts ...pulumi.ResourceOption) (*CustomLayer, error) {
	if args == nil || args.ShortName == nil {
		return nil, errors.New("missing required argument 'ShortName'")
	}
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	if args == nil {
		args = &CustomLayerArgs{}
	}
	var resource CustomLayer
	err := ctx.RegisterResource("aws:opsworks/customLayer:CustomLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomLayer gets an existing CustomLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomLayerState, opts ...pulumi.ResourceOption) (*CustomLayer, error) {
	var resource CustomLayer
	err := ctx.ReadResource("aws:opsworks/customLayer:CustomLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomLayer resources.
type customLayerState struct {
	Arn                      *string                `pulumi:"arn"`
	AutoAssignElasticIps     *bool                  `pulumi:"autoAssignElasticIps"`
	AutoAssignPublicIps      *bool                  `pulumi:"autoAssignPublicIps"`
	AutoHealing              *bool                  `pulumi:"autoHealing"`
	CustomConfigureRecipes   []string               `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes      []string               `pulumi:"customDeployRecipes"`
	CustomInstanceProfileArn *string                `pulumi:"customInstanceProfileArn"`
	CustomJson               *string                `pulumi:"customJson"`
	CustomSecurityGroupIds   []string               `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes       []string               `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes    []string               `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes    []string               `pulumi:"customUndeployRecipes"`
	DrainElbOnShutdown       *bool                  `pulumi:"drainElbOnShutdown"`
	EbsVolumes               []CustomLayerEbsVolume `pulumi:"ebsVolumes"`
	ElasticLoadBalancer      *string                `pulumi:"elasticLoadBalancer"`
	InstallUpdatesOnBoot     *bool                  `pulumi:"installUpdatesOnBoot"`
	InstanceShutdownTimeout  *int                   `pulumi:"instanceShutdownTimeout"`
	Name                     *string                `pulumi:"name"`
	ShortName                *string                `pulumi:"shortName"`
	StackId                  *string                `pulumi:"stackId"`
	SystemPackages           []string               `pulumi:"systemPackages"`
	Tags                     map[string]string      `pulumi:"tags"`
	UseEbsOptimizedInstances *bool                  `pulumi:"useEbsOptimizedInstances"`
}

type CustomLayerState struct {
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
	EbsVolumes               CustomLayerEbsVolumeArrayInput
	ElasticLoadBalancer      pulumi.StringPtrInput
	InstallUpdatesOnBoot     pulumi.BoolPtrInput
	InstanceShutdownTimeout  pulumi.IntPtrInput
	Name                     pulumi.StringPtrInput
	ShortName                pulumi.StringPtrInput
	StackId                  pulumi.StringPtrInput
	SystemPackages           pulumi.StringArrayInput
	Tags                     pulumi.StringMapInput
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (CustomLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*customLayerState)(nil)).Elem()
}

type customLayerArgs struct {
	AutoAssignElasticIps     *bool                  `pulumi:"autoAssignElasticIps"`
	AutoAssignPublicIps      *bool                  `pulumi:"autoAssignPublicIps"`
	AutoHealing              *bool                  `pulumi:"autoHealing"`
	CustomConfigureRecipes   []string               `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes      []string               `pulumi:"customDeployRecipes"`
	CustomInstanceProfileArn *string                `pulumi:"customInstanceProfileArn"`
	CustomJson               *string                `pulumi:"customJson"`
	CustomSecurityGroupIds   []string               `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes       []string               `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes    []string               `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes    []string               `pulumi:"customUndeployRecipes"`
	DrainElbOnShutdown       *bool                  `pulumi:"drainElbOnShutdown"`
	EbsVolumes               []CustomLayerEbsVolume `pulumi:"ebsVolumes"`
	ElasticLoadBalancer      *string                `pulumi:"elasticLoadBalancer"`
	InstallUpdatesOnBoot     *bool                  `pulumi:"installUpdatesOnBoot"`
	InstanceShutdownTimeout  *int                   `pulumi:"instanceShutdownTimeout"`
	Name                     *string                `pulumi:"name"`
	ShortName                string                 `pulumi:"shortName"`
	StackId                  string                 `pulumi:"stackId"`
	SystemPackages           []string               `pulumi:"systemPackages"`
	Tags                     map[string]string      `pulumi:"tags"`
	UseEbsOptimizedInstances *bool                  `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a CustomLayer resource.
type CustomLayerArgs struct {
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
	EbsVolumes               CustomLayerEbsVolumeArrayInput
	ElasticLoadBalancer      pulumi.StringPtrInput
	InstallUpdatesOnBoot     pulumi.BoolPtrInput
	InstanceShutdownTimeout  pulumi.IntPtrInput
	Name                     pulumi.StringPtrInput
	ShortName                pulumi.StringInput
	StackId                  pulumi.StringInput
	SystemPackages           pulumi.StringArrayInput
	Tags                     pulumi.StringMapInput
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (CustomLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customLayerArgs)(nil)).Elem()
}

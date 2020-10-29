// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks haproxy layer resource.
type HaproxyLayer struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrOutput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrOutput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrOutput     `pulumi:"autoHealing"`
	CustomConfigureRecipes pulumi.StringArrayOutput `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    pulumi.StringArrayOutput `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrOutput `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrOutput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayOutput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     pulumi.StringArrayOutput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  pulumi.StringArrayOutput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  pulumi.StringArrayOutput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrOutput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes HaproxyLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod pulumi.StringPtrOutput `pulumi:"healthcheckMethod"`
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl pulumi.StringPtrOutput `pulumi:"healthcheckUrl"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Whether to enable HAProxy stats.
	StatsEnabled pulumi.BoolPtrOutput `pulumi:"statsEnabled"`
	// The password to use for HAProxy stats.
	StatsPassword pulumi.StringOutput `pulumi:"statsPassword"`
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl pulumi.StringPtrOutput `pulumi:"statsUrl"`
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser pulumi.StringPtrOutput `pulumi:"statsUser"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewHaproxyLayer registers a new resource with the given unique name, arguments, and options.
func NewHaproxyLayer(ctx *pulumi.Context,
	name string, args *HaproxyLayerArgs, opts ...pulumi.ResourceOption) (*HaproxyLayer, error) {
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	if args == nil || args.StatsPassword == nil {
		return nil, errors.New("missing required argument 'StatsPassword'")
	}
	if args == nil {
		args = &HaproxyLayerArgs{}
	}
	var resource HaproxyLayer
	err := ctx.RegisterResource("aws:opsworks/haproxyLayer:HaproxyLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHaproxyLayer gets an existing HaproxyLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaproxyLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HaproxyLayerState, opts ...pulumi.ResourceOption) (*HaproxyLayer, error) {
	var resource HaproxyLayer
	err := ctx.ReadResource("aws:opsworks/haproxyLayer:HaproxyLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HaproxyLayer resources.
type haproxyLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            *bool    `pulumi:"autoHealing"`
	CustomConfigureRecipes []string `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    []string `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []HaproxyLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod *string `pulumi:"healthcheckMethod"`
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl *string `pulumi:"healthcheckUrl"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Whether to enable HAProxy stats.
	StatsEnabled *bool `pulumi:"statsEnabled"`
	// The password to use for HAProxy stats.
	StatsPassword *string `pulumi:"statsPassword"`
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl *string `pulumi:"statsUrl"`
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser *string `pulumi:"statsUser"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

type HaproxyLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrInput
	CustomConfigureRecipes pulumi.StringArrayInput
	CustomDeployRecipes    pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes HaproxyLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod pulumi.StringPtrInput
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Whether to enable HAProxy stats.
	StatsEnabled pulumi.BoolPtrInput
	// The password to use for HAProxy stats.
	StatsPassword pulumi.StringPtrInput
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl pulumi.StringPtrInput
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (HaproxyLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*haproxyLayerState)(nil)).Elem()
}

type haproxyLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            *bool    `pulumi:"autoHealing"`
	CustomConfigureRecipes []string `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    []string `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []HaproxyLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod *string `pulumi:"healthcheckMethod"`
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl *string `pulumi:"healthcheckUrl"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Whether to enable HAProxy stats.
	StatsEnabled *bool `pulumi:"statsEnabled"`
	// The password to use for HAProxy stats.
	StatsPassword string `pulumi:"statsPassword"`
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl *string `pulumi:"statsUrl"`
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser *string `pulumi:"statsUser"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a HaproxyLayer resource.
type HaproxyLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrInput
	CustomConfigureRecipes pulumi.StringArrayInput
	CustomDeployRecipes    pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes HaproxyLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod pulumi.StringPtrInput
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Whether to enable HAProxy stats.
	StatsEnabled pulumi.BoolPtrInput
	// The password to use for HAProxy stats.
	StatsPassword pulumi.StringInput
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl pulumi.StringPtrInput
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (HaproxyLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*haproxyLayerArgs)(nil)).Elem()
}

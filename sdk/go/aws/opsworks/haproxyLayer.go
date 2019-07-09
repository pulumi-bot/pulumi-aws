// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks haproxy layer resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_haproxy_layer.html.markdown.
type HaproxyLayer struct {
	s *pulumi.ResourceState
}

// NewHaproxyLayer registers a new resource with the given unique name, arguments, and options.
func NewHaproxyLayer(ctx *pulumi.Context,
	name string, args *HaproxyLayerArgs, opts ...pulumi.ResourceOpt) (*HaproxyLayer, error) {
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	if args == nil || args.StatsPassword == nil {
		return nil, errors.New("missing required argument 'StatsPassword'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoAssignElasticIps"] = nil
		inputs["autoAssignPublicIps"] = nil
		inputs["autoHealing"] = nil
		inputs["customConfigureRecipes"] = nil
		inputs["customDeployRecipes"] = nil
		inputs["customInstanceProfileArn"] = nil
		inputs["customJson"] = nil
		inputs["customSecurityGroupIds"] = nil
		inputs["customSetupRecipes"] = nil
		inputs["customShutdownRecipes"] = nil
		inputs["customUndeployRecipes"] = nil
		inputs["drainElbOnShutdown"] = nil
		inputs["ebsVolumes"] = nil
		inputs["elasticLoadBalancer"] = nil
		inputs["healthcheckMethod"] = nil
		inputs["healthcheckUrl"] = nil
		inputs["installUpdatesOnBoot"] = nil
		inputs["instanceShutdownTimeout"] = nil
		inputs["name"] = nil
		inputs["stackId"] = nil
		inputs["statsEnabled"] = nil
		inputs["statsPassword"] = nil
		inputs["statsUrl"] = nil
		inputs["statsUser"] = nil
		inputs["systemPackages"] = nil
		inputs["useEbsOptimizedInstances"] = nil
	} else {
		inputs["autoAssignElasticIps"] = args.AutoAssignElasticIps
		inputs["autoAssignPublicIps"] = args.AutoAssignPublicIps
		inputs["autoHealing"] = args.AutoHealing
		inputs["customConfigureRecipes"] = args.CustomConfigureRecipes
		inputs["customDeployRecipes"] = args.CustomDeployRecipes
		inputs["customInstanceProfileArn"] = args.CustomInstanceProfileArn
		inputs["customJson"] = args.CustomJson
		inputs["customSecurityGroupIds"] = args.CustomSecurityGroupIds
		inputs["customSetupRecipes"] = args.CustomSetupRecipes
		inputs["customShutdownRecipes"] = args.CustomShutdownRecipes
		inputs["customUndeployRecipes"] = args.CustomUndeployRecipes
		inputs["drainElbOnShutdown"] = args.DrainElbOnShutdown
		inputs["ebsVolumes"] = args.EbsVolumes
		inputs["elasticLoadBalancer"] = args.ElasticLoadBalancer
		inputs["healthcheckMethod"] = args.HealthcheckMethod
		inputs["healthcheckUrl"] = args.HealthcheckUrl
		inputs["installUpdatesOnBoot"] = args.InstallUpdatesOnBoot
		inputs["instanceShutdownTimeout"] = args.InstanceShutdownTimeout
		inputs["name"] = args.Name
		inputs["stackId"] = args.StackId
		inputs["statsEnabled"] = args.StatsEnabled
		inputs["statsPassword"] = args.StatsPassword
		inputs["statsUrl"] = args.StatsUrl
		inputs["statsUser"] = args.StatsUser
		inputs["systemPackages"] = args.SystemPackages
		inputs["useEbsOptimizedInstances"] = args.UseEbsOptimizedInstances
	}
	s, err := ctx.RegisterResource("aws:opsworks/haproxyLayer:HaproxyLayer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HaproxyLayer{s: s}, nil
}

// GetHaproxyLayer gets an existing HaproxyLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaproxyLayer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HaproxyLayerState, opts ...pulumi.ResourceOpt) (*HaproxyLayer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoAssignElasticIps"] = state.AutoAssignElasticIps
		inputs["autoAssignPublicIps"] = state.AutoAssignPublicIps
		inputs["autoHealing"] = state.AutoHealing
		inputs["customConfigureRecipes"] = state.CustomConfigureRecipes
		inputs["customDeployRecipes"] = state.CustomDeployRecipes
		inputs["customInstanceProfileArn"] = state.CustomInstanceProfileArn
		inputs["customJson"] = state.CustomJson
		inputs["customSecurityGroupIds"] = state.CustomSecurityGroupIds
		inputs["customSetupRecipes"] = state.CustomSetupRecipes
		inputs["customShutdownRecipes"] = state.CustomShutdownRecipes
		inputs["customUndeployRecipes"] = state.CustomUndeployRecipes
		inputs["drainElbOnShutdown"] = state.DrainElbOnShutdown
		inputs["ebsVolumes"] = state.EbsVolumes
		inputs["elasticLoadBalancer"] = state.ElasticLoadBalancer
		inputs["healthcheckMethod"] = state.HealthcheckMethod
		inputs["healthcheckUrl"] = state.HealthcheckUrl
		inputs["installUpdatesOnBoot"] = state.InstallUpdatesOnBoot
		inputs["instanceShutdownTimeout"] = state.InstanceShutdownTimeout
		inputs["name"] = state.Name
		inputs["stackId"] = state.StackId
		inputs["statsEnabled"] = state.StatsEnabled
		inputs["statsPassword"] = state.StatsPassword
		inputs["statsUrl"] = state.StatsUrl
		inputs["statsUser"] = state.StatsUser
		inputs["systemPackages"] = state.SystemPackages
		inputs["useEbsOptimizedInstances"] = state.UseEbsOptimizedInstances
	}
	s, err := ctx.ReadResource("aws:opsworks/haproxyLayer:HaproxyLayer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HaproxyLayer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HaproxyLayer) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HaproxyLayer) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Whether to automatically assign an elastic IP address to the layer's instances.
func (r *HaproxyLayer) AutoAssignElasticIps() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoAssignElasticIps"])
}

// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
func (r *HaproxyLayer) AutoAssignPublicIps() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoAssignPublicIps"])
}

// Whether to enable auto-healing for the layer.
func (r *HaproxyLayer) AutoHealing() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoHealing"])
}

func (r *HaproxyLayer) CustomConfigureRecipes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customConfigureRecipes"])
}

func (r *HaproxyLayer) CustomDeployRecipes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customDeployRecipes"])
}

// The ARN of an IAM profile that will be used for the layer's instances.
func (r *HaproxyLayer) CustomInstanceProfileArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["customInstanceProfileArn"])
}

// Custom JSON attributes to apply to the layer.
func (r *HaproxyLayer) CustomJson() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["customJson"])
}

// Ids for a set of security groups to apply to the layer's instances.
func (r *HaproxyLayer) CustomSecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customSecurityGroupIds"])
}

func (r *HaproxyLayer) CustomSetupRecipes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customSetupRecipes"])
}

func (r *HaproxyLayer) CustomShutdownRecipes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customShutdownRecipes"])
}

func (r *HaproxyLayer) CustomUndeployRecipes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customUndeployRecipes"])
}

// Whether to enable Elastic Load Balancing connection draining.
func (r *HaproxyLayer) DrainElbOnShutdown() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["drainElbOnShutdown"])
}

// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
func (r *HaproxyLayer) EbsVolumes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ebsVolumes"])
}

// Name of an Elastic Load Balancer to attach to this layer
func (r *HaproxyLayer) ElasticLoadBalancer() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["elasticLoadBalancer"])
}

// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
func (r *HaproxyLayer) HealthcheckMethod() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["healthcheckMethod"])
}

// URL path to use for instance healthchecks. Defaults to "/".
func (r *HaproxyLayer) HealthcheckUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["healthcheckUrl"])
}

// Whether to install OS and package updates on each instance when it boots.
func (r *HaproxyLayer) InstallUpdatesOnBoot() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["installUpdatesOnBoot"])
}

// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
func (r *HaproxyLayer) InstanceShutdownTimeout() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["instanceShutdownTimeout"])
}

// A human-readable name for the layer.
func (r *HaproxyLayer) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The id of the stack the layer will belong to.
func (r *HaproxyLayer) StackId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["stackId"])
}

// Whether to enable HAProxy stats.
func (r *HaproxyLayer) StatsEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["statsEnabled"])
}

// The password to use for HAProxy stats.
func (r *HaproxyLayer) StatsPassword() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["statsPassword"])
}

// The HAProxy stats URL. Defaults to "/haproxy?stats".
func (r *HaproxyLayer) StatsUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["statsUrl"])
}

// The username for HAProxy stats. Defaults to "opsworks".
func (r *HaproxyLayer) StatsUser() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["statsUser"])
}

// Names of a set of system packages to install on the layer's instances.
func (r *HaproxyLayer) SystemPackages() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["systemPackages"])
}

// Whether to use EBS-optimized instances.
func (r *HaproxyLayer) UseEbsOptimizedInstances() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useEbsOptimizedInstances"])
}

// Input properties used for looking up and filtering HaproxyLayer resources.
type HaproxyLayerState struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps interface{}
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps interface{}
	// Whether to enable auto-healing for the layer.
	AutoHealing interface{}
	CustomConfigureRecipes interface{}
	CustomDeployRecipes interface{}
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn interface{}
	// Custom JSON attributes to apply to the layer.
	CustomJson interface{}
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds interface{}
	CustomSetupRecipes interface{}
	CustomShutdownRecipes interface{}
	CustomUndeployRecipes interface{}
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown interface{}
	// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes interface{}
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer interface{}
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod interface{}
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl interface{}
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot interface{}
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout interface{}
	// A human-readable name for the layer.
	Name interface{}
	// The id of the stack the layer will belong to.
	StackId interface{}
	// Whether to enable HAProxy stats.
	StatsEnabled interface{}
	// The password to use for HAProxy stats.
	StatsPassword interface{}
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl interface{}
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser interface{}
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages interface{}
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances interface{}
}

// The set of arguments for constructing a HaproxyLayer resource.
type HaproxyLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps interface{}
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps interface{}
	// Whether to enable auto-healing for the layer.
	AutoHealing interface{}
	CustomConfigureRecipes interface{}
	CustomDeployRecipes interface{}
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn interface{}
	// Custom JSON attributes to apply to the layer.
	CustomJson interface{}
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds interface{}
	CustomSetupRecipes interface{}
	CustomShutdownRecipes interface{}
	CustomUndeployRecipes interface{}
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown interface{}
	// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes interface{}
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer interface{}
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod interface{}
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl interface{}
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot interface{}
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout interface{}
	// A human-readable name for the layer.
	Name interface{}
	// The id of the stack the layer will belong to.
	StackId interface{}
	// Whether to enable HAProxy stats.
	StatsEnabled interface{}
	// The password to use for HAProxy stats.
	StatsPassword interface{}
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl interface{}
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser interface{}
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages interface{}
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances interface{}
}

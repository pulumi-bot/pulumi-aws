// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks static web server layer resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewStaticWebLayer(ctx, "web", &opsworks.StaticWebLayerArgs{
// 			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
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
// OpsWorks static web server Layers can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:opsworks/staticWebLayer:StaticWebLayer bar 00000000-0000-0000-0000-000000000000
// ```
type StaticWebLayer struct {
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
	CustomJson               pulumi.StringPtrOutput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayOutput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     pulumi.StringArrayOutput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  pulumi.StringArrayOutput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  pulumi.StringArrayOutput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrOutput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes StaticWebLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewStaticWebLayer registers a new resource with the given unique name, arguments, and options.
func NewStaticWebLayer(ctx *pulumi.Context,
	name string, args *StaticWebLayerArgs, opts ...pulumi.ResourceOption) (*StaticWebLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource StaticWebLayer
	err := ctx.RegisterResource("aws:opsworks/staticWebLayer:StaticWebLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticWebLayer gets an existing StaticWebLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticWebLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticWebLayerState, opts ...pulumi.ResourceOption) (*StaticWebLayer, error) {
	var resource StaticWebLayer
	err := ctx.ReadResource("aws:opsworks/staticWebLayer:StaticWebLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticWebLayer resources.
type staticWebLayerState struct {
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
	CustomJson               *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []StaticWebLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

type StaticWebLayerState struct {
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
	CustomJson               pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes StaticWebLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (StaticWebLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticWebLayerState)(nil)).Elem()
}

type staticWebLayerArgs struct {
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
	CustomJson               *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []StaticWebLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a StaticWebLayer resource.
type StaticWebLayerArgs struct {
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
	CustomJson               pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes StaticWebLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (StaticWebLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticWebLayerArgs)(nil)).Elem()
}

type StaticWebLayerInput interface {
	pulumi.Input

	ToStaticWebLayerOutput() StaticWebLayerOutput
	ToStaticWebLayerOutputWithContext(ctx context.Context) StaticWebLayerOutput
}

func (*StaticWebLayer) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticWebLayer)(nil))
}

func (i *StaticWebLayer) ToStaticWebLayerOutput() StaticWebLayerOutput {
	return i.ToStaticWebLayerOutputWithContext(context.Background())
}

func (i *StaticWebLayer) ToStaticWebLayerOutputWithContext(ctx context.Context) StaticWebLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebLayerOutput)
}

func (i *StaticWebLayer) ToStaticWebLayerPtrOutput() StaticWebLayerPtrOutput {
	return i.ToStaticWebLayerPtrOutputWithContext(context.Background())
}

func (i *StaticWebLayer) ToStaticWebLayerPtrOutputWithContext(ctx context.Context) StaticWebLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebLayerPtrOutput)
}

type StaticWebLayerPtrInput interface {
	pulumi.Input

	ToStaticWebLayerPtrOutput() StaticWebLayerPtrOutput
	ToStaticWebLayerPtrOutputWithContext(ctx context.Context) StaticWebLayerPtrOutput
}

type staticWebLayerPtrType StaticWebLayerArgs

func (*staticWebLayerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticWebLayer)(nil))
}

func (i *staticWebLayerPtrType) ToStaticWebLayerPtrOutput() StaticWebLayerPtrOutput {
	return i.ToStaticWebLayerPtrOutputWithContext(context.Background())
}

func (i *staticWebLayerPtrType) ToStaticWebLayerPtrOutputWithContext(ctx context.Context) StaticWebLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebLayerOutput).ToStaticWebLayerPtrOutput()
}

// StaticWebLayerArrayInput is an input type that accepts StaticWebLayerArray and StaticWebLayerArrayOutput values.
// You can construct a concrete instance of `StaticWebLayerArrayInput` via:
//
//          StaticWebLayerArray{ StaticWebLayerArgs{...} }
type StaticWebLayerArrayInput interface {
	pulumi.Input

	ToStaticWebLayerArrayOutput() StaticWebLayerArrayOutput
	ToStaticWebLayerArrayOutputWithContext(context.Context) StaticWebLayerArrayOutput
}

type StaticWebLayerArray []StaticWebLayerInput

func (StaticWebLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticWebLayer)(nil))
}

func (i StaticWebLayerArray) ToStaticWebLayerArrayOutput() StaticWebLayerArrayOutput {
	return i.ToStaticWebLayerArrayOutputWithContext(context.Background())
}

func (i StaticWebLayerArray) ToStaticWebLayerArrayOutputWithContext(ctx context.Context) StaticWebLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebLayerArrayOutput)
}

// StaticWebLayerMapInput is an input type that accepts StaticWebLayerMap and StaticWebLayerMapOutput values.
// You can construct a concrete instance of `StaticWebLayerMapInput` via:
//
//          StaticWebLayerMap{ "key": StaticWebLayerArgs{...} }
type StaticWebLayerMapInput interface {
	pulumi.Input

	ToStaticWebLayerMapOutput() StaticWebLayerMapOutput
	ToStaticWebLayerMapOutputWithContext(context.Context) StaticWebLayerMapOutput
}

type StaticWebLayerMap map[string]StaticWebLayerInput

func (StaticWebLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StaticWebLayer)(nil))
}

func (i StaticWebLayerMap) ToStaticWebLayerMapOutput() StaticWebLayerMapOutput {
	return i.ToStaticWebLayerMapOutputWithContext(context.Background())
}

func (i StaticWebLayerMap) ToStaticWebLayerMapOutputWithContext(ctx context.Context) StaticWebLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebLayerMapOutput)
}

type StaticWebLayerOutput struct {
	*pulumi.OutputState
}

func (StaticWebLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticWebLayer)(nil))
}

func (o StaticWebLayerOutput) ToStaticWebLayerOutput() StaticWebLayerOutput {
	return o
}

func (o StaticWebLayerOutput) ToStaticWebLayerOutputWithContext(ctx context.Context) StaticWebLayerOutput {
	return o
}

func (o StaticWebLayerOutput) ToStaticWebLayerPtrOutput() StaticWebLayerPtrOutput {
	return o.ToStaticWebLayerPtrOutputWithContext(context.Background())
}

func (o StaticWebLayerOutput) ToStaticWebLayerPtrOutputWithContext(ctx context.Context) StaticWebLayerPtrOutput {
	return o.ApplyT(func(v StaticWebLayer) *StaticWebLayer {
		return &v
	}).(StaticWebLayerPtrOutput)
}

type StaticWebLayerPtrOutput struct {
	*pulumi.OutputState
}

func (StaticWebLayerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticWebLayer)(nil))
}

func (o StaticWebLayerPtrOutput) ToStaticWebLayerPtrOutput() StaticWebLayerPtrOutput {
	return o
}

func (o StaticWebLayerPtrOutput) ToStaticWebLayerPtrOutputWithContext(ctx context.Context) StaticWebLayerPtrOutput {
	return o
}

type StaticWebLayerArrayOutput struct{ *pulumi.OutputState }

func (StaticWebLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticWebLayer)(nil))
}

func (o StaticWebLayerArrayOutput) ToStaticWebLayerArrayOutput() StaticWebLayerArrayOutput {
	return o
}

func (o StaticWebLayerArrayOutput) ToStaticWebLayerArrayOutputWithContext(ctx context.Context) StaticWebLayerArrayOutput {
	return o
}

func (o StaticWebLayerArrayOutput) Index(i pulumi.IntInput) StaticWebLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticWebLayer {
		return vs[0].([]StaticWebLayer)[vs[1].(int)]
	}).(StaticWebLayerOutput)
}

type StaticWebLayerMapOutput struct{ *pulumi.OutputState }

func (StaticWebLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StaticWebLayer)(nil))
}

func (o StaticWebLayerMapOutput) ToStaticWebLayerMapOutput() StaticWebLayerMapOutput {
	return o
}

func (o StaticWebLayerMapOutput) ToStaticWebLayerMapOutputWithContext(ctx context.Context) StaticWebLayerMapOutput {
	return o
}

func (o StaticWebLayerMapOutput) MapIndex(k pulumi.StringInput) StaticWebLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StaticWebLayer {
		return vs[0].(map[string]StaticWebLayer)[vs[1].(string)]
	}).(StaticWebLayerOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticWebLayerOutput{})
	pulumi.RegisterOutputType(StaticWebLayerPtrOutput{})
	pulumi.RegisterOutputType(StaticWebLayerArrayOutput{})
	pulumi.RegisterOutputType(StaticWebLayerMapOutput{})
}

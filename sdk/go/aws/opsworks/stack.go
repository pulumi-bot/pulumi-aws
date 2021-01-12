// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks stack resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewStack(ctx, "main", &opsworks.StackArgs{
// 			Region:                    pulumi.String("us-west-1"),
// 			ServiceRoleArn:            pulumi.Any(aws_iam_role.Opsworks.Arn),
// 			DefaultInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Opsworks.Arn),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("foobar-stack"),
// 			},
// 			CustomJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", " \"foobar\": {\n", "    \"version\": \"1.0.0\"\n", "  }\n", "}\n")),
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
// OpsWorks stacks can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:opsworks/stack:Stack bar 00000000-0000-0000-0000-000000000000
// ```
type Stack struct {
	pulumi.CustomResourceState

	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`
	Arn          pulumi.StringOutput `pulumi:"arn"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringPtrOutput `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringPtrOutput `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringPtrOutput `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringPtrOutput `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources StackCustomCookbooksSourceArrayOutput `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringPtrOutput `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringOutput `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringOutput `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringPtrOutput `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringPtrOutput `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringPtrOutput `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringOutput `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringPtrOutput `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolPtrOutput `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`
	StackEndpoint  pulumi.StringOutput `pulumi:"stackEndpoint"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolPtrOutput `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolPtrOutput `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultInstanceProfileArn == nil {
		return nil, errors.New("invalid value for required argument 'DefaultInstanceProfileArn'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ServiceRoleArn'")
	}
	var resource Stack
	err := ctx.RegisterResource("aws:opsworks/stack:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("aws:opsworks/stack:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion *string `pulumi:"agentVersion"`
	Arn          *string `pulumi:"arn"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion *string `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color *string `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName *string `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion *string `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources []StackCustomCookbooksSource `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson *string `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone *string `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn *string `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs *string `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType *string `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName *string `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId *string `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme *string `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf *bool `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name *string `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region *string `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	StackEndpoint  *string `pulumi:"stackEndpoint"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks *bool `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups *bool `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId *string `pulumi:"vpcId"`
}

type StackState struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringPtrInput
	Arn          pulumi.StringPtrInput
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringPtrInput
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringPtrInput
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringPtrInput
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringPtrInput
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources StackCustomCookbooksSourceArrayInput
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringPtrInput
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringPtrInput
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringPtrInput
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringPtrInput
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringPtrInput
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringPtrInput
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringPtrInput
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringPtrInput
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolPtrInput
	// The name of the stack.
	Name pulumi.StringPtrInput
	// The name of the region where the stack will exist.
	Region pulumi.StringPtrInput
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringPtrInput
	StackEndpoint  pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolPtrInput
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolPtrInput
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion *string `pulumi:"agentVersion"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion *string `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color *string `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName *string `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion *string `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources []StackCustomCookbooksSource `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson *string `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone *string `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn string `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs *string `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType *string `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName *string `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId *string `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme *string `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf *bool `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name *string `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region string `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn string `pulumi:"serviceRoleArn"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks *bool `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups *bool `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringPtrInput
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringPtrInput
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringPtrInput
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringPtrInput
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringPtrInput
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources StackCustomCookbooksSourceArrayInput
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringPtrInput
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringPtrInput
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringInput
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringPtrInput
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringPtrInput
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringPtrInput
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringPtrInput
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringPtrInput
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolPtrInput
	// The name of the stack.
	Name pulumi.StringPtrInput
	// The name of the region where the stack will exist.
	Region pulumi.StringInput
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolPtrInput
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolPtrInput
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

type StackInput interface {
	pulumi.Input

	ToStackOutput() StackOutput
	ToStackOutputWithContext(ctx context.Context) StackOutput
}

func (*Stack) ElementType() reflect.Type {
	return reflect.TypeOf((*Stack)(nil))
}

func (i *Stack) ToStackOutput() StackOutput {
	return i.ToStackOutputWithContext(context.Background())
}

func (i *Stack) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackOutput)
}

func (i *Stack) ToStackPtrOutput() StackPtrOutput {
	return i.ToStackPtrOutputWithContext(context.Background())
}

func (i *Stack) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPtrOutput)
}

type StackPtrInput interface {
	pulumi.Input

	ToStackPtrOutput() StackPtrOutput
	ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput
}

type stackPtrType StackArgs

func (*stackPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil))
}

func (i *stackPtrType) ToStackPtrOutput() StackPtrOutput {
	return i.ToStackPtrOutputWithContext(context.Background())
}

func (i *stackPtrType) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackOutput).ToStackPtrOutput()
}

// StackArrayInput is an input type that accepts StackArray and StackArrayOutput values.
// You can construct a concrete instance of `StackArrayInput` via:
//
//          StackArray{ StackArgs{...} }
type StackArrayInput interface {
	pulumi.Input

	ToStackArrayOutput() StackArrayOutput
	ToStackArrayOutputWithContext(context.Context) StackArrayOutput
}

type StackArray []StackInput

func (StackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Stack)(nil))
}

func (i StackArray) ToStackArrayOutput() StackArrayOutput {
	return i.ToStackArrayOutputWithContext(context.Background())
}

func (i StackArray) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackArrayOutput)
}

// StackMapInput is an input type that accepts StackMap and StackMapOutput values.
// You can construct a concrete instance of `StackMapInput` via:
//
//          StackMap{ "key": StackArgs{...} }
type StackMapInput interface {
	pulumi.Input

	ToStackMapOutput() StackMapOutput
	ToStackMapOutputWithContext(context.Context) StackMapOutput
}

type StackMap map[string]StackInput

func (StackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Stack)(nil))
}

func (i StackMap) ToStackMapOutput() StackMapOutput {
	return i.ToStackMapOutputWithContext(context.Background())
}

func (i StackMap) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackMapOutput)
}

type StackOutput struct {
	*pulumi.OutputState
}

func (StackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Stack)(nil))
}

func (o StackOutput) ToStackOutput() StackOutput {
	return o
}

func (o StackOutput) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return o
}

func (o StackOutput) ToStackPtrOutput() StackPtrOutput {
	return o.ToStackPtrOutputWithContext(context.Background())
}

func (o StackOutput) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return o.ApplyT(func(v Stack) *Stack {
		return &v
	}).(StackPtrOutput)
}

type StackPtrOutput struct {
	*pulumi.OutputState
}

func (StackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil))
}

func (o StackPtrOutput) ToStackPtrOutput() StackPtrOutput {
	return o
}

func (o StackPtrOutput) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return o
}

type StackArrayOutput struct{ *pulumi.OutputState }

func (StackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Stack)(nil))
}

func (o StackArrayOutput) ToStackArrayOutput() StackArrayOutput {
	return o
}

func (o StackArrayOutput) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return o
}

func (o StackArrayOutput) Index(i pulumi.IntInput) StackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Stack {
		return vs[0].([]Stack)[vs[1].(int)]
	}).(StackOutput)
}

type StackMapOutput struct{ *pulumi.OutputState }

func (StackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Stack)(nil))
}

func (o StackMapOutput) ToStackMapOutput() StackMapOutput {
	return o
}

func (o StackMapOutput) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return o
}

func (o StackMapOutput) MapIndex(k pulumi.StringInput) StackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Stack {
		return vs[0].(map[string]Stack)[vs[1].(string)]
	}).(StackOutput)
}

func init() {
	pulumi.RegisterOutputType(StackOutput{})
	pulumi.RegisterOutputType(StackPtrOutput{})
	pulumi.RegisterOutputType(StackArrayOutput{})
	pulumi.RegisterOutputType(StackMapOutput{})
}

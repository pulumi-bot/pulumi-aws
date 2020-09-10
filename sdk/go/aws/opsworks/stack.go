// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Stack struct {
	pulumi.CustomResourceState

	AgentVersion                pulumi.StringOutput                   `pulumi:"agentVersion"`
	Arn                         pulumi.StringOutput                   `pulumi:"arn"`
	BerkshelfVersion            pulumi.StringPtrOutput                `pulumi:"berkshelfVersion"`
	Color                       pulumi.StringPtrOutput                `pulumi:"color"`
	ConfigurationManagerName    pulumi.StringPtrOutput                `pulumi:"configurationManagerName"`
	ConfigurationManagerVersion pulumi.StringPtrOutput                `pulumi:"configurationManagerVersion"`
	CustomCookbooksSources      StackCustomCookbooksSourceArrayOutput `pulumi:"customCookbooksSources"`
	CustomJson                  pulumi.StringPtrOutput                `pulumi:"customJson"`
	DefaultAvailabilityZone     pulumi.StringOutput                   `pulumi:"defaultAvailabilityZone"`
	DefaultInstanceProfileArn   pulumi.StringOutput                   `pulumi:"defaultInstanceProfileArn"`
	DefaultOs                   pulumi.StringPtrOutput                `pulumi:"defaultOs"`
	DefaultRootDeviceType       pulumi.StringPtrOutput                `pulumi:"defaultRootDeviceType"`
	DefaultSshKeyName           pulumi.StringPtrOutput                `pulumi:"defaultSshKeyName"`
	DefaultSubnetId             pulumi.StringOutput                   `pulumi:"defaultSubnetId"`
	HostnameTheme               pulumi.StringPtrOutput                `pulumi:"hostnameTheme"`
	ManageBerkshelf             pulumi.BoolPtrOutput                  `pulumi:"manageBerkshelf"`
	Name                        pulumi.StringOutput                   `pulumi:"name"`
	Region                      pulumi.StringOutput                   `pulumi:"region"`
	ServiceRoleArn              pulumi.StringOutput                   `pulumi:"serviceRoleArn"`
	StackEndpoint               pulumi.StringOutput                   `pulumi:"stackEndpoint"`
	Tags                        pulumi.StringMapOutput                `pulumi:"tags"`
	UseCustomCookbooks          pulumi.BoolPtrOutput                  `pulumi:"useCustomCookbooks"`
	UseOpsworksSecurityGroups   pulumi.BoolPtrOutput                  `pulumi:"useOpsworksSecurityGroups"`
	VpcId                       pulumi.StringOutput                   `pulumi:"vpcId"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil || args.DefaultInstanceProfileArn == nil {
		return nil, errors.New("missing required argument 'DefaultInstanceProfileArn'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.ServiceRoleArn == nil {
		return nil, errors.New("missing required argument 'ServiceRoleArn'")
	}
	if args == nil {
		args = &StackArgs{}
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
	AgentVersion                *string                      `pulumi:"agentVersion"`
	Arn                         *string                      `pulumi:"arn"`
	BerkshelfVersion            *string                      `pulumi:"berkshelfVersion"`
	Color                       *string                      `pulumi:"color"`
	ConfigurationManagerName    *string                      `pulumi:"configurationManagerName"`
	ConfigurationManagerVersion *string                      `pulumi:"configurationManagerVersion"`
	CustomCookbooksSources      []StackCustomCookbooksSource `pulumi:"customCookbooksSources"`
	CustomJson                  *string                      `pulumi:"customJson"`
	DefaultAvailabilityZone     *string                      `pulumi:"defaultAvailabilityZone"`
	DefaultInstanceProfileArn   *string                      `pulumi:"defaultInstanceProfileArn"`
	DefaultOs                   *string                      `pulumi:"defaultOs"`
	DefaultRootDeviceType       *string                      `pulumi:"defaultRootDeviceType"`
	DefaultSshKeyName           *string                      `pulumi:"defaultSshKeyName"`
	DefaultSubnetId             *string                      `pulumi:"defaultSubnetId"`
	HostnameTheme               *string                      `pulumi:"hostnameTheme"`
	ManageBerkshelf             *bool                        `pulumi:"manageBerkshelf"`
	Name                        *string                      `pulumi:"name"`
	Region                      *string                      `pulumi:"region"`
	ServiceRoleArn              *string                      `pulumi:"serviceRoleArn"`
	StackEndpoint               *string                      `pulumi:"stackEndpoint"`
	Tags                        map[string]string            `pulumi:"tags"`
	UseCustomCookbooks          *bool                        `pulumi:"useCustomCookbooks"`
	UseOpsworksSecurityGroups   *bool                        `pulumi:"useOpsworksSecurityGroups"`
	VpcId                       *string                      `pulumi:"vpcId"`
}

type StackState struct {
	AgentVersion                pulumi.StringPtrInput
	Arn                         pulumi.StringPtrInput
	BerkshelfVersion            pulumi.StringPtrInput
	Color                       pulumi.StringPtrInput
	ConfigurationManagerName    pulumi.StringPtrInput
	ConfigurationManagerVersion pulumi.StringPtrInput
	CustomCookbooksSources      StackCustomCookbooksSourceArrayInput
	CustomJson                  pulumi.StringPtrInput
	DefaultAvailabilityZone     pulumi.StringPtrInput
	DefaultInstanceProfileArn   pulumi.StringPtrInput
	DefaultOs                   pulumi.StringPtrInput
	DefaultRootDeviceType       pulumi.StringPtrInput
	DefaultSshKeyName           pulumi.StringPtrInput
	DefaultSubnetId             pulumi.StringPtrInput
	HostnameTheme               pulumi.StringPtrInput
	ManageBerkshelf             pulumi.BoolPtrInput
	Name                        pulumi.StringPtrInput
	Region                      pulumi.StringPtrInput
	ServiceRoleArn              pulumi.StringPtrInput
	StackEndpoint               pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	UseCustomCookbooks          pulumi.BoolPtrInput
	UseOpsworksSecurityGroups   pulumi.BoolPtrInput
	VpcId                       pulumi.StringPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	AgentVersion                *string                      `pulumi:"agentVersion"`
	BerkshelfVersion            *string                      `pulumi:"berkshelfVersion"`
	Color                       *string                      `pulumi:"color"`
	ConfigurationManagerName    *string                      `pulumi:"configurationManagerName"`
	ConfigurationManagerVersion *string                      `pulumi:"configurationManagerVersion"`
	CustomCookbooksSources      []StackCustomCookbooksSource `pulumi:"customCookbooksSources"`
	CustomJson                  *string                      `pulumi:"customJson"`
	DefaultAvailabilityZone     *string                      `pulumi:"defaultAvailabilityZone"`
	DefaultInstanceProfileArn   string                       `pulumi:"defaultInstanceProfileArn"`
	DefaultOs                   *string                      `pulumi:"defaultOs"`
	DefaultRootDeviceType       *string                      `pulumi:"defaultRootDeviceType"`
	DefaultSshKeyName           *string                      `pulumi:"defaultSshKeyName"`
	DefaultSubnetId             *string                      `pulumi:"defaultSubnetId"`
	HostnameTheme               *string                      `pulumi:"hostnameTheme"`
	ManageBerkshelf             *bool                        `pulumi:"manageBerkshelf"`
	Name                        *string                      `pulumi:"name"`
	Region                      string                       `pulumi:"region"`
	ServiceRoleArn              string                       `pulumi:"serviceRoleArn"`
	Tags                        map[string]string            `pulumi:"tags"`
	UseCustomCookbooks          *bool                        `pulumi:"useCustomCookbooks"`
	UseOpsworksSecurityGroups   *bool                        `pulumi:"useOpsworksSecurityGroups"`
	VpcId                       *string                      `pulumi:"vpcId"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	AgentVersion                pulumi.StringPtrInput
	BerkshelfVersion            pulumi.StringPtrInput
	Color                       pulumi.StringPtrInput
	ConfigurationManagerName    pulumi.StringPtrInput
	ConfigurationManagerVersion pulumi.StringPtrInput
	CustomCookbooksSources      StackCustomCookbooksSourceArrayInput
	CustomJson                  pulumi.StringPtrInput
	DefaultAvailabilityZone     pulumi.StringPtrInput
	DefaultInstanceProfileArn   pulumi.StringInput
	DefaultOs                   pulumi.StringPtrInput
	DefaultRootDeviceType       pulumi.StringPtrInput
	DefaultSshKeyName           pulumi.StringPtrInput
	DefaultSubnetId             pulumi.StringPtrInput
	HostnameTheme               pulumi.StringPtrInput
	ManageBerkshelf             pulumi.BoolPtrInput
	Name                        pulumi.StringPtrInput
	Region                      pulumi.StringInput
	ServiceRoleArn              pulumi.StringInput
	Tags                        pulumi.StringMapInput
	UseCustomCookbooks          pulumi.BoolPtrInput
	UseOpsworksSecurityGroups   pulumi.BoolPtrInput
	VpcId                       pulumi.StringPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

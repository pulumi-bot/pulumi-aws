// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic Beanstalk Environment Resource. Elastic Beanstalk allows
// you to deploy and manage applications in the AWS cloud without worrying about
// the infrastructure that runs those applications.
//
// Environments are often things such as `development`, `integration`, or
// `production`.
//
//
// ## Option Settings
//
// Some options can be stack-specific, check [AWS Docs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html)
// for supported options and examples.
//
// The `setting` and `allSettings` mappings support the following format:
//
// * `namespace` - unique namespace identifying the option's associated AWS resource
// * `name` - name of the configuration option
// * `value` - value for the configuration option
// * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)
type Environment struct {
	pulumi.CustomResourceState

	// List of all option settings configured in this Environment. These
	// are a combination of default settings and their overrides from `setting` in
	// the configuration.
	AllSettings EnvironmentAllSettingArrayOutput `pulumi:"allSettings"`
	// Name of the application that contains the version
	// to be deployed
	Application pulumi.StringOutput `pulumi:"application"`
	Arn         pulumi.StringOutput `pulumi:"arn"`
	// The autoscaling groups used by this Environment.
	AutoscalingGroups pulumi.StringArrayOutput `pulumi:"autoscalingGroups"`
	// Fully qualified DNS name for this Environment.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix pulumi.StringOutput `pulumi:"cnamePrefix"`
	// Short description of the Environment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The URL to the Load Balancer for this Environment
	EndpointUrl pulumi.StringOutput `pulumi:"endpointUrl"`
	// Instances used by this Environment.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// Launch configurations in use by this Environment.
	LaunchConfigurations pulumi.StringArrayOutput `pulumi:"launchConfigurations"`
	// Elastic load balancers in use by this Environment.
	LoadBalancers pulumi.StringArrayOutput `pulumi:"loadBalancers"`
	// A unique name for this Environment. This name is used
	// in the application URL
	Name pulumi.StringOutput `pulumi:"name"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn pulumi.StringOutput `pulumi:"platformArn"`
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval pulumi.StringPtrOutput `pulumi:"pollInterval"`
	// SQS queues in use by this Environment.
	Queues pulumi.StringArrayOutput `pulumi:"queues"`
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings EnvironmentSettingArrayOutput `pulumi:"settings"`
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName pulumi.StringOutput `pulumi:"solutionStackName"`
	// A set of tags to apply to the Environment.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName pulumi.StringPtrOutput `pulumi:"templateName"`
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// Autoscaling triggers in use by this Environment.
	Triggers pulumi.StringArrayOutput `pulumi:"triggers"`
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version pulumi.StringOutput `pulumi:"version"`
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout pulumi.StringPtrOutput `pulumi:"waitForReadyTimeout"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil || args.Application == nil {
		return nil, errors.New("missing required argument 'Application'")
	}
	if args == nil {
		args = &EnvironmentArgs{}
	}
	var resource Environment
	err := ctx.RegisterResource("aws:elasticbeanstalk/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("aws:elasticbeanstalk/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// List of all option settings configured in this Environment. These
	// are a combination of default settings and their overrides from `setting` in
	// the configuration.
	AllSettings []EnvironmentAllSetting `pulumi:"allSettings"`
	// Name of the application that contains the version
	// to be deployed
	Application *string `pulumi:"application"`
	Arn         *string `pulumi:"arn"`
	// The autoscaling groups used by this Environment.
	AutoscalingGroups []string `pulumi:"autoscalingGroups"`
	// Fully qualified DNS name for this Environment.
	Cname *string `pulumi:"cname"`
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix *string `pulumi:"cnamePrefix"`
	// Short description of the Environment
	Description *string `pulumi:"description"`
	// The URL to the Load Balancer for this Environment
	EndpointUrl *string `pulumi:"endpointUrl"`
	// Instances used by this Environment.
	Instances []string `pulumi:"instances"`
	// Launch configurations in use by this Environment.
	LaunchConfigurations []string `pulumi:"launchConfigurations"`
	// Elastic load balancers in use by this Environment.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// A unique name for this Environment. This name is used
	// in the application URL
	Name *string `pulumi:"name"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn *string `pulumi:"platformArn"`
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval *string `pulumi:"pollInterval"`
	// SQS queues in use by this Environment.
	Queues []string `pulumi:"queues"`
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings []EnvironmentSetting `pulumi:"settings"`
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName *string `pulumi:"solutionStackName"`
	// A set of tags to apply to the Environment.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName *string `pulumi:"templateName"`
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier *string `pulumi:"tier"`
	// Autoscaling triggers in use by this Environment.
	Triggers []string `pulumi:"triggers"`
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version *string `pulumi:"version"`
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout *string `pulumi:"waitForReadyTimeout"`
}

type EnvironmentState struct {
	// List of all option settings configured in this Environment. These
	// are a combination of default settings and their overrides from `setting` in
	// the configuration.
	AllSettings EnvironmentAllSettingArrayInput
	// Name of the application that contains the version
	// to be deployed
	Application pulumi.StringPtrInput
	Arn         pulumi.StringPtrInput
	// The autoscaling groups used by this Environment.
	AutoscalingGroups pulumi.StringArrayInput
	// Fully qualified DNS name for this Environment.
	Cname pulumi.StringPtrInput
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix pulumi.StringPtrInput
	// Short description of the Environment
	Description pulumi.StringPtrInput
	// The URL to the Load Balancer for this Environment
	EndpointUrl pulumi.StringPtrInput
	// Instances used by this Environment.
	Instances pulumi.StringArrayInput
	// Launch configurations in use by this Environment.
	LaunchConfigurations pulumi.StringArrayInput
	// Elastic load balancers in use by this Environment.
	LoadBalancers pulumi.StringArrayInput
	// A unique name for this Environment. This name is used
	// in the application URL
	Name pulumi.StringPtrInput
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn pulumi.StringPtrInput
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval pulumi.StringPtrInput
	// SQS queues in use by this Environment.
	Queues pulumi.StringArrayInput
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings EnvironmentSettingArrayInput
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName pulumi.StringPtrInput
	// A set of tags to apply to the Environment.
	Tags pulumi.MapInput
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName pulumi.StringPtrInput
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier pulumi.StringPtrInput
	// Autoscaling triggers in use by this Environment.
	Triggers pulumi.StringArrayInput
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version pulumi.StringPtrInput
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout pulumi.StringPtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Name of the application that contains the version
	// to be deployed
	Application interface{} `pulumi:"application"`
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix *string `pulumi:"cnamePrefix"`
	// Short description of the Environment
	Description *string `pulumi:"description"`
	// A unique name for this Environment. This name is used
	// in the application URL
	Name *string `pulumi:"name"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn *string `pulumi:"platformArn"`
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval *string `pulumi:"pollInterval"`
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings []EnvironmentSetting `pulumi:"settings"`
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName *string `pulumi:"solutionStackName"`
	// A set of tags to apply to the Environment.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName *string `pulumi:"templateName"`
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier *string `pulumi:"tier"`
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version *string `pulumi:"version"`
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout *string `pulumi:"waitForReadyTimeout"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Name of the application that contains the version
	// to be deployed
	Application pulumi.Input
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix pulumi.StringPtrInput
	// Short description of the Environment
	Description pulumi.StringPtrInput
	// A unique name for this Environment. This name is used
	// in the application URL
	Name pulumi.StringPtrInput
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn pulumi.StringPtrInput
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval pulumi.StringPtrInput
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings EnvironmentSettingArrayInput
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName pulumi.StringPtrInput
	// A set of tags to apply to the Environment.
	Tags pulumi.MapInput
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName pulumi.StringPtrInput
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier pulumi.StringPtrInput
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version pulumi.StringPtrInput
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout pulumi.StringPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

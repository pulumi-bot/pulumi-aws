// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Image Builder Infrastructure Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/imagebuilder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := imagebuilder.NewInfrastructureConfiguration(ctx, "example", &imagebuilder.InfrastructureConfigurationArgs{
// 			Description:         pulumi.String("example description"),
// 			InstanceProfileName: pulumi.Any(aws_iam_instance_profile.Example.Name),
// 			InstanceTypes: pulumi.StringArray{
// 				pulumi.String("t2.nano"),
// 				pulumi.String("t3.micro"),
// 			},
// 			KeyPair: pulumi.Any(aws_key_pair.Example.Key_name),
// 			SecurityGroupIds: pulumi.StringArray{
// 				pulumi.Any(aws_security_group.Example.Id),
// 			},
// 			SnsTopicArn:                pulumi.Any(aws_sns_topic.Example.Arn),
// 			SubnetId:                   pulumi.Any(aws_subnet.Main.Id),
// 			TerminateInstanceOnFailure: pulumi.Bool(true),
// 			Logging: &imagebuilder.InfrastructureConfigurationLoggingArgs{
// 				S3Logs: &imagebuilder.InfrastructureConfigurationLoggingS3LogsArgs{
// 					S3BucketName: pulumi.Any(aws_s3_bucket.Example.Bucket),
// 					S3KeyPrefix:  pulumi.String("logs"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
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
// `aws_imagebuilder_infrastructure_configuration` can be imported using the Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-component/example
// ```
type InfrastructureConfiguration struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Date when the configuration was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Date when the configuration was updated.
	DateUpdated pulumi.StringOutput `pulumi:"dateUpdated"`
	// Description for the configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of IAM Instance Profile.
	InstanceProfileName pulumi.StringOutput `pulumi:"instanceProfileName"`
	// Set of EC2 Instance Types.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// Name of EC2 Key Pair.
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// Configuration block with logging settings. Detailed below.
	Logging InfrastructureConfigurationLoggingPtrOutput `pulumi:"logging"`
	// Name for the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags pulumi.StringMapOutput `pulumi:"resourceTags"`
	// Set of EC2 Security Group identifiers.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn pulumi.StringPtrOutput `pulumi:"snsTopicArn"`
	// EC2 Subnet identifier. Also requires `securityGroupIds` argument.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
	TerminateInstanceOnFailure pulumi.BoolPtrOutput `pulumi:"terminateInstanceOnFailure"`
}

// NewInfrastructureConfiguration registers a new resource with the given unique name, arguments, and options.
func NewInfrastructureConfiguration(ctx *pulumi.Context,
	name string, args *InfrastructureConfigurationArgs, opts ...pulumi.ResourceOption) (*InfrastructureConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceProfileName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceProfileName'")
	}
	var resource InfrastructureConfiguration
	err := ctx.RegisterResource("aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInfrastructureConfiguration gets an existing InfrastructureConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInfrastructureConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InfrastructureConfigurationState, opts ...pulumi.ResourceOption) (*InfrastructureConfiguration, error) {
	var resource InfrastructureConfiguration
	err := ctx.ReadResource("aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InfrastructureConfiguration resources.
type infrastructureConfigurationState struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn *string `pulumi:"arn"`
	// Date when the configuration was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Date when the configuration was updated.
	DateUpdated *string `pulumi:"dateUpdated"`
	// Description for the configuration.
	Description *string `pulumi:"description"`
	// Name of IAM Instance Profile.
	InstanceProfileName *string `pulumi:"instanceProfileName"`
	// Set of EC2 Instance Types.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Name of EC2 Key Pair.
	KeyPair *string `pulumi:"keyPair"`
	// Configuration block with logging settings. Detailed below.
	Logging *InfrastructureConfigurationLogging `pulumi:"logging"`
	// Name for the configuration.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags map[string]string `pulumi:"resourceTags"`
	// Set of EC2 Security Group identifiers.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// EC2 Subnet identifier. Also requires `securityGroupIds` argument.
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
	TerminateInstanceOnFailure *bool `pulumi:"terminateInstanceOnFailure"`
}

type InfrastructureConfigurationState struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn pulumi.StringPtrInput
	// Date when the configuration was created.
	DateCreated pulumi.StringPtrInput
	// Date when the configuration was updated.
	DateUpdated pulumi.StringPtrInput
	// Description for the configuration.
	Description pulumi.StringPtrInput
	// Name of IAM Instance Profile.
	InstanceProfileName pulumi.StringPtrInput
	// Set of EC2 Instance Types.
	InstanceTypes pulumi.StringArrayInput
	// Name of EC2 Key Pair.
	KeyPair pulumi.StringPtrInput
	// Configuration block with logging settings. Detailed below.
	Logging InfrastructureConfigurationLoggingPtrInput
	// Name for the configuration.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags pulumi.StringMapInput
	// Set of EC2 Security Group identifiers.
	SecurityGroupIds pulumi.StringArrayInput
	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn pulumi.StringPtrInput
	// EC2 Subnet identifier. Also requires `securityGroupIds` argument.
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
	TerminateInstanceOnFailure pulumi.BoolPtrInput
}

func (InfrastructureConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*infrastructureConfigurationState)(nil)).Elem()
}

type infrastructureConfigurationArgs struct {
	// Description for the configuration.
	Description *string `pulumi:"description"`
	// Name of IAM Instance Profile.
	InstanceProfileName string `pulumi:"instanceProfileName"`
	// Set of EC2 Instance Types.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Name of EC2 Key Pair.
	KeyPair *string `pulumi:"keyPair"`
	// Configuration block with logging settings. Detailed below.
	Logging *InfrastructureConfigurationLogging `pulumi:"logging"`
	// Name for the configuration.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags map[string]string `pulumi:"resourceTags"`
	// Set of EC2 Security Group identifiers.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// EC2 Subnet identifier. Also requires `securityGroupIds` argument.
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
	TerminateInstanceOnFailure *bool `pulumi:"terminateInstanceOnFailure"`
}

// The set of arguments for constructing a InfrastructureConfiguration resource.
type InfrastructureConfigurationArgs struct {
	// Description for the configuration.
	Description pulumi.StringPtrInput
	// Name of IAM Instance Profile.
	InstanceProfileName pulumi.StringInput
	// Set of EC2 Instance Types.
	InstanceTypes pulumi.StringArrayInput
	// Name of EC2 Key Pair.
	KeyPair pulumi.StringPtrInput
	// Configuration block with logging settings. Detailed below.
	Logging InfrastructureConfigurationLoggingPtrInput
	// Name for the configuration.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags pulumi.StringMapInput
	// Set of EC2 Security Group identifiers.
	SecurityGroupIds pulumi.StringArrayInput
	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn pulumi.StringPtrInput
	// EC2 Subnet identifier. Also requires `securityGroupIds` argument.
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
	TerminateInstanceOnFailure pulumi.BoolPtrInput
}

func (InfrastructureConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*infrastructureConfigurationArgs)(nil)).Elem()
}

type InfrastructureConfigurationInput interface {
	pulumi.Input

	ToInfrastructureConfigurationOutput() InfrastructureConfigurationOutput
	ToInfrastructureConfigurationOutputWithContext(ctx context.Context) InfrastructureConfigurationOutput
}

func (*InfrastructureConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureConfiguration)(nil))
}

func (i *InfrastructureConfiguration) ToInfrastructureConfigurationOutput() InfrastructureConfigurationOutput {
	return i.ToInfrastructureConfigurationOutputWithContext(context.Background())
}

func (i *InfrastructureConfiguration) ToInfrastructureConfigurationOutputWithContext(ctx context.Context) InfrastructureConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfrastructureConfigurationOutput)
}

func (i *InfrastructureConfiguration) ToInfrastructureConfigurationPtrOutput() InfrastructureConfigurationPtrOutput {
	return i.ToInfrastructureConfigurationPtrOutputWithContext(context.Background())
}

func (i *InfrastructureConfiguration) ToInfrastructureConfigurationPtrOutputWithContext(ctx context.Context) InfrastructureConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfrastructureConfigurationPtrOutput)
}

type InfrastructureConfigurationPtrInput interface {
	pulumi.Input

	ToInfrastructureConfigurationPtrOutput() InfrastructureConfigurationPtrOutput
	ToInfrastructureConfigurationPtrOutputWithContext(ctx context.Context) InfrastructureConfigurationPtrOutput
}

type infrastructureConfigurationPtrType InfrastructureConfigurationArgs

func (*infrastructureConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InfrastructureConfiguration)(nil))
}

func (i *infrastructureConfigurationPtrType) ToInfrastructureConfigurationPtrOutput() InfrastructureConfigurationPtrOutput {
	return i.ToInfrastructureConfigurationPtrOutputWithContext(context.Background())
}

func (i *infrastructureConfigurationPtrType) ToInfrastructureConfigurationPtrOutputWithContext(ctx context.Context) InfrastructureConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfrastructureConfigurationPtrOutput)
}

// InfrastructureConfigurationArrayInput is an input type that accepts InfrastructureConfigurationArray and InfrastructureConfigurationArrayOutput values.
// You can construct a concrete instance of `InfrastructureConfigurationArrayInput` via:
//
//          InfrastructureConfigurationArray{ InfrastructureConfigurationArgs{...} }
type InfrastructureConfigurationArrayInput interface {
	pulumi.Input

	ToInfrastructureConfigurationArrayOutput() InfrastructureConfigurationArrayOutput
	ToInfrastructureConfigurationArrayOutputWithContext(context.Context) InfrastructureConfigurationArrayOutput
}

type InfrastructureConfigurationArray []InfrastructureConfigurationInput

func (InfrastructureConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InfrastructureConfiguration)(nil))
}

func (i InfrastructureConfigurationArray) ToInfrastructureConfigurationArrayOutput() InfrastructureConfigurationArrayOutput {
	return i.ToInfrastructureConfigurationArrayOutputWithContext(context.Background())
}

func (i InfrastructureConfigurationArray) ToInfrastructureConfigurationArrayOutputWithContext(ctx context.Context) InfrastructureConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfrastructureConfigurationArrayOutput)
}

// InfrastructureConfigurationMapInput is an input type that accepts InfrastructureConfigurationMap and InfrastructureConfigurationMapOutput values.
// You can construct a concrete instance of `InfrastructureConfigurationMapInput` via:
//
//          InfrastructureConfigurationMap{ "key": InfrastructureConfigurationArgs{...} }
type InfrastructureConfigurationMapInput interface {
	pulumi.Input

	ToInfrastructureConfigurationMapOutput() InfrastructureConfigurationMapOutput
	ToInfrastructureConfigurationMapOutputWithContext(context.Context) InfrastructureConfigurationMapOutput
}

type InfrastructureConfigurationMap map[string]InfrastructureConfigurationInput

func (InfrastructureConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InfrastructureConfiguration)(nil))
}

func (i InfrastructureConfigurationMap) ToInfrastructureConfigurationMapOutput() InfrastructureConfigurationMapOutput {
	return i.ToInfrastructureConfigurationMapOutputWithContext(context.Background())
}

func (i InfrastructureConfigurationMap) ToInfrastructureConfigurationMapOutputWithContext(ctx context.Context) InfrastructureConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfrastructureConfigurationMapOutput)
}

type InfrastructureConfigurationOutput struct {
	*pulumi.OutputState
}

func (InfrastructureConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureConfiguration)(nil))
}

func (o InfrastructureConfigurationOutput) ToInfrastructureConfigurationOutput() InfrastructureConfigurationOutput {
	return o
}

func (o InfrastructureConfigurationOutput) ToInfrastructureConfigurationOutputWithContext(ctx context.Context) InfrastructureConfigurationOutput {
	return o
}

func (o InfrastructureConfigurationOutput) ToInfrastructureConfigurationPtrOutput() InfrastructureConfigurationPtrOutput {
	return o.ToInfrastructureConfigurationPtrOutputWithContext(context.Background())
}

func (o InfrastructureConfigurationOutput) ToInfrastructureConfigurationPtrOutputWithContext(ctx context.Context) InfrastructureConfigurationPtrOutput {
	return o.ApplyT(func(v InfrastructureConfiguration) *InfrastructureConfiguration {
		return &v
	}).(InfrastructureConfigurationPtrOutput)
}

type InfrastructureConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (InfrastructureConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InfrastructureConfiguration)(nil))
}

func (o InfrastructureConfigurationPtrOutput) ToInfrastructureConfigurationPtrOutput() InfrastructureConfigurationPtrOutput {
	return o
}

func (o InfrastructureConfigurationPtrOutput) ToInfrastructureConfigurationPtrOutputWithContext(ctx context.Context) InfrastructureConfigurationPtrOutput {
	return o
}

type InfrastructureConfigurationArrayOutput struct{ *pulumi.OutputState }

func (InfrastructureConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InfrastructureConfiguration)(nil))
}

func (o InfrastructureConfigurationArrayOutput) ToInfrastructureConfigurationArrayOutput() InfrastructureConfigurationArrayOutput {
	return o
}

func (o InfrastructureConfigurationArrayOutput) ToInfrastructureConfigurationArrayOutputWithContext(ctx context.Context) InfrastructureConfigurationArrayOutput {
	return o
}

func (o InfrastructureConfigurationArrayOutput) Index(i pulumi.IntInput) InfrastructureConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InfrastructureConfiguration {
		return vs[0].([]InfrastructureConfiguration)[vs[1].(int)]
	}).(InfrastructureConfigurationOutput)
}

type InfrastructureConfigurationMapOutput struct{ *pulumi.OutputState }

func (InfrastructureConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InfrastructureConfiguration)(nil))
}

func (o InfrastructureConfigurationMapOutput) ToInfrastructureConfigurationMapOutput() InfrastructureConfigurationMapOutput {
	return o
}

func (o InfrastructureConfigurationMapOutput) ToInfrastructureConfigurationMapOutputWithContext(ctx context.Context) InfrastructureConfigurationMapOutput {
	return o
}

func (o InfrastructureConfigurationMapOutput) MapIndex(k pulumi.StringInput) InfrastructureConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InfrastructureConfiguration {
		return vs[0].(map[string]InfrastructureConfiguration)[vs[1].(string)]
	}).(InfrastructureConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(InfrastructureConfigurationOutput{})
	pulumi.RegisterOutputType(InfrastructureConfigurationPtrOutput{})
	pulumi.RegisterOutputType(InfrastructureConfigurationArrayOutput{})
	pulumi.RegisterOutputType(InfrastructureConfigurationMapOutput{})
}

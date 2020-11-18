// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon Managed Streaming for Kafka configuration. More information can be found on the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/msk"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := msk.NewConfiguration(ctx, "example", &msk.ConfigurationArgs{
// 			KafkaVersions: pulumi.StringArray{
// 				pulumi.String("2.1.0"),
// 			},
// 			ServerProperties: pulumi.String(fmt.Sprintf("%v%v%v", "auto.create.topics.enable = true\n", "delete.topic.enable = true\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Configuration struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions pulumi.StringArrayOutput `pulumi:"kafkaVersions"`
	// Latest revision of the configuration.
	LatestRevision pulumi.IntOutput `pulumi:"latestRevision"`
	// Name of the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties pulumi.StringOutput `pulumi:"serverProperties"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KafkaVersions == nil {
		return nil, errors.New("invalid value for required argument 'KafkaVersions'")
	}
	if args.ServerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ServerProperties'")
	}
	var resource Configuration
	err := ctx.RegisterResource("aws:msk/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("aws:msk/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn *string `pulumi:"arn"`
	// Description of the configuration.
	Description *string `pulumi:"description"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []string `pulumi:"kafkaVersions"`
	// Latest revision of the configuration.
	LatestRevision *int `pulumi:"latestRevision"`
	// Name of the configuration.
	Name *string `pulumi:"name"`
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties *string `pulumi:"serverProperties"`
}

type ConfigurationState struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn pulumi.StringPtrInput
	// Description of the configuration.
	Description pulumi.StringPtrInput
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions pulumi.StringArrayInput
	// Latest revision of the configuration.
	LatestRevision pulumi.IntPtrInput
	// Name of the configuration.
	Name pulumi.StringPtrInput
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties pulumi.StringPtrInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// Description of the configuration.
	Description *string `pulumi:"description"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []string `pulumi:"kafkaVersions"`
	// Name of the configuration.
	Name *string `pulumi:"name"`
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties string `pulumi:"serverProperties"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Description of the configuration.
	Description pulumi.StringPtrInput
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions pulumi.StringArrayInput
	// Name of the configuration.
	Name pulumi.StringPtrInput
	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties pulumi.StringInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil)).Elem()
}

func (i Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

type ConfigurationOutput struct {
	*pulumi.OutputState
}

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationOutput)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SES event destination
//
// ## Example Usage
// ### CloudWatch Destination
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewEventDestination(ctx, "cloudwatch", &ses.EventDestinationArgs{
// 			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
// 			Enabled:              pulumi.Bool(true),
// 			MatchingTypes: pulumi.StringArray{
// 				pulumi.String("bounce"),
// 				pulumi.String("send"),
// 			},
// 			CloudwatchDestinations: ses.EventDestinationCloudwatchDestinationArray{
// 				&ses.EventDestinationCloudwatchDestinationArgs{
// 					DefaultValue:  pulumi.String("default"),
// 					DimensionName: pulumi.String("dimension"),
// 					ValueSource:   pulumi.String("emailHeader"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Kinesis Destination
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewEventDestination(ctx, "kinesis", &ses.EventDestinationArgs{
// 			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
// 			Enabled:              pulumi.Bool(true),
// 			MatchingTypes: pulumi.StringArray{
// 				pulumi.String("bounce"),
// 				pulumi.String("send"),
// 			},
// 			KinesisDestination: &ses.EventDestinationKinesisDestinationArgs{
// 				StreamArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
// 				RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### SNS Destination
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewEventDestination(ctx, "sns", &ses.EventDestinationArgs{
// 			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
// 			Enabled:              pulumi.Bool(true),
// 			MatchingTypes: pulumi.StringArray{
// 				pulumi.String("bounce"),
// 				pulumi.String("send"),
// 			},
// 			SnsDestination: &ses.EventDestinationSnsDestinationArgs{
// 				TopicArn: pulumi.Any(aws_sns_topic.Example.Arn),
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
// SES event destinations can be imported using `configuration_set_name` together with the event destination's `name`, e.g.
//
// ```sh
//  $ pulumi import aws:ses/eventDestination:EventDestination sns some-configuration-set-test/event-destination-sns
// ```
type EventDestination struct {
	pulumi.CustomResourceState

	// CloudWatch destination for the events
	CloudwatchDestinations EventDestinationCloudwatchDestinationArrayOutput `pulumi:"cloudwatchDestinations"`
	// The name of the configuration set
	ConfigurationSetName pulumi.StringOutput `pulumi:"configurationSetName"`
	// If true, the event destination will be enabled
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Send the events to a kinesis firehose destination
	KinesisDestination EventDestinationKinesisDestinationPtrOutput `pulumi:"kinesisDestination"`
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes pulumi.StringArrayOutput `pulumi:"matchingTypes"`
	// The name of the event destination
	Name pulumi.StringOutput `pulumi:"name"`
	// Send the events to an SNS Topic destination
	SnsDestination EventDestinationSnsDestinationPtrOutput `pulumi:"snsDestination"`
}

// NewEventDestination registers a new resource with the given unique name, arguments, and options.
func NewEventDestination(ctx *pulumi.Context,
	name string, args *EventDestinationArgs, opts ...pulumi.ResourceOption) (*EventDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationSetName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationSetName'")
	}
	if args.MatchingTypes == nil {
		return nil, errors.New("invalid value for required argument 'MatchingTypes'")
	}
	var resource EventDestination
	err := ctx.RegisterResource("aws:ses/eventDestination:EventDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventDestination gets an existing EventDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventDestinationState, opts ...pulumi.ResourceOption) (*EventDestination, error) {
	var resource EventDestination
	err := ctx.ReadResource("aws:ses/eventDestination:EventDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventDestination resources.
type eventDestinationState struct {
	// CloudWatch destination for the events
	CloudwatchDestinations []EventDestinationCloudwatchDestination `pulumi:"cloudwatchDestinations"`
	// The name of the configuration set
	ConfigurationSetName *string `pulumi:"configurationSetName"`
	// If true, the event destination will be enabled
	Enabled *bool `pulumi:"enabled"`
	// Send the events to a kinesis firehose destination
	KinesisDestination *EventDestinationKinesisDestination `pulumi:"kinesisDestination"`
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes []string `pulumi:"matchingTypes"`
	// The name of the event destination
	Name *string `pulumi:"name"`
	// Send the events to an SNS Topic destination
	SnsDestination *EventDestinationSnsDestination `pulumi:"snsDestination"`
}

type EventDestinationState struct {
	// CloudWatch destination for the events
	CloudwatchDestinations EventDestinationCloudwatchDestinationArrayInput
	// The name of the configuration set
	ConfigurationSetName pulumi.StringPtrInput
	// If true, the event destination will be enabled
	Enabled pulumi.BoolPtrInput
	// Send the events to a kinesis firehose destination
	KinesisDestination EventDestinationKinesisDestinationPtrInput
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes pulumi.StringArrayInput
	// The name of the event destination
	Name pulumi.StringPtrInput
	// Send the events to an SNS Topic destination
	SnsDestination EventDestinationSnsDestinationPtrInput
}

func (EventDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventDestinationState)(nil)).Elem()
}

type eventDestinationArgs struct {
	// CloudWatch destination for the events
	CloudwatchDestinations []EventDestinationCloudwatchDestination `pulumi:"cloudwatchDestinations"`
	// The name of the configuration set
	ConfigurationSetName string `pulumi:"configurationSetName"`
	// If true, the event destination will be enabled
	Enabled *bool `pulumi:"enabled"`
	// Send the events to a kinesis firehose destination
	KinesisDestination *EventDestinationKinesisDestination `pulumi:"kinesisDestination"`
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes []string `pulumi:"matchingTypes"`
	// The name of the event destination
	Name *string `pulumi:"name"`
	// Send the events to an SNS Topic destination
	SnsDestination *EventDestinationSnsDestination `pulumi:"snsDestination"`
}

// The set of arguments for constructing a EventDestination resource.
type EventDestinationArgs struct {
	// CloudWatch destination for the events
	CloudwatchDestinations EventDestinationCloudwatchDestinationArrayInput
	// The name of the configuration set
	ConfigurationSetName pulumi.StringInput
	// If true, the event destination will be enabled
	Enabled pulumi.BoolPtrInput
	// Send the events to a kinesis firehose destination
	KinesisDestination EventDestinationKinesisDestinationPtrInput
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes pulumi.StringArrayInput
	// The name of the event destination
	Name pulumi.StringPtrInput
	// Send the events to an SNS Topic destination
	SnsDestination EventDestinationSnsDestinationPtrInput
}

func (EventDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventDestinationArgs)(nil)).Elem()
}

type EventDestinationInput interface {
	pulumi.Input

	ToEventDestinationOutput() EventDestinationOutput
	ToEventDestinationOutputWithContext(ctx context.Context) EventDestinationOutput
}

func (EventDestination) ElementType() reflect.Type {
	return reflect.TypeOf((*EventDestination)(nil)).Elem()
}

func (i EventDestination) ToEventDestinationOutput() EventDestinationOutput {
	return i.ToEventDestinationOutputWithContext(context.Background())
}

func (i EventDestination) ToEventDestinationOutputWithContext(ctx context.Context) EventDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventDestinationOutput)
}

type EventDestinationOutput struct {
	*pulumi.OutputState
}

func (EventDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventDestinationOutput)(nil)).Elem()
}

func (o EventDestinationOutput) ToEventDestinationOutput() EventDestinationOutput {
	return o
}

func (o EventDestinationOutput) ToEventDestinationOutputWithContext(ctx context.Context) EventDestinationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventDestinationOutput{})
}

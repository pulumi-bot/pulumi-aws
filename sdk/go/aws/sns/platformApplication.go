// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SNS platform application resource
//
// ## Example Usage
// ### Apple Push Notification Service (APNS)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sns.NewPlatformApplication(ctx, "apnsApplication", &sns.PlatformApplicationArgs{
// 			Platform:           pulumi.String("APNS"),
// 			PlatformCredential: pulumi.String("<APNS PRIVATE KEY>"),
// 			PlatformPrincipal:  pulumi.String("<APNS CERTIFICATE>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Google Cloud Messaging (GCM)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sns.NewPlatformApplication(ctx, "gcmApplication", &sns.PlatformApplicationArgs{
// 			Platform:           pulumi.String("GCM"),
// 			PlatformCredential: pulumi.String("<GCM API KEY>"),
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
// SNS platform applications can be imported using the ARN, e.g.
//
// ```sh
//  $ pulumi import aws:sns/platformApplication:PlatformApplication gcm_application arn:aws:sns:us-west-2:0123456789012:app/GCM/gcm_application
// ```
type PlatformApplication struct {
	pulumi.CustomResourceState

	// The ARN of the SNS platform application
	Arn pulumi.StringOutput `pulumi:"arn"`
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn pulumi.StringPtrOutput `pulumi:"eventDeliveryFailureTopicArn"`
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn pulumi.StringPtrOutput `pulumi:"eventEndpointCreatedTopicArn"`
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn pulumi.StringPtrOutput `pulumi:"eventEndpointDeletedTopicArn"`
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn pulumi.StringPtrOutput `pulumi:"eventEndpointUpdatedTopicArn"`
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn pulumi.StringPtrOutput `pulumi:"failureFeedbackRoleArn"`
	// The friendly name for the SNS platform application
	Name pulumi.StringOutput `pulumi:"name"`
	// The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential pulumi.StringOutput `pulumi:"platformCredential"`
	// Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal pulumi.StringPtrOutput `pulumi:"platformPrincipal"`
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn pulumi.StringPtrOutput `pulumi:"successFeedbackRoleArn"`
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate pulumi.StringPtrOutput `pulumi:"successFeedbackSampleRate"`
}

// NewPlatformApplication registers a new resource with the given unique name, arguments, and options.
func NewPlatformApplication(ctx *pulumi.Context,
	name string, args *PlatformApplicationArgs, opts ...pulumi.ResourceOption) (*PlatformApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.PlatformCredential == nil {
		return nil, errors.New("invalid value for required argument 'PlatformCredential'")
	}
	var resource PlatformApplication
	err := ctx.RegisterResource("aws:sns/platformApplication:PlatformApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlatformApplication gets an existing PlatformApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlatformApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlatformApplicationState, opts ...pulumi.ResourceOption) (*PlatformApplication, error) {
	var resource PlatformApplication
	err := ctx.ReadResource("aws:sns/platformApplication:PlatformApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlatformApplication resources.
type platformApplicationState struct {
	// The ARN of the SNS platform application
	Arn *string `pulumi:"arn"`
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn *string `pulumi:"eventDeliveryFailureTopicArn"`
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn *string `pulumi:"eventEndpointCreatedTopicArn"`
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn *string `pulumi:"eventEndpointDeletedTopicArn"`
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn *string `pulumi:"eventEndpointUpdatedTopicArn"`
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn *string `pulumi:"failureFeedbackRoleArn"`
	// The friendly name for the SNS platform application
	Name *string `pulumi:"name"`
	// The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
	Platform *string `pulumi:"platform"`
	// Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential *string `pulumi:"platformCredential"`
	// Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal *string `pulumi:"platformPrincipal"`
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn *string `pulumi:"successFeedbackRoleArn"`
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate *string `pulumi:"successFeedbackSampleRate"`
}

type PlatformApplicationState struct {
	// The ARN of the SNS platform application
	Arn pulumi.StringPtrInput
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn pulumi.StringPtrInput
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn pulumi.StringPtrInput
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn pulumi.StringPtrInput
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn pulumi.StringPtrInput
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn pulumi.StringPtrInput
	// The friendly name for the SNS platform application
	Name pulumi.StringPtrInput
	// The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
	Platform pulumi.StringPtrInput
	// Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential pulumi.StringPtrInput
	// Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal pulumi.StringPtrInput
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn pulumi.StringPtrInput
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate pulumi.StringPtrInput
}

func (PlatformApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*platformApplicationState)(nil)).Elem()
}

type platformApplicationArgs struct {
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn *string `pulumi:"eventDeliveryFailureTopicArn"`
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn *string `pulumi:"eventEndpointCreatedTopicArn"`
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn *string `pulumi:"eventEndpointDeletedTopicArn"`
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn *string `pulumi:"eventEndpointUpdatedTopicArn"`
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn *string `pulumi:"failureFeedbackRoleArn"`
	// The friendly name for the SNS platform application
	Name *string `pulumi:"name"`
	// The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
	Platform string `pulumi:"platform"`
	// Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential string `pulumi:"platformCredential"`
	// Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal *string `pulumi:"platformPrincipal"`
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn *string `pulumi:"successFeedbackRoleArn"`
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate *string `pulumi:"successFeedbackSampleRate"`
}

// The set of arguments for constructing a PlatformApplication resource.
type PlatformApplicationArgs struct {
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn pulumi.StringPtrInput
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn pulumi.StringPtrInput
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn pulumi.StringPtrInput
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn pulumi.StringPtrInput
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn pulumi.StringPtrInput
	// The friendly name for the SNS platform application
	Name pulumi.StringPtrInput
	// The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
	Platform pulumi.StringInput
	// Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential pulumi.StringInput
	// Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal pulumi.StringPtrInput
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn pulumi.StringPtrInput
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate pulumi.StringPtrInput
}

func (PlatformApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*platformApplicationArgs)(nil)).Elem()
}

type PlatformApplicationInput interface {
	pulumi.Input

	ToPlatformApplicationOutput() PlatformApplicationOutput
	ToPlatformApplicationOutputWithContext(ctx context.Context) PlatformApplicationOutput
}

func (*PlatformApplication) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformApplication)(nil))
}

func (i *PlatformApplication) ToPlatformApplicationOutput() PlatformApplicationOutput {
	return i.ToPlatformApplicationOutputWithContext(context.Background())
}

func (i *PlatformApplication) ToPlatformApplicationOutputWithContext(ctx context.Context) PlatformApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformApplicationOutput)
}

func (i *PlatformApplication) ToPlatformApplicationPtrOutput() PlatformApplicationPtrOutput {
	return i.ToPlatformApplicationPtrOutputWithContext(context.Background())
}

func (i *PlatformApplication) ToPlatformApplicationPtrOutputWithContext(ctx context.Context) PlatformApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformApplicationPtrOutput)
}

type PlatformApplicationPtrInput interface {
	pulumi.Input

	ToPlatformApplicationPtrOutput() PlatformApplicationPtrOutput
	ToPlatformApplicationPtrOutputWithContext(ctx context.Context) PlatformApplicationPtrOutput
}

type platformApplicationPtrType PlatformApplicationArgs

func (*platformApplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformApplication)(nil))
}

func (i *platformApplicationPtrType) ToPlatformApplicationPtrOutput() PlatformApplicationPtrOutput {
	return i.ToPlatformApplicationPtrOutputWithContext(context.Background())
}

func (i *platformApplicationPtrType) ToPlatformApplicationPtrOutputWithContext(ctx context.Context) PlatformApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformApplicationPtrOutput)
}

// PlatformApplicationArrayInput is an input type that accepts PlatformApplicationArray and PlatformApplicationArrayOutput values.
// You can construct a concrete instance of `PlatformApplicationArrayInput` via:
//
//          PlatformApplicationArray{ PlatformApplicationArgs{...} }
type PlatformApplicationArrayInput interface {
	pulumi.Input

	ToPlatformApplicationArrayOutput() PlatformApplicationArrayOutput
	ToPlatformApplicationArrayOutputWithContext(context.Context) PlatformApplicationArrayOutput
}

type PlatformApplicationArray []PlatformApplicationInput

func (PlatformApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PlatformApplication)(nil))
}

func (i PlatformApplicationArray) ToPlatformApplicationArrayOutput() PlatformApplicationArrayOutput {
	return i.ToPlatformApplicationArrayOutputWithContext(context.Background())
}

func (i PlatformApplicationArray) ToPlatformApplicationArrayOutputWithContext(ctx context.Context) PlatformApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformApplicationArrayOutput)
}

// PlatformApplicationMapInput is an input type that accepts PlatformApplicationMap and PlatformApplicationMapOutput values.
// You can construct a concrete instance of `PlatformApplicationMapInput` via:
//
//          PlatformApplicationMap{ "key": PlatformApplicationArgs{...} }
type PlatformApplicationMapInput interface {
	pulumi.Input

	ToPlatformApplicationMapOutput() PlatformApplicationMapOutput
	ToPlatformApplicationMapOutputWithContext(context.Context) PlatformApplicationMapOutput
}

type PlatformApplicationMap map[string]PlatformApplicationInput

func (PlatformApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PlatformApplication)(nil))
}

func (i PlatformApplicationMap) ToPlatformApplicationMapOutput() PlatformApplicationMapOutput {
	return i.ToPlatformApplicationMapOutputWithContext(context.Background())
}

func (i PlatformApplicationMap) ToPlatformApplicationMapOutputWithContext(ctx context.Context) PlatformApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformApplicationMapOutput)
}

type PlatformApplicationOutput struct {
	*pulumi.OutputState
}

func (PlatformApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformApplication)(nil))
}

func (o PlatformApplicationOutput) ToPlatformApplicationOutput() PlatformApplicationOutput {
	return o
}

func (o PlatformApplicationOutput) ToPlatformApplicationOutputWithContext(ctx context.Context) PlatformApplicationOutput {
	return o
}

func (o PlatformApplicationOutput) ToPlatformApplicationPtrOutput() PlatformApplicationPtrOutput {
	return o.ToPlatformApplicationPtrOutputWithContext(context.Background())
}

func (o PlatformApplicationOutput) ToPlatformApplicationPtrOutputWithContext(ctx context.Context) PlatformApplicationPtrOutput {
	return o.ApplyT(func(v PlatformApplication) *PlatformApplication {
		return &v
	}).(PlatformApplicationPtrOutput)
}

type PlatformApplicationPtrOutput struct {
	*pulumi.OutputState
}

func (PlatformApplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformApplication)(nil))
}

func (o PlatformApplicationPtrOutput) ToPlatformApplicationPtrOutput() PlatformApplicationPtrOutput {
	return o
}

func (o PlatformApplicationPtrOutput) ToPlatformApplicationPtrOutputWithContext(ctx context.Context) PlatformApplicationPtrOutput {
	return o
}

type PlatformApplicationArrayOutput struct{ *pulumi.OutputState }

func (PlatformApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlatformApplication)(nil))
}

func (o PlatformApplicationArrayOutput) ToPlatformApplicationArrayOutput() PlatformApplicationArrayOutput {
	return o
}

func (o PlatformApplicationArrayOutput) ToPlatformApplicationArrayOutputWithContext(ctx context.Context) PlatformApplicationArrayOutput {
	return o
}

func (o PlatformApplicationArrayOutput) Index(i pulumi.IntInput) PlatformApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlatformApplication {
		return vs[0].([]PlatformApplication)[vs[1].(int)]
	}).(PlatformApplicationOutput)
}

type PlatformApplicationMapOutput struct{ *pulumi.OutputState }

func (PlatformApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PlatformApplication)(nil))
}

func (o PlatformApplicationMapOutput) ToPlatformApplicationMapOutput() PlatformApplicationMapOutput {
	return o
}

func (o PlatformApplicationMapOutput) ToPlatformApplicationMapOutputWithContext(ctx context.Context) PlatformApplicationMapOutput {
	return o
}

func (o PlatformApplicationMapOutput) MapIndex(k pulumi.StringInput) PlatformApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PlatformApplication {
		return vs[0].(map[string]PlatformApplication)[vs[1].(string)]
	}).(PlatformApplicationOutput)
}

func init() {
	pulumi.RegisterOutputType(PlatformApplicationOutput{})
	pulumi.RegisterOutputType(PlatformApplicationPtrOutput{})
	pulumi.RegisterOutputType(PlatformApplicationArrayOutput{})
	pulumi.RegisterOutputType(PlatformApplicationMapOutput{})
}

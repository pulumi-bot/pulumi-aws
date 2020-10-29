// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Amazon Lex Intent resource. For more information see
// [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lex"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lex.NewIntent(ctx, "orderFlowersIntent", &lex.IntentArgs{
// 			ConfirmationPrompt: &lex.IntentConfirmationPromptArgs{
// 				MaxAttempts: pulumi.Int(2),
// 				Messages: lex.IntentConfirmationPromptMessageArray{
// 					&lex.IntentConfirmationPromptMessageArgs{
// 						Content:     pulumi.String("Okay, your {FlowerType} will be ready for pickup by {PickupTime} on {PickupDate}.  Does this sound okay?"),
// 						ContentType: pulumi.String("PlainText"),
// 					},
// 				},
// 			},
// 			CreateVersion: pulumi.Bool(false),
// 			Description:   pulumi.String("Intent to order a bouquet of flowers for pick up"),
// 			FulfillmentActivity: &lex.IntentFulfillmentActivityArgs{
// 				Type: pulumi.String("ReturnIntent"),
// 			},
// 			RejectionStatement: &lex.IntentRejectionStatementArgs{
// 				Messages: lex.IntentRejectionStatementMessageArray{
// 					&lex.IntentRejectionStatementMessageArgs{
// 						Content:     pulumi.String("Okay, I will not place your order."),
// 						ContentType: pulumi.String("PlainText"),
// 					},
// 				},
// 			},
// 			SampleUtterances: pulumi.StringArray{
// 				pulumi.String("I would like to order some flowers"),
// 				pulumi.String("I would like to pick up flowers"),
// 			},
// 			Slots: lex.IntentSlotArray{
// 				&lex.IntentSlotArgs{
// 					Description: pulumi.String("The type of flowers to pick up"),
// 					Name:        pulumi.String("FlowerType"),
// 					Priority:    pulumi.Int(1),
// 					SampleUtterances: pulumi.StringArray{
// 						pulumi.String("I would like to order {FlowerType}"),
// 					},
// 					SlotConstraint:  pulumi.String("Required"),
// 					SlotType:        pulumi.String("FlowerTypes"),
// 					SlotTypeVersion: pulumi.String(fmt.Sprintf("%v%v", "$", "LATEST")),
// 					ValueElicitationPrompt: &lex.IntentSlotValueElicitationPromptArgs{
// 						MaxAttempts: pulumi.Int(2),
// 						Message: pulumi.StringMapArray{
// 							pulumi.StringMap{
// 								"content":     pulumi.String("What type of flowers would you like to order?"),
// 								"contentType": pulumi.String("PlainText"),
// 							},
// 						},
// 					},
// 				},
// 				&lex.IntentSlotArgs{
// 					Description: pulumi.String("The date to pick up the flowers"),
// 					Name:        pulumi.String("PickupDate"),
// 					Priority:    pulumi.Int(2),
// 					SampleUtterances: pulumi.StringArray{
// 						pulumi.String("I would like to order {FlowerType}"),
// 					},
// 					SlotConstraint:  pulumi.String("Required"),
// 					SlotType:        pulumi.String("AMAZON.DATE"),
// 					SlotTypeVersion: pulumi.String(fmt.Sprintf("%v%v", "$", "LATEST")),
// 					ValueElicitationPrompt: &lex.IntentSlotValueElicitationPromptArgs{
// 						MaxAttempts: pulumi.Int(2),
// 						Message: pulumi.StringMapArray{
// 							pulumi.StringMap{
// 								"content":     pulumi.String("What day do you want the {FlowerType} to be picked up?"),
// 								"contentType": pulumi.String("PlainText"),
// 							},
// 						},
// 					},
// 				},
// 				&lex.IntentSlotArgs{
// 					Description: pulumi.String("The time to pick up the flowers"),
// 					Name:        pulumi.String("PickupTime"),
// 					Priority:    pulumi.Int(3),
// 					SampleUtterances: pulumi.StringArray{
// 						pulumi.String("I would like to order {FlowerType}"),
// 					},
// 					SlotConstraint:  pulumi.String("Required"),
// 					SlotType:        pulumi.String("AMAZON.TIME"),
// 					SlotTypeVersion: pulumi.String(fmt.Sprintf("%v%v", "$", "LATEST")),
// 					ValueElicitationPrompt: &lex.IntentSlotValueElicitationPromptArgs{
// 						MaxAttempts: pulumi.Int(2),
// 						Message: pulumi.StringMapArray{
// 							pulumi.StringMap{
// 								"content":     pulumi.String("Pick up the {FlowerType} at what time on {PickupDate}?"),
// 								"contentType": pulumi.String("PlainText"),
// 							},
// 						},
// 					},
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
type Intent struct {
	pulumi.CustomResourceState

	// The ARN of the Lex intent.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Checksum identifying the version of the intent that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the intent.
	Checksum pulumi.StringOutput `pulumi:"checksum"`
	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the `fulfillmentActivity`. If you return the intent to the client
	// application, you can't specify this element. The `followUpPrompt` and `conclusionStatement` are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	ConclusionStatement IntentConclusionStatementPtrOutput `pulumi:"conclusionStatement"`
	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the `rejectionStatement` and `confirmationPrompt`,
	// or neither. Attributes are documented under prompt.
	ConfirmationPrompt IntentConfirmationPromptPtrOutput `pulumi:"confirmationPrompt"`
	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrOutput `pulumi:"createVersion"`
	// The date when the intent version was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// A description of the bot.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	DialogCodeHook IntentDialogCodeHookPtrOutput `pulumi:"dialogCodeHook"`
	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The `followUpPrompt` field and the `conclusionStatement` field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	FollowUpPrompt IntentFollowUpPromptPtrOutput `pulumi:"followUpPrompt"`
	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, `fulfillmentActivity` defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	FulfillmentActivity IntentFulfillmentActivityOutput `pulumi:"fulfillmentActivity"`
	// The date when the $LATEST version of this intent was updated.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the intent slot that you want to create. The name is case sensitive.
	Name pulumi.StringOutput `pulumi:"name"`
	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature pulumi.StringPtrOutput `pulumi:"parentIntentSignature"`
	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	RejectionStatement IntentRejectionStatementPtrOutput `pulumi:"rejectionStatement"`
	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances pulumi.StringArrayOutput `pulumi:"sampleUtterances"`
	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	Slots IntentSlotArrayOutput `pulumi:"slots"`
	// The version of the bot.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewIntent registers a new resource with the given unique name, arguments, and options.
func NewIntent(ctx *pulumi.Context,
	name string, args *IntentArgs, opts ...pulumi.ResourceOption) (*Intent, error) {
	if args == nil || args.FulfillmentActivity == nil {
		return nil, errors.New("missing required argument 'FulfillmentActivity'")
	}
	if args == nil {
		args = &IntentArgs{}
	}
	var resource Intent
	err := ctx.RegisterResource("aws:lex/intent:Intent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntent gets an existing Intent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntentState, opts ...pulumi.ResourceOption) (*Intent, error) {
	var resource Intent
	err := ctx.ReadResource("aws:lex/intent:Intent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Intent resources.
type intentState struct {
	// The ARN of the Lex intent.
	Arn *string `pulumi:"arn"`
	// Checksum identifying the version of the intent that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the intent.
	Checksum *string `pulumi:"checksum"`
	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the `fulfillmentActivity`. If you return the intent to the client
	// application, you can't specify this element. The `followUpPrompt` and `conclusionStatement` are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	ConclusionStatement *IntentConclusionStatement `pulumi:"conclusionStatement"`
	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the `rejectionStatement` and `confirmationPrompt`,
	// or neither. Attributes are documented under prompt.
	ConfirmationPrompt *IntentConfirmationPrompt `pulumi:"confirmationPrompt"`
	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to `false`.
	CreateVersion *bool `pulumi:"createVersion"`
	// The date when the intent version was created.
	CreatedDate *string `pulumi:"createdDate"`
	// A description of the bot.
	Description *string `pulumi:"description"`
	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	DialogCodeHook *IntentDialogCodeHook `pulumi:"dialogCodeHook"`
	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The `followUpPrompt` field and the `conclusionStatement` field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	FollowUpPrompt *IntentFollowUpPrompt `pulumi:"followUpPrompt"`
	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, `fulfillmentActivity` defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	FulfillmentActivity *IntentFulfillmentActivity `pulumi:"fulfillmentActivity"`
	// The date when the $LATEST version of this intent was updated.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the intent slot that you want to create. The name is case sensitive.
	Name *string `pulumi:"name"`
	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature *string `pulumi:"parentIntentSignature"`
	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	RejectionStatement *IntentRejectionStatement `pulumi:"rejectionStatement"`
	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances []string `pulumi:"sampleUtterances"`
	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	Slots []IntentSlot `pulumi:"slots"`
	// The version of the bot.
	Version *string `pulumi:"version"`
}

type IntentState struct {
	// The ARN of the Lex intent.
	Arn pulumi.StringPtrInput
	// Checksum identifying the version of the intent that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the intent.
	Checksum pulumi.StringPtrInput
	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the `fulfillmentActivity`. If you return the intent to the client
	// application, you can't specify this element. The `followUpPrompt` and `conclusionStatement` are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	ConclusionStatement IntentConclusionStatementPtrInput
	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the `rejectionStatement` and `confirmationPrompt`,
	// or neither. Attributes are documented under prompt.
	ConfirmationPrompt IntentConfirmationPromptPtrInput
	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrInput
	// The date when the intent version was created.
	CreatedDate pulumi.StringPtrInput
	// A description of the bot.
	Description pulumi.StringPtrInput
	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	DialogCodeHook IntentDialogCodeHookPtrInput
	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The `followUpPrompt` field and the `conclusionStatement` field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	FollowUpPrompt IntentFollowUpPromptPtrInput
	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, `fulfillmentActivity` defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	FulfillmentActivity IntentFulfillmentActivityPtrInput
	// The date when the $LATEST version of this intent was updated.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the intent slot that you want to create. The name is case sensitive.
	Name pulumi.StringPtrInput
	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature pulumi.StringPtrInput
	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	RejectionStatement IntentRejectionStatementPtrInput
	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances pulumi.StringArrayInput
	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	Slots IntentSlotArrayInput
	// The version of the bot.
	Version pulumi.StringPtrInput
}

func (IntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*intentState)(nil)).Elem()
}

type intentArgs struct {
	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the `fulfillmentActivity`. If you return the intent to the client
	// application, you can't specify this element. The `followUpPrompt` and `conclusionStatement` are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	ConclusionStatement *IntentConclusionStatement `pulumi:"conclusionStatement"`
	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the `rejectionStatement` and `confirmationPrompt`,
	// or neither. Attributes are documented under prompt.
	ConfirmationPrompt *IntentConfirmationPrompt `pulumi:"confirmationPrompt"`
	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to `false`.
	CreateVersion *bool `pulumi:"createVersion"`
	// A description of the bot.
	Description *string `pulumi:"description"`
	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	DialogCodeHook *IntentDialogCodeHook `pulumi:"dialogCodeHook"`
	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The `followUpPrompt` field and the `conclusionStatement` field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	FollowUpPrompt *IntentFollowUpPrompt `pulumi:"followUpPrompt"`
	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, `fulfillmentActivity` defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	FulfillmentActivity IntentFulfillmentActivity `pulumi:"fulfillmentActivity"`
	// The name of the intent slot that you want to create. The name is case sensitive.
	Name *string `pulumi:"name"`
	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature *string `pulumi:"parentIntentSignature"`
	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	RejectionStatement *IntentRejectionStatement `pulumi:"rejectionStatement"`
	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances []string `pulumi:"sampleUtterances"`
	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	Slots []IntentSlot `pulumi:"slots"`
}

// The set of arguments for constructing a Intent resource.
type IntentArgs struct {
	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the `fulfillmentActivity`. If you return the intent to the client
	// application, you can't specify this element. The `followUpPrompt` and `conclusionStatement` are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	ConclusionStatement IntentConclusionStatementPtrInput
	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the `rejectionStatement` and `confirmationPrompt`,
	// or neither. Attributes are documented under prompt.
	ConfirmationPrompt IntentConfirmationPromptPtrInput
	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to `false`.
	CreateVersion pulumi.BoolPtrInput
	// A description of the bot.
	Description pulumi.StringPtrInput
	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	DialogCodeHook IntentDialogCodeHookPtrInput
	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The `followUpPrompt` field and the `conclusionStatement` field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	FollowUpPrompt IntentFollowUpPromptPtrInput
	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, `fulfillmentActivity` defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	FulfillmentActivity IntentFulfillmentActivityInput
	// The name of the intent slot that you want to create. The name is case sensitive.
	Name pulumi.StringPtrInput
	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature pulumi.StringPtrInput
	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	RejectionStatement IntentRejectionStatementPtrInput
	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances pulumi.StringArrayInput
	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	Slots IntentSlotArrayInput
}

func (IntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*intentArgs)(nil)).Elem()
}

type IntentInput interface {
	pulumi.Input

	ToIntentOutput() IntentOutput
	ToIntentOutputWithContext(ctx context.Context) IntentOutput
}

func (Intent) ElementType() reflect.Type {
	return reflect.TypeOf((*Intent)(nil)).Elem()
}

func (i Intent) ToIntentOutput() IntentOutput {
	return i.ToIntentOutputWithContext(context.Background())
}

func (i Intent) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentOutput)
}

type IntentOutput struct {
	*pulumi.OutputState
}

func (IntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntentOutput)(nil)).Elem()
}

func (o IntentOutput) ToIntentOutput() IntentOutput {
	return o
}

func (o IntentOutput) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntentOutput{})
}

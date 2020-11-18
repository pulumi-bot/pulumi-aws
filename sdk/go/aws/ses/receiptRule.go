// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SES receipt rule resource
//
// ## Example Usage
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
// 		_, err := ses.NewReceiptRule(ctx, "store", &ses.ReceiptRuleArgs{
// 			AddHeaderActions: ses.ReceiptRuleAddHeaderActionArray{
// 				&ses.ReceiptRuleAddHeaderActionArgs{
// 					HeaderName:  pulumi.String("Custom-Header"),
// 					HeaderValue: pulumi.String("Added by SES"),
// 					Position:    pulumi.Int(1),
// 				},
// 			},
// 			Enabled: pulumi.Bool(true),
// 			Recipients: pulumi.StringArray{
// 				pulumi.String("karen@example.com"),
// 			},
// 			RuleSetName: pulumi.String("default-rule-set"),
// 			S3Actions: ses.ReceiptRuleS3ActionArray{
// 				&ses.ReceiptRuleS3ActionArgs{
// 					BucketName: pulumi.String("emails"),
// 					Position:   pulumi.Int(2),
// 				},
// 			},
// 			ScanEnabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ReceiptRule struct {
	pulumi.CustomResourceState

	// A list of Add Header Action blocks. Documented below.
	AddHeaderActions ReceiptRuleAddHeaderActionArrayOutput `pulumi:"addHeaderActions"`
	// The name of the rule to place this rule after
	After pulumi.StringPtrOutput `pulumi:"after"`
	// A list of Bounce Action blocks. Documented below.
	BounceActions ReceiptRuleBounceActionArrayOutput `pulumi:"bounceActions"`
	// If true, the rule will be enabled
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// A list of Lambda Action blocks. Documented below.
	LambdaActions ReceiptRuleLambdaActionArrayOutput `pulumi:"lambdaActions"`
	// The name of the rule
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of email addresses
	Recipients pulumi.StringArrayOutput `pulumi:"recipients"`
	// The name of the rule set
	RuleSetName pulumi.StringOutput `pulumi:"ruleSetName"`
	// A list of S3 Action blocks. Documented below.
	S3Actions ReceiptRuleS3ActionArrayOutput `pulumi:"s3Actions"`
	// If true, incoming emails will be scanned for spam and viruses
	ScanEnabled pulumi.BoolOutput `pulumi:"scanEnabled"`
	// A list of SNS Action blocks. Documented below.
	SnsActions ReceiptRuleSnsActionArrayOutput `pulumi:"snsActions"`
	// A list of Stop Action blocks. Documented below.
	StopActions ReceiptRuleStopActionArrayOutput `pulumi:"stopActions"`
	// Require or Optional
	TlsPolicy pulumi.StringOutput `pulumi:"tlsPolicy"`
	// A list of WorkMail Action blocks. Documented below.
	WorkmailActions ReceiptRuleWorkmailActionArrayOutput `pulumi:"workmailActions"`
}

// NewReceiptRule registers a new resource with the given unique name, arguments, and options.
func NewReceiptRule(ctx *pulumi.Context,
	name string, args *ReceiptRuleArgs, opts ...pulumi.ResourceOption) (*ReceiptRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuleSetName == nil {
		return nil, errors.New("invalid value for required argument 'RuleSetName'")
	}
	var resource ReceiptRule
	err := ctx.RegisterResource("aws:ses/receiptRule:ReceiptRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiptRule gets an existing ReceiptRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiptRuleState, opts ...pulumi.ResourceOption) (*ReceiptRule, error) {
	var resource ReceiptRule
	err := ctx.ReadResource("aws:ses/receiptRule:ReceiptRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReceiptRule resources.
type receiptRuleState struct {
	// A list of Add Header Action blocks. Documented below.
	AddHeaderActions []ReceiptRuleAddHeaderAction `pulumi:"addHeaderActions"`
	// The name of the rule to place this rule after
	After *string `pulumi:"after"`
	// A list of Bounce Action blocks. Documented below.
	BounceActions []ReceiptRuleBounceAction `pulumi:"bounceActions"`
	// If true, the rule will be enabled
	Enabled *bool `pulumi:"enabled"`
	// A list of Lambda Action blocks. Documented below.
	LambdaActions []ReceiptRuleLambdaAction `pulumi:"lambdaActions"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// A list of email addresses
	Recipients []string `pulumi:"recipients"`
	// The name of the rule set
	RuleSetName *string `pulumi:"ruleSetName"`
	// A list of S3 Action blocks. Documented below.
	S3Actions []ReceiptRuleS3Action `pulumi:"s3Actions"`
	// If true, incoming emails will be scanned for spam and viruses
	ScanEnabled *bool `pulumi:"scanEnabled"`
	// A list of SNS Action blocks. Documented below.
	SnsActions []ReceiptRuleSnsAction `pulumi:"snsActions"`
	// A list of Stop Action blocks. Documented below.
	StopActions []ReceiptRuleStopAction `pulumi:"stopActions"`
	// Require or Optional
	TlsPolicy *string `pulumi:"tlsPolicy"`
	// A list of WorkMail Action blocks. Documented below.
	WorkmailActions []ReceiptRuleWorkmailAction `pulumi:"workmailActions"`
}

type ReceiptRuleState struct {
	// A list of Add Header Action blocks. Documented below.
	AddHeaderActions ReceiptRuleAddHeaderActionArrayInput
	// The name of the rule to place this rule after
	After pulumi.StringPtrInput
	// A list of Bounce Action blocks. Documented below.
	BounceActions ReceiptRuleBounceActionArrayInput
	// If true, the rule will be enabled
	Enabled pulumi.BoolPtrInput
	// A list of Lambda Action blocks. Documented below.
	LambdaActions ReceiptRuleLambdaActionArrayInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// A list of email addresses
	Recipients pulumi.StringArrayInput
	// The name of the rule set
	RuleSetName pulumi.StringPtrInput
	// A list of S3 Action blocks. Documented below.
	S3Actions ReceiptRuleS3ActionArrayInput
	// If true, incoming emails will be scanned for spam and viruses
	ScanEnabled pulumi.BoolPtrInput
	// A list of SNS Action blocks. Documented below.
	SnsActions ReceiptRuleSnsActionArrayInput
	// A list of Stop Action blocks. Documented below.
	StopActions ReceiptRuleStopActionArrayInput
	// Require or Optional
	TlsPolicy pulumi.StringPtrInput
	// A list of WorkMail Action blocks. Documented below.
	WorkmailActions ReceiptRuleWorkmailActionArrayInput
}

func (ReceiptRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleState)(nil)).Elem()
}

type receiptRuleArgs struct {
	// A list of Add Header Action blocks. Documented below.
	AddHeaderActions []ReceiptRuleAddHeaderAction `pulumi:"addHeaderActions"`
	// The name of the rule to place this rule after
	After *string `pulumi:"after"`
	// A list of Bounce Action blocks. Documented below.
	BounceActions []ReceiptRuleBounceAction `pulumi:"bounceActions"`
	// If true, the rule will be enabled
	Enabled *bool `pulumi:"enabled"`
	// A list of Lambda Action blocks. Documented below.
	LambdaActions []ReceiptRuleLambdaAction `pulumi:"lambdaActions"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// A list of email addresses
	Recipients []string `pulumi:"recipients"`
	// The name of the rule set
	RuleSetName string `pulumi:"ruleSetName"`
	// A list of S3 Action blocks. Documented below.
	S3Actions []ReceiptRuleS3Action `pulumi:"s3Actions"`
	// If true, incoming emails will be scanned for spam and viruses
	ScanEnabled *bool `pulumi:"scanEnabled"`
	// A list of SNS Action blocks. Documented below.
	SnsActions []ReceiptRuleSnsAction `pulumi:"snsActions"`
	// A list of Stop Action blocks. Documented below.
	StopActions []ReceiptRuleStopAction `pulumi:"stopActions"`
	// Require or Optional
	TlsPolicy *string `pulumi:"tlsPolicy"`
	// A list of WorkMail Action blocks. Documented below.
	WorkmailActions []ReceiptRuleWorkmailAction `pulumi:"workmailActions"`
}

// The set of arguments for constructing a ReceiptRule resource.
type ReceiptRuleArgs struct {
	// A list of Add Header Action blocks. Documented below.
	AddHeaderActions ReceiptRuleAddHeaderActionArrayInput
	// The name of the rule to place this rule after
	After pulumi.StringPtrInput
	// A list of Bounce Action blocks. Documented below.
	BounceActions ReceiptRuleBounceActionArrayInput
	// If true, the rule will be enabled
	Enabled pulumi.BoolPtrInput
	// A list of Lambda Action blocks. Documented below.
	LambdaActions ReceiptRuleLambdaActionArrayInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// A list of email addresses
	Recipients pulumi.StringArrayInput
	// The name of the rule set
	RuleSetName pulumi.StringInput
	// A list of S3 Action blocks. Documented below.
	S3Actions ReceiptRuleS3ActionArrayInput
	// If true, incoming emails will be scanned for spam and viruses
	ScanEnabled pulumi.BoolPtrInput
	// A list of SNS Action blocks. Documented below.
	SnsActions ReceiptRuleSnsActionArrayInput
	// A list of Stop Action blocks. Documented below.
	StopActions ReceiptRuleStopActionArrayInput
	// Require or Optional
	TlsPolicy pulumi.StringPtrInput
	// A list of WorkMail Action blocks. Documented below.
	WorkmailActions ReceiptRuleWorkmailActionArrayInput
}

func (ReceiptRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleArgs)(nil)).Elem()
}

type ReceiptRuleInput interface {
	pulumi.Input

	ToReceiptRuleOutput() ReceiptRuleOutput
	ToReceiptRuleOutputWithContext(ctx context.Context) ReceiptRuleOutput
}

func (ReceiptRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRule)(nil)).Elem()
}

func (i ReceiptRule) ToReceiptRuleOutput() ReceiptRuleOutput {
	return i.ToReceiptRuleOutputWithContext(context.Background())
}

func (i ReceiptRule) ToReceiptRuleOutputWithContext(ctx context.Context) ReceiptRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleOutput)
}

type ReceiptRuleOutput struct {
	*pulumi.OutputState
}

func (ReceiptRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRuleOutput)(nil)).Elem()
}

func (o ReceiptRuleOutput) ToReceiptRuleOutput() ReceiptRuleOutput {
	return o
}

func (o ReceiptRuleOutput) ToReceiptRuleOutputWithContext(ctx context.Context) ReceiptRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReceiptRuleOutput{})
}

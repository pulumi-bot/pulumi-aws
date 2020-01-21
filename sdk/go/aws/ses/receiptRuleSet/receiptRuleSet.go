// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package receiptRuleSet

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES receipt rule set resource
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_receipt_rule_set.html.markdown.
type ReceiptRuleSet struct {
	pulumi.CustomResourceState

	// The name of the rule set
	RuleSetName pulumi.StringOutput `pulumi:"ruleSetName"`
}

// NewReceiptRuleSet registers a new resource with the given unique name, arguments, and options.
func NewReceiptRuleSet(ctx *pulumi.Context,
	name string, args *ReceiptRuleSetArgs, opts ...pulumi.ResourceOption) (*ReceiptRuleSet, error) {
	if args == nil || args.RuleSetName == nil {
		return nil, errors.New("missing required argument 'RuleSetName'")
	}
	if args == nil {
		args = &ReceiptRuleSetArgs{}
	}
	var resource ReceiptRuleSet
	err := ctx.RegisterResource("aws:ses/receiptRuleSet:ReceiptRuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiptRuleSet gets an existing ReceiptRuleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiptRuleSetState, opts ...pulumi.ResourceOption) (*ReceiptRuleSet, error) {
	var resource ReceiptRuleSet
	err := ctx.ReadResource("aws:ses/receiptRuleSet:ReceiptRuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReceiptRuleSet resources.
type receiptRuleSetState struct {
	// The name of the rule set
	RuleSetName *string `pulumi:"ruleSetName"`
}

type ReceiptRuleSetState struct {
	// The name of the rule set
	RuleSetName pulumi.StringPtrInput
}

func (ReceiptRuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleSetState)(nil)).Elem()
}

type receiptRuleSetArgs struct {
	// The name of the rule set
	RuleSetName string `pulumi:"ruleSetName"`
}

// The set of arguments for constructing a ReceiptRuleSet resource.
type ReceiptRuleSetArgs struct {
	// The name of the rule set
	RuleSetName pulumi.StringInput
}

func (ReceiptRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleSetArgs)(nil)).Elem()
}


// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Load Balancer Listener Rule resource.
//
// > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.
//
//
//
// Deprecated: aws.elasticloadbalancingv2.ListenerRule has been deprecated in favour of aws.lb.ListenerRule
//
// aws.elasticloadbalancingv2.ListenerRule has been deprecated in favour of aws.lb.ListenerRule
type ListenerRule struct {
	pulumi.CustomResourceState

	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayOutput `pulumi:"actions"`
	// The ARN of the rule (matches `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayOutput `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
}

// NewListenerRule registers a new resource with the given unique name, arguments, and options.
func NewListenerRule(ctx *pulumi.Context,
	name string, args *ListenerRuleArgs, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	if args == nil || args.Actions == nil {
		return nil, errors.New("missing required argument 'Actions'")
	}
	if args == nil || args.Conditions == nil {
		return nil, errors.New("missing required argument 'Conditions'")
	}
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	if args == nil {
		args = &ListenerRuleArgs{}
	}
	var resource ListenerRule
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerRule gets an existing ListenerRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerRuleState, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	var resource ListenerRule
	err := ctx.ReadResource("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerRule resources.
type listenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions []ListenerRuleAction `pulumi:"actions"`
	// The ARN of the rule (matches `id`)
	Arn *string `pulumi:"arn"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []ListenerRuleCondition `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn *string `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *int `pulumi:"priority"`
}

type ListenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayInput
	// The ARN of the rule (matches `id`)
	Arn pulumi.StringPtrInput
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayInput
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringPtrInput
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntPtrInput
}

func (ListenerRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleState)(nil)).Elem()
}

type listenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions []ListenerRuleAction `pulumi:"actions"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []ListenerRuleCondition `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn string `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *int `pulumi:"priority"`
}

// The set of arguments for constructing a ListenerRule resource.
type ListenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayInput
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayInput
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringInput
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntPtrInput
}

func (ListenerRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleArgs)(nil)).Elem()
}

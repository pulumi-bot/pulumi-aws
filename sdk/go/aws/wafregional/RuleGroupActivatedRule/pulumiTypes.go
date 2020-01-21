// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RuleGroupActivatedRule

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/wafregional/RuleGroupActivatedRuleAction"
)

type RuleGroupActivatedRule struct {
	// Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
	Action wafregionalRuleGroupActivatedRuleAction.RuleGroupActivatedRuleAction `pulumi:"action"`
	// Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
	Priority int `pulumi:"priority"`
	// The ID of a [rule](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)
	RuleId string `pulumi:"ruleId"`
	// The rule type, either [`REGULAR`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html), [`RATE_BASED`](https://www.terraform.io/docs/providers/aws/r/wafregional_rate_based_rule.html), or `GROUP`. Defaults to `REGULAR`.
	Type *string `pulumi:"type"`
}

type RuleGroupActivatedRuleInput interface {
	pulumi.Input

	ToRuleGroupActivatedRuleOutput() RuleGroupActivatedRuleOutput
	ToRuleGroupActivatedRuleOutputWithContext(context.Context) RuleGroupActivatedRuleOutput
}

type RuleGroupActivatedRuleArgs struct {
	// Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
	Action wafregionalRuleGroupActivatedRuleAction.RuleGroupActivatedRuleActionInput `pulumi:"action"`
	// Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
	Priority pulumi.IntInput `pulumi:"priority"`
	// The ID of a [rule](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The rule type, either [`REGULAR`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html), [`RATE_BASED`](https://www.terraform.io/docs/providers/aws/r/wafregional_rate_based_rule.html), or `GROUP`. Defaults to `REGULAR`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (RuleGroupActivatedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroupActivatedRule)(nil)).Elem()
}

func (i RuleGroupActivatedRuleArgs) ToRuleGroupActivatedRuleOutput() RuleGroupActivatedRuleOutput {
	return i.ToRuleGroupActivatedRuleOutputWithContext(context.Background())
}

func (i RuleGroupActivatedRuleArgs) ToRuleGroupActivatedRuleOutputWithContext(ctx context.Context) RuleGroupActivatedRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupActivatedRuleOutput)
}

type RuleGroupActivatedRuleArrayInput interface {
	pulumi.Input

	ToRuleGroupActivatedRuleArrayOutput() RuleGroupActivatedRuleArrayOutput
	ToRuleGroupActivatedRuleArrayOutputWithContext(context.Context) RuleGroupActivatedRuleArrayOutput
}

type RuleGroupActivatedRuleArray []RuleGroupActivatedRuleInput

func (RuleGroupActivatedRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleGroupActivatedRule)(nil)).Elem()
}

func (i RuleGroupActivatedRuleArray) ToRuleGroupActivatedRuleArrayOutput() RuleGroupActivatedRuleArrayOutput {
	return i.ToRuleGroupActivatedRuleArrayOutputWithContext(context.Background())
}

func (i RuleGroupActivatedRuleArray) ToRuleGroupActivatedRuleArrayOutputWithContext(ctx context.Context) RuleGroupActivatedRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupActivatedRuleArrayOutput)
}

type RuleGroupActivatedRuleOutput struct { *pulumi.OutputState }

func (RuleGroupActivatedRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroupActivatedRule)(nil)).Elem()
}

func (o RuleGroupActivatedRuleOutput) ToRuleGroupActivatedRuleOutput() RuleGroupActivatedRuleOutput {
	return o
}

func (o RuleGroupActivatedRuleOutput) ToRuleGroupActivatedRuleOutputWithContext(ctx context.Context) RuleGroupActivatedRuleOutput {
	return o
}

// Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
func (o RuleGroupActivatedRuleOutput) Action() wafregionalRuleGroupActivatedRuleAction.RuleGroupActivatedRuleActionOutput {
	return o.ApplyT(func (v RuleGroupActivatedRule) wafregionalRuleGroupActivatedRuleAction.RuleGroupActivatedRuleAction { return v.Action }).(wafregionalRuleGroupActivatedRuleAction.RuleGroupActivatedRuleActionOutput)
}

// Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
func (o RuleGroupActivatedRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func (v RuleGroupActivatedRule) int { return v.Priority }).(pulumi.IntOutput)
}

// The ID of a [rule](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)
func (o RuleGroupActivatedRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func (v RuleGroupActivatedRule) string { return v.RuleId }).(pulumi.StringOutput)
}

// The rule type, either [`REGULAR`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html), [`RATE_BASED`](https://www.terraform.io/docs/providers/aws/r/wafregional_rate_based_rule.html), or `GROUP`. Defaults to `REGULAR`.
func (o RuleGroupActivatedRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RuleGroupActivatedRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RuleGroupActivatedRuleArrayOutput struct { *pulumi.OutputState}

func (RuleGroupActivatedRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleGroupActivatedRule)(nil)).Elem()
}

func (o RuleGroupActivatedRuleArrayOutput) ToRuleGroupActivatedRuleArrayOutput() RuleGroupActivatedRuleArrayOutput {
	return o
}

func (o RuleGroupActivatedRuleArrayOutput) ToRuleGroupActivatedRuleArrayOutputWithContext(ctx context.Context) RuleGroupActivatedRuleArrayOutput {
	return o
}

func (o RuleGroupActivatedRuleArrayOutput) Index(i pulumi.IntInput) RuleGroupActivatedRuleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) RuleGroupActivatedRule {
		return vs[0].([]RuleGroupActivatedRule)[vs[1].(int)]
	}).(RuleGroupActivatedRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleGroupActivatedRuleOutput{})
	pulumi.RegisterOutputType(RuleGroupActivatedRuleArrayOutput{})
}

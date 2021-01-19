// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an WAF Regional Rule Resource for use with Application Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ipset, err := wafregional.NewIpSet(ctx, "ipset", &wafregional.IpSetArgs{
// 			IpSetDescriptors: wafregional.IpSetIpSetDescriptorArray{
// 				&wafregional.IpSetIpSetDescriptorArgs{
// 					Type:  pulumi.String("IPV4"),
// 					Value: pulumi.String("192.0.7.0/24"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = wafregional.NewRule(ctx, "wafrule", &wafregional.RuleArgs{
// 			MetricName: pulumi.String("tfWAFRule"),
// 			Predicates: wafregional.RulePredicateArray{
// 				&wafregional.RulePredicateArgs{
// 					Type:    pulumi.String("IPMatch"),
// 					DataId:  ipset.ID(),
// 					Negated: pulumi.Bool(false),
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
// ## Nested Fields
//
// ### `predicate`
//
// See the [WAF Documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_Predicate.html) for more information.
//
// #### Arguments
//
// * `type` - (Required) The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`
// * `dataId` - (Required) The unique identifier of a predicate, such as the ID of a `ByteMatchSet` or `IPSet`.
// * `negated` - (Required) Whether to use the settings or the negated settings that you specified in the objects.
//
// ## Import
//
// WAF Regional Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:wafregional/rule:Rule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The ARN of the WAF Regional Rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RulePredicateArrayOutput `pulumi:"predicates"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	var resource Rule
	err := ctx.RegisterResource("aws:wafregional/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("aws:wafregional/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The ARN of the WAF Regional Rule.
	Arn *string `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RulePredicate `pulumi:"predicates"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type RuleState struct {
	// The ARN of the WAF Regional Rule.
	Arn pulumi.StringPtrInput
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringPtrInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RulePredicateArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RulePredicate `pulumi:"predicates"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RulePredicateArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

func (i *Rule) ToRulePtrOutput() RulePtrOutput {
	return i.ToRulePtrOutputWithContext(context.Background())
}

func (i *Rule) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulePtrOutput)
}

type RulePtrInput interface {
	pulumi.Input

	ToRulePtrOutput() RulePtrOutput
	ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput
}

type rulePtrType RuleArgs

func (*rulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil))
}

func (i *rulePtrType) ToRulePtrOutput() RulePtrOutput {
	return i.ToRulePtrOutputWithContext(context.Background())
}

func (i *rulePtrType) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulePtrOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//          RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Rule)(nil))
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//          RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Rule)(nil))
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct {
	*pulumi.OutputState
}

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func (o RuleOutput) ToRulePtrOutput() RulePtrOutput {
	return o.ToRulePtrOutputWithContext(context.Background())
}

func (o RuleOutput) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return o.ApplyT(func(v Rule) *Rule {
		return &v
	}).(RulePtrOutput)
}

type RulePtrOutput struct {
	*pulumi.OutputState
}

func (RulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil))
}

func (o RulePtrOutput) ToRulePtrOutput() RulePtrOutput {
	return o
}

func (o RulePtrOutput) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return o
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Rule)(nil))
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Rule {
		return vs[0].([]Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Rule)(nil))
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Rule {
		return vs[0].(map[string]Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RulePtrOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}

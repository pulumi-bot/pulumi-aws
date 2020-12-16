// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Rate Based Rule Resource
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
// 		_, err = wafregional.NewRateBasedRule(ctx, "wafrule", &wafregional.RateBasedRuleArgs{
// 			MetricName: pulumi.String("tfWAFRule"),
// 			RateKey:    pulumi.String("IP"),
// 			RateLimit:  pulumi.Int(100),
// 			Predicates: wafregional.RateBasedRulePredicateArray{
// 				&wafregional.RateBasedRulePredicateArgs{
// 					DataId:  ipset.ID(),
// 					Negated: pulumi.Bool(false),
// 					Type:    pulumi.String("IPMatch"),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			ipset,
// 		}))
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
// WAF Regional Rate Based Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:wafregional/rateBasedRule:RateBasedRule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type RateBasedRule struct {
	pulumi.CustomResourceState

	// The ARN of the WAF Regional Rate Based Rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayOutput `pulumi:"predicates"`
	// Valid value is IP.
	RateKey pulumi.StringOutput `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntOutput `pulumi:"rateLimit"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRateBasedRule registers a new resource with the given unique name, arguments, and options.
func NewRateBasedRule(ctx *pulumi.Context,
	name string, args *RateBasedRuleArgs, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.RateKey == nil {
		return nil, errors.New("invalid value for required argument 'RateKey'")
	}
	if args.RateLimit == nil {
		return nil, errors.New("invalid value for required argument 'RateLimit'")
	}
	var resource RateBasedRule
	err := ctx.RegisterResource("aws:wafregional/rateBasedRule:RateBasedRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRateBasedRule gets an existing RateBasedRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRateBasedRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RateBasedRuleState, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	var resource RateBasedRule
	err := ctx.ReadResource("aws:wafregional/rateBasedRule:RateBasedRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RateBasedRule resources.
type rateBasedRuleState struct {
	// The ARN of the WAF Regional Rate Based Rule.
	Arn *string `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RateBasedRulePredicate `pulumi:"predicates"`
	// Valid value is IP.
	RateKey *string `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit *int `pulumi:"rateLimit"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type RateBasedRuleState struct {
	// The ARN of the WAF Regional Rate Based Rule.
	Arn pulumi.StringPtrInput
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringPtrInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayInput
	// Valid value is IP.
	RateKey pulumi.StringPtrInput
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RateBasedRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleState)(nil)).Elem()
}

type rateBasedRuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RateBasedRulePredicate `pulumi:"predicates"`
	// Valid value is IP.
	RateKey string `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit int `pulumi:"rateLimit"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RateBasedRule resource.
type RateBasedRuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayInput
	// Valid value is IP.
	RateKey pulumi.StringInput
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RateBasedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleArgs)(nil)).Elem()
}

type RateBasedRuleInput interface {
	pulumi.Input

	ToRateBasedRuleOutput() RateBasedRuleOutput
	ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput
}

func (*RateBasedRule) ElementType() reflect.Type {
	return reflect.TypeOf((*RateBasedRule)(nil))
}

func (i *RateBasedRule) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return i.ToRateBasedRuleOutputWithContext(context.Background())
}

func (i *RateBasedRule) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRuleOutput)
}

type RateBasedRuleOutput struct {
	*pulumi.OutputState
}

func (RateBasedRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateBasedRule)(nil))
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return o
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RateBasedRuleOutput{})
}

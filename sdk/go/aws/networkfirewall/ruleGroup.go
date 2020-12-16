// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Network Firewall Rule Group Resource
//
// ## Example Usage
// ### Stateful Inspection
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity: pulumi.Int(100),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					RulesSourceList: &networkfirewall.RuleGroupRuleGroupRulesSourceRulesSourceListArgs{
// 						GeneratedRulesType: pulumi.String("DENYLIST"),
// 						TargetTypes: pulumi.StringArray{
// 							pulumi.String("HTTP_HOST"),
// 						},
// 						Targets: pulumi.StringArray{
// 							pulumi.String("test.example.com"),
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			Type: pulumi.String("STATEFUL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateful Inspection compatible with intrusion detection systems like Snort or Suricata
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity: pulumi.Int(100),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					StatefulRule: pulumi.MapArray{
// 						pulumi.Map{
// 							"action": pulumi.String("DROP"),
// 							"header": pulumi.Map{
// 								"destination":     pulumi.String("124.1.1.24/32"),
// 								"destinationPort": pulumi.Float64(53),
// 								"direction":       pulumi.String("ANY"),
// 								"protocol":        pulumi.String("TCP"),
// 								"source":          pulumi.String("1.2.3.4/32"),
// 								"sourcePort":      pulumi.Float64(53),
// 							},
// 							"ruleOption": pulumi.StringMapArray{
// 								pulumi.StringMap{
// 									"keyword": pulumi.String("sid:1"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			Type: pulumi.String("STATEFUL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateless Inspection with a Custom Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity:    pulumi.Int(100),
// 			Description: pulumi.String("Stateless Rate Limiting Rule"),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					StatelessRulesAndCustomActions: &networkfirewall.RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsArgs{
// 						CustomAction: pulumi.MapArray{
// 							pulumi.Map{
// 								"actionDefinition": pulumi.StringMapArrayMapMap{
// 									"publishMetricAction": pulumi.StringMapArrayMap{
// 										"dimension": pulumi.StringMapArray{
// 											pulumi.StringMap{
// 												"value": pulumi.String("2"),
// 											},
// 										},
// 									},
// 								},
// 								"actionName": pulumi.String("ExampleMetricsAction"),
// 							},
// 						},
// 						StatelessRule: pulumi.MapArray{
// 							pulumi.Map{
// 								"priority": pulumi.Float64(1),
// 								"ruleDefinition": pulumi.Map{
// 									"actions": pulumi.StringArray{
// 										pulumi.String("aws:pass"),
// 										pulumi.String("ExampleMetricsAction"),
// 									},
// 									"matchAttributes": pulumi.Map{
// 										"destination": pulumi.StringMapArray{
// 											pulumi.StringMap{
// 												"addressDefinition": pulumi.String("124.1.1.5/32"),
// 											},
// 										},
// 										"destinationPort": pulumi.Float64MapArray{
// 											pulumi.Float64Map{
// 												"fromPort": pulumi.Float64(443),
// 												"toPort":   pulumi.Float64(443),
// 											},
// 										},
// 										"protocols": pulumi.Float64Array{
// 											pulumi.Float64(6),
// 										},
// 										"source": pulumi.StringMapArray{
// 											pulumi.StringMap{
// 												"addressDefinition": pulumi.String("1.2.3.4/32"),
// 											},
// 										},
// 										"sourcePort": pulumi.Float64MapArray{
// 											pulumi.Float64Map{
// 												"fromPort": pulumi.Float64(443),
// 												"toPort":   pulumi.Float64(443),
// 											},
// 										},
// 										"tcpFlag": pulumi.StringArrayMapArray{
// 											pulumi.StringArrayMap{
// 												"flags": pulumi.StringArray{
// 													pulumi.String("SYN"),
// 												},
// 												"masks": pulumi.StringArray{
// 													pulumi.String("SYN"),
// 													pulumi.String("ACK"),
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			Type: pulumi.String("STATELESS"),
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
// Network Firewall Rule Groups can be imported using their `ARN`.
//
// ```sh
//  $ pulumi import aws:networkfirewall/ruleGroup:RuleGroup example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
// ```
type RuleGroup struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name of the rule group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup RuleGroupRuleGroupOutput `pulumi:"ruleGroup"`
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules pulumi.StringPtrOutput `pulumi:"rules"`
	// A map of key:value pairs to associate with the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type pulumi.StringOutput `pulumi:"type"`
	// A string token used when updating the rule group.
	UpdateToken pulumi.StringOutput `pulumi:"updateToken"`
}

// NewRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewRuleGroup(ctx *pulumi.Context,
	name string, args *RuleGroupArgs, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capacity == nil {
		return nil, errors.New("invalid value for required argument 'Capacity'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource RuleGroup
	err := ctx.RegisterResource("aws:networkfirewall/ruleGroup:RuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleGroup gets an existing RuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleGroupState, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	var resource RuleGroup
	err := ctx.ReadResource("aws:networkfirewall/ruleGroup:RuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleGroup resources.
type ruleGroupState struct {
	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn *string `pulumi:"arn"`
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity *int `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description *string `pulumi:"description"`
	// A friendly name of the rule group.
	Name *string `pulumi:"name"`
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup *RuleGroupRuleGroup `pulumi:"ruleGroup"`
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules *string `pulumi:"rules"`
	// A map of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type *string `pulumi:"type"`
	// A string token used when updating the rule group.
	UpdateToken *string `pulumi:"updateToken"`
}

type RuleGroupState struct {
	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn pulumi.StringPtrInput
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity pulumi.IntPtrInput
	// A friendly description of the rule group.
	Description pulumi.StringPtrInput
	// A friendly name of the rule group.
	Name pulumi.StringPtrInput
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup RuleGroupRuleGroupPtrInput
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules pulumi.StringPtrInput
	// A map of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type pulumi.StringPtrInput
	// A string token used when updating the rule group.
	UpdateToken pulumi.StringPtrInput
}

func (RuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupState)(nil)).Elem()
}

type ruleGroupArgs struct {
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity int `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description *string `pulumi:"description"`
	// A friendly name of the rule group.
	Name *string `pulumi:"name"`
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup *RuleGroupRuleGroup `pulumi:"ruleGroup"`
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules *string `pulumi:"rules"`
	// A map of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a RuleGroup resource.
type RuleGroupArgs struct {
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity pulumi.IntInput
	// A friendly description of the rule group.
	Description pulumi.StringPtrInput
	// A friendly name of the rule group.
	Name pulumi.StringPtrInput
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup RuleGroupRuleGroupPtrInput
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules pulumi.StringPtrInput
	// A map of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type pulumi.StringInput
}

func (RuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupArgs)(nil)).Elem()
}

type RuleGroupInput interface {
	pulumi.Input

	ToRuleGroupOutput() RuleGroupOutput
	ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput
}

func (*RuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroup)(nil))
}

func (i *RuleGroup) ToRuleGroupOutput() RuleGroupOutput {
	return i.ToRuleGroupOutputWithContext(context.Background())
}

func (i *RuleGroup) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupOutput)
}

type RuleGroupOutput struct {
	*pulumi.OutputState
}

func (RuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroup)(nil))
}

func (o RuleGroupOutput) ToRuleGroupOutput() RuleGroupOutput {
	return o
}

func (o RuleGroupOutput) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RuleGroupOutput{})
}

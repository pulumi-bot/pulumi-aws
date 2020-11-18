// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a WAFv2 Web ACL resource.
//
// ## Example Usage
//
// This resource is based on `wafv2.RuleGroup`, check the documentation of the `wafv2.RuleGroup` resource to see examples of the various available statements.
// ### Managed Rule
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafv2.NewWebAcl(ctx, "example", &wafv2.WebAclArgs{
// 			DefaultAction: &wafv2.WebAclDefaultActionArgs{
// 				Allow: nil,
// 			},
// 			Description: pulumi.String("Example of a managed rule."),
// 			Rules: wafv2.WebAclRuleArray{
// 				&wafv2.WebAclRuleArgs{
// 					Name: pulumi.String("rule-1"),
// 					OverrideAction: &wafv2.WebAclRuleOverrideActionArgs{
// 						Count: nil,
// 					},
// 					Priority: pulumi.Int(1),
// 					Statement: &wafv2.WebAclRuleStatementArgs{
// 						ManagedRuleGroupStatement: &wafv2.WebAclRuleStatementManagedRuleGroupStatementArgs{
// 							ExcludedRule: pulumi.StringMapArray{
// 								pulumi.StringMap{
// 									"name": pulumi.String("SizeRestrictions_QUERYSTRING"),
// 								},
// 								pulumi.StringMap{
// 									"name": pulumi.String("NoUserAgent_HEADER"),
// 								},
// 							},
// 							Name:       pulumi.String("AWSManagedRulesCommonRuleSet"),
// 							VendorName: pulumi.String("AWS"),
// 						},
// 					},
// 					VisibilityConfig: &wafv2.WebAclRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 			},
// 			Scope: pulumi.String("REGIONAL"),
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			VisibilityConfig: &wafv2.WebAclVisibilityConfigArgs{
// 				CloudwatchMetricsEnabled: pulumi.Bool(false),
// 				MetricName:               pulumi.String("friendly-metric-name"),
// 				SampledRequestsEnabled:   pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Rate Based
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafv2.NewWebAcl(ctx, "example", &wafv2.WebAclArgs{
// 			DefaultAction: &wafv2.WebAclDefaultActionArgs{
// 				Block: nil,
// 			},
// 			Description: pulumi.String("Example of a rate based statement."),
// 			Rules: wafv2.WebAclRuleArray{
// 				&wafv2.WebAclRuleArgs{
// 					Action: &wafv2.WebAclRuleActionArgs{
// 						Count: nil,
// 					},
// 					Name:     pulumi.String("rule-1"),
// 					Priority: pulumi.Int(1),
// 					Statement: &wafv2.WebAclRuleStatementArgs{
// 						RateBasedStatement: &wafv2.WebAclRuleStatementRateBasedStatementArgs{
// 							AggregateKeyType: pulumi.String("IP"),
// 							Limit:            pulumi.Int(10000),
// 							ScopeDownStatement: &wafv2.WebAclRuleStatementRateBasedStatementScopeDownStatementArgs{
// 								GeoMatchStatement: &wafv2.WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementArgs{
// 									CountryCodes: pulumi.StringArray{
// 										pulumi.String("US"),
// 										pulumi.String("NL"),
// 									},
// 								},
// 							},
// 						},
// 					},
// 					VisibilityConfig: &wafv2.WebAclRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 			},
// 			Scope: pulumi.String("REGIONAL"),
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			VisibilityConfig: &wafv2.WebAclVisibilityConfigArgs{
// 				CloudwatchMetricsEnabled: pulumi.Bool(false),
// 				MetricName:               pulumi.String("friendly-metric-name"),
// 				SampledRequestsEnabled:   pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Rule Group Reference
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := wafv2.NewRuleGroup(ctx, "example", &wafv2.RuleGroupArgs{
// 			Capacity: pulumi.Int(10),
// 			Scope:    pulumi.String("REGIONAL"),
// 			Rules: wafv2.RuleGroupRuleArray{
// 				&wafv2.RuleGroupRuleArgs{
// 					Name:     pulumi.String("rule-1"),
// 					Priority: pulumi.Int(1),
// 					Action: &wafv2.RuleGroupRuleActionArgs{
// 						Count: nil,
// 					},
// 					Statement: &wafv2.RuleGroupRuleStatementArgs{
// 						GeoMatchStatement: &wafv2.RuleGroupRuleStatementGeoMatchStatementArgs{
// 							CountryCodes: pulumi.StringArray{
// 								pulumi.String("NL"),
// 							},
// 						},
// 					},
// 					VisibilityConfig: &wafv2.RuleGroupRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 				&wafv2.RuleGroupRuleArgs{
// 					Name:     pulumi.String("rule-to-exclude-a"),
// 					Priority: pulumi.Int(10),
// 					Action: &wafv2.RuleGroupRuleActionArgs{
// 						Allow: nil,
// 					},
// 					Statement: &wafv2.RuleGroupRuleStatementArgs{
// 						GeoMatchStatement: &wafv2.RuleGroupRuleStatementGeoMatchStatementArgs{
// 							CountryCodes: pulumi.StringArray{
// 								pulumi.String("US"),
// 							},
// 						},
// 					},
// 					VisibilityConfig: &wafv2.RuleGroupRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 				&wafv2.RuleGroupRuleArgs{
// 					Name:     pulumi.String("rule-to-exclude-b"),
// 					Priority: pulumi.Int(15),
// 					Action: &wafv2.RuleGroupRuleActionArgs{
// 						Allow: nil,
// 					},
// 					Statement: &wafv2.RuleGroupRuleStatementArgs{
// 						GeoMatchStatement: &wafv2.RuleGroupRuleStatementGeoMatchStatementArgs{
// 							CountryCodes: pulumi.StringArray{
// 								pulumi.String("GB"),
// 							},
// 						},
// 					},
// 					VisibilityConfig: &wafv2.RuleGroupRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 			},
// 			VisibilityConfig: &wafv2.RuleGroupVisibilityConfigArgs{
// 				CloudwatchMetricsEnabled: pulumi.Bool(false),
// 				MetricName:               pulumi.String("friendly-metric-name"),
// 				SampledRequestsEnabled:   pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = wafv2.NewWebAcl(ctx, "test", &wafv2.WebAclArgs{
// 			Scope: pulumi.String("REGIONAL"),
// 			DefaultAction: &wafv2.WebAclDefaultActionArgs{
// 				Block: nil,
// 			},
// 			Rules: wafv2.WebAclRuleArray{
// 				&wafv2.WebAclRuleArgs{
// 					Name:     pulumi.String("rule-1"),
// 					Priority: pulumi.Int(1),
// 					OverrideAction: &wafv2.WebAclRuleOverrideActionArgs{
// 						Count: nil,
// 					},
// 					Statement: &wafv2.WebAclRuleStatementArgs{
// 						RuleGroupReferenceStatement: &wafv2.WebAclRuleStatementRuleGroupReferenceStatementArgs{
// 							Arn: example.Arn,
// 							ExcludedRules: wafv2.WebAclRuleStatementRuleGroupReferenceStatementExcludedRuleArray{
// 								&wafv2.WebAclRuleStatementRuleGroupReferenceStatementExcludedRuleArgs{
// 									Name: pulumi.String("rule-to-exclude-b"),
// 								},
// 								&wafv2.WebAclRuleStatementRuleGroupReferenceStatementExcludedRuleArgs{
// 									Name: pulumi.String("rule-to-exclude-a"),
// 								},
// 							},
// 						},
// 					},
// 					VisibilityConfig: &wafv2.WebAclRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			VisibilityConfig: &wafv2.WebAclVisibilityConfigArgs{
// 				CloudwatchMetricsEnabled: pulumi.Bool(false),
// 				MetricName:               pulumi.String("friendly-metric-name"),
// 				SampledRequestsEnabled:   pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type WebAcl struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The web ACL capacity units (WCUs) currently being used by this web ACL.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
	DefaultAction WebAclDefaultActionOutput `pulumi:"defaultAction"`
	// A friendly description of the WebACL.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	LockToken   pulumi.StringOutput    `pulumi:"lockToken"`
	// A friendly name of the WebACL.
	Name pulumi.StringOutput `pulumi:"name"`
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules WebAclRuleArrayOutput `pulumi:"rules"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig WebAclVisibilityConfigOutput `pulumi:"visibilityConfig"`
}

// NewWebAcl registers a new resource with the given unique name, arguments, and options.
func NewWebAcl(ctx *pulumi.Context,
	name string, args *WebAclArgs, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.VisibilityConfig == nil {
		return nil, errors.New("invalid value for required argument 'VisibilityConfig'")
	}
	var resource WebAcl
	err := ctx.RegisterResource("aws:wafv2/webAcl:WebAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAcl gets an existing WebAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclState, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	var resource WebAcl
	err := ctx.ReadResource("aws:wafv2/webAcl:WebAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type webAclState struct {
	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn *string `pulumi:"arn"`
	// The web ACL capacity units (WCUs) currently being used by this web ACL.
	Capacity *int `pulumi:"capacity"`
	// The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
	DefaultAction *WebAclDefaultAction `pulumi:"defaultAction"`
	// A friendly description of the WebACL.
	Description *string `pulumi:"description"`
	LockToken   *string `pulumi:"lockToken"`
	// A friendly name of the WebACL.
	Name *string `pulumi:"name"`
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules []WebAclRule `pulumi:"rules"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope *string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig *WebAclVisibilityConfig `pulumi:"visibilityConfig"`
}

type WebAclState struct {
	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn pulumi.StringPtrInput
	// The web ACL capacity units (WCUs) currently being used by this web ACL.
	Capacity pulumi.IntPtrInput
	// The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
	DefaultAction WebAclDefaultActionPtrInput
	// A friendly description of the WebACL.
	Description pulumi.StringPtrInput
	LockToken   pulumi.StringPtrInput
	// A friendly name of the WebACL.
	Name pulumi.StringPtrInput
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules WebAclRuleArrayInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringPtrInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig WebAclVisibilityConfigPtrInput
}

func (WebAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclState)(nil)).Elem()
}

type webAclArgs struct {
	// The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
	DefaultAction WebAclDefaultAction `pulumi:"defaultAction"`
	// A friendly description of the WebACL.
	Description *string `pulumi:"description"`
	// A friendly name of the WebACL.
	Name *string `pulumi:"name"`
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules []WebAclRule `pulumi:"rules"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig WebAclVisibilityConfig `pulumi:"visibilityConfig"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
	DefaultAction WebAclDefaultActionInput
	// A friendly description of the WebACL.
	Description pulumi.StringPtrInput
	// A friendly name of the WebACL.
	Name pulumi.StringPtrInput
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules WebAclRuleArrayInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig WebAclVisibilityConfigInput
}

func (WebAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclArgs)(nil)).Elem()
}

type WebAclInput interface {
	pulumi.Input

	ToWebAclOutput() WebAclOutput
	ToWebAclOutputWithContext(ctx context.Context) WebAclOutput
}

func (WebAcl) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAcl)(nil)).Elem()
}

func (i WebAcl) ToWebAclOutput() WebAclOutput {
	return i.ToWebAclOutputWithContext(context.Background())
}

func (i WebAcl) ToWebAclOutputWithContext(ctx context.Context) WebAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclOutput)
}

type WebAclOutput struct {
	*pulumi.OutputState
}

func (WebAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAclOutput)(nil)).Elem()
}

func (o WebAclOutput) ToWebAclOutput() WebAclOutput {
	return o
}

func (o WebAclOutput) ToWebAclOutputWithContext(ctx context.Context) WebAclOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAclOutput{})
}

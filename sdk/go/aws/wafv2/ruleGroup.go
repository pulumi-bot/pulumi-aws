// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a WAFv2 Rule Group resource.
//
// ## Example Usage
// ### Simple
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
// 		_, err := wafv2.NewRuleGroup(ctx, "example", &wafv2.RuleGroupArgs{
// 			Capacity: pulumi.Int(2),
// 			Rules: wafv2.RuleGroupRuleArray{
// 				&wafv2.RuleGroupRuleArgs{
// 					Action: &wafv2.RuleGroupRuleActionArgs{
// 						Allow: nil,
// 					},
// 					Name:     pulumi.String("rule-1"),
// 					Priority: pulumi.Int(1),
// 					Statement: &wafv2.RuleGroupRuleStatementArgs{
// 						GeoMatchStatement: &wafv2.RuleGroupRuleStatementGeoMatchStatementArgs{
// 							CountryCodes: pulumi.StringArray{
// 								pulumi.String("US"),
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
// 			},
// 			Scope: pulumi.String("REGIONAL"),
// 			VisibilityConfig: &wafv2.RuleGroupVisibilityConfigArgs{
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
//
// ## Import
//
// WAFv2 Rule Group can be imported using `ID/name/scope` e.g.
//
// ```sh
//  $ pulumi import aws:wafv2/ruleGroup:RuleGroup example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
// ```
type RuleGroup struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	LockToken   pulumi.StringOutput    `pulumi:"lockToken"`
	// A friendly name of the rule group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules RuleGroupRuleArrayOutput `pulumi:"rules"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig RuleGroupVisibilityConfigOutput `pulumi:"visibilityConfig"`
}

// NewRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewRuleGroup(ctx *pulumi.Context,
	name string, args *RuleGroupArgs, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	if args == nil || args.Capacity == nil {
		return nil, errors.New("missing required argument 'Capacity'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.VisibilityConfig == nil {
		return nil, errors.New("missing required argument 'VisibilityConfig'")
	}
	if args == nil {
		args = &RuleGroupArgs{}
	}
	var resource RuleGroup
	err := ctx.RegisterResource("aws:wafv2/ruleGroup:RuleGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:wafv2/ruleGroup:RuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleGroup resources.
type ruleGroupState struct {
	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn *string `pulumi:"arn"`
	// The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
	Capacity *int `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description *string `pulumi:"description"`
	LockToken   *string `pulumi:"lockToken"`
	// A friendly name of the rule group.
	Name *string `pulumi:"name"`
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules []RuleGroupRule `pulumi:"rules"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope *string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig *RuleGroupVisibilityConfig `pulumi:"visibilityConfig"`
}

type RuleGroupState struct {
	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn pulumi.StringPtrInput
	// The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
	Capacity pulumi.IntPtrInput
	// A friendly description of the rule group.
	Description pulumi.StringPtrInput
	LockToken   pulumi.StringPtrInput
	// A friendly name of the rule group.
	Name pulumi.StringPtrInput
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules RuleGroupRuleArrayInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringPtrInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig RuleGroupVisibilityConfigPtrInput
}

func (RuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupState)(nil)).Elem()
}

type ruleGroupArgs struct {
	// The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
	Capacity int `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description *string `pulumi:"description"`
	// A friendly name of the rule group.
	Name *string `pulumi:"name"`
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules []RuleGroupRule `pulumi:"rules"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig RuleGroupVisibilityConfig `pulumi:"visibilityConfig"`
}

// The set of arguments for constructing a RuleGroup resource.
type RuleGroupArgs struct {
	// The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
	Capacity pulumi.IntInput
	// A friendly description of the rule group.
	Description pulumi.StringPtrInput
	// A friendly name of the rule group.
	Name pulumi.StringPtrInput
	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules RuleGroupRuleArrayInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig RuleGroupVisibilityConfigInput
}

func (RuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupArgs)(nil)).Elem()
}

type RuleGroupInput interface {
	pulumi.Input

	ToRuleGroupOutput() RuleGroupOutput
	ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput
}

func (RuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroup)(nil)).Elem()
}

func (i RuleGroup) ToRuleGroupOutput() RuleGroupOutput {
	return i.ToRuleGroupOutputWithContext(context.Background())
}

func (i RuleGroup) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupOutput)
}

type RuleGroupOutput struct {
	*pulumi.OutputState
}

func (RuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroupOutput)(nil)).Elem()
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

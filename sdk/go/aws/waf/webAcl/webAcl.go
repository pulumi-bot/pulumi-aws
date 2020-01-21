// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package webAcl

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclDefaultAction"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclLoggingConfiguration"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclLoggingConfigurationRedactedFields"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclLoggingConfigurationRedactedFieldsFieldToMatch"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclRule"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclRuleAction"
	"https:/github.com/pulumi/pulumi-aws/waf/WebAclRuleOverrideAction"
)

// Provides a WAF Web ACL Resource
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_web_acl.html.markdown.
type WebAcl struct {
	pulumi.CustomResourceState

	// The ARN of the WAF WebACL.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction wafWebAclDefaultAction.WebAclDefaultActionOutput `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration wafWebAclLoggingConfiguration.WebAclLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules wafWebAclRule.WebAclRuleArrayOutput `pulumi:"rules"`
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewWebAcl registers a new resource with the given unique name, arguments, and options.
func NewWebAcl(ctx *pulumi.Context,
	name string, args *WebAclArgs, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	if args == nil || args.DefaultAction == nil {
		return nil, errors.New("missing required argument 'DefaultAction'")
	}
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	if args == nil {
		args = &WebAclArgs{}
	}
	var resource WebAcl
	err := ctx.RegisterResource("aws:waf/webAcl:WebAcl", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:waf/webAcl:WebAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type webAclState struct {
	// The ARN of the WAF WebACL.
	Arn *string `pulumi:"arn"`
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction *wafWebAclDefaultAction.WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *wafWebAclLoggingConfiguration.WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules []wafWebAclRule.WebAclRule `pulumi:"rules"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

type WebAclState struct {
	// The ARN of the WAF WebACL.
	Arn pulumi.StringPtrInput
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction wafWebAclDefaultAction.WebAclDefaultActionPtrInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration wafWebAclLoggingConfiguration.WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringPtrInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules wafWebAclRule.WebAclRuleArrayInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (WebAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclState)(nil)).Elem()
}

type webAclArgs struct {
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction wafWebAclDefaultAction.WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *wafWebAclLoggingConfiguration.WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules []wafWebAclRule.WebAclRule `pulumi:"rules"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction wafWebAclDefaultAction.WebAclDefaultActionInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration wafWebAclLoggingConfiguration.WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules wafWebAclRule.WebAclRuleArrayInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (WebAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclArgs)(nil)).Elem()
}


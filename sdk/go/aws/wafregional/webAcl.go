// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
//
// ## Example Usage
type WebAcl struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionOutput `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRuleArrayOutput `pulumi:"rules"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
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
	err := ctx.RegisterResource("aws:wafregional/webAcl:WebAcl", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:wafregional/webAcl:WebAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type webAclState struct {
	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn *string `pulumi:"arn"`
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction *WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules []WebAclRule `pulumi:"rules"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type WebAclState struct {
	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn pulumi.StringPtrInput
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionPtrInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringPtrInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRuleArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (WebAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclState)(nil)).Elem()
}

type webAclArgs struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules []WebAclRule `pulumi:"rules"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRuleArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (WebAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclArgs)(nil)).Elem()
}

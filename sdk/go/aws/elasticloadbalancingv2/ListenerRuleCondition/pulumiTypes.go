// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ListenerRuleCondition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/elasticloadbalancingv2/ListenerRuleConditionHostHeader"
	"https:/github.com/pulumi/pulumi-aws/elasticloadbalancingv2/ListenerRuleConditionHttpHeader"
	"https:/github.com/pulumi/pulumi-aws/elasticloadbalancingv2/ListenerRuleConditionHttpRequestMethod"
	"https:/github.com/pulumi/pulumi-aws/elasticloadbalancingv2/ListenerRuleConditionPathPattern"
	"https:/github.com/pulumi/pulumi-aws/elasticloadbalancingv2/ListenerRuleConditionQueryString"
	"https:/github.com/pulumi/pulumi-aws/elasticloadbalancingv2/ListenerRuleConditionSourceIp"
)

type ListenerRuleCondition struct {
	// The type of condition. Valid values are `host-header` or `path-pattern`. Must also set `values`.
	Field *string `pulumi:"field"`
	// Contains a single `value` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
	HostHeader *elasticloadbalancingv2ListenerRuleConditionHostHeader.ListenerRuleConditionHostHeader `pulumi:"hostHeader"`
	// HTTP headers to match. HTTP Header block fields documented below.
	HttpHeader *elasticloadbalancingv2ListenerRuleConditionHttpHeader.ListenerRuleConditionHttpHeader `pulumi:"httpHeader"`
	// Contains a single `value` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
	HttpRequestMethod *elasticloadbalancingv2ListenerRuleConditionHttpRequestMethod.ListenerRuleConditionHttpRequestMethod `pulumi:"httpRequestMethod"`
	// Contains a single `value` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard charaters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query-string` condition.
	PathPattern *elasticloadbalancingv2ListenerRuleConditionPathPattern.ListenerRuleConditionPathPattern `pulumi:"pathPattern"`
	// Query strings to match. Query String block fields documented below.
	QueryStrings []elasticloadbalancingv2ListenerRuleConditionQueryString.ListenerRuleConditionQueryString `pulumi:"queryStrings"`
	// Contains a single `value` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http-header` condition instead.
	SourceIp *elasticloadbalancingv2ListenerRuleConditionSourceIp.ListenerRuleConditionSourceIp `pulumi:"sourceIp"`
	// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
	Values *string `pulumi:"values"`
}

type ListenerRuleConditionInput interface {
	pulumi.Input

	ToListenerRuleConditionOutput() ListenerRuleConditionOutput
	ToListenerRuleConditionOutputWithContext(context.Context) ListenerRuleConditionOutput
}

type ListenerRuleConditionArgs struct {
	// The type of condition. Valid values are `host-header` or `path-pattern`. Must also set `values`.
	Field pulumi.StringPtrInput `pulumi:"field"`
	// Contains a single `value` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
	HostHeader elasticloadbalancingv2ListenerRuleConditionHostHeader.ListenerRuleConditionHostHeaderPtrInput `pulumi:"hostHeader"`
	// HTTP headers to match. HTTP Header block fields documented below.
	HttpHeader elasticloadbalancingv2ListenerRuleConditionHttpHeader.ListenerRuleConditionHttpHeaderPtrInput `pulumi:"httpHeader"`
	// Contains a single `value` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
	HttpRequestMethod elasticloadbalancingv2ListenerRuleConditionHttpRequestMethod.ListenerRuleConditionHttpRequestMethodPtrInput `pulumi:"httpRequestMethod"`
	// Contains a single `value` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard charaters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query-string` condition.
	PathPattern elasticloadbalancingv2ListenerRuleConditionPathPattern.ListenerRuleConditionPathPatternPtrInput `pulumi:"pathPattern"`
	// Query strings to match. Query String block fields documented below.
	QueryStrings elasticloadbalancingv2ListenerRuleConditionQueryString.ListenerRuleConditionQueryStringArrayInput `pulumi:"queryStrings"`
	// Contains a single `value` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http-header` condition instead.
	SourceIp elasticloadbalancingv2ListenerRuleConditionSourceIp.ListenerRuleConditionSourceIpPtrInput `pulumi:"sourceIp"`
	// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
	Values pulumi.StringPtrInput `pulumi:"values"`
}

func (ListenerRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleCondition)(nil)).Elem()
}

func (i ListenerRuleConditionArgs) ToListenerRuleConditionOutput() ListenerRuleConditionOutput {
	return i.ToListenerRuleConditionOutputWithContext(context.Background())
}

func (i ListenerRuleConditionArgs) ToListenerRuleConditionOutputWithContext(ctx context.Context) ListenerRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleConditionOutput)
}

type ListenerRuleConditionArrayInput interface {
	pulumi.Input

	ToListenerRuleConditionArrayOutput() ListenerRuleConditionArrayOutput
	ToListenerRuleConditionArrayOutputWithContext(context.Context) ListenerRuleConditionArrayOutput
}

type ListenerRuleConditionArray []ListenerRuleConditionInput

func (ListenerRuleConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerRuleCondition)(nil)).Elem()
}

func (i ListenerRuleConditionArray) ToListenerRuleConditionArrayOutput() ListenerRuleConditionArrayOutput {
	return i.ToListenerRuleConditionArrayOutputWithContext(context.Background())
}

func (i ListenerRuleConditionArray) ToListenerRuleConditionArrayOutputWithContext(ctx context.Context) ListenerRuleConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleConditionArrayOutput)
}

type ListenerRuleConditionOutput struct { *pulumi.OutputState }

func (ListenerRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleCondition)(nil)).Elem()
}

func (o ListenerRuleConditionOutput) ToListenerRuleConditionOutput() ListenerRuleConditionOutput {
	return o
}

func (o ListenerRuleConditionOutput) ToListenerRuleConditionOutputWithContext(ctx context.Context) ListenerRuleConditionOutput {
	return o
}

// The type of condition. Valid values are `host-header` or `path-pattern`. Must also set `values`.
func (o ListenerRuleConditionOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *string { return v.Field }).(pulumi.StringPtrOutput)
}

// Contains a single `value` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
func (o ListenerRuleConditionOutput) HostHeader() elasticloadbalancingv2ListenerRuleConditionHostHeader.ListenerRuleConditionHostHeaderPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *elasticloadbalancingv2ListenerRuleConditionHostHeader.ListenerRuleConditionHostHeader { return v.HostHeader }).(elasticloadbalancingv2ListenerRuleConditionHostHeader.ListenerRuleConditionHostHeaderPtrOutput)
}

// HTTP headers to match. HTTP Header block fields documented below.
func (o ListenerRuleConditionOutput) HttpHeader() elasticloadbalancingv2ListenerRuleConditionHttpHeader.ListenerRuleConditionHttpHeaderPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *elasticloadbalancingv2ListenerRuleConditionHttpHeader.ListenerRuleConditionHttpHeader { return v.HttpHeader }).(elasticloadbalancingv2ListenerRuleConditionHttpHeader.ListenerRuleConditionHttpHeaderPtrOutput)
}

// Contains a single `value` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
func (o ListenerRuleConditionOutput) HttpRequestMethod() elasticloadbalancingv2ListenerRuleConditionHttpRequestMethod.ListenerRuleConditionHttpRequestMethodPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *elasticloadbalancingv2ListenerRuleConditionHttpRequestMethod.ListenerRuleConditionHttpRequestMethod { return v.HttpRequestMethod }).(elasticloadbalancingv2ListenerRuleConditionHttpRequestMethod.ListenerRuleConditionHttpRequestMethodPtrOutput)
}

// Contains a single `value` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard charaters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query-string` condition.
func (o ListenerRuleConditionOutput) PathPattern() elasticloadbalancingv2ListenerRuleConditionPathPattern.ListenerRuleConditionPathPatternPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *elasticloadbalancingv2ListenerRuleConditionPathPattern.ListenerRuleConditionPathPattern { return v.PathPattern }).(elasticloadbalancingv2ListenerRuleConditionPathPattern.ListenerRuleConditionPathPatternPtrOutput)
}

// Query strings to match. Query String block fields documented below.
func (o ListenerRuleConditionOutput) QueryStrings() elasticloadbalancingv2ListenerRuleConditionQueryString.ListenerRuleConditionQueryStringArrayOutput {
	return o.ApplyT(func (v ListenerRuleCondition) []elasticloadbalancingv2ListenerRuleConditionQueryString.ListenerRuleConditionQueryString { return v.QueryStrings }).(elasticloadbalancingv2ListenerRuleConditionQueryString.ListenerRuleConditionQueryStringArrayOutput)
}

// Contains a single `value` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http-header` condition instead.
func (o ListenerRuleConditionOutput) SourceIp() elasticloadbalancingv2ListenerRuleConditionSourceIp.ListenerRuleConditionSourceIpPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *elasticloadbalancingv2ListenerRuleConditionSourceIp.ListenerRuleConditionSourceIp { return v.SourceIp }).(elasticloadbalancingv2ListenerRuleConditionSourceIp.ListenerRuleConditionSourceIpPtrOutput)
}

// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
func (o ListenerRuleConditionOutput) Values() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleCondition) *string { return v.Values }).(pulumi.StringPtrOutput)
}

type ListenerRuleConditionArrayOutput struct { *pulumi.OutputState}

func (ListenerRuleConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerRuleCondition)(nil)).Elem()
}

func (o ListenerRuleConditionArrayOutput) ToListenerRuleConditionArrayOutput() ListenerRuleConditionArrayOutput {
	return o
}

func (o ListenerRuleConditionArrayOutput) ToListenerRuleConditionArrayOutputWithContext(ctx context.Context) ListenerRuleConditionArrayOutput {
	return o
}

func (o ListenerRuleConditionArrayOutput) Index(i pulumi.IntInput) ListenerRuleConditionOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ListenerRuleCondition {
		return vs[0].([]ListenerRuleCondition)[vs[1].(int)]
	}).(ListenerRuleConditionOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerRuleConditionOutput{})
	pulumi.RegisterOutputType(ListenerRuleConditionArrayOutput{})
}

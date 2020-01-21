// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ListenerRuleConditionHttpHeader

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ListenerRuleConditionHttpHeader struct {
	// Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
	HttpHeaderName string `pulumi:"httpHeaderName"`
	// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
	Values []string `pulumi:"values"`
}

type ListenerRuleConditionHttpHeaderInput interface {
	pulumi.Input

	ToListenerRuleConditionHttpHeaderOutput() ListenerRuleConditionHttpHeaderOutput
	ToListenerRuleConditionHttpHeaderOutputWithContext(context.Context) ListenerRuleConditionHttpHeaderOutput
}

type ListenerRuleConditionHttpHeaderArgs struct {
	// Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
	HttpHeaderName pulumi.StringInput `pulumi:"httpHeaderName"`
	// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (ListenerRuleConditionHttpHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleConditionHttpHeader)(nil)).Elem()
}

func (i ListenerRuleConditionHttpHeaderArgs) ToListenerRuleConditionHttpHeaderOutput() ListenerRuleConditionHttpHeaderOutput {
	return i.ToListenerRuleConditionHttpHeaderOutputWithContext(context.Background())
}

func (i ListenerRuleConditionHttpHeaderArgs) ToListenerRuleConditionHttpHeaderOutputWithContext(ctx context.Context) ListenerRuleConditionHttpHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleConditionHttpHeaderOutput)
}

func (i ListenerRuleConditionHttpHeaderArgs) ToListenerRuleConditionHttpHeaderPtrOutput() ListenerRuleConditionHttpHeaderPtrOutput {
	return i.ToListenerRuleConditionHttpHeaderPtrOutputWithContext(context.Background())
}

func (i ListenerRuleConditionHttpHeaderArgs) ToListenerRuleConditionHttpHeaderPtrOutputWithContext(ctx context.Context) ListenerRuleConditionHttpHeaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleConditionHttpHeaderOutput).ToListenerRuleConditionHttpHeaderPtrOutputWithContext(ctx)
}

type ListenerRuleConditionHttpHeaderPtrInput interface {
	pulumi.Input

	ToListenerRuleConditionHttpHeaderPtrOutput() ListenerRuleConditionHttpHeaderPtrOutput
	ToListenerRuleConditionHttpHeaderPtrOutputWithContext(context.Context) ListenerRuleConditionHttpHeaderPtrOutput
}

type listenerRuleConditionHttpHeaderPtrType ListenerRuleConditionHttpHeaderArgs

func ListenerRuleConditionHttpHeaderPtr(v *ListenerRuleConditionHttpHeaderArgs) ListenerRuleConditionHttpHeaderPtrInput {	return (*listenerRuleConditionHttpHeaderPtrType)(v)
}

func (*listenerRuleConditionHttpHeaderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRuleConditionHttpHeader)(nil)).Elem()
}

func (i *listenerRuleConditionHttpHeaderPtrType) ToListenerRuleConditionHttpHeaderPtrOutput() ListenerRuleConditionHttpHeaderPtrOutput {
	return i.ToListenerRuleConditionHttpHeaderPtrOutputWithContext(context.Background())
}

func (i *listenerRuleConditionHttpHeaderPtrType) ToListenerRuleConditionHttpHeaderPtrOutputWithContext(ctx context.Context) ListenerRuleConditionHttpHeaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleConditionHttpHeaderPtrOutput)
}

type ListenerRuleConditionHttpHeaderOutput struct { *pulumi.OutputState }

func (ListenerRuleConditionHttpHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleConditionHttpHeader)(nil)).Elem()
}

func (o ListenerRuleConditionHttpHeaderOutput) ToListenerRuleConditionHttpHeaderOutput() ListenerRuleConditionHttpHeaderOutput {
	return o
}

func (o ListenerRuleConditionHttpHeaderOutput) ToListenerRuleConditionHttpHeaderOutputWithContext(ctx context.Context) ListenerRuleConditionHttpHeaderOutput {
	return o
}

func (o ListenerRuleConditionHttpHeaderOutput) ToListenerRuleConditionHttpHeaderPtrOutput() ListenerRuleConditionHttpHeaderPtrOutput {
	return o.ToListenerRuleConditionHttpHeaderPtrOutputWithContext(context.Background())
}

func (o ListenerRuleConditionHttpHeaderOutput) ToListenerRuleConditionHttpHeaderPtrOutputWithContext(ctx context.Context) ListenerRuleConditionHttpHeaderPtrOutput {
	return o.ApplyT(func(v ListenerRuleConditionHttpHeader) *ListenerRuleConditionHttpHeader {
		return &v
	}).(ListenerRuleConditionHttpHeaderPtrOutput)
}
// Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
func (o ListenerRuleConditionHttpHeaderOutput) HttpHeaderName() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerRuleConditionHttpHeader) string { return v.HttpHeaderName }).(pulumi.StringOutput)
}

// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
func (o ListenerRuleConditionHttpHeaderOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ListenerRuleConditionHttpHeader) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ListenerRuleConditionHttpHeaderPtrOutput struct { *pulumi.OutputState}

func (ListenerRuleConditionHttpHeaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRuleConditionHttpHeader)(nil)).Elem()
}

func (o ListenerRuleConditionHttpHeaderPtrOutput) ToListenerRuleConditionHttpHeaderPtrOutput() ListenerRuleConditionHttpHeaderPtrOutput {
	return o
}

func (o ListenerRuleConditionHttpHeaderPtrOutput) ToListenerRuleConditionHttpHeaderPtrOutputWithContext(ctx context.Context) ListenerRuleConditionHttpHeaderPtrOutput {
	return o
}

func (o ListenerRuleConditionHttpHeaderPtrOutput) Elem() ListenerRuleConditionHttpHeaderOutput {
	return o.ApplyT(func (v *ListenerRuleConditionHttpHeader) ListenerRuleConditionHttpHeader { return *v }).(ListenerRuleConditionHttpHeaderOutput)
}

// Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
func (o ListenerRuleConditionHttpHeaderPtrOutput) HttpHeaderName() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerRuleConditionHttpHeader) string { return v.HttpHeaderName }).(pulumi.StringOutput)
}

// Query string pairs or values to match. Query String Value blocks documented below. Multiple `values` blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '\*' or '?' character in a query string, escape the character with a backslash (\\). Only one pair needs to match for the condition to be satisfied.
func (o ListenerRuleConditionHttpHeaderPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ListenerRuleConditionHttpHeader) []string { return v.Values }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerRuleConditionHttpHeaderOutput{})
	pulumi.RegisterOutputType(ListenerRuleConditionHttpHeaderPtrOutput{})
}

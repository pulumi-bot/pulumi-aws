// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ListenerRuleActionRedirect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ListenerRuleActionRedirect struct {
	// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
	Host *string `pulumi:"host"`
	// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
	Path *string `pulumi:"path"`
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port *string `pulumi:"port"`
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol *string `pulumi:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
	Query *string `pulumi:"query"`
	// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
	StatusCode string `pulumi:"statusCode"`
}

type ListenerRuleActionRedirectInput interface {
	pulumi.Input

	ToListenerRuleActionRedirectOutput() ListenerRuleActionRedirectOutput
	ToListenerRuleActionRedirectOutputWithContext(context.Context) ListenerRuleActionRedirectOutput
}

type ListenerRuleActionRedirectArgs struct {
	// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.StringPtrInput `pulumi:"port"`
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
	Query pulumi.StringPtrInput `pulumi:"query"`
	// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
	StatusCode pulumi.StringInput `pulumi:"statusCode"`
}

func (ListenerRuleActionRedirectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleActionRedirect)(nil)).Elem()
}

func (i ListenerRuleActionRedirectArgs) ToListenerRuleActionRedirectOutput() ListenerRuleActionRedirectOutput {
	return i.ToListenerRuleActionRedirectOutputWithContext(context.Background())
}

func (i ListenerRuleActionRedirectArgs) ToListenerRuleActionRedirectOutputWithContext(ctx context.Context) ListenerRuleActionRedirectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleActionRedirectOutput)
}

func (i ListenerRuleActionRedirectArgs) ToListenerRuleActionRedirectPtrOutput() ListenerRuleActionRedirectPtrOutput {
	return i.ToListenerRuleActionRedirectPtrOutputWithContext(context.Background())
}

func (i ListenerRuleActionRedirectArgs) ToListenerRuleActionRedirectPtrOutputWithContext(ctx context.Context) ListenerRuleActionRedirectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleActionRedirectOutput).ToListenerRuleActionRedirectPtrOutputWithContext(ctx)
}

type ListenerRuleActionRedirectPtrInput interface {
	pulumi.Input

	ToListenerRuleActionRedirectPtrOutput() ListenerRuleActionRedirectPtrOutput
	ToListenerRuleActionRedirectPtrOutputWithContext(context.Context) ListenerRuleActionRedirectPtrOutput
}

type listenerRuleActionRedirectPtrType ListenerRuleActionRedirectArgs

func ListenerRuleActionRedirectPtr(v *ListenerRuleActionRedirectArgs) ListenerRuleActionRedirectPtrInput {	return (*listenerRuleActionRedirectPtrType)(v)
}

func (*listenerRuleActionRedirectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRuleActionRedirect)(nil)).Elem()
}

func (i *listenerRuleActionRedirectPtrType) ToListenerRuleActionRedirectPtrOutput() ListenerRuleActionRedirectPtrOutput {
	return i.ToListenerRuleActionRedirectPtrOutputWithContext(context.Background())
}

func (i *listenerRuleActionRedirectPtrType) ToListenerRuleActionRedirectPtrOutputWithContext(ctx context.Context) ListenerRuleActionRedirectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleActionRedirectPtrOutput)
}

type ListenerRuleActionRedirectOutput struct { *pulumi.OutputState }

func (ListenerRuleActionRedirectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleActionRedirect)(nil)).Elem()
}

func (o ListenerRuleActionRedirectOutput) ToListenerRuleActionRedirectOutput() ListenerRuleActionRedirectOutput {
	return o
}

func (o ListenerRuleActionRedirectOutput) ToListenerRuleActionRedirectOutputWithContext(ctx context.Context) ListenerRuleActionRedirectOutput {
	return o
}

func (o ListenerRuleActionRedirectOutput) ToListenerRuleActionRedirectPtrOutput() ListenerRuleActionRedirectPtrOutput {
	return o.ToListenerRuleActionRedirectPtrOutputWithContext(context.Background())
}

func (o ListenerRuleActionRedirectOutput) ToListenerRuleActionRedirectPtrOutputWithContext(ctx context.Context) ListenerRuleActionRedirectPtrOutput {
	return o.ApplyT(func(v ListenerRuleActionRedirect) *ListenerRuleActionRedirect {
		return &v
	}).(ListenerRuleActionRedirectPtrOutput)
}
// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
func (o ListenerRuleActionRedirectOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
func (o ListenerRuleActionRedirectOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
func (o ListenerRuleActionRedirectOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Port }).(pulumi.StringPtrOutput)
}

// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
func (o ListenerRuleActionRedirectOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
func (o ListenerRuleActionRedirectOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
func (o ListenerRuleActionRedirectOutput) StatusCode() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) string { return v.StatusCode }).(pulumi.StringOutput)
}

type ListenerRuleActionRedirectPtrOutput struct { *pulumi.OutputState}

func (ListenerRuleActionRedirectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRuleActionRedirect)(nil)).Elem()
}

func (o ListenerRuleActionRedirectPtrOutput) ToListenerRuleActionRedirectPtrOutput() ListenerRuleActionRedirectPtrOutput {
	return o
}

func (o ListenerRuleActionRedirectPtrOutput) ToListenerRuleActionRedirectPtrOutputWithContext(ctx context.Context) ListenerRuleActionRedirectPtrOutput {
	return o
}

func (o ListenerRuleActionRedirectPtrOutput) Elem() ListenerRuleActionRedirectOutput {
	return o.ApplyT(func (v *ListenerRuleActionRedirect) ListenerRuleActionRedirect { return *v }).(ListenerRuleActionRedirectOutput)
}

// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
func (o ListenerRuleActionRedirectPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
func (o ListenerRuleActionRedirectPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
func (o ListenerRuleActionRedirectPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Port }).(pulumi.StringPtrOutput)
}

// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
func (o ListenerRuleActionRedirectPtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
func (o ListenerRuleActionRedirectPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
func (o ListenerRuleActionRedirectPtrOutput) StatusCode() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerRuleActionRedirect) string { return v.StatusCode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerRuleActionRedirectOutput{})
	pulumi.RegisterOutputType(ListenerRuleActionRedirectPtrOutput{})
}

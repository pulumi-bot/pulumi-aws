// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ListenerDefaultActionFixedResponse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ListenerDefaultActionFixedResponse struct {
	// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
	ContentType string `pulumi:"contentType"`
	// The message body.
	MessageBody *string `pulumi:"messageBody"`
	// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
	StatusCode *string `pulumi:"statusCode"`
}

type ListenerDefaultActionFixedResponseInput interface {
	pulumi.Input

	ToListenerDefaultActionFixedResponseOutput() ListenerDefaultActionFixedResponseOutput
	ToListenerDefaultActionFixedResponseOutputWithContext(context.Context) ListenerDefaultActionFixedResponseOutput
}

type ListenerDefaultActionFixedResponseArgs struct {
	// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
	ContentType pulumi.StringInput `pulumi:"contentType"`
	// The message body.
	MessageBody pulumi.StringPtrInput `pulumi:"messageBody"`
	// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
	StatusCode pulumi.StringPtrInput `pulumi:"statusCode"`
}

func (ListenerDefaultActionFixedResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerDefaultActionFixedResponse)(nil)).Elem()
}

func (i ListenerDefaultActionFixedResponseArgs) ToListenerDefaultActionFixedResponseOutput() ListenerDefaultActionFixedResponseOutput {
	return i.ToListenerDefaultActionFixedResponseOutputWithContext(context.Background())
}

func (i ListenerDefaultActionFixedResponseArgs) ToListenerDefaultActionFixedResponseOutputWithContext(ctx context.Context) ListenerDefaultActionFixedResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerDefaultActionFixedResponseOutput)
}

func (i ListenerDefaultActionFixedResponseArgs) ToListenerDefaultActionFixedResponsePtrOutput() ListenerDefaultActionFixedResponsePtrOutput {
	return i.ToListenerDefaultActionFixedResponsePtrOutputWithContext(context.Background())
}

func (i ListenerDefaultActionFixedResponseArgs) ToListenerDefaultActionFixedResponsePtrOutputWithContext(ctx context.Context) ListenerDefaultActionFixedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerDefaultActionFixedResponseOutput).ToListenerDefaultActionFixedResponsePtrOutputWithContext(ctx)
}

type ListenerDefaultActionFixedResponsePtrInput interface {
	pulumi.Input

	ToListenerDefaultActionFixedResponsePtrOutput() ListenerDefaultActionFixedResponsePtrOutput
	ToListenerDefaultActionFixedResponsePtrOutputWithContext(context.Context) ListenerDefaultActionFixedResponsePtrOutput
}

type listenerDefaultActionFixedResponsePtrType ListenerDefaultActionFixedResponseArgs

func ListenerDefaultActionFixedResponsePtr(v *ListenerDefaultActionFixedResponseArgs) ListenerDefaultActionFixedResponsePtrInput {	return (*listenerDefaultActionFixedResponsePtrType)(v)
}

func (*listenerDefaultActionFixedResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerDefaultActionFixedResponse)(nil)).Elem()
}

func (i *listenerDefaultActionFixedResponsePtrType) ToListenerDefaultActionFixedResponsePtrOutput() ListenerDefaultActionFixedResponsePtrOutput {
	return i.ToListenerDefaultActionFixedResponsePtrOutputWithContext(context.Background())
}

func (i *listenerDefaultActionFixedResponsePtrType) ToListenerDefaultActionFixedResponsePtrOutputWithContext(ctx context.Context) ListenerDefaultActionFixedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerDefaultActionFixedResponsePtrOutput)
}

type ListenerDefaultActionFixedResponseOutput struct { *pulumi.OutputState }

func (ListenerDefaultActionFixedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerDefaultActionFixedResponse)(nil)).Elem()
}

func (o ListenerDefaultActionFixedResponseOutput) ToListenerDefaultActionFixedResponseOutput() ListenerDefaultActionFixedResponseOutput {
	return o
}

func (o ListenerDefaultActionFixedResponseOutput) ToListenerDefaultActionFixedResponseOutputWithContext(ctx context.Context) ListenerDefaultActionFixedResponseOutput {
	return o
}

func (o ListenerDefaultActionFixedResponseOutput) ToListenerDefaultActionFixedResponsePtrOutput() ListenerDefaultActionFixedResponsePtrOutput {
	return o.ToListenerDefaultActionFixedResponsePtrOutputWithContext(context.Background())
}

func (o ListenerDefaultActionFixedResponseOutput) ToListenerDefaultActionFixedResponsePtrOutputWithContext(ctx context.Context) ListenerDefaultActionFixedResponsePtrOutput {
	return o.ApplyT(func(v ListenerDefaultActionFixedResponse) *ListenerDefaultActionFixedResponse {
		return &v
	}).(ListenerDefaultActionFixedResponsePtrOutput)
}
// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
func (o ListenerDefaultActionFixedResponseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerDefaultActionFixedResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

// The message body.
func (o ListenerDefaultActionFixedResponseOutput) MessageBody() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerDefaultActionFixedResponse) *string { return v.MessageBody }).(pulumi.StringPtrOutput)
}

// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
func (o ListenerDefaultActionFixedResponseOutput) StatusCode() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerDefaultActionFixedResponse) *string { return v.StatusCode }).(pulumi.StringPtrOutput)
}

type ListenerDefaultActionFixedResponsePtrOutput struct { *pulumi.OutputState}

func (ListenerDefaultActionFixedResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerDefaultActionFixedResponse)(nil)).Elem()
}

func (o ListenerDefaultActionFixedResponsePtrOutput) ToListenerDefaultActionFixedResponsePtrOutput() ListenerDefaultActionFixedResponsePtrOutput {
	return o
}

func (o ListenerDefaultActionFixedResponsePtrOutput) ToListenerDefaultActionFixedResponsePtrOutputWithContext(ctx context.Context) ListenerDefaultActionFixedResponsePtrOutput {
	return o
}

func (o ListenerDefaultActionFixedResponsePtrOutput) Elem() ListenerDefaultActionFixedResponseOutput {
	return o.ApplyT(func (v *ListenerDefaultActionFixedResponse) ListenerDefaultActionFixedResponse { return *v }).(ListenerDefaultActionFixedResponseOutput)
}

// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
func (o ListenerDefaultActionFixedResponsePtrOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerDefaultActionFixedResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

// The message body.
func (o ListenerDefaultActionFixedResponsePtrOutput) MessageBody() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerDefaultActionFixedResponse) *string { return v.MessageBody }).(pulumi.StringPtrOutput)
}

// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
func (o ListenerDefaultActionFixedResponsePtrOutput) StatusCode() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerDefaultActionFixedResponse) *string { return v.StatusCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerDefaultActionFixedResponseOutput{})
	pulumi.RegisterOutputType(ListenerDefaultActionFixedResponsePtrOutput{})
}

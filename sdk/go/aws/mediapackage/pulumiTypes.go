// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ChannelHlsIngest struct {
	// A list of the ingest endpoints
	IngestEndpoints []ChannelHlsIngestIngestEndpoint `pulumi:"ingestEndpoints"`
}

type ChannelHlsIngestInput interface {
	pulumi.Input

	ToChannelHlsIngestOutput() ChannelHlsIngestOutput
	ToChannelHlsIngestOutputWithContext(context.Context) ChannelHlsIngestOutput
}

type ChannelHlsIngestArgs struct {
	// A list of the ingest endpoints
	IngestEndpoints ChannelHlsIngestIngestEndpointArrayInput `pulumi:"ingestEndpoints"`
}

func (ChannelHlsIngestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngest)(nil)).Elem()
}

func (i ChannelHlsIngestArgs) ToChannelHlsIngestOutput() ChannelHlsIngestOutput {
	return i.ToChannelHlsIngestOutputWithContext(context.Background())
}

func (i ChannelHlsIngestArgs) ToChannelHlsIngestOutputWithContext(ctx context.Context) ChannelHlsIngestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestOutput)
}

type ChannelHlsIngestArrayInput interface {
	pulumi.Input

	ToChannelHlsIngestArrayOutput() ChannelHlsIngestArrayOutput
	ToChannelHlsIngestArrayOutputWithContext(context.Context) ChannelHlsIngestArrayOutput
}

type ChannelHlsIngestArray []ChannelHlsIngestInput

func (ChannelHlsIngestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngest)(nil)).Elem()
}

func (i ChannelHlsIngestArray) ToChannelHlsIngestArrayOutput() ChannelHlsIngestArrayOutput {
	return i.ToChannelHlsIngestArrayOutputWithContext(context.Background())
}

func (i ChannelHlsIngestArray) ToChannelHlsIngestArrayOutputWithContext(ctx context.Context) ChannelHlsIngestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestArrayOutput)
}

type ChannelHlsIngestOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngest)(nil)).Elem()
}

func (o ChannelHlsIngestOutput) ToChannelHlsIngestOutput() ChannelHlsIngestOutput {
	return o
}

func (o ChannelHlsIngestOutput) ToChannelHlsIngestOutputWithContext(ctx context.Context) ChannelHlsIngestOutput {
	return o
}

// A list of the ingest endpoints
func (o ChannelHlsIngestOutput) IngestEndpoints() ChannelHlsIngestIngestEndpointArrayOutput {
	return o.ApplyT(func(v ChannelHlsIngest) []ChannelHlsIngestIngestEndpoint { return v.IngestEndpoints }).(ChannelHlsIngestIngestEndpointArrayOutput)
}

type ChannelHlsIngestArrayOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngest)(nil)).Elem()
}

func (o ChannelHlsIngestArrayOutput) ToChannelHlsIngestArrayOutput() ChannelHlsIngestArrayOutput {
	return o
}

func (o ChannelHlsIngestArrayOutput) ToChannelHlsIngestArrayOutputWithContext(ctx context.Context) ChannelHlsIngestArrayOutput {
	return o
}

func (o ChannelHlsIngestArrayOutput) Index(i pulumi.IntInput) ChannelHlsIngestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelHlsIngest {
		return vs[0].([]ChannelHlsIngest)[vs[1].(int)]
	}).(ChannelHlsIngestOutput)
}

type ChannelHlsIngestIngestEndpoint struct {
	// The password
	Password string `pulumi:"password"`
	// The URL
	Url string `pulumi:"url"`
	// The username
	Username string `pulumi:"username"`
}

type ChannelHlsIngestIngestEndpointInput interface {
	pulumi.Input

	ToChannelHlsIngestIngestEndpointOutput() ChannelHlsIngestIngestEndpointOutput
	ToChannelHlsIngestIngestEndpointOutputWithContext(context.Context) ChannelHlsIngestIngestEndpointOutput
}

type ChannelHlsIngestIngestEndpointArgs struct {
	// The password
	Password pulumi.StringInput `pulumi:"password"`
	// The URL
	Url pulumi.StringInput `pulumi:"url"`
	// The username
	Username pulumi.StringInput `pulumi:"username"`
}

func (ChannelHlsIngestIngestEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngestIngestEndpoint)(nil)).Elem()
}

func (i ChannelHlsIngestIngestEndpointArgs) ToChannelHlsIngestIngestEndpointOutput() ChannelHlsIngestIngestEndpointOutput {
	return i.ToChannelHlsIngestIngestEndpointOutputWithContext(context.Background())
}

func (i ChannelHlsIngestIngestEndpointArgs) ToChannelHlsIngestIngestEndpointOutputWithContext(ctx context.Context) ChannelHlsIngestIngestEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestIngestEndpointOutput)
}

type ChannelHlsIngestIngestEndpointArrayInput interface {
	pulumi.Input

	ToChannelHlsIngestIngestEndpointArrayOutput() ChannelHlsIngestIngestEndpointArrayOutput
	ToChannelHlsIngestIngestEndpointArrayOutputWithContext(context.Context) ChannelHlsIngestIngestEndpointArrayOutput
}

type ChannelHlsIngestIngestEndpointArray []ChannelHlsIngestIngestEndpointInput

func (ChannelHlsIngestIngestEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngestIngestEndpoint)(nil)).Elem()
}

func (i ChannelHlsIngestIngestEndpointArray) ToChannelHlsIngestIngestEndpointArrayOutput() ChannelHlsIngestIngestEndpointArrayOutput {
	return i.ToChannelHlsIngestIngestEndpointArrayOutputWithContext(context.Background())
}

func (i ChannelHlsIngestIngestEndpointArray) ToChannelHlsIngestIngestEndpointArrayOutputWithContext(ctx context.Context) ChannelHlsIngestIngestEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestIngestEndpointArrayOutput)
}

type ChannelHlsIngestIngestEndpointOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestIngestEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngestIngestEndpoint)(nil)).Elem()
}

func (o ChannelHlsIngestIngestEndpointOutput) ToChannelHlsIngestIngestEndpointOutput() ChannelHlsIngestIngestEndpointOutput {
	return o
}

func (o ChannelHlsIngestIngestEndpointOutput) ToChannelHlsIngestIngestEndpointOutputWithContext(ctx context.Context) ChannelHlsIngestIngestEndpointOutput {
	return o
}

// The password
func (o ChannelHlsIngestIngestEndpointOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ChannelHlsIngestIngestEndpoint) string { return v.Password }).(pulumi.StringOutput)
}

// The URL
func (o ChannelHlsIngestIngestEndpointOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v ChannelHlsIngestIngestEndpoint) string { return v.Url }).(pulumi.StringOutput)
}

// The username
func (o ChannelHlsIngestIngestEndpointOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ChannelHlsIngestIngestEndpoint) string { return v.Username }).(pulumi.StringOutput)
}

type ChannelHlsIngestIngestEndpointArrayOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestIngestEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngestIngestEndpoint)(nil)).Elem()
}

func (o ChannelHlsIngestIngestEndpointArrayOutput) ToChannelHlsIngestIngestEndpointArrayOutput() ChannelHlsIngestIngestEndpointArrayOutput {
	return o
}

func (o ChannelHlsIngestIngestEndpointArrayOutput) ToChannelHlsIngestIngestEndpointArrayOutputWithContext(ctx context.Context) ChannelHlsIngestIngestEndpointArrayOutput {
	return o
}

func (o ChannelHlsIngestIngestEndpointArrayOutput) Index(i pulumi.IntInput) ChannelHlsIngestIngestEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelHlsIngestIngestEndpoint {
		return vs[0].([]ChannelHlsIngestIngestEndpoint)[vs[1].(int)]
	}).(ChannelHlsIngestIngestEndpointOutput)
}

type ChannelHlsIngestState struct {
	// A list of the ingest endpoints
	IngestEndpoints []ChannelHlsIngestStateIngestEndpoint `pulumi:"ingestEndpoints"`
}

type ChannelHlsIngestStateInput interface {
	pulumi.Input

	ToChannelHlsIngestStateOutput() ChannelHlsIngestStateOutput
	ToChannelHlsIngestStateOutputWithContext(context.Context) ChannelHlsIngestStateOutput
}

type ChannelHlsIngestStateArgs struct {
	// A list of the ingest endpoints
	IngestEndpoints ChannelHlsIngestStateIngestEndpointArrayInput `pulumi:"ingestEndpoints"`
}

func (ChannelHlsIngestStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngestState)(nil)).Elem()
}

func (i ChannelHlsIngestStateArgs) ToChannelHlsIngestStateOutput() ChannelHlsIngestStateOutput {
	return i.ToChannelHlsIngestStateOutputWithContext(context.Background())
}

func (i ChannelHlsIngestStateArgs) ToChannelHlsIngestStateOutputWithContext(ctx context.Context) ChannelHlsIngestStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestStateOutput)
}

type ChannelHlsIngestStateArrayInput interface {
	pulumi.Input

	ToChannelHlsIngestStateArrayOutput() ChannelHlsIngestStateArrayOutput
	ToChannelHlsIngestStateArrayOutputWithContext(context.Context) ChannelHlsIngestStateArrayOutput
}

type ChannelHlsIngestStateArray []ChannelHlsIngestStateInput

func (ChannelHlsIngestStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngestState)(nil)).Elem()
}

func (i ChannelHlsIngestStateArray) ToChannelHlsIngestStateArrayOutput() ChannelHlsIngestStateArrayOutput {
	return i.ToChannelHlsIngestStateArrayOutputWithContext(context.Background())
}

func (i ChannelHlsIngestStateArray) ToChannelHlsIngestStateArrayOutputWithContext(ctx context.Context) ChannelHlsIngestStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestStateArrayOutput)
}

type ChannelHlsIngestStateOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngestState)(nil)).Elem()
}

func (o ChannelHlsIngestStateOutput) ToChannelHlsIngestStateOutput() ChannelHlsIngestStateOutput {
	return o
}

func (o ChannelHlsIngestStateOutput) ToChannelHlsIngestStateOutputWithContext(ctx context.Context) ChannelHlsIngestStateOutput {
	return o
}

// A list of the ingest endpoints
func (o ChannelHlsIngestStateOutput) IngestEndpoints() ChannelHlsIngestStateIngestEndpointArrayOutput {
	return o.ApplyT(func(v ChannelHlsIngestState) []ChannelHlsIngestStateIngestEndpoint { return v.IngestEndpoints }).(ChannelHlsIngestStateIngestEndpointArrayOutput)
}

type ChannelHlsIngestStateArrayOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngestState)(nil)).Elem()
}

func (o ChannelHlsIngestStateArrayOutput) ToChannelHlsIngestStateArrayOutput() ChannelHlsIngestStateArrayOutput {
	return o
}

func (o ChannelHlsIngestStateArrayOutput) ToChannelHlsIngestStateArrayOutputWithContext(ctx context.Context) ChannelHlsIngestStateArrayOutput {
	return o
}

func (o ChannelHlsIngestStateArrayOutput) Index(i pulumi.IntInput) ChannelHlsIngestStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelHlsIngestState {
		return vs[0].([]ChannelHlsIngestState)[vs[1].(int)]
	}).(ChannelHlsIngestStateOutput)
}

type ChannelHlsIngestStateIngestEndpoint struct {
	// The password
	Password *string `pulumi:"password"`
	// The URL
	Url *string `pulumi:"url"`
	// The username
	Username *string `pulumi:"username"`
}

type ChannelHlsIngestStateIngestEndpointInput interface {
	pulumi.Input

	ToChannelHlsIngestStateIngestEndpointOutput() ChannelHlsIngestStateIngestEndpointOutput
	ToChannelHlsIngestStateIngestEndpointOutputWithContext(context.Context) ChannelHlsIngestStateIngestEndpointOutput
}

type ChannelHlsIngestStateIngestEndpointArgs struct {
	// The password
	Password pulumi.StringPtrInput `pulumi:"password"`
	// The URL
	Url pulumi.StringPtrInput `pulumi:"url"`
	// The username
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (ChannelHlsIngestStateIngestEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngestStateIngestEndpoint)(nil)).Elem()
}

func (i ChannelHlsIngestStateIngestEndpointArgs) ToChannelHlsIngestStateIngestEndpointOutput() ChannelHlsIngestStateIngestEndpointOutput {
	return i.ToChannelHlsIngestStateIngestEndpointOutputWithContext(context.Background())
}

func (i ChannelHlsIngestStateIngestEndpointArgs) ToChannelHlsIngestStateIngestEndpointOutputWithContext(ctx context.Context) ChannelHlsIngestStateIngestEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestStateIngestEndpointOutput)
}

type ChannelHlsIngestStateIngestEndpointArrayInput interface {
	pulumi.Input

	ToChannelHlsIngestStateIngestEndpointArrayOutput() ChannelHlsIngestStateIngestEndpointArrayOutput
	ToChannelHlsIngestStateIngestEndpointArrayOutputWithContext(context.Context) ChannelHlsIngestStateIngestEndpointArrayOutput
}

type ChannelHlsIngestStateIngestEndpointArray []ChannelHlsIngestStateIngestEndpointInput

func (ChannelHlsIngestStateIngestEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngestStateIngestEndpoint)(nil)).Elem()
}

func (i ChannelHlsIngestStateIngestEndpointArray) ToChannelHlsIngestStateIngestEndpointArrayOutput() ChannelHlsIngestStateIngestEndpointArrayOutput {
	return i.ToChannelHlsIngestStateIngestEndpointArrayOutputWithContext(context.Background())
}

func (i ChannelHlsIngestStateIngestEndpointArray) ToChannelHlsIngestStateIngestEndpointArrayOutputWithContext(ctx context.Context) ChannelHlsIngestStateIngestEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelHlsIngestStateIngestEndpointArrayOutput)
}

type ChannelHlsIngestStateIngestEndpointOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestStateIngestEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelHlsIngestStateIngestEndpoint)(nil)).Elem()
}

func (o ChannelHlsIngestStateIngestEndpointOutput) ToChannelHlsIngestStateIngestEndpointOutput() ChannelHlsIngestStateIngestEndpointOutput {
	return o
}

func (o ChannelHlsIngestStateIngestEndpointOutput) ToChannelHlsIngestStateIngestEndpointOutputWithContext(ctx context.Context) ChannelHlsIngestStateIngestEndpointOutput {
	return o
}

// The password
func (o ChannelHlsIngestStateIngestEndpointOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelHlsIngestStateIngestEndpoint) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// The URL
func (o ChannelHlsIngestStateIngestEndpointOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelHlsIngestStateIngestEndpoint) *string { return v.Url }).(pulumi.StringPtrOutput)
}

// The username
func (o ChannelHlsIngestStateIngestEndpointOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelHlsIngestStateIngestEndpoint) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ChannelHlsIngestStateIngestEndpointArrayOutput struct{ *pulumi.OutputState }

func (ChannelHlsIngestStateIngestEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelHlsIngestStateIngestEndpoint)(nil)).Elem()
}

func (o ChannelHlsIngestStateIngestEndpointArrayOutput) ToChannelHlsIngestStateIngestEndpointArrayOutput() ChannelHlsIngestStateIngestEndpointArrayOutput {
	return o
}

func (o ChannelHlsIngestStateIngestEndpointArrayOutput) ToChannelHlsIngestStateIngestEndpointArrayOutputWithContext(ctx context.Context) ChannelHlsIngestStateIngestEndpointArrayOutput {
	return o
}

func (o ChannelHlsIngestStateIngestEndpointArrayOutput) Index(i pulumi.IntInput) ChannelHlsIngestStateIngestEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelHlsIngestStateIngestEndpoint {
		return vs[0].([]ChannelHlsIngestStateIngestEndpoint)[vs[1].(int)]
	}).(ChannelHlsIngestStateIngestEndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelHlsIngestOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestArrayOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestIngestEndpointOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestIngestEndpointArrayOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestStateOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestStateArrayOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestStateIngestEndpointOutput{})
	pulumi.RegisterOutputType(ChannelHlsIngestStateIngestEndpointArrayOutput{})
}

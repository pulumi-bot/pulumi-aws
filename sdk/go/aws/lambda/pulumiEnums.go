// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// See https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
type Runtime string

const (
	RuntimeDotnetCore2d1 = Runtime("dotnetcore2.1")
	RuntimeDotnetCore3d1 = Runtime("dotnetcore3.1")
	RuntimeGo1dx         = Runtime("go1.x")
	RuntimeJava8         = Runtime("java8")
	RuntimeJava8AL2      = Runtime("java8.al2")
	RuntimeJava11        = Runtime("java11")
	RuntimeRuby2d5       = Runtime("ruby2.5")
	RuntimeRuby2d7       = Runtime("ruby2.7")
	RuntimeNodeJS10dX    = Runtime("nodejs10.x")
	RuntimeNodeJS12dX    = Runtime("nodejs12.x")
	RuntimeNodeJS14dX    = Runtime("nodejs14.x")
	RuntimePython2d7     = Runtime("python2.7")
	RuntimePython3d6     = Runtime("python3.6")
	RuntimePython3d7     = Runtime("python3.7")
	RuntimePython3d8     = Runtime("python3.8")
	RuntimeCustom        = Runtime("provided")
	RuntimeCustomAL2     = Runtime("provided.al2")
)

func (Runtime) ElementType() reflect.Type {
	return reflect.TypeOf((*Runtime)(nil)).Elem()
}

func (e Runtime) ToRuntimeOutput() RuntimeOutput {
	return pulumi.ToOutput(e).(RuntimeOutput)
}

func (e Runtime) ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RuntimeOutput)
}

func (e Runtime) ToRuntimePtrOutput() RuntimePtrOutput {
	return e.ToRuntimePtrOutputWithContext(context.Background())
}

func (e Runtime) ToRuntimePtrOutputWithContext(ctx context.Context) RuntimePtrOutput {
	return Runtime(e).ToRuntimeOutputWithContext(ctx).ToRuntimePtrOutputWithContext(ctx)
}

func (e Runtime) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Runtime) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Runtime) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Runtime) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RuntimeOutput struct{ *pulumi.OutputState }

func (RuntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Runtime)(nil)).Elem()
}

func (o RuntimeOutput) ToRuntimeOutput() RuntimeOutput {
	return o
}

func (o RuntimeOutput) ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput {
	return o
}

func (o RuntimeOutput) ToRuntimePtrOutput() RuntimePtrOutput {
	return o.ToRuntimePtrOutputWithContext(context.Background())
}

func (o RuntimeOutput) ToRuntimePtrOutputWithContext(ctx context.Context) RuntimePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Runtime) *Runtime {
		return &v
	}).(RuntimePtrOutput)
}

type RuntimePtrOutput struct{ *pulumi.OutputState }

func (RuntimePtrOutput) ElementType() reflect.Type {
	return runtimePtrType
}

func (o RuntimePtrOutput) ToRuntimePtrOutput() RuntimePtrOutput {
	return o
}

func (o RuntimePtrOutput) ToRuntimePtrOutputWithContext(ctx context.Context) RuntimePtrOutput {
	return o
}

func (o RuntimePtrOutput) Elem() RuntimeOutput {
	return o.ApplyT(func(v *Runtime) Runtime {
		var ret Runtime
		if v != nil {
			ret = *v
		}
		return ret
	}).(RuntimeOutput)
}

// RuntimeInput is an input type that accepts RuntimeArgs and RuntimeOutput values.
// You can construct a concrete instance of `RuntimeInput` via:
//
//          RuntimeArgs{...}
type RuntimeInput interface {
	pulumi.Input

	ToRuntimeOutput() RuntimeOutput
	ToRuntimeOutputWithContext(context.Context) RuntimeOutput
}

var runtimePtrType = reflect.TypeOf((**Runtime)(nil)).Elem()

type RuntimePtrInput interface {
	pulumi.Input

	ToRuntimePtrOutput() RuntimePtrOutput
	ToRuntimePtrOutputWithContext(context.Context) RuntimePtrOutput
}

type runtimePtr string

func RuntimePtr(v string) RuntimePtrInput {
	return (*runtimePtr)(&v)
}

func (*runtimePtr) ElementType() reflect.Type {
	return runtimePtrType
}

func (in *runtimePtr) ToRuntimePtrOutput() RuntimePtrOutput {
	return pulumi.ToOutput(in).(RuntimePtrOutput)
}

func (in *runtimePtr) ToRuntimePtrOutputWithContext(ctx context.Context) RuntimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RuntimePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RuntimeOutput{})
	pulumi.RegisterOutputType(RuntimePtrOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
//
// Provides information about a Load Balancer Listener.
//
// This data source can prove useful when a module accepts an LB Listener as an input variable and needs to know the LB it is attached to, or other information specific to the listener in question.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		listenerArn := cfg.Require("listenerArn")
// 		opt0 := listenerArn
// 		_, err := lb.LookupListener(ctx, &lb.LookupListenerArgs{
// 			Arn: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "default-public"
// 		selected, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
// 			Name: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt2 := selected.Arn
// 		opt3 := 443
// 		_, err = lb.LookupListener(ctx, &lb.LookupListenerArgs{
// 			LoadBalancerArn: &opt2,
// 			Port:            &opt3,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupListener(ctx *pulumi.Context, args *LookupListenerArgs, opts ...pulumi.InvokeOption) (*LookupListenerResult, error) {
	var rv LookupListenerResult
	err := ctx.Invoke("aws:lb/getListener:getListener", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListener.
type LookupListenerArgs struct {
	// ARN of the listener. Required if `loadBalancerArn` and `port` is not set.
	Arn *string `pulumi:"arn"`
	// ARN of the load balancer. Required if `arn` is not set.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// Port of the listener. Required if `arn` is not set.
	Port *int              `pulumi:"port"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getListener.
type LookupListenerResult struct {
	AlpnPolicy     string                     `pulumi:"alpnPolicy"`
	Arn            string                     `pulumi:"arn"`
	CertificateArn string                     `pulumi:"certificateArn"`
	DefaultActions []GetListenerDefaultAction `pulumi:"defaultActions"`
	// The provider-assigned unique ID for this managed resource.
	Id              string            `pulumi:"id"`
	LoadBalancerArn string            `pulumi:"loadBalancerArn"`
	Port            int               `pulumi:"port"`
	Protocol        string            `pulumi:"protocol"`
	SslPolicy       string            `pulumi:"sslPolicy"`
	Tags            map[string]string `pulumi:"tags"`
}

func LookupListenerApply(ctx *pulumi.Context, args LookupListenerApplyInput, opts ...pulumi.InvokeOption) LookupListenerResultOutput {
	return args.ToLookupListenerApplyOutput().ApplyT(func(v LookupListenerArgs) (LookupListenerResult, error) {
		r, err := LookupListener(ctx, &v, opts...)
		return *r, err

	}).(LookupListenerResultOutput)
}

// LookupListenerApplyInput is an input type that accepts LookupListenerApplyArgs and LookupListenerApplyOutput values.
// You can construct a concrete instance of `LookupListenerApplyInput` via:
//
//          LookupListenerApplyArgs{...}
type LookupListenerApplyInput interface {
	pulumi.Input

	ToLookupListenerApplyOutput() LookupListenerApplyOutput
	ToLookupListenerApplyOutputWithContext(context.Context) LookupListenerApplyOutput
}

// A collection of arguments for invoking getListener.
type LookupListenerApplyArgs struct {
	// ARN of the listener. Required if `loadBalancerArn` and `port` is not set.
	Arn pulumi.StringPtrInput `pulumi:"arn"`
	// ARN of the load balancer. Required if `arn` is not set.
	LoadBalancerArn pulumi.StringPtrInput `pulumi:"loadBalancerArn"`
	// Port of the listener. Required if `arn` is not set.
	Port pulumi.IntPtrInput    `pulumi:"port"`
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupListenerApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerArgs)(nil)).Elem()
}

func (i LookupListenerApplyArgs) ToLookupListenerApplyOutput() LookupListenerApplyOutput {
	return i.ToLookupListenerApplyOutputWithContext(context.Background())
}

func (i LookupListenerApplyArgs) ToLookupListenerApplyOutputWithContext(ctx context.Context) LookupListenerApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupListenerApplyOutput)
}

// A collection of arguments for invoking getListener.
type LookupListenerApplyOutput struct{ *pulumi.OutputState }

func (LookupListenerApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerArgs)(nil)).Elem()
}

func (o LookupListenerApplyOutput) ToLookupListenerApplyOutput() LookupListenerApplyOutput {
	return o
}

func (o LookupListenerApplyOutput) ToLookupListenerApplyOutputWithContext(ctx context.Context) LookupListenerApplyOutput {
	return o
}

// ARN of the listener. Required if `loadBalancerArn` and `port` is not set.
func (o LookupListenerApplyOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupListenerArgs) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// ARN of the load balancer. Required if `arn` is not set.
func (o LookupListenerApplyOutput) LoadBalancerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupListenerArgs) *string { return v.LoadBalancerArn }).(pulumi.StringPtrOutput)
}

// Port of the listener. Required if `arn` is not set.
func (o LookupListenerApplyOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupListenerArgs) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupListenerApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupListenerArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getListener.
type LookupListenerResultOutput struct{ *pulumi.OutputState }

func (LookupListenerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerResult)(nil)).Elem()
}

func (o LookupListenerResultOutput) ToLookupListenerResultOutput() LookupListenerResultOutput {
	return o
}

func (o LookupListenerResultOutput) ToLookupListenerResultOutputWithContext(ctx context.Context) LookupListenerResultOutput {
	return o
}

func (o LookupListenerResultOutput) AlpnPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.AlpnPolicy }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) CertificateArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.CertificateArn }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) DefaultActions() GetListenerDefaultActionArrayOutput {
	return o.ApplyT(func(v LookupListenerResult) []GetListenerDefaultAction { return v.DefaultActions }).(GetListenerDefaultActionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupListenerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) LoadBalancerArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.LoadBalancerArn }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupListenerResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupListenerResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) SslPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupListenerResult) string { return v.SslPolicy }).(pulumi.StringOutput)
}

func (o LookupListenerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupListenerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupListenerApplyOutput{})
	pulumi.RegisterOutputType(LookupListenerResultOutput{})
}

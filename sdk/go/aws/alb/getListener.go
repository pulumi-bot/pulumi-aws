// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
//
// Provides information about a Load Balancer Listener.
//
// This data source can prove useful when a module accepts an LB Listener as an
// input variable and needs to know the LB it is attached to, or other
// information specific to the listener in question.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
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
// 		_, err := lb.LookupListener(ctx, &lb.LookupListenerArgs{
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
	err := ctx.Invoke("aws:alb/getListener:getListener", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListener.
type LookupListenerArgs struct {
	// The arn of the listener. Required if `loadBalancerArn` and `port` is not set.
	Arn *string `pulumi:"arn"`
	// The arn of the load balancer. Required if `arn` is not set.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// The port of the listener. Required if `arn` is not set.
	Port *int `pulumi:"port"`
}

// A collection of values returned by getListener.
type LookupListenerResult struct {
	Arn            string                     `pulumi:"arn"`
	CertificateArn string                     `pulumi:"certificateArn"`
	DefaultActions []GetListenerDefaultAction `pulumi:"defaultActions"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	Port            int    `pulumi:"port"`
	Protocol        string `pulumi:"protocol"`
	SslPolicy       string `pulumi:"sslPolicy"`
}

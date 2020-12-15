// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.
//
// Provides information about a Load Balancer.
//
// This data source can prove useful when a module accepts an LB as an input
// variable and needs to, for example, determine the security groups associated
// with it, etc.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		lbArn := ""
// 		if param := cfg.Get("lbArn"); param != "" {
// 			lbArn = param
// 		}
// 		lbName := ""
// 		if param := cfg.Get("lbName"); param != "" {
// 			lbName = param
// 		}
// 		opt0 := lbArn
// 		opt1 := lbName
// 		_, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
// 			Arn:  _opt0,
// 			Name: _opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.elasticloadbalancingv2.getLoadBalancer has been deprecated in favor of aws.lb.getLoadBalancer
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws:elasticloadbalancingv2/getLoadBalancer:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerArgs struct {
	// The full ARN of the load balancer.
	Arn *string `pulumi:"arn"`
	// The unique name of the load balancer.
	Name *string           `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLoadBalancer.
type LookupLoadBalancerResult struct {
	AccessLogs               GetLoadBalancerAccessLogs `pulumi:"accessLogs"`
	Arn                      string                    `pulumi:"arn"`
	ArnSuffix                string                    `pulumi:"arnSuffix"`
	CustomerOwnedIpv4Pool    string                    `pulumi:"customerOwnedIpv4Pool"`
	DnsName                  string                    `pulumi:"dnsName"`
	DropInvalidHeaderFields  bool                      `pulumi:"dropInvalidHeaderFields"`
	EnableDeletionProtection bool                      `pulumi:"enableDeletionProtection"`
	EnableHttp2              bool                      `pulumi:"enableHttp2"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	IdleTimeout      int                            `pulumi:"idleTimeout"`
	Internal         bool                           `pulumi:"internal"`
	IpAddressType    string                         `pulumi:"ipAddressType"`
	LoadBalancerType string                         `pulumi:"loadBalancerType"`
	Name             string                         `pulumi:"name"`
	SecurityGroups   []string                       `pulumi:"securityGroups"`
	SubnetMappings   []GetLoadBalancerSubnetMapping `pulumi:"subnetMappings"`
	Subnets          []string                       `pulumi:"subnets"`
	Tags             map[string]string              `pulumi:"tags"`
	VpcId            string                         `pulumi:"vpcId"`
	ZoneId           string                         `pulumi:"zoneId"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about a "classic" Elastic Load Balancer (ELB).
// See [LB Data Source](https://www.terraform.io/docs/providers/aws/d/lb.html) if you are looking for "v2"
// Application Load Balancer (ALB) or Network Load Balancer (NLB).
//
// This data source can prove useful when a module accepts an LB as an input
// variable and needs to, for example, determine the security groups associated
// with it, etc.
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws:elasticloadbalancing/getLoadBalancer:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerArgs struct {
	// The unique name of the load balancer.
	Name string                 `pulumi:"name"`
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getLoadBalancer.
type LookupLoadBalancerResult struct {
	AccessLogs                GetLoadBalancerAccessLogs  `pulumi:"accessLogs"`
	Arn                       string                     `pulumi:"arn"`
	AvailabilityZones         []string                   `pulumi:"availabilityZones"`
	ConnectionDraining        bool                       `pulumi:"connectionDraining"`
	ConnectionDrainingTimeout int                        `pulumi:"connectionDrainingTimeout"`
	CrossZoneLoadBalancing    bool                       `pulumi:"crossZoneLoadBalancing"`
	DnsName                   string                     `pulumi:"dnsName"`
	HealthCheck               GetLoadBalancerHealthCheck `pulumi:"healthCheck"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string                    `pulumi:"id"`
	IdleTimeout           int                       `pulumi:"idleTimeout"`
	Instances             []string                  `pulumi:"instances"`
	Internal              bool                      `pulumi:"internal"`
	Listeners             []GetLoadBalancerListener `pulumi:"listeners"`
	Name                  string                    `pulumi:"name"`
	SecurityGroups        []string                  `pulumi:"securityGroups"`
	SourceSecurityGroup   string                    `pulumi:"sourceSecurityGroup"`
	SourceSecurityGroupId string                    `pulumi:"sourceSecurityGroupId"`
	Subnets               []string                  `pulumi:"subnets"`
	Tags                  map[string]interface{}    `pulumi:"tags"`
	ZoneId                string                    `pulumi:"zoneId"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws:elb/getLoadBalancer:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerArgs struct {
	Name string            `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
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
	Tags                  map[string]string         `pulumi:"tags"`
	ZoneId                string                    `pulumi:"zoneId"`
}

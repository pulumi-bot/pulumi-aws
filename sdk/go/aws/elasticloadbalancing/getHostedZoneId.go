// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
// in a given region for the purpose of using in an AWS Route53 Alias.
//
//
//
// Deprecated: aws.elasticloadbalancing.getHostedZoneId has been deprecated in favour of aws.elb.getHostedZoneId
//
// aws.elasticloadbalancing.getHostedZoneId has been deprecated in favour of aws.elb.getHostedZoneId
func GetHostedZoneId(ctx *pulumi.Context, args *GetHostedZoneIdArgs, opts ...pulumi.InvokeOption) (*GetHostedZoneIdResult, error) {
	var rv GetHostedZoneIdResult
	err := ctx.Invoke("aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostedZoneId.
type GetHostedZoneIdArgs struct {
	// Name of the region whose AWS ELB HostedZoneId is desired.
	// Defaults to the region from the AWS provider configuration.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getHostedZoneId.
type GetHostedZoneIdResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
}

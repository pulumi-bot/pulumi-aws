// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
// in a given region for the purpose of using in an AWS Route53 Alias.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_hosted_zone_id.html.markdown.
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
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Region *string `pulumi:"region"`
}


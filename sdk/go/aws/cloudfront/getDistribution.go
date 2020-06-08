// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about a CloudFront distribution.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := cloudfront.LookupDistribution(ctx, &cloudfront.LookupDistributionArgs{
// 			Id: "EDFDVBD632BHDS5",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDistribution(ctx *pulumi.Context, args *LookupDistributionArgs, opts ...pulumi.InvokeOption) (*LookupDistributionResult, error) {
	var rv LookupDistributionResult
	err := ctx.Invoke("aws:cloudfront/getDistribution:getDistribution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDistribution.
type LookupDistributionArgs struct {
	// The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
	Id   string                 `pulumi:"id"`
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getDistribution.
type LookupDistributionResult struct {
	// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
	Arn string `pulumi:"arn"`
	// The domain name corresponding to the distribution. For
	// example: `d604721fxaaqy9.cloudfront.net`.
	DomainName string `pulumi:"domainName"`
	Enabled    bool   `pulumi:"enabled"`
	// The current version of the distribution's information. For example:
	// `E2QWRUHAPOMQZL`.
	Etag string `pulumi:"etag"`
	// The CloudFront Route 53 zone ID that can be used to
	// route an [Alias Resource Record Set][7] to. This attribute is simply an
	// alias for the zone ID `Z2FDTNDATAQYW2`.
	HostedZoneId string `pulumi:"hostedZoneId"`
	// The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
	Id string `pulumi:"id"`
	// The number of invalidation batches
	// currently in progress.
	InProgressValidationBatches int `pulumi:"inProgressValidationBatches"`
	// The date and time the distribution was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The current status of the distribution. `Deployed` if the
	// distribution's information is fully propagated throughout the Amazon
	// CloudFront system.
	Status string                 `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupBucket(ctx *pulumi.Context, args *LookupBucketArgs, opts ...pulumi.InvokeOption) (*LookupBucketResult, error) {
	var rv LookupBucketResult
	err := ctx.Invoke("aws:s3/getBucket:getBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucket.
type LookupBucketArgs struct {
	Bucket string `pulumi:"bucket"`
}

// A collection of values returned by getBucket.
type LookupBucketResult struct {
	Arn                      string `pulumi:"arn"`
	Bucket                   string `pulumi:"bucket"`
	BucketDomainName         string `pulumi:"bucketDomainName"`
	BucketRegionalDomainName string `pulumi:"bucketRegionalDomainName"`
	HostedZoneId             string `pulumi:"hostedZoneId"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	Region          string `pulumi:"region"`
	WebsiteDomain   string `pulumi:"websiteDomain"`
	WebsiteEndpoint string `pulumi:"websiteEndpoint"`
}

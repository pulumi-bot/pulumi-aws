// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides details about a specific S3 bucket.
//
// This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
// Distribution.
//
// ## Example Usage
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
	// The name of the bucket
	Bucket string `pulumi:"bucket"`
}

// A collection of values returned by getBucket.
type LookupBucketResult struct {
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn    string `pulumi:"arn"`
	Bucket string `pulumi:"bucket"`
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName string `pulumi:"bucketDomainName"`
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName string `pulumi:"bucketRegionalDomainName"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId string `pulumi:"hostedZoneId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The AWS region this bucket resides in.
	Region string `pulumi:"region"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint string `pulumi:"websiteEndpoint"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    public static class GetBucket
    {
        /// <summary>
        /// Provides details about a specific S3 bucket.
        /// 
        /// This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
        /// Distribution.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Route53 Record
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var selected = Output.Create(Aws.S3.GetBucket.InvokeAsync(new Aws.S3.GetBucketArgs
        ///         {
        ///             Bucket = "bucket.test.com",
        ///         }));
        ///         var testZone = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        ///         {
        ///             Name = "test.com.",
        ///         }));
        ///         var example = new Aws.Route53.Record("example", new Aws.Route53.RecordArgs
        ///         {
        ///             ZoneId = testZone.Apply(testZone =&gt; testZone.Id),
        ///             Name = "bucket",
        ///             Type = "A",
        ///             Aliases = 
        ///             {
        ///                 new Aws.Route53.Inputs.RecordAliasArgs
        ///                 {
        ///                     Name = selected.Apply(selected =&gt; selected.WebsiteDomain),
        ///                     ZoneId = selected.Apply(selected =&gt; selected.HostedZoneId),
        ///                 },
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### CloudFront Origin
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var selected = Output.Create(Aws.S3.GetBucket.InvokeAsync(new Aws.S3.GetBucketArgs
        ///         {
        ///             Bucket = "a-test-bucket",
        ///         }));
        ///         var test = new Aws.CloudFront.Distribution("test", new Aws.CloudFront.DistributionArgs
        ///         {
        ///             Origins = 
        ///             {
        ///                 new Aws.CloudFront.Inputs.DistributionOriginArgs
        ///                 {
        ///                     DomainName = selected.Apply(selected =&gt; selected.BucketDomainName),
        ///                     OriginId = "s3-selected-bucket",
        ///                 },
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketResult> InvokeAsync(GetBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketResult>("aws:s3/getBucket:getBucket", args ?? new GetBucketArgs(), options.WithVersion());

        public static Output<GetBucketResult> Invoke(GetBucketOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Bucket.Box()
            ).Apply(a => {
                    var args = new GetBucketArgs();
                    a[0].Set(args, nameof(args.Bucket));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetBucketArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the bucket
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        public GetBucketArgs()
        {
        }
    }

    public sealed class GetBucketOutputArgs
    {
        /// <summary>
        /// The name of the bucket
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        public GetBucketOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketResult
    {
        /// <summary>
        /// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
        /// </summary>
        public readonly string Arn;
        public readonly string Bucket;
        /// <summary>
        /// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
        /// </summary>
        public readonly string BucketDomainName;
        /// <summary>
        /// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
        /// </summary>
        public readonly string BucketRegionalDomainName;
        /// <summary>
        /// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
        /// </summary>
        public readonly string HostedZoneId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The AWS region this bucket resides in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
        /// </summary>
        public readonly string WebsiteDomain;
        /// <summary>
        /// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
        /// </summary>
        public readonly string WebsiteEndpoint;

        [OutputConstructor]
        private GetBucketResult(
            string arn,

            string bucket,

            string bucketDomainName,

            string bucketRegionalDomainName,

            string hostedZoneId,

            string id,

            string region,

            string websiteDomain,

            string websiteEndpoint)
        {
            Arn = arn;
            Bucket = bucket;
            BucketDomainName = bucketDomainName;
            BucketRegionalDomainName = bucketRegionalDomainName;
            HostedZoneId = hostedZoneId;
            Id = id;
            Region = region;
            WebsiteDomain = websiteDomain;
            WebsiteEndpoint = websiteEndpoint;
        }
    }
}

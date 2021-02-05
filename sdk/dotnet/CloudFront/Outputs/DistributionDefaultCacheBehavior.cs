// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Outputs
{

    [OutputType]
    public sealed class DistributionDefaultCacheBehavior
    {
        /// <summary>
        /// Controls which HTTP methods CloudFront
        /// processes and forwards to your Amazon S3 bucket or your custom origin.
        /// </summary>
        public readonly ImmutableArray<string> AllowedMethods;
        /// <summary>
        /// Controls whether CloudFront caches the
        /// response to requests using the specified HTTP methods.
        /// </summary>
        public readonly ImmutableArray<string> CachedMethods;
        /// <summary>
        /// Whether you want CloudFront to automatically
        /// compress content for web requests that include `Accept-Encoding: gzip` in
        /// the request header (default: `false`).
        /// </summary>
        public readonly bool? Compress;
        /// <summary>
        /// The default amount of time (in seconds) that an
        /// object is in a CloudFront cache before CloudFront forwards another request
        /// in the absence of an `Cache-Control max-age` or `Expires` header. Defaults to
        /// 1 day.
        /// </summary>
        public readonly int? DefaultTtl;
        /// <summary>
        /// Field level encryption configuration ID
        /// </summary>
        public readonly string? FieldLevelEncryptionId;
        /// <summary>
        /// The forwarded values configuration that specifies how CloudFront
        /// handles query strings, cookies and headers (maximum one).
        /// </summary>
        public readonly Outputs.DistributionDefaultCacheBehaviorForwardedValues ForwardedValues;
        /// <summary>
        /// A config block that triggers a lambda function with
        /// specific actions. Defined below, maximum 4.
        /// </summary>
        public readonly ImmutableArray<Outputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociation> LambdaFunctionAssociations;
        /// <summary>
        /// The maximum amount of time (in seconds) that an
        /// object is in a CloudFront cache before CloudFront forwards another request
        /// to your origin to determine whether the object has been updated. Only
        /// effective in the presence of `Cache-Control max-age`, `Cache-Control
        /// s-maxage`, and `Expires` headers. Defaults to 365 days.
        /// </summary>
        public readonly int? MaxTtl;
        /// <summary>
        /// The minimum amount of time that you want objects to
        /// stay in CloudFront caches before CloudFront queries your origin to see
        /// whether the object has been updated. Defaults to 0 seconds.
        /// </summary>
        public readonly int? MinTtl;
        public readonly string? OriginRequestPolicyId;
        /// <summary>
        /// Indicates whether you want to distribute
        /// media files in Microsoft Smooth Streaming format using the origin that is
        /// associated with this cache behavior.
        /// </summary>
        public readonly bool? SmoothStreaming;
        /// <summary>
        /// The value of ID for the origin that you want
        /// CloudFront to route requests to when a request matches the path pattern
        /// either for a cache behavior or for the default cache behavior.
        /// </summary>
        public readonly string TargetOriginId;
        /// <summary>
        /// List of AWS account IDs (or `self`) that you want to allow to create signed URLs for private content.
        /// See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
        /// </summary>
        public readonly ImmutableArray<string> TrustedSigners;
        /// <summary>
        /// Use this element to specify the
        /// protocol that users can use to access the files in the origin specified by
        /// TargetOriginId when a request matches the path pattern in PathPattern. One
        /// of `allow-all`, `https-only`, or `redirect-to-https`.
        /// </summary>
        public readonly string ViewerProtocolPolicy;

        [OutputConstructor]
        private DistributionDefaultCacheBehavior(
            ImmutableArray<string> allowedMethods,

            ImmutableArray<string> cachedMethods,

            bool? compress,

            int? defaultTtl,

            string? fieldLevelEncryptionId,

            Outputs.DistributionDefaultCacheBehaviorForwardedValues forwardedValues,

            ImmutableArray<Outputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociation> lambdaFunctionAssociations,

            int? maxTtl,

            int? minTtl,

            string? originRequestPolicyId,

            bool? smoothStreaming,

            string targetOriginId,

            ImmutableArray<string> trustedSigners,

            string viewerProtocolPolicy)
        {
            AllowedMethods = allowedMethods;
            CachedMethods = cachedMethods;
            Compress = compress;
            DefaultTtl = defaultTtl;
            FieldLevelEncryptionId = fieldLevelEncryptionId;
            ForwardedValues = forwardedValues;
            LambdaFunctionAssociations = lambdaFunctionAssociations;
            MaxTtl = maxTtl;
            MinTtl = minTtl;
            OriginRequestPolicyId = originRequestPolicyId;
            SmoothStreaming = smoothStreaming;
            TargetOriginId = targetOriginId;
            TrustedSigners = trustedSigners;
            ViewerProtocolPolicy = viewerProtocolPolicy;
        }
    }
}

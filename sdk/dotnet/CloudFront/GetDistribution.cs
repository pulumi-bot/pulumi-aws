// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront
{
    public static class GetDistribution
    {
        /// <summary>
        /// Use this data source to retrieve information about a CloudFront distribution.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.CloudFront.GetDistribution.InvokeAsync(new Aws.CloudFront.GetDistributionArgs
        ///         {
        ///             Id = "EDFDVBD632BHDS5",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDistributionResult> InvokeAsync(GetDistributionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDistributionResult>("aws:cloudfront/getDistribution:getDistribution", args ?? new GetDistributionArgs(), options.WithVersion());

        public static Output<GetDistributionResult> Invoke(GetDistributionOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Id.Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetDistributionArgs();
                    a[0].Set(args, nameof(args.Id));
                    a[1].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetDistributionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetDistributionArgs()
        {
        }
    }

    public sealed class GetDistributionOutputArgs
    {
        /// <summary>
        /// The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetDistributionOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDistributionResult
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The domain name corresponding to the distribution. For
        /// example: `d604721fxaaqy9.cloudfront.net`.
        /// </summary>
        public readonly string DomainName;
        public readonly bool Enabled;
        /// <summary>
        /// The current version of the distribution's information. For example:
        /// `E2QWRUHAPOMQZL`.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The CloudFront Route 53 zone ID that can be used to
        /// route an [Alias Resource Record Set][7] to. This attribute is simply an
        /// alias for the zone ID `Z2FDTNDATAQYW2`.
        /// </summary>
        public readonly string HostedZoneId;
        /// <summary>
        /// The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The number of invalidation batches
        /// currently in progress.
        /// </summary>
        public readonly int InProgressValidationBatches;
        /// <summary>
        /// The date and time the distribution was last modified.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The current status of the distribution. `Deployed` if the
        /// distribution's information is fully propagated throughout the Amazon
        /// CloudFront system.
        /// </summary>
        public readonly string Status;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetDistributionResult(
            string arn,

            string domainName,

            bool enabled,

            string etag,

            string hostedZoneId,

            string id,

            int inProgressValidationBatches,

            string lastModifiedTime,

            string status,

            ImmutableDictionary<string, string>? tags)
        {
            Arn = arn;
            DomainName = domainName;
            Enabled = enabled;
            Etag = etag;
            HostedZoneId = hostedZoneId;
            Id = id;
            InProgressValidationBatches = inProgressValidationBatches;
            LastModifiedTime = lastModifiedTime;
            Status = status;
            Tags = tags;
        }
    }
}

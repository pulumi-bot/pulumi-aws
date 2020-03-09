// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Elb
{
    public static partial class GetServiceAccount
    {
        /// <summary>
        /// Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
        /// in a given region for the purpose of whitelisting in S3 bucket policy.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_service_account.html.markdown.
        /// </summary>
        public static Task<GetServiceAccountResult> InvokeAsync(GetServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceAccountResult>("aws:elb/getServiceAccount:getServiceAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetServiceAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the region whose AWS ELB account ID is desired.
        /// Defaults to the region from the AWS provider configuration.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetServiceAccountArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetServiceAccountResult
    {
        /// <summary>
        /// The ARN of the AWS ELB service account in the selected region.
        /// </summary>
        public readonly string Arn;
        public readonly string? Region;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetServiceAccountResult(
            string arn,
            string? region,
            string id)
        {
            Arn = arn;
            Region = region;
            Id = id;
        }
    }
}

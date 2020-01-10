// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudTrail
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the Account ID of the [AWS CloudTrail Service Account](http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-supported-regions.html)
        /// in a given region for the purpose of allowing CloudTrail to store trail data in S3.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudtrail_service_account.html.markdown.
        /// </summary>
        public static Task<GetServiceAccountResult> GetServiceAccount(GetServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceAccountResult>("aws:cloudtrail/getServiceAccount:getServiceAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetServiceAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the region whose AWS CloudTrail account ID is desired.
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
        /// The ARN of the AWS CloudTrail service account in the selected region.
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

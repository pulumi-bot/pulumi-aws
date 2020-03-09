// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cur
{
    public static partial class GetReportDefinition
    {
        /// <summary>
        /// Use this data source to get information on an AWS Cost and Usage Report Definition.
        /// 
        /// &gt; *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.
        /// 
        /// &gt; *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cur_report_definition.html.markdown.
        /// </summary>
        public static Task<GetReportDefinitionResult> InvokeAsync(GetReportDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReportDefinitionResult>("aws:cur/getReportDefinition:getReportDefinition", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetReportDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the report definition to match.
        /// </summary>
        [Input("reportName", required: true)]
        public string ReportName { get; set; } = null!;

        public GetReportDefinitionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetReportDefinitionResult
    {
        /// <summary>
        /// A list of additional artifacts.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalArtifacts;
        /// <summary>
        /// A list of schema elements.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalSchemaElements;
        /// <summary>
        /// Preferred format for report.
        /// </summary>
        public readonly string Compression;
        /// <summary>
        /// Preferred compression format for report.
        /// </summary>
        public readonly string Format;
        public readonly string ReportName;
        /// <summary>
        /// Name of customer S3 bucket.
        /// </summary>
        public readonly string S3Bucket;
        /// <summary>
        /// Preferred report path prefix.
        /// </summary>
        public readonly string S3Prefix;
        /// <summary>
        /// Region of customer S3 bucket.
        /// </summary>
        public readonly string S3Region;
        /// <summary>
        /// The frequency on which report data are measured and displayed.
        /// </summary>
        public readonly string TimeUnit;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetReportDefinitionResult(
            ImmutableArray<string> additionalArtifacts,
            ImmutableArray<string> additionalSchemaElements,
            string compression,
            string format,
            string reportName,
            string s3Bucket,
            string s3Prefix,
            string s3Region,
            string timeUnit,
            string id)
        {
            AdditionalArtifacts = additionalArtifacts;
            AdditionalSchemaElements = additionalSchemaElements;
            Compression = compression;
            Format = format;
            ReportName = reportName;
            S3Bucket = s3Bucket;
            S3Prefix = s3Prefix;
            S3Region = s3Region;
            TimeUnit = timeUnit;
            Id = id;
        }
    }
}

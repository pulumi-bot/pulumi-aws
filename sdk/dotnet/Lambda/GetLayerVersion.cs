// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about a Lambda Layer Version.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_layer_version.html.markdown.
        /// </summary>
        public static Task<GetLayerVersionResult> GetLayerVersion(GetLayerVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLayerVersionResult>("aws:lambda/getLayerVersion:getLayerVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetLayerVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specific runtime the layer version must support. Conflicts with `version`. If specified, the latest available layer version supporting the provided runtime will be used.
        /// </summary>
        [Input("compatibleRuntime")]
        public string? CompatibleRuntime { get; set; }

        /// <summary>
        /// Name of the lambda layer.
        /// </summary>
        [Input("layerName", required: true)]
        public string LayerName { get; set; } = null!;

        /// <summary>
        /// Specific layer version. Conflicts with `compatible_runtime`. If omitted, the latest available layer version will be used.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetLayerVersionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetLayerVersionResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lambda Layer with version.
        /// </summary>
        public readonly string Arn;
        public readonly string? CompatibleRuntime;
        /// <summary>
        /// A list of [Runtimes][1] the specific Lambda Layer version is compatible with.
        /// </summary>
        public readonly ImmutableArray<string> CompatibleRuntimes;
        /// <summary>
        /// The date this resource was created.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// Description of the specific Lambda Layer version.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lambda Layer without version.
        /// </summary>
        public readonly string LayerArn;
        public readonly string LayerName;
        /// <summary>
        /// License info associated with the specific Lambda Layer version.
        /// </summary>
        public readonly string LicenseInfo;
        /// <summary>
        /// Base64-encoded representation of raw SHA-256 sum of the zip file.
        /// </summary>
        public readonly string SourceCodeHash;
        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        public readonly int SourceCodeSize;
        /// <summary>
        /// This Lamba Layer version.
        /// </summary>
        public readonly int Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLayerVersionResult(
            string arn,
            string? compatibleRuntime,
            ImmutableArray<string> compatibleRuntimes,
            string createdDate,
            string description,
            string layerArn,
            string layerName,
            string licenseInfo,
            string sourceCodeHash,
            int sourceCodeSize,
            int version,
            string id)
        {
            Arn = arn;
            CompatibleRuntime = compatibleRuntime;
            CompatibleRuntimes = compatibleRuntimes;
            CreatedDate = createdDate;
            Description = description;
            LayerArn = layerArn;
            LayerName = layerName;
            LicenseInfo = licenseInfo;
            SourceCodeHash = sourceCodeHash;
            SourceCodeSize = sourceCodeSize;
            Version = version;
            Id = id;
        }
    }
}

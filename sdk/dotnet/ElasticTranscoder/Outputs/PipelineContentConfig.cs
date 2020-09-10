// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticTranscoder.Outputs
{

    [OutputType]
    public sealed class PipelineContentConfig
    {
        public readonly string? Bucket;
        public readonly string? StorageClass;

        [OutputConstructor]
        private PipelineContentConfig(
            string? bucket,

            string? storageClass)
        {
            Bucket = bucket;
            StorageClass = storageClass;
        }
    }
}

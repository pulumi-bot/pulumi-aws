// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticTranscoder.Inputs
{

    public sealed class PipelineThumbnailConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        public PipelineThumbnailConfigGetArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm.Inputs
{

    public sealed class ResourceDataSyncS3DestinationGetArgs : Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("syncFormat")]
        public Input<string>? SyncFormat { get; set; }

        public ResourceDataSyncS3DestinationGetArgs()
        {
        }
    }
}

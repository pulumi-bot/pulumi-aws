// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class SpotInstanceRequestEbsBlockDeviceArgs : Pulumi.ResourceArgs
    {
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public SpotInstanceRequestEbsBlockDeviceArgs()
        {
        }
    }
}

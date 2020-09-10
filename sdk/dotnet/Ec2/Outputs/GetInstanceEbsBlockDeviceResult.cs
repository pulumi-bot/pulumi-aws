// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class GetInstanceEbsBlockDeviceResult
    {
        public readonly bool DeleteOnTermination;
        public readonly string DeviceName;
        public readonly bool Encrypted;
        public readonly int Iops;
        public readonly string KmsKeyId;
        public readonly string SnapshotId;
        public readonly string VolumeId;
        public readonly int VolumeSize;
        public readonly string VolumeType;

        [OutputConstructor]
        private GetInstanceEbsBlockDeviceResult(
            bool deleteOnTermination,

            string deviceName,

            bool encrypted,

            int iops,

            string kmsKeyId,

            string snapshotId,

            string volumeId,

            int volumeSize,

            string volumeType)
        {
            DeleteOnTermination = deleteOnTermination;
            DeviceName = deviceName;
            Encrypted = encrypted;
            Iops = iops;
            KmsKeyId = kmsKeyId;
            SnapshotId = snapshotId;
            VolumeId = volumeId;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }
}

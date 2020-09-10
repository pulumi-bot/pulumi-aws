// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Outputs
{

    [OutputType]
    public sealed class InventoryDestinationBucket
    {
        public readonly string? AccountId;
        public readonly string BucketArn;
        public readonly Outputs.InventoryDestinationBucketEncryption? Encryption;
        public readonly string Format;
        public readonly string? Prefix;

        [OutputConstructor]
        private InventoryDestinationBucket(
            string? accountId,

            string bucketArn,

            Outputs.InventoryDestinationBucketEncryption? encryption,

            string format,

            string? prefix)
        {
            AccountId = accountId;
            BucketArn = bucketArn;
            Encryption = encryption;
            Format = format;
            Prefix = prefix;
        }
    }
}

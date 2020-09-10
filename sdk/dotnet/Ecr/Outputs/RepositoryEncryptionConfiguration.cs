// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr.Outputs
{

    [OutputType]
    public sealed class RepositoryEncryptionConfiguration
    {
        public readonly string? EncryptionType;
        public readonly string? KmsKey;

        [OutputConstructor]
        private RepositoryEncryptionConfiguration(
            string? encryptionType,

            string? kmsKey)
        {
            EncryptionType = encryptionType;
            KmsKey = kmsKey;
        }
    }
}

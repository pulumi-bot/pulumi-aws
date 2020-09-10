// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr.Inputs
{

    public sealed class RepositoryEncryptionConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("encryptionType")]
        public Input<string>? EncryptionType { get; set; }

        [Input("kmsKey")]
        public Input<string>? KmsKey { get; set; }

        public RepositoryEncryptionConfigurationArgs()
        {
        }
    }
}

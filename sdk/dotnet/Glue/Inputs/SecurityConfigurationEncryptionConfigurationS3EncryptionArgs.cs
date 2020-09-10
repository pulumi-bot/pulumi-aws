// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class SecurityConfigurationEncryptionConfigurationS3EncryptionArgs : Pulumi.ResourceArgs
    {
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("s3EncryptionMode")]
        public Input<string>? S3EncryptionMode { get; set; }

        public SecurityConfigurationEncryptionConfigurationS3EncryptionArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolSoftwareTokenMfaConfigurationGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public UserPoolSoftwareTokenMfaConfigurationGetArgs()
        {
        }
    }
}

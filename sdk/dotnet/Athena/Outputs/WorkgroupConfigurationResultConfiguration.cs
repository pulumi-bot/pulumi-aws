// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Athena.Outputs
{

    [OutputType]
    public sealed class WorkgroupConfigurationResultConfiguration
    {
        public readonly Outputs.WorkgroupConfigurationResultConfigurationEncryptionConfiguration? EncryptionConfiguration;
        public readonly string? OutputLocation;

        [OutputConstructor]
        private WorkgroupConfigurationResultConfiguration(
            Outputs.WorkgroupConfigurationResultConfigurationEncryptionConfiguration? encryptionConfiguration,

            string? outputLocation)
        {
            EncryptionConfiguration = encryptionConfiguration;
            OutputLocation = outputLocation;
        }
    }
}

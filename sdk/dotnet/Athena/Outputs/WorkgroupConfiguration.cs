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
    public sealed class WorkgroupConfiguration
    {
        public readonly int? BytesScannedCutoffPerQuery;
        public readonly bool? EnforceWorkgroupConfiguration;
        public readonly bool? PublishCloudwatchMetricsEnabled;
        public readonly Outputs.WorkgroupConfigurationResultConfiguration? ResultConfiguration;

        [OutputConstructor]
        private WorkgroupConfiguration(
            int? bytesScannedCutoffPerQuery,

            bool? enforceWorkgroupConfiguration,

            bool? publishCloudwatchMetricsEnabled,

            Outputs.WorkgroupConfigurationResultConfiguration? resultConfiguration)
        {
            BytesScannedCutoffPerQuery = bytesScannedCutoffPerQuery;
            EnforceWorkgroupConfiguration = enforceWorkgroupConfiguration;
            PublishCloudwatchMetricsEnabled = publishCloudwatchMetricsEnabled;
            ResultConfiguration = resultConfiguration;
        }
    }
}

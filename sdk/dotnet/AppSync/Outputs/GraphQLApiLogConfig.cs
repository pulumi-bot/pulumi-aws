// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync.Outputs
{

    [OutputType]
    public sealed class GraphQLApiLogConfig
    {
        public readonly string CloudwatchLogsRoleArn;
        public readonly bool? ExcludeVerboseContent;
        public readonly string FieldLogLevel;

        [OutputConstructor]
        private GraphQLApiLogConfig(
            string cloudwatchLogsRoleArn,

            bool? excludeVerboseContent,

            string fieldLogLevel)
        {
            CloudwatchLogsRoleArn = cloudwatchLogsRoleArn;
            ExcludeVerboseContent = excludeVerboseContent;
            FieldLogLevel = fieldLogLevel;
        }
    }
}

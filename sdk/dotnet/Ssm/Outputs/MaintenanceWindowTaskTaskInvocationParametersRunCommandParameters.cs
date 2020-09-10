// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm.Outputs
{

    [OutputType]
    public sealed class MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
    {
        public readonly string? Comment;
        public readonly string? DocumentHash;
        public readonly string? DocumentHashType;
        public readonly Outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig? NotificationConfig;
        public readonly string? OutputS3Bucket;
        public readonly string? OutputS3KeyPrefix;
        public readonly ImmutableArray<Outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter> Parameters;
        public readonly string? ServiceRoleArn;
        public readonly int? TimeoutSeconds;

        [OutputConstructor]
        private MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters(
            string? comment,

            string? documentHash,

            string? documentHashType,

            Outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig? notificationConfig,

            string? outputS3Bucket,

            string? outputS3KeyPrefix,

            ImmutableArray<Outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter> parameters,

            string? serviceRoleArn,

            int? timeoutSeconds)
        {
            Comment = comment;
            DocumentHash = documentHash;
            DocumentHashType = documentHashType;
            NotificationConfig = notificationConfig;
            OutputS3Bucket = outputS3Bucket;
            OutputS3KeyPrefix = outputS3KeyPrefix;
            Parameters = parameters;
            ServiceRoleArn = serviceRoleArn;
            TimeoutSeconds = timeoutSeconds;
        }
    }
}

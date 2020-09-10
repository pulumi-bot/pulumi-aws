// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Outputs
{

    [OutputType]
    public sealed class EventTargetEcsTarget
    {
        public readonly string? Group;
        public readonly string? LaunchType;
        public readonly Outputs.EventTargetEcsTargetNetworkConfiguration? NetworkConfiguration;
        public readonly string? PlatformVersion;
        public readonly int? TaskCount;
        public readonly string TaskDefinitionArn;

        [OutputConstructor]
        private EventTargetEcsTarget(
            string? group,

            string? launchType,

            Outputs.EventTargetEcsTargetNetworkConfiguration? networkConfiguration,

            string? platformVersion,

            int? taskCount,

            string taskDefinitionArn)
        {
            Group = group;
            LaunchType = launchType;
            NetworkConfiguration = networkConfiguration;
            PlatformVersion = platformVersion;
            TaskCount = taskCount;
            TaskDefinitionArn = taskDefinitionArn;
        }
    }
}

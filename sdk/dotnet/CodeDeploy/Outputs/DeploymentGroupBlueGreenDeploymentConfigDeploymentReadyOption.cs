// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption
    {
        public readonly string? ActionOnTimeout;
        public readonly int? WaitTimeInMinutes;

        [OutputConstructor]
        private DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption(
            string? actionOnTimeout,

            int? waitTimeInMinutes)
        {
            ActionOnTimeout = actionOnTimeout;
            WaitTimeInMinutes = waitTimeInMinutes;
        }
    }
}

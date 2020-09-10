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
    public sealed class DeploymentConfigTrafficRoutingConfig
    {
        public readonly Outputs.DeploymentConfigTrafficRoutingConfigTimeBasedCanary? TimeBasedCanary;
        public readonly Outputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinear? TimeBasedLinear;
        public readonly string? Type;

        [OutputConstructor]
        private DeploymentConfigTrafficRoutingConfig(
            Outputs.DeploymentConfigTrafficRoutingConfigTimeBasedCanary? timeBasedCanary,

            Outputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinear? timeBasedLinear,

            string? type)
        {
            TimeBasedCanary = timeBasedCanary;
            TimeBasedLinear = timeBasedLinear;
            Type = type;
        }
    }
}

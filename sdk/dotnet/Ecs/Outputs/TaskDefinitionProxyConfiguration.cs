// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionProxyConfiguration
    {
        /// <summary>
        /// The name of the container that will serve as the App Mesh proxy.
        /// </summary>
        public readonly string ContainerName;
        /// <summary>
        /// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;
        /// <summary>
        /// The type of constraint. Use `memberOf` to restrict selection to a group of valid candidates.
        /// Note that `distinctInstance` is not supported in task definitions.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private TaskDefinitionProxyConfiguration(
            string containerName,

            ImmutableDictionary<string, string>? properties,

            string? type)
        {
            ContainerName = containerName;
            Properties = properties;
            Type = type;
        }
    }
}

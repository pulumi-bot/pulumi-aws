// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class ClusterDefaultCapacityProviderStrategyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
        /// </summary>
        [Input("base")]
        public Input<int>? Base { get; set; }

        /// <summary>
        /// The short name of the capacity provider.
        /// </summary>
        [Input("capacityProvider", required: true)]
        public Input<string> CapacityProvider { get; set; } = null!;

        /// <summary>
        /// The relative percentage of the total number of launched tasks that should use the specified capacity provider.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ClusterDefaultCapacityProviderStrategyGetArgs()
        {
        }
    }
}

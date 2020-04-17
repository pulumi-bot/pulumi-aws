// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Inputs
{

    public sealed class ClusterParameterGroupParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// "immediate" (default), or "pending-reboot". Some
        /// engines can't apply some parameters without a reboot, and you will need to
        /// specify "pending-reboot" here.
        /// </summary>
        [Input("applyMethod")]
        public Input<string>? ApplyMethod { get; set; }

        /// <summary>
        /// The name of the DB cluster parameter group. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the DB parameter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ClusterParameterGroupParameterArgs()
        {
        }
    }
}

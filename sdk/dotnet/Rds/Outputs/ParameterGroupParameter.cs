// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Outputs
{

    [OutputType]
    public sealed class ParameterGroupParameter
    {
        /// <summary>
        /// "immediate" (default), or "pending-reboot". Some
        /// engines can't apply some parameters without a reboot, and you will need to
        /// specify "pending-reboot" here.
        /// </summary>
        public readonly string? ApplyMethod;
        /// <summary>
        /// The name of the DB parameter group. If omitted, this provider will assign a random, unique name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the DB parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ParameterGroupParameter(
            string? applyMethod,

            string name,

            string value)
        {
            ApplyMethod = applyMethod;
            Name = name;
            Value = value;
        }
    }
}

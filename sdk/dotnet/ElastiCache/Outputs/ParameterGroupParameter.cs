// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache.Outputs
{

    [OutputType]
    public sealed class ParameterGroupParameter
    {
        /// <summary>
        /// The name of the ElastiCache parameter group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the ElastiCache parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ParameterGroupParameter(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}

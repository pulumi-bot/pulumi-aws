// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sfn.Outputs
{

    [OutputType]
    public sealed class StateMachineLoggingConfiguration
    {
        /// <summary>
        /// Determines whether execution data is included in your log. When set to `false`, data is excluded.
        /// </summary>
        public readonly bool? IncludeExecutionData;
        /// <summary>
        /// Defines which category of execution history events are logged. Valid values: `ALL`, `ERROR`, `FATAL`, `OFF`
        /// </summary>
        public readonly string? Level;
        /// <summary>
        /// Amazon Resource Name (ARN) of a CloudWatch log group. Make sure the State Machine has the correct IAM policies for logging. The ARN must end with `:*`
        /// </summary>
        public readonly string? LogDestination;

        [OutputConstructor]
        private StateMachineLoggingConfiguration(
            bool? includeExecutionData,

            string? level,

            string? logDestination)
        {
            IncludeExecutionData = includeExecutionData;
            Level = level;
            LogDestination = logDestination;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sfn.Inputs
{

    public sealed class StateMachineTracingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to `true`, AWS X-Ray tracing is enabled. Make sure the State Machine has the correct IAM policies for logging. See the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg/xray-iam.html) for details.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public StateMachineTracingConfigurationArgs()
        {
        }
    }
}

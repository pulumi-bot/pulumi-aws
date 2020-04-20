// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy.Inputs
{

    public sealed class DeploymentGroupTriggerConfigurationGetArgs : Pulumi.ResourceArgs
    {
        [Input("triggerEvents", required: true)]
        private InputList<string>? _triggerEvents;

        /// <summary>
        /// The event type or types for which notifications are triggered. Some values that are supported: `DeploymentStart`, `DeploymentSuccess`, `DeploymentFailure`, `DeploymentStop`, `DeploymentRollback`, `InstanceStart`, `InstanceSuccess`, `InstanceFailure`.  See [the CodeDeploy documentation](http://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-sns-event-notifications-create-trigger.html) for all possible values.
        /// </summary>
        public InputList<string> TriggerEvents
        {
            get => _triggerEvents ?? (_triggerEvents = new InputList<string>());
            set => _triggerEvents = value;
        }

        /// <summary>
        /// The name of the notification trigger.
        /// </summary>
        [Input("triggerName", required: true)]
        public Input<string> TriggerName { get; set; } = null!;

        /// <summary>
        /// The ARN of the SNS topic through which notifications are sent.
        /// </summary>
        [Input("triggerTargetArn", required: true)]
        public Input<string> TriggerTargetArn { get; set; } = null!;

        public DeploymentGroupTriggerConfigurationGetArgs()
        {
        }
    }
}

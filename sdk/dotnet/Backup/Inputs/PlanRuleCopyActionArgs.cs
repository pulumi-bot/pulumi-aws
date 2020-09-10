// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup.Inputs
{

    public sealed class PlanRuleCopyActionArgs : Pulumi.ResourceArgs
    {
        [Input("destinationVaultArn", required: true)]
        public Input<string> DestinationVaultArn { get; set; } = null!;

        [Input("lifecycle")]
        public Input<Inputs.PlanRuleCopyActionLifecycleArgs>? Lifecycle { get; set; }

        public PlanRuleCopyActionArgs()
        {
        }
    }
}

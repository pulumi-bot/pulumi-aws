// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs()
        {
        }
    }
}

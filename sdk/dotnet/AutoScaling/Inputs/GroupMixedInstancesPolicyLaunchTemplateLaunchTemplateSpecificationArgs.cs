// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs()
        {
        }
    }
}

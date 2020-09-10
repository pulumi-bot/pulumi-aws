// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class FleetLaunchTemplateConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("launchTemplateSpecification", required: true)]
        public Input<Inputs.FleetLaunchTemplateConfigLaunchTemplateSpecificationGetArgs> LaunchTemplateSpecification { get; set; } = null!;

        [Input("overrides")]
        private InputList<Inputs.FleetLaunchTemplateConfigOverrideGetArgs>? _overrides;
        public InputList<Inputs.FleetLaunchTemplateConfigOverrideGetArgs> Overrides
        {
            get => _overrides ?? (_overrides = new InputList<Inputs.FleetLaunchTemplateConfigOverrideGetArgs>());
            set => _overrides = value;
        }

        public FleetLaunchTemplateConfigGetArgs()
        {
        }
    }
}

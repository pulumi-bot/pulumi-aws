// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Batch.Inputs
{

    public sealed class ComputeEnvironmentComputeResourcesArgs : Pulumi.ResourceArgs
    {
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        [Input("bidPercentage")]
        public Input<int>? BidPercentage { get; set; }

        [Input("desiredVcpus")]
        public Input<int>? DesiredVcpus { get; set; }

        [Input("ec2KeyPair")]
        public Input<string>? Ec2KeyPair { get; set; }

        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("instanceRole", required: true)]
        public Input<string> InstanceRole { get; set; } = null!;

        [Input("instanceTypes", required: true)]
        private InputList<string>? _instanceTypes;
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        [Input("launchTemplate")]
        public Input<Inputs.ComputeEnvironmentComputeResourcesLaunchTemplateArgs>? LaunchTemplate { get; set; }

        [Input("maxVcpus", required: true)]
        public Input<int> MaxVcpus { get; set; } = null!;

        [Input("minVcpus", required: true)]
        public Input<int> MinVcpus { get; set; } = null!;

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("spotIamFleetRole")]
        public Input<string>? SpotIamFleetRole { get; set; }

        [Input("subnets", required: true)]
        private InputList<string>? _subnets;
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ComputeEnvironmentComputeResourcesArgs()
        {
        }
    }
}

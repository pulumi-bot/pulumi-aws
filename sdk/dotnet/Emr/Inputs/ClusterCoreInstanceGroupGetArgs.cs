// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Inputs
{

    public sealed class ClusterCoreInstanceGroupGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<string>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// Bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        /// </summary>
        [Input("bidPrice")]
        public Input<string>? BidPrice { get; set; }

        [Input("ebsConfigs")]
        private InputList<Inputs.ClusterCoreInstanceGroupEbsConfigGetArgs>? _ebsConfigs;

        /// <summary>
        /// Configuration block(s) for EBS volumes attached to each instance in the instance group. Detailed below.
        /// </summary>
        public InputList<Inputs.ClusterCoreInstanceGroupEbsConfigGetArgs> EbsConfigs
        {
            get => _ebsConfigs ?? (_ebsConfigs = new InputList<Inputs.ClusterCoreInstanceGroupEbsConfigGetArgs>());
            set => _ebsConfigs = value;
        }

        /// <summary>
        /// The ID of the EMR Cluster
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Target number of instances for the instance group. Must be 1 or 3. Defaults to 1. Launching with multiple master nodes is only supported in EMR version 5.23.0+, and requires this resource's `core_instance_group` to be configured. Public (Internet accessible) instances must be created in VPC subnets that have [map public IP on launch](https://www.terraform.io/docs/providers/aws/r/subnet.html#map_public_ip_on_launch) enabled. Termination protection is automatically enabled when launched with multiple master nodes and this provider must have the `termination_protection = false` configuration applied before destroying this resource.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// EC2 instance type for all instances in the instance group.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The name of the step.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ClusterCoreInstanceGroupGetArgs()
        {
        }
    }
}

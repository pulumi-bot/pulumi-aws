// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class ServiceLoadBalancerArgs : Pulumi.ResourceArgs
    {
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        [Input("containerPort", required: true)]
        public Input<int> ContainerPort { get; set; } = null!;

        [Input("elbName")]
        public Input<string>? ElbName { get; set; }

        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        public ServiceLoadBalancerArgs()
        {
        }
    }
}

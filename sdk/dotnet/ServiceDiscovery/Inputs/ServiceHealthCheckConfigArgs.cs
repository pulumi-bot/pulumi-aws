// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceDiscovery.Inputs
{

    public sealed class ServiceHealthCheckConfigArgs : Pulumi.ResourceArgs
    {
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        [Input("resourcePath")]
        public Input<string>? ResourcePath { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ServiceHealthCheckConfigArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class TaskDefinitionInferenceAcceleratorGetArgs : Pulumi.ResourceArgs
    {
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        [Input("deviceType", required: true)]
        public Input<string> DeviceType { get; set; } = null!;

        public TaskDefinitionInferenceAcceleratorGetArgs()
        {
        }
    }
}

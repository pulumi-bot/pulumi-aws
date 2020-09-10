// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualNodeSpecServiceDiscoveryAwsCloudMapGetArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public VirtualNodeSpecServiceDiscoveryAwsCloudMapGetArgs()
        {
        }
    }
}

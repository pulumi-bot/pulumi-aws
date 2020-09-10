// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class RouteSpecHttpRouteMatchArgs : Pulumi.ResourceArgs
    {
        [Input("headers")]
        private InputList<Inputs.RouteSpecHttpRouteMatchHeaderArgs>? _headers;
        public InputList<Inputs.RouteSpecHttpRouteMatchHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.RouteSpecHttpRouteMatchHeaderArgs>());
            set => _headers = value;
        }

        [Input("method")]
        public Input<string>? Method { get; set; }

        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        public RouteSpecHttpRouteMatchArgs()
        {
        }
    }
}

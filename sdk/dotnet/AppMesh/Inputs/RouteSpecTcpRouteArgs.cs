// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class RouteSpecTcpRouteArgs : Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<Inputs.RouteSpecTcpRouteActionArgs> Action { get; set; } = null!;

        public RouteSpecTcpRouteArgs()
        {
        }
    }
}

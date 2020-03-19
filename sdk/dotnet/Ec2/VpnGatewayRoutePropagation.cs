// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Requests automatic route propagation between a VPN gateway and a route table.
    /// 
    /// &gt; **Note:** This resource should not be used with a route table that has
    /// the `propagating_vgws` argument set. If that argument is set, any route
    /// propagation not explicitly listed in its value will be removed.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpn_gateway_route_propagation.html.markdown.
    /// </summary>
    public partial class VpnGatewayRoutePropagation : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the `aws.ec2.RouteTable` to propagate routes into.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// The id of the `aws.ec2.VpnGateway` to propagate routes from.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string> VpnGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a VpnGatewayRoutePropagation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnGatewayRoutePropagation(string name, VpnGatewayRoutePropagationArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VpnGatewayRoutePropagation(string name, Input<string> id, VpnGatewayRoutePropagationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpnGatewayRoutePropagation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnGatewayRoutePropagation Get(string name, Input<string> id, VpnGatewayRoutePropagationState? state = null, CustomResourceOptions? options = null)
        {
            return new VpnGatewayRoutePropagation(name, id, state, options);
        }
    }

    public sealed class VpnGatewayRoutePropagationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the `aws.ec2.RouteTable` to propagate routes into.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        /// <summary>
        /// The id of the `aws.ec2.VpnGateway` to propagate routes from.
        /// </summary>
        [Input("vpnGatewayId", required: true)]
        public Input<string> VpnGatewayId { get; set; } = null!;

        public VpnGatewayRoutePropagationArgs()
        {
        }
    }

    public sealed class VpnGatewayRoutePropagationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the `aws.ec2.RouteTable` to propagate routes into.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// The id of the `aws.ec2.VpnGateway` to propagate routes from.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public VpnGatewayRoutePropagationState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2ClientVpn
{
    /// <summary>
    /// Provides additional routes for AWS Client VPN endpoints. For more information on usage, please see the
    /// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEndpoint = new Aws.Ec2ClientVpn.Endpoint("exampleEndpoint", new Aws.Ec2ClientVpn.EndpointArgs
    ///         {
    ///             Description = "Example Client VPN endpoint",
    ///             ServerCertificateArn = aws_acm_certificate.Example.Arn,
    ///             ClientCidrBlock = "10.0.0.0/16",
    ///             AuthenticationOptions = 
    ///             {
    ///                 new Aws.Ec2ClientVpn.Inputs.EndpointAuthenticationOptionArgs
    ///                 {
    ///                     Type = "certificate-authentication",
    ///                     RootCertificateChainArn = aws_acm_certificate.Example.Arn,
    ///                 },
    ///             },
    ///             ConnectionLogOptions = new Aws.Ec2ClientVpn.Inputs.EndpointConnectionLogOptionsArgs
    ///             {
    ///                 Enabled = false,
    ///             },
    ///         });
    ///         var exampleNetworkAssociation = new Aws.Ec2ClientVpn.NetworkAssociation("exampleNetworkAssociation", new Aws.Ec2ClientVpn.NetworkAssociationArgs
    ///         {
    ///             ClientVpnEndpointId = exampleEndpoint.Id,
    ///             SubnetId = aws_subnet.Example.Id,
    ///         });
    ///         var exampleRoute = new Aws.Ec2ClientVpn.Route("exampleRoute", new Aws.Ec2ClientVpn.RouteArgs
    ///         {
    ///             ClientVpnEndpointId = exampleEndpoint.Id,
    ///             DestinationCidrBlock = "0.0.0.0/0",
    ///             TargetVpcSubnetId = exampleNetworkAssociation.SubnetId,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AWS Client VPN routes can be imported using the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a `,`.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2clientvpn/route:Route example cvpn-endpoint-1234567890abcdef,subnet-9876543210fedcba,10.1.0.0/24
    /// ```
    /// </summary>
    public partial class Route : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Client VPN endpoint.
        /// </summary>
        [Output("clientVpnEndpointId")]
        public Output<string> ClientVpnEndpointId { get; private set; } = null!;

        /// <summary>
        /// A brief description of the authorization rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address range, in CIDR notation, of the route destination.
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Indicates how the Client VPN route was added. Will be `add-route` for routes created by this resource.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
        /// </summary>
        [Output("targetVpcSubnetId")]
        public Output<string> TargetVpcSubnetId { get; private set; } = null!;

        /// <summary>
        /// The type of the route.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2clientvpn/route:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2clientvpn/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Client VPN endpoint.
        /// </summary>
        [Input("clientVpnEndpointId", required: true)]
        public Input<string> ClientVpnEndpointId { get; set; } = null!;

        /// <summary>
        /// A brief description of the authorization rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IPv4 address range, in CIDR notation, of the route destination.
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
        /// </summary>
        [Input("targetVpcSubnetId", required: true)]
        public Input<string> TargetVpcSubnetId { get; set; } = null!;

        public RouteArgs()
        {
        }
    }

    public sealed class RouteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Client VPN endpoint.
        /// </summary>
        [Input("clientVpnEndpointId")]
        public Input<string>? ClientVpnEndpointId { get; set; }

        /// <summary>
        /// A brief description of the authorization rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IPv4 address range, in CIDR notation, of the route destination.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// Indicates how the Client VPN route was added. Will be `add-route` for routes created by this resource.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
        /// </summary>
        [Input("targetVpcSubnetId")]
        public Input<string>? TargetVpcSubnetId { get; set; }

        /// <summary>
        /// The type of the route.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RouteState()
        {
        }
    }
}

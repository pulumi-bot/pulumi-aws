// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectConnect
{
    /// <summary>
    /// Provides a Direct Connect public virtual interface resource.
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
    ///         var foo = new Aws.DirectConnect.PublicVirtualInterface("foo", new Aws.DirectConnect.PublicVirtualInterfaceArgs
    ///         {
    ///             AddressFamily = "ipv4",
    ///             AmazonAddress = "175.45.176.2/30",
    ///             BgpAsn = 65352,
    ///             ConnectionId = "dxcon-zzzzzzzz",
    ///             CustomerAddress = "175.45.176.1/30",
    ///             RouteFilterPrefixes = 
    ///             {
    ///                 "210.52.109.0/24",
    ///                 "175.45.176.0/22",
    ///             },
    ///             Vlan = 4094,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Direct Connect public virtual interfaces can be imported using the `vif id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:directconnect/publicVirtualInterface:PublicVirtualInterface test dxvif-33cc44dd
    /// ```
    /// </summary>
    public partial class PublicVirtualInterface : Pulumi.CustomResource
    {
        /// <summary>
        /// The address family for the BGP peer. `ipv4 ` or `ipv6`.
        /// </summary>
        [Output("addressFamily")]
        public Output<string> AddressFamily { get; private set; } = null!;

        /// <summary>
        /// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        /// </summary>
        [Output("amazonAddress")]
        public Output<string> AmazonAddress { get; private set; } = null!;

        [Output("amazonSideAsn")]
        public Output<string> AmazonSideAsn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the virtual interface.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Direct Connect endpoint on which the virtual interface terminates.
        /// </summary>
        [Output("awsDevice")]
        public Output<string> AwsDevice { get; private set; } = null!;

        /// <summary>
        /// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        /// </summary>
        [Output("bgpAsn")]
        public Output<int> BgpAsn { get; private set; } = null!;

        /// <summary>
        /// The authentication key for BGP configuration.
        /// </summary>
        [Output("bgpAuthKey")]
        public Output<string> BgpAuthKey { get; private set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        /// </summary>
        [Output("connectionId")]
        public Output<string> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        /// </summary>
        [Output("customerAddress")]
        public Output<string> CustomerAddress { get; private set; } = null!;

        /// <summary>
        /// The name for the virtual interface.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of routes to be advertised to the AWS network in this region.
        /// </summary>
        [Output("routeFilterPrefixes")]
        public Output<ImmutableArray<string>> RouteFilterPrefixes { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Output("vlan")]
        public Output<int> Vlan { get; private set; } = null!;


        /// <summary>
        /// Create a PublicVirtualInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicVirtualInterface(string name, PublicVirtualInterfaceArgs args, CustomResourceOptions? options = null)
            : base("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, args ?? new PublicVirtualInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicVirtualInterface(string name, Input<string> id, PublicVirtualInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PublicVirtualInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicVirtualInterface Get(string name, Input<string> id, PublicVirtualInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicVirtualInterface(name, id, state, options);
        }
    }

    public sealed class PublicVirtualInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address family for the BGP peer. `ipv4 ` or `ipv6`.
        /// </summary>
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        /// <summary>
        /// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        /// </summary>
        [Input("amazonAddress")]
        public Input<string>? AmazonAddress { get; set; }

        /// <summary>
        /// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        /// </summary>
        [Input("bgpAsn", required: true)]
        public Input<int> BgpAsn { get; set; } = null!;

        /// <summary>
        /// The authentication key for BGP configuration.
        /// </summary>
        [Input("bgpAuthKey")]
        public Input<string>? BgpAuthKey { get; set; }

        /// <summary>
        /// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        /// </summary>
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        /// <summary>
        /// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        /// </summary>
        [Input("customerAddress")]
        public Input<string>? CustomerAddress { get; set; }

        /// <summary>
        /// The name for the virtual interface.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("routeFilterPrefixes", required: true)]
        private InputList<string>? _routeFilterPrefixes;

        /// <summary>
        /// A list of routes to be advertised to the AWS network in this region.
        /// </summary>
        public InputList<string> RouteFilterPrefixes
        {
            get => _routeFilterPrefixes ?? (_routeFilterPrefixes = new InputList<string>());
            set => _routeFilterPrefixes = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Input("vlan", required: true)]
        public Input<int> Vlan { get; set; } = null!;

        public PublicVirtualInterfaceArgs()
        {
        }
    }

    public sealed class PublicVirtualInterfaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address family for the BGP peer. `ipv4 ` or `ipv6`.
        /// </summary>
        [Input("addressFamily")]
        public Input<string>? AddressFamily { get; set; }

        /// <summary>
        /// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        /// </summary>
        [Input("amazonAddress")]
        public Input<string>? AmazonAddress { get; set; }

        [Input("amazonSideAsn")]
        public Input<string>? AmazonSideAsn { get; set; }

        /// <summary>
        /// The ARN of the virtual interface.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Direct Connect endpoint on which the virtual interface terminates.
        /// </summary>
        [Input("awsDevice")]
        public Input<string>? AwsDevice { get; set; }

        /// <summary>
        /// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        /// </summary>
        [Input("bgpAsn")]
        public Input<int>? BgpAsn { get; set; }

        /// <summary>
        /// The authentication key for BGP configuration.
        /// </summary>
        [Input("bgpAuthKey")]
        public Input<string>? BgpAuthKey { get; set; }

        /// <summary>
        /// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        /// </summary>
        [Input("customerAddress")]
        public Input<string>? CustomerAddress { get; set; }

        /// <summary>
        /// The name for the virtual interface.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("routeFilterPrefixes")]
        private InputList<string>? _routeFilterPrefixes;

        /// <summary>
        /// A list of routes to be advertised to the AWS network in this region.
        /// </summary>
        public InputList<string> RouteFilterPrefixes
        {
            get => _routeFilterPrefixes ?? (_routeFilterPrefixes = new InputList<string>());
            set => _routeFilterPrefixes = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public PublicVirtualInterfaceState()
        {
        }
    }
}

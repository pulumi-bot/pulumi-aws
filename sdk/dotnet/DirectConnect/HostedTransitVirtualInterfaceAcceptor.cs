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
    /// Provides a resource to manage the accepter's side of a Direct Connect hosted transit virtual interface.
    /// This resource accepts ownership of a transit virtual interface created by another AWS account.
    /// 
    /// &gt; **NOTE:** AWS allows a Direct Connect hosted transit virtual interface to be deleted from either the allocator's or accepter's side. However, this provider only allows the Direct Connect hosted transit virtual interface to be deleted from the allocator's side by removing the corresponding `aws.directconnect.HostedTransitVirtualInterface` resource from your configuration. Removing a `aws.directconnect.HostedTransitVirtualInterfaceAcceptor` resource from your configuration will remove it from your statefile and management, **but will not delete the Direct Connect virtual interface.**
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
    ///         var accepter = new Aws.Provider("accepter", new Aws.ProviderArgs
    ///         {
    ///         });
    ///         // Accepter's credentials.
    ///         var accepterCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
    ///         // Accepter's side of the VIF.
    ///         var example = new Aws.DirectConnect.Gateway("example", new Aws.DirectConnect.GatewayArgs
    ///         {
    ///             AmazonSideAsn = "64512",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Accepter,
    ///         });
    ///         // Creator's side of the VIF
    ///         var creator = new Aws.DirectConnect.HostedTransitVirtualInterface("creator", new Aws.DirectConnect.HostedTransitVirtualInterfaceArgs
    ///         {
    ///             ConnectionId = "dxcon-zzzzzzzz",
    ///             OwnerAccountId = accepterCallerIdentity.Apply(accepterCallerIdentity =&gt; accepterCallerIdentity.AccountId),
    ///             Vlan = 4094,
    ///             AddressFamily = "ipv4",
    ///             BgpAsn = 65352,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 example,
    ///             },
    ///         });
    ///         var accepterHostedTransitVirtualInterfaceAcceptor = new Aws.DirectConnect.HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor", new Aws.DirectConnect.HostedTransitVirtualInterfaceAcceptorArgs
    ///         {
    ///             VirtualInterfaceId = creator.Id,
    ///             DxGatewayId = example.Id,
    ///             Tags = 
    ///             {
    ///                 { "Side", "Accepter" },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Accepter,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Direct Connect hosted transit virtual interfaces can be imported using the `vif id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor test dxvif-33cc44dd
    /// ```
    /// </summary>
    public partial class HostedTransitVirtualInterfaceAcceptor : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the virtual interface.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect gateway to which to connect the virtual interface.
        /// </summary>
        [Output("dxGatewayId")]
        public Output<string> DxGatewayId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect virtual interface to accept.
        /// </summary>
        [Output("virtualInterfaceId")]
        public Output<string> VirtualInterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a HostedTransitVirtualInterfaceAcceptor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostedTransitVirtualInterfaceAcceptor(string name, HostedTransitVirtualInterfaceAcceptorArgs args, CustomResourceOptions? options = null)
            : base("aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor", name, args ?? new HostedTransitVirtualInterfaceAcceptorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostedTransitVirtualInterfaceAcceptor(string name, Input<string> id, HostedTransitVirtualInterfaceAcceptorState? state = null, CustomResourceOptions? options = null)
            : base("aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HostedTransitVirtualInterfaceAcceptor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostedTransitVirtualInterfaceAcceptor Get(string name, Input<string> id, HostedTransitVirtualInterfaceAcceptorState? state = null, CustomResourceOptions? options = null)
        {
            return new HostedTransitVirtualInterfaceAcceptor(name, id, state, options);
        }
    }

    public sealed class HostedTransitVirtualInterfaceAcceptorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Direct Connect gateway to which to connect the virtual interface.
        /// </summary>
        [Input("dxGatewayId", required: true)]
        public Input<string> DxGatewayId { get; set; } = null!;

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
        /// The ID of the Direct Connect virtual interface to accept.
        /// </summary>
        [Input("virtualInterfaceId", required: true)]
        public Input<string> VirtualInterfaceId { get; set; } = null!;

        public HostedTransitVirtualInterfaceAcceptorArgs()
        {
        }
    }

    public sealed class HostedTransitVirtualInterfaceAcceptorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the virtual interface.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ID of the Direct Connect gateway to which to connect the virtual interface.
        /// </summary>
        [Input("dxGatewayId")]
        public Input<string>? DxGatewayId { get; set; }

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
        /// The ID of the Direct Connect virtual interface to accept.
        /// </summary>
        [Input("virtualInterfaceId")]
        public Input<string>? VirtualInterfaceId { get; set; }

        public HostedTransitVirtualInterfaceAcceptorState()
        {
        }
    }
}

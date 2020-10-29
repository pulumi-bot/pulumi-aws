// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    /// <summary>
    /// Manages an EC2 Transit Gateway Peering Attachment.
    /// For examples of custom route table association and propagation, see the [EC2 Transit Gateway Networking Examples Guide](https://docs.aws.amazon.com/vpc/latest/tgw/TGW_Scenarios.html).
    /// </summary>
    public partial class PeeringAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        /// </summary>
        [Output("peerAccountId")]
        public Output<string> PeerAccountId { get; private set; } = null!;

        /// <summary>
        /// Region of EC2 Transit Gateway to peer with.
        /// </summary>
        [Output("peerRegion")]
        public Output<string> PeerRegion { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway to peer with.
        /// </summary>
        [Output("peerTransitGatewayId")]
        public Output<string> PeerTransitGatewayId { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Peering Attachment.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Output("transitGatewayId")]
        public Output<string> TransitGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a PeeringAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PeeringAttachment(string name, PeeringAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/peeringAttachment:PeeringAttachment", name, args ?? new PeeringAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PeeringAttachment(string name, Input<string> id, PeeringAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/peeringAttachment:PeeringAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PeeringAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PeeringAttachment Get(string name, Input<string> id, PeeringAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new PeeringAttachment(name, id, state, options);
        }
    }

    public sealed class PeeringAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        /// </summary>
        [Input("peerAccountId")]
        public Input<string>? PeerAccountId { get; set; }

        /// <summary>
        /// Region of EC2 Transit Gateway to peer with.
        /// </summary>
        [Input("peerRegion", required: true)]
        public Input<string> PeerRegion { get; set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway to peer with.
        /// </summary>
        [Input("peerTransitGatewayId", required: true)]
        public Input<string> PeerTransitGatewayId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Peering Attachment.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId", required: true)]
        public Input<string> TransitGatewayId { get; set; } = null!;

        public PeeringAttachmentArgs()
        {
        }
    }

    public sealed class PeeringAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        /// </summary>
        [Input("peerAccountId")]
        public Input<string>? PeerAccountId { get; set; }

        /// <summary>
        /// Region of EC2 Transit Gateway to peer with.
        /// </summary>
        [Input("peerRegion")]
        public Input<string>? PeerRegion { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway to peer with.
        /// </summary>
        [Input("peerTransitGatewayId")]
        public Input<string>? PeerTransitGatewayId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Peering Attachment.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        public PeeringAttachmentState()
        {
        }
    }
}

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
    /// Manages the accepter's side of an EC2 Transit Gateway Peering Attachment.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2.TransitGatewayPeeringAttachmentAccepter("example", new Aws.Ec2.TransitGatewayPeeringAttachmentAccepterArgs
    ///         {
    ///             Tags = 
    ///             {
    ///                 { "Name", "Example cross-account attachment" },
    ///             },
    ///             TransitGatewayAttachmentId = aws_ec2_transit_gateway_peering_attachment.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TransitGatewayPeeringAttachmentAccepter : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier of the AWS account that owns the EC2 TGW peering.
        /// </summary>
        [Output("peerAccountId")]
        public Output<string> PeerAccountId { get; private set; } = null!;

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
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the EC2 Transit Gateway Peering Attachment to manage.
        /// </summary>
        [Output("transitGatewayAttachmentId")]
        public Output<string> TransitGatewayAttachmentId { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Output("transitGatewayId")]
        public Output<string> TransitGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitGatewayPeeringAttachmentAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitGatewayPeeringAttachmentAccepter(string name, TransitGatewayPeeringAttachmentAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/transitGatewayPeeringAttachmentAccepter:TransitGatewayPeeringAttachmentAccepter", name, args ?? new TransitGatewayPeeringAttachmentAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitGatewayPeeringAttachmentAccepter(string name, Input<string> id, TransitGatewayPeeringAttachmentAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/transitGatewayPeeringAttachmentAccepter:TransitGatewayPeeringAttachmentAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitGatewayPeeringAttachmentAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitGatewayPeeringAttachmentAccepter Get(string name, Input<string> id, TransitGatewayPeeringAttachmentAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitGatewayPeeringAttachmentAccepter(name, id, state, options);
        }
    }

    public sealed class TransitGatewayPeeringAttachmentAccepterArgs : Pulumi.ResourceArgs
    {
        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Peering Attachment.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the EC2 Transit Gateway Peering Attachment to manage.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public Input<string> TransitGatewayAttachmentId { get; set; } = null!;

        public TransitGatewayPeeringAttachmentAccepterArgs()
        {
        }
    }

    public sealed class TransitGatewayPeeringAttachmentAccepterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the AWS account that owns the EC2 TGW peering.
        /// </summary>
        [Input("peerAccountId")]
        public Input<string>? PeerAccountId { get; set; }

        [Input("peerRegion")]
        public Input<string>? PeerRegion { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway to peer with.
        /// </summary>
        [Input("peerTransitGatewayId")]
        public Input<string>? PeerTransitGatewayId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway Peering Attachment.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the EC2 Transit Gateway Peering Attachment to manage.
        /// </summary>
        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        public TransitGatewayPeeringAttachmentAccepterState()
        {
        }
    }
}

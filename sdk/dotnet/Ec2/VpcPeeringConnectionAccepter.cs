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
    /// Provides a resource to manage the accepter's side of a VPC Peering Connection.
    /// 
    /// When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
    /// VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
    /// accepter's account.
    /// The requester can use the `aws.ec2.VpcPeeringConnection` resource to manage its side of the connection
    /// and the accepter can use the `aws.ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
    /// connection into management.
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
    ///         var peer = new Aws.Provider("peer", new Aws.ProviderArgs
    ///         {
    ///             Region = "us-west-2",
    ///         });
    ///         var main = new Aws.Ec2.Vpc("main", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var peerVpc = new Aws.Ec2.Vpc("peerVpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.1.0.0/16",
    ///         });
    ///         var peerCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
    ///         // Requester's side of the connection.
    ///         var peerVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("peerVpcPeeringConnection", new Aws.Ec2.VpcPeeringConnectionArgs
    ///         {
    ///             AutoAccept = false,
    ///             PeerOwnerId = peerCallerIdentity.Apply(peerCallerIdentity =&gt; peerCallerIdentity.AccountId),
    ///             PeerRegion = "us-west-2",
    ///             PeerVpcId = peerVpc.Id,
    ///             Tags = 
    ///             {
    ///                 { "Side", "Requester" },
    ///             },
    ///             VpcId = main.Id,
    ///         });
    ///         // Accepter's side of the connection.
    ///         var peerVpcPeeringConnectionAccepter = new Aws.Ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", new Aws.Ec2.VpcPeeringConnectionAccepterArgs
    ///         {
    ///             AutoAccept = true,
    ///             Tags = 
    ///             {
    ///                 { "Side", "Accepter" },
    ///             },
    ///             VpcPeeringConnectionId = peerVpcPeeringConnection.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class VpcPeeringConnectionAccepter : Pulumi.CustomResource
    {
        /// <summary>
        /// The status of the VPC Peering Connection request.
        /// </summary>
        [Output("acceptStatus")]
        public Output<string> AcceptStatus { get; private set; } = null!;

        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        /// </summary>
        [Output("accepter")]
        public Output<Outputs.VpcPeeringConnectionAccepterAccepter> Accepter { get; private set; } = null!;

        /// <summary>
        /// Whether or not to accept the peering request. Defaults to `false`.
        /// </summary>
        [Output("autoAccept")]
        public Output<bool?> AutoAccept { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID of the owner of the requester VPC.
        /// </summary>
        [Output("peerOwnerId")]
        public Output<string> PeerOwnerId { get; private set; } = null!;

        /// <summary>
        /// The region of the accepter VPC.
        /// </summary>
        [Output("peerRegion")]
        public Output<string> PeerRegion { get; private set; } = null!;

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Output("peerVpcId")]
        public Output<string> PeerVpcId { get; private set; } = null!;

        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        /// </summary>
        [Output("requester")]
        public Output<Outputs.VpcPeeringConnectionAccepterRequester> Requester { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the accepter VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The VPC Peering Connection ID to manage.
        /// </summary>
        [Output("vpcPeeringConnectionId")]
        public Output<string> VpcPeeringConnectionId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPeeringConnectionAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPeeringConnectionAccepter(string name, VpcPeeringConnectionAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, args ?? new VpcPeeringConnectionAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPeeringConnectionAccepter(string name, Input<string> id, VpcPeeringConnectionAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcPeeringConnectionAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPeeringConnectionAccepter Get(string name, Input<string> id, VpcPeeringConnectionAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPeeringConnectionAccepter(name, id, state, options);
        }
    }

    public sealed class VpcPeeringConnectionAccepterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        /// </summary>
        [Input("accepter")]
        public Input<Inputs.VpcPeeringConnectionAccepterAccepterArgs>? Accepter { get; set; }

        /// <summary>
        /// Whether or not to accept the peering request. Defaults to `false`.
        /// </summary>
        [Input("autoAccept")]
        public Input<bool>? AutoAccept { get; set; }

        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        /// </summary>
        [Input("requester")]
        public Input<Inputs.VpcPeeringConnectionAccepterRequesterArgs>? Requester { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC Peering Connection ID to manage.
        /// </summary>
        [Input("vpcPeeringConnectionId", required: true)]
        public Input<string> VpcPeeringConnectionId { get; set; } = null!;

        public VpcPeeringConnectionAccepterArgs()
        {
        }
    }

    public sealed class VpcPeeringConnectionAccepterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status of the VPC Peering Connection request.
        /// </summary>
        [Input("acceptStatus")]
        public Input<string>? AcceptStatus { get; set; }

        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        /// </summary>
        [Input("accepter")]
        public Input<Inputs.VpcPeeringConnectionAccepterAccepterGetArgs>? Accepter { get; set; }

        /// <summary>
        /// Whether or not to accept the peering request. Defaults to `false`.
        /// </summary>
        [Input("autoAccept")]
        public Input<bool>? AutoAccept { get; set; }

        /// <summary>
        /// The AWS account ID of the owner of the requester VPC.
        /// </summary>
        [Input("peerOwnerId")]
        public Input<string>? PeerOwnerId { get; set; }

        /// <summary>
        /// The region of the accepter VPC.
        /// </summary>
        [Input("peerRegion")]
        public Input<string>? PeerRegion { get; set; }

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Input("peerVpcId")]
        public Input<string>? PeerVpcId { get; set; }

        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        /// </summary>
        [Input("requester")]
        public Input<Inputs.VpcPeeringConnectionAccepterRequesterGetArgs>? Requester { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the accepter VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The VPC Peering Connection ID to manage.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public Input<string>? VpcPeeringConnectionId { get; set; }

        public VpcPeeringConnectionAccepterState()
        {
        }
    }
}

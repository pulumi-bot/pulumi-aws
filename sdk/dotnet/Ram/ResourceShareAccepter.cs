// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ram
{
    /// <summary>
    /// Manage accepting a Resource Access Manager (RAM) Resource Share invitation. From a _receiver_ AWS account, accept an invitation to share resources that were shared by a _sender_ AWS account. To create a resource share in the _sender_, see the `aws.ram.ResourceShare` resource.
    /// 
    /// &gt; **Note:** If both AWS accounts are in the same Organization and [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), this resource is not necessary as RAM Resource Share invitations are not used.
    /// </summary>
    public partial class ResourceShareAccepter : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the resource share invitation.
        /// </summary>
        [Output("invitationArn")]
        public Output<string> InvitationArn { get; private set; } = null!;

        /// <summary>
        /// The account ID of the receiver account which accepts the invitation.
        /// </summary>
        [Output("receiverAccountId")]
        public Output<string> ReceiverAccountId { get; private set; } = null!;

        /// <summary>
        /// A list of the resource ARNs shared via the resource share.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<string>> Resources { get; private set; } = null!;

        /// <summary>
        /// The account ID of the sender account which submits the invitation.
        /// </summary>
        [Output("senderAccountId")]
        public Output<string> SenderAccountId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the resource share.
        /// </summary>
        [Output("shareArn")]
        public Output<string> ShareArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource share as displayed in the console.
        /// </summary>
        [Output("shareId")]
        public Output<string> ShareId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource share.
        /// </summary>
        [Output("shareName")]
        public Output<string> ShareName { get; private set; } = null!;

        /// <summary>
        /// The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceShareAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceShareAccepter(string name, ResourceShareAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:ram/resourceShareAccepter:ResourceShareAccepter", name, args ?? new ResourceShareAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceShareAccepter(string name, Input<string> id, ResourceShareAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ram/resourceShareAccepter:ResourceShareAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceShareAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceShareAccepter Get(string name, Input<string> id, ResourceShareAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceShareAccepter(name, id, state, options);
        }
    }

    public sealed class ResourceShareAccepterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the resource share.
        /// </summary>
        [Input("shareArn", required: true)]
        public Input<string> ShareArn { get; set; } = null!;

        public ResourceShareAccepterArgs()
        {
        }
    }

    public sealed class ResourceShareAccepterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the resource share invitation.
        /// </summary>
        [Input("invitationArn")]
        public Input<string>? InvitationArn { get; set; }

        /// <summary>
        /// The account ID of the receiver account which accepts the invitation.
        /// </summary>
        [Input("receiverAccountId")]
        public Input<string>? ReceiverAccountId { get; set; }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// A list of the resource ARNs shared via the resource share.
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        /// <summary>
        /// The account ID of the sender account which submits the invitation.
        /// </summary>
        [Input("senderAccountId")]
        public Input<string>? SenderAccountId { get; set; }

        /// <summary>
        /// The ARN of the resource share.
        /// </summary>
        [Input("shareArn")]
        public Input<string>? ShareArn { get; set; }

        /// <summary>
        /// The ID of the resource share as displayed in the console.
        /// </summary>
        [Input("shareId")]
        public Input<string>? ShareId { get; set; }

        /// <summary>
        /// The name of the resource share.
        /// </summary>
        [Input("shareName")]
        public Input<string>? ShareName { get; set; }

        /// <summary>
        /// The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ResourceShareAccepterState()
        {
        }
    }
}

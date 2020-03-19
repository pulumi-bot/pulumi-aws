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
    /// Associates a Direct Connect Gateway with a VGW or transit gateway.
    /// 
    /// To create a cross-account association, create an [`aws.directconnect.GatewayAssociationProposal` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association_proposal.html)
    /// in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
    /// by creating an `aws.directconnect.GatewayAssociation` resource with the `proposal_id` and `associated_gateway_owner_account_id` attributes set.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_gateway_association.html.markdown.
    /// </summary>
    public partial class GatewayAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        /// </summary>
        [Output("allowedPrefixes")]
        public Output<ImmutableArray<string>> AllowedPrefixes { get; private set; } = null!;

        /// <summary>
        /// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
        /// Used for single account Direct Connect gateway associations.
        /// </summary>
        [Output("associatedGatewayId")]
        public Output<string> AssociatedGatewayId { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
        /// Used for cross-account Direct Connect gateway associations.
        /// </summary>
        [Output("associatedGatewayOwnerAccountId")]
        public Output<string> AssociatedGatewayOwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
        /// </summary>
        [Output("associatedGatewayType")]
        public Output<string> AssociatedGatewayType { get; private set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect gateway association.
        /// </summary>
        [Output("dxGatewayAssociationId")]
        public Output<string> DxGatewayAssociationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect gateway.
        /// </summary>
        [Output("dxGatewayId")]
        public Output<string> DxGatewayId { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the Direct Connect gateway.
        /// </summary>
        [Output("dxGatewayOwnerAccountId")]
        public Output<string> DxGatewayOwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect gateway association proposal.
        /// Used for cross-account Direct Connect gateway associations.
        /// </summary>
        [Output("proposalId")]
        public Output<string?> ProposalId { get; private set; } = null!;

        /// <summary>
        /// *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
        /// Used for single account Direct Connect gateway associations.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string?> VpnGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayAssociation(string name, GatewayAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:directconnect/gatewayAssociation:GatewayAssociation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GatewayAssociation(string name, Input<string> id, GatewayAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:directconnect/gatewayAssociation:GatewayAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayAssociation Get(string name, Input<string> id, GatewayAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayAssociation(name, id, state, options);
        }
    }

    public sealed class GatewayAssociationArgs : Pulumi.ResourceArgs
    {
        [Input("allowedPrefixes")]
        private InputList<string>? _allowedPrefixes;

        /// <summary>
        /// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        /// </summary>
        public InputList<string> AllowedPrefixes
        {
            get => _allowedPrefixes ?? (_allowedPrefixes = new InputList<string>());
            set => _allowedPrefixes = value;
        }

        /// <summary>
        /// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
        /// Used for single account Direct Connect gateway associations.
        /// </summary>
        [Input("associatedGatewayId")]
        public Input<string>? AssociatedGatewayId { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
        /// Used for cross-account Direct Connect gateway associations.
        /// </summary>
        [Input("associatedGatewayOwnerAccountId")]
        public Input<string>? AssociatedGatewayOwnerAccountId { get; set; }

        /// <summary>
        /// The ID of the Direct Connect gateway.
        /// </summary>
        [Input("dxGatewayId", required: true)]
        public Input<string> DxGatewayId { get; set; } = null!;

        /// <summary>
        /// The ID of the Direct Connect gateway association proposal.
        /// Used for cross-account Direct Connect gateway associations.
        /// </summary>
        [Input("proposalId")]
        public Input<string>? ProposalId { get; set; }

        /// <summary>
        /// *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
        /// Used for single account Direct Connect gateway associations.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public GatewayAssociationArgs()
        {
        }
    }

    public sealed class GatewayAssociationState : Pulumi.ResourceArgs
    {
        [Input("allowedPrefixes")]
        private InputList<string>? _allowedPrefixes;

        /// <summary>
        /// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        /// </summary>
        public InputList<string> AllowedPrefixes
        {
            get => _allowedPrefixes ?? (_allowedPrefixes = new InputList<string>());
            set => _allowedPrefixes = value;
        }

        /// <summary>
        /// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
        /// Used for single account Direct Connect gateway associations.
        /// </summary>
        [Input("associatedGatewayId")]
        public Input<string>? AssociatedGatewayId { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
        /// Used for cross-account Direct Connect gateway associations.
        /// </summary>
        [Input("associatedGatewayOwnerAccountId")]
        public Input<string>? AssociatedGatewayOwnerAccountId { get; set; }

        /// <summary>
        /// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
        /// </summary>
        [Input("associatedGatewayType")]
        public Input<string>? AssociatedGatewayType { get; set; }

        /// <summary>
        /// The ID of the Direct Connect gateway association.
        /// </summary>
        [Input("dxGatewayAssociationId")]
        public Input<string>? DxGatewayAssociationId { get; set; }

        /// <summary>
        /// The ID of the Direct Connect gateway.
        /// </summary>
        [Input("dxGatewayId")]
        public Input<string>? DxGatewayId { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the Direct Connect gateway.
        /// </summary>
        [Input("dxGatewayOwnerAccountId")]
        public Input<string>? DxGatewayOwnerAccountId { get; set; }

        /// <summary>
        /// The ID of the Direct Connect gateway association proposal.
        /// Used for cross-account Direct Connect gateway associations.
        /// </summary>
        [Input("proposalId")]
        public Input<string>? ProposalId { get; set; }

        /// <summary>
        /// *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
        /// Used for single account Direct Connect gateway associations.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public GatewayAssociationState()
        {
        }
    }
}

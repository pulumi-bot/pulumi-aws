// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    public static class GetTransitGateway
    {
        /// <summary>
        /// Get information on an EC2 Transit Gateway.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTransitGatewayResult> InvokeAsync(GetTransitGatewayArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayResult>("aws:ec2transitgateway/getTransitGateway:getTransitGateway", args ?? new GetTransitGatewayArgs(), options.WithVersion());
    }


    public sealed class GetTransitGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetTransitGatewayFilterArgs>? _filters;

        /// <summary>
        /// One or more configuration blocks containing name-values filters. Detailed below.
        /// </summary>
        public List<Inputs.GetTransitGatewayFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetTransitGatewayFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Identifier of the EC2 Transit Gateway.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetTransitGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransitGatewayResult
    {
        /// <summary>
        /// Private Autonomous System Number (ASN) for the Amazon side of a BGP session
        /// </summary>
        public readonly int AmazonSideAsn;
        /// <summary>
        /// EC2 Transit Gateway Amazon Resource Name (ARN)
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Identifier of the default association route table
        /// </summary>
        public readonly string AssociationDefaultRouteTableId;
        /// <summary>
        /// Whether resource attachment requests are automatically accepted.
        /// </summary>
        public readonly string AutoAcceptSharedAttachments;
        /// <summary>
        /// Whether resource attachments are automatically associated with the default association route table.
        /// </summary>
        public readonly string DefaultRouteTableAssociation;
        /// <summary>
        /// Whether resource attachments automatically propagate routes to the default propagation route table.
        /// </summary>
        public readonly string DefaultRouteTablePropagation;
        /// <summary>
        /// Description of the EC2 Transit Gateway
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether DNS support is enabled.
        /// </summary>
        public readonly string DnsSupport;
        public readonly ImmutableArray<Outputs.GetTransitGatewayFilterResult> Filters;
        /// <summary>
        /// EC2 Transit Gateway identifier
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Identifier of the AWS account that owns the EC2 Transit Gateway
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// Identifier of the default propagation route table.
        /// </summary>
        public readonly string PropagationDefaultRouteTableId;
        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Whether VPN Equal Cost Multipath Protocol support is enabled.
        /// </summary>
        public readonly string VpnEcmpSupport;

        [OutputConstructor]
        private GetTransitGatewayResult(
            int amazonSideAsn,

            string arn,

            string associationDefaultRouteTableId,

            string autoAcceptSharedAttachments,

            string defaultRouteTableAssociation,

            string defaultRouteTablePropagation,

            string description,

            string dnsSupport,

            ImmutableArray<Outputs.GetTransitGatewayFilterResult> filters,

            string? id,

            string ownerId,

            string propagationDefaultRouteTableId,

            ImmutableDictionary<string, string> tags,

            string vpnEcmpSupport)
        {
            AmazonSideAsn = amazonSideAsn;
            Arn = arn;
            AssociationDefaultRouteTableId = associationDefaultRouteTableId;
            AutoAcceptSharedAttachments = autoAcceptSharedAttachments;
            DefaultRouteTableAssociation = defaultRouteTableAssociation;
            DefaultRouteTablePropagation = defaultRouteTablePropagation;
            Description = description;
            DnsSupport = dnsSupport;
            Filters = filters;
            Id = id;
            OwnerId = ownerId;
            PropagationDefaultRouteTableId = propagationDefaultRouteTableId;
            Tags = tags;
            VpnEcmpSupport = vpnEcmpSupport;
        }
    }
}

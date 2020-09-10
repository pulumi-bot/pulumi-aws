// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    public static class GetVpcAttachment
    {
        public static Task<GetVpcAttachmentResult> InvokeAsync(GetVpcAttachmentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcAttachmentResult>("aws:ec2transitgateway/getVpcAttachment:getVpcAttachment", args ?? new GetVpcAttachmentArgs(), options.WithVersion());
    }


    public sealed class GetVpcAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcAttachmentFilterArgs>? _filters;
        public List<Inputs.GetVpcAttachmentFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcAttachmentFilterArgs>());
            set => _filters = value;
        }

        [Input("id")]
        public string? Id { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetVpcAttachmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcAttachmentResult
    {
        public readonly string DnsSupport;
        public readonly ImmutableArray<Outputs.GetVpcAttachmentFilterResult> Filters;
        public readonly string? Id;
        public readonly string Ipv6Support;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string TransitGatewayId;
        public readonly string VpcId;
        public readonly string VpcOwnerId;

        [OutputConstructor]
        private GetVpcAttachmentResult(
            string dnsSupport,

            ImmutableArray<Outputs.GetVpcAttachmentFilterResult> filters,

            string? id,

            string ipv6Support,

            ImmutableArray<string> subnetIds,

            ImmutableDictionary<string, string> tags,

            string transitGatewayId,

            string vpcId,

            string vpcOwnerId)
        {
            DnsSupport = dnsSupport;
            Filters = filters;
            Id = id;
            Ipv6Support = ipv6Support;
            SubnetIds = subnetIds;
            Tags = tags;
            TransitGatewayId = transitGatewayId;
            VpcId = vpcId;
            VpcOwnerId = vpcOwnerId;
        }
    }
}

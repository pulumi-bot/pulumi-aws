// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    public static class GetDirectConnectGatewayAttachment
    {
        public static Task<GetDirectConnectGatewayAttachmentResult> InvokeAsync(GetDirectConnectGatewayAttachmentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDirectConnectGatewayAttachmentResult>("aws:ec2transitgateway/getDirectConnectGatewayAttachment:getDirectConnectGatewayAttachment", args ?? new GetDirectConnectGatewayAttachmentArgs(), options.WithVersion());
    }


    public sealed class GetDirectConnectGatewayAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("dxGatewayId")]
        public string? DxGatewayId { get; set; }

        [Input("filters")]
        private List<Inputs.GetDirectConnectGatewayAttachmentFilterArgs>? _filters;
        public List<Inputs.GetDirectConnectGatewayAttachmentFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDirectConnectGatewayAttachmentFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        [Input("transitGatewayId")]
        public string? TransitGatewayId { get; set; }

        public GetDirectConnectGatewayAttachmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDirectConnectGatewayAttachmentResult
    {
        public readonly string? DxGatewayId;
        public readonly ImmutableArray<Outputs.GetDirectConnectGatewayAttachmentFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string? TransitGatewayId;

        [OutputConstructor]
        private GetDirectConnectGatewayAttachmentResult(
            string? dxGatewayId,

            ImmutableArray<Outputs.GetDirectConnectGatewayAttachmentFilterResult> filters,

            string id,

            ImmutableDictionary<string, string> tags,

            string? transitGatewayId)
        {
            DxGatewayId = dxGatewayId;
            Filters = filters;
            Id = id;
            Tags = tags;
            TransitGatewayId = transitGatewayId;
        }
    }
}

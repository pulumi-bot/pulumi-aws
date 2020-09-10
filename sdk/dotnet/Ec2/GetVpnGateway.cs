// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpnGateway
    {
        public static Task<GetVpnGatewayResult> InvokeAsync(GetVpnGatewayArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpnGatewayResult>("aws:ec2/getVpnGateway:getVpnGateway", args ?? new GetVpnGatewayArgs(), options.WithVersion());
    }


    public sealed class GetVpnGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("amazonSideAsn")]
        public string? AmazonSideAsn { get; set; }

        [Input("attachedVpcId")]
        public string? AttachedVpcId { get; set; }

        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("filters")]
        private List<Inputs.GetVpnGatewayFilterArgs>? _filters;
        public List<Inputs.GetVpnGatewayFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpnGatewayFilterArgs>());
            set => _filters = value;
        }

        [Input("id")]
        public string? Id { get; set; }

        [Input("state")]
        public string? State { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetVpnGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpnGatewayResult
    {
        public readonly string AmazonSideAsn;
        public readonly string Arn;
        public readonly string AttachedVpcId;
        public readonly string AvailabilityZone;
        public readonly ImmutableArray<Outputs.GetVpnGatewayFilterResult> Filters;
        public readonly string Id;
        public readonly string State;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetVpnGatewayResult(
            string amazonSideAsn,

            string arn,

            string attachedVpcId,

            string availabilityZone,

            ImmutableArray<Outputs.GetVpnGatewayFilterResult> filters,

            string id,

            string state,

            ImmutableDictionary<string, string> tags)
        {
            AmazonSideAsn = amazonSideAsn;
            Arn = arn;
            AttachedVpcId = attachedVpcId;
            AvailabilityZone = availabilityZone;
            Filters = filters;
            Id = id;
            State = state;
            Tags = tags;
        }
    }
}

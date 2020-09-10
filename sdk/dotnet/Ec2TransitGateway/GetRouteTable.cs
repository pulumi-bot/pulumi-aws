// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    public static class GetRouteTable
    {
        public static Task<GetRouteTableResult> InvokeAsync(GetRouteTableArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteTableResult>("aws:ec2transitgateway/getRouteTable:getRouteTable", args ?? new GetRouteTableArgs(), options.WithVersion());
    }


    public sealed class GetRouteTableArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetRouteTableFilterArgs>? _filters;
        public List<Inputs.GetRouteTableFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetRouteTableFilterArgs>());
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

        public GetRouteTableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteTableResult
    {
        public readonly bool DefaultAssociationRouteTable;
        public readonly bool DefaultPropagationRouteTable;
        public readonly ImmutableArray<Outputs.GetRouteTableFilterResult> Filters;
        public readonly string? Id;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string TransitGatewayId;

        [OutputConstructor]
        private GetRouteTableResult(
            bool defaultAssociationRouteTable,

            bool defaultPropagationRouteTable,

            ImmutableArray<Outputs.GetRouteTableFilterResult> filters,

            string? id,

            ImmutableDictionary<string, string> tags,

            string transitGatewayId)
        {
            DefaultAssociationRouteTable = defaultAssociationRouteTable;
            DefaultPropagationRouteTable = defaultPropagationRouteTable;
            Filters = filters;
            Id = id;
            Tags = tags;
            TransitGatewayId = transitGatewayId;
        }
    }
}

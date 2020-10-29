// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetNatGateway
    {
        /// <summary>
        /// Provides details about a specific Nat Gateway.
        /// </summary>
        public static Task<GetNatGatewayResult> InvokeAsync(GetNatGatewayArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNatGatewayResult>("aws:ec2/getNatGateway:getNatGateway", args ?? new GetNatGatewayArgs(), options.WithVersion());
    }


    public sealed class GetNatGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetNatGatewayFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetNatGatewayFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetNatGatewayFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The id of the specific Nat Gateway to retrieve.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The state of the NAT gateway (pending | failed | available | deleting | deleted ).
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// The id of subnet that the Nat Gateway resides in.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired Nat Gateway.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the VPC that the Nat Gateway resides in.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetNatGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNatGatewayResult
    {
        /// <summary>
        /// The Id of the EIP allocated to the selected Nat Gateway.
        /// </summary>
        public readonly string AllocationId;
        public readonly ImmutableArray<Outputs.GetNatGatewayFilterResult> Filters;
        public readonly string Id;
        /// <summary>
        /// The Id of the ENI allocated to the selected Nat Gateway.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// The private Ip address of the selected Nat Gateway.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// The public Ip (EIP) address of the selected Nat Gateway.
        /// </summary>
        public readonly string PublicIp;
        public readonly string State;
        public readonly string SubnetId;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetNatGatewayResult(
            string allocationId,

            ImmutableArray<Outputs.GetNatGatewayFilterResult> filters,

            string id,

            string networkInterfaceId,

            string privateIp,

            string publicIp,

            string state,

            string subnetId,

            ImmutableDictionary<string, string> tags,

            string vpcId)
        {
            AllocationId = allocationId;
            Filters = filters;
            Id = id;
            NetworkInterfaceId = networkInterfaceId;
            PrivateIp = privateIp;
            PublicIp = publicIp;
            State = state;
            SubnetId = subnetId;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}

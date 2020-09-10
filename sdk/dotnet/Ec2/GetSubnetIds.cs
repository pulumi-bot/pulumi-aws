// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetSubnetIds
    {
        public static Task<GetSubnetIdsResult> InvokeAsync(GetSubnetIdsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetIdsResult>("aws:ec2/getSubnetIds:getSubnetIds", args ?? new GetSubnetIdsArgs(), options.WithVersion());
    }


    public sealed class GetSubnetIdsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSubnetIdsFilterArgs>? _filters;
        public List<Inputs.GetSubnetIdsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSubnetIdsFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        [Input("vpcId", required: true)]
        public string VpcId { get; set; } = null!;

        public GetSubnetIdsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubnetIdsResult
    {
        public readonly ImmutableArray<Outputs.GetSubnetIdsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetSubnetIdsResult(
            ImmutableArray<Outputs.GetSubnetIdsFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableDictionary<string, string> tags,

            string vpcId)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}

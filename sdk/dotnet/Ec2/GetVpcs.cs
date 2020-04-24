// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpcs
    {
        /// <summary>
        /// This resource can be useful for getting back a list of VPC Ids for a region.
        /// 
        /// The following example retrieves a list of VPC Ids with a custom tag of `service` set to a value of "production".
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcsResult> InvokeAsync(GetVpcsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcsResult>("aws:ec2/getVpcs:getVpcs", args ?? new GetVpcsArgs(), options.WithVersion());
    }


    public sealed class GetVpcsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcsFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcsFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match
        /// a pair on the desired vpcs.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetVpcsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcsResult
    {
        public readonly ImmutableArray<Outputs.GetVpcsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of all the VPC Ids found. This data source will fail if none are found.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetVpcsResult(
            ImmutableArray<Outputs.GetVpcsFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableDictionary<string, object> tags)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Tags = tags;
        }
    }
}

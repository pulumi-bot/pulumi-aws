// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetPrefixList
    {
        /// <summary>
        /// `aws..getPrefixList` provides details about a specific prefix list (PL)
        /// in the current region.
        /// 
        /// This can be used both to validate a prefix list given in a variable
        /// and to obtain the CIDR blocks (IP address ranges) for the associated
        /// AWS service. The latter may be useful e.g. for adding network ACL
        /// rules.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPrefixListResult> InvokeAsync(GetPrefixListArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrefixListResult>("aws:index/getPrefixList:getPrefixList", args ?? new GetPrefixListArgs(), options.WithVersion());
    }


    public sealed class GetPrefixListArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetPrefixListFilterArgs>? _filters;

        /// <summary>
        /// Configuration block(s) for filtering. Detailed below.
        /// </summary>
        public List<Inputs.GetPrefixListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetPrefixListFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the prefix list to select.
        /// </summary>
        [Input("prefixListId")]
        public string? PrefixListId { get; set; }

        public GetPrefixListArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrefixListResult
    {
        /// <summary>
        /// The list of CIDR blocks for the AWS service associated with the prefix list.
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocks;
        public readonly ImmutableArray<Outputs.GetPrefixListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the selected prefix list.
        /// </summary>
        public readonly string Name;
        public readonly string? PrefixListId;

        [OutputConstructor]
        private GetPrefixListResult(
            ImmutableArray<string> cidrBlocks,

            ImmutableArray<Outputs.GetPrefixListFilterResult> filters,

            string id,

            string name,

            string? prefixListId)
        {
            CidrBlocks = cidrBlocks;
            Filters = filters;
            Id = id;
            Name = name;
            PrefixListId = prefixListId;
        }
    }
}

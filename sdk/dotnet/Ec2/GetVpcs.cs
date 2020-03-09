// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static partial class GetVpcs
    {
        /// <summary>
        /// This resource can be useful for getting back a list of VPC Ids for a region.
        /// 
        /// The following example retrieves a list of VPC Ids with a custom tag of `service` set to a value of "production".
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpcs.html.markdown.
        /// </summary>
        public static Task<GetVpcsResult> InvokeAsync(GetVpcsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcsResult>("aws:ec2/getVpcs:getVpcs", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVpcsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcsFiltersArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcsFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcsFiltersArgs>());
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
        public readonly ImmutableArray<Outputs.GetVpcsFiltersResult> Filters;
        /// <summary>
        /// A list of all the VPC Ids found. This data source will fail if none are found.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVpcsResult(
            ImmutableArray<Outputs.GetVpcsFiltersResult> filters,
            ImmutableArray<string> ids,
            ImmutableDictionary<string, object> tags,
            string id)
        {
            Filters = filters;
            Ids = ids;
            Tags = tags;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetVpcsFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// Set of values that are accepted for the given field.
        /// A VPC will be selected if any one of the given values matches.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetVpcsFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetVpcsFiltersResult
    {
        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set of values that are accepted for the given field.
        /// A VPC will be selected if any one of the given values matches.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetVpcsFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}

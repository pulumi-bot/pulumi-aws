// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static partial class GetSubnetIds
    {
        /// <summary>
        /// `aws.ec2.getSubnetIds` provides a set of ids for a vpc_id
        /// 
        /// This resource can be useful for getting back a set of subnet ids for a vpc.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/subnet_ids.html.markdown.
        /// </summary>
        public static Task<GetSubnetIdsResult> InvokeAsync(GetSubnetIdsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetIdsResult>("aws:ec2/getSubnetIds:getSubnetIds", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSubnetIdsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSubnetIdsFiltersArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetSubnetIdsFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSubnetIdsFiltersArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match
        /// a pair on the desired subnets.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID that you want to filter from.
        /// </summary>
        [Input("vpcId", required: true)]
        public string VpcId { get; set; } = null!;

        public GetSubnetIdsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSubnetIdsResult
    {
        public readonly ImmutableArray<Outputs.GetSubnetIdsFiltersResult> Filters;
        /// <summary>
        /// A set of all the subnet ids found. This data source will fail if none are found.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string VpcId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSubnetIdsResult(
            ImmutableArray<Outputs.GetSubnetIdsFiltersResult> filters,
            ImmutableArray<string> ids,
            ImmutableDictionary<string, object> tags,
            string vpcId,
            string id)
        {
            Filters = filters;
            Ids = ids;
            Tags = tags;
            VpcId = vpcId;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetSubnetIdsFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html).
        /// For example, if matching against tag `Name`, use:
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// Set of values that are accepted for the given field.
        /// Subnet IDs will be selected if any one of the given values match.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetSubnetIdsFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSubnetIdsFiltersResult
    {
        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html).
        /// For example, if matching against tag `Name`, use:
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set of values that are accepted for the given field.
        /// Subnet IDs will be selected if any one of the given values match.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetSubnetIdsFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}

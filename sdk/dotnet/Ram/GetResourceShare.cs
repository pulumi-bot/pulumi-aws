// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ram
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.ram.ResourceShare` Retrieve information about a RAM Resource Share.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ram_resource_share.html.markdown.
        /// </summary>
        public static Task<GetResourceShareResult> GetResourceShare(GetResourceShareArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceShareResult>("aws:ram/getResourceShare:getResourceShare", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetResourceShareArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetResourceShareFiltersArgs>? _filters;

        /// <summary>
        /// A filter used to scope the list e.g. by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
        /// </summary>
        public List<Inputs.GetResourceShareFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetResourceShareFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the tag key to filter on.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The owner of the resource share. Valid values are SELF or OTHER-ACCOUNTS
        /// </summary>
        [Input("resourceOwner", required: true)]
        public string ResourceOwner { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetResourceShareArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetResourceShareResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        public readonly string Arn;
        public readonly ImmutableArray<Outputs.GetResourceShareFiltersResult> Filters;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceOwner;
        /// <summary>
        /// The Status of the RAM share.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The Tags attached to the RAM share
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetResourceShareResult(
            string arn,
            ImmutableArray<Outputs.GetResourceShareFiltersResult> filters,
            string id,
            string name,
            string resourceOwner,
            string status,
            ImmutableDictionary<string, object> tags)
        {
            Arn = arn;
            Filters = filters;
            Id = id;
            Name = name;
            ResourceOwner = resourceOwner;
            Status = status;
            Tags = tags;
        }
    }

    namespace Inputs
    {

    public sealed class GetResourceShareFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the tag key to filter on.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// The value of the tag key.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetResourceShareFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetResourceShareFiltersResult
    {
        /// <summary>
        /// The name of the tag key to filter on.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the tag key.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetResourceShareFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}

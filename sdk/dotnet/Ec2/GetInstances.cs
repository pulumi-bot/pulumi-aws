// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static partial class GetInstances
    {
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("aws:ec2/getInstances:getInstances", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstancesFiltersArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are
        /// several valid keys, for a full reference, check out
        /// [describe-instances in the AWS CLI reference][1].
        /// </summary>
        public List<Inputs.GetInstancesFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstancesFiltersArgs>());
            set => _filters = value;
        }

        [Input("instanceStateNames")]
        private List<string>? _instanceStateNames;

        /// <summary>
        /// A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
        /// </summary>
        public List<string> InstanceStateNames
        {
            get => _instanceStateNames ?? (_instanceStateNames = new List<string>());
            set => _instanceStateNames = value;
        }

        [Input("instanceTags")]
        private Dictionary<string, object>? _instanceTags;

        /// <summary>
        /// A mapping of tags, each pair of which must
        /// exactly match a pair on desired instances.
        /// </summary>
        public Dictionary<string, object> InstanceTags
        {
            get => _instanceTags ?? (_instanceTags = new Dictionary<string, object>());
            set => _instanceTags = value;
        }

        public GetInstancesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly ImmutableArray<Outputs.GetInstancesFiltersResult> Filters;
        /// <summary>
        /// IDs of instances found through the filter
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> InstanceStateNames;
        public readonly ImmutableDictionary<string, object> InstanceTags;
        /// <summary>
        /// Private IP addresses of instances found through the filter
        /// </summary>
        public readonly ImmutableArray<string> PrivateIps;
        /// <summary>
        /// Public IP addresses of instances found through the filter
        /// </summary>
        public readonly ImmutableArray<string> PublicIps;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstancesResult(
            ImmutableArray<Outputs.GetInstancesFiltersResult> filters,
            ImmutableArray<string> ids,
            ImmutableArray<string> instanceStateNames,
            ImmutableDictionary<string, object> instanceTags,
            ImmutableArray<string> privateIps,
            ImmutableArray<string> publicIps,
            string id)
        {
            Filters = filters;
            Ids = ids;
            InstanceStateNames = instanceStateNames;
            InstanceTags = instanceTags;
            PrivateIps = privateIps;
            PublicIps = publicIps;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetInstancesFiltersArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetInstancesFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstancesFiltersResult
    {
        public readonly string Name;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetInstancesFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ResourceGroupsTaggingApi
{
    public static class GetResources
    {
        /// <summary>
        /// Provides details about resource tagging.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Get All Resource Tag Mappings
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.ResourceGroupsTaggingApi.GetResources.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Filter By Tag Key and Value
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.ResourceGroupsTaggingApi.GetResources.InvokeAsync(new Aws.ResourceGroupsTaggingApi.GetResourcesArgs
        ///         {
        ///             TagFilters = 
        ///             {
        ///                 new Aws.ResourceGroupsTaggingApi.Inputs.GetResourcesTagFilterArgs
        ///                 {
        ///                     Key = "tag-key",
        ///                     Values = 
        ///                     {
        ///                         "tag-value-1",
        ///                         "tag-value-2",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Filter By Resource Type
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.ResourceGroupsTaggingApi.GetResources.InvokeAsync(new Aws.ResourceGroupsTaggingApi.GetResourcesArgs
        ///         {
        ///             ResourceTypeFilters = 
        ///             {
        ///                 "ec2:instance",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourcesResult> InvokeAsync(GetResourcesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcesResult>("aws:resourcegroupstaggingapi/getResources:getResources", args ?? new GetResourcesArgs(), options.WithVersion());

        public static Output<GetResourcesResult> Apply(GetResourcesApplyArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetResourcesApplyArgs();
            return Pulumi.Output.All(
                args.ExcludeCompliantResources.Box(),
                args.IncludeComplianceDetails.Box(),
                args.ResourceArnLists.Box(),
                args.ResourceTypeFilters.Box(),
                args.TagFilters.Box()
            ).Apply(a => {
                    var args = new GetResourcesArgs();
                    a[0].Set(args, nameof(args.ExcludeCompliantResources));
                    a[1].Set(args, nameof(args.IncludeComplianceDetails));
                    a[2].Set(args, nameof(args.ResourceArnLists));
                    a[3].Set(args, nameof(args.ResourceTypeFilters));
                    a[4].Set(args, nameof(args.TagFilters));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetResourcesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether to exclude resources that are compliant with the tag policy. You can use this parameter only if the `include_compliance_details` argument is also set to `true`.
        /// </summary>
        [Input("excludeCompliantResources")]
        public bool? ExcludeCompliantResources { get; set; }

        /// <summary>
        /// Specifies whether to include details regarding the compliance with the effective tag policy.
        /// </summary>
        [Input("includeComplianceDetails")]
        public bool? IncludeComplianceDetails { get; set; }

        [Input("resourceArnLists")]
        private List<string>? _resourceArnLists;

        /// <summary>
        /// Specifies a list of ARNs of resources for which you want to retrieve tag data. Conflicts with `filter`.
        /// </summary>
        public List<string> ResourceArnLists
        {
            get => _resourceArnLists ?? (_resourceArnLists = new List<string>());
            set => _resourceArnLists = value;
        }

        [Input("resourceTypeFilters")]
        private List<string>? _resourceTypeFilters;

        /// <summary>
        /// The constraints on the resources that you want returned. The format of each resource type is `service:resourceType`. For example, specifying a resource type of `ec2` returns all Amazon EC2 resources (which includes EC2 instances). Specifying a resource type of `ec2:instance` returns only EC2 instances.
        /// </summary>
        public List<string> ResourceTypeFilters
        {
            get => _resourceTypeFilters ?? (_resourceTypeFilters = new List<string>());
            set => _resourceTypeFilters = value;
        }

        [Input("tagFilters")]
        private List<Inputs.GetResourcesTagFilterArgs>? _tagFilters;

        /// <summary>
        /// Specifies a list of Tag Filters (keys and values) to restrict the output to only those resources that have the specified tag and, if included, the specified value. See Tag Filter below. Conflicts with `resource_arn_list`.
        /// </summary>
        public List<Inputs.GetResourcesTagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new List<Inputs.GetResourcesTagFilterArgs>());
            set => _tagFilters = value;
        }

        public GetResourcesArgs()
        {
        }
    }

    public sealed class GetResourcesApplyArgs
    {
        /// <summary>
        /// Specifies whether to exclude resources that are compliant with the tag policy. You can use this parameter only if the `include_compliance_details` argument is also set to `true`.
        /// </summary>
        [Input("excludeCompliantResources")]
        public Input<bool>? ExcludeCompliantResources { get; set; }

        /// <summary>
        /// Specifies whether to include details regarding the compliance with the effective tag policy.
        /// </summary>
        [Input("includeComplianceDetails")]
        public Input<bool>? IncludeComplianceDetails { get; set; }

        [Input("resourceArnLists")]
        private InputList<string>? _resourceArnLists;

        /// <summary>
        /// Specifies a list of ARNs of resources for which you want to retrieve tag data. Conflicts with `filter`.
        /// </summary>
        public InputList<string> ResourceArnLists
        {
            get => _resourceArnLists ?? (_resourceArnLists = new InputList<string>());
            set => _resourceArnLists = value;
        }

        [Input("resourceTypeFilters")]
        private InputList<string>? _resourceTypeFilters;

        /// <summary>
        /// The constraints on the resources that you want returned. The format of each resource type is `service:resourceType`. For example, specifying a resource type of `ec2` returns all Amazon EC2 resources (which includes EC2 instances). Specifying a resource type of `ec2:instance` returns only EC2 instances.
        /// </summary>
        public InputList<string> ResourceTypeFilters
        {
            get => _resourceTypeFilters ?? (_resourceTypeFilters = new InputList<string>());
            set => _resourceTypeFilters = value;
        }

        [Input("tagFilters")]
        private InputList<Inputs.GetResourcesTagFilterArgs>? _tagFilters;

        /// <summary>
        /// Specifies a list of Tag Filters (keys and values) to restrict the output to only those resources that have the specified tag and, if included, the specified value. See Tag Filter below. Conflicts with `resource_arn_list`.
        /// </summary>
        public InputList<Inputs.GetResourcesTagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new InputList<Inputs.GetResourcesTagFilterArgs>());
            set => _tagFilters = value;
        }

        public GetResourcesApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcesResult
    {
        public readonly bool? ExcludeCompliantResources;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeComplianceDetails;
        public readonly ImmutableArray<string> ResourceArnLists;
        /// <summary>
        /// List of objects matching the search criteria.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourcesResourceTagMappingListResult> ResourceTagMappingLists;
        public readonly ImmutableArray<string> ResourceTypeFilters;
        public readonly ImmutableArray<Outputs.GetResourcesTagFilterResult> TagFilters;

        [OutputConstructor]
        private GetResourcesResult(
            bool? excludeCompliantResources,

            string id,

            bool? includeComplianceDetails,

            ImmutableArray<string> resourceArnLists,

            ImmutableArray<Outputs.GetResourcesResourceTagMappingListResult> resourceTagMappingLists,

            ImmutableArray<string> resourceTypeFilters,

            ImmutableArray<Outputs.GetResourcesTagFilterResult> tagFilters)
        {
            ExcludeCompliantResources = excludeCompliantResources;
            Id = id;
            IncludeComplianceDetails = includeComplianceDetails;
            ResourceArnLists = resourceArnLists;
            ResourceTagMappingLists = resourceTagMappingLists;
            ResourceTypeFilters = resourceTypeFilters;
            TagFilters = tagFilters;
        }
    }
}

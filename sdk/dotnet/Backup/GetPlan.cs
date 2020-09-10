// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup
{
    public static class GetPlan
    {
        public static Task<GetPlanResult> InvokeAsync(GetPlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPlanResult>("aws:backup/getPlan:getPlan", args ?? new GetPlanArgs(), options.WithVersion());
    }


    public sealed class GetPlanArgs : Pulumi.InvokeArgs
    {
        [Input("planId", required: true)]
        public string PlanId { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetPlanArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPlanResult
    {
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string PlanId;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string Version;

        [OutputConstructor]
        private GetPlanResult(
            string arn,

            string id,

            string name,

            string planId,

            ImmutableDictionary<string, string> tags,

            string version)
        {
            Arn = arn;
            Id = id;
            Name = name;
            PlanId = planId;
            Tags = tags;
            Version = version;
        }
    }
}

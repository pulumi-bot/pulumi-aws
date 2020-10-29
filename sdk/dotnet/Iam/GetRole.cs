// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    public static class GetRole
    {
        /// <summary>
        /// This data source can be used to fetch information about a specific
        /// IAM role. By using this data source, you can reference IAM role
        /// properties without having to hard code ARNs as input.
        /// </summary>
        public static Task<GetRoleResult> InvokeAsync(GetRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("aws:iam/getRole:getRole", args ?? new GetRoleArgs(), options.WithVersion());
    }


    public sealed class GetRoleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly IAM role name to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// The tags attached to the role.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetRoleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoleResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the role.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The policy document associated with the role.
        /// </summary>
        public readonly string AssumeRolePolicy;
        /// <summary>
        /// Creation date of the role in RFC 3339 format.
        /// </summary>
        public readonly string CreateDate;
        /// <summary>
        /// Description for the role.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Maximum session duration.
        /// </summary>
        public readonly int MaxSessionDuration;
        public readonly string Name;
        /// <summary>
        /// The path to the role.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The ARN of the policy that is used to set the permissions boundary for the role.
        /// </summary>
        public readonly string PermissionsBoundary;
        /// <summary>
        /// The tags attached to the role.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The stable and unique string identifying the role.
        /// </summary>
        public readonly string UniqueId;

        [OutputConstructor]
        private GetRoleResult(
            string arn,

            string assumeRolePolicy,

            string createDate,

            string description,

            string id,

            int maxSessionDuration,

            string name,

            string path,

            string permissionsBoundary,

            ImmutableDictionary<string, string> tags,

            string uniqueId)
        {
            Arn = arn;
            AssumeRolePolicy = assumeRolePolicy;
            CreateDate = createDate;
            Description = description;
            Id = id;
            MaxSessionDuration = maxSessionDuration;
            Name = name;
            Path = path;
            PermissionsBoundary = permissionsBoundary;
            Tags = tags;
            UniqueId = uniqueId;
        }
    }
}

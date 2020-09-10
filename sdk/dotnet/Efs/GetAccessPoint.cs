// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs
{
    public static class GetAccessPoint
    {
        public static Task<GetAccessPointResult> InvokeAsync(GetAccessPointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessPointResult>("aws:efs/getAccessPoint:getAccessPoint", args ?? new GetAccessPointArgs(), options.WithVersion());
    }


    public sealed class GetAccessPointArgs : Pulumi.InvokeArgs
    {
        [Input("accessPointId", required: true)]
        public string AccessPointId { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetAccessPointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessPointResult
    {
        public readonly string AccessPointId;
        public readonly string Arn;
        public readonly string FileSystemArn;
        public readonly string FileSystemId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OwnerId;
        public readonly ImmutableArray<Outputs.GetAccessPointPosixUserResult> PosixUsers;
        public readonly ImmutableArray<Outputs.GetAccessPointRootDirectoryResult> RootDirectories;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetAccessPointResult(
            string accessPointId,

            string arn,

            string fileSystemArn,

            string fileSystemId,

            string id,

            string ownerId,

            ImmutableArray<Outputs.GetAccessPointPosixUserResult> posixUsers,

            ImmutableArray<Outputs.GetAccessPointRootDirectoryResult> rootDirectories,

            ImmutableDictionary<string, string>? tags)
        {
            AccessPointId = accessPointId;
            Arn = arn;
            FileSystemArn = fileSystemArn;
            FileSystemId = fileSystemId;
            Id = id;
            OwnerId = ownerId;
            PosixUsers = posixUsers;
            RootDirectories = rootDirectories;
            Tags = tags;
        }
    }
}

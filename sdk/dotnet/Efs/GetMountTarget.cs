// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about an Elastic File System Mount Target (EFS).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/efs_mount_target.html.markdown.
        /// </summary>
        [Obsolete("Use GetMountTarget.InvokeAsync() instead")]
        public static Task<GetMountTargetResult> GetMountTarget(GetMountTargetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMountTargetResult>("aws:efs/getMountTarget:getMountTarget", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetMountTarget
    {
        /// <summary>
        /// Provides information about an Elastic File System Mount Target (EFS).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/efs_mount_target.html.markdown.
        /// </summary>
        public static Task<GetMountTargetResult> InvokeAsync(GetMountTargetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMountTargetResult>("aws:efs/getMountTarget:getMountTarget", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetMountTargetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the mount target that you want to have described
        /// </summary>
        [Input("mountTargetId", required: true)]
        public string MountTargetId { get; set; } = null!;

        public GetMountTargetArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetMountTargetResult
    {
        /// <summary>
        /// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        /// </summary>
        public readonly string DnsName;
        /// <summary>
        /// Amazon Resource Name of the file system for which the mount target is intended.
        /// </summary>
        public readonly string FileSystemArn;
        /// <summary>
        /// ID of the file system for which the mount target is intended.
        /// </summary>
        public readonly string FileSystemId;
        /// <summary>
        /// Address at which the file system may be mounted via the mount target.
        /// </summary>
        public readonly string IpAddress;
        public readonly string MountTargetId;
        /// <summary>
        /// The ID of the network interface that Amazon EFS created when it created the mount target.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// List of VPC security group IDs attached to the mount target.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// ID of the mount target's subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetMountTargetResult(
            string dnsName,
            string fileSystemArn,
            string fileSystemId,
            string ipAddress,
            string mountTargetId,
            string networkInterfaceId,
            ImmutableArray<string> securityGroups,
            string subnetId,
            string id)
        {
            DnsName = dnsName;
            FileSystemArn = fileSystemArn;
            FileSystemId = fileSystemId;
            IpAddress = ipAddress;
            MountTargetId = mountTargetId;
            NetworkInterfaceId = networkInterfaceId;
            SecurityGroups = securityGroups;
            SubnetId = subnetId;
            Id = id;
        }
    }
}

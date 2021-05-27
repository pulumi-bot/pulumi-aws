// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    public static class GetAddon
    {
        /// <summary>
        /// Retrieve information about an EKS add-on.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Eks.GetAddon.InvokeAsync(new Aws.Eks.GetAddonArgs
        ///         {
        ///             AddonName = "vpc-cni",
        ///             ClusterName = aws_eks_cluster.Example.Name,
        ///         }));
        ///         this.EksAddonOutputs = aws_eks_addon.Example;
        ///     }
        /// 
        ///     [Output("eksAddonOutputs")]
        ///     public Output&lt;string&gt; EksAddonOutputs { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAddonResult> InvokeAsync(GetAddonArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAddonResult>("aws:eks/getAddon:getAddon", args ?? new GetAddonArgs(), options.WithVersion());

        public static Output<GetAddonResult> Invoke(GetAddonOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.AddonName.Box(),
                args.ClusterName.Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetAddonArgs();
                    a[0].Set(args, nameof(args.AddonName));
                    a[1].Set(args, nameof(args.ClusterName));
                    a[2].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAddonArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the EKS add-on. The name must match one of
        /// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
        /// </summary>
        [Input("addonName", required: true)]
        public string AddonName { get; set; } = null!;

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetAddonArgs()
        {
        }
    }

    public sealed class GetAddonOutputArgs
    {
        /// <summary>
        /// Name of the EKS add-on. The name must match one of
        /// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
        /// </summary>
        [Input("addonName", required: true)]
        public Input<string> AddonName { get; set; } = null!;

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetAddonOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAddonResult
    {
        public readonly string AddonName;
        /// <summary>
        /// The version of EKS add-on.
        /// </summary>
        public readonly string AddonVersion;
        /// <summary>
        /// Amazon Resource Name (ARN) of the EKS add-on.
        /// </summary>
        public readonly string Arn;
        public readonly string ClusterName;
        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
        /// </summary>
        public readonly string ModifiedAt;
        /// <summary>
        /// ARN of IAM role used for EKS add-on. If value is empty -
        /// then add-on uses the IAM role assigned to the EKS Cluster node.
        /// </summary>
        public readonly string ServiceAccountRoleArn;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAddonResult(
            string addonName,

            string addonVersion,

            string arn,

            string clusterName,

            string createdAt,

            string id,

            string modifiedAt,

            string serviceAccountRoleArn,

            ImmutableDictionary<string, string> tags)
        {
            AddonName = addonName;
            AddonVersion = addonVersion;
            Arn = arn;
            ClusterName = clusterName;
            CreatedAt = createdAt;
            Id = id;
            ModifiedAt = modifiedAt;
            ServiceAccountRoleArn = serviceAccountRoleArn;
            Tags = tags;
        }
    }
}

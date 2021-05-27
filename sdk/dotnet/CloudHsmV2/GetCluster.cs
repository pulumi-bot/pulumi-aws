// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudHsmV2
{
    public static class GetCluster
    {
        /// <summary>
        /// Use this data source to get information about a CloudHSM v2 cluster
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
        ///         var cluster = Output.Create(Aws.CloudHsmV2.GetCluster.InvokeAsync(new Aws.CloudHsmV2.GetClusterArgs
        ///         {
        ///             ClusterId = "cluster-testclusterid",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:cloudhsmv2/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithVersion());

        public static Output<GetClusterResult> Invoke(GetClusterOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.ClusterId.Box(),
                args.ClusterState.Box()
            ).Apply(a => {
                    var args = new GetClusterArgs();
                    a[0].Set(args, nameof(args.ClusterId));
                    a[1].Set(args, nameof(args.ClusterState));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of Cloud HSM v2 cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The state of the cluster to be found.
        /// </summary>
        [Input("clusterState")]
        public string? ClusterState { get; set; }

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterOutputArgs
    {
        /// <summary>
        /// The id of Cloud HSM v2 cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The state of the cluster to be found.
        /// </summary>
        [Input("clusterState")]
        public Input<string>? ClusterState { get; set; }

        public GetClusterOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The list of cluster certificates.
        /// * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
        /// * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state.
        /// * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
        /// * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
        /// * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
        /// The number of available cluster certificates may vary depending on state of the cluster.
        /// </summary>
        public readonly Outputs.GetClusterClusterCertificatesResult ClusterCertificates;
        public readonly string ClusterId;
        public readonly string ClusterState;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the security group associated with the CloudHSM cluster.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The IDs of subnets in which cluster operates.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The id of the VPC that the CloudHSM cluster resides in.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetClusterResult(
            Outputs.GetClusterClusterCertificatesResult clusterCertificates,

            string clusterId,

            string clusterState,

            string id,

            string securityGroupId,

            ImmutableArray<string> subnetIds,

            string vpcId)
        {
            ClusterCertificates = clusterCertificates;
            ClusterId = clusterId;
            ClusterState = clusterState;
            Id = id;
            SecurityGroupId = securityGroupId;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    public static partial class Invokes
    {
        /// <summary>
        /// Retrieve information about an EKS Cluster.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/eks_cluster.html.markdown.
        /// </summary>
        public static Task<GetClusterResult> GetCluster(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:eks/getCluster:getCluster", args, options.WithVersion());
    }

    public sealed class GetClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cluster
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Nested attribute containing `certificate-authority-data` for your cluster.
        /// </summary>
        public readonly Outputs.GetClusterCertificateAuthorityResult CertificateAuthority;
        /// <summary>
        /// The Unix epoch time stamp in seconds for when the cluster was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The enabled control plane logs.
        /// </summary>
        public readonly ImmutableArray<string> EnabledClusterLogTypes;
        /// <summary>
        /// The endpoint for your Kubernetes API server.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. For an example using this information to enable IAM Roles for Service Accounts, see the [`aws.eks.Cluster` resource documentation](https://www.terraform.io/docs/providers/aws/r/eks_cluster.html).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterIdentitiesResult> Identities;
        public readonly string Name;
        /// <summary>
        /// The platform version for the cluster.
        /// </summary>
        public readonly string PlatformVersion;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The Kubernetes server version for the cluster.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// Nested attribute containing VPC configuration for the cluster.
        /// </summary>
        public readonly Outputs.GetClusterVpcConfigResult VpcConfig;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterResult(
            string arn,
            Outputs.GetClusterCertificateAuthorityResult certificateAuthority,
            string createdAt,
            ImmutableArray<string> enabledClusterLogTypes,
            string endpoint,
            ImmutableArray<Outputs.GetClusterIdentitiesResult> identities,
            string name,
            string platformVersion,
            string roleArn,
            string status,
            ImmutableDictionary<string, string> tags,
            string version,
            Outputs.GetClusterVpcConfigResult vpcConfig,
            string id)
        {
            Arn = arn;
            CertificateAuthority = certificateAuthority;
            CreatedAt = createdAt;
            EnabledClusterLogTypes = enabledClusterLogTypes;
            Endpoint = endpoint;
            Identities = identities;
            Name = name;
            PlatformVersion = platformVersion;
            RoleArn = roleArn;
            Status = status;
            Tags = tags;
            Version = version;
            VpcConfig = vpcConfig;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClusterCertificateAuthorityResult
    {
        /// <summary>
        /// The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
        /// </summary>
        public readonly string Data;

        [OutputConstructor]
        private GetClusterCertificateAuthorityResult(string data)
        {
            Data = data;
        }
    }

    [OutputType]
    public sealed class GetClusterIdentitiesOidcsResult
    {
        /// <summary>
        /// Issuer URL for the OpenID Connect identity provider.
        /// </summary>
        public readonly string Issuer;

        [OutputConstructor]
        private GetClusterIdentitiesOidcsResult(string issuer)
        {
            Issuer = issuer;
        }
    }

    [OutputType]
    public sealed class GetClusterIdentitiesResult
    {
        /// <summary>
        /// Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster.
        /// </summary>
        public readonly ImmutableArray<GetClusterIdentitiesOidcsResult> Oidcs;

        [OutputConstructor]
        private GetClusterIdentitiesResult(ImmutableArray<GetClusterIdentitiesOidcsResult> oidcs)
        {
            Oidcs = oidcs;
        }
    }

    [OutputType]
    public sealed class GetClusterVpcConfigResult
    {
        /// <summary>
        /// Indicates whether or not the Amazon EKS private API server endpoint is enabled.
        /// </summary>
        public readonly bool EndpointPrivateAccess;
        /// <summary>
        /// Indicates whether or not the Amazon EKS public API server endpoint is enabled.
        /// </summary>
        public readonly bool EndpointPublicAccess;
        /// <summary>
        /// List of security group IDs
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// List of subnet IDs
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The VPC associated with your cluster.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetClusterVpcConfigResult(
            bool endpointPrivateAccess,
            bool endpointPublicAccess,
            ImmutableArray<string> securityGroupIds,
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            EndpointPrivateAccess = endpointPrivateAccess;
            EndpointPublicAccess = endpointPublicAccess;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
    }
}

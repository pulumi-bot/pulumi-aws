// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch
{
    public static class GetDomain
    {
        /// <summary>
        /// Use this data source to get information about an Elasticsearch Domain
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
        ///         var myDomain = Output.Create(Aws.ElasticSearch.GetDomain.InvokeAsync(new Aws.ElasticSearch.GetDomainArgs
        ///         {
        ///             DomainName = "my-domain-name",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws:elasticsearch/getDomain:getDomain", args ?? new GetDomainArgs(), options.WithVersion());
    }


    public sealed class GetDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tags assigned to the domain.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetDomainArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// The policy document attached to the domain.
        /// </summary>
        public readonly string AccessPolicies;
        /// <summary>
        /// Key-value string pairs to specify advanced configuration options.
        /// </summary>
        public readonly ImmutableDictionary<string, object> AdvancedOptions;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the domain.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Cluster configuration of the domain.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainClusterConfigResult> ClusterConfigs;
        /// <summary>
        /// Domain Amazon Cognito Authentication options for Kibana.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainCognitoOptionResult> CognitoOptions;
        /// <summary>
        /// Status of the creation of the domain.
        /// </summary>
        public readonly bool Created;
        /// <summary>
        /// Status of the deletion of the domain.
        /// </summary>
        public readonly bool Deleted;
        /// <summary>
        /// Unique identifier for the domain.
        /// </summary>
        public readonly string DomainId;
        public readonly string DomainName;
        /// <summary>
        /// EBS Options for the instances in the domain.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainEbsOptionResult> EbsOptions;
        /// <summary>
        /// ElasticSearch version for the domain.
        /// </summary>
        public readonly string ElasticsearchVersion;
        /// <summary>
        /// Domain encryption at rest related options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainEncryptionAtRestResult> EncryptionAtRests;
        /// <summary>
        /// Domain-specific endpoint used to submit index, search, and data upload requests.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Domain-specific endpoint used to access the Kibana application.
        /// </summary>
        public readonly string KibanaEndpoint;
        /// <summary>
        /// Domain log publishing related options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainLogPublishingOptionResult> LogPublishingOptions;
        /// <summary>
        /// Domain in transit encryption related options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainNodeToNodeEncryptionResult> NodeToNodeEncryptions;
        /// <summary>
        /// Status of a configuration change in the domain.
        /// * `snapshot_options` – Domain snapshot related options.
        /// </summary>
        public readonly string Processing;
        public readonly ImmutableArray<Outputs.GetDomainSnapshotOptionResult> SnapshotOptions;
        /// <summary>
        /// The tags assigned to the domain.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// VPC Options for private Elasticsearch domains.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainVpcOptionResult> VpcOptions;

        [OutputConstructor]
        private GetDomainResult(
            string accessPolicies,

            ImmutableDictionary<string, object> advancedOptions,

            string arn,

            ImmutableArray<Outputs.GetDomainClusterConfigResult> clusterConfigs,

            ImmutableArray<Outputs.GetDomainCognitoOptionResult> cognitoOptions,

            bool created,

            bool deleted,

            string domainId,

            string domainName,

            ImmutableArray<Outputs.GetDomainEbsOptionResult> ebsOptions,

            string elasticsearchVersion,

            ImmutableArray<Outputs.GetDomainEncryptionAtRestResult> encryptionAtRests,

            string endpoint,

            string id,

            string kibanaEndpoint,

            ImmutableArray<Outputs.GetDomainLogPublishingOptionResult> logPublishingOptions,

            ImmutableArray<Outputs.GetDomainNodeToNodeEncryptionResult> nodeToNodeEncryptions,

            string processing,

            ImmutableArray<Outputs.GetDomainSnapshotOptionResult> snapshotOptions,

            ImmutableDictionary<string, object> tags,

            ImmutableArray<Outputs.GetDomainVpcOptionResult> vpcOptions)
        {
            AccessPolicies = accessPolicies;
            AdvancedOptions = advancedOptions;
            Arn = arn;
            ClusterConfigs = clusterConfigs;
            CognitoOptions = cognitoOptions;
            Created = created;
            Deleted = deleted;
            DomainId = domainId;
            DomainName = domainName;
            EbsOptions = ebsOptions;
            ElasticsearchVersion = elasticsearchVersion;
            EncryptionAtRests = encryptionAtRests;
            Endpoint = endpoint;
            Id = id;
            KibanaEndpoint = kibanaEndpoint;
            LogPublishingOptions = logPublishingOptions;
            NodeToNodeEncryptions = nodeToNodeEncryptions;
            Processing = processing;
            SnapshotOptions = snapshotOptions;
            Tags = tags;
            VpcOptions = vpcOptions;
        }
    }
}

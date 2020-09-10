// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB
{
    public static class GetTable
    {
        public static Task<GetTableResult> InvokeAsync(GetTableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTableResult>("aws:dynamodb/getTable:getTable", args ?? new GetTableArgs(), options.WithVersion());
    }


    public sealed class GetTableArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("serverSideEncryption")]
        public Inputs.GetTableServerSideEncryptionArgs? ServerSideEncryption { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetTableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTableResult
    {
        public readonly string Arn;
        public readonly ImmutableArray<Outputs.GetTableAttributeResult> Attributes;
        public readonly string BillingMode;
        public readonly ImmutableArray<Outputs.GetTableGlobalSecondaryIndexResult> GlobalSecondaryIndexes;
        public readonly string HashKey;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetTableLocalSecondaryIndexResult> LocalSecondaryIndexes;
        public readonly string Name;
        public readonly Outputs.GetTablePointInTimeRecoveryResult PointInTimeRecovery;
        public readonly string RangeKey;
        public readonly int ReadCapacity;
        public readonly ImmutableArray<Outputs.GetTableReplicaResult> Replicas;
        public readonly Outputs.GetTableServerSideEncryptionResult ServerSideEncryption;
        public readonly string StreamArn;
        public readonly bool StreamEnabled;
        public readonly string StreamLabel;
        public readonly string StreamViewType;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly Outputs.GetTableTtlResult Ttl;
        public readonly int WriteCapacity;

        [OutputConstructor]
        private GetTableResult(
            string arn,

            ImmutableArray<Outputs.GetTableAttributeResult> attributes,

            string billingMode,

            ImmutableArray<Outputs.GetTableGlobalSecondaryIndexResult> globalSecondaryIndexes,

            string hashKey,

            string id,

            ImmutableArray<Outputs.GetTableLocalSecondaryIndexResult> localSecondaryIndexes,

            string name,

            Outputs.GetTablePointInTimeRecoveryResult pointInTimeRecovery,

            string rangeKey,

            int readCapacity,

            ImmutableArray<Outputs.GetTableReplicaResult> replicas,

            Outputs.GetTableServerSideEncryptionResult serverSideEncryption,

            string streamArn,

            bool streamEnabled,

            string streamLabel,

            string streamViewType,

            ImmutableDictionary<string, string> tags,

            Outputs.GetTableTtlResult ttl,

            int writeCapacity)
        {
            Arn = arn;
            Attributes = attributes;
            BillingMode = billingMode;
            GlobalSecondaryIndexes = globalSecondaryIndexes;
            HashKey = hashKey;
            Id = id;
            LocalSecondaryIndexes = localSecondaryIndexes;
            Name = name;
            PointInTimeRecovery = pointInTimeRecovery;
            RangeKey = rangeKey;
            ReadCapacity = readCapacity;
            Replicas = replicas;
            ServerSideEncryption = serverSideEncryption;
            StreamArn = streamArn;
            StreamEnabled = streamEnabled;
            StreamLabel = streamLabel;
            StreamViewType = streamViewType;
            Tags = tags;
            Ttl = ttl;
            WriteCapacity = writeCapacity;
        }
    }
}

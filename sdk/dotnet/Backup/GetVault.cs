// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup
{
    public static class GetVault
    {
        public static Task<GetVaultResult> InvokeAsync(GetVaultArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVaultResult>("aws:backup/getVault:getVault", args ?? new GetVaultArgs(), options.WithVersion());
    }


    public sealed class GetVaultArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetVaultArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVaultResult
    {
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KmsKeyArn;
        public readonly string Name;
        public readonly int RecoveryPoints;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetVaultResult(
            string arn,

            string id,

            string kmsKeyArn,

            string name,

            int recoveryPoints,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            Id = id;
            KmsKeyArn = kmsKeyArn;
            Name = name;
            RecoveryPoints = recoveryPoints;
            Tags = tags;
        }
    }
}

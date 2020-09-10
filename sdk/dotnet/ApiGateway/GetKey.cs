// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    public static class GetKey
    {
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("aws:apigateway/getKey:getKey", args ?? new GetKeyArgs(), options.WithVersion());
    }


    public sealed class GetKeyArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        public readonly string CreatedDate;
        public readonly string Description;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string LastUpdatedDate;
        public readonly string Name;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string Value;

        [OutputConstructor]
        private GetKeyResult(
            string createdDate,

            string description,

            bool enabled,

            string id,

            string lastUpdatedDate,

            string name,

            ImmutableDictionary<string, string> tags,

            string value)
        {
            CreatedDate = createdDate;
            Description = description;
            Enabled = enabled;
            Id = id;
            LastUpdatedDate = lastUpdatedDate;
            Name = name;
            Tags = tags;
            Value = value;
        }
    }
}

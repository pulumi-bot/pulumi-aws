// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    public static class GetRestApi
    {
        /// <summary>
        /// Use this data source to get the id and root_resource_id of a REST API in
        /// API Gateway. To fetch the REST API you must provide a name to match against. 
        /// As there is no unique name constraint on REST APIs this data source will 
        /// error if there is more than one match.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRestApiResult> InvokeAsync(GetRestApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRestApiResult>("aws:apigateway/getRestApi:getRestApi", args ?? new GetRestApiArgs(), options.WithVersion());
    }


    public sealed class GetRestApiArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetRestApiArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRestApiResult
    {
        /// <summary>
        /// The source of the API key for requests.
        /// </summary>
        public readonly string ApiKeySource;
        /// <summary>
        /// The ARN of the REST API.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The list of binary media types supported by the REST API.
        /// </summary>
        public readonly ImmutableArray<string> BinaryMediaTypes;
        /// <summary>
        /// The description of the REST API.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The endpoint configuration of this RestApi showing the endpoint types of the API.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRestApiEndpointConfigurationResult> EndpointConfigurations;
        /// <summary>
        /// The execution ARN part to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn` when allowing API Gateway to invoke a Lambda function, e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
        /// </summary>
        public readonly string ExecutionArn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Minimum response size to compress for the REST API.
        /// </summary>
        public readonly int MinimumCompressionSize;
        public readonly string Name;
        /// <summary>
        /// JSON formatted policy document that controls access to the API Gateway.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
        /// </summary>
        public readonly string RootResourceId;
        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetRestApiResult(
            string apiKeySource,

            string arn,

            ImmutableArray<string> binaryMediaTypes,

            string description,

            ImmutableArray<Outputs.GetRestApiEndpointConfigurationResult> endpointConfigurations,

            string executionArn,

            string id,

            int minimumCompressionSize,

            string name,

            string policy,

            string rootResourceId,

            ImmutableDictionary<string, object> tags)
        {
            ApiKeySource = apiKeySource;
            Arn = arn;
            BinaryMediaTypes = binaryMediaTypes;
            Description = description;
            EndpointConfigurations = endpointConfigurations;
            ExecutionArn = executionArn;
            Id = id;
            MinimumCompressionSize = minimumCompressionSize;
            Name = name;
            Policy = policy;
            RootResourceId = rootResourceId;
            Tags = tags;
        }
    }
}

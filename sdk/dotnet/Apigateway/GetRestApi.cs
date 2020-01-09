// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the id and root_resource_id of a REST API in
        /// API Gateway. To fetch the REST API you must provide a name to match against. 
        /// As there is no unique name constraint on REST APIs this data source will 
        /// error if there is more than one match.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_rest_api.html.markdown.
        /// </summary>
        public static Task<GetRestApiResult> GetRestApi(GetRestApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRestApiResult>("aws:apigateway/getRestApi:getRestApi", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRestApiArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the REST API to look up. If no REST API is found with this name, an error will be returned. 
        /// If multiple REST APIs are found with this name, an error will be returned.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRestApiArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRestApiResult
    {
        public readonly string Name;
        /// <summary>
        /// Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
        /// </summary>
        public readonly string RootResourceId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRestApiResult(
            string name,
            string rootResourceId,
            string id)
        {
            Name = name;
            RootResourceId = rootResourceId;
            Id = id;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    public static class GetUserPools
    {
        /// <summary>
        /// Use this data source to get a list of cognito user pools.
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
        ///         var selectedRestApi = Output.Create(Aws.ApiGateway.GetRestApi.InvokeAsync(new Aws.ApiGateway.GetRestApiArgs
        ///         {
        ///             Name = var.Api_gateway_name,
        ///         }));
        ///         var selectedUserPools = Output.Create(Aws.Cognito.GetUserPools.InvokeAsync(new Aws.Cognito.GetUserPoolsArgs
        ///         {
        ///             Name = var.Cognito_user_pool_name,
        ///         }));
        ///         var cognito = new Aws.ApiGateway.Authorizer("cognito", new Aws.ApiGateway.AuthorizerArgs
        ///         {
        ///             ProviderArns = selectedUserPools.Apply(selectedUserPools =&gt; selectedUserPools.Arns),
        ///             RestApi = selectedRestApi.Apply(selectedRestApi =&gt; selectedRestApi.Id),
        ///             Type = "COGNITO_USER_POOLS",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserPoolsResult> InvokeAsync(GetUserPoolsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserPoolsResult>("aws:cognito/getUserPools:getUserPools", args ?? new GetUserPoolsArgs(), options.WithVersion());
    }


    public sealed class GetUserPoolsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetUserPoolsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserPoolsResult
    {
        public readonly ImmutableArray<string> Arns;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of cognito user pool ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string Name;

        [OutputConstructor]
        private GetUserPoolsResult(
            ImmutableArray<string> arns,

            string id,

            ImmutableArray<string> ids,

            string name)
        {
            Arns = arns;
            Id = id;
            Ids = ids;
            Name = name;
        }
    }
}

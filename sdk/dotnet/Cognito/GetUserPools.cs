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
        ///             Name = @var.Api_gateway_name,
        ///         }));
        ///         var selectedUserPools = Output.Create(Aws.Cognito.GetUserPools.InvokeAsync(new Aws.Cognito.GetUserPoolsArgs
        ///         {
        ///             Name = @var.Cognito_user_pool_name,
        ///         }));
        ///         var cognito = new Aws.ApiGateway.Authorizer("cognito", new Aws.ApiGateway.AuthorizerArgs
        ///         {
        ///             Type = "COGNITO_USER_POOLS",
        ///             RestApi = selectedRestApi.Apply(selectedRestApi =&gt; selectedRestApi.Id),
        ///             ProviderArns = selectedUserPools.Apply(selectedUserPools =&gt; selectedUserPools.Arns),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserPoolsResult> InvokeAsync(GetUserPoolsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserPoolsResult>("aws:cognito/getUserPools:getUserPools", args ?? new GetUserPoolsArgs(), options.WithVersion());

        public static Output<GetUserPoolsResult> Invoke(GetUserPoolsOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetUserPoolsArgs();
                    a[0].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
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

    public sealed class GetUserPoolsOutputArgs
    {
        /// <summary>
        /// Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetUserPoolsOutputArgs()
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

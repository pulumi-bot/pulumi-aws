// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    public static class GetResource
    {
        /// <summary>
        /// Use this data source to get the id of a Resource in API Gateway.
        /// To fetch the Resource, you must provide the REST API id as well as the full path.  
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
        ///         var myRestApi = Output.Create(Aws.ApiGateway.GetRestApi.InvokeAsync(new Aws.ApiGateway.GetRestApiArgs
        ///         {
        ///             Name = "my-rest-api",
        ///         }));
        ///         var myResource = myRestApi.Apply(myRestApi =&gt; Output.Create(Aws.ApiGateway.GetResource.InvokeAsync(new Aws.ApiGateway.GetResourceArgs
        ///         {
        ///             RestApiId = myRestApi.Id,
        ///             Path = "/endpoint/path",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourceResult> InvokeAsync(GetResourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceResult>("aws:apigateway/getResource:getResource", args ?? new GetResourceArgs(), options.WithVersion());

        public static Output<GetResourceResult> Apply(GetResourceApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Path.Box(),
                args.RestApiId.Box()
            ).Apply(a => {
                    var args = new GetResourceArgs();
                    a[0].Set(args, nameof(args.Path));
                    a[1].Set(args, nameof(args.RestApiId));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetResourceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full path of the resource.  If no path is found, an error will be returned.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// The REST API id that owns the resource. If no REST API is found, an error will be returned.
        /// </summary>
        [Input("restApiId", required: true)]
        public string RestApiId { get; set; } = null!;

        public GetResourceArgs()
        {
        }
    }

    public sealed class GetResourceApplyArgs
    {
        /// <summary>
        /// The full path of the resource.  If no path is found, an error will be returned.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The REST API id that owns the resource. If no REST API is found, an error will be returned.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public GetResourceApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set to the ID of the parent Resource.
        /// </summary>
        public readonly string ParentId;
        public readonly string Path;
        /// <summary>
        /// Set to the path relative to the parent Resource.
        /// </summary>
        public readonly string PathPart;
        public readonly string RestApiId;

        [OutputConstructor]
        private GetResourceResult(
            string id,

            string parentId,

            string path,

            string pathPart,

            string restApiId)
        {
            Id = id;
            ParentId = parentId;
            Path = path;
            PathPart = pathPart;
            RestApiId = restApiId;
        }
    }
}

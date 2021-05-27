// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetRegion
    {
        /// <summary>
        /// `aws.getRegion` provides details about a specific AWS region.
        /// 
        /// As well as validating a given region name this resource can be used to
        /// discover the name of the region configured within the provider. The latter
        /// can be useful in a child module which is inheriting an AWS provider
        /// configuration from its parent module.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how the resource might be used to obtain
        /// the name of the AWS region configured on the provider.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var current = Output.Create(Aws.GetRegion.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegionResult> InvokeAsync(GetRegionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionResult>("aws:index/getRegion:getRegion", args ?? new GetRegionArgs(), options.WithVersion());

        public static Output<GetRegionResult> Invoke(GetRegionOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetRegionOutputArgs();
            return Pulumi.Output.All(
                args.Endpoint.Box(),
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetRegionArgs();
                    a[0].Set(args, nameof(args.Endpoint));
                    a[1].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetRegionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The EC2 endpoint of the region to select.
        /// </summary>
        [Input("endpoint")]
        public string? Endpoint { get; set; }

        /// <summary>
        /// The full name of the region to select.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetRegionArgs()
        {
        }
    }

    public sealed class GetRegionOutputArgs
    {
        /// <summary>
        /// The EC2 endpoint of the region to select.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The full name of the region to select.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetRegionOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionResult
    {
        /// <summary>
        /// The region's description in this format: "Location (Region name)".
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The EC2 endpoint for the selected region.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the selected region.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetRegionResult(
            string description,

            string endpoint,

            string id,

            string name)
        {
            Description = description;
            Endpoint = endpoint;
            Id = id;
            Name = name;
        }
    }
}

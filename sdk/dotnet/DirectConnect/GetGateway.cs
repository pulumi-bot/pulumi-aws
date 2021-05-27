// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectConnect
{
    public static class GetGateway
    {
        /// <summary>
        /// Retrieve information about a Direct Connect Gateway.
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
        ///         var example = Output.Create(Aws.DirectConnect.GetGateway.InvokeAsync(new Aws.DirectConnect.GetGatewayArgs
        ///         {
        ///             Name = "example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGatewayResult> InvokeAsync(GetGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayResult>("aws:directconnect/getGateway:getGateway", args ?? new GetGatewayArgs(), options.WithVersion());

        public static Output<GetGatewayResult> Invoke(GetGatewayOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetGatewayArgs();
                    a[0].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the gateway to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetGatewayArgs()
        {
        }
    }

    public sealed class GetGatewayOutputArgs
    {
        /// <summary>
        /// The name of the gateway to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetGatewayOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewayResult
    {
        /// <summary>
        /// The ASN on the Amazon side of the connection.
        /// </summary>
        public readonly string AmazonSideAsn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// AWS Account ID of the gateway.
        /// </summary>
        public readonly string OwnerAccountId;

        [OutputConstructor]
        private GetGatewayResult(
            string amazonSideAsn,

            string id,

            string name,

            string ownerAccountId)
        {
            AmazonSideAsn = amazonSideAsn;
            Id = id;
            Name = name;
            OwnerAccountId = ownerAccountId;
        }
    }
}

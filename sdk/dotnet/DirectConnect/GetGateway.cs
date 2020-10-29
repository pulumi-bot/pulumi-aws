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
        /// </summary>
        public static Task<GetGatewayResult> InvokeAsync(GetGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayResult>("aws:directconnect/getGateway:getGateway", args ?? new GetGatewayArgs(), options.WithVersion());
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

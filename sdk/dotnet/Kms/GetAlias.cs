// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kms
{
    public static class GetAlias
    {
        public static Task<GetAliasResult> InvokeAsync(GetAliasArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAliasResult>("aws:kms/getAlias:getAlias", args ?? new GetAliasArgs(), options.WithVersion());
    }


    public sealed class GetAliasArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetAliasArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAliasResult
    {
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string TargetKeyArn;
        public readonly string TargetKeyId;

        [OutputConstructor]
        private GetAliasResult(
            string arn,

            string id,

            string name,

            string targetKeyArn,

            string targetKeyId)
        {
            Arn = arn;
            Id = id;
            Name = name;
            TargetKeyArn = targetKeyArn;
            TargetKeyId = targetKeyId;
        }
    }
}

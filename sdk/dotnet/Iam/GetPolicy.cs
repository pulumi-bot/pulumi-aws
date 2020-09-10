// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    public static class GetPolicy
    {
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("aws:iam/getPolicy:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        public readonly string Arn;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Path;
        public readonly string Policy;

        [OutputConstructor]
        private GetPolicyResult(
            string arn,

            string description,

            string id,

            string name,

            string path,

            string policy)
        {
            Arn = arn;
            Description = description;
            Id = id;
            Name = name;
            Path = path;
            Policy = policy;
        }
    }
}

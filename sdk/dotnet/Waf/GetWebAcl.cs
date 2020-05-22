// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    public static class GetWebAcl
    {
        /// <summary>
        /// `aws.waf.WebAcl` Retrieves a WAF Web ACL Resource Id.
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
        ///         var example = Output.Create(Aws.Waf.GetWebAcl.InvokeAsync(new Aws.Waf.GetWebAclArgs
        ///         {
        ///             Name = "tfWAFWebACL",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWebAclResult> InvokeAsync(GetWebAclArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebAclResult>("aws:waf/getWebAcl:getWebAcl", args ?? new GetWebAclArgs(), options.WithVersion());
    }


    public sealed class GetWebAclArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the WAF Web ACL.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetWebAclArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebAclResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetWebAclResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Inspector
{
    public static partial class Invokes
    {
        /// <summary>
        /// The AWS Inspector Rules Packages data source allows access to the list of AWS
        /// Inspector Rules Packages which can be used by AWS Inspector within the region
        /// configured in the provider.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/inspector_rules_packages.html.markdown.
        /// </summary>
        public static Task<GetRulesPackagesResult> GetRulesPackages(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRulesPackagesResult>("aws:inspector/getRulesPackages:getRulesPackages", InvokeArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetRulesPackagesResult
    {
        /// <summary>
        /// A list of the AWS Inspector Rules Packages arns available in the AWS region.
        /// </summary>
        public readonly ImmutableArray<string> Arns;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRulesPackagesResult(
            ImmutableArray<string> arns,
            string id)
        {
            Arns = arns;
            Id = id;
        }
    }
}

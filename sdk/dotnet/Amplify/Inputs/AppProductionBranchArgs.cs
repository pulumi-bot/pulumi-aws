// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Amplify.Inputs
{

    public sealed class AppProductionBranchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The branch name for the production branch.
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        /// <summary>
        /// The last deploy time of the production branch.
        /// </summary>
        [Input("lastDeployTime")]
        public Input<string>? LastDeployTime { get; set; }

        /// <summary>
        /// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The thumbnail URL for the production branch.
        /// </summary>
        [Input("thumbnailUrl")]
        public Input<string>? ThumbnailUrl { get; set; }

        public AppProductionBranchArgs()
        {
        }
    }
}

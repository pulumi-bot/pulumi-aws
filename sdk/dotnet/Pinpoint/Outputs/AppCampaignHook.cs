// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pinpoint.Outputs
{

    [OutputType]
    public sealed class AppCampaignHook
    {
        public readonly string? LambdaFunctionName;
        public readonly string? Mode;
        public readonly string? WebUrl;

        [OutputConstructor]
        private AppCampaignHook(
            string? lambdaFunctionName,

            string? mode,

            string? webUrl)
        {
            LambdaFunctionName = lambdaFunctionName;
            Mode = mode;
            WebUrl = webUrl;
        }
    }
}

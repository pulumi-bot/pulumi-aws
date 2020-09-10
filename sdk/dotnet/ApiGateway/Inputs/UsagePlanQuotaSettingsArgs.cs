// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway.Inputs
{

    public sealed class UsagePlanQuotaSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("limit", required: true)]
        public Input<int> Limit { get; set; } = null!;

        [Input("offset")]
        public Input<int>? Offset { get; set; }

        [Input("period", required: true)]
        public Input<string> Period { get; set; } = null!;

        public UsagePlanQuotaSettingsArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway.Inputs
{

    public sealed class AccountThrottleSettingsGetArgs : Pulumi.ResourceArgs
    {
        [Input("burstLimit")]
        public Input<int>? BurstLimit { get; set; }

        [Input("rateLimit")]
        public Input<double>? RateLimit { get; set; }

        public AccountThrottleSettingsGetArgs()
        {
        }
    }
}

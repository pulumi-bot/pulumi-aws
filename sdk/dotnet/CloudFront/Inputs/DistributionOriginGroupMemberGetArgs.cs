// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class DistributionOriginGroupMemberGetArgs : Pulumi.ResourceArgs
    {
        [Input("originId", required: true)]
        public Input<string> OriginId { get; set; } = null!;

        public DistributionOriginGroupMemberGetArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway.Inputs
{

    public sealed class StageAccessLogSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public StageAccessLogSettingsArgs()
        {
        }
    }
}

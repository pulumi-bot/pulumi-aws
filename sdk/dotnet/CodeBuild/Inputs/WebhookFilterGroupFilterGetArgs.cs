// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class WebhookFilterGroupFilterGetArgs : Pulumi.ResourceArgs
    {
        [Input("excludeMatchedPattern")]
        public Input<bool>? ExcludeMatchedPattern { get; set; }

        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public WebhookFilterGroupFilterGetArgs()
        {
        }
    }
}

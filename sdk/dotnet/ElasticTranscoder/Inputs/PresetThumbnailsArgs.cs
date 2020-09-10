// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticTranscoder.Inputs
{

    public sealed class PresetThumbnailsArgs : Pulumi.ResourceArgs
    {
        [Input("aspectRatio")]
        public Input<string>? AspectRatio { get; set; }

        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("maxHeight")]
        public Input<string>? MaxHeight { get; set; }

        [Input("maxWidth")]
        public Input<string>? MaxWidth { get; set; }

        [Input("paddingPolicy")]
        public Input<string>? PaddingPolicy { get; set; }

        [Input("resolution")]
        public Input<string>? Resolution { get; set; }

        [Input("sizingPolicy")]
        public Input<string>? SizingPolicy { get; set; }

        public PresetThumbnailsArgs()
        {
        }
    }
}

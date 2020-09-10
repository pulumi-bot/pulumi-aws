// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticTranscoder.Outputs
{

    [OutputType]
    public sealed class PresetVideoWatermark
    {
        public readonly string? HorizontalAlign;
        public readonly string? HorizontalOffset;
        public readonly string? Id;
        public readonly string? MaxHeight;
        public readonly string? MaxWidth;
        public readonly string? Opacity;
        public readonly string? SizingPolicy;
        public readonly string? Target;
        public readonly string? VerticalAlign;
        public readonly string? VerticalOffset;

        [OutputConstructor]
        private PresetVideoWatermark(
            string? horizontalAlign,

            string? horizontalOffset,

            string? id,

            string? maxHeight,

            string? maxWidth,

            string? opacity,

            string? sizingPolicy,

            string? target,

            string? verticalAlign,

            string? verticalOffset)
        {
            HorizontalAlign = horizontalAlign;
            HorizontalOffset = horizontalOffset;
            Id = id;
            MaxHeight = maxHeight;
            MaxWidth = maxWidth;
            Opacity = opacity;
            SizingPolicy = sizingPolicy;
            Target = target;
            VerticalAlign = verticalAlign;
            VerticalOffset = verticalOffset;
        }
    }
}

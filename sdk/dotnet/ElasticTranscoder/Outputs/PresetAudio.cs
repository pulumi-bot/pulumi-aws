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
    public sealed class PresetAudio
    {
        public readonly string? AudioPackingMode;
        public readonly string? BitRate;
        public readonly string? Channels;
        public readonly string? Codec;
        public readonly string? SampleRate;

        [OutputConstructor]
        private PresetAudio(
            string? audioPackingMode,

            string? bitRate,

            string? channels,

            string? codec,

            string? sampleRate)
        {
            AudioPackingMode = audioPackingMode;
            BitRate = bitRate;
            Channels = channels;
            Codec = codec;
            SampleRate = sampleRate;
        }
    }
}

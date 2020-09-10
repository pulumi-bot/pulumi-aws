// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("processors")]
        private InputList<Inputs.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorGetArgs>? _processors;
        public InputList<Inputs.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorGetArgs> Processors
        {
            get => _processors ?? (_processors = new InputList<Inputs.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorGetArgs>());
            set => _processors = value;
        }

        public FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationGetArgs()
        {
        }
    }
}

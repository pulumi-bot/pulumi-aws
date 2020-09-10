// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Outputs
{

    [OutputType]
    public sealed class FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor
    {
        public readonly ImmutableArray<Outputs.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorParameter> Parameters;
        public readonly string Type;

        [OutputConstructor]
        private FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor(
            ImmutableArray<Outputs.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorParameter> parameters,

            string type)
        {
            Parameters = parameters;
            Type = type;
        }
    }
}

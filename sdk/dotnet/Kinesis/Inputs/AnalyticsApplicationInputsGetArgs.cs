// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class AnalyticsApplicationInputsGetArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("kinesisFirehose")]
        public Input<Inputs.AnalyticsApplicationInputsKinesisFirehoseGetArgs>? KinesisFirehose { get; set; }

        [Input("kinesisStream")]
        public Input<Inputs.AnalyticsApplicationInputsKinesisStreamGetArgs>? KinesisStream { get; set; }

        [Input("namePrefix", required: true)]
        public Input<string> NamePrefix { get; set; } = null!;

        [Input("parallelism")]
        public Input<Inputs.AnalyticsApplicationInputsParallelismGetArgs>? Parallelism { get; set; }

        [Input("processingConfiguration")]
        public Input<Inputs.AnalyticsApplicationInputsProcessingConfigurationGetArgs>? ProcessingConfiguration { get; set; }

        [Input("schema", required: true)]
        public Input<Inputs.AnalyticsApplicationInputsSchemaGetArgs> Schema { get; set; } = null!;

        [Input("startingPositionConfigurations")]
        private InputList<Inputs.AnalyticsApplicationInputsStartingPositionConfigurationGetArgs>? _startingPositionConfigurations;
        public InputList<Inputs.AnalyticsApplicationInputsStartingPositionConfigurationGetArgs> StartingPositionConfigurations
        {
            get => _startingPositionConfigurations ?? (_startingPositionConfigurations = new InputList<Inputs.AnalyticsApplicationInputsStartingPositionConfigurationGetArgs>());
            set => _startingPositionConfigurations = value;
        }

        [Input("streamNames")]
        private InputList<string>? _streamNames;
        public InputList<string> StreamNames
        {
            get => _streamNames ?? (_streamNames = new InputList<string>());
            set => _streamNames = value;
        }

        public AnalyticsApplicationInputsGetArgs()
        {
        }
    }
}

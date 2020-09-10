// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArgs : Pulumi.ResourceArgs
    {
        [Input("mapping")]
        public Input<string>? Mapping { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("sqlType", required: true)]
        public Input<string> SqlType { get; set; } = null!;

        public AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArgs()
        {
        }
    }
}

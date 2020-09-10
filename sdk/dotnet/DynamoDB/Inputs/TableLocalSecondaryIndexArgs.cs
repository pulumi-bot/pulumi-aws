// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Inputs
{

    public sealed class TableLocalSecondaryIndexArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("nonKeyAttributes")]
        private InputList<string>? _nonKeyAttributes;
        public InputList<string> NonKeyAttributes
        {
            get => _nonKeyAttributes ?? (_nonKeyAttributes = new InputList<string>());
            set => _nonKeyAttributes = value;
        }

        [Input("projectionType", required: true)]
        public Input<string> ProjectionType { get; set; } = null!;

        [Input("rangeKey", required: true)]
        public Input<string> RangeKey { get; set; } = null!;

        public TableLocalSecondaryIndexArgs()
        {
        }
    }
}

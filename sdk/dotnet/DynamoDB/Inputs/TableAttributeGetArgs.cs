// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Inputs
{

    public sealed class TableAttributeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the table, this needs to be unique
        /// within a region.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TableAttributeGetArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class ClassifierXmlClassifierGetArgs : Pulumi.ResourceArgs
    {
        [Input("classification", required: true)]
        public Input<string> Classification { get; set; } = null!;

        [Input("rowTag", required: true)]
        public Input<string> RowTag { get; set; } = null!;

        public ClassifierXmlClassifierGetArgs()
        {
        }
    }
}

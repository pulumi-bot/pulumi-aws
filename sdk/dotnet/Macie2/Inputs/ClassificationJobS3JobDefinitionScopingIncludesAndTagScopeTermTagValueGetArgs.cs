// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2.Inputs
{

    public sealed class ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValueGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object property to use in the condition.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValueGetArgs()
        {
        }
    }
}

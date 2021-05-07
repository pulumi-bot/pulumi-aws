// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2.Inputs
{

    public sealed class ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operator to use in the condition.
        /// </summary>
        [Input("comparator")]
        public Input<string>? Comparator { get; set; }

        /// <summary>
        /// The tag key to use in the condition.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("tagValues")]
        private InputList<Inputs.ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValueArgs>? _tagValues;

        /// <summary>
        /// The tag keys or tag key and value pairs to use in the condition.
        /// </summary>
        public InputList<Inputs.ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValueArgs> TagValues
        {
            get => _tagValues ?? (_tagValues = new InputList<Inputs.ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValueArgs>());
            set => _tagValues = value;
        }

        /// <summary>
        /// The type of object to apply the condition to.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermArgs()
        {
        }
    }
}

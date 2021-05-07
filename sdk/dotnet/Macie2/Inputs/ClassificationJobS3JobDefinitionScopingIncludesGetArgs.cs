// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2.Inputs
{

    public sealed class ClassificationJobS3JobDefinitionScopingIncludesGetArgs : Pulumi.ResourceArgs
    {
        [Input("ands")]
        private InputList<Inputs.ClassificationJobS3JobDefinitionScopingIncludesAndGetArgs>? _ands;

        /// <summary>
        /// An array of conditions, one for each condition that determines which objects to include or exclude from the job. (documented below)
        /// </summary>
        public InputList<Inputs.ClassificationJobS3JobDefinitionScopingIncludesAndGetArgs> Ands
        {
            get => _ands ?? (_ands = new InputList<Inputs.ClassificationJobS3JobDefinitionScopingIncludesAndGetArgs>());
            set => _ands = value;
        }

        public ClassificationJobS3JobDefinitionScopingIncludesGetArgs()
        {
        }
    }
}

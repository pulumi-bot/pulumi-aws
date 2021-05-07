// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2.Outputs
{

    [OutputType]
    public sealed class ClassificationJobS3JobDefinitionScopingIncludes
    {
        /// <summary>
        /// An array of conditions, one for each condition that determines which objects to include or exclude from the job. (documented below)
        /// </summary>
        public readonly ImmutableArray<Outputs.ClassificationJobS3JobDefinitionScopingIncludesAnd> Ands;

        [OutputConstructor]
        private ClassificationJobS3JobDefinitionScopingIncludes(ImmutableArray<Outputs.ClassificationJobS3JobDefinitionScopingIncludesAnd> ands)
        {
            Ands = ands;
        }
    }
}

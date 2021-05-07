// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2.Inputs
{

    public sealed class ClassificationJobS3JobDefinitionGetArgs : Pulumi.ResourceArgs
    {
        [Input("bucketDefinitions")]
        private InputList<Inputs.ClassificationJobS3JobDefinitionBucketDefinitionGetArgs>? _bucketDefinitions;

        /// <summary>
        /// An array of objects, one for each AWS account that owns buckets to analyze. Each object specifies the account ID for an account and one or more buckets to analyze for the account. (documented below)
        /// </summary>
        public InputList<Inputs.ClassificationJobS3JobDefinitionBucketDefinitionGetArgs> BucketDefinitions
        {
            get => _bucketDefinitions ?? (_bucketDefinitions = new InputList<Inputs.ClassificationJobS3JobDefinitionBucketDefinitionGetArgs>());
            set => _bucketDefinitions = value;
        }

        /// <summary>
        /// The property- and tag-based conditions that determine which objects to include or exclude from the analysis. (documented below)
        /// </summary>
        [Input("scoping")]
        public Input<Inputs.ClassificationJobS3JobDefinitionScopingGetArgs>? Scoping { get; set; }

        public ClassificationJobS3JobDefinitionGetArgs()
        {
        }
    }
}

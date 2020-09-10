// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg.Outputs
{

    [OutputType]
    public sealed class RuleSource
    {
        public readonly string Owner;
        public readonly ImmutableArray<Outputs.RuleSourceSourceDetail> SourceDetails;
        public readonly string SourceIdentifier;

        [OutputConstructor]
        private RuleSource(
            string owner,

            ImmutableArray<Outputs.RuleSourceSourceDetail> sourceDetails,

            string sourceIdentifier)
        {
            Owner = owner;
            SourceDetails = sourceDetails;
            SourceIdentifier = sourceIdentifier;
        }
    }
}

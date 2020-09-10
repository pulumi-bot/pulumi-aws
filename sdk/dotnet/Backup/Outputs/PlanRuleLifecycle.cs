// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup.Outputs
{

    [OutputType]
    public sealed class PlanRuleLifecycle
    {
        public readonly int? ColdStorageAfter;
        public readonly int? DeleteAfter;

        [OutputConstructor]
        private PlanRuleLifecycle(
            int? coldStorageAfter,

            int? deleteAfter)
        {
            ColdStorageAfter = coldStorageAfter;
            DeleteAfter = deleteAfter;
        }
    }
}

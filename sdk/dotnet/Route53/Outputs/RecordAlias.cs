// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53.Outputs
{

    [OutputType]
    public sealed class RecordAlias
    {
        public readonly bool EvaluateTargetHealth;
        public readonly string Name;
        public readonly string ZoneId;

        [OutputConstructor]
        private RecordAlias(
            bool evaluateTargetHealth,

            string name,

            string zoneId)
        {
            EvaluateTargetHealth = evaluateTargetHealth;
            Name = name;
            ZoneId = zoneId;
        }
    }
}

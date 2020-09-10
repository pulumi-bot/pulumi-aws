// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaConvert.Outputs
{

    [OutputType]
    public sealed class QueueReservationPlanSettings
    {
        public readonly string Commitment;
        public readonly string RenewalType;
        public readonly int ReservedSlots;

        [OutputConstructor]
        private QueueReservationPlanSettings(
            string commitment,

            string renewalType,

            int reservedSlots)
        {
            Commitment = commitment;
            RenewalType = renewalType;
            ReservedSlots = reservedSlots;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Outputs
{

    [OutputType]
    public sealed class MetricAlarmMetricQueryMetric
    {
        public readonly ImmutableDictionary<string, string>? Dimensions;
        public readonly string MetricName;
        public readonly string? Namespace;
        public readonly int Period;
        public readonly string Stat;
        public readonly string? Unit;

        [OutputConstructor]
        private MetricAlarmMetricQueryMetric(
            ImmutableDictionary<string, string>? dimensions,

            string metricName,

            string? @namespace,

            int period,

            string stat,

            string? unit)
        {
            Dimensions = dimensions;
            MetricName = metricName;
            Namespace = @namespace;
            Period = period;
            Stat = stat;
            Unit = unit;
        }
    }
}

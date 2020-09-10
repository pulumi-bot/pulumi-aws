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
    public sealed class LogMetricFilterMetricTransformation
    {
        public readonly string? DefaultValue;
        public readonly string Name;
        public readonly string Namespace;
        public readonly string Value;

        [OutputConstructor]
        private LogMetricFilterMetricTransformation(
            string? defaultValue,

            string name,

            string @namespace,

            string value)
        {
            DefaultValue = defaultValue;
            Name = name;
            Namespace = @namespace;
            Value = value;
        }
    }
}

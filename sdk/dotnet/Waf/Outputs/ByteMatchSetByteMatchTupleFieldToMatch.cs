// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf.Outputs
{

    [OutputType]
    public sealed class ByteMatchSetByteMatchTupleFieldToMatch
    {
        public readonly string? Data;
        public readonly string Type;

        [OutputConstructor]
        private ByteMatchSetByteMatchTupleFieldToMatch(
            string? data,

            string type)
        {
            Data = data;
            Type = type;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda.Outputs
{

    [OutputType]
    public sealed class FunctionEventInvokeConfigDestinationConfigOnFailure
    {
        public readonly string Destination;

        [OutputConstructor]
        private FunctionEventInvokeConfigDestinationConfigOnFailure(string destination)
        {
            Destination = destination;
        }
    }
}

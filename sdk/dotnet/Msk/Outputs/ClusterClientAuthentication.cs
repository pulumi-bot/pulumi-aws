// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterClientAuthentication
    {
        public readonly Outputs.ClusterClientAuthenticationTls? Tls;

        [OutputConstructor]
        private ClusterClientAuthentication(Outputs.ClusterClientAuthenticationTls? tls)
        {
            Tls = tls;
        }
    }
}

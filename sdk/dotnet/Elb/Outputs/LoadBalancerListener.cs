// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Elb.Outputs
{

    [OutputType]
    public sealed class LoadBalancerListener
    {
        public readonly int InstancePort;
        public readonly string InstanceProtocol;
        public readonly int LbPort;
        public readonly string LbProtocol;
        public readonly string? SslCertificateId;

        [OutputConstructor]
        private LoadBalancerListener(
            int instancePort,

            string instanceProtocol,

            int lbPort,

            string lbProtocol,

            string? sslCertificateId)
        {
            InstancePort = instancePort;
            InstanceProtocol = instanceProtocol;
            LbPort = lbPort;
            LbProtocol = lbProtocol;
            SslCertificateId = sslCertificateId;
        }
    }
}

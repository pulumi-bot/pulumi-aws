// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Acmpca.Inputs
{

    public sealed class CertificateAuthorityCertificateAuthorityConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("keyAlgorithm", required: true)]
        public Input<string> KeyAlgorithm { get; set; } = null!;

        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        [Input("subject", required: true)]
        public Input<Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs> Subject { get; set; } = null!;

        public CertificateAuthorityCertificateAuthorityConfigurationArgs()
        {
        }
    }
}

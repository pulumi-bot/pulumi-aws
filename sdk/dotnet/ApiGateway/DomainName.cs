// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// ## Import
    /// 
    /// API Gateway domain names can be imported using their `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigateway/domainName:DomainName example dev.example.com
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigateway/domainName:DomainName")]
    public partial class DomainName : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Output("certificateArn")]
        public Output<string?> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// The certificate issued for the domain name
        /// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`.
        /// </summary>
        [Output("certificateBody")]
        public Output<string?> CertificateBody { get; private set; } = null!;

        /// <summary>
        /// The certificate for the CA that issued the
        /// certificate, along with any intermediate CA certificates required to
        /// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        /// `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// The unique name to use when registering this
        /// certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`. Required if `certificate_arn` is not set.
        /// </summary>
        [Output("certificateName")]
        public Output<string?> CertificateName { get; private set; } = null!;

        /// <summary>
        /// The private key associated with the
        /// domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Output("certificatePrivateKey")]
        public Output<string?> CertificatePrivateKey { get; private set; } = null!;

        /// <summary>
        /// The upload date associated with the domain certificate.
        /// </summary>
        [Output("certificateUploadDate")]
        public Output<string> CertificateUploadDate { get; private set; } = null!;

        /// <summary>
        /// The hostname created by Cloudfront to represent
        /// the distribution that implements this domain name mapping.
        /// </summary>
        [Output("cloudfrontDomainName")]
        public Output<string> CloudfrontDomainName { get; private set; } = null!;

        /// <summary>
        /// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
        /// that can be used to create a Route53 alias record for the distribution.
        /// </summary>
        [Output("cloudfrontZoneId")]
        public Output<string> CloudfrontZoneId { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified domain name to register
        /// </summary>
        [Output("domainName")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Configuration block defining API endpoint information including type. Defined below.
        /// </summary>
        [Output("endpointConfiguration")]
        public Output<Outputs.DomainNameEndpointConfiguration> EndpointConfiguration { get; private set; } = null!;

        /// <summary>
        /// The mutual TLS authentication configuration for the domain name. Defined below.
        /// </summary>
        [Output("mutualTlsAuthentication")]
        public Output<Outputs.DomainNameMutualTlsAuthentication?> MutualTlsAuthentication { get; private set; } = null!;

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        /// </summary>
        [Output("regionalCertificateArn")]
        public Output<string?> RegionalCertificateArn { get; private set; } = null!;

        /// <summary>
        /// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        /// `certificate_private_key`.
        /// </summary>
        [Output("regionalCertificateName")]
        public Output<string?> RegionalCertificateName { get; private set; } = null!;

        /// <summary>
        /// The hostname for the custom domain's regional endpoint.
        /// </summary>
        [Output("regionalDomainName")]
        public Output<string> RegionalDomainName { get; private set; } = null!;

        /// <summary>
        /// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
        /// </summary>
        [Output("regionalZoneId")]
        public Output<string> RegionalZoneId { get; private set; } = null!;

        /// <summary>
        /// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        /// </summary>
        [Output("securityPolicy")]
        public Output<string> SecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DomainName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainName(string name, DomainNameArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/domainName:DomainName", name, args ?? new DomainNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainName(string name, Input<string> id, DomainNameState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/domainName:DomainName", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainName Get(string name, Input<string> id, DomainNameState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainName(name, id, state, options);
        }
    }

    public sealed class DomainNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The certificate issued for the domain name
        /// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`.
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The certificate for the CA that issued the
        /// certificate, along with any intermediate CA certificates required to
        /// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        /// `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The unique name to use when registering this
        /// certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`. Required if `certificate_arn` is not set.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// The private key associated with the
        /// domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificatePrivateKey")]
        public Input<string>? CertificatePrivateKey { get; set; }

        /// <summary>
        /// The fully-qualified domain name to register
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Configuration block defining API endpoint information including type. Defined below.
        /// </summary>
        [Input("endpointConfiguration")]
        public Input<Inputs.DomainNameEndpointConfigurationArgs>? EndpointConfiguration { get; set; }

        /// <summary>
        /// The mutual TLS authentication configuration for the domain name. Defined below.
        /// </summary>
        [Input("mutualTlsAuthentication")]
        public Input<Inputs.DomainNameMutualTlsAuthenticationArgs>? MutualTlsAuthentication { get; set; }

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateArn")]
        public Input<string>? RegionalCertificateArn { get; set; }

        /// <summary>
        /// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        /// `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateName")]
        public Input<string>? RegionalCertificateName { get; set; }

        /// <summary>
        /// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DomainNameArgs()
        {
        }
    }

    public sealed class DomainNameState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The certificate issued for the domain name
        /// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`.
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The certificate for the CA that issued the
        /// certificate, along with any intermediate CA certificates required to
        /// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        /// `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The unique name to use when registering this
        /// certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`. Required if `certificate_arn` is not set.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// The private key associated with the
        /// domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificatePrivateKey")]
        public Input<string>? CertificatePrivateKey { get; set; }

        /// <summary>
        /// The upload date associated with the domain certificate.
        /// </summary>
        [Input("certificateUploadDate")]
        public Input<string>? CertificateUploadDate { get; set; }

        /// <summary>
        /// The hostname created by Cloudfront to represent
        /// the distribution that implements this domain name mapping.
        /// </summary>
        [Input("cloudfrontDomainName")]
        public Input<string>? CloudfrontDomainName { get; set; }

        /// <summary>
        /// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
        /// that can be used to create a Route53 alias record for the distribution.
        /// </summary>
        [Input("cloudfrontZoneId")]
        public Input<string>? CloudfrontZoneId { get; set; }

        /// <summary>
        /// The fully-qualified domain name to register
        /// </summary>
        [Input("domainName")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Configuration block defining API endpoint information including type. Defined below.
        /// </summary>
        [Input("endpointConfiguration")]
        public Input<Inputs.DomainNameEndpointConfigurationGetArgs>? EndpointConfiguration { get; set; }

        /// <summary>
        /// The mutual TLS authentication configuration for the domain name. Defined below.
        /// </summary>
        [Input("mutualTlsAuthentication")]
        public Input<Inputs.DomainNameMutualTlsAuthenticationGetArgs>? MutualTlsAuthentication { get; set; }

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateArn")]
        public Input<string>? RegionalCertificateArn { get; set; }

        /// <summary>
        /// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        /// `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateName")]
        public Input<string>? RegionalCertificateName { get; set; }

        /// <summary>
        /// The hostname for the custom domain's regional endpoint.
        /// </summary>
        [Input("regionalDomainName")]
        public Input<string>? RegionalDomainName { get; set; }

        /// <summary>
        /// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
        /// </summary>
        [Input("regionalZoneId")]
        public Input<string>? RegionalZoneId { get; set; }

        /// <summary>
        /// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DomainNameState()
        {
        }
    }
}

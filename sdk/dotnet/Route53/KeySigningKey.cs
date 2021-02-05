// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    /// <summary>
    /// Manages an Route 53 Key Signing Key. For more information about managing Domain Name System Security Extensions (DNSSEC)in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleKey = new Aws.Kms.Key("exampleKey", new Aws.Kms.KeyArgs
    ///         {
    ///             CustomerMasterKeySpec = "ECC_NIST_P256",
    ///             DeletionWindowInDays = 7,
    ///             KeyUsage = "SIGN_VERIFY",
    ///             Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", new[]
    ///                                 {
    ///                                     "kms:DescribeKey",
    ///                                     "kms:GetPublicKey",
    ///                                     "kms:Sign",
    ///                                 }
    ///                              },
    ///                             { "Effect", "Allow" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "Service", "api-service.dnssec.route53.aws.internal" },
    ///                             } },
    ///                             { "Sid", "Route 53 DNSSEC Permissions" },
    ///                         },
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "kms:*" },
    ///                             { "Effect", "Allow" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "AWS", "*" },
    ///                             } },
    ///                             { "Resource", "*" },
    ///                             { "Sid", "IAM User Permissions" },
    ///                         },
    ///                     }
    ///                  },
    ///                 { "Version", "2012-10-17" },
    ///             }),
    ///         });
    ///         var exampleZone = new Aws.Route53.Zone("exampleZone", new Aws.Route53.ZoneArgs
    ///         {
    ///         });
    ///         var exampleKeySigningKey = new Aws.Route53.KeySigningKey("exampleKeySigningKey", new Aws.Route53.KeySigningKeyArgs
    ///         {
    ///             HostedZoneId = aws_route53_zone.Test.Id,
    ///             KeyManagementServiceArn = aws_kms_key.Test.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_route53_key_signing_key` resources can be imported by using the Route 53 Hosted Zone identifier and KMS Key identifier, separated by a comma (`,`), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:route53/keySigningKey:KeySigningKey example Z1D633PJN98FT9,example
    /// ```
    /// </summary>
    [AwsResourceType("aws:route53/keySigningKey:KeySigningKey")]
    public partial class KeySigningKey : Pulumi.CustomResource
    {
        /// <summary>
        /// A string used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
        /// </summary>
        [Output("digestAlgorithmMnemonic")]
        public Output<string> DigestAlgorithmMnemonic { get; private set; } = null!;

        /// <summary>
        /// An integer used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
        /// </summary>
        [Output("digestAlgorithmType")]
        public Output<int> DigestAlgorithmType { get; private set; } = null!;

        /// <summary>
        /// A cryptographic digest of a DNSKEY resource record (RR). DNSKEY records are used to publish the public key that resolvers can use to verify DNSSEC signatures that are used to secure certain kinds of information provided by the DNS system.
        /// </summary>
        [Output("digestValue")]
        public Output<string> DigestValue { get; private set; } = null!;

        /// <summary>
        /// A string that represents a DNSKEY record.
        /// </summary>
        [Output("dnskeyRecord")]
        public Output<string> DnskeyRecord { get; private set; } = null!;

        /// <summary>
        /// A string that represents a delegation signer (DS) record.
        /// </summary>
        [Output("dsRecord")]
        public Output<string> DsRecord { get; private set; } = null!;

        /// <summary>
        /// An integer that specifies how the key is used. For key-signing key (KSK), this value is always 257.
        /// </summary>
        [Output("flag")]
        public Output<int> Flag { get; private set; } = null!;

        /// <summary>
        /// Identifier of the Route 53 Hosted Zone.
        /// </summary>
        [Output("hostedZoneId")]
        public Output<string> HostedZoneId { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
        /// </summary>
        [Output("keyManagementServiceArn")]
        public Output<string> KeyManagementServiceArn { get; private set; } = null!;

        /// <summary>
        /// An integer used to identify the DNSSEC record for the domain name. The process used to calculate the value is described in [RFC-4034 Appendix B](https://tools.ietf.org/rfc/rfc4034.txt).
        /// </summary>
        [Output("keyTag")]
        public Output<int> KeyTag { get; private set; } = null!;

        /// <summary>
        /// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The public key, represented as a Base64 encoding, as required by [RFC-4034 Page 5](https://tools.ietf.org/rfc/rfc4034.txt).
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// A string used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
        /// </summary>
        [Output("signingAlgorithmMnemonic")]
        public Output<string> SigningAlgorithmMnemonic { get; private set; } = null!;

        /// <summary>
        /// An integer used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
        /// </summary>
        [Output("signingAlgorithmType")]
        public Output<int> SigningAlgorithmType { get; private set; } = null!;

        /// <summary>
        /// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a KeySigningKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeySigningKey(string name, KeySigningKeyArgs args, CustomResourceOptions? options = null)
            : base("aws:route53/keySigningKey:KeySigningKey", name, args ?? new KeySigningKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeySigningKey(string name, Input<string> id, KeySigningKeyState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/keySigningKey:KeySigningKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeySigningKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeySigningKey Get(string name, Input<string> id, KeySigningKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new KeySigningKey(name, id, state, options);
        }
    }

    public sealed class KeySigningKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the Route 53 Hosted Zone.
        /// </summary>
        [Input("hostedZoneId", required: true)]
        public Input<string> HostedZoneId { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
        /// </summary>
        [Input("keyManagementServiceArn", required: true)]
        public Input<string> KeyManagementServiceArn { get; set; } = null!;

        /// <summary>
        /// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public KeySigningKeyArgs()
        {
        }
    }

    public sealed class KeySigningKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A string used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
        /// </summary>
        [Input("digestAlgorithmMnemonic")]
        public Input<string>? DigestAlgorithmMnemonic { get; set; }

        /// <summary>
        /// An integer used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
        /// </summary>
        [Input("digestAlgorithmType")]
        public Input<int>? DigestAlgorithmType { get; set; }

        /// <summary>
        /// A cryptographic digest of a DNSKEY resource record (RR). DNSKEY records are used to publish the public key that resolvers can use to verify DNSSEC signatures that are used to secure certain kinds of information provided by the DNS system.
        /// </summary>
        [Input("digestValue")]
        public Input<string>? DigestValue { get; set; }

        /// <summary>
        /// A string that represents a DNSKEY record.
        /// </summary>
        [Input("dnskeyRecord")]
        public Input<string>? DnskeyRecord { get; set; }

        /// <summary>
        /// A string that represents a delegation signer (DS) record.
        /// </summary>
        [Input("dsRecord")]
        public Input<string>? DsRecord { get; set; }

        /// <summary>
        /// An integer that specifies how the key is used. For key-signing key (KSK), this value is always 257.
        /// </summary>
        [Input("flag")]
        public Input<int>? Flag { get; set; }

        /// <summary>
        /// Identifier of the Route 53 Hosted Zone.
        /// </summary>
        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
        /// </summary>
        [Input("keyManagementServiceArn")]
        public Input<string>? KeyManagementServiceArn { get; set; }

        /// <summary>
        /// An integer used to identify the DNSSEC record for the domain name. The process used to calculate the value is described in [RFC-4034 Appendix B](https://tools.ietf.org/rfc/rfc4034.txt).
        /// </summary>
        [Input("keyTag")]
        public Input<int>? KeyTag { get; set; }

        /// <summary>
        /// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public key, represented as a Base64 encoding, as required by [RFC-4034 Page 5](https://tools.ietf.org/rfc/rfc4034.txt).
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// A string used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
        /// </summary>
        [Input("signingAlgorithmMnemonic")]
        public Input<string>? SigningAlgorithmMnemonic { get; set; }

        /// <summary>
        /// An integer used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
        /// </summary>
        [Input("signingAlgorithmType")]
        public Input<int>? SigningAlgorithmType { get; set; }

        /// <summary>
        /// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public KeySigningKeyState()
        {
        }
    }
}

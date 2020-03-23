// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES domain identity resource
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_domain_identity.html.markdown.
    /// </summary>
    public partial class DomainIdentity : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the domain identity.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The domain name to assign to SES
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// A code which when added to the domain as a TXT record
        /// will signal to SES that the owner of the domain has authorised SES to act on
        /// their behalf. The domain identity will be in state "verification pending"
        /// until this is done. See below for an example of how this might be achieved
        /// when the domain is hosted in Route 53 and managed by this provider.  Find out
        /// more about verifying domains in Amazon SES in the [AWS SES
        /// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
        /// </summary>
        [Output("verificationToken")]
        public Output<string> VerificationToken { get; private set; } = null!;


        /// <summary>
        /// Create a DomainIdentity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainIdentity(string name, DomainIdentityArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/domainIdentity:DomainIdentity", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DomainIdentity(string name, Input<string> id, DomainIdentityState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/domainIdentity:DomainIdentity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainIdentity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainIdentity Get(string name, Input<string> id, DomainIdentityState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainIdentity(name, id, state, options);
        }
    }

    public sealed class DomainIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name to assign to SES
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        public DomainIdentityArgs()
        {
        }
    }

    public sealed class DomainIdentityState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the domain identity.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The domain name to assign to SES
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// A code which when added to the domain as a TXT record
        /// will signal to SES that the owner of the domain has authorised SES to act on
        /// their behalf. The domain identity will be in state "verification pending"
        /// until this is done. See below for an example of how this might be achieved
        /// when the domain is hosted in Route 53 and managed by this provider.  Find out
        /// more about verifying domains in Amazon SES in the [AWS SES
        /// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
        /// </summary>
        [Input("verificationToken")]
        public Input<string>? VerificationToken { get; set; }

        public DomainIdentityState()
        {
        }
    }
}

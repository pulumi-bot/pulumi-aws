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
    /// Provides an SES email identity resource
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ses.EmailIdentity("example", new Aws.Ses.EmailIdentityArgs
    ///         {
    ///             Email = "email@example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class EmailIdentity : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the email identity.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The email address to assign to SES
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;


        /// <summary>
        /// Create a EmailIdentity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EmailIdentity(string name, EmailIdentityArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/emailIdentity:EmailIdentity", name, args ?? new EmailIdentityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EmailIdentity(string name, Input<string> id, EmailIdentityState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/emailIdentity:EmailIdentity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EmailIdentity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EmailIdentity Get(string name, Input<string> id, EmailIdentityState? state = null, CustomResourceOptions? options = null)
        {
            return new EmailIdentity(name, id, state, options);
        }
    }

    public sealed class EmailIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email address to assign to SES
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public EmailIdentityArgs()
        {
        }
    }

    public sealed class EmailIdentityState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the email identity.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The email address to assign to SES
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        public EmailIdentityState()
        {
        }
    }
}

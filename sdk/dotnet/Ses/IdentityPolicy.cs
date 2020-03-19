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
    /// Manages a SES Identity Policy. More information about SES Sending Authorization Policies can be found in the [SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_identity_policy.html.markdown.
    /// </summary>
    public partial class IdentityPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Name or Amazon Resource Name (ARN) of the SES Identity.
        /// </summary>
        [Output("identity")]
        public Output<string> Identity { get; private set; } = null!;

        /// <summary>
        /// Name of the policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityPolicy(string name, IdentityPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/identityPolicy:IdentityPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IdentityPolicy(string name, Input<string> id, IdentityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/identityPolicy:IdentityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityPolicy Get(string name, Input<string> id, IdentityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityPolicy(name, id, state, options);
        }
    }

    public sealed class IdentityPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name or Amazon Resource Name (ARN) of the SES Identity.
        /// </summary>
        [Input("identity", required: true)]
        public Input<string> Identity { get; set; } = null!;

        /// <summary>
        /// Name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public IdentityPolicyArgs()
        {
        }
    }

    public sealed class IdentityPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name or Amazon Resource Name (ARN) of the SES Identity.
        /// </summary>
        [Input("identity")]
        public Input<string>? Identity { get; set; }

        /// <summary>
        /// Name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public IdentityPolicyState()
        {
        }
    }
}

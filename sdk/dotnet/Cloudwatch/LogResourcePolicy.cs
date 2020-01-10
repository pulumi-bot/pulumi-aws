// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides a resource to manage a CloudWatch log resource policy.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_resource_policy.html.markdown.
    /// </summary>
    public partial class LogResourcePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        /// </summary>
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Name of the resource policy.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;


        /// <summary>
        /// Create a LogResourcePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogResourcePolicy(string name, LogResourcePolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/logResourcePolicy:LogResourcePolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LogResourcePolicy(string name, Input<string> id, LogResourcePolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/logResourcePolicy:LogResourcePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogResourcePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogResourcePolicy Get(string name, Input<string> id, LogResourcePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LogResourcePolicy(name, id, state, options);
        }
    }

    public sealed class LogResourcePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<string> PolicyDocument { get; set; } = null!;

        /// <summary>
        /// Name of the resource policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        public LogResourcePolicyArgs()
        {
        }
    }

    public sealed class LogResourcePolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// Name of the resource policy.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public LogResourcePolicyState()
        {
        }
    }
}

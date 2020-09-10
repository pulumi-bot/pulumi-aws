// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    public partial class LogResourcePolicy : Pulumi.CustomResource
    {
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

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
            : base("aws:cloudwatch/logResourcePolicy:LogResourcePolicy", name, args ?? new LogResourcePolicyArgs(), MakeResourceOptions(options, ""))
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
        [Input("policyDocument", required: true)]
        public Input<string> PolicyDocument { get; set; } = null!;

        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        public LogResourcePolicyArgs()
        {
        }
    }

    public sealed class LogResourcePolicyState : Pulumi.ResourceArgs
    {
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public LogResourcePolicyState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// Attaches a Managed IAM Policy to an IAM group
    /// 
    /// &gt; **NOTE:** The usage of this resource conflicts with the `aws.iam.PolicyAttachment` resource and will permanently show a difference if both are defined.
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
    ///         var @group = new Aws.Iam.Group("group", new Aws.Iam.GroupArgs
    ///         {
    ///         });
    ///         var policy = new Aws.Iam.Policy("policy", new Aws.Iam.PolicyArgs
    ///         {
    ///             Description = "A test policy",
    ///             Policy = "",
    ///         });
    ///         // insert policy here
    ///         var test_attach = new Aws.Iam.GroupPolicyAttachment("test-attach", new Aws.Iam.GroupPolicyAttachmentArgs
    ///         {
    ///             Group = @group.Name,
    ///             PolicyArn = policy.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class GroupPolicyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The group the policy should be applied to
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The ARN of the policy you want to apply
        /// </summary>
        [Output("policyArn")]
        public Output<string> PolicyArn { get; private set; } = null!;


        /// <summary>
        /// Create a GroupPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupPolicyAttachment(string name, GroupPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/groupPolicyAttachment:GroupPolicyAttachment", name, args ?? new GroupPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupPolicyAttachment(string name, Input<string> id, GroupPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/groupPolicyAttachment:GroupPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupPolicyAttachment Get(string name, Input<string> id, GroupPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class GroupPolicyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group the policy should be applied to
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The ARN of the policy you want to apply
        /// </summary>
        [Input("policyArn", required: true)]
        public Input<string> PolicyArn { get; set; } = null!;

        public GroupPolicyAttachmentArgs()
        {
        }
    }

    public sealed class GroupPolicyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group the policy should be applied to
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The ARN of the policy you want to apply
        /// </summary>
        [Input("policyArn")]
        public Input<string>? PolicyArn { get; set; }

        public GroupPolicyAttachmentState()
        {
        }
    }
}

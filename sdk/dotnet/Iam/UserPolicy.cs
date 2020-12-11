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
    /// Provides an IAM policy attached to a user.
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
    ///         var lbUser = new Aws.Iam.User("lbUser", new Aws.Iam.UserArgs
    ///         {
    ///             Path = "/system/",
    ///         });
    ///         var lbRo = new Aws.Iam.UserPolicy("lbRo", new Aws.Iam.UserPolicyArgs
    ///         {
    ///             User = lbUser.Name,
    ///             Policy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": [
    ///         ""ec2:Describe*""
    ///       ],
    ///       ""Effect"": ""Allow"",
    ///       ""Resource"": ""*""
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var lbAccessKey = new Aws.Iam.AccessKey("lbAccessKey", new Aws.Iam.AccessKeyArgs
    ///         {
    ///             User = lbUser.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM User Policies can be imported using the `user_name:user_policy_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:iam/userPolicy:UserPolicy mypolicy user_of_mypolicy_name:mypolicy_name
    /// ```
    /// </summary>
    [AwsResourceType("aws:iam/userPolicy:UserPolicy")]
    public partial class UserPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the policy. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// IAM user to which to attach this policy.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a UserPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPolicy(string name, UserPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/userPolicy:UserPolicy", name, args ?? new UserPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPolicy(string name, Input<string> id, UserPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/userPolicy:UserPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPolicy Get(string name, Input<string> id, UserPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new UserPolicy(name, id, state, options);
        }
    }

    public sealed class UserPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the policy. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// IAM user to which to attach this policy.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public UserPolicyArgs()
        {
        }
    }

    public sealed class UserPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the policy. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// IAM user to which to attach this policy.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public UserPolicyState()
        {
        }
    }
}

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
    /// Provides a resource for adding an [IAM User](https://www.terraform.io/docs/providers/aws/r/iam_user.html) to [IAM Groups](https://www.terraform.io/docs/providers/aws/r/iam_group.html). This
    /// resource can be used multiple times with the same user for non-overlapping
    /// groups.
    /// 
    /// To exclusively manage the users in a group, see the
    /// [`aws.iam.GroupMembership` resource][3].
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var user1 = new Aws.Iam.User("user1", new Aws.Iam.UserArgs
    ///         {
    ///         });
    ///         var group1 = new Aws.Iam.Group("group1", new Aws.Iam.GroupArgs
    ///         {
    ///         });
    ///         var group2 = new Aws.Iam.Group("group2", new Aws.Iam.GroupArgs
    ///         {
    ///         });
    ///         var example1 = new Aws.Iam.UserGroupMembership("example1", new Aws.Iam.UserGroupMembershipArgs
    ///         {
    ///             Groups = 
    ///             {
    ///                 group1.Name,
    ///                 group2.Name,
    ///             },
    ///             User = user1.Name,
    ///         });
    ///         var group3 = new Aws.Iam.Group("group3", new Aws.Iam.GroupArgs
    ///         {
    ///         });
    ///         var example2 = new Aws.Iam.UserGroupMembership("example2", new Aws.Iam.UserGroupMembershipArgs
    ///         {
    ///             Groups = 
    ///             {
    ///                 group3.Name,
    ///             },
    ///             User = user1.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class UserGroupMembership : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of [IAM Groups](https://www.terraform.io/docs/providers/aws/r/iam_group.html) to add the user to
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// The name of the [IAM User](https://www.terraform.io/docs/providers/aws/r/iam_user.html) to add to groups
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a UserGroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGroupMembership(string name, UserGroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/userGroupMembership:UserGroupMembership", name, args ?? new UserGroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGroupMembership(string name, Input<string> id, UserGroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/userGroupMembership:UserGroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserGroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGroupMembership Get(string name, Input<string> id, UserGroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGroupMembership(name, id, state, options);
        }
    }

    public sealed class UserGroupMembershipArgs : Pulumi.ResourceArgs
    {
        [Input("groups", required: true)]
        private InputList<string>? _groups;

        /// <summary>
        /// A list of [IAM Groups](https://www.terraform.io/docs/providers/aws/r/iam_group.html) to add the user to
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The name of the [IAM User](https://www.terraform.io/docs/providers/aws/r/iam_user.html) to add to groups
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public UserGroupMembershipArgs()
        {
        }
    }

    public sealed class UserGroupMembershipState : Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// A list of [IAM Groups](https://www.terraform.io/docs/providers/aws/r/iam_group.html) to add the user to
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The name of the [IAM User](https://www.terraform.io/docs/providers/aws/r/iam_user.html) to add to groups
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public UserGroupMembershipState()
        {
        }
    }
}

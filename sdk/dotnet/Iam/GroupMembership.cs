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
    /// &gt; **WARNING:** Multiple aws.iam.GroupMembership resources with the same group name will produce inconsistent behavior!
    /// 
    /// Provides a top level resource to manage IAM Group membership for IAM Users. For
    /// more information on managing IAM Groups or IAM Users, see [IAM Groups](https://www.terraform.io/docs/providers/aws/r/iam_group.html) or
    /// [IAM Users](https://www.terraform.io/docs/providers/aws/r/iam_user.html)
    /// 
    /// &gt; **Note:** `aws.iam.GroupMembership` will conflict with itself if used more than once with the same group. To non-exclusively manage the users in a group, see the
    /// [`aws.iam.UserGroupMembership` resource][3].
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
    ///         var userOne = new Aws.Iam.User("userOne", new Aws.Iam.UserArgs
    ///         {
    ///         });
    ///         var userTwo = new Aws.Iam.User("userTwo", new Aws.Iam.UserArgs
    ///         {
    ///         });
    ///         var team = new Aws.Iam.GroupMembership("team", new Aws.Iam.GroupMembershipArgs
    ///         {
    ///             Users = 
    ///             {
    ///                 userOne.Name,
    ///                 userTwo.Name,
    ///             },
    ///             Group = @group.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:iam/groupMembership:GroupMembership")]
    public partial class GroupMembership : Pulumi.CustomResource
    {
        /// <summary>
        /// The IAM Group name to attach the list of `users` to
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name to identify the Group Membership
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of IAM User names to associate with the Group
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMembership(string name, GroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/groupMembership:GroupMembership", name, args ?? new GroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMembership(string name, Input<string> id, GroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/groupMembership:GroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMembership Get(string name, Input<string> id, GroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMembership(name, id, state, options);
        }
    }

    public sealed class GroupMembershipArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM Group name to attach the list of `users` to
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name to identify the Group Membership
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("users", required: true)]
        private InputList<string>? _users;

        /// <summary>
        /// A list of IAM User names to associate with the Group
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public GroupMembershipArgs()
        {
        }
    }

    public sealed class GroupMembershipState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM Group name to attach the list of `users` to
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name to identify the Group Membership
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// A list of IAM User names to associate with the Group
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public GroupMembershipState()
        {
        }
    }
}

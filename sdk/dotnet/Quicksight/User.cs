// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight
{
    /// <summary>
    /// Resource for managing QuickSight User
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
    ///         var example = new Aws.Quicksight.User("example", new Aws.Quicksight.UserArgs
    ///         {
    ///             Email = "author@example.com",
    ///             IdentityType = "IAM",
    ///             UserName = "an-author",
    ///             UserRole = "AUTHOR",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Importing is currently not supported on this resource.
    /// </summary>
    [AwsResourceType("aws:quicksight/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the user
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// The email address of the user that you want to register.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
        /// </summary>
        [Output("iamArn")]
        public Output<string?> IamArn { get; private set; } = null!;

        /// <summary>
        /// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
        /// </summary>
        [Output("identityType")]
        public Output<string> IdentityType { get; private set; } = null!;

        /// <summary>
        /// The namespace. Currently, you should set this to `default`.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
        /// </summary>
        [Output("sessionName")]
        public Output<string?> SessionName { get; private set; } = null!;

        /// <summary>
        /// The Amazon QuickSight user name that you want to create for the user you are registering.
        /// </summary>
        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;

        /// <summary>
        /// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
        /// </summary>
        [Output("userRole")]
        public Output<string> UserRole { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("aws:quicksight/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("aws:quicksight/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// The email address of the user that you want to register.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
        /// </summary>
        [Input("iamArn")]
        public Input<string>? IamArn { get; set; }

        /// <summary>
        /// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
        /// </summary>
        [Input("identityType", required: true)]
        public Input<string> IdentityType { get; set; } = null!;

        /// <summary>
        /// The namespace. Currently, you should set this to `default`.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
        /// </summary>
        [Input("sessionName")]
        public Input<string>? SessionName { get; set; }

        /// <summary>
        /// The Amazon QuickSight user name that you want to create for the user you are registering.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
        /// </summary>
        [Input("userRole", required: true)]
        public Input<string> UserRole { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the user
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// The email address of the user that you want to register.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
        /// </summary>
        [Input("iamArn")]
        public Input<string>? IamArn { get; set; }

        /// <summary>
        /// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
        /// </summary>
        [Input("identityType")]
        public Input<string>? IdentityType { get; set; }

        /// <summary>
        /// The namespace. Currently, you should set this to `default`.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
        /// </summary>
        [Input("sessionName")]
        public Input<string>? SessionName { get; set; }

        /// <summary>
        /// The Amazon QuickSight user name that you want to create for the user you are registering.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
        /// </summary>
        [Input("userRole")]
        public Input<string>? UserRole { get; set; }

        public UserState()
        {
        }
    }
}

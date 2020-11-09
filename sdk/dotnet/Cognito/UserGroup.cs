// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    /// <summary>
    /// Provides a Cognito User Group resource.
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
    ///         var mainUserPool = new Aws.Cognito.UserPool("mainUserPool", new Aws.Cognito.UserPoolArgs
    ///         {
    ///         });
    ///         var groupRole = new Aws.Iam.Role("groupRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Sid"": """",
    ///       ""Effect"": ""Allow"",
    ///       ""Principal"": {
    ///         ""Federated"": ""cognito-identity.amazonaws.com""
    ///       },
    ///       ""Action"": ""sts:AssumeRoleWithWebIdentity"",
    ///       ""Condition"": {
    ///         ""StringEquals"": {
    ///           ""cognito-identity.amazonaws.com:aud"": ""us-east-1:12345678-dead-beef-cafe-123456790ab""
    ///         },
    ///         ""ForAnyValue:StringLike"": {
    ///           ""cognito-identity.amazonaws.com:amr"": ""authenticated""
    ///         }
    ///       }
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var mainUserGroup = new Aws.Cognito.UserGroup("mainUserGroup", new Aws.Cognito.UserGroupArgs
    ///         {
    ///             UserPoolId = mainUserPool.Id,
    ///             Description = "Managed by Pulumi",
    ///             Precedence = 42,
    ///             RoleArn = groupRole.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cognito User Groups can be imported using the `user_pool_id`/`name` attributes concatenated, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:cognito/userGroup:UserGroup group us-east-1_vG78M4goG/user-group
    /// ```
    /// </summary>
    public partial class UserGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the user group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the user group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The precedence of the user group.
        /// </summary>
        [Output("precedence")]
        public Output<int?> Precedence { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM role to be associated with the user group.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Output("userPoolId")]
        public Output<string> UserPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a UserGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGroup(string name, UserGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/userGroup:UserGroup", name, args ?? new UserGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGroup(string name, Input<string> id, UserGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/userGroup:UserGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGroup Get(string name, Input<string> id, UserGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGroup(name, id, state, options);
        }
    }

    public sealed class UserGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the user group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the user group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The precedence of the user group.
        /// </summary>
        [Input("precedence")]
        public Input<int>? Precedence { get; set; }

        /// <summary>
        /// The ARN of the IAM role to be associated with the user group.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public UserGroupArgs()
        {
        }
    }

    public sealed class UserGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the user group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the user group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The precedence of the user group.
        /// </summary>
        [Input("precedence")]
        public Input<int>? Precedence { get; set; }

        /// <summary>
        /// The ARN of the IAM role to be associated with the user group.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The user pool ID.
        /// </summary>
        [Input("userPoolId")]
        public Input<string>? UserPoolId { get; set; }

        public UserGroupState()
        {
        }
    }
}

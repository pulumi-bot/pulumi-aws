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
    /// &gt; **Note:** There is only a single policy allowed per AWS account. An existing policy will be lost when using this resource as an effect of this limitation.
    /// 
    /// Manages Password Policy for the AWS Account.
    /// See more about [Account Password Policy](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html)
    /// in the official AWS docs.
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
    ///         var strict = new Aws.Iam.AccountPasswordPolicy("strict", new Aws.Iam.AccountPasswordPolicyArgs
    ///         {
    ///             AllowUsersToChangePassword = true,
    ///             MinimumPasswordLength = 8,
    ///             RequireLowercaseCharacters = true,
    ///             RequireNumbers = true,
    ///             RequireSymbols = true,
    ///             RequireUppercaseCharacters = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM Account Password Policy can be imported using the word `iam-account-password-policy`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:iam/accountPasswordPolicy:AccountPasswordPolicy strict iam-account-password-policy
    /// ```
    /// </summary>
    public partial class AccountPasswordPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to allow users to change their own password
        /// </summary>
        [Output("allowUsersToChangePassword")]
        public Output<bool?> AllowUsersToChangePassword { get; private set; } = null!;

        /// <summary>
        /// Indicates whether passwords in the account expire.
        /// Returns `true` if `max_password_age` contains a value greater than `0`.
        /// Returns `false` if it is `0` or _not present_.
        /// </summary>
        [Output("expirePasswords")]
        public Output<bool> ExpirePasswords { get; private set; } = null!;

        /// <summary>
        /// Whether users are prevented from setting a new password after their password has expired
        /// (i.e. require administrator reset)
        /// </summary>
        [Output("hardExpiry")]
        public Output<bool> HardExpiry { get; private set; } = null!;

        /// <summary>
        /// The number of days that an user password is valid.
        /// </summary>
        [Output("maxPasswordAge")]
        public Output<int> MaxPasswordAge { get; private set; } = null!;

        /// <summary>
        /// Minimum length to require for user passwords.
        /// </summary>
        [Output("minimumPasswordLength")]
        public Output<int?> MinimumPasswordLength { get; private set; } = null!;

        /// <summary>
        /// The number of previous passwords that users are prevented from reusing.
        /// </summary>
        [Output("passwordReusePrevention")]
        public Output<int> PasswordReusePrevention { get; private set; } = null!;

        /// <summary>
        /// Whether to require lowercase characters for user passwords.
        /// </summary>
        [Output("requireLowercaseCharacters")]
        public Output<bool> RequireLowercaseCharacters { get; private set; } = null!;

        /// <summary>
        /// Whether to require numbers for user passwords.
        /// </summary>
        [Output("requireNumbers")]
        public Output<bool> RequireNumbers { get; private set; } = null!;

        /// <summary>
        /// Whether to require symbols for user passwords.
        /// </summary>
        [Output("requireSymbols")]
        public Output<bool> RequireSymbols { get; private set; } = null!;

        /// <summary>
        /// Whether to require uppercase characters for user passwords.
        /// </summary>
        [Output("requireUppercaseCharacters")]
        public Output<bool> RequireUppercaseCharacters { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPasswordPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPasswordPolicy(string name, AccountPasswordPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:iam/accountPasswordPolicy:AccountPasswordPolicy", name, args ?? new AccountPasswordPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountPasswordPolicy(string name, Input<string> id, AccountPasswordPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/accountPasswordPolicy:AccountPasswordPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountPasswordPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPasswordPolicy Get(string name, Input<string> id, AccountPasswordPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountPasswordPolicy(name, id, state, options);
        }
    }

    public sealed class AccountPasswordPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow users to change their own password
        /// </summary>
        [Input("allowUsersToChangePassword")]
        public Input<bool>? AllowUsersToChangePassword { get; set; }

        /// <summary>
        /// Whether users are prevented from setting a new password after their password has expired
        /// (i.e. require administrator reset)
        /// </summary>
        [Input("hardExpiry")]
        public Input<bool>? HardExpiry { get; set; }

        /// <summary>
        /// The number of days that an user password is valid.
        /// </summary>
        [Input("maxPasswordAge")]
        public Input<int>? MaxPasswordAge { get; set; }

        /// <summary>
        /// Minimum length to require for user passwords.
        /// </summary>
        [Input("minimumPasswordLength")]
        public Input<int>? MinimumPasswordLength { get; set; }

        /// <summary>
        /// The number of previous passwords that users are prevented from reusing.
        /// </summary>
        [Input("passwordReusePrevention")]
        public Input<int>? PasswordReusePrevention { get; set; }

        /// <summary>
        /// Whether to require lowercase characters for user passwords.
        /// </summary>
        [Input("requireLowercaseCharacters")]
        public Input<bool>? RequireLowercaseCharacters { get; set; }

        /// <summary>
        /// Whether to require numbers for user passwords.
        /// </summary>
        [Input("requireNumbers")]
        public Input<bool>? RequireNumbers { get; set; }

        /// <summary>
        /// Whether to require symbols for user passwords.
        /// </summary>
        [Input("requireSymbols")]
        public Input<bool>? RequireSymbols { get; set; }

        /// <summary>
        /// Whether to require uppercase characters for user passwords.
        /// </summary>
        [Input("requireUppercaseCharacters")]
        public Input<bool>? RequireUppercaseCharacters { get; set; }

        public AccountPasswordPolicyArgs()
        {
        }
    }

    public sealed class AccountPasswordPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow users to change their own password
        /// </summary>
        [Input("allowUsersToChangePassword")]
        public Input<bool>? AllowUsersToChangePassword { get; set; }

        /// <summary>
        /// Indicates whether passwords in the account expire.
        /// Returns `true` if `max_password_age` contains a value greater than `0`.
        /// Returns `false` if it is `0` or _not present_.
        /// </summary>
        [Input("expirePasswords")]
        public Input<bool>? ExpirePasswords { get; set; }

        /// <summary>
        /// Whether users are prevented from setting a new password after their password has expired
        /// (i.e. require administrator reset)
        /// </summary>
        [Input("hardExpiry")]
        public Input<bool>? HardExpiry { get; set; }

        /// <summary>
        /// The number of days that an user password is valid.
        /// </summary>
        [Input("maxPasswordAge")]
        public Input<int>? MaxPasswordAge { get; set; }

        /// <summary>
        /// Minimum length to require for user passwords.
        /// </summary>
        [Input("minimumPasswordLength")]
        public Input<int>? MinimumPasswordLength { get; set; }

        /// <summary>
        /// The number of previous passwords that users are prevented from reusing.
        /// </summary>
        [Input("passwordReusePrevention")]
        public Input<int>? PasswordReusePrevention { get; set; }

        /// <summary>
        /// Whether to require lowercase characters for user passwords.
        /// </summary>
        [Input("requireLowercaseCharacters")]
        public Input<bool>? RequireLowercaseCharacters { get; set; }

        /// <summary>
        /// Whether to require numbers for user passwords.
        /// </summary>
        [Input("requireNumbers")]
        public Input<bool>? RequireNumbers { get; set; }

        /// <summary>
        /// Whether to require symbols for user passwords.
        /// </summary>
        [Input("requireSymbols")]
        public Input<bool>? RequireSymbols { get; set; }

        /// <summary>
        /// Whether to require uppercase characters for user passwords.
        /// </summary>
        [Input("requireUppercaseCharacters")]
        public Input<bool>? RequireUppercaseCharacters { get; set; }

        public AccountPasswordPolicyState()
        {
        }
    }
}

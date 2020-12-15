// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub
{
    /// <summary>
    /// Provides a Security Hub member resource.
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
    ///         var exampleAccount = new Aws.SecurityHub.Account("exampleAccount", new Aws.SecurityHub.AccountArgs
    ///         {
    ///         });
    ///         var exampleMember = new Aws.SecurityHub.Member("exampleMember", new Aws.SecurityHub.MemberArgs
    ///         {
    ///             AccountId = "123456789012",
    ///             Email = "example@example.com",
    ///             Invite = true,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleAccount,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Security Hub members can be imported using their account ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:securityhub/member:Member example 123456789012
    /// ```
    /// </summary>
    [AwsResourceType("aws:securityhub/member:Member")]
    public partial class Member : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the member AWS account.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The email of the member AWS account.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
        /// </summary>
        [Output("invite")]
        public Output<bool?> Invite { get; private set; } = null!;

        /// <summary>
        /// The ID of the master Security Hub AWS account.
        /// </summary>
        [Output("masterId")]
        public Output<string> MasterId { get; private set; } = null!;

        /// <summary>
        /// The status of the member account relationship.
        /// </summary>
        [Output("memberStatus")]
        public Output<string> MemberStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Member resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Member(string name, MemberArgs args, CustomResourceOptions? options = null)
            : base("aws:securityhub/member:Member", name, args ?? new MemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Member(string name, Input<string> id, MemberState? state = null, CustomResourceOptions? options = null)
            : base("aws:securityhub/member:Member", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Member resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Member Get(string name, Input<string> id, MemberState? state = null, CustomResourceOptions? options = null)
        {
            return new Member(name, id, state, options);
        }
    }

    public sealed class MemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the member AWS account.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The email of the member AWS account.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
        /// </summary>
        [Input("invite")]
        public Input<bool>? Invite { get; set; }

        public MemberArgs()
        {
        }
    }

    public sealed class MemberState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the member AWS account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The email of the member AWS account.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
        /// </summary>
        [Input("invite")]
        public Input<bool>? Invite { get; set; }

        /// <summary>
        /// The ID of the master Security Hub AWS account.
        /// </summary>
        [Input("masterId")]
        public Input<string>? MasterId { get; set; }

        /// <summary>
        /// The status of the member account relationship.
        /// </summary>
        [Input("memberStatus")]
        public Input<string>? MemberStatus { get; set; }

        public MemberState()
        {
        }
    }
}

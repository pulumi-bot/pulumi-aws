// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeStarNotifications
{
    /// <summary>
    /// Provides a CodeStar Notifications Rule.
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
    ///         var code = new Aws.CodeCommit.Repository("code", new Aws.CodeCommit.RepositoryArgs
    ///         {
    ///             RepositoryName = "example-code-repo",
    ///         });
    ///         var notif = new Aws.Sns.Topic("notif", new Aws.Sns.TopicArgs
    ///         {
    ///         });
    ///         var notifAccess = notif.Arn.Apply(arn =&gt; Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Actions = 
    ///                     {
    ///                         "sns:Publish",
    ///                     },
    ///                     Principals = 
    ///                     {
    ///                         new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                         {
    ///                             Type = "Service",
    ///                             Identifiers = 
    ///                             {
    ///                                 "codestar-notifications.amazonaws.com",
    ///                             },
    ///                         },
    ///                     },
    ///                     Resources = 
    ///                     {
    ///                         arn,
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var @default = new Aws.Sns.TopicPolicy("default", new Aws.Sns.TopicPolicyArgs
    ///         {
    ///             Arn = notif.Arn,
    ///             Policy = notifAccess.Apply(notifAccess =&gt; notifAccess.Json),
    ///         });
    ///         var commits = new Aws.CodeStarNotifications.NotificationRule("commits", new Aws.CodeStarNotifications.NotificationRuleArgs
    ///         {
    ///             DetailType = "BASIC",
    ///             EventTypeIds = 
    ///             {
    ///                 "codecommit-repository-comments-on-commits",
    ///             },
    ///             Resource = code.Arn,
    ///             Targets = 
    ///             {
    ///                 new Aws.CodeStarNotifications.Inputs.NotificationRuleTargetArgs
    ///                 {
    ///                     Address = notif.Arn,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeStar notification rule can be imported using the ARN, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:codestarnotifications/notificationRule:NotificationRule foo arn:aws:codestar-notifications:us-west-1:0123456789:notificationrule/2cdc68a3-8f7c-4893-b6a5-45b362bd4f2b
    /// ```
    /// </summary>
    [AwsResourceType("aws:codestarnotifications/notificationRule:NotificationRule")]
    public partial class NotificationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The codestar notification rule ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
        /// </summary>
        [Output("detailType")]
        public Output<string> DetailType { get; private set; } = null!;

        /// <summary>
        /// A list of event types associated with this notification rule.
        /// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
        /// </summary>
        [Output("eventTypeIds")]
        public Output<ImmutableArray<string>> EventTypeIds { get; private set; } = null!;

        /// <summary>
        /// The name of notification rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the resource to associate with the notification rule.
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        /// <summary>
        /// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.NotificationRuleTarget>> Targets { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationRule(string name, NotificationRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:codestarnotifications/notificationRule:NotificationRule", name, args ?? new NotificationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationRule(string name, Input<string> id, NotificationRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:codestarnotifications/notificationRule:NotificationRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NotificationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationRule Get(string name, Input<string> id, NotificationRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new NotificationRule(name, id, state, options);
        }
    }

    public sealed class NotificationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
        /// </summary>
        [Input("detailType", required: true)]
        public Input<string> DetailType { get; set; } = null!;

        [Input("eventTypeIds", required: true)]
        private InputList<string>? _eventTypeIds;

        /// <summary>
        /// A list of event types associated with this notification rule.
        /// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
        /// </summary>
        public InputList<string> EventTypeIds
        {
            get => _eventTypeIds ?? (_eventTypeIds = new InputList<string>());
            set => _eventTypeIds = value;
        }

        /// <summary>
        /// The name of notification rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the resource to associate with the notification rule.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targets")]
        private InputList<Inputs.NotificationRuleTargetArgs>? _targets;

        /// <summary>
        /// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
        /// </summary>
        public InputList<Inputs.NotificationRuleTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.NotificationRuleTargetArgs>());
            set => _targets = value;
        }

        public NotificationRuleArgs()
        {
        }
    }

    public sealed class NotificationRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The codestar notification rule ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
        /// </summary>
        [Input("detailType")]
        public Input<string>? DetailType { get; set; }

        [Input("eventTypeIds")]
        private InputList<string>? _eventTypeIds;

        /// <summary>
        /// A list of event types associated with this notification rule.
        /// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
        /// </summary>
        public InputList<string> EventTypeIds
        {
            get => _eventTypeIds ?? (_eventTypeIds = new InputList<string>());
            set => _eventTypeIds = value;
        }

        /// <summary>
        /// The name of notification rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the resource to associate with the notification rule.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targets")]
        private InputList<Inputs.NotificationRuleTargetGetArgs>? _targets;

        /// <summary>
        /// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
        /// </summary>
        public InputList<Inputs.NotificationRuleTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.NotificationRuleTargetGetArgs>());
            set => _targets = value;
        }

        public NotificationRuleState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES receipt rule resource
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_receipt_rule.html.markdown.
    /// </summary>
    public partial class ReceiptRule : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of Add Header Action blocks. Documented below.
        /// </summary>
        [Output("addHeaderActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleAddHeaderActions>> AddHeaderActions { get; private set; } = null!;

        /// <summary>
        /// The name of the rule to place this rule after
        /// </summary>
        [Output("after")]
        public Output<string?> After { get; private set; } = null!;

        /// <summary>
        /// A list of Bounce Action blocks. Documented below.
        /// </summary>
        [Output("bounceActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleBounceActions>> BounceActions { get; private set; } = null!;

        /// <summary>
        /// If true, the rule will be enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// A list of Lambda Action blocks. Documented below.
        /// </summary>
        [Output("lambdaActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleLambdaActions>> LambdaActions { get; private set; } = null!;

        /// <summary>
        /// The name of the rule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of email addresses
        /// </summary>
        [Output("recipients")]
        public Output<ImmutableArray<string>> Recipients { get; private set; } = null!;

        /// <summary>
        /// The name of the rule set
        /// </summary>
        [Output("ruleSetName")]
        public Output<string> RuleSetName { get; private set; } = null!;

        /// <summary>
        /// A list of S3 Action blocks. Documented below.
        /// </summary>
        [Output("s3Actions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleS3Actions>> S3Actions { get; private set; } = null!;

        /// <summary>
        /// If true, incoming emails will be scanned for spam and viruses
        /// </summary>
        [Output("scanEnabled")]
        public Output<bool> ScanEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of SNS Action blocks. Documented below.
        /// </summary>
        [Output("snsActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleSnsActions>> SnsActions { get; private set; } = null!;

        /// <summary>
        /// A list of Stop Action blocks. Documented below.
        /// </summary>
        [Output("stopActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleStopActions>> StopActions { get; private set; } = null!;

        /// <summary>
        /// Require or Optional
        /// </summary>
        [Output("tlsPolicy")]
        public Output<string> TlsPolicy { get; private set; } = null!;

        /// <summary>
        /// A list of WorkMail Action blocks. Documented below.
        /// </summary>
        [Output("workmailActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleWorkmailActions>> WorkmailActions { get; private set; } = null!;


        /// <summary>
        /// Create a ReceiptRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReceiptRule(string name, ReceiptRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/receiptRule:ReceiptRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ReceiptRule(string name, Input<string> id, ReceiptRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/receiptRule:ReceiptRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReceiptRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReceiptRule Get(string name, Input<string> id, ReceiptRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ReceiptRule(name, id, state, options);
        }
    }

    public sealed class ReceiptRuleArgs : Pulumi.ResourceArgs
    {
        [Input("addHeaderActions")]
        private InputList<Inputs.ReceiptRuleAddHeaderActionsArgs>? _addHeaderActions;

        /// <summary>
        /// A list of Add Header Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleAddHeaderActionsArgs> AddHeaderActions
        {
            get => _addHeaderActions ?? (_addHeaderActions = new InputList<Inputs.ReceiptRuleAddHeaderActionsArgs>());
            set => _addHeaderActions = value;
        }

        /// <summary>
        /// The name of the rule to place this rule after
        /// </summary>
        [Input("after")]
        public Input<string>? After { get; set; }

        [Input("bounceActions")]
        private InputList<Inputs.ReceiptRuleBounceActionsArgs>? _bounceActions;

        /// <summary>
        /// A list of Bounce Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleBounceActionsArgs> BounceActions
        {
            get => _bounceActions ?? (_bounceActions = new InputList<Inputs.ReceiptRuleBounceActionsArgs>());
            set => _bounceActions = value;
        }

        /// <summary>
        /// If true, the rule will be enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("lambdaActions")]
        private InputList<Inputs.ReceiptRuleLambdaActionsArgs>? _lambdaActions;

        /// <summary>
        /// A list of Lambda Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleLambdaActionsArgs> LambdaActions
        {
            get => _lambdaActions ?? (_lambdaActions = new InputList<Inputs.ReceiptRuleLambdaActionsArgs>());
            set => _lambdaActions = value;
        }

        /// <summary>
        /// The name of the rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("recipients")]
        private InputList<string>? _recipients;

        /// <summary>
        /// A list of email addresses
        /// </summary>
        public InputList<string> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<string>());
            set => _recipients = value;
        }

        /// <summary>
        /// The name of the rule set
        /// </summary>
        [Input("ruleSetName", required: true)]
        public Input<string> RuleSetName { get; set; } = null!;

        [Input("s3Actions")]
        private InputList<Inputs.ReceiptRuleS3ActionsArgs>? _s3Actions;

        /// <summary>
        /// A list of S3 Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleS3ActionsArgs> S3Actions
        {
            get => _s3Actions ?? (_s3Actions = new InputList<Inputs.ReceiptRuleS3ActionsArgs>());
            set => _s3Actions = value;
        }

        /// <summary>
        /// If true, incoming emails will be scanned for spam and viruses
        /// </summary>
        [Input("scanEnabled")]
        public Input<bool>? ScanEnabled { get; set; }

        [Input("snsActions")]
        private InputList<Inputs.ReceiptRuleSnsActionsArgs>? _snsActions;

        /// <summary>
        /// A list of SNS Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleSnsActionsArgs> SnsActions
        {
            get => _snsActions ?? (_snsActions = new InputList<Inputs.ReceiptRuleSnsActionsArgs>());
            set => _snsActions = value;
        }

        [Input("stopActions")]
        private InputList<Inputs.ReceiptRuleStopActionsArgs>? _stopActions;

        /// <summary>
        /// A list of Stop Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleStopActionsArgs> StopActions
        {
            get => _stopActions ?? (_stopActions = new InputList<Inputs.ReceiptRuleStopActionsArgs>());
            set => _stopActions = value;
        }

        /// <summary>
        /// Require or Optional
        /// </summary>
        [Input("tlsPolicy")]
        public Input<string>? TlsPolicy { get; set; }

        [Input("workmailActions")]
        private InputList<Inputs.ReceiptRuleWorkmailActionsArgs>? _workmailActions;

        /// <summary>
        /// A list of WorkMail Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleWorkmailActionsArgs> WorkmailActions
        {
            get => _workmailActions ?? (_workmailActions = new InputList<Inputs.ReceiptRuleWorkmailActionsArgs>());
            set => _workmailActions = value;
        }

        public ReceiptRuleArgs()
        {
        }
    }

    public sealed class ReceiptRuleState : Pulumi.ResourceArgs
    {
        [Input("addHeaderActions")]
        private InputList<Inputs.ReceiptRuleAddHeaderActionsGetArgs>? _addHeaderActions;

        /// <summary>
        /// A list of Add Header Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleAddHeaderActionsGetArgs> AddHeaderActions
        {
            get => _addHeaderActions ?? (_addHeaderActions = new InputList<Inputs.ReceiptRuleAddHeaderActionsGetArgs>());
            set => _addHeaderActions = value;
        }

        /// <summary>
        /// The name of the rule to place this rule after
        /// </summary>
        [Input("after")]
        public Input<string>? After { get; set; }

        [Input("bounceActions")]
        private InputList<Inputs.ReceiptRuleBounceActionsGetArgs>? _bounceActions;

        /// <summary>
        /// A list of Bounce Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleBounceActionsGetArgs> BounceActions
        {
            get => _bounceActions ?? (_bounceActions = new InputList<Inputs.ReceiptRuleBounceActionsGetArgs>());
            set => _bounceActions = value;
        }

        /// <summary>
        /// If true, the rule will be enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("lambdaActions")]
        private InputList<Inputs.ReceiptRuleLambdaActionsGetArgs>? _lambdaActions;

        /// <summary>
        /// A list of Lambda Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleLambdaActionsGetArgs> LambdaActions
        {
            get => _lambdaActions ?? (_lambdaActions = new InputList<Inputs.ReceiptRuleLambdaActionsGetArgs>());
            set => _lambdaActions = value;
        }

        /// <summary>
        /// The name of the rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("recipients")]
        private InputList<string>? _recipients;

        /// <summary>
        /// A list of email addresses
        /// </summary>
        public InputList<string> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<string>());
            set => _recipients = value;
        }

        /// <summary>
        /// The name of the rule set
        /// </summary>
        [Input("ruleSetName")]
        public Input<string>? RuleSetName { get; set; }

        [Input("s3Actions")]
        private InputList<Inputs.ReceiptRuleS3ActionsGetArgs>? _s3Actions;

        /// <summary>
        /// A list of S3 Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleS3ActionsGetArgs> S3Actions
        {
            get => _s3Actions ?? (_s3Actions = new InputList<Inputs.ReceiptRuleS3ActionsGetArgs>());
            set => _s3Actions = value;
        }

        /// <summary>
        /// If true, incoming emails will be scanned for spam and viruses
        /// </summary>
        [Input("scanEnabled")]
        public Input<bool>? ScanEnabled { get; set; }

        [Input("snsActions")]
        private InputList<Inputs.ReceiptRuleSnsActionsGetArgs>? _snsActions;

        /// <summary>
        /// A list of SNS Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleSnsActionsGetArgs> SnsActions
        {
            get => _snsActions ?? (_snsActions = new InputList<Inputs.ReceiptRuleSnsActionsGetArgs>());
            set => _snsActions = value;
        }

        [Input("stopActions")]
        private InputList<Inputs.ReceiptRuleStopActionsGetArgs>? _stopActions;

        /// <summary>
        /// A list of Stop Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleStopActionsGetArgs> StopActions
        {
            get => _stopActions ?? (_stopActions = new InputList<Inputs.ReceiptRuleStopActionsGetArgs>());
            set => _stopActions = value;
        }

        /// <summary>
        /// Require or Optional
        /// </summary>
        [Input("tlsPolicy")]
        public Input<string>? TlsPolicy { get; set; }

        [Input("workmailActions")]
        private InputList<Inputs.ReceiptRuleWorkmailActionsGetArgs>? _workmailActions;

        /// <summary>
        /// A list of WorkMail Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleWorkmailActionsGetArgs> WorkmailActions
        {
            get => _workmailActions ?? (_workmailActions = new InputList<Inputs.ReceiptRuleWorkmailActionsGetArgs>());
            set => _workmailActions = value;
        }

        public ReceiptRuleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ReceiptRuleAddHeaderActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the header to add
        /// </summary>
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        /// <summary>
        /// The value of the header to add
        /// </summary>
        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        public ReceiptRuleAddHeaderActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleAddHeaderActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the header to add
        /// </summary>
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        /// <summary>
        /// The value of the header to add
        /// </summary>
        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        public ReceiptRuleAddHeaderActionsGetArgs()
        {
        }
    }

    public sealed class ReceiptRuleBounceActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The message to send
        /// </summary>
        [Input("message", required: true)]
        public Input<string> Message { get; set; } = null!;

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The email address of the sender
        /// </summary>
        [Input("sender", required: true)]
        public Input<string> Sender { get; set; } = null!;

        /// <summary>
        /// The RFC 5321 SMTP reply code
        /// </summary>
        [Input("smtpReplyCode", required: true)]
        public Input<string> SmtpReplyCode { get; set; } = null!;

        /// <summary>
        /// The RFC 3463 SMTP enhanced status code
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleBounceActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleBounceActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The message to send
        /// </summary>
        [Input("message", required: true)]
        public Input<string> Message { get; set; } = null!;

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The email address of the sender
        /// </summary>
        [Input("sender", required: true)]
        public Input<string> Sender { get; set; } = null!;

        /// <summary>
        /// The RFC 5321 SMTP reply code
        /// </summary>
        [Input("smtpReplyCode", required: true)]
        public Input<string> SmtpReplyCode { get; set; } = null!;

        /// <summary>
        /// The RFC 3463 SMTP enhanced status code
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleBounceActionsGetArgs()
        {
        }
    }

    public sealed class ReceiptRuleLambdaActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Lambda function to invoke
        /// </summary>
        [Input("functionArn", required: true)]
        public Input<string> FunctionArn { get; set; } = null!;

        /// <summary>
        /// Event or RequestResponse
        /// </summary>
        [Input("invocationType")]
        public Input<string>? InvocationType { get; set; }

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleLambdaActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleLambdaActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Lambda function to invoke
        /// </summary>
        [Input("functionArn", required: true)]
        public Input<string> FunctionArn { get; set; } = null!;

        /// <summary>
        /// Event or RequestResponse
        /// </summary>
        [Input("invocationType")]
        public Input<string>? InvocationType { get; set; }

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleLambdaActionsGetArgs()
        {
        }
    }

    public sealed class ReceiptRuleS3ActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the S3 bucket
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The ARN of the KMS key
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The key prefix of the S3 bucket
        /// </summary>
        [Input("objectKeyPrefix")]
        public Input<string>? ObjectKeyPrefix { get; set; }

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleS3ActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleS3ActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the S3 bucket
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The ARN of the KMS key
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The key prefix of the S3 bucket
        /// </summary>
        [Input("objectKeyPrefix")]
        public Input<string>? ObjectKeyPrefix { get; set; }

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleS3ActionsGetArgs()
        {
        }
    }

    public sealed class ReceiptRuleSnsActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn", required: true)]
        public Input<string> TopicArn { get; set; } = null!;

        public ReceiptRuleSnsActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleSnsActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn", required: true)]
        public Input<string> TopicArn { get; set; } = null!;

        public ReceiptRuleSnsActionsGetArgs()
        {
        }
    }

    public sealed class ReceiptRuleStopActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The scope to apply
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleStopActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleStopActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The scope to apply
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleStopActionsGetArgs()
        {
        }
    }

    public sealed class ReceiptRuleWorkmailActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the WorkMail organization
        /// </summary>
        [Input("organizationArn", required: true)]
        public Input<string> OrganizationArn { get; set; } = null!;

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleWorkmailActionsArgs()
        {
        }
    }

    public sealed class ReceiptRuleWorkmailActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the WorkMail organization
        /// </summary>
        [Input("organizationArn", required: true)]
        public Input<string> OrganizationArn { get; set; } = null!;

        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleWorkmailActionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ReceiptRuleAddHeaderActions
    {
        /// <summary>
        /// The name of the header to add
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// The value of the header to add
        /// </summary>
        public readonly string HeaderValue;
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;

        [OutputConstructor]
        private ReceiptRuleAddHeaderActions(
            string headerName,
            string headerValue,
            int position)
        {
            HeaderName = headerName;
            HeaderValue = headerValue;
            Position = position;
        }
    }

    [OutputType]
    public sealed class ReceiptRuleBounceActions
    {
        /// <summary>
        /// The message to send
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;
        /// <summary>
        /// The email address of the sender
        /// </summary>
        public readonly string Sender;
        /// <summary>
        /// The RFC 5321 SMTP reply code
        /// </summary>
        public readonly string SmtpReplyCode;
        /// <summary>
        /// The RFC 3463 SMTP enhanced status code
        /// </summary>
        public readonly string? StatusCode;
        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleBounceActions(
            string message,
            int position,
            string sender,
            string smtpReplyCode,
            string? statusCode,
            string? topicArn)
        {
            Message = message;
            Position = position;
            Sender = sender;
            SmtpReplyCode = smtpReplyCode;
            StatusCode = statusCode;
            TopicArn = topicArn;
        }
    }

    [OutputType]
    public sealed class ReceiptRuleLambdaActions
    {
        /// <summary>
        /// The ARN of the Lambda function to invoke
        /// </summary>
        public readonly string FunctionArn;
        /// <summary>
        /// Event or RequestResponse
        /// </summary>
        public readonly string InvocationType;
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;
        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleLambdaActions(
            string functionArn,
            string invocationType,
            int position,
            string? topicArn)
        {
            FunctionArn = functionArn;
            InvocationType = invocationType;
            Position = position;
            TopicArn = topicArn;
        }
    }

    [OutputType]
    public sealed class ReceiptRuleS3Actions
    {
        /// <summary>
        /// The name of the S3 bucket
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// The ARN of the KMS key
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// The key prefix of the S3 bucket
        /// </summary>
        public readonly string? ObjectKeyPrefix;
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;
        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleS3Actions(
            string bucketName,
            string? kmsKeyArn,
            string? objectKeyPrefix,
            int position,
            string? topicArn)
        {
            BucketName = bucketName;
            KmsKeyArn = kmsKeyArn;
            ObjectKeyPrefix = objectKeyPrefix;
            Position = position;
            TopicArn = topicArn;
        }
    }

    [OutputType]
    public sealed class ReceiptRuleSnsActions
    {
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;
        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        public readonly string TopicArn;

        [OutputConstructor]
        private ReceiptRuleSnsActions(
            int position,
            string topicArn)
        {
            Position = position;
            TopicArn = topicArn;
        }
    }

    [OutputType]
    public sealed class ReceiptRuleStopActions
    {
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;
        /// <summary>
        /// The scope to apply
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleStopActions(
            int position,
            string scope,
            string? topicArn)
        {
            Position = position;
            Scope = scope;
            TopicArn = topicArn;
        }
    }

    [OutputType]
    public sealed class ReceiptRuleWorkmailActions
    {
        /// <summary>
        /// The ARN of the WorkMail organization
        /// </summary>
        public readonly string OrganizationArn;
        /// <summary>
        /// The position of the action in the receipt rule
        /// </summary>
        public readonly int Position;
        /// <summary>
        /// The ARN of an SNS topic to notify
        /// </summary>
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleWorkmailActions(
            string organizationArn,
            int position,
            string? topicArn)
        {
            OrganizationArn = organizationArn;
            Position = position;
            TopicArn = topicArn;
        }
    }
    }
}

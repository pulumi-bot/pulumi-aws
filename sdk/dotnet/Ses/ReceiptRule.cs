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
    /// </summary>
    public partial class ReceiptRule : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of Add Header Action blocks. Documented below.
        /// </summary>
        [Output("addHeaderActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleAddHeaderAction>> AddHeaderActions { get; private set; } = null!;

        /// <summary>
        /// The name of the rule to place this rule after
        /// </summary>
        [Output("after")]
        public Output<string?> After { get; private set; } = null!;

        /// <summary>
        /// A list of Bounce Action blocks. Documented below.
        /// </summary>
        [Output("bounceActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleBounceAction>> BounceActions { get; private set; } = null!;

        /// <summary>
        /// If true, the rule will be enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// A list of Lambda Action blocks. Documented below.
        /// </summary>
        [Output("lambdaActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleLambdaAction>> LambdaActions { get; private set; } = null!;

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
        public Output<ImmutableArray<Outputs.ReceiptRuleS3Action>> S3Actions { get; private set; } = null!;

        /// <summary>
        /// If true, incoming emails will be scanned for spam and viruses
        /// </summary>
        [Output("scanEnabled")]
        public Output<bool> ScanEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of SNS Action blocks. Documented below.
        /// </summary>
        [Output("snsActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleSnsAction>> SnsActions { get; private set; } = null!;

        /// <summary>
        /// A list of Stop Action blocks. Documented below.
        /// </summary>
        [Output("stopActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleStopAction>> StopActions { get; private set; } = null!;

        /// <summary>
        /// Require or Optional
        /// </summary>
        [Output("tlsPolicy")]
        public Output<string> TlsPolicy { get; private set; } = null!;

        /// <summary>
        /// A list of WorkMail Action blocks. Documented below.
        /// </summary>
        [Output("workmailActions")]
        public Output<ImmutableArray<Outputs.ReceiptRuleWorkmailAction>> WorkmailActions { get; private set; } = null!;


        /// <summary>
        /// Create a ReceiptRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReceiptRule(string name, ReceiptRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/receiptRule:ReceiptRule", name, args ?? new ReceiptRuleArgs(), MakeResourceOptions(options, ""))
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
        private InputList<Inputs.ReceiptRuleAddHeaderActionArgs>? _addHeaderActions;

        /// <summary>
        /// A list of Add Header Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleAddHeaderActionArgs> AddHeaderActions
        {
            get => _addHeaderActions ?? (_addHeaderActions = new InputList<Inputs.ReceiptRuleAddHeaderActionArgs>());
            set => _addHeaderActions = value;
        }

        /// <summary>
        /// The name of the rule to place this rule after
        /// </summary>
        [Input("after")]
        public Input<string>? After { get; set; }

        [Input("bounceActions")]
        private InputList<Inputs.ReceiptRuleBounceActionArgs>? _bounceActions;

        /// <summary>
        /// A list of Bounce Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleBounceActionArgs> BounceActions
        {
            get => _bounceActions ?? (_bounceActions = new InputList<Inputs.ReceiptRuleBounceActionArgs>());
            set => _bounceActions = value;
        }

        /// <summary>
        /// If true, the rule will be enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("lambdaActions")]
        private InputList<Inputs.ReceiptRuleLambdaActionArgs>? _lambdaActions;

        /// <summary>
        /// A list of Lambda Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleLambdaActionArgs> LambdaActions
        {
            get => _lambdaActions ?? (_lambdaActions = new InputList<Inputs.ReceiptRuleLambdaActionArgs>());
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
        private InputList<Inputs.ReceiptRuleS3ActionArgs>? _s3Actions;

        /// <summary>
        /// A list of S3 Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleS3ActionArgs> S3Actions
        {
            get => _s3Actions ?? (_s3Actions = new InputList<Inputs.ReceiptRuleS3ActionArgs>());
            set => _s3Actions = value;
        }

        /// <summary>
        /// If true, incoming emails will be scanned for spam and viruses
        /// </summary>
        [Input("scanEnabled")]
        public Input<bool>? ScanEnabled { get; set; }

        [Input("snsActions")]
        private InputList<Inputs.ReceiptRuleSnsActionArgs>? _snsActions;

        /// <summary>
        /// A list of SNS Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleSnsActionArgs> SnsActions
        {
            get => _snsActions ?? (_snsActions = new InputList<Inputs.ReceiptRuleSnsActionArgs>());
            set => _snsActions = value;
        }

        [Input("stopActions")]
        private InputList<Inputs.ReceiptRuleStopActionArgs>? _stopActions;

        /// <summary>
        /// A list of Stop Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleStopActionArgs> StopActions
        {
            get => _stopActions ?? (_stopActions = new InputList<Inputs.ReceiptRuleStopActionArgs>());
            set => _stopActions = value;
        }

        /// <summary>
        /// Require or Optional
        /// </summary>
        [Input("tlsPolicy")]
        public Input<string>? TlsPolicy { get; set; }

        [Input("workmailActions")]
        private InputList<Inputs.ReceiptRuleWorkmailActionArgs>? _workmailActions;

        /// <summary>
        /// A list of WorkMail Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleWorkmailActionArgs> WorkmailActions
        {
            get => _workmailActions ?? (_workmailActions = new InputList<Inputs.ReceiptRuleWorkmailActionArgs>());
            set => _workmailActions = value;
        }

        public ReceiptRuleArgs()
        {
        }
    }

    public sealed class ReceiptRuleState : Pulumi.ResourceArgs
    {
        [Input("addHeaderActions")]
        private InputList<Inputs.ReceiptRuleAddHeaderActionGetArgs>? _addHeaderActions;

        /// <summary>
        /// A list of Add Header Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleAddHeaderActionGetArgs> AddHeaderActions
        {
            get => _addHeaderActions ?? (_addHeaderActions = new InputList<Inputs.ReceiptRuleAddHeaderActionGetArgs>());
            set => _addHeaderActions = value;
        }

        /// <summary>
        /// The name of the rule to place this rule after
        /// </summary>
        [Input("after")]
        public Input<string>? After { get; set; }

        [Input("bounceActions")]
        private InputList<Inputs.ReceiptRuleBounceActionGetArgs>? _bounceActions;

        /// <summary>
        /// A list of Bounce Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleBounceActionGetArgs> BounceActions
        {
            get => _bounceActions ?? (_bounceActions = new InputList<Inputs.ReceiptRuleBounceActionGetArgs>());
            set => _bounceActions = value;
        }

        /// <summary>
        /// If true, the rule will be enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("lambdaActions")]
        private InputList<Inputs.ReceiptRuleLambdaActionGetArgs>? _lambdaActions;

        /// <summary>
        /// A list of Lambda Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleLambdaActionGetArgs> LambdaActions
        {
            get => _lambdaActions ?? (_lambdaActions = new InputList<Inputs.ReceiptRuleLambdaActionGetArgs>());
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
        private InputList<Inputs.ReceiptRuleS3ActionGetArgs>? _s3Actions;

        /// <summary>
        /// A list of S3 Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleS3ActionGetArgs> S3Actions
        {
            get => _s3Actions ?? (_s3Actions = new InputList<Inputs.ReceiptRuleS3ActionGetArgs>());
            set => _s3Actions = value;
        }

        /// <summary>
        /// If true, incoming emails will be scanned for spam and viruses
        /// </summary>
        [Input("scanEnabled")]
        public Input<bool>? ScanEnabled { get; set; }

        [Input("snsActions")]
        private InputList<Inputs.ReceiptRuleSnsActionGetArgs>? _snsActions;

        /// <summary>
        /// A list of SNS Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleSnsActionGetArgs> SnsActions
        {
            get => _snsActions ?? (_snsActions = new InputList<Inputs.ReceiptRuleSnsActionGetArgs>());
            set => _snsActions = value;
        }

        [Input("stopActions")]
        private InputList<Inputs.ReceiptRuleStopActionGetArgs>? _stopActions;

        /// <summary>
        /// A list of Stop Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleStopActionGetArgs> StopActions
        {
            get => _stopActions ?? (_stopActions = new InputList<Inputs.ReceiptRuleStopActionGetArgs>());
            set => _stopActions = value;
        }

        /// <summary>
        /// Require or Optional
        /// </summary>
        [Input("tlsPolicy")]
        public Input<string>? TlsPolicy { get; set; }

        [Input("workmailActions")]
        private InputList<Inputs.ReceiptRuleWorkmailActionGetArgs>? _workmailActions;

        /// <summary>
        /// A list of WorkMail Action blocks. Documented below.
        /// </summary>
        public InputList<Inputs.ReceiptRuleWorkmailActionGetArgs> WorkmailActions
        {
            get => _workmailActions ?? (_workmailActions = new InputList<Inputs.ReceiptRuleWorkmailActionGetArgs>());
            set => _workmailActions = value;
        }

        public ReceiptRuleState()
        {
        }
    }
}

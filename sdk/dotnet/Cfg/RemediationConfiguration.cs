// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg
{
    /// <summary>
    /// Provides an AWS Config Remediation Configuration.
    /// 
    /// &gt; **Note:** Config Remediation Configuration requires an existing [Config Rule](https://www.terraform.io/docs/providers/aws/r/config_config_rule.html) to be present.
    /// 
    /// ## Example Usage
    /// 
    /// AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var thisRule = new Aws.Cfg.Rule("thisRule", new Aws.Cfg.RuleArgs
    ///         {
    ///             Source = new Aws.Cfg.Inputs.RuleSourceArgs
    ///             {
    ///                 Owner = "AWS",
    ///                 SourceIdentifier = "S3_BUCKET_VERSIONING_ENABLED",
    ///             },
    ///         });
    ///         var thisRemediationConfiguration = new Aws.Cfg.RemediationConfiguration("thisRemediationConfiguration", new Aws.Cfg.RemediationConfigurationArgs
    ///         {
    ///             ConfigRuleName = thisRule.Name,
    ///             ResourceType = "AWS::S3::Bucket",
    ///             TargetType = "SSM_DOCUMENT",
    ///             TargetId = "AWS-EnableS3BucketEncryption",
    ///             TargetVersion = "1",
    ///             Parameters = 
    ///             {
    ///                 new Aws.Cfg.Inputs.RemediationConfigurationParameterArgs
    ///                 {
    ///                     Name = "AutomationAssumeRole",
    ///                     StaticValue = "arn:aws:iam::875924563244:role/security_config",
    ///                 },
    ///                 new Aws.Cfg.Inputs.RemediationConfigurationParameterArgs
    ///                 {
    ///                     Name = "BucketName",
    ///                     ResourceValue = "RESOURCE_ID",
    ///                 },
    ///                 new Aws.Cfg.Inputs.RemediationConfigurationParameterArgs
    ///                 {
    ///                     Name = "SSEAlgorithm",
    ///                     StaticValue = "AES256",
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
    /// Remediation Configurations can be imported using the name config_rule_name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:cfg/remediationConfiguration:RemediationConfiguration this example
    /// ```
    /// </summary>
    [AwsResourceType("aws:cfg/remediationConfiguration:RemediationConfiguration")]
    public partial class RemediationConfiguration : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the AWS Config rule
        /// </summary>
        [Output("configRuleName")]
        public Output<string> ConfigRuleName { get; private set; } = null!;

        /// <summary>
        /// Can be specified multiple times for each
        /// parameter. Each parameter block supports fields documented below.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.RemediationConfigurationParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// The type of a resource
        /// </summary>
        [Output("resourceType")]
        public Output<string?> ResourceType { get; private set; } = null!;

        /// <summary>
        /// Target ID is the name of the public document
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;

        /// <summary>
        /// The type of the target. Target executes remediation. For example, SSM document
        /// </summary>
        [Output("targetType")]
        public Output<string> TargetType { get; private set; } = null!;

        /// <summary>
        /// Version of the target. For example, version of the SSM document
        /// </summary>
        [Output("targetVersion")]
        public Output<string?> TargetVersion { get; private set; } = null!;


        /// <summary>
        /// Create a RemediationConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RemediationConfiguration(string name, RemediationConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:cfg/remediationConfiguration:RemediationConfiguration", name, args ?? new RemediationConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RemediationConfiguration(string name, Input<string> id, RemediationConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:cfg/remediationConfiguration:RemediationConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RemediationConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RemediationConfiguration Get(string name, Input<string> id, RemediationConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new RemediationConfiguration(name, id, state, options);
        }
    }

    public sealed class RemediationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the AWS Config rule
        /// </summary>
        [Input("configRuleName", required: true)]
        public Input<string> ConfigRuleName { get; set; } = null!;

        [Input("parameters")]
        private InputList<Inputs.RemediationConfigurationParameterArgs>? _parameters;

        /// <summary>
        /// Can be specified multiple times for each
        /// parameter. Each parameter block supports fields documented below.
        /// </summary>
        public InputList<Inputs.RemediationConfigurationParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RemediationConfigurationParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The type of a resource
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// Target ID is the name of the public document
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        /// <summary>
        /// The type of the target. Target executes remediation. For example, SSM document
        /// </summary>
        [Input("targetType", required: true)]
        public Input<string> TargetType { get; set; } = null!;

        /// <summary>
        /// Version of the target. For example, version of the SSM document
        /// </summary>
        [Input("targetVersion")]
        public Input<string>? TargetVersion { get; set; }

        public RemediationConfigurationArgs()
        {
        }
    }

    public sealed class RemediationConfigurationState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the AWS Config rule
        /// </summary>
        [Input("configRuleName")]
        public Input<string>? ConfigRuleName { get; set; }

        [Input("parameters")]
        private InputList<Inputs.RemediationConfigurationParameterGetArgs>? _parameters;

        /// <summary>
        /// Can be specified multiple times for each
        /// parameter. Each parameter block supports fields documented below.
        /// </summary>
        public InputList<Inputs.RemediationConfigurationParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RemediationConfigurationParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The type of a resource
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// Target ID is the name of the public document
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        /// <summary>
        /// The type of the target. Target executes remediation. For example, SSM document
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        /// <summary>
        /// Version of the target. For example, version of the SSM document
        /// </summary>
        [Input("targetVersion")]
        public Input<string>? TargetVersion { get; set; }

        public RemediationConfigurationState()
        {
        }
    }
}

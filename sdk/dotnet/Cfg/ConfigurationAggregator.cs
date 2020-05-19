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
    /// Manages an AWS Config Configuration Aggregator
    /// 
    /// ## Example Usage
    /// 
    /// ### Account Based Aggregation
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var account = new Aws.Cfg.ConfigurationAggregator("account", new Aws.Cfg.ConfigurationAggregatorArgs
    ///         {
    ///             AccountAggregationSource = new Aws.Cfg.Inputs.ConfigurationAggregatorAccountAggregationSourceArgs
    ///             {
    ///                 AccountIds = 
    ///                 {
    ///                     "123456789012",
    ///                 },
    ///                 Regions = 
    ///                 {
    ///                     "us-west-2",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Organization Based Aggregation
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var organizationRole = new Aws.Iam.Role("organizationRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Sid"": """",
    ///       ""Effect"": ""Allow"",
    ///       ""Principal"": {
    ///         ""Service"": ""config.amazonaws.com""
    ///       },
    ///       ""Action"": ""sts:AssumeRole""
    ///     }
    ///   ]
    /// }
    /// 
    /// ",
    ///         });
    ///         var organizationConfigurationAggregator = new Aws.Cfg.ConfigurationAggregator("organizationConfigurationAggregator", new Aws.Cfg.ConfigurationAggregatorArgs
    ///         {
    ///             OrganizationAggregationSource = new Aws.Cfg.Inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs
    ///             {
    ///                 AllRegions = true,
    ///                 RoleArn = organizationRole.Arn,
    ///             },
    ///         });
    ///         var organizationRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("organizationRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations",
    ///             Role = organizationRole.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ConfigurationAggregator : Pulumi.CustomResource
    {
        /// <summary>
        /// The account(s) to aggregate config data from as documented below.
        /// </summary>
        [Output("accountAggregationSource")]
        public Output<Outputs.ConfigurationAggregatorAccountAggregationSource?> AccountAggregationSource { get; private set; } = null!;

        /// <summary>
        /// The ARN of the aggregator
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the configuration aggregator.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization to aggregate config data from as documented below.
        /// </summary>
        [Output("organizationAggregationSource")]
        public Output<Outputs.ConfigurationAggregatorOrganizationAggregationSource?> OrganizationAggregationSource { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationAggregator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationAggregator(string name, ConfigurationAggregatorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:cfg/configurationAggregator:ConfigurationAggregator", name, args ?? new ConfigurationAggregatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationAggregator(string name, Input<string> id, ConfigurationAggregatorState? state = null, CustomResourceOptions? options = null)
            : base("aws:cfg/configurationAggregator:ConfigurationAggregator", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationAggregator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationAggregator Get(string name, Input<string> id, ConfigurationAggregatorState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigurationAggregator(name, id, state, options);
        }
    }

    public sealed class ConfigurationAggregatorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account(s) to aggregate config data from as documented below.
        /// </summary>
        [Input("accountAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorAccountAggregationSourceArgs>? AccountAggregationSource { get; set; }

        /// <summary>
        /// The name of the configuration aggregator.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization to aggregate config data from as documented below.
        /// </summary>
        [Input("organizationAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs>? OrganizationAggregationSource { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ConfigurationAggregatorArgs()
        {
        }
    }

    public sealed class ConfigurationAggregatorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account(s) to aggregate config data from as documented below.
        /// </summary>
        [Input("accountAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorAccountAggregationSourceGetArgs>? AccountAggregationSource { get; set; }

        /// <summary>
        /// The ARN of the aggregator
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the configuration aggregator.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization to aggregate config data from as documented below.
        /// </summary>
        [Input("organizationAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorOrganizationAggregationSourceGetArgs>? OrganizationAggregationSource { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ConfigurationAggregatorState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk
{
    /// <summary>
    /// Provides an Elastic Beanstalk Configuration Template, which are associated with
    /// a specific application and are used to deploy different versions of the
    /// application with the same configuration settings.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// ## Option Settings
    /// 
    /// The `setting` field supports the following format:
    /// 
    /// * `namespace` - unique namespace identifying the option's associated AWS resource
    /// * `name` - name of the configuration option
    /// * `value` - value for the configuration option
    /// * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastic_beanstalk_configuration_template.html.markdown.
    /// </summary>
    public partial class ConfigurationTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// name of the application to associate with this configuration template
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// Short description of the Template
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the environment used with this configuration template
        /// </summary>
        [Output("environmentId")]
        public Output<string?> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// A unique name for this Template.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Option settings to configure the new Environment. These
        /// override specific values that are set as defaults. The format is detailed
        /// below in Option Settings
        /// </summary>
        [Output("settings")]
        public Output<ImmutableArray<Outputs.ConfigurationTemplateSettings>> Settings { get; private set; } = null!;

        /// <summary>
        /// A solution stack to base your Template
        /// off of. Example stacks can be found in the [Amazon API documentation][1]
        /// </summary>
        [Output("solutionStackName")]
        public Output<string?> SolutionStackName { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationTemplate(string name, ConfigurationTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationTemplate(string name, Input<string> id, ConfigurationTemplateState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationTemplate Get(string name, Input<string> id, ConfigurationTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigurationTemplate(name, id, state, options);
        }
    }

    public sealed class ConfigurationTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the application to associate with this configuration template
        /// </summary>
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        /// <summary>
        /// Short description of the Template
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the environment used with this configuration template
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// A unique name for this Template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("settings")]
        private InputList<Inputs.ConfigurationTemplateSettingsArgs>? _settings;

        /// <summary>
        /// Option settings to configure the new Environment. These
        /// override specific values that are set as defaults. The format is detailed
        /// below in Option Settings
        /// </summary>
        public InputList<Inputs.ConfigurationTemplateSettingsArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.ConfigurationTemplateSettingsArgs>());
            set => _settings = value;
        }

        /// <summary>
        /// A solution stack to base your Template
        /// off of. Example stacks can be found in the [Amazon API documentation][1]
        /// </summary>
        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        public ConfigurationTemplateArgs()
        {
        }
    }

    public sealed class ConfigurationTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the application to associate with this configuration template
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Short description of the Template
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the environment used with this configuration template
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// A unique name for this Template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("settings")]
        private InputList<Inputs.ConfigurationTemplateSettingsGetArgs>? _settings;

        /// <summary>
        /// Option settings to configure the new Environment. These
        /// override specific values that are set as defaults. The format is detailed
        /// below in Option Settings
        /// </summary>
        public InputList<Inputs.ConfigurationTemplateSettingsGetArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.ConfigurationTemplateSettingsGetArgs>());
            set => _settings = value;
        }

        /// <summary>
        /// A solution stack to base your Template
        /// off of. Example stacks can be found in the [Amazon API documentation][1]
        /// </summary>
        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        public ConfigurationTemplateState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ConfigurationTemplateSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name for this Template.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("resource")]
        public Input<string>? Resource { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ConfigurationTemplateSettingsArgs()
        {
        }
    }

    public sealed class ConfigurationTemplateSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name for this Template.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("resource")]
        public Input<string>? Resource { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ConfigurationTemplateSettingsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ConfigurationTemplateSettings
    {
        /// <summary>
        /// A unique name for this Template.
        /// </summary>
        public readonly string Name;
        public readonly string Namespace;
        public readonly string? Resource;
        public readonly string Value;

        [OutputConstructor]
        private ConfigurationTemplateSettings(
            string name,
            string @namespace,
            string? resource,
            string value)
        {
            Name = name;
            Namespace = @namespace;
            Resource = resource;
            Value = value;
        }
    }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk
{
    public partial class Environment : Pulumi.CustomResource
    {
        [Output("allSettings")]
        public Output<ImmutableArray<Outputs.EnvironmentAllSetting>> AllSettings { get; private set; } = null!;

        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("autoscalingGroups")]
        public Output<ImmutableArray<string>> AutoscalingGroups { get; private set; } = null!;

        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        [Output("cnamePrefix")]
        public Output<string> CnamePrefix { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("endpointUrl")]
        public Output<string> EndpointUrl { get; private set; } = null!;

        [Output("instances")]
        public Output<ImmutableArray<string>> Instances { get; private set; } = null!;

        [Output("launchConfigurations")]
        public Output<ImmutableArray<string>> LaunchConfigurations { get; private set; } = null!;

        [Output("loadBalancers")]
        public Output<ImmutableArray<string>> LoadBalancers { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platformArn")]
        public Output<string> PlatformArn { get; private set; } = null!;

        [Output("pollInterval")]
        public Output<string?> PollInterval { get; private set; } = null!;

        [Output("queues")]
        public Output<ImmutableArray<string>> Queues { get; private set; } = null!;

        [Output("settings")]
        public Output<ImmutableArray<Outputs.EnvironmentSetting>> Settings { get; private set; } = null!;

        [Output("solutionStackName")]
        public Output<string> SolutionStackName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("templateName")]
        public Output<string?> TemplateName { get; private set; } = null!;

        [Output("tier")]
        public Output<string?> Tier { get; private set; } = null!;

        [Output("triggers")]
        public Output<ImmutableArray<string>> Triggers { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        [Output("waitForReadyTimeout")]
        public Output<string?> WaitForReadyTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/environment:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/environment:Environment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, state, options);
        }
    }

    public sealed class EnvironmentArgs : Pulumi.ResourceArgs
    {
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        [Input("cnamePrefix")]
        public Input<string>? CnamePrefix { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformArn")]
        public Input<string>? PlatformArn { get; set; }

        [Input("pollInterval")]
        public Input<string>? PollInterval { get; set; }

        [Input("settings")]
        private InputList<Inputs.EnvironmentSettingArgs>? _settings;
        public InputList<Inputs.EnvironmentSettingArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.EnvironmentSettingArgs>());
            set => _settings = value;
        }

        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("tier")]
        public Input<string>? Tier { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("waitForReadyTimeout")]
        public Input<string>? WaitForReadyTimeout { get; set; }

        public EnvironmentArgs()
        {
        }
    }

    public sealed class EnvironmentState : Pulumi.ResourceArgs
    {
        [Input("allSettings")]
        private InputList<Inputs.EnvironmentAllSettingGetArgs>? _allSettings;
        public InputList<Inputs.EnvironmentAllSettingGetArgs> AllSettings
        {
            get => _allSettings ?? (_allSettings = new InputList<Inputs.EnvironmentAllSettingGetArgs>());
            set => _allSettings = value;
        }

        [Input("application")]
        public Input<string>? Application { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("autoscalingGroups")]
        private InputList<string>? _autoscalingGroups;
        public InputList<string> AutoscalingGroups
        {
            get => _autoscalingGroups ?? (_autoscalingGroups = new InputList<string>());
            set => _autoscalingGroups = value;
        }

        [Input("cname")]
        public Input<string>? Cname { get; set; }

        [Input("cnamePrefix")]
        public Input<string>? CnamePrefix { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("endpointUrl")]
        public Input<string>? EndpointUrl { get; set; }

        [Input("instances")]
        private InputList<string>? _instances;
        public InputList<string> Instances
        {
            get => _instances ?? (_instances = new InputList<string>());
            set => _instances = value;
        }

        [Input("launchConfigurations")]
        private InputList<string>? _launchConfigurations;
        public InputList<string> LaunchConfigurations
        {
            get => _launchConfigurations ?? (_launchConfigurations = new InputList<string>());
            set => _launchConfigurations = value;
        }

        [Input("loadBalancers")]
        private InputList<string>? _loadBalancers;
        public InputList<string> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<string>());
            set => _loadBalancers = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformArn")]
        public Input<string>? PlatformArn { get; set; }

        [Input("pollInterval")]
        public Input<string>? PollInterval { get; set; }

        [Input("queues")]
        private InputList<string>? _queues;
        public InputList<string> Queues
        {
            get => _queues ?? (_queues = new InputList<string>());
            set => _queues = value;
        }

        [Input("settings")]
        private InputList<Inputs.EnvironmentSettingGetArgs>? _settings;
        public InputList<Inputs.EnvironmentSettingGetArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.EnvironmentSettingGetArgs>());
            set => _settings = value;
        }

        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("tier")]
        public Input<string>? Tier { get; set; }

        [Input("triggers")]
        private InputList<string>? _triggers;
        public InputList<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<string>());
            set => _triggers = value;
        }

        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("waitForReadyTimeout")]
        public Input<string>? WaitForReadyTimeout { get; set; }

        public EnvironmentState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg
{
    public partial class ConfigurationAggregator : Pulumi.CustomResource
    {
        [Output("accountAggregationSource")]
        public Output<Outputs.ConfigurationAggregatorAccountAggregationSource?> AccountAggregationSource { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("organizationAggregationSource")]
        public Output<Outputs.ConfigurationAggregatorOrganizationAggregationSource?> OrganizationAggregationSource { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


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
        [Input("accountAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorAccountAggregationSourceArgs>? AccountAggregationSource { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs>? OrganizationAggregationSource { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ConfigurationAggregatorArgs()
        {
        }
    }

    public sealed class ConfigurationAggregatorState : Pulumi.ResourceArgs
    {
        [Input("accountAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorAccountAggregationSourceGetArgs>? AccountAggregationSource { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorOrganizationAggregationSourceGetArgs>? OrganizationAggregationSource { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ConfigurationAggregatorState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LicenseManager
{
    /// <summary>
    /// Provides a License Manager license configuration resource.
    /// 
    /// &gt; **Note:** Removing the `license_count` attribute is not supported by the License Manager API - recreate the resource instead.
    /// 
    /// ## Rules
    /// 
    /// License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:
    /// 
    /// * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
    /// * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
    /// * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
    /// * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
    /// * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
    /// * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
    /// * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`
    /// </summary>
    public partial class LicenseConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the license configuration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Number of licenses managed by the license configuration.
        /// </summary>
        [Output("licenseCount")]
        public Output<int?> LicenseCount { get; private set; } = null!;

        /// <summary>
        /// Sets the number of available licenses as a hard limit.
        /// </summary>
        [Output("licenseCountHardLimit")]
        public Output<bool?> LicenseCountHardLimit { get; private set; } = null!;

        /// <summary>
        /// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        /// </summary>
        [Output("licenseCountingType")]
        public Output<string> LicenseCountingType { get; private set; } = null!;

        /// <summary>
        /// Array of configured License Manager rules.
        /// </summary>
        [Output("licenseRules")]
        public Output<ImmutableArray<string>> LicenseRules { get; private set; } = null!;

        /// <summary>
        /// Name of the license configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseConfiguration(string name, LicenseConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, args ?? new LicenseConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseConfiguration(string name, Input<string> id, LicenseConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseConfiguration Get(string name, Input<string> id, LicenseConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseConfiguration(name, id, state, options);
        }
    }

    public sealed class LicenseConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the license configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of licenses managed by the license configuration.
        /// </summary>
        [Input("licenseCount")]
        public Input<int>? LicenseCount { get; set; }

        /// <summary>
        /// Sets the number of available licenses as a hard limit.
        /// </summary>
        [Input("licenseCountHardLimit")]
        public Input<bool>? LicenseCountHardLimit { get; set; }

        /// <summary>
        /// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        /// </summary>
        [Input("licenseCountingType", required: true)]
        public Input<string> LicenseCountingType { get; set; } = null!;

        [Input("licenseRules")]
        private InputList<string>? _licenseRules;

        /// <summary>
        /// Array of configured License Manager rules.
        /// </summary>
        public InputList<string> LicenseRules
        {
            get => _licenseRules ?? (_licenseRules = new InputList<string>());
            set => _licenseRules = value;
        }

        /// <summary>
        /// Name of the license configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public LicenseConfigurationArgs()
        {
        }
    }

    public sealed class LicenseConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the license configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of licenses managed by the license configuration.
        /// </summary>
        [Input("licenseCount")]
        public Input<int>? LicenseCount { get; set; }

        /// <summary>
        /// Sets the number of available licenses as a hard limit.
        /// </summary>
        [Input("licenseCountHardLimit")]
        public Input<bool>? LicenseCountHardLimit { get; set; }

        /// <summary>
        /// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        /// </summary>
        [Input("licenseCountingType")]
        public Input<string>? LicenseCountingType { get; set; }

        [Input("licenseRules")]
        private InputList<string>? _licenseRules;

        /// <summary>
        /// Array of configured License Manager rules.
        /// </summary>
        public InputList<string> LicenseRules
        {
            get => _licenseRules ?? (_licenseRules = new InputList<string>());
            set => _licenseRules = value;
        }

        /// <summary>
        /// Name of the license configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public LicenseConfigurationState()
        {
        }
    }
}

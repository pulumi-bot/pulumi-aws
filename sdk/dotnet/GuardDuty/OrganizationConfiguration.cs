// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty
{
    /// <summary>
    /// Manages the GuardDuty Organization Configuration in the current AWS Region. The AWS account utilizing this resource must have been assigned as a delegated Organization administrator account, e.g. via the `aws.guardduty.OrganizationAdminAccount` resource. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).
    /// 
    /// &gt; **NOTE:** This is an advanced resource. The provider will automatically assume management of the GuardDuty Organization Configuration without import and perform no actions on removal from the resource configuration.
    /// </summary>
    public partial class OrganizationConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
        /// </summary>
        [Output("autoEnable")]
        public Output<bool> AutoEnable { get; private set; } = null!;

        /// <summary>
        /// The detector ID of the GuardDuty account.
        /// </summary>
        [Output("detectorId")]
        public Output<string> DetectorId { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationConfiguration(string name, OrganizationConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:guardduty/organizationConfiguration:OrganizationConfiguration", name, args ?? new OrganizationConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationConfiguration(string name, Input<string> id, OrganizationConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/organizationConfiguration:OrganizationConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationConfiguration Get(string name, Input<string> id, OrganizationConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationConfiguration(name, id, state, options);
        }
    }

    public sealed class OrganizationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
        /// </summary>
        [Input("autoEnable", required: true)]
        public Input<bool> AutoEnable { get; set; } = null!;

        /// <summary>
        /// The detector ID of the GuardDuty account.
        /// </summary>
        [Input("detectorId", required: true)]
        public Input<string> DetectorId { get; set; } = null!;

        public OrganizationConfigurationArgs()
        {
        }
    }

    public sealed class OrganizationConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
        /// </summary>
        [Input("autoEnable")]
        public Input<bool>? AutoEnable { get; set; }

        /// <summary>
        /// The detector ID of the GuardDuty account.
        /// </summary>
        [Input("detectorId")]
        public Input<string>? DetectorId { get; set; }

        public OrganizationConfigurationState()
        {
        }
    }
}

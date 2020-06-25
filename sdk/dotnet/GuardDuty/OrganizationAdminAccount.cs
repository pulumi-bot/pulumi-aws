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
    /// Manages a GuardDuty Organization Admin Account. The AWS account utilizing this resource must be an Organizations master account. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).
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
    ///         var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new Aws.Organizations.OrganizationArgs
    ///         {
    ///             AwsServiceAccessPrincipals = 
    ///             {
    ///                 "guardduty.amazonaws.com",
    ///             },
    ///             FeatureSet = "ALL",
    ///         });
    ///         var exampleDetector = new Aws.GuardDuty.Detector("exampleDetector", new Aws.GuardDuty.DetectorArgs
    ///         {
    ///         });
    ///         var exampleOrganizationAdminAccount = new Aws.GuardDuty.OrganizationAdminAccount("exampleOrganizationAdminAccount", new Aws.GuardDuty.OrganizationAdminAccountArgs
    ///         {
    ///             AdminAccountId = "123456789012",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class OrganizationAdminAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// AWS account identifier to designate as a delegated administrator for GuardDuty.
        /// </summary>
        [Output("adminAccountId")]
        public Output<string> AdminAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationAdminAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationAdminAccount(string name, OrganizationAdminAccountArgs args, CustomResourceOptions? options = null)
            : base("aws:guardduty/organizationAdminAccount:OrganizationAdminAccount", name, args ?? new OrganizationAdminAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationAdminAccount(string name, Input<string> id, OrganizationAdminAccountState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/organizationAdminAccount:OrganizationAdminAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationAdminAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationAdminAccount Get(string name, Input<string> id, OrganizationAdminAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationAdminAccount(name, id, state, options);
        }
    }

    public sealed class OrganizationAdminAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS account identifier to designate as a delegated administrator for GuardDuty.
        /// </summary>
        [Input("adminAccountId", required: true)]
        public Input<string> AdminAccountId { get; set; } = null!;

        public OrganizationAdminAccountArgs()
        {
        }
    }

    public sealed class OrganizationAdminAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS account identifier to designate as a delegated administrator for GuardDuty.
        /// </summary>
        [Input("adminAccountId")]
        public Input<string>? AdminAccountId { get; set; }

        public OrganizationAdminAccountState()
        {
        }
    }
}

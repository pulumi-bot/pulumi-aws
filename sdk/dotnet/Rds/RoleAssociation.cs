// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public partial class RoleAssociation : Pulumi.CustomResource
    {
        [Output("dbInstanceIdentifier")]
        public Output<string> DbInstanceIdentifier { get; private set; } = null!;

        [Output("featureName")]
        public Output<string> FeatureName { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a RoleAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleAssociation(string name, RoleAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/roleAssociation:RoleAssociation", name, args ?? new RoleAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoleAssociation(string name, Input<string> id, RoleAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/roleAssociation:RoleAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoleAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleAssociation Get(string name, Input<string> id, RoleAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new RoleAssociation(name, id, state, options);
        }
    }

    public sealed class RoleAssociationArgs : Pulumi.ResourceArgs
    {
        [Input("dbInstanceIdentifier", required: true)]
        public Input<string> DbInstanceIdentifier { get; set; } = null!;

        [Input("featureName", required: true)]
        public Input<string> FeatureName { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public RoleAssociationArgs()
        {
        }
    }

    public sealed class RoleAssociationState : Pulumi.ResourceArgs
    {
        [Input("dbInstanceIdentifier")]
        public Input<string>? DbInstanceIdentifier { get; set; }

        [Input("featureName")]
        public Input<string>? FeatureName { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public RoleAssociationState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie
{
    public partial class S3BucketAssociation : Pulumi.CustomResource
    {
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        [Output("classificationType")]
        public Output<Outputs.S3BucketAssociationClassificationType> ClassificationType { get; private set; } = null!;

        [Output("memberAccountId")]
        public Output<string?> MemberAccountId { get; private set; } = null!;

        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;


        /// <summary>
        /// Create a S3BucketAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public S3BucketAssociation(string name, S3BucketAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:macie/s3BucketAssociation:S3BucketAssociation", name, args ?? new S3BucketAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private S3BucketAssociation(string name, Input<string> id, S3BucketAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie/s3BucketAssociation:S3BucketAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing S3BucketAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static S3BucketAssociation Get(string name, Input<string> id, S3BucketAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new S3BucketAssociation(name, id, state, options);
        }
    }

    public sealed class S3BucketAssociationArgs : Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("classificationType")]
        public Input<Inputs.S3BucketAssociationClassificationTypeArgs>? ClassificationType { get; set; }

        [Input("memberAccountId")]
        public Input<string>? MemberAccountId { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public S3BucketAssociationArgs()
        {
        }
    }

    public sealed class S3BucketAssociationState : Pulumi.ResourceArgs
    {
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("classificationType")]
        public Input<Inputs.S3BucketAssociationClassificationTypeGetArgs>? ClassificationType { get; set; }

        [Input("memberAccountId")]
        public Input<string>? MemberAccountId { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public S3BucketAssociationState()
        {
        }
    }
}

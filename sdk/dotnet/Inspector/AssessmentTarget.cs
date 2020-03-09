// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Inspector
{
    /// <summary>
    /// Provides a Inspector assessment target
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/inspector_assessment_target.html.markdown.
    /// </summary>
    public partial class AssessmentTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The target assessment ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the assessment target.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        /// </summary>
        [Output("resourceGroupArn")]
        public Output<string?> ResourceGroupArn { get; private set; } = null!;


        /// <summary>
        /// Create a AssessmentTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssessmentTarget(string name, AssessmentTargetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:inspector/assessmentTarget:AssessmentTarget", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AssessmentTarget(string name, Input<string> id, AssessmentTargetState? state = null, CustomResourceOptions? options = null)
            : base("aws:inspector/assessmentTarget:AssessmentTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AssessmentTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssessmentTarget Get(string name, Input<string> id, AssessmentTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new AssessmentTarget(name, id, state, options);
        }
    }

    public sealed class AssessmentTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the assessment target.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        /// </summary>
        [Input("resourceGroupArn")]
        public Input<string>? ResourceGroupArn { get; set; }

        public AssessmentTargetArgs()
        {
        }
    }

    public sealed class AssessmentTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target assessment ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the assessment target.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        /// </summary>
        [Input("resourceGroupArn")]
        public Input<string>? ResourceGroupArn { get; set; }

        public AssessmentTargetState()
        {
        }
    }
}

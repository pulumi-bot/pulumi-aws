// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling
{
    /// <summary>
    /// Provides an Auto Scaling Attachment resource.
    /// 
    /// &gt; **NOTE on Auto Scaling Groups and ASG Attachments:** This provider currently provides
    /// both a standalone `aws.autoscaling.Attachment` resource
    /// (describing an ASG attached to an ELB or ALB), and an `aws.autoscaling.Group`
    /// with `load_balancers` and `target_group_arns` defined in-line. These two methods are not
    /// mutually-exclusive. If `aws.autoscaling.Attachment` resources are used, either alone or with inline
    /// `load_balancers` or `target_group_arns`, the `aws.autoscaling.Group` resource must be configured
    /// to [ignore changes](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to the `load_balancers` and `target_group_arns` arguments.
    /// </summary>
    public partial class Attachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of an ALB Target Group.
        /// </summary>
        [Output("albTargetGroupArn")]
        public Output<string?> AlbTargetGroupArn { get; private set; } = null!;

        /// <summary>
        /// Name of ASG to associate with the ELB.
        /// </summary>
        [Output("autoscalingGroupName")]
        public Output<string> AutoscalingGroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the ELB.
        /// </summary>
        [Output("elb")]
        public Output<string?> Elb { get; private set; } = null!;


        /// <summary>
        /// Create a Attachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attachment(string name, AttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:autoscaling/attachment:Attachment", name, args ?? new AttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Attachment(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:autoscaling/attachment:Attachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attachment Get(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Attachment(name, id, state, options);
        }
    }

    public sealed class AttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of an ALB Target Group.
        /// </summary>
        [Input("albTargetGroupArn")]
        public Input<string>? AlbTargetGroupArn { get; set; }

        /// <summary>
        /// Name of ASG to associate with the ELB.
        /// </summary>
        [Input("autoscalingGroupName", required: true)]
        public Input<string> AutoscalingGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the ELB.
        /// </summary>
        [Input("elb")]
        public Input<string>? Elb { get; set; }

        public AttachmentArgs()
        {
        }
    }

    public sealed class AttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of an ALB Target Group.
        /// </summary>
        [Input("albTargetGroupArn")]
        public Input<string>? AlbTargetGroupArn { get; set; }

        /// <summary>
        /// Name of ASG to associate with the ELB.
        /// </summary>
        [Input("autoscalingGroupName")]
        public Input<string>? AutoscalingGroupName { get; set; }

        /// <summary>
        /// The name of the ELB.
        /// </summary>
        [Input("elb")]
        public Input<string>? Elb { get; set; }

        public AttachmentState()
        {
        }
    }
}

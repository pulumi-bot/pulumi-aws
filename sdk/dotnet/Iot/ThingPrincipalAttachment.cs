// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot
{
    /// <summary>
    /// Attaches Principal to AWS IoT Thing.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing_principal_attachment.html.markdown.
    /// </summary>
    public partial class ThingPrincipalAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The name of the thing.
        /// </summary>
        [Output("thing")]
        public Output<string> Thing { get; private set; } = null!;


        /// <summary>
        /// Create a ThingPrincipalAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThingPrincipalAttachment(string name, ThingPrincipalAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ThingPrincipalAttachment(string name, Input<string> id, ThingPrincipalAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ThingPrincipalAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThingPrincipalAttachment Get(string name, Input<string> id, ThingPrincipalAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ThingPrincipalAttachment(name, id, state, options);
        }
    }

    public sealed class ThingPrincipalAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The name of the thing.
        /// </summary>
        [Input("thing", required: true)]
        public Input<string> Thing { get; set; } = null!;

        public ThingPrincipalAttachmentArgs()
        {
        }
    }

    public sealed class ThingPrincipalAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The name of the thing.
        /// </summary>
        [Input("thing")]
        public Input<string>? Thing { get; set; }

        public ThingPrincipalAttachmentState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppAutoScaling
{
    public partial class ScheduledAction : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        [Output("scalableDimension")]
        public Output<string?> ScalableDimension { get; private set; } = null!;

        [Output("scalableTargetAction")]
        public Output<Outputs.ScheduledActionScalableTargetAction?> ScalableTargetAction { get; private set; } = null!;

        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        [Output("serviceNamespace")]
        public Output<string> ServiceNamespace { get; private set; } = null!;

        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduledAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduledAction(string name, ScheduledActionArgs args, CustomResourceOptions? options = null)
            : base("aws:appautoscaling/scheduledAction:ScheduledAction", name, args ?? new ScheduledActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduledAction(string name, Input<string> id, ScheduledActionState? state = null, CustomResourceOptions? options = null)
            : base("aws:appautoscaling/scheduledAction:ScheduledAction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScheduledAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduledAction Get(string name, Input<string> id, ScheduledActionState? state = null, CustomResourceOptions? options = null)
        {
            return new ScheduledAction(name, id, state, options);
        }
    }

    public sealed class ScheduledActionArgs : Pulumi.ResourceArgs
    {
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        [Input("scalableDimension")]
        public Input<string>? ScalableDimension { get; set; }

        [Input("scalableTargetAction")]
        public Input<Inputs.ScheduledActionScalableTargetActionArgs>? ScalableTargetAction { get; set; }

        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        [Input("serviceNamespace", required: true)]
        public Input<string> ServiceNamespace { get; set; } = null!;

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public ScheduledActionArgs()
        {
        }
    }

    public sealed class ScheduledActionState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        [Input("scalableDimension")]
        public Input<string>? ScalableDimension { get; set; }

        [Input("scalableTargetAction")]
        public Input<Inputs.ScheduledActionScalableTargetActionGetArgs>? ScalableTargetAction { get; set; }

        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        [Input("serviceNamespace")]
        public Input<string>? ServiceNamespace { get; set; }

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public ScheduledActionState()
        {
        }
    }
}

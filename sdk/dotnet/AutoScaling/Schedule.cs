// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling
{
    public partial class Schedule : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("autoscalingGroupName")]
        public Output<string> AutoscalingGroupName { get; private set; } = null!;

        [Output("desiredCapacity")]
        public Output<int> DesiredCapacity { get; private set; } = null!;

        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        [Output("maxSize")]
        public Output<int> MaxSize { get; private set; } = null!;

        [Output("minSize")]
        public Output<int> MinSize { get; private set; } = null!;

        [Output("recurrence")]
        public Output<string> Recurrence { get; private set; } = null!;

        [Output("scheduledActionName")]
        public Output<string> ScheduledActionName { get; private set; } = null!;

        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a Schedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schedule(string name, ScheduleArgs args, CustomResourceOptions? options = null)
            : base("aws:autoscaling/schedule:Schedule", name, args ?? new ScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schedule(string name, Input<string> id, ScheduleState? state = null, CustomResourceOptions? options = null)
            : base("aws:autoscaling/schedule:Schedule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Schedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schedule Get(string name, Input<string> id, ScheduleState? state = null, CustomResourceOptions? options = null)
        {
            return new Schedule(name, id, state, options);
        }
    }

    public sealed class ScheduleArgs : Pulumi.ResourceArgs
    {
        [Input("autoscalingGroupName", required: true)]
        public Input<string> AutoscalingGroupName { get; set; } = null!;

        [Input("desiredCapacity")]
        public Input<int>? DesiredCapacity { get; set; }

        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        [Input("recurrence")]
        public Input<string>? Recurrence { get; set; }

        [Input("scheduledActionName", required: true)]
        public Input<string> ScheduledActionName { get; set; } = null!;

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public ScheduleArgs()
        {
        }
    }

    public sealed class ScheduleState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("autoscalingGroupName")]
        public Input<string>? AutoscalingGroupName { get; set; }

        [Input("desiredCapacity")]
        public Input<int>? DesiredCapacity { get; set; }

        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        [Input("recurrence")]
        public Input<string>? Recurrence { get; set; }

        [Input("scheduledActionName")]
        public Input<string>? ScheduledActionName { get; set; }

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public ScheduleState()
        {
        }
    }
}

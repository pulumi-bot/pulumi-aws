// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    public partial class MaintenanceWindowTask : Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("maxConcurrency")]
        public Output<string> MaxConcurrency { get; private set; } = null!;

        [Output("maxErrors")]
        public Output<string> MaxErrors { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        [Output("serviceRoleArn")]
        public Output<string> ServiceRoleArn { get; private set; } = null!;

        [Output("targets")]
        public Output<ImmutableArray<Outputs.MaintenanceWindowTaskTarget>> Targets { get; private set; } = null!;

        [Output("taskArn")]
        public Output<string> TaskArn { get; private set; } = null!;

        [Output("taskInvocationParameters")]
        public Output<Outputs.MaintenanceWindowTaskTaskInvocationParameters?> TaskInvocationParameters { get; private set; } = null!;

        [Output("taskType")]
        public Output<string> TaskType { get; private set; } = null!;

        [Output("windowId")]
        public Output<string> WindowId { get; private set; } = null!;


        /// <summary>
        /// Create a MaintenanceWindowTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MaintenanceWindowTask(string name, MaintenanceWindowTaskArgs args, CustomResourceOptions? options = null)
            : base("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, args ?? new MaintenanceWindowTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MaintenanceWindowTask(string name, Input<string> id, MaintenanceWindowTaskState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MaintenanceWindowTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MaintenanceWindowTask Get(string name, Input<string> id, MaintenanceWindowTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new MaintenanceWindowTask(name, id, state, options);
        }
    }

    public sealed class MaintenanceWindowTaskArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("maxConcurrency", required: true)]
        public Input<string> MaxConcurrency { get; set; } = null!;

        [Input("maxErrors", required: true)]
        public Input<string> MaxErrors { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("serviceRoleArn", required: true)]
        public Input<string> ServiceRoleArn { get; set; } = null!;

        [Input("targets", required: true)]
        private InputList<Inputs.MaintenanceWindowTaskTargetArgs>? _targets;
        public InputList<Inputs.MaintenanceWindowTaskTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MaintenanceWindowTaskTargetArgs>());
            set => _targets = value;
        }

        [Input("taskArn", required: true)]
        public Input<string> TaskArn { get; set; } = null!;

        [Input("taskInvocationParameters")]
        public Input<Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs>? TaskInvocationParameters { get; set; }

        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        [Input("windowId", required: true)]
        public Input<string> WindowId { get; set; } = null!;

        public MaintenanceWindowTaskArgs()
        {
        }
    }

    public sealed class MaintenanceWindowTaskState : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("maxConcurrency")]
        public Input<string>? MaxConcurrency { get; set; }

        [Input("maxErrors")]
        public Input<string>? MaxErrors { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("serviceRoleArn")]
        public Input<string>? ServiceRoleArn { get; set; }

        [Input("targets")]
        private InputList<Inputs.MaintenanceWindowTaskTargetGetArgs>? _targets;
        public InputList<Inputs.MaintenanceWindowTaskTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MaintenanceWindowTaskTargetGetArgs>());
            set => _targets = value;
        }

        [Input("taskArn")]
        public Input<string>? TaskArn { get; set; }

        [Input("taskInvocationParameters")]
        public Input<Inputs.MaintenanceWindowTaskTaskInvocationParametersGetArgs>? TaskInvocationParameters { get; set; }

        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        [Input("windowId")]
        public Input<string>? WindowId { get; set; }

        public MaintenanceWindowTaskState()
        {
        }
    }
}

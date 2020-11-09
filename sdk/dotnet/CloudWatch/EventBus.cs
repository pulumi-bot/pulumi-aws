// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides an EventBridge event bus resource.
    /// 
    /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
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
    ///         var messenger = new Aws.CloudWatch.EventBus("messenger", new Aws.CloudWatch.EventBusArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EventBridge event buses can be imported using the `name`, e.g. console
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudwatch/eventBus:EventBus messenger chat-messages
    /// ```
    /// </summary>
    public partial class EventBus : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the event bus.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the new event bus.
        /// The names of custom event buses can't contain the / character.
        /// Please note that a partner event bus is not supported at the moment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EventBus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventBus(string name, EventBusArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventBus:EventBus", name, args ?? new EventBusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventBus(string name, Input<string> id, EventBusState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventBus:EventBus", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventBus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventBus Get(string name, Input<string> id, EventBusState? state = null, CustomResourceOptions? options = null)
        {
            return new EventBus(name, id, state, options);
        }
    }

    public sealed class EventBusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the new event bus.
        /// The names of custom event buses can't contain the / character.
        /// Please note that a partner event bus is not supported at the moment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EventBusArgs()
        {
        }
    }

    public sealed class EventBusState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the event bus.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the new event bus.
        /// The names of custom event buses can't contain the / character.
        /// Please note that a partner event bus is not supported at the moment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EventBusState()
        {
        }
    }
}

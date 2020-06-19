// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Attach an Elastic network interface (ENI) resource with EC2 instance.
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
    ///         var test = new Aws.Ec2.NetworkInterfaceAttachment("test", new Aws.Ec2.NetworkInterfaceAttachmentArgs
    ///         {
    ///             DeviceIndex = 0,
    ///             InstanceId = aws_instance.Test.Id,
    ///             NetworkInterfaceId = aws_network_interface.Test.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NetworkInterfaceAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ENI Attachment ID.
        /// </summary>
        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// Network interface index (int).
        /// </summary>
        [Output("deviceIndex")]
        public Output<int> DeviceIndex { get; private set; } = null!;

        /// <summary>
        /// Instance ID to attach.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// ENI ID to attach.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The status of the Network Interface Attachment.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceAttachment(string name, NetworkInterfaceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment", name, args ?? new NetworkInterfaceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceAttachment(string name, Input<string> id, NetworkInterfaceAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceAttachment Get(string name, Input<string> id, NetworkInterfaceAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceAttachment(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network interface index (int).
        /// </summary>
        [Input("deviceIndex", required: true)]
        public Input<int> DeviceIndex { get; set; } = null!;

        /// <summary>
        /// Instance ID to attach.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// ENI ID to attach.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        public NetworkInterfaceAttachmentArgs()
        {
        }
    }

    public sealed class NetworkInterfaceAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ENI Attachment ID.
        /// </summary>
        [Input("attachmentId")]
        public Input<string>? AttachmentId { get; set; }

        /// <summary>
        /// Network interface index (int).
        /// </summary>
        [Input("deviceIndex")]
        public Input<int>? DeviceIndex { get; set; }

        /// <summary>
        /// Instance ID to attach.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// ENI ID to attach.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The status of the Network Interface Attachment.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public NetworkInterfaceAttachmentState()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.StorageGateway
{
    /// <summary>
    /// Manages an AWS Storage Gateway upload buffer.
    /// 
    /// &gt; **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this resource does not perform any Storage Gateway actions.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.StorageGateway.UploadBuffer("example", new Aws.StorageGateway.UploadBufferArgs
    ///         {
    ///             DiskId = data.Aws_storagegateway_local_disk.Example.Id,
    ///             GatewayArn = aws_storagegateway_gateway.Example.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class UploadBuffer : Pulumi.CustomResource
    {
        /// <summary>
        /// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        /// </summary>
        [Output("diskId")]
        public Output<string> DiskId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Output("gatewayArn")]
        public Output<string> GatewayArn { get; private set; } = null!;


        /// <summary>
        /// Create a UploadBuffer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UploadBuffer(string name, UploadBufferArgs args, CustomResourceOptions? options = null)
            : base("aws:storagegateway/uploadBuffer:UploadBuffer", name, args ?? new UploadBufferArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UploadBuffer(string name, Input<string> id, UploadBufferState? state = null, CustomResourceOptions? options = null)
            : base("aws:storagegateway/uploadBuffer:UploadBuffer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UploadBuffer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UploadBuffer Get(string name, Input<string> id, UploadBufferState? state = null, CustomResourceOptions? options = null)
        {
            return new UploadBuffer(name, id, state, options);
        }
    }

    public sealed class UploadBufferArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Input("gatewayArn", required: true)]
        public Input<string> GatewayArn { get; set; } = null!;

        public UploadBufferArgs()
        {
        }
    }

    public sealed class UploadBufferState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Input("gatewayArn")]
        public Input<string>? GatewayArn { get; set; }

        public UploadBufferState()
        {
        }
    }
}

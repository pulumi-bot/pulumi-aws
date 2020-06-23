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
    /// Manages an AWS Storage Gateway file, tape, or volume gateway in the provider region.
    /// 
    /// &gt; NOTE: The Storage Gateway API requires the gateway to be connected to properly return information after activation. If you are receiving `The specified gateway is not connected` errors during resource creation (gateway activation), ensure your gateway instance meets the [Storage Gateway requirements](https://docs.aws.amazon.com/storagegateway/latest/userguide/Requirements.html).
    /// 
    /// ## Example Usage
    /// ### File Gateway
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.StorageGateway.Gateway("example", new Aws.StorageGateway.GatewayArgs
    ///         {
    ///             GatewayIpAddress = "1.2.3.4",
    ///             GatewayName = "example",
    ///             GatewayTimezone = "GMT",
    ///             GatewayType = "FILE_S3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Volume Gateway (Cached)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.StorageGateway.Gateway("example", new Aws.StorageGateway.GatewayArgs
    ///         {
    ///             GatewayIpAddress = "1.2.3.4",
    ///             GatewayName = "example",
    ///             GatewayTimezone = "GMT",
    ///             GatewayType = "CACHED",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Volume Gateway (Stored)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.StorageGateway.Gateway("example", new Aws.StorageGateway.GatewayArgs
    ///         {
    ///             GatewayIpAddress = "1.2.3.4",
    ///             GatewayName = "example",
    ///             GatewayTimezone = "GMT",
    ///             GatewayType = "STORED",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Gateway : Pulumi.CustomResource
    {
        /// <summary>
        /// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Output("activationKey")]
        public Output<string> ActivationKey { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
        /// </summary>
        [Output("cloudwatchLogGroupArn")]
        public Output<string?> CloudwatchLogGroupArn { get; private set; } = null!;

        /// <summary>
        /// Identifier of the gateway.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Output("gatewayIpAddress")]
        public Output<string> GatewayIpAddress { get; private set; } = null!;

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Output("gatewayName")]
        public Output<string> GatewayName { get; private set; } = null!;

        /// <summary>
        /// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
        /// </summary>
        [Output("gatewayTimezone")]
        public Output<string> GatewayTimezone { get; private set; } = null!;

        /// <summary>
        /// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
        /// </summary>
        [Output("gatewayType")]
        public Output<string?> GatewayType { get; private set; } = null!;

        /// <summary>
        /// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
        /// </summary>
        [Output("gatewayVpcEndpoint")]
        public Output<string?> GatewayVpcEndpoint { get; private set; } = null!;

        [Output("mediumChangerType")]
        public Output<string?> MediumChangerType { get; private set; } = null!;

        /// <summary>
        /// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
        /// </summary>
        [Output("smbActiveDirectorySettings")]
        public Output<Outputs.GatewaySmbActiveDirectorySettings?> SmbActiveDirectorySettings { get; private set; } = null!;

        /// <summary>
        /// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
        /// </summary>
        [Output("smbGuestPassword")]
        public Output<string?> SmbGuestPassword { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
        /// </summary>
        [Output("tapeDriveType")]
        public Output<string?> TapeDriveType { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("aws:storagegateway/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("aws:storagegateway/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("activationKey")]
        public Input<string>? ActivationKey { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
        /// </summary>
        [Input("cloudwatchLogGroupArn")]
        public Input<string>? CloudwatchLogGroupArn { get; set; }

        /// <summary>
        /// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("gatewayIpAddress")]
        public Input<string>? GatewayIpAddress { get; set; }

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
        /// </summary>
        [Input("gatewayTimezone", required: true)]
        public Input<string> GatewayTimezone { get; set; } = null!;

        /// <summary>
        /// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
        /// </summary>
        [Input("gatewayType")]
        public Input<string>? GatewayType { get; set; }

        /// <summary>
        /// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
        /// </summary>
        [Input("gatewayVpcEndpoint")]
        public Input<string>? GatewayVpcEndpoint { get; set; }

        [Input("mediumChangerType")]
        public Input<string>? MediumChangerType { get; set; }

        /// <summary>
        /// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
        /// </summary>
        [Input("smbActiveDirectorySettings")]
        public Input<Inputs.GatewaySmbActiveDirectorySettingsArgs>? SmbActiveDirectorySettings { get; set; }

        /// <summary>
        /// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
        /// </summary>
        [Input("smbGuestPassword")]
        public Input<string>? SmbGuestPassword { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
        /// </summary>
        [Input("tapeDriveType")]
        public Input<string>? TapeDriveType { get; set; }

        public GatewayArgs()
        {
        }
    }

    public sealed class GatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("activationKey")]
        public Input<string>? ActivationKey { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
        /// </summary>
        [Input("cloudwatchLogGroupArn")]
        public Input<string>? CloudwatchLogGroupArn { get; set; }

        /// <summary>
        /// Identifier of the gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("gatewayIpAddress")]
        public Input<string>? GatewayIpAddress { get; set; }

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        /// <summary>
        /// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
        /// </summary>
        [Input("gatewayTimezone")]
        public Input<string>? GatewayTimezone { get; set; }

        /// <summary>
        /// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_S3`, `STORED`, `VTL`.
        /// </summary>
        [Input("gatewayType")]
        public Input<string>? GatewayType { get; set; }

        /// <summary>
        /// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running Pulumi. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
        /// </summary>
        [Input("gatewayVpcEndpoint")]
        public Input<string>? GatewayVpcEndpoint { get; set; }

        [Input("mediumChangerType")]
        public Input<string>? MediumChangerType { get; set; }

        /// <summary>
        /// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
        /// </summary>
        [Input("smbActiveDirectorySettings")]
        public Input<Inputs.GatewaySmbActiveDirectorySettingsGetArgs>? SmbActiveDirectorySettings { get; set; }

        /// <summary>
        /// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` gateway type. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
        /// </summary>
        [Input("smbGuestPassword")]
        public Input<string>? SmbGuestPassword { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
        /// </summary>
        [Input("tapeDriveType")]
        public Input<string>? TapeDriveType { get; set; }

        public GatewayState()
        {
        }
    }
}

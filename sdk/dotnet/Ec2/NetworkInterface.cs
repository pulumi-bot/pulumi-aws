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
    /// Provides an Elastic network interface (ENI) resource.
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
    ///         var test = new Aws.Ec2.NetworkInterface("test", new Aws.Ec2.NetworkInterfaceArgs
    ///         {
    ///             Attachments = 
    ///             {
    ///                 new Aws.Ec2.Inputs.NetworkInterfaceAttachmentArgs
    ///                 {
    ///                     DeviceIndex = 1,
    ///                     Instance = aws_instance.Test.Id,
    ///                 },
    ///             },
    ///             PrivateIps = 
    ///             {
    ///                 "10.0.0.50",
    ///             },
    ///             SecurityGroups = 
    ///             {
    ///                 aws_security_group.Web.Id,
    ///             },
    ///             SubnetId = aws_subnet.Public_a.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NetworkInterface : Pulumi.CustomResource
    {
        /// <summary>
        /// Block to define the attachment of the ENI. Documented below.
        /// </summary>
        [Output("attachments")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceAttachment>> Attachments { get; private set; } = null!;

        /// <summary>
        /// A description for the network interface.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        [Output("outpostArn")]
        public Output<string> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// The private DNS name of the network interface (IPv4).
        /// </summary>
        [Output("privateDnsName")]
        public Output<string> PrivateDnsName { get; private set; } = null!;

        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// List of private IPs to assign to the ENI.
        /// </summary>
        [Output("privateIps")]
        public Output<ImmutableArray<string>> PrivateIps { get; private set; } = null!;

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default. 
        /// </summary>
        [Output("privateIpsCount")]
        public Output<int> PrivateIpsCount { get; private set; } = null!;

        /// <summary>
        /// List of security group IDs to assign to the ENI.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Whether to enable source destination checking for the ENI. Default true.
        /// </summary>
        [Output("sourceDestCheck")]
        public Output<bool?> SourceDestCheck { get; private set; } = null!;

        /// <summary>
        /// Subnet ID to create the ENI in.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterface:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterface:NetworkInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputList<Inputs.NetworkInterfaceAttachmentArgs>? _attachments;

        /// <summary>
        /// Block to define the attachment of the ENI. Documented below.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceAttachmentArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.NetworkInterfaceAttachmentArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// A description for the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of private IPs to assign to the ENI.
        /// </summary>
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default. 
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of security group IDs to assign to the ENI.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Whether to enable source destination checking for the ENI. Default true.
        /// </summary>
        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        /// <summary>
        /// Subnet ID to create the ENI in.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public NetworkInterfaceArgs()
        {
        }
    }

    public sealed class NetworkInterfaceState : Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputList<Inputs.NetworkInterfaceAttachmentGetArgs>? _attachments;

        /// <summary>
        /// Block to define the attachment of the ENI. Documented below.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceAttachmentGetArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.NetworkInterfaceAttachmentGetArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// A description for the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The private DNS name of the network interface (IPv4).
        /// </summary>
        [Input("privateDnsName")]
        public Input<string>? PrivateDnsName { get; set; }

        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of private IPs to assign to the ENI.
        /// </summary>
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default. 
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of security group IDs to assign to the ENI.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Whether to enable source destination checking for the ENI. Default true.
        /// </summary>
        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        /// <summary>
        /// Subnet ID to create the ENI in.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public NetworkInterfaceState()
        {
        }
    }
}

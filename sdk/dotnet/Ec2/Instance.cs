// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public partial class Instance : Pulumi.CustomResource
    {
        [Output("ami")]
        public Output<string> Ami { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("associatePublicIpAddress")]
        public Output<bool> AssociatePublicIpAddress { get; private set; } = null!;

        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        [Output("cpuCoreCount")]
        public Output<int> CpuCoreCount { get; private set; } = null!;

        [Output("cpuThreadsPerCore")]
        public Output<int> CpuThreadsPerCore { get; private set; } = null!;

        [Output("creditSpecification")]
        public Output<Outputs.InstanceCreditSpecification?> CreditSpecification { get; private set; } = null!;

        [Output("disableApiTermination")]
        public Output<bool?> DisableApiTermination { get; private set; } = null!;

        [Output("ebsBlockDevices")]
        public Output<ImmutableArray<Outputs.InstanceEbsBlockDevice>> EbsBlockDevices { get; private set; } = null!;

        [Output("ebsOptimized")]
        public Output<bool?> EbsOptimized { get; private set; } = null!;

        [Output("ephemeralBlockDevices")]
        public Output<ImmutableArray<Outputs.InstanceEphemeralBlockDevice>> EphemeralBlockDevices { get; private set; } = null!;

        [Output("getPasswordData")]
        public Output<bool?> GetPasswordData { get; private set; } = null!;

        [Output("hibernation")]
        public Output<bool?> Hibernation { get; private set; } = null!;

        [Output("hostId")]
        public Output<string> HostId { get; private set; } = null!;

        [Output("iamInstanceProfile")]
        public Output<string?> IamInstanceProfile { get; private set; } = null!;

        [Output("instanceInitiatedShutdownBehavior")]
        public Output<string?> InstanceInitiatedShutdownBehavior { get; private set; } = null!;

        [Output("instanceState")]
        public Output<string> State { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        [Output("ipv6AddressCount")]
        public Output<int> Ipv6AddressCount { get; private set; } = null!;

        [Output("ipv6Addresses")]
        public Output<ImmutableArray<string>> Ipv6Addresses { get; private set; } = null!;

        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        [Output("metadataOptions")]
        public Output<Outputs.InstanceMetadataOptions> MetadataOptions { get; private set; } = null!;

        [Output("monitoring")]
        public Output<bool?> Monitoring { get; private set; } = null!;

        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.InstanceNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        [Output("outpostArn")]
        public Output<string> OutpostArn { get; private set; } = null!;

        [Output("passwordData")]
        public Output<string> PasswordData { get; private set; } = null!;

        [Output("placementGroup")]
        public Output<string> PlacementGroup { get; private set; } = null!;

        [Output("primaryNetworkInterfaceId")]
        public Output<string> PrimaryNetworkInterfaceId { get; private set; } = null!;

        [Output("privateDns")]
        public Output<string> PrivateDns { get; private set; } = null!;

        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        [Output("publicDns")]
        public Output<string> PublicDns { get; private set; } = null!;

        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        [Output("rootBlockDevice")]
        public Output<Outputs.InstanceRootBlockDevice> RootBlockDevice { get; private set; } = null!;

        [Output("secondaryPrivateIps")]
        public Output<ImmutableArray<string>> SecondaryPrivateIps { get; private set; } = null!;

        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        [Output("sourceDestCheck")]
        public Output<bool?> SourceDestCheck { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tenancy")]
        public Output<string> Tenancy { get; private set; } = null!;

        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        [Output("userDataBase64")]
        public Output<string?> UserDataBase64 { get; private set; } = null!;

        [Output("volumeTags")]
        public Output<ImmutableDictionary<string, string>> VolumeTags { get; private set; } = null!;

        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        [Input("ami", required: true)]
        public Input<string> Ami { get; set; } = null!;

        [Input("associatePublicIpAddress")]
        public Input<bool>? AssociatePublicIpAddress { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cpuCoreCount")]
        public Input<int>? CpuCoreCount { get; set; }

        [Input("cpuThreadsPerCore")]
        public Input<int>? CpuThreadsPerCore { get; set; }

        [Input("creditSpecification")]
        public Input<Inputs.InstanceCreditSpecificationArgs>? CreditSpecification { get; set; }

        [Input("disableApiTermination")]
        public Input<bool>? DisableApiTermination { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.InstanceEbsBlockDeviceArgs>? _ebsBlockDevices;
        public InputList<Inputs.InstanceEbsBlockDeviceArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.InstanceEbsBlockDeviceArgs>());
            set => _ebsBlockDevices = value;
        }

        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.InstanceEphemeralBlockDeviceArgs>? _ephemeralBlockDevices;
        public InputList<Inputs.InstanceEphemeralBlockDeviceArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.InstanceEphemeralBlockDeviceArgs>());
            set => _ephemeralBlockDevices = value;
        }

        [Input("getPasswordData")]
        public Input<bool>? GetPasswordData { get; set; }

        [Input("hibernation")]
        public Input<bool>? Hibernation { get; set; }

        [Input("hostId")]
        public Input<string>? HostId { get; set; }

        [Input("iamInstanceProfile")]
        public Input<string>? IamInstanceProfile { get; set; }

        [Input("instanceInitiatedShutdownBehavior")]
        public Input<string>? InstanceInitiatedShutdownBehavior { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("metadataOptions")]
        public Input<Inputs.InstanceMetadataOptionsArgs>? MetadataOptions { get; set; }

        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceNetworkInterfaceArgs>? _networkInterfaces;
        public InputList<Inputs.InstanceNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("placementGroup")]
        public Input<string>? PlacementGroup { get; set; }

        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("rootBlockDevice")]
        public Input<Inputs.InstanceRootBlockDeviceArgs>? RootBlockDevice { get; set; }

        [Input("secondaryPrivateIps")]
        private InputList<string>? _secondaryPrivateIps;
        public InputList<string> SecondaryPrivateIps
        {
            get => _secondaryPrivateIps ?? (_secondaryPrivateIps = new InputList<string>());
            set => _secondaryPrivateIps = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        [Obsolete(@"Use of `securityGroups` is discouraged as it does not allow for changes and will force your instance to be replaced if changes are made. To avoid this, use `vpcSecurityGroupIds` which allows for updates.")]
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tenancy")]
        public Input<string>? Tenancy { get; set; }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("userDataBase64")]
        public Input<string>? UserDataBase64 { get; set; }

        [Input("volumeTags")]
        private InputMap<string>? _volumeTags;
        public InputMap<string> VolumeTags
        {
            get => _volumeTags ?? (_volumeTags = new InputMap<string>());
            set => _volumeTags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        [Input("ami")]
        public Input<string>? Ami { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("associatePublicIpAddress")]
        public Input<bool>? AssociatePublicIpAddress { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cpuCoreCount")]
        public Input<int>? CpuCoreCount { get; set; }

        [Input("cpuThreadsPerCore")]
        public Input<int>? CpuThreadsPerCore { get; set; }

        [Input("creditSpecification")]
        public Input<Inputs.InstanceCreditSpecificationGetArgs>? CreditSpecification { get; set; }

        [Input("disableApiTermination")]
        public Input<bool>? DisableApiTermination { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.InstanceEbsBlockDeviceGetArgs>? _ebsBlockDevices;
        public InputList<Inputs.InstanceEbsBlockDeviceGetArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.InstanceEbsBlockDeviceGetArgs>());
            set => _ebsBlockDevices = value;
        }

        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.InstanceEphemeralBlockDeviceGetArgs>? _ephemeralBlockDevices;
        public InputList<Inputs.InstanceEphemeralBlockDeviceGetArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.InstanceEphemeralBlockDeviceGetArgs>());
            set => _ephemeralBlockDevices = value;
        }

        [Input("getPasswordData")]
        public Input<bool>? GetPasswordData { get; set; }

        [Input("hibernation")]
        public Input<bool>? Hibernation { get; set; }

        [Input("hostId")]
        public Input<string>? HostId { get; set; }

        [Input("iamInstanceProfile")]
        public Input<string>? IamInstanceProfile { get; set; }

        [Input("instanceInitiatedShutdownBehavior")]
        public Input<string>? InstanceInitiatedShutdownBehavior { get; set; }

        [Input("instanceState")]
        public Input<string>? State { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("metadataOptions")]
        public Input<Inputs.InstanceMetadataOptionsGetArgs>? MetadataOptions { get; set; }

        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceNetworkInterfaceGetArgs>? _networkInterfaces;
        public InputList<Inputs.InstanceNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        [Input("passwordData")]
        public Input<string>? PasswordData { get; set; }

        [Input("placementGroup")]
        public Input<string>? PlacementGroup { get; set; }

        [Input("primaryNetworkInterfaceId")]
        public Input<string>? PrimaryNetworkInterfaceId { get; set; }

        [Input("privateDns")]
        public Input<string>? PrivateDns { get; set; }

        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("publicDns")]
        public Input<string>? PublicDns { get; set; }

        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        [Input("rootBlockDevice")]
        public Input<Inputs.InstanceRootBlockDeviceGetArgs>? RootBlockDevice { get; set; }

        [Input("secondaryPrivateIps")]
        private InputList<string>? _secondaryPrivateIps;
        public InputList<string> SecondaryPrivateIps
        {
            get => _secondaryPrivateIps ?? (_secondaryPrivateIps = new InputList<string>());
            set => _secondaryPrivateIps = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        [Obsolete(@"Use of `securityGroups` is discouraged as it does not allow for changes and will force your instance to be replaced if changes are made. To avoid this, use `vpcSecurityGroupIds` which allows for updates.")]
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tenancy")]
        public Input<string>? Tenancy { get; set; }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("userDataBase64")]
        public Input<string>? UserDataBase64 { get; set; }

        [Input("volumeTags")]
        private InputMap<string>? _volumeTags;
        public InputMap<string> VolumeTags
        {
            get => _volumeTags ?? (_volumeTags = new InputMap<string>());
            set => _volumeTags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public InstanceState()
        {
        }
    }
}

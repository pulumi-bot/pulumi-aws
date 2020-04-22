// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetInstance
    {
        /// <summary>
        /// Use this data source to get the ID of an Amazon EC2 Instance for use in other
        /// resources.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("aws:ec2/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithVersion());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstanceFilterArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are
        /// several valid keys, for a full reference, check out
        /// [describe-instances in the AWS CLI reference][1].
        /// </summary>
        public List<Inputs.GetInstanceFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
        /// </summary>
        [Input("getPasswordData")]
        public bool? GetPasswordData { get; set; }

        /// <summary>
        /// Retrieve Base64 encoded User Data contents into the `user_data_base64` attribute. A SHA-1 hash of the User Data contents will always be present in the `user_data` attribute. Defaults to `false`.
        /// </summary>
        [Input("getUserData")]
        public bool? GetUserData { get; set; }

        /// <summary>
        /// Specify the exact Instance ID with which to populate the data source.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("instanceTags")]
        private Dictionary<string, object>? _instanceTags;

        /// <summary>
        /// A mapping of tags, each pair of which must
        /// exactly match a pair on the desired Instance.
        /// </summary>
        public Dictionary<string, object> InstanceTags
        {
            get => _instanceTags ?? (_instanceTags = new Dictionary<string, object>());
            set => _instanceTags = value;
        }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the Instance.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// The ID of the AMI used to launch the instance.
        /// </summary>
        public readonly string Ami;
        /// <summary>
        /// The ARN of the instance.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Whether or not the Instance is associated with a public IP address or not (Boolean).
        /// </summary>
        public readonly bool AssociatePublicIpAddress;
        /// <summary>
        /// The availability zone of the Instance.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The credit specification of the Instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceCreditSpecificationResult> CreditSpecifications;
        public readonly bool DisableApiTermination;
        /// <summary>
        /// The EBS block device mappings of the Instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceEbsBlockDeviceResult> EbsBlockDevices;
        /// <summary>
        /// Whether the Instance is EBS optimized or not (Boolean).
        /// </summary>
        public readonly bool EbsOptimized;
        /// <summary>
        /// The ephemeral block device mappings of the Instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceEphemeralBlockDeviceResult> EphemeralBlockDevices;
        public readonly ImmutableArray<Outputs.GetInstanceFilterResult> Filters;
        public readonly bool? GetPasswordData;
        public readonly bool? GetUserData;
        /// <summary>
        /// The Id of the dedicated host the instance will be assigned to.
        /// </summary>
        public readonly string HostId;
        /// <summary>
        /// The name of the instance profile associated with the Instance.
        /// </summary>
        public readonly string IamInstanceProfile;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        /// <summary>
        /// The state of the instance. One of: `pending`, `running`, `shutting-down`, `terminated`, `stopping`, `stopped`. See [Instance Lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html) for more information.
        /// </summary>
        public readonly string InstanceState;
        public readonly ImmutableDictionary<string, object> InstanceTags;
        /// <summary>
        /// The type of the Instance.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The key name of the Instance.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// The metadata options of the Instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceMetadataOptionResult> MetadataOptions;
        /// <summary>
        /// Whether detailed monitoring is enabled or disabled for the Instance (Boolean).
        /// </summary>
        public readonly bool Monitoring;
        /// <summary>
        /// The ID of the network interface that was created with the Instance.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// Base-64 encoded encrypted password data for the instance.
        /// Useful for getting the administrator password for instances running Microsoft Windows.
        /// This attribute is only exported if `get_password_data` is true.
        /// See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
        /// </summary>
        public readonly string PasswordData;
        /// <summary>
        /// The placement group of the Instance.
        /// </summary>
        public readonly string PlacementGroup;
        /// <summary>
        /// The private DNS name assigned to the Instance. Can only be
        /// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
        /// for your VPC.
        /// </summary>
        public readonly string PrivateDns;
        /// <summary>
        /// The private IP address assigned to the Instance.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// The public DNS name assigned to the Instance. For EC2-VPC, this
        /// is only available if you've enabled DNS hostnames for your VPC.
        /// </summary>
        public readonly string PublicDns;
        /// <summary>
        /// The public IP address assigned to the Instance, if applicable. **NOTE**: If you are using an [`aws.ec2.Eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// The root block device mappings of the Instance
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceRootBlockDeviceResult> RootBlockDevices;
        /// <summary>
        /// The associated security groups.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// Whether the network interface performs source/destination checking (Boolean).
        /// </summary>
        public readonly bool SourceDestCheck;
        /// <summary>
        /// The VPC subnet ID.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// A mapping of tags assigned to the Instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The tenancy of the instance: `dedicated`, `default`, `host`.
        /// </summary>
        public readonly string Tenancy;
        /// <summary>
        /// SHA-1 hash of User Data supplied to the Instance.
        /// </summary>
        public readonly string UserData;
        /// <summary>
        /// Base64 encoded contents of User Data supplied to the Instance. This attribute is only exported if `get_user_data` is true.
        /// </summary>
        public readonly string UserDataBase64;
        /// <summary>
        /// The associated security groups in a non-default VPC.
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private GetInstanceResult(
            string ami,

            string arn,

            bool associatePublicIpAddress,

            string availabilityZone,

            ImmutableArray<Outputs.GetInstanceCreditSpecificationResult> creditSpecifications,

            bool disableApiTermination,

            ImmutableArray<Outputs.GetInstanceEbsBlockDeviceResult> ebsBlockDevices,

            bool ebsOptimized,

            ImmutableArray<Outputs.GetInstanceEphemeralBlockDeviceResult> ephemeralBlockDevices,

            ImmutableArray<Outputs.GetInstanceFilterResult> filters,

            bool? getPasswordData,

            bool? getUserData,

            string hostId,

            string iamInstanceProfile,

            string id,

            string? instanceId,

            string instanceState,

            ImmutableDictionary<string, object> instanceTags,

            string instanceType,

            string keyName,

            ImmutableArray<Outputs.GetInstanceMetadataOptionResult> metadataOptions,

            bool monitoring,

            string networkInterfaceId,

            string passwordData,

            string placementGroup,

            string privateDns,

            string privateIp,

            string publicDns,

            string publicIp,

            ImmutableArray<Outputs.GetInstanceRootBlockDeviceResult> rootBlockDevices,

            ImmutableArray<string> securityGroups,

            bool sourceDestCheck,

            string subnetId,

            ImmutableDictionary<string, object> tags,

            string tenancy,

            string userData,

            string userDataBase64,

            ImmutableArray<string> vpcSecurityGroupIds)
        {
            Ami = ami;
            Arn = arn;
            AssociatePublicIpAddress = associatePublicIpAddress;
            AvailabilityZone = availabilityZone;
            CreditSpecifications = creditSpecifications;
            DisableApiTermination = disableApiTermination;
            EbsBlockDevices = ebsBlockDevices;
            EbsOptimized = ebsOptimized;
            EphemeralBlockDevices = ephemeralBlockDevices;
            Filters = filters;
            GetPasswordData = getPasswordData;
            GetUserData = getUserData;
            HostId = hostId;
            IamInstanceProfile = iamInstanceProfile;
            Id = id;
            InstanceId = instanceId;
            InstanceState = instanceState;
            InstanceTags = instanceTags;
            InstanceType = instanceType;
            KeyName = keyName;
            MetadataOptions = metadataOptions;
            Monitoring = monitoring;
            NetworkInterfaceId = networkInterfaceId;
            PasswordData = passwordData;
            PlacementGroup = placementGroup;
            PrivateDns = privateDns;
            PrivateIp = privateIp;
            PublicDns = publicDns;
            PublicIp = publicIp;
            RootBlockDevices = rootBlockDevices;
            SecurityGroups = securityGroups;
            SourceDestCheck = sourceDestCheck;
            SubnetId = subnetId;
            Tags = tags;
            Tenancy = tenancy;
            UserData = userData;
            UserDataBase64 = userDataBase64;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
        }
    }
}

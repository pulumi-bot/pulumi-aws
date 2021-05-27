// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetLaunchConfiguration
    {
        /// <summary>
        /// Provides information about a Launch Configuration.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ubuntu = Output.Create(Aws.Ec2.GetLaunchConfiguration.InvokeAsync(new Aws.Ec2.GetLaunchConfigurationArgs
        ///         {
        ///             Name = "test-launch-config",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLaunchConfigurationResult> InvokeAsync(GetLaunchConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLaunchConfigurationResult>("aws:ec2/getLaunchConfiguration:getLaunchConfiguration", args ?? new GetLaunchConfigurationArgs(), options.WithVersion());

        public static Output<GetLaunchConfigurationResult> Invoke(GetLaunchConfigurationOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box()
            ).Apply(a => {
                    var args = new GetLaunchConfigurationArgs();
                    a[0].Set(args, nameof(args.Name));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetLaunchConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the launch configuration.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetLaunchConfigurationArgs()
        {
        }
    }

    public sealed class GetLaunchConfigurationOutputArgs
    {
        /// <summary>
        /// The name of the launch configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetLaunchConfigurationOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLaunchConfigurationResult
    {
        /// <summary>
        /// The Amazon Resource Name of the launch configuration.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Whether a Public IP address is associated with the instance.
        /// </summary>
        public readonly bool AssociatePublicIpAddress;
        /// <summary>
        /// The EBS Block Devices attached to the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchConfigurationEbsBlockDeviceResult> EbsBlockDevices;
        /// <summary>
        /// Whether the launched EC2 instance will be EBS-optimized.
        /// </summary>
        public readonly bool EbsOptimized;
        /// <summary>
        /// Whether Detailed Monitoring is Enabled.
        /// </summary>
        public readonly bool EnableMonitoring;
        /// <summary>
        /// The Ephemeral volumes on the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchConfigurationEphemeralBlockDeviceResult> EphemeralBlockDevices;
        /// <summary>
        /// The IAM Instance Profile to associate with launched instances.
        /// </summary>
        public readonly string IamInstanceProfile;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The EC2 Image ID of the instance.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The Instance Type of the instance to launch.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The Key Name that should be used for the instance.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// The metadata options for the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchConfigurationMetadataOptionResult> MetadataOptions;
        /// <summary>
        /// The Name of the launch configuration.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Tenancy of the instance.
        /// </summary>
        public readonly string PlacementTenancy;
        /// <summary>
        /// The Root Block Device of the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchConfigurationRootBlockDeviceResult> RootBlockDevices;
        /// <summary>
        /// A list of associated Security Group IDS.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// The Price to use for reserving Spot instances.
        /// </summary>
        public readonly string SpotPrice;
        /// <summary>
        /// The User Data of the instance.
        /// </summary>
        public readonly string UserData;
        /// <summary>
        /// The ID of a ClassicLink-enabled VPC.
        /// </summary>
        public readonly string VpcClassicLinkId;
        /// <summary>
        /// The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.
        /// </summary>
        public readonly ImmutableArray<string> VpcClassicLinkSecurityGroups;

        [OutputConstructor]
        private GetLaunchConfigurationResult(
            string arn,

            bool associatePublicIpAddress,

            ImmutableArray<Outputs.GetLaunchConfigurationEbsBlockDeviceResult> ebsBlockDevices,

            bool ebsOptimized,

            bool enableMonitoring,

            ImmutableArray<Outputs.GetLaunchConfigurationEphemeralBlockDeviceResult> ephemeralBlockDevices,

            string iamInstanceProfile,

            string id,

            string imageId,

            string instanceType,

            string keyName,

            ImmutableArray<Outputs.GetLaunchConfigurationMetadataOptionResult> metadataOptions,

            string name,

            string placementTenancy,

            ImmutableArray<Outputs.GetLaunchConfigurationRootBlockDeviceResult> rootBlockDevices,

            ImmutableArray<string> securityGroups,

            string spotPrice,

            string userData,

            string vpcClassicLinkId,

            ImmutableArray<string> vpcClassicLinkSecurityGroups)
        {
            Arn = arn;
            AssociatePublicIpAddress = associatePublicIpAddress;
            EbsBlockDevices = ebsBlockDevices;
            EbsOptimized = ebsOptimized;
            EnableMonitoring = enableMonitoring;
            EphemeralBlockDevices = ephemeralBlockDevices;
            IamInstanceProfile = iamInstanceProfile;
            Id = id;
            ImageId = imageId;
            InstanceType = instanceType;
            KeyName = keyName;
            MetadataOptions = metadataOptions;
            Name = name;
            PlacementTenancy = placementTenancy;
            RootBlockDevices = rootBlockDevices;
            SecurityGroups = securityGroups;
            SpotPrice = spotPrice;
            UserData = userData;
            VpcClassicLinkId = vpcClassicLinkId;
            VpcClassicLinkSecurityGroups = vpcClassicLinkSecurityGroups;
        }
    }
}

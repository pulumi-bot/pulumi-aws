// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Batch.Outputs
{

    [OutputType]
    public sealed class ComputeEnvironmentComputeResources
    {
        /// <summary>
        /// The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated. Valid items are `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED` or `BEST_FIT`. Defaults to `BEST_FIT`. See [AWS docs](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) for details.
        /// </summary>
        public readonly string? AllocationStrategy;
        /// <summary>
        /// Integer of minimum percentage that a Spot Instance price must be when compared with the On-Demand price for that instance type before instances are launched. For example, if your bid percentage is 20% (`20`), then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. This parameter is required for SPOT compute environments.
        /// </summary>
        public readonly int? BidPercentage;
        /// <summary>
        /// The desired number of EC2 vCPUS in the compute environment.
        /// </summary>
        public readonly int? DesiredVcpus;
        /// <summary>
        /// The EC2 key pair that is used for instances launched in the compute environment.
        /// </summary>
        public readonly string? Ec2KeyPair;
        /// <summary>
        /// The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
        /// </summary>
        public readonly string? ImageId;
        /// <summary>
        /// The Amazon ECS instance role applied to Amazon EC2 instances in a compute environment.
        /// </summary>
        public readonly string InstanceRole;
        /// <summary>
        /// A list of instance types that may be launched.
        /// </summary>
        public readonly ImmutableArray<string> InstanceTypes;
        /// <summary>
        /// The launch template to use for your compute resources. See details below.
        /// </summary>
        public readonly Outputs.ComputeEnvironmentComputeResourcesLaunchTemplate? LaunchTemplate;
        /// <summary>
        /// The maximum number of EC2 vCPUs that an environment can reach.
        /// </summary>
        public readonly int MaxVcpus;
        /// <summary>
        /// The minimum number of EC2 vCPUs that an environment should maintain.
        /// </summary>
        public readonly int MinVcpus;
        /// <summary>
        /// A list of EC2 security group that are associated with instances launched in the compute environment.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment. This parameter is required for SPOT compute environments.
        /// </summary>
        public readonly string? SpotIamFleetRole;
        /// <summary>
        /// A list of VPC subnets into which the compute resources are launched.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;
        /// <summary>
        /// Key-value pair tags to be applied to resources that are launched in the compute environment.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// The type of the compute environment. Valid items are `MANAGED` or `UNMANAGED`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ComputeEnvironmentComputeResources(
            string? allocationStrategy,

            int? bidPercentage,

            int? desiredVcpus,

            string? ec2KeyPair,

            string? imageId,

            string instanceRole,

            ImmutableArray<string> instanceTypes,

            Outputs.ComputeEnvironmentComputeResourcesLaunchTemplate? launchTemplate,

            int maxVcpus,

            int minVcpus,

            ImmutableArray<string> securityGroupIds,

            string? spotIamFleetRole,

            ImmutableArray<string> subnets,

            ImmutableDictionary<string, object>? tags,

            string type)
        {
            AllocationStrategy = allocationStrategy;
            BidPercentage = bidPercentage;
            DesiredVcpus = desiredVcpus;
            Ec2KeyPair = ec2KeyPair;
            ImageId = imageId;
            InstanceRole = instanceRole;
            InstanceTypes = instanceTypes;
            LaunchTemplate = launchTemplate;
            MaxVcpus = maxVcpus;
            MinVcpus = minVcpus;
            SecurityGroupIds = securityGroupIds;
            SpotIamFleetRole = spotIamFleetRole;
            Subnets = subnets;
            Tags = tags;
            Type = type;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling
{
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to get information on an existing autoscaling group.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("aws:autoscaling/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithVersion());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the exact name of the desired autoscaling group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Auto Scaling group.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// One or more Availability Zones for the group.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly int DefaultCooldown;
        /// <summary>
        /// The desired size of the group.
        /// </summary>
        public readonly int DesiredCapacity;
        /// <summary>
        /// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
        /// </summary>
        public readonly int HealthCheckGracePeriod;
        /// <summary>
        /// The service to use for the health checks. The valid values are EC2 and ELB.
        /// </summary>
        public readonly string HealthCheckType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the associated launch configuration.
        /// </summary>
        public readonly string LaunchConfiguration;
        /// <summary>
        /// One or more load balancers associated with the group.
        /// </summary>
        public readonly ImmutableArray<string> LoadBalancers;
        /// <summary>
        /// The maximum size of the group.
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// The minimum size of the group.
        /// </summary>
        public readonly int MinSize;
        /// <summary>
        /// Name of the Auto Scaling Group.
        /// </summary>
        public readonly string Name;
        public readonly bool NewInstancesProtectedFromScaleIn;
        /// <summary>
        /// The name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
        /// </summary>
        public readonly string PlacementGroup;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
        /// </summary>
        public readonly string ServiceLinkedRoleArn;
        /// <summary>
        /// The current state of the group when DeleteAutoScalingGroup is in progress.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The Amazon Resource Names (ARN) of the target groups for your load balancer.
        /// </summary>
        public readonly ImmutableArray<string> TargetGroupArns;
        /// <summary>
        /// The termination policies for the group.
        /// </summary>
        public readonly ImmutableArray<string> TerminationPolicies;
        /// <summary>
        /// VPC ID for the group.
        /// </summary>
        public readonly string VpcZoneIdentifier;

        [OutputConstructor]
        private GetGroupResult(
            string arn,

            ImmutableArray<string> availabilityZones,

            int defaultCooldown,

            int desiredCapacity,

            int healthCheckGracePeriod,

            string healthCheckType,

            string id,

            string launchConfiguration,

            ImmutableArray<string> loadBalancers,

            int maxSize,

            int minSize,

            string name,

            bool newInstancesProtectedFromScaleIn,

            string placementGroup,

            string serviceLinkedRoleArn,

            string status,

            ImmutableArray<string> targetGroupArns,

            ImmutableArray<string> terminationPolicies,

            string vpcZoneIdentifier)
        {
            Arn = arn;
            AvailabilityZones = availabilityZones;
            DefaultCooldown = defaultCooldown;
            DesiredCapacity = desiredCapacity;
            HealthCheckGracePeriod = healthCheckGracePeriod;
            HealthCheckType = healthCheckType;
            Id = id;
            LaunchConfiguration = launchConfiguration;
            LoadBalancers = loadBalancers;
            MaxSize = maxSize;
            MinSize = minSize;
            Name = name;
            NewInstancesProtectedFromScaleIn = newInstancesProtectedFromScaleIn;
            PlacementGroup = placementGroup;
            ServiceLinkedRoleArn = serviceLinkedRoleArn;
            Status = status;
            TargetGroupArns = targetGroupArns;
            TerminationPolicies = terminationPolicies;
            VpcZoneIdentifier = vpcZoneIdentifier;
        }
    }
}

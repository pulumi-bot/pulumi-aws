// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs
{
    public static class GetContainerDefinition
    {
        /// <summary>
        /// The ECS container definition data source allows access to details of
        /// a specific container within an AWS ECS service.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetContainerDefinitionResult> InvokeAsync(GetContainerDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContainerDefinitionResult>("aws:ecs/getContainerDefinition:getContainerDefinition", args ?? new GetContainerDefinitionArgs(), options.WithVersion());
    }


    public sealed class GetContainerDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container definition
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// The ARN of the task definition which contains the container
        /// </summary>
        [Input("taskDefinition", required: true)]
        public string TaskDefinition { get; set; } = null!;

        public GetContainerDefinitionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContainerDefinitionResult
    {
        public readonly string ContainerName;
        /// <summary>
        /// The CPU limit for this container definition
        /// </summary>
        public readonly int Cpu;
        /// <summary>
        /// Indicator if networking is disabled
        /// </summary>
        public readonly bool DisableNetworking;
        /// <summary>
        /// Set docker labels
        /// </summary>
        public readonly ImmutableDictionary<string, string> DockerLabels;
        /// <summary>
        /// The environment in use
        /// </summary>
        public readonly ImmutableDictionary<string, string> Environment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The docker image in use, including the digest
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The digest of the docker image in use
        /// </summary>
        public readonly string ImageDigest;
        /// <summary>
        /// The memory limit for this container definition
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit
        /// </summary>
        public readonly int MemoryReservation;
        public readonly string TaskDefinition;

        [OutputConstructor]
        private GetContainerDefinitionResult(
            string containerName,

            int cpu,

            bool disableNetworking,

            ImmutableDictionary<string, string> dockerLabels,

            ImmutableDictionary<string, string> environment,

            string id,

            string image,

            string imageDigest,

            int memory,

            int memoryReservation,

            string taskDefinition)
        {
            ContainerName = containerName;
            Cpu = cpu;
            DisableNetworking = disableNetworking;
            DockerLabels = dockerLabels;
            Environment = environment;
            Id = id;
            Image = image;
            ImageDigest = imageDigest;
            Memory = memory;
            MemoryReservation = memoryReservation;
            TaskDefinition = taskDefinition;
        }
    }
}

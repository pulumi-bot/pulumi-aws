// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetInstanceType
    {
        /// <summary>
        /// Get characteristics for a single EC2 Instance Type.
        /// </summary>
        public static Task<GetInstanceTypeResult> InvokeAsync(GetInstanceTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypeResult>("aws:ec2/getInstanceType:getInstanceType", args ?? new GetInstanceTypeArgs(), options.WithVersion());
    }


    public sealed class GetInstanceTypeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The default number of cores for the instance type.
        /// </summary>
        [Input("defaultCores")]
        public int? DefaultCores { get; set; }

        /// <summary>
        /// The  default  number of threads per core for the instance type.
        /// </summary>
        [Input("defaultThreadsPerCore")]
        public int? DefaultThreadsPerCore { get; set; }

        [Input("fpgas")]
        private List<Inputs.GetInstanceTypeFpgaArgs>? _fpgas;

        /// <summary>
        /// Describes the FPGA accelerator settings for the instance type.
        /// * `fpgas.#.count` - The count of FPGA accelerators for the instance type.
        /// * `fpgas.#.manufacturer` - The manufacturer of the FPGA accelerator.
        /// * `fpgas.#.memory_size` - The size (in MiB) for the memory available to the FPGA accelerator.
        /// * `fpgas.#.name` - The name of the FPGA accelerator.
        /// </summary>
        public List<Inputs.GetInstanceTypeFpgaArgs> Fpgas
        {
            get => _fpgas ?? (_fpgas = new List<Inputs.GetInstanceTypeFpgaArgs>());
            set => _fpgas = value;
        }

        [Input("gpuses")]
        private List<Inputs.GetInstanceTypeGpusArgs>? _gpuses;

        /// <summary>
        /// Describes the GPU accelerators for the instance type.
        /// * `gpus.#.count` - The number of GPUs for the instance type.
        /// * `gpus.#.manufacturer` - The manufacturer of the GPU accelerator.
        /// * `gpus.#.memory_size` - The size (in MiB) for the memory available to the GPU accelerator.
        /// * `gpus.#.name` - The name of the GPU accelerator.
        /// </summary>
        public List<Inputs.GetInstanceTypeGpusArgs> Gpuses
        {
            get => _gpuses ?? (_gpuses = new List<Inputs.GetInstanceTypeGpusArgs>());
            set => _gpuses = value;
        }

        /// <summary>
        /// Indicates the hypervisor used for the instance type.
        /// * `inference_accelerators` Describes the Inference accelerators for the instance type.
        /// * `inference_accelerators.#.count` - The number of Inference accelerators for the instance type.
        /// * `inference_accelerators.#.manufacturer` - The manufacturer of the Inference accelerator.
        /// * `inference_accelerators.#.name` - The name of the Inference accelerator.
        /// </summary>
        [Input("hypervisor")]
        public string? Hypervisor { get; set; }

        [Input("inferenceAccelerators")]
        private List<Inputs.GetInstanceTypeInferenceAcceleratorArgs>? _inferenceAccelerators;
        public List<Inputs.GetInstanceTypeInferenceAcceleratorArgs> InferenceAccelerators
        {
            get => _inferenceAccelerators ?? (_inferenceAccelerators = new List<Inputs.GetInstanceTypeInferenceAcceleratorArgs>());
            set => _inferenceAccelerators = value;
        }

        [Input("instanceDisks")]
        private List<Inputs.GetInstanceTypeInstanceDiskArgs>? _instanceDisks;

        /// <summary>
        /// Describes the disks for the instance type.
        /// * `instance_disks.#.count` - The number of disks with this configuration.
        /// * `instance_disks.#.size` - The size of the disk in GB.
        /// * `instance_disks.#.type` - The type of disk.
        /// </summary>
        public List<Inputs.GetInstanceTypeInstanceDiskArgs> InstanceDisks
        {
            get => _instanceDisks ?? (_instanceDisks = new List<Inputs.GetInstanceTypeInstanceDiskArgs>());
            set => _instanceDisks = value;
        }

        /// <summary>
        /// Instance
        /// </summary>
        [Input("instanceType", required: true)]
        public string InstanceType { get; set; } = null!;

        /// <summary>
        /// The maximum number of IPv6 addresses per network interface.
        /// </summary>
        [Input("maximumIpv6AddressesPerInterface")]
        public int? MaximumIpv6AddressesPerInterface { get; set; }

        /// <summary>
        /// The total memory of all FPGA accelerators for the instance type (in MiB).
        /// </summary>
        [Input("totalFpgaMemory")]
        public int? TotalFpgaMemory { get; set; }

        /// <summary>
        /// The total size of the memory for the GPU accelerators for the instance type (in MiB).
        /// </summary>
        [Input("totalGpuMemory")]
        public int? TotalGpuMemory { get; set; }

        /// <summary>
        /// The total size of the instance disks, in GB.
        /// </summary>
        [Input("totalInstanceStorage")]
        public int? TotalInstanceStorage { get; set; }

        public GetInstanceTypeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceTypeResult
    {
        /// <summary>
        /// `true` if auto recovery is supported.
        /// </summary>
        public readonly bool AutoRecoverySupported;
        /// <summary>
        /// `true` if it is a bare metal instance type.
        /// </summary>
        public readonly bool BareMetal;
        /// <summary>
        /// `true` if the instance type is a burstable performance instance type.
        /// </summary>
        public readonly bool BurstablePerformanceSupported;
        /// <summary>
        /// `true`  if the instance type is a current generation.
        /// </summary>
        public readonly bool CurrentGeneration;
        /// <summary>
        /// `true` if Dedicated Hosts are supported on the instance type.
        /// </summary>
        public readonly bool DedicatedHostsSupported;
        /// <summary>
        /// The default number of cores for the instance type.
        /// </summary>
        public readonly int DefaultCores;
        /// <summary>
        /// The  default  number of threads per core for the instance type.
        /// </summary>
        public readonly int DefaultThreadsPerCore;
        /// <summary>
        /// The default number of vCPUs for the instance type.
        /// </summary>
        public readonly int DefaultVcpus;
        /// <summary>
        /// Indicates whether Amazon EBS encryption is supported.
        /// </summary>
        public readonly string EbsEncryptionSupport;
        /// <summary>
        /// Indicates whether non-volatile memory express (NVMe) is supported.
        /// </summary>
        public readonly string EbsNvmeSupport;
        /// <summary>
        /// Indicates that the instance type is Amazon EBS-optimized.
        /// </summary>
        public readonly string EbsOptimizedSupport;
        /// <summary>
        /// The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.
        /// </summary>
        public readonly int EbsPerformanceBaselineBandwidth;
        /// <summary>
        /// The baseline input/output storage operations per seconds for an EBS-optimized instance type.
        /// </summary>
        public readonly int EbsPerformanceBaselineIops;
        /// <summary>
        /// The baseline throughput performance for an EBS-optimized instance type, in MBps.
        /// </summary>
        public readonly double EbsPerformanceBaselineThroughput;
        /// <summary>
        /// The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.
        /// </summary>
        public readonly int EbsPerformanceMaximumBandwidth;
        /// <summary>
        /// The maximum input/output storage operations per second for an EBS-optimized instance type.
        /// </summary>
        public readonly int EbsPerformanceMaximumIops;
        /// <summary>
        /// The maximum throughput performance for an EBS-optimized instance type, in MBps.
        /// </summary>
        public readonly double EbsPerformanceMaximumThroughput;
        /// <summary>
        /// Indicates whether Elastic Fabric Adapter (EFA) is supported.
        /// </summary>
        public readonly bool EfaSupported;
        /// <summary>
        /// Indicates whether Elastic Network Adapter (ENA) is supported.
        /// </summary>
        public readonly string EnaSupport;
        /// <summary>
        /// Describes the FPGA accelerator settings for the instance type.
        /// * `fpgas.#.count` - The count of FPGA accelerators for the instance type.
        /// * `fpgas.#.manufacturer` - The manufacturer of the FPGA accelerator.
        /// * `fpgas.#.memory_size` - The size (in MiB) for the memory available to the FPGA accelerator.
        /// * `fpgas.#.name` - The name of the FPGA accelerator.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypeFpgaResult> Fpgas;
        /// <summary>
        /// `true` if the instance type is eligible for the free tier.
        /// </summary>
        public readonly bool FreeTierEligible;
        /// <summary>
        /// Describes the GPU accelerators for the instance type.
        /// * `gpus.#.count` - The number of GPUs for the instance type.
        /// * `gpus.#.manufacturer` - The manufacturer of the GPU accelerator.
        /// * `gpus.#.memory_size` - The size (in MiB) for the memory available to the GPU accelerator.
        /// * `gpus.#.name` - The name of the GPU accelerator.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypeGpusResult> Gpuses;
        /// <summary>
        /// `true` if On-Demand hibernation is supported.
        /// </summary>
        public readonly bool HibernationSupported;
        /// <summary>
        /// Indicates the hypervisor used for the instance type.
        /// * `inference_accelerators` Describes the Inference accelerators for the instance type.
        /// * `inference_accelerators.#.count` - The number of Inference accelerators for the instance type.
        /// * `inference_accelerators.#.manufacturer` - The manufacturer of the Inference accelerator.
        /// * `inference_accelerators.#.name` - The name of the Inference accelerator.
        /// </summary>
        public readonly string Hypervisor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetInstanceTypeInferenceAcceleratorResult> InferenceAccelerators;
        /// <summary>
        /// Describes the disks for the instance type.
        /// * `instance_disks.#.count` - The number of disks with this configuration.
        /// * `instance_disks.#.size` - The size of the disk in GB.
        /// * `instance_disks.#.type` - The type of disk.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypeInstanceDiskResult> InstanceDisks;
        /// <summary>
        /// `true` if instance storage is supported.
        /// </summary>
        public readonly bool InstanceStorageSupported;
        public readonly string InstanceType;
        /// <summary>
        /// `true` if IPv6 is supported.
        /// </summary>
        public readonly bool Ipv6Supported;
        /// <summary>
        /// The maximum number of IPv4 addresses per network interface.
        /// </summary>
        public readonly int MaximumIpv4AddressesPerInterface;
        /// <summary>
        /// The maximum number of IPv6 addresses per network interface.
        /// </summary>
        public readonly int MaximumIpv6AddressesPerInterface;
        /// <summary>
        /// The maximum number of network interfaces for the instance type.
        /// </summary>
        public readonly int MaximumNetworkInterfaces;
        /// <summary>
        /// Size of the instance memory, in MiB.
        /// </summary>
        public readonly int MemorySize;
        /// <summary>
        /// Describes the network performance.
        /// </summary>
        public readonly string NetworkPerformance;
        /// <summary>
        /// A list of architectures supported by the instance type.
        /// </summary>
        public readonly ImmutableArray<string> SupportedArchitectures;
        /// <summary>
        /// A list of supported placement groups types.
        /// </summary>
        public readonly ImmutableArray<string> SupportedPlacementStrategies;
        /// <summary>
        /// Indicates the supported root device types.
        /// </summary>
        public readonly ImmutableArray<string> SupportedRootDeviceTypes;
        /// <summary>
        /// Indicates whether the instance type is offered for spot or On-Demand.
        /// </summary>
        public readonly ImmutableArray<string> SupportedUsagesClasses;
        /// <summary>
        /// The supported virtualization types.
        /// </summary>
        public readonly ImmutableArray<string> SupportedVirtualizationTypes;
        /// <summary>
        /// The speed of the processor, in GHz.
        /// </summary>
        public readonly double SustainedClockSpeed;
        /// <summary>
        /// The total memory of all FPGA accelerators for the instance type (in MiB).
        /// </summary>
        public readonly int TotalFpgaMemory;
        /// <summary>
        /// The total size of the memory for the GPU accelerators for the instance type (in MiB).
        /// </summary>
        public readonly int TotalGpuMemory;
        /// <summary>
        /// The total size of the instance disks, in GB.
        /// </summary>
        public readonly int TotalInstanceStorage;
        /// <summary>
        /// List of the valid number of cores that can be configured for the instance type.
        /// </summary>
        public readonly ImmutableArray<int> ValidCores;
        /// <summary>
        /// List of the valid number of threads per core that can be configured for the instance type.
        /// </summary>
        public readonly ImmutableArray<int> ValidThreadsPerCores;

        [OutputConstructor]
        private GetInstanceTypeResult(
            bool autoRecoverySupported,

            bool bareMetal,

            bool burstablePerformanceSupported,

            bool currentGeneration,

            bool dedicatedHostsSupported,

            int defaultCores,

            int defaultThreadsPerCore,

            int defaultVcpus,

            string ebsEncryptionSupport,

            string ebsNvmeSupport,

            string ebsOptimizedSupport,

            int ebsPerformanceBaselineBandwidth,

            int ebsPerformanceBaselineIops,

            double ebsPerformanceBaselineThroughput,

            int ebsPerformanceMaximumBandwidth,

            int ebsPerformanceMaximumIops,

            double ebsPerformanceMaximumThroughput,

            bool efaSupported,

            string enaSupport,

            ImmutableArray<Outputs.GetInstanceTypeFpgaResult> fpgas,

            bool freeTierEligible,

            ImmutableArray<Outputs.GetInstanceTypeGpusResult> gpuses,

            bool hibernationSupported,

            string hypervisor,

            string id,

            ImmutableArray<Outputs.GetInstanceTypeInferenceAcceleratorResult> inferenceAccelerators,

            ImmutableArray<Outputs.GetInstanceTypeInstanceDiskResult> instanceDisks,

            bool instanceStorageSupported,

            string instanceType,

            bool ipv6Supported,

            int maximumIpv4AddressesPerInterface,

            int maximumIpv6AddressesPerInterface,

            int maximumNetworkInterfaces,

            int memorySize,

            string networkPerformance,

            ImmutableArray<string> supportedArchitectures,

            ImmutableArray<string> supportedPlacementStrategies,

            ImmutableArray<string> supportedRootDeviceTypes,

            ImmutableArray<string> supportedUsagesClasses,

            ImmutableArray<string> supportedVirtualizationTypes,

            double sustainedClockSpeed,

            int totalFpgaMemory,

            int totalGpuMemory,

            int totalInstanceStorage,

            ImmutableArray<int> validCores,

            ImmutableArray<int> validThreadsPerCores)
        {
            AutoRecoverySupported = autoRecoverySupported;
            BareMetal = bareMetal;
            BurstablePerformanceSupported = burstablePerformanceSupported;
            CurrentGeneration = currentGeneration;
            DedicatedHostsSupported = dedicatedHostsSupported;
            DefaultCores = defaultCores;
            DefaultThreadsPerCore = defaultThreadsPerCore;
            DefaultVcpus = defaultVcpus;
            EbsEncryptionSupport = ebsEncryptionSupport;
            EbsNvmeSupport = ebsNvmeSupport;
            EbsOptimizedSupport = ebsOptimizedSupport;
            EbsPerformanceBaselineBandwidth = ebsPerformanceBaselineBandwidth;
            EbsPerformanceBaselineIops = ebsPerformanceBaselineIops;
            EbsPerformanceBaselineThroughput = ebsPerformanceBaselineThroughput;
            EbsPerformanceMaximumBandwidth = ebsPerformanceMaximumBandwidth;
            EbsPerformanceMaximumIops = ebsPerformanceMaximumIops;
            EbsPerformanceMaximumThroughput = ebsPerformanceMaximumThroughput;
            EfaSupported = efaSupported;
            EnaSupport = enaSupport;
            Fpgas = fpgas;
            FreeTierEligible = freeTierEligible;
            Gpuses = gpuses;
            HibernationSupported = hibernationSupported;
            Hypervisor = hypervisor;
            Id = id;
            InferenceAccelerators = inferenceAccelerators;
            InstanceDisks = instanceDisks;
            InstanceStorageSupported = instanceStorageSupported;
            InstanceType = instanceType;
            Ipv6Supported = ipv6Supported;
            MaximumIpv4AddressesPerInterface = maximumIpv4AddressesPerInterface;
            MaximumIpv6AddressesPerInterface = maximumIpv6AddressesPerInterface;
            MaximumNetworkInterfaces = maximumNetworkInterfaces;
            MemorySize = memorySize;
            NetworkPerformance = networkPerformance;
            SupportedArchitectures = supportedArchitectures;
            SupportedPlacementStrategies = supportedPlacementStrategies;
            SupportedRootDeviceTypes = supportedRootDeviceTypes;
            SupportedUsagesClasses = supportedUsagesClasses;
            SupportedVirtualizationTypes = supportedVirtualizationTypes;
            SustainedClockSpeed = sustainedClockSpeed;
            TotalFpgaMemory = totalFpgaMemory;
            TotalGpuMemory = totalGpuMemory;
            TotalInstanceStorage = totalInstanceStorage;
            ValidCores = validCores;
            ValidThreadsPerCores = validThreadsPerCores;
        }
    }
}

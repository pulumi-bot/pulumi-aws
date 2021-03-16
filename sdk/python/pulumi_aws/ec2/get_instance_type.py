# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetInstanceTypeResult',
    'AwaitableGetInstanceTypeResult',
    'get_instance_type',
]

@pulumi.output_type
class GetInstanceTypeResult:
    """
    A collection of values returned by getInstanceType.
    """
    def __init__(__self__, auto_recovery_supported=None, bare_metal=None, burstable_performance_supported=None, current_generation=None, dedicated_hosts_supported=None, default_cores=None, default_threads_per_core=None, default_vcpus=None, ebs_encryption_support=None, ebs_nvme_support=None, ebs_optimized_support=None, ebs_performance_baseline_bandwidth=None, ebs_performance_baseline_iops=None, ebs_performance_baseline_throughput=None, ebs_performance_maximum_bandwidth=None, ebs_performance_maximum_iops=None, ebs_performance_maximum_throughput=None, efa_supported=None, ena_support=None, fpgas=None, free_tier_eligible=None, gpuses=None, hibernation_supported=None, hypervisor=None, id=None, inference_accelerators=None, instance_disks=None, instance_storage_supported=None, instance_type=None, ipv6_supported=None, maximum_ipv4_addresses_per_interface=None, maximum_ipv6_addresses_per_interface=None, maximum_network_interfaces=None, memory_size=None, network_performance=None, supported_architectures=None, supported_placement_strategies=None, supported_root_device_types=None, supported_usages_classes=None, supported_virtualization_types=None, sustained_clock_speed=None, total_fpga_memory=None, total_gpu_memory=None, total_instance_storage=None, valid_cores=None, valid_threads_per_cores=None):
        if auto_recovery_supported and not isinstance(auto_recovery_supported, bool):
            raise TypeError("Expected argument 'auto_recovery_supported' to be a bool")
        pulumi.set(__self__, "auto_recovery_supported", auto_recovery_supported)
        if bare_metal and not isinstance(bare_metal, bool):
            raise TypeError("Expected argument 'bare_metal' to be a bool")
        pulumi.set(__self__, "bare_metal", bare_metal)
        if burstable_performance_supported and not isinstance(burstable_performance_supported, bool):
            raise TypeError("Expected argument 'burstable_performance_supported' to be a bool")
        pulumi.set(__self__, "burstable_performance_supported", burstable_performance_supported)
        if current_generation and not isinstance(current_generation, bool):
            raise TypeError("Expected argument 'current_generation' to be a bool")
        pulumi.set(__self__, "current_generation", current_generation)
        if dedicated_hosts_supported and not isinstance(dedicated_hosts_supported, bool):
            raise TypeError("Expected argument 'dedicated_hosts_supported' to be a bool")
        pulumi.set(__self__, "dedicated_hosts_supported", dedicated_hosts_supported)
        if default_cores and not isinstance(default_cores, int):
            raise TypeError("Expected argument 'default_cores' to be a int")
        pulumi.set(__self__, "default_cores", default_cores)
        if default_threads_per_core and not isinstance(default_threads_per_core, int):
            raise TypeError("Expected argument 'default_threads_per_core' to be a int")
        pulumi.set(__self__, "default_threads_per_core", default_threads_per_core)
        if default_vcpus and not isinstance(default_vcpus, int):
            raise TypeError("Expected argument 'default_vcpus' to be a int")
        pulumi.set(__self__, "default_vcpus", default_vcpus)
        if ebs_encryption_support and not isinstance(ebs_encryption_support, str):
            raise TypeError("Expected argument 'ebs_encryption_support' to be a str")
        pulumi.set(__self__, "ebs_encryption_support", ebs_encryption_support)
        if ebs_nvme_support and not isinstance(ebs_nvme_support, str):
            raise TypeError("Expected argument 'ebs_nvme_support' to be a str")
        pulumi.set(__self__, "ebs_nvme_support", ebs_nvme_support)
        if ebs_optimized_support and not isinstance(ebs_optimized_support, str):
            raise TypeError("Expected argument 'ebs_optimized_support' to be a str")
        pulumi.set(__self__, "ebs_optimized_support", ebs_optimized_support)
        if ebs_performance_baseline_bandwidth and not isinstance(ebs_performance_baseline_bandwidth, int):
            raise TypeError("Expected argument 'ebs_performance_baseline_bandwidth' to be a int")
        pulumi.set(__self__, "ebs_performance_baseline_bandwidth", ebs_performance_baseline_bandwidth)
        if ebs_performance_baseline_iops and not isinstance(ebs_performance_baseline_iops, int):
            raise TypeError("Expected argument 'ebs_performance_baseline_iops' to be a int")
        pulumi.set(__self__, "ebs_performance_baseline_iops", ebs_performance_baseline_iops)
        if ebs_performance_baseline_throughput and not isinstance(ebs_performance_baseline_throughput, float):
            raise TypeError("Expected argument 'ebs_performance_baseline_throughput' to be a float")
        pulumi.set(__self__, "ebs_performance_baseline_throughput", ebs_performance_baseline_throughput)
        if ebs_performance_maximum_bandwidth and not isinstance(ebs_performance_maximum_bandwidth, int):
            raise TypeError("Expected argument 'ebs_performance_maximum_bandwidth' to be a int")
        pulumi.set(__self__, "ebs_performance_maximum_bandwidth", ebs_performance_maximum_bandwidth)
        if ebs_performance_maximum_iops and not isinstance(ebs_performance_maximum_iops, int):
            raise TypeError("Expected argument 'ebs_performance_maximum_iops' to be a int")
        pulumi.set(__self__, "ebs_performance_maximum_iops", ebs_performance_maximum_iops)
        if ebs_performance_maximum_throughput and not isinstance(ebs_performance_maximum_throughput, float):
            raise TypeError("Expected argument 'ebs_performance_maximum_throughput' to be a float")
        pulumi.set(__self__, "ebs_performance_maximum_throughput", ebs_performance_maximum_throughput)
        if efa_supported and not isinstance(efa_supported, bool):
            raise TypeError("Expected argument 'efa_supported' to be a bool")
        pulumi.set(__self__, "efa_supported", efa_supported)
        if ena_support and not isinstance(ena_support, str):
            raise TypeError("Expected argument 'ena_support' to be a str")
        pulumi.set(__self__, "ena_support", ena_support)
        if fpgas and not isinstance(fpgas, list):
            raise TypeError("Expected argument 'fpgas' to be a list")
        pulumi.set(__self__, "fpgas", fpgas)
        if free_tier_eligible and not isinstance(free_tier_eligible, bool):
            raise TypeError("Expected argument 'free_tier_eligible' to be a bool")
        pulumi.set(__self__, "free_tier_eligible", free_tier_eligible)
        if gpuses and not isinstance(gpuses, list):
            raise TypeError("Expected argument 'gpuses' to be a list")
        pulumi.set(__self__, "gpuses", gpuses)
        if hibernation_supported and not isinstance(hibernation_supported, bool):
            raise TypeError("Expected argument 'hibernation_supported' to be a bool")
        pulumi.set(__self__, "hibernation_supported", hibernation_supported)
        if hypervisor and not isinstance(hypervisor, str):
            raise TypeError("Expected argument 'hypervisor' to be a str")
        pulumi.set(__self__, "hypervisor", hypervisor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inference_accelerators and not isinstance(inference_accelerators, list):
            raise TypeError("Expected argument 'inference_accelerators' to be a list")
        pulumi.set(__self__, "inference_accelerators", inference_accelerators)
        if instance_disks and not isinstance(instance_disks, list):
            raise TypeError("Expected argument 'instance_disks' to be a list")
        pulumi.set(__self__, "instance_disks", instance_disks)
        if instance_storage_supported and not isinstance(instance_storage_supported, bool):
            raise TypeError("Expected argument 'instance_storage_supported' to be a bool")
        pulumi.set(__self__, "instance_storage_supported", instance_storage_supported)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if ipv6_supported and not isinstance(ipv6_supported, bool):
            raise TypeError("Expected argument 'ipv6_supported' to be a bool")
        pulumi.set(__self__, "ipv6_supported", ipv6_supported)
        if maximum_ipv4_addresses_per_interface and not isinstance(maximum_ipv4_addresses_per_interface, int):
            raise TypeError("Expected argument 'maximum_ipv4_addresses_per_interface' to be a int")
        pulumi.set(__self__, "maximum_ipv4_addresses_per_interface", maximum_ipv4_addresses_per_interface)
        if maximum_ipv6_addresses_per_interface and not isinstance(maximum_ipv6_addresses_per_interface, int):
            raise TypeError("Expected argument 'maximum_ipv6_addresses_per_interface' to be a int")
        pulumi.set(__self__, "maximum_ipv6_addresses_per_interface", maximum_ipv6_addresses_per_interface)
        if maximum_network_interfaces and not isinstance(maximum_network_interfaces, int):
            raise TypeError("Expected argument 'maximum_network_interfaces' to be a int")
        pulumi.set(__self__, "maximum_network_interfaces", maximum_network_interfaces)
        if memory_size and not isinstance(memory_size, int):
            raise TypeError("Expected argument 'memory_size' to be a int")
        pulumi.set(__self__, "memory_size", memory_size)
        if network_performance and not isinstance(network_performance, str):
            raise TypeError("Expected argument 'network_performance' to be a str")
        pulumi.set(__self__, "network_performance", network_performance)
        if supported_architectures and not isinstance(supported_architectures, list):
            raise TypeError("Expected argument 'supported_architectures' to be a list")
        pulumi.set(__self__, "supported_architectures", supported_architectures)
        if supported_placement_strategies and not isinstance(supported_placement_strategies, list):
            raise TypeError("Expected argument 'supported_placement_strategies' to be a list")
        pulumi.set(__self__, "supported_placement_strategies", supported_placement_strategies)
        if supported_root_device_types and not isinstance(supported_root_device_types, list):
            raise TypeError("Expected argument 'supported_root_device_types' to be a list")
        pulumi.set(__self__, "supported_root_device_types", supported_root_device_types)
        if supported_usages_classes and not isinstance(supported_usages_classes, list):
            raise TypeError("Expected argument 'supported_usages_classes' to be a list")
        pulumi.set(__self__, "supported_usages_classes", supported_usages_classes)
        if supported_virtualization_types and not isinstance(supported_virtualization_types, list):
            raise TypeError("Expected argument 'supported_virtualization_types' to be a list")
        pulumi.set(__self__, "supported_virtualization_types", supported_virtualization_types)
        if sustained_clock_speed and not isinstance(sustained_clock_speed, float):
            raise TypeError("Expected argument 'sustained_clock_speed' to be a float")
        pulumi.set(__self__, "sustained_clock_speed", sustained_clock_speed)
        if total_fpga_memory and not isinstance(total_fpga_memory, int):
            raise TypeError("Expected argument 'total_fpga_memory' to be a int")
        pulumi.set(__self__, "total_fpga_memory", total_fpga_memory)
        if total_gpu_memory and not isinstance(total_gpu_memory, int):
            raise TypeError("Expected argument 'total_gpu_memory' to be a int")
        pulumi.set(__self__, "total_gpu_memory", total_gpu_memory)
        if total_instance_storage and not isinstance(total_instance_storage, int):
            raise TypeError("Expected argument 'total_instance_storage' to be a int")
        pulumi.set(__self__, "total_instance_storage", total_instance_storage)
        if valid_cores and not isinstance(valid_cores, list):
            raise TypeError("Expected argument 'valid_cores' to be a list")
        pulumi.set(__self__, "valid_cores", valid_cores)
        if valid_threads_per_cores and not isinstance(valid_threads_per_cores, list):
            raise TypeError("Expected argument 'valid_threads_per_cores' to be a list")
        pulumi.set(__self__, "valid_threads_per_cores", valid_threads_per_cores)

    @property
    @pulumi.getter(name="autoRecoverySupported")
    def auto_recovery_supported(self) -> bool:
        """
        `true` if auto recovery is supported.
        """
        return pulumi.get(self, "auto_recovery_supported")

    @property
    @pulumi.getter(name="bareMetal")
    def bare_metal(self) -> bool:
        """
        `true` if it is a bare metal instance type.
        """
        return pulumi.get(self, "bare_metal")

    @property
    @pulumi.getter(name="burstablePerformanceSupported")
    def burstable_performance_supported(self) -> bool:
        """
        `true` if the instance type is a burstable performance instance type.
        """
        return pulumi.get(self, "burstable_performance_supported")

    @property
    @pulumi.getter(name="currentGeneration")
    def current_generation(self) -> bool:
        """
        `true`  if the instance type is a current generation.
        """
        return pulumi.get(self, "current_generation")

    @property
    @pulumi.getter(name="dedicatedHostsSupported")
    def dedicated_hosts_supported(self) -> bool:
        """
        `true` if Dedicated Hosts are supported on the instance type.
        """
        return pulumi.get(self, "dedicated_hosts_supported")

    @property
    @pulumi.getter(name="defaultCores")
    def default_cores(self) -> int:
        """
        The default number of cores for the instance type.
        """
        return pulumi.get(self, "default_cores")

    @property
    @pulumi.getter(name="defaultThreadsPerCore")
    def default_threads_per_core(self) -> int:
        """
        The  default  number of threads per core for the instance type.
        """
        return pulumi.get(self, "default_threads_per_core")

    @property
    @pulumi.getter(name="defaultVcpus")
    def default_vcpus(self) -> int:
        """
        The default number of vCPUs for the instance type.
        """
        return pulumi.get(self, "default_vcpus")

    @property
    @pulumi.getter(name="ebsEncryptionSupport")
    def ebs_encryption_support(self) -> str:
        """
        Indicates whether Amazon EBS encryption is supported.
        """
        return pulumi.get(self, "ebs_encryption_support")

    @property
    @pulumi.getter(name="ebsNvmeSupport")
    def ebs_nvme_support(self) -> str:
        """
        Indicates whether non-volatile memory express (NVMe) is supported.
        """
        return pulumi.get(self, "ebs_nvme_support")

    @property
    @pulumi.getter(name="ebsOptimizedSupport")
    def ebs_optimized_support(self) -> str:
        """
        Indicates that the instance type is Amazon EBS-optimized.
        """
        return pulumi.get(self, "ebs_optimized_support")

    @property
    @pulumi.getter(name="ebsPerformanceBaselineBandwidth")
    def ebs_performance_baseline_bandwidth(self) -> int:
        """
        The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.
        """
        return pulumi.get(self, "ebs_performance_baseline_bandwidth")

    @property
    @pulumi.getter(name="ebsPerformanceBaselineIops")
    def ebs_performance_baseline_iops(self) -> int:
        """
        The baseline input/output storage operations per seconds for an EBS-optimized instance type.
        """
        return pulumi.get(self, "ebs_performance_baseline_iops")

    @property
    @pulumi.getter(name="ebsPerformanceBaselineThroughput")
    def ebs_performance_baseline_throughput(self) -> float:
        """
        The baseline throughput performance for an EBS-optimized instance type, in MBps.
        """
        return pulumi.get(self, "ebs_performance_baseline_throughput")

    @property
    @pulumi.getter(name="ebsPerformanceMaximumBandwidth")
    def ebs_performance_maximum_bandwidth(self) -> int:
        """
        The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.
        """
        return pulumi.get(self, "ebs_performance_maximum_bandwidth")

    @property
    @pulumi.getter(name="ebsPerformanceMaximumIops")
    def ebs_performance_maximum_iops(self) -> int:
        """
        The maximum input/output storage operations per second for an EBS-optimized instance type.
        """
        return pulumi.get(self, "ebs_performance_maximum_iops")

    @property
    @pulumi.getter(name="ebsPerformanceMaximumThroughput")
    def ebs_performance_maximum_throughput(self) -> float:
        """
        The maximum throughput performance for an EBS-optimized instance type, in MBps.
        """
        return pulumi.get(self, "ebs_performance_maximum_throughput")

    @property
    @pulumi.getter(name="efaSupported")
    def efa_supported(self) -> bool:
        """
        Indicates whether Elastic Fabric Adapter (EFA) is supported.
        """
        return pulumi.get(self, "efa_supported")

    @property
    @pulumi.getter(name="enaSupport")
    def ena_support(self) -> str:
        """
        Indicates whether Elastic Network Adapter (ENA) is supported.
        """
        return pulumi.get(self, "ena_support")

    @property
    @pulumi.getter
    def fpgas(self) -> Sequence['outputs.GetInstanceTypeFpgaResult']:
        """
        Describes the FPGA accelerator settings for the instance type.
        * `fpgas.#.count` - The count of FPGA accelerators for the instance type.
        * `fpgas.#.manufacturer` - The manufacturer of the FPGA accelerator.
        * `fpgas.#.memory_size` - The size (in MiB) for the memory available to the FPGA accelerator.
        * `fpgas.#.name` - The name of the FPGA accelerator.
        """
        return pulumi.get(self, "fpgas")

    @property
    @pulumi.getter(name="freeTierEligible")
    def free_tier_eligible(self) -> bool:
        """
        `true` if the instance type is eligible for the free tier.
        """
        return pulumi.get(self, "free_tier_eligible")

    @property
    @pulumi.getter
    def gpuses(self) -> Sequence['outputs.GetInstanceTypeGpusResult']:
        """
        Describes the GPU accelerators for the instance type.
        * `gpus.#.count` - The number of GPUs for the instance type.
        * `gpus.#.manufacturer` - The manufacturer of the GPU accelerator.
        * `gpus.#.memory_size` - The size (in MiB) for the memory available to the GPU accelerator.
        * `gpus.#.name` - The name of the GPU accelerator.
        """
        return pulumi.get(self, "gpuses")

    @property
    @pulumi.getter(name="hibernationSupported")
    def hibernation_supported(self) -> bool:
        """
        `true` if On-Demand hibernation is supported.
        """
        return pulumi.get(self, "hibernation_supported")

    @property
    @pulumi.getter
    def hypervisor(self) -> str:
        """
        Indicates the hypervisor used for the instance type.
        * `inference_accelerators` Describes the Inference accelerators for the instance type.
        * `inference_accelerators.#.count` - The number of Inference accelerators for the instance type.
        * `inference_accelerators.#.manufacturer` - The manufacturer of the Inference accelerator.
        * `inference_accelerators.#.name` - The name of the Inference accelerator.
        """
        return pulumi.get(self, "hypervisor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inferenceAccelerators")
    def inference_accelerators(self) -> Sequence['outputs.GetInstanceTypeInferenceAcceleratorResult']:
        return pulumi.get(self, "inference_accelerators")

    @property
    @pulumi.getter(name="instanceDisks")
    def instance_disks(self) -> Sequence['outputs.GetInstanceTypeInstanceDiskResult']:
        """
        Describes the disks for the instance type.
        * `instance_disks.#.count` - The number of disks with this configuration.
        * `instance_disks.#.size` - The size of the disk in GB.
        * `instance_disks.#.type` - The type of disk.
        """
        return pulumi.get(self, "instance_disks")

    @property
    @pulumi.getter(name="instanceStorageSupported")
    def instance_storage_supported(self) -> bool:
        """
        `true` if instance storage is supported.
        """
        return pulumi.get(self, "instance_storage_supported")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="ipv6Supported")
    def ipv6_supported(self) -> bool:
        """
        `true` if IPv6 is supported.
        """
        return pulumi.get(self, "ipv6_supported")

    @property
    @pulumi.getter(name="maximumIpv4AddressesPerInterface")
    def maximum_ipv4_addresses_per_interface(self) -> int:
        """
        The maximum number of IPv4 addresses per network interface.
        """
        return pulumi.get(self, "maximum_ipv4_addresses_per_interface")

    @property
    @pulumi.getter(name="maximumIpv6AddressesPerInterface")
    def maximum_ipv6_addresses_per_interface(self) -> int:
        """
        The maximum number of IPv6 addresses per network interface.
        """
        return pulumi.get(self, "maximum_ipv6_addresses_per_interface")

    @property
    @pulumi.getter(name="maximumNetworkInterfaces")
    def maximum_network_interfaces(self) -> int:
        """
        The maximum number of network interfaces for the instance type.
        """
        return pulumi.get(self, "maximum_network_interfaces")

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> int:
        """
        Size of the instance memory, in MiB.
        """
        return pulumi.get(self, "memory_size")

    @property
    @pulumi.getter(name="networkPerformance")
    def network_performance(self) -> str:
        """
        Describes the network performance.
        """
        return pulumi.get(self, "network_performance")

    @property
    @pulumi.getter(name="supportedArchitectures")
    def supported_architectures(self) -> Sequence[str]:
        """
        A list of architectures supported by the instance type.
        """
        return pulumi.get(self, "supported_architectures")

    @property
    @pulumi.getter(name="supportedPlacementStrategies")
    def supported_placement_strategies(self) -> Sequence[str]:
        """
        A list of supported placement groups types.
        """
        return pulumi.get(self, "supported_placement_strategies")

    @property
    @pulumi.getter(name="supportedRootDeviceTypes")
    def supported_root_device_types(self) -> Sequence[str]:
        """
        Indicates the supported root device types.
        """
        return pulumi.get(self, "supported_root_device_types")

    @property
    @pulumi.getter(name="supportedUsagesClasses")
    def supported_usages_classes(self) -> Sequence[str]:
        """
        Indicates whether the instance type is offered for spot or On-Demand.
        """
        return pulumi.get(self, "supported_usages_classes")

    @property
    @pulumi.getter(name="supportedVirtualizationTypes")
    def supported_virtualization_types(self) -> Sequence[str]:
        """
        The supported virtualization types.
        """
        return pulumi.get(self, "supported_virtualization_types")

    @property
    @pulumi.getter(name="sustainedClockSpeed")
    def sustained_clock_speed(self) -> float:
        """
        The speed of the processor, in GHz.
        """
        return pulumi.get(self, "sustained_clock_speed")

    @property
    @pulumi.getter(name="totalFpgaMemory")
    def total_fpga_memory(self) -> int:
        """
        The total memory of all FPGA accelerators for the instance type (in MiB).
        """
        return pulumi.get(self, "total_fpga_memory")

    @property
    @pulumi.getter(name="totalGpuMemory")
    def total_gpu_memory(self) -> int:
        """
        The total size of the memory for the GPU accelerators for the instance type (in MiB).
        """
        return pulumi.get(self, "total_gpu_memory")

    @property
    @pulumi.getter(name="totalInstanceStorage")
    def total_instance_storage(self) -> int:
        """
        The total size of the instance disks, in GB.
        """
        return pulumi.get(self, "total_instance_storage")

    @property
    @pulumi.getter(name="validCores")
    def valid_cores(self) -> Sequence[int]:
        """
        List of the valid number of cores that can be configured for the instance type.
        """
        return pulumi.get(self, "valid_cores")

    @property
    @pulumi.getter(name="validThreadsPerCores")
    def valid_threads_per_cores(self) -> Sequence[int]:
        """
        List of the valid number of threads per core that can be configured for the instance type.
        """
        return pulumi.get(self, "valid_threads_per_cores")


class AwaitableGetInstanceTypeResult(GetInstanceTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceTypeResult(
            auto_recovery_supported=self.auto_recovery_supported,
            bare_metal=self.bare_metal,
            burstable_performance_supported=self.burstable_performance_supported,
            current_generation=self.current_generation,
            dedicated_hosts_supported=self.dedicated_hosts_supported,
            default_cores=self.default_cores,
            default_threads_per_core=self.default_threads_per_core,
            default_vcpus=self.default_vcpus,
            ebs_encryption_support=self.ebs_encryption_support,
            ebs_nvme_support=self.ebs_nvme_support,
            ebs_optimized_support=self.ebs_optimized_support,
            ebs_performance_baseline_bandwidth=self.ebs_performance_baseline_bandwidth,
            ebs_performance_baseline_iops=self.ebs_performance_baseline_iops,
            ebs_performance_baseline_throughput=self.ebs_performance_baseline_throughput,
            ebs_performance_maximum_bandwidth=self.ebs_performance_maximum_bandwidth,
            ebs_performance_maximum_iops=self.ebs_performance_maximum_iops,
            ebs_performance_maximum_throughput=self.ebs_performance_maximum_throughput,
            efa_supported=self.efa_supported,
            ena_support=self.ena_support,
            fpgas=self.fpgas,
            free_tier_eligible=self.free_tier_eligible,
            gpuses=self.gpuses,
            hibernation_supported=self.hibernation_supported,
            hypervisor=self.hypervisor,
            id=self.id,
            inference_accelerators=self.inference_accelerators,
            instance_disks=self.instance_disks,
            instance_storage_supported=self.instance_storage_supported,
            instance_type=self.instance_type,
            ipv6_supported=self.ipv6_supported,
            maximum_ipv4_addresses_per_interface=self.maximum_ipv4_addresses_per_interface,
            maximum_ipv6_addresses_per_interface=self.maximum_ipv6_addresses_per_interface,
            maximum_network_interfaces=self.maximum_network_interfaces,
            memory_size=self.memory_size,
            network_performance=self.network_performance,
            supported_architectures=self.supported_architectures,
            supported_placement_strategies=self.supported_placement_strategies,
            supported_root_device_types=self.supported_root_device_types,
            supported_usages_classes=self.supported_usages_classes,
            supported_virtualization_types=self.supported_virtualization_types,
            sustained_clock_speed=self.sustained_clock_speed,
            total_fpga_memory=self.total_fpga_memory,
            total_gpu_memory=self.total_gpu_memory,
            total_instance_storage=self.total_instance_storage,
            valid_cores=self.valid_cores,
            valid_threads_per_cores=self.valid_threads_per_cores)


def get_instance_type(default_cores: Optional[int] = None,
                      default_threads_per_core: Optional[int] = None,
                      fpgas: Optional[Sequence[pulumi.InputType['GetInstanceTypeFpgaArgs']]] = None,
                      gpuses: Optional[Sequence[pulumi.InputType['GetInstanceTypeGpusArgs']]] = None,
                      hypervisor: Optional[str] = None,
                      inference_accelerators: Optional[Sequence[pulumi.InputType['GetInstanceTypeInferenceAcceleratorArgs']]] = None,
                      instance_disks: Optional[Sequence[pulumi.InputType['GetInstanceTypeInstanceDiskArgs']]] = None,
                      instance_type: Optional[str] = None,
                      maximum_ipv6_addresses_per_interface: Optional[int] = None,
                      total_fpga_memory: Optional[int] = None,
                      total_gpu_memory: Optional[int] = None,
                      total_instance_storage: Optional[int] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceTypeResult:
    """
    Get characteristics for a single EC2 Instance Type.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2.get_instance_type(instance_type="t2.micro")
    ```


    :param int default_cores: The default number of cores for the instance type.
    :param int default_threads_per_core: The  default  number of threads per core for the instance type.
    :param Sequence[pulumi.InputType['GetInstanceTypeFpgaArgs']] fpgas: Describes the FPGA accelerator settings for the instance type.
           * `fpgas.#.count` - The count of FPGA accelerators for the instance type.
           * `fpgas.#.manufacturer` - The manufacturer of the FPGA accelerator.
           * `fpgas.#.memory_size` - The size (in MiB) for the memory available to the FPGA accelerator.
           * `fpgas.#.name` - The name of the FPGA accelerator.
    :param Sequence[pulumi.InputType['GetInstanceTypeGpusArgs']] gpuses: Describes the GPU accelerators for the instance type.
           * `gpus.#.count` - The number of GPUs for the instance type.
           * `gpus.#.manufacturer` - The manufacturer of the GPU accelerator.
           * `gpus.#.memory_size` - The size (in MiB) for the memory available to the GPU accelerator.
           * `gpus.#.name` - The name of the GPU accelerator.
    :param str hypervisor: Indicates the hypervisor used for the instance type.
           * `inference_accelerators` Describes the Inference accelerators for the instance type.
           * `inference_accelerators.#.count` - The number of Inference accelerators for the instance type.
           * `inference_accelerators.#.manufacturer` - The manufacturer of the Inference accelerator.
           * `inference_accelerators.#.name` - The name of the Inference accelerator.
    :param Sequence[pulumi.InputType['GetInstanceTypeInstanceDiskArgs']] instance_disks: Describes the disks for the instance type.
           * `instance_disks.#.count` - The number of disks with this configuration.
           * `instance_disks.#.size` - The size of the disk in GB.
           * `instance_disks.#.type` - The type of disk.
    :param str instance_type: Instance
    :param int maximum_ipv6_addresses_per_interface: The maximum number of IPv6 addresses per network interface.
    :param int total_fpga_memory: The total memory of all FPGA accelerators for the instance type (in MiB).
    :param int total_gpu_memory: The total size of the memory for the GPU accelerators for the instance type (in MiB).
    :param int total_instance_storage: The total size of the instance disks, in GB.
    """
    __args__ = dict()
    __args__['defaultCores'] = default_cores
    __args__['defaultThreadsPerCore'] = default_threads_per_core
    __args__['fpgas'] = fpgas
    __args__['gpuses'] = gpuses
    __args__['hypervisor'] = hypervisor
    __args__['inferenceAccelerators'] = inference_accelerators
    __args__['instanceDisks'] = instance_disks
    __args__['instanceType'] = instance_type
    __args__['maximumIpv6AddressesPerInterface'] = maximum_ipv6_addresses_per_interface
    __args__['totalFpgaMemory'] = total_fpga_memory
    __args__['totalGpuMemory'] = total_gpu_memory
    __args__['totalInstanceStorage'] = total_instance_storage
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getInstanceType:getInstanceType', __args__, opts=opts, typ=GetInstanceTypeResult).value

    return AwaitableGetInstanceTypeResult(
        auto_recovery_supported=__ret__.auto_recovery_supported,
        bare_metal=__ret__.bare_metal,
        burstable_performance_supported=__ret__.burstable_performance_supported,
        current_generation=__ret__.current_generation,
        dedicated_hosts_supported=__ret__.dedicated_hosts_supported,
        default_cores=__ret__.default_cores,
        default_threads_per_core=__ret__.default_threads_per_core,
        default_vcpus=__ret__.default_vcpus,
        ebs_encryption_support=__ret__.ebs_encryption_support,
        ebs_nvme_support=__ret__.ebs_nvme_support,
        ebs_optimized_support=__ret__.ebs_optimized_support,
        ebs_performance_baseline_bandwidth=__ret__.ebs_performance_baseline_bandwidth,
        ebs_performance_baseline_iops=__ret__.ebs_performance_baseline_iops,
        ebs_performance_baseline_throughput=__ret__.ebs_performance_baseline_throughput,
        ebs_performance_maximum_bandwidth=__ret__.ebs_performance_maximum_bandwidth,
        ebs_performance_maximum_iops=__ret__.ebs_performance_maximum_iops,
        ebs_performance_maximum_throughput=__ret__.ebs_performance_maximum_throughput,
        efa_supported=__ret__.efa_supported,
        ena_support=__ret__.ena_support,
        fpgas=__ret__.fpgas,
        free_tier_eligible=__ret__.free_tier_eligible,
        gpuses=__ret__.gpuses,
        hibernation_supported=__ret__.hibernation_supported,
        hypervisor=__ret__.hypervisor,
        id=__ret__.id,
        inference_accelerators=__ret__.inference_accelerators,
        instance_disks=__ret__.instance_disks,
        instance_storage_supported=__ret__.instance_storage_supported,
        instance_type=__ret__.instance_type,
        ipv6_supported=__ret__.ipv6_supported,
        maximum_ipv4_addresses_per_interface=__ret__.maximum_ipv4_addresses_per_interface,
        maximum_ipv6_addresses_per_interface=__ret__.maximum_ipv6_addresses_per_interface,
        maximum_network_interfaces=__ret__.maximum_network_interfaces,
        memory_size=__ret__.memory_size,
        network_performance=__ret__.network_performance,
        supported_architectures=__ret__.supported_architectures,
        supported_placement_strategies=__ret__.supported_placement_strategies,
        supported_root_device_types=__ret__.supported_root_device_types,
        supported_usages_classes=__ret__.supported_usages_classes,
        supported_virtualization_types=__ret__.supported_virtualization_types,
        sustained_clock_speed=__ret__.sustained_clock_speed,
        total_fpga_memory=__ret__.total_fpga_memory,
        total_gpu_memory=__ret__.total_gpu_memory,
        total_instance_storage=__ret__.total_instance_storage,
        valid_cores=__ret__.valid_cores,
        valid_threads_per_cores=__ret__.valid_threads_per_cores)

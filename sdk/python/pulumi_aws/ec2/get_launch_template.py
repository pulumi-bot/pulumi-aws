# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetLaunchTemplateResult',
    'AwaitableGetLaunchTemplateResult',
    'get_launch_template',
]

@pulumi.output_type
class GetLaunchTemplateResult:
    """
    A collection of values returned by getLaunchTemplate.
    """
    def __init__(__self__, arn=None, block_device_mappings=None, credit_specifications=None, default_version=None, description=None, disable_api_termination=None, ebs_optimized=None, elastic_gpu_specifications=None, filters=None, hibernation_options=None, iam_instance_profiles=None, id=None, image_id=None, instance_initiated_shutdown_behavior=None, instance_market_options=None, instance_type=None, kernel_id=None, key_name=None, latest_version=None, metadata_options=None, monitorings=None, name=None, network_interfaces=None, placements=None, ram_disk_id=None, security_group_names=None, tag_specifications=None, tags=None, user_data=None, vpc_security_group_ids=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if block_device_mappings and not isinstance(block_device_mappings, list):
            raise TypeError("Expected argument 'block_device_mappings' to be a list")
        pulumi.set(__self__, "block_device_mappings", block_device_mappings)
        if credit_specifications and not isinstance(credit_specifications, list):
            raise TypeError("Expected argument 'credit_specifications' to be a list")
        pulumi.set(__self__, "credit_specifications", credit_specifications)
        if default_version and not isinstance(default_version, float):
            raise TypeError("Expected argument 'default_version' to be a float")
        pulumi.set(__self__, "default_version", default_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disable_api_termination and not isinstance(disable_api_termination, bool):
            raise TypeError("Expected argument 'disable_api_termination' to be a bool")
        pulumi.set(__self__, "disable_api_termination", disable_api_termination)
        if ebs_optimized and not isinstance(ebs_optimized, str):
            raise TypeError("Expected argument 'ebs_optimized' to be a str")
        pulumi.set(__self__, "ebs_optimized", ebs_optimized)
        if elastic_gpu_specifications and not isinstance(elastic_gpu_specifications, list):
            raise TypeError("Expected argument 'elastic_gpu_specifications' to be a list")
        pulumi.set(__self__, "elastic_gpu_specifications", elastic_gpu_specifications)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if hibernation_options and not isinstance(hibernation_options, list):
            raise TypeError("Expected argument 'hibernation_options' to be a list")
        pulumi.set(__self__, "hibernation_options", hibernation_options)
        if iam_instance_profiles and not isinstance(iam_instance_profiles, list):
            raise TypeError("Expected argument 'iam_instance_profiles' to be a list")
        pulumi.set(__self__, "iam_instance_profiles", iam_instance_profiles)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if instance_initiated_shutdown_behavior and not isinstance(instance_initiated_shutdown_behavior, str):
            raise TypeError("Expected argument 'instance_initiated_shutdown_behavior' to be a str")
        pulumi.set(__self__, "instance_initiated_shutdown_behavior", instance_initiated_shutdown_behavior)
        if instance_market_options and not isinstance(instance_market_options, list):
            raise TypeError("Expected argument 'instance_market_options' to be a list")
        pulumi.set(__self__, "instance_market_options", instance_market_options)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if kernel_id and not isinstance(kernel_id, str):
            raise TypeError("Expected argument 'kernel_id' to be a str")
        pulumi.set(__self__, "kernel_id", kernel_id)
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        pulumi.set(__self__, "key_name", key_name)
        if latest_version and not isinstance(latest_version, float):
            raise TypeError("Expected argument 'latest_version' to be a float")
        pulumi.set(__self__, "latest_version", latest_version)
        if metadata_options and not isinstance(metadata_options, list):
            raise TypeError("Expected argument 'metadata_options' to be a list")
        pulumi.set(__self__, "metadata_options", metadata_options)
        if monitorings and not isinstance(monitorings, list):
            raise TypeError("Expected argument 'monitorings' to be a list")
        pulumi.set(__self__, "monitorings", monitorings)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        pulumi.set(__self__, "network_interfaces", network_interfaces)
        if placements and not isinstance(placements, list):
            raise TypeError("Expected argument 'placements' to be a list")
        pulumi.set(__self__, "placements", placements)
        if ram_disk_id and not isinstance(ram_disk_id, str):
            raise TypeError("Expected argument 'ram_disk_id' to be a str")
        pulumi.set(__self__, "ram_disk_id", ram_disk_id)
        if security_group_names and not isinstance(security_group_names, list):
            raise TypeError("Expected argument 'security_group_names' to be a list")
        pulumi.set(__self__, "security_group_names", security_group_names)
        if tag_specifications and not isinstance(tag_specifications, list):
            raise TypeError("Expected argument 'tag_specifications' to be a list")
        pulumi.set(__self__, "tag_specifications", tag_specifications)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)
        if vpc_security_group_ids and not isinstance(vpc_security_group_ids, list):
            raise TypeError("Expected argument 'vpc_security_group_ids' to be a list")
        pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the launch template.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="blockDeviceMappings")
    def block_device_mappings(self) -> List['outputs.GetLaunchTemplateBlockDeviceMappingResult']:
        """
        Specify volumes to attach to the instance besides the volumes specified by the AMI.
        """
        return pulumi.get(self, "block_device_mappings")

    @property
    @pulumi.getter(name="creditSpecifications")
    def credit_specifications(self) -> List['outputs.GetLaunchTemplateCreditSpecificationResult']:
        """
        Customize the credit specification of the instance. See Credit
        Specification below for more details.
        """
        return pulumi.get(self, "credit_specifications")

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> float:
        """
        The default version of the launch template.
        """
        return pulumi.get(self, "default_version")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the launch template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableApiTermination")
    def disable_api_termination(self) -> bool:
        """
        If `true`, enables [EC2 Instance
        Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
        """
        return pulumi.get(self, "disable_api_termination")

    @property
    @pulumi.getter(name="ebsOptimized")
    def ebs_optimized(self) -> str:
        """
        If `true`, the launched EC2 instance will be EBS-optimized.
        """
        return pulumi.get(self, "ebs_optimized")

    @property
    @pulumi.getter(name="elasticGpuSpecifications")
    def elastic_gpu_specifications(self) -> List['outputs.GetLaunchTemplateElasticGpuSpecificationResult']:
        """
        The elastic GPU to attach to the instance. See Elastic GPU
        below for more details.
        """
        return pulumi.get(self, "elastic_gpu_specifications")

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetLaunchTemplateFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="hibernationOptions")
    def hibernation_options(self) -> List['outputs.GetLaunchTemplateHibernationOptionResult']:
        """
        The hibernation options for the instance.
        """
        return pulumi.get(self, "hibernation_options")

    @property
    @pulumi.getter(name="iamInstanceProfiles")
    def iam_instance_profiles(self) -> List['outputs.GetLaunchTemplateIamInstanceProfileResult']:
        """
        The IAM Instance Profile to launch the instance with. See Instance Profile
        below for more details.
        """
        return pulumi.get(self, "iam_instance_profiles")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The AMI from which to launch the instance.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceInitiatedShutdownBehavior")
    def instance_initiated_shutdown_behavior(self) -> str:
        """
        Shutdown behavior for the instance. Can be `stop` or `terminate`.
        (Default: `stop`).
        """
        return pulumi.get(self, "instance_initiated_shutdown_behavior")

    @property
    @pulumi.getter(name="instanceMarketOptions")
    def instance_market_options(self) -> List['outputs.GetLaunchTemplateInstanceMarketOptionResult']:
        """
        The market (purchasing) option for the instance.
        below for details.
        """
        return pulumi.get(self, "instance_market_options")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The type of the instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="kernelId")
    def kernel_id(self) -> str:
        """
        The kernel ID.
        """
        return pulumi.get(self, "kernel_id")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        The key name to use for the instance.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="latestVersion")
    def latest_version(self) -> float:
        """
        The latest version of the launch template.
        """
        return pulumi.get(self, "latest_version")

    @property
    @pulumi.getter(name="metadataOptions")
    def metadata_options(self) -> List['outputs.GetLaunchTemplateMetadataOptionResult']:
        """
        The metadata options for the instance.
        """
        return pulumi.get(self, "metadata_options")

    @property
    @pulumi.getter
    def monitorings(self) -> List['outputs.GetLaunchTemplateMonitoringResult']:
        """
        The monitoring option for the instance.
        """
        return pulumi.get(self, "monitorings")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> List['outputs.GetLaunchTemplateNetworkInterfaceResult']:
        """
        Customize network interfaces to be attached at instance boot time. See Network
        Interfaces below for more details.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter
    def placements(self) -> List['outputs.GetLaunchTemplatePlacementResult']:
        """
        The placement of the instance.
        """
        return pulumi.get(self, "placements")

    @property
    @pulumi.getter(name="ramDiskId")
    def ram_disk_id(self) -> str:
        """
        The ID of the RAM disk.
        """
        return pulumi.get(self, "ram_disk_id")

    @property
    @pulumi.getter(name="securityGroupNames")
    def security_group_names(self) -> List[str]:
        """
        A list of security group names to associate with. If you are creating Instances in a VPC, use
        `vpc_security_group_ids` instead.
        """
        return pulumi.get(self, "security_group_names")

    @property
    @pulumi.getter(name="tagSpecifications")
    def tag_specifications(self) -> List['outputs.GetLaunchTemplateTagSpecificationResult']:
        """
        The tags to apply to the resources during launch.
        """
        return pulumi.get(self, "tag_specifications")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        (Optional) A map of tags to assign to the launch template.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> str:
        """
        The Base64-encoded user data to provide when launching the instance.
        """
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> List[str]:
        """
        A list of security group IDs to associate with.
        """
        return pulumi.get(self, "vpc_security_group_ids")


class AwaitableGetLaunchTemplateResult(GetLaunchTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLaunchTemplateResult(
            arn=self.arn,
            block_device_mappings=self.block_device_mappings,
            credit_specifications=self.credit_specifications,
            default_version=self.default_version,
            description=self.description,
            disable_api_termination=self.disable_api_termination,
            ebs_optimized=self.ebs_optimized,
            elastic_gpu_specifications=self.elastic_gpu_specifications,
            filters=self.filters,
            hibernation_options=self.hibernation_options,
            iam_instance_profiles=self.iam_instance_profiles,
            id=self.id,
            image_id=self.image_id,
            instance_initiated_shutdown_behavior=self.instance_initiated_shutdown_behavior,
            instance_market_options=self.instance_market_options,
            instance_type=self.instance_type,
            kernel_id=self.kernel_id,
            key_name=self.key_name,
            latest_version=self.latest_version,
            metadata_options=self.metadata_options,
            monitorings=self.monitorings,
            name=self.name,
            network_interfaces=self.network_interfaces,
            placements=self.placements,
            ram_disk_id=self.ram_disk_id,
            security_group_names=self.security_group_names,
            tag_specifications=self.tag_specifications,
            tags=self.tags,
            user_data=self.user_data,
            vpc_security_group_ids=self.vpc_security_group_ids)


def get_launch_template(filters: Optional[List[pulumi.InputType['GetLaunchTemplateFilterArgs']]] = None,
                        name: Optional[str] = None,
                        tags: Optional[Mapping[str, str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLaunchTemplateResult:
    """
    Provides information about a Launch Template.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    default = aws.ec2.get_launch_template(name="my-launch-template")
    ```
    ### Filter

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.ec2.get_launch_template(filters=[{
        "name": "launch-template-name",
        "values": ["some-template"],
    }])
    ```


    :param List[pulumi.InputType['GetLaunchTemplateFilterArgs']] filters: Configuration block(s) for filtering. Detailed below.
    :param str name: The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match a pair on the desired Launch Template.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLaunchTemplate:getLaunchTemplate', __args__, opts=opts, typ=GetLaunchTemplateResult).value

    return AwaitableGetLaunchTemplateResult(
        arn=__ret__.arn,
        block_device_mappings=__ret__.block_device_mappings,
        credit_specifications=__ret__.credit_specifications,
        default_version=__ret__.default_version,
        description=__ret__.description,
        disable_api_termination=__ret__.disable_api_termination,
        ebs_optimized=__ret__.ebs_optimized,
        elastic_gpu_specifications=__ret__.elastic_gpu_specifications,
        filters=__ret__.filters,
        hibernation_options=__ret__.hibernation_options,
        iam_instance_profiles=__ret__.iam_instance_profiles,
        id=__ret__.id,
        image_id=__ret__.image_id,
        instance_initiated_shutdown_behavior=__ret__.instance_initiated_shutdown_behavior,
        instance_market_options=__ret__.instance_market_options,
        instance_type=__ret__.instance_type,
        kernel_id=__ret__.kernel_id,
        key_name=__ret__.key_name,
        latest_version=__ret__.latest_version,
        metadata_options=__ret__.metadata_options,
        monitorings=__ret__.monitorings,
        name=__ret__.name,
        network_interfaces=__ret__.network_interfaces,
        placements=__ret__.placements,
        ram_disk_id=__ret__.ram_disk_id,
        security_group_names=__ret__.security_group_names,
        tag_specifications=__ret__.tag_specifications,
        tags=__ret__.tags,
        user_data=__ret__.user_data,
        vpc_security_group_ids=__ret__.vpc_security_group_ids)

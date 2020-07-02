# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetLaunchTemplateResult:
    """
    A collection of values returned by getLaunchTemplate.
    """
    def __init__(__self__, arn=None, block_device_mappings=None, credit_specifications=None, default_version=None, description=None, disable_api_termination=None, ebs_optimized=None, elastic_gpu_specifications=None, filters=None, hibernation_options=None, iam_instance_profiles=None, id=None, image_id=None, instance_initiated_shutdown_behavior=None, instance_market_options=None, instance_type=None, kernel_id=None, key_name=None, latest_version=None, metadata_options=None, monitorings=None, name=None, network_interfaces=None, placements=None, ram_disk_id=None, security_group_names=None, tag_specifications=None, tags=None, user_data=None, vpc_security_group_ids=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Amazon Resource Name (ARN) of the launch template.
        """
        if block_device_mappings and not isinstance(block_device_mappings, list):
            raise TypeError("Expected argument 'block_device_mappings' to be a list")
        __self__.block_device_mappings = block_device_mappings
        """
        Specify volumes to attach to the instance besides the volumes specified by the AMI.
        """
        if credit_specifications and not isinstance(credit_specifications, list):
            raise TypeError("Expected argument 'credit_specifications' to be a list")
        __self__.credit_specifications = credit_specifications
        """
        Customize the credit specification of the instance. See Credit
        Specification below for more details.
        """
        if default_version and not isinstance(default_version, float):
            raise TypeError("Expected argument 'default_version' to be a float")
        __self__.default_version = default_version
        """
        The default version of the launch template.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of the launch template.
        """
        if disable_api_termination and not isinstance(disable_api_termination, bool):
            raise TypeError("Expected argument 'disable_api_termination' to be a bool")
        __self__.disable_api_termination = disable_api_termination
        """
        If `true`, enables [EC2 Instance
        Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
        """
        if ebs_optimized and not isinstance(ebs_optimized, str):
            raise TypeError("Expected argument 'ebs_optimized' to be a str")
        __self__.ebs_optimized = ebs_optimized
        """
        If `true`, the launched EC2 instance will be EBS-optimized.
        """
        if elastic_gpu_specifications and not isinstance(elastic_gpu_specifications, list):
            raise TypeError("Expected argument 'elastic_gpu_specifications' to be a list")
        __self__.elastic_gpu_specifications = elastic_gpu_specifications
        """
        The elastic GPU to attach to the instance. See Elastic GPU
        below for more details.
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if hibernation_options and not isinstance(hibernation_options, list):
            raise TypeError("Expected argument 'hibernation_options' to be a list")
        __self__.hibernation_options = hibernation_options
        """
        The hibernation options for the instance.
        """
        if iam_instance_profiles and not isinstance(iam_instance_profiles, list):
            raise TypeError("Expected argument 'iam_instance_profiles' to be a list")
        __self__.iam_instance_profiles = iam_instance_profiles
        """
        The IAM Instance Profile to launch the instance with. See Instance Profile
        below for more details.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        __self__.image_id = image_id
        """
        The AMI from which to launch the instance.
        """
        if instance_initiated_shutdown_behavior and not isinstance(instance_initiated_shutdown_behavior, str):
            raise TypeError("Expected argument 'instance_initiated_shutdown_behavior' to be a str")
        __self__.instance_initiated_shutdown_behavior = instance_initiated_shutdown_behavior
        """
        Shutdown behavior for the instance. Can be `stop` or `terminate`.
        (Default: `stop`).
        """
        if instance_market_options and not isinstance(instance_market_options, list):
            raise TypeError("Expected argument 'instance_market_options' to be a list")
        __self__.instance_market_options = instance_market_options
        """
        The market (purchasing) option for the instance.
        below for details.
        """
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        __self__.instance_type = instance_type
        """
        The type of the instance.
        """
        if kernel_id and not isinstance(kernel_id, str):
            raise TypeError("Expected argument 'kernel_id' to be a str")
        __self__.kernel_id = kernel_id
        """
        The kernel ID.
        """
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        __self__.key_name = key_name
        """
        The key name to use for the instance.
        """
        if latest_version and not isinstance(latest_version, float):
            raise TypeError("Expected argument 'latest_version' to be a float")
        __self__.latest_version = latest_version
        """
        The latest version of the launch template.
        """
        if metadata_options and not isinstance(metadata_options, list):
            raise TypeError("Expected argument 'metadata_options' to be a list")
        __self__.metadata_options = metadata_options
        """
        The metadata options for the instance.
        """
        if monitorings and not isinstance(monitorings, list):
            raise TypeError("Expected argument 'monitorings' to be a list")
        __self__.monitorings = monitorings
        """
        The monitoring option for the instance.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        __self__.network_interfaces = network_interfaces
        """
        Customize network interfaces to be attached at instance boot time. See Network
        Interfaces below for more details.
        """
        if placements and not isinstance(placements, list):
            raise TypeError("Expected argument 'placements' to be a list")
        __self__.placements = placements
        """
        The placement of the instance.
        """
        if ram_disk_id and not isinstance(ram_disk_id, str):
            raise TypeError("Expected argument 'ram_disk_id' to be a str")
        __self__.ram_disk_id = ram_disk_id
        """
        The ID of the RAM disk.
        """
        if security_group_names and not isinstance(security_group_names, list):
            raise TypeError("Expected argument 'security_group_names' to be a list")
        __self__.security_group_names = security_group_names
        """
        A list of security group names to associate with. If you are creating Instances in a VPC, use
        `vpc_security_group_ids` instead.
        """
        if tag_specifications and not isinstance(tag_specifications, list):
            raise TypeError("Expected argument 'tag_specifications' to be a list")
        __self__.tag_specifications = tag_specifications
        """
        The tags to apply to the resources during launch.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        (Optional) A map of tags to assign to the launch template.
        """
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        __self__.user_data = user_data
        """
        The Base64-encoded user data to provide when launching the instance.
        """
        if vpc_security_group_ids and not isinstance(vpc_security_group_ids, list):
            raise TypeError("Expected argument 'vpc_security_group_ids' to be a list")
        __self__.vpc_security_group_ids = vpc_security_group_ids
        """
        A list of security group IDs to associate with.
        """


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


def get_launch_template(filters=None,name=None,tags=None,opts=None):
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


    :param list filters: Configuration block(s) for filtering. Detailed below.
    :param str name: The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
    :param dict tags: A map of tags, each pair of which must exactly match a pair on the desired Launch Template.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
      * `values` (`list`) - Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLaunchTemplate:getLaunchTemplate', __args__, opts=opts).value

    return AwaitableGetLaunchTemplateResult(
        arn=__ret__.get('arn'),
        block_device_mappings=__ret__.get('blockDeviceMappings'),
        credit_specifications=__ret__.get('creditSpecifications'),
        default_version=__ret__.get('defaultVersion'),
        description=__ret__.get('description'),
        disable_api_termination=__ret__.get('disableApiTermination'),
        ebs_optimized=__ret__.get('ebsOptimized'),
        elastic_gpu_specifications=__ret__.get('elasticGpuSpecifications'),
        filters=__ret__.get('filters'),
        hibernation_options=__ret__.get('hibernationOptions'),
        iam_instance_profiles=__ret__.get('iamInstanceProfiles'),
        id=__ret__.get('id'),
        image_id=__ret__.get('imageId'),
        instance_initiated_shutdown_behavior=__ret__.get('instanceInitiatedShutdownBehavior'),
        instance_market_options=__ret__.get('instanceMarketOptions'),
        instance_type=__ret__.get('instanceType'),
        kernel_id=__ret__.get('kernelId'),
        key_name=__ret__.get('keyName'),
        latest_version=__ret__.get('latestVersion'),
        metadata_options=__ret__.get('metadataOptions'),
        monitorings=__ret__.get('monitorings'),
        name=__ret__.get('name'),
        network_interfaces=__ret__.get('networkInterfaces'),
        placements=__ret__.get('placements'),
        ram_disk_id=__ret__.get('ramDiskId'),
        security_group_names=__ret__.get('securityGroupNames'),
        tag_specifications=__ret__.get('tagSpecifications'),
        tags=__ret__.get('tags'),
        user_data=__ret__.get('userData'),
        vpc_security_group_ids=__ret__.get('vpcSecurityGroupIds'))

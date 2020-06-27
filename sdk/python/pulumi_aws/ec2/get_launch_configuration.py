# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLaunchConfigurationResult:
    """
    A collection of values returned by getLaunchConfiguration.
    """
    def __init__(__self__, arn=None, associate_public_ip_address=None, ebs_block_devices=None, ebs_optimized=None, enable_monitoring=None, ephemeral_block_devices=None, iam_instance_profile=None, id=None, image_id=None, instance_type=None, key_name=None, name=None, placement_tenancy=None, root_block_devices=None, security_groups=None, spot_price=None, user_data=None, vpc_classic_link_id=None, vpc_classic_link_security_groups=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name of the launch configuration.
        """
        if associate_public_ip_address and not isinstance(associate_public_ip_address, bool):
            raise TypeError("Expected argument 'associate_public_ip_address' to be a bool")
        __self__.associate_public_ip_address = associate_public_ip_address
        """
        Whether a Public IP address is associated with the instance.
        """
        if ebs_block_devices and not isinstance(ebs_block_devices, list):
            raise TypeError("Expected argument 'ebs_block_devices' to be a list")
        __self__.ebs_block_devices = ebs_block_devices
        """
        The EBS Block Devices attached to the instance.
        """
        if ebs_optimized and not isinstance(ebs_optimized, bool):
            raise TypeError("Expected argument 'ebs_optimized' to be a bool")
        __self__.ebs_optimized = ebs_optimized
        """
        Whether the launched EC2 instance will be EBS-optimized.
        """
        if enable_monitoring and not isinstance(enable_monitoring, bool):
            raise TypeError("Expected argument 'enable_monitoring' to be a bool")
        __self__.enable_monitoring = enable_monitoring
        """
        Whether Detailed Monitoring is Enabled.
        """
        if ephemeral_block_devices and not isinstance(ephemeral_block_devices, list):
            raise TypeError("Expected argument 'ephemeral_block_devices' to be a list")
        __self__.ephemeral_block_devices = ephemeral_block_devices
        """
        The Ephemeral volumes on the instance.
        """
        if iam_instance_profile and not isinstance(iam_instance_profile, str):
            raise TypeError("Expected argument 'iam_instance_profile' to be a str")
        __self__.iam_instance_profile = iam_instance_profile
        """
        The IAM Instance Profile to associate with launched instances.
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
        The EC2 Image ID of the instance.
        """
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        __self__.instance_type = instance_type
        """
        The Instance Type of the instance to launch.
        """
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        __self__.key_name = key_name
        """
        The Key Name that should be used for the instance.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The Name of the launch configuration.
        """
        if placement_tenancy and not isinstance(placement_tenancy, str):
            raise TypeError("Expected argument 'placement_tenancy' to be a str")
        __self__.placement_tenancy = placement_tenancy
        """
        The Tenancy of the instance.
        """
        if root_block_devices and not isinstance(root_block_devices, list):
            raise TypeError("Expected argument 'root_block_devices' to be a list")
        __self__.root_block_devices = root_block_devices
        """
        The Root Block Device of the instance.
        """
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        __self__.security_groups = security_groups
        """
        A list of associated Security Group IDS.
        """
        if spot_price and not isinstance(spot_price, str):
            raise TypeError("Expected argument 'spot_price' to be a str")
        __self__.spot_price = spot_price
        """
        The Price to use for reserving Spot instances.
        """
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        __self__.user_data = user_data
        """
        The User Data of the instance.
        """
        if vpc_classic_link_id and not isinstance(vpc_classic_link_id, str):
            raise TypeError("Expected argument 'vpc_classic_link_id' to be a str")
        __self__.vpc_classic_link_id = vpc_classic_link_id
        """
        The ID of a ClassicLink-enabled VPC.
        """
        if vpc_classic_link_security_groups and not isinstance(vpc_classic_link_security_groups, list):
            raise TypeError("Expected argument 'vpc_classic_link_security_groups' to be a list")
        __self__.vpc_classic_link_security_groups = vpc_classic_link_security_groups
        """
        The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.
        """
class AwaitableGetLaunchConfigurationResult(GetLaunchConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLaunchConfigurationResult(
            arn=self.arn,
            associate_public_ip_address=self.associate_public_ip_address,
            ebs_block_devices=self.ebs_block_devices,
            ebs_optimized=self.ebs_optimized,
            enable_monitoring=self.enable_monitoring,
            ephemeral_block_devices=self.ephemeral_block_devices,
            iam_instance_profile=self.iam_instance_profile,
            id=self.id,
            image_id=self.image_id,
            instance_type=self.instance_type,
            key_name=self.key_name,
            name=self.name,
            placement_tenancy=self.placement_tenancy,
            root_block_devices=self.root_block_devices,
            security_groups=self.security_groups,
            spot_price=self.spot_price,
            user_data=self.user_data,
            vpc_classic_link_id=self.vpc_classic_link_id,
            vpc_classic_link_security_groups=self.vpc_classic_link_security_groups)

def get_launch_configuration(name=None,opts=None):
    """
    Provides information about a Launch Configuration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    ubuntu = aws.ec2.get_launch_configuration(name="test-launch-config")
    ```


    :param str name: The name of the launch configuration.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLaunchConfiguration:getLaunchConfiguration', __args__, opts=opts).value

    return AwaitableGetLaunchConfigurationResult(
        arn=__ret__.get('arn'),
        associate_public_ip_address=__ret__.get('associatePublicIpAddress'),
        ebs_block_devices=__ret__.get('ebsBlockDevices'),
        ebs_optimized=__ret__.get('ebsOptimized'),
        enable_monitoring=__ret__.get('enableMonitoring'),
        ephemeral_block_devices=__ret__.get('ephemeralBlockDevices'),
        iam_instance_profile=__ret__.get('iamInstanceProfile'),
        id=__ret__.get('id'),
        image_id=__ret__.get('imageId'),
        instance_type=__ret__.get('instanceType'),
        key_name=__ret__.get('keyName'),
        name=__ret__.get('name'),
        placement_tenancy=__ret__.get('placementTenancy'),
        root_block_devices=__ret__.get('rootBlockDevices'),
        security_groups=__ret__.get('securityGroups'),
        spot_price=__ret__.get('spotPrice'),
        user_data=__ret__.get('userData'),
        vpc_classic_link_id=__ret__.get('vpcClassicLinkId'),
        vpc_classic_link_security_groups=__ret__.get('vpcClassicLinkSecurityGroups'))

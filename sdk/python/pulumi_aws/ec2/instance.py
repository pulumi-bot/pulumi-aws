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

__all__ = ['Instance']


class Instance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ami: Optional[pulumi.Input[str]] = None,
                 associate_public_ip_address: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cpu_core_count: Optional[pulumi.Input[float]] = None,
                 cpu_threads_per_core: Optional[pulumi.Input[float]] = None,
                 credit_specification: Optional[pulumi.Input[pulumi.InputType['InstanceCreditSpecificationArgs']]] = None,
                 disable_api_termination: Optional[pulumi.Input[bool]] = None,
                 ebs_block_devices: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceEbsBlockDeviceArgs']]]]] = None,
                 ebs_optimized: Optional[pulumi.Input[bool]] = None,
                 ephemeral_block_devices: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceEphemeralBlockDeviceArgs']]]]] = None,
                 get_password_data: Optional[pulumi.Input[bool]] = None,
                 hibernation: Optional[pulumi.Input[bool]] = None,
                 host_id: Optional[pulumi.Input[str]] = None,
                 iam_instance_profile: Optional[pulumi.Input[str]] = None,
                 instance_initiated_shutdown_behavior: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 ipv6_address_count: Optional[pulumi.Input[float]] = None,
                 ipv6_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 metadata_options: Optional[pulumi.Input[pulumi.InputType['InstanceMetadataOptionsArgs']]] = None,
                 monitoring: Optional[pulumi.Input[bool]] = None,
                 network_interfaces: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceNetworkInterfaceArgs']]]]] = None,
                 placement_group: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 root_block_device: Optional[pulumi.Input[pulumi.InputType['InstanceRootBlockDeviceArgs']]] = None,
                 secondary_private_ips: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 security_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 source_dest_check: Optional[pulumi.Input[bool]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tenancy: Optional[pulumi.Input[str]] = None,
                 user_data: Optional[pulumi.Input[str]] = None,
                 user_data_base64: Optional[pulumi.Input[str]] = None,
                 volume_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Instance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if ami is None:
                raise TypeError("Missing required property 'ami'")
            __props__['ami'] = ami
            __props__['associate_public_ip_address'] = associate_public_ip_address
            __props__['availability_zone'] = availability_zone
            __props__['cpu_core_count'] = cpu_core_count
            __props__['cpu_threads_per_core'] = cpu_threads_per_core
            __props__['credit_specification'] = credit_specification
            __props__['disable_api_termination'] = disable_api_termination
            __props__['ebs_block_devices'] = ebs_block_devices
            __props__['ebs_optimized'] = ebs_optimized
            __props__['ephemeral_block_devices'] = ephemeral_block_devices
            __props__['get_password_data'] = get_password_data
            __props__['hibernation'] = hibernation
            __props__['host_id'] = host_id
            __props__['iam_instance_profile'] = iam_instance_profile
            __props__['instance_initiated_shutdown_behavior'] = instance_initiated_shutdown_behavior
            if instance_type is None:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            __props__['ipv6_address_count'] = ipv6_address_count
            __props__['ipv6_addresses'] = ipv6_addresses
            __props__['key_name'] = key_name
            __props__['metadata_options'] = metadata_options
            __props__['monitoring'] = monitoring
            __props__['network_interfaces'] = network_interfaces
            __props__['placement_group'] = placement_group
            __props__['private_ip'] = private_ip
            __props__['root_block_device'] = root_block_device
            __props__['secondary_private_ips'] = secondary_private_ips
            if security_groups is not None:
                warnings.warn("Use of `securityGroups` is discouraged as it does not allow for changes and will force your instance to be replaced if changes are made. To avoid this, use `vpcSecurityGroupIds` which allows for updates.", DeprecationWarning)
                pulumi.log.warn("security_groups is deprecated: Use of `securityGroups` is discouraged as it does not allow for changes and will force your instance to be replaced if changes are made. To avoid this, use `vpcSecurityGroupIds` which allows for updates.")
            __props__['security_groups'] = security_groups
            __props__['source_dest_check'] = source_dest_check
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['tenancy'] = tenancy
            __props__['user_data'] = user_data
            __props__['user_data_base64'] = user_data_base64
            __props__['volume_tags'] = volume_tags
            __props__['vpc_security_group_ids'] = vpc_security_group_ids
            __props__['arn'] = None
            __props__['instance_state'] = None
            __props__['outpost_arn'] = None
            __props__['password_data'] = None
            __props__['primary_network_interface_id'] = None
            __props__['private_dns'] = None
            __props__['public_dns'] = None
            __props__['public_ip'] = None
        super(Instance, __self__).__init__(
            'aws:ec2/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ami: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            associate_public_ip_address: Optional[pulumi.Input[bool]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            cpu_core_count: Optional[pulumi.Input[float]] = None,
            cpu_threads_per_core: Optional[pulumi.Input[float]] = None,
            credit_specification: Optional[pulumi.Input[pulumi.InputType['InstanceCreditSpecificationArgs']]] = None,
            disable_api_termination: Optional[pulumi.Input[bool]] = None,
            ebs_block_devices: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceEbsBlockDeviceArgs']]]]] = None,
            ebs_optimized: Optional[pulumi.Input[bool]] = None,
            ephemeral_block_devices: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceEphemeralBlockDeviceArgs']]]]] = None,
            get_password_data: Optional[pulumi.Input[bool]] = None,
            hibernation: Optional[pulumi.Input[bool]] = None,
            host_id: Optional[pulumi.Input[str]] = None,
            iam_instance_profile: Optional[pulumi.Input[str]] = None,
            instance_initiated_shutdown_behavior: Optional[pulumi.Input[str]] = None,
            instance_state: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            ipv6_address_count: Optional[pulumi.Input[float]] = None,
            ipv6_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            metadata_options: Optional[pulumi.Input[pulumi.InputType['InstanceMetadataOptionsArgs']]] = None,
            monitoring: Optional[pulumi.Input[bool]] = None,
            network_interfaces: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceNetworkInterfaceArgs']]]]] = None,
            outpost_arn: Optional[pulumi.Input[str]] = None,
            password_data: Optional[pulumi.Input[str]] = None,
            placement_group: Optional[pulumi.Input[str]] = None,
            primary_network_interface_id: Optional[pulumi.Input[str]] = None,
            private_dns: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            public_dns: Optional[pulumi.Input[str]] = None,
            public_ip: Optional[pulumi.Input[str]] = None,
            root_block_device: Optional[pulumi.Input[pulumi.InputType['InstanceRootBlockDeviceArgs']]] = None,
            secondary_private_ips: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            security_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            source_dest_check: Optional[pulumi.Input[bool]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tenancy: Optional[pulumi.Input[str]] = None,
            user_data: Optional[pulumi.Input[str]] = None,
            user_data_base64: Optional[pulumi.Input[str]] = None,
            volume_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ami"] = ami
        __props__["arn"] = arn
        __props__["associate_public_ip_address"] = associate_public_ip_address
        __props__["availability_zone"] = availability_zone
        __props__["cpu_core_count"] = cpu_core_count
        __props__["cpu_threads_per_core"] = cpu_threads_per_core
        __props__["credit_specification"] = credit_specification
        __props__["disable_api_termination"] = disable_api_termination
        __props__["ebs_block_devices"] = ebs_block_devices
        __props__["ebs_optimized"] = ebs_optimized
        __props__["ephemeral_block_devices"] = ephemeral_block_devices
        __props__["get_password_data"] = get_password_data
        __props__["hibernation"] = hibernation
        __props__["host_id"] = host_id
        __props__["iam_instance_profile"] = iam_instance_profile
        __props__["instance_initiated_shutdown_behavior"] = instance_initiated_shutdown_behavior
        __props__["instance_state"] = instance_state
        __props__["instance_type"] = instance_type
        __props__["ipv6_address_count"] = ipv6_address_count
        __props__["ipv6_addresses"] = ipv6_addresses
        __props__["key_name"] = key_name
        __props__["metadata_options"] = metadata_options
        __props__["monitoring"] = monitoring
        __props__["network_interfaces"] = network_interfaces
        __props__["outpost_arn"] = outpost_arn
        __props__["password_data"] = password_data
        __props__["placement_group"] = placement_group
        __props__["primary_network_interface_id"] = primary_network_interface_id
        __props__["private_dns"] = private_dns
        __props__["private_ip"] = private_ip
        __props__["public_dns"] = public_dns
        __props__["public_ip"] = public_ip
        __props__["root_block_device"] = root_block_device
        __props__["secondary_private_ips"] = secondary_private_ips
        __props__["security_groups"] = security_groups
        __props__["source_dest_check"] = source_dest_check
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["tenancy"] = tenancy
        __props__["user_data"] = user_data
        __props__["user_data_base64"] = user_data_base64
        __props__["volume_tags"] = volume_tags
        __props__["vpc_security_group_ids"] = vpc_security_group_ids
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ami(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ami")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associatePublicIpAddress")
    def associate_public_ip_address(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "associate_public_ip_address")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cpuCoreCount")
    def cpu_core_count(self) -> pulumi.Output[float]:
        return pulumi.get(self, "cpu_core_count")

    @property
    @pulumi.getter(name="cpuThreadsPerCore")
    def cpu_threads_per_core(self) -> pulumi.Output[float]:
        return pulumi.get(self, "cpu_threads_per_core")

    @property
    @pulumi.getter(name="creditSpecification")
    def credit_specification(self) -> pulumi.Output[Optional['outputs.InstanceCreditSpecification']]:
        return pulumi.get(self, "credit_specification")

    @property
    @pulumi.getter(name="disableApiTermination")
    def disable_api_termination(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "disable_api_termination")

    @property
    @pulumi.getter(name="ebsBlockDevices")
    def ebs_block_devices(self) -> pulumi.Output[List['outputs.InstanceEbsBlockDevice']]:
        return pulumi.get(self, "ebs_block_devices")

    @property
    @pulumi.getter(name="ebsOptimized")
    def ebs_optimized(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "ebs_optimized")

    @property
    @pulumi.getter(name="ephemeralBlockDevices")
    def ephemeral_block_devices(self) -> pulumi.Output[List['outputs.InstanceEphemeralBlockDevice']]:
        return pulumi.get(self, "ephemeral_block_devices")

    @property
    @pulumi.getter(name="getPasswordData")
    def get_password_data(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "get_password_data")

    @property
    @pulumi.getter
    def hibernation(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "hibernation")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter(name="iamInstanceProfile")
    def iam_instance_profile(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "iam_instance_profile")

    @property
    @pulumi.getter(name="instanceInitiatedShutdownBehavior")
    def instance_initiated_shutdown_behavior(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "instance_initiated_shutdown_behavior")

    @property
    @pulumi.getter(name="instanceState")
    def instance_state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_state")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="ipv6AddressCount")
    def ipv6_address_count(self) -> pulumi.Output[float]:
        return pulumi.get(self, "ipv6_address_count")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="metadataOptions")
    def metadata_options(self) -> pulumi.Output['outputs.InstanceMetadataOptions']:
        return pulumi.get(self, "metadata_options")

    @property
    @pulumi.getter
    def monitoring(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "monitoring")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[List['outputs.InstanceNetworkInterface']]:
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter(name="passwordData")
    def password_data(self) -> pulumi.Output[str]:
        return pulumi.get(self, "password_data")

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> pulumi.Output[str]:
        return pulumi.get(self, "placement_group")

    @property
    @pulumi.getter(name="primaryNetworkInterfaceId")
    def primary_network_interface_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "primary_network_interface_id")

    @property
    @pulumi.getter(name="privateDns")
    def private_dns(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_dns")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="publicDns")
    def public_dns(self) -> pulumi.Output[str]:
        return pulumi.get(self, "public_dns")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="rootBlockDevice")
    def root_block_device(self) -> pulumi.Output['outputs.InstanceRootBlockDevice']:
        return pulumi.get(self, "root_block_device")

    @property
    @pulumi.getter(name="secondaryPrivateIps")
    def secondary_private_ips(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "secondary_private_ips")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="sourceDestCheck")
    def source_dest_check(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "source_dest_check")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tenancy(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tenancy")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter(name="userDataBase64")
    def user_data_base64(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_data_base64")

    @property
    @pulumi.getter(name="volumeTags")
    def volume_tags(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "volume_tags")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "vpc_security_group_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


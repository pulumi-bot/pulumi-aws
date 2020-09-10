# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetMountTargetResult',
    'AwaitableGetMountTargetResult',
    'get_mount_target',
]

@pulumi.output_type
class GetMountTargetResult:
    """
    A collection of values returned by getMountTarget.
    """
    def __init__(__self__, availability_zone_id=None, availability_zone_name=None, dns_name=None, file_system_arn=None, file_system_id=None, id=None, ip_address=None, mount_target_dns_name=None, mount_target_id=None, network_interface_id=None, owner_id=None, security_groups=None, subnet_id=None):
        if availability_zone_id and not isinstance(availability_zone_id, str):
            raise TypeError("Expected argument 'availability_zone_id' to be a str")
        pulumi.set(__self__, "availability_zone_id", availability_zone_id)
        if availability_zone_name and not isinstance(availability_zone_name, str):
            raise TypeError("Expected argument 'availability_zone_name' to be a str")
        pulumi.set(__self__, "availability_zone_name", availability_zone_name)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if file_system_arn and not isinstance(file_system_arn, str):
            raise TypeError("Expected argument 'file_system_arn' to be a str")
        pulumi.set(__self__, "file_system_arn", file_system_arn)
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if mount_target_dns_name and not isinstance(mount_target_dns_name, str):
            raise TypeError("Expected argument 'mount_target_dns_name' to be a str")
        pulumi.set(__self__, "mount_target_dns_name", mount_target_dns_name)
        if mount_target_id and not isinstance(mount_target_id, str):
            raise TypeError("Expected argument 'mount_target_id' to be a str")
        pulumi.set(__self__, "mount_target_id", mount_target_id)
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError("Expected argument 'network_interface_id' to be a str")
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="availabilityZoneId")
    def availability_zone_id(self) -> str:
        return pulumi.get(self, "availability_zone_id")

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> str:
        return pulumi.get(self, "availability_zone_name")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="fileSystemArn")
    def file_system_arn(self) -> str:
        return pulumi.get(self, "file_system_arn")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="mountTargetDnsName")
    def mount_target_dns_name(self) -> str:
        return pulumi.get(self, "mount_target_dns_name")

    @property
    @pulumi.getter(name="mountTargetId")
    def mount_target_id(self) -> str:
        return pulumi.get(self, "mount_target_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> List[str]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        return pulumi.get(self, "subnet_id")


class AwaitableGetMountTargetResult(GetMountTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMountTargetResult(
            availability_zone_id=self.availability_zone_id,
            availability_zone_name=self.availability_zone_name,
            dns_name=self.dns_name,
            file_system_arn=self.file_system_arn,
            file_system_id=self.file_system_id,
            id=self.id,
            ip_address=self.ip_address,
            mount_target_dns_name=self.mount_target_dns_name,
            mount_target_id=self.mount_target_id,
            network_interface_id=self.network_interface_id,
            owner_id=self.owner_id,
            security_groups=self.security_groups,
            subnet_id=self.subnet_id)


def get_mount_target(mount_target_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMountTargetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['mountTargetId'] = mount_target_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:efs/getMountTarget:getMountTarget', __args__, opts=opts, typ=GetMountTargetResult).value

    return AwaitableGetMountTargetResult(
        availability_zone_id=__ret__.availability_zone_id,
        availability_zone_name=__ret__.availability_zone_name,
        dns_name=__ret__.dns_name,
        file_system_arn=__ret__.file_system_arn,
        file_system_id=__ret__.file_system_id,
        id=__ret__.id,
        ip_address=__ret__.ip_address,
        mount_target_dns_name=__ret__.mount_target_dns_name,
        mount_target_id=__ret__.mount_target_id,
        network_interface_id=__ret__.network_interface_id,
        owner_id=__ret__.owner_id,
        security_groups=__ret__.security_groups,
        subnet_id=__ret__.subnet_id)

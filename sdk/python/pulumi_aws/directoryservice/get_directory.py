# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetDirectoryResult',
    'AwaitableGetDirectoryResult',
    'get_directory',
]

@pulumi.output_type
class GetDirectoryResult:
    """
    A collection of values returned by getDirectory.
    """
    def __init__(__self__, access_url=None, alias=None, connect_settings=None, description=None, directory_id=None, dns_ip_addresses=None, edition=None, enable_sso=None, id=None, name=None, security_group_id=None, short_name=None, size=None, tags=None, type=None, vpc_settings=None):
        if access_url and not isinstance(access_url, str):
            raise TypeError("Expected argument 'access_url' to be a str")
        pulumi.set(__self__, "access_url", access_url)
        if alias and not isinstance(alias, str):
            raise TypeError("Expected argument 'alias' to be a str")
        pulumi.set(__self__, "alias", alias)
        if connect_settings and not isinstance(connect_settings, list):
            raise TypeError("Expected argument 'connect_settings' to be a list")
        pulumi.set(__self__, "connect_settings", connect_settings)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if directory_id and not isinstance(directory_id, str):
            raise TypeError("Expected argument 'directory_id' to be a str")
        pulumi.set(__self__, "directory_id", directory_id)
        if dns_ip_addresses and not isinstance(dns_ip_addresses, list):
            raise TypeError("Expected argument 'dns_ip_addresses' to be a list")
        pulumi.set(__self__, "dns_ip_addresses", dns_ip_addresses)
        if edition and not isinstance(edition, str):
            raise TypeError("Expected argument 'edition' to be a str")
        pulumi.set(__self__, "edition", edition)
        if enable_sso and not isinstance(enable_sso, bool):
            raise TypeError("Expected argument 'enable_sso' to be a bool")
        pulumi.set(__self__, "enable_sso", enable_sso)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if short_name and not isinstance(short_name, str):
            raise TypeError("Expected argument 'short_name' to be a str")
        pulumi.set(__self__, "short_name", short_name)
        if size and not isinstance(size, str):
            raise TypeError("Expected argument 'size' to be a str")
        pulumi.set(__self__, "size", size)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vpc_settings and not isinstance(vpc_settings, list):
            raise TypeError("Expected argument 'vpc_settings' to be a list")
        pulumi.set(__self__, "vpc_settings", vpc_settings)

    @property
    @pulumi.getter(name="accessUrl")
    def access_url(self) -> str:
        return pulumi.get(self, "access_url")

    @property
    @pulumi.getter
    def alias(self) -> str:
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="connectSettings")
    def connect_settings(self) -> List['outputs.GetDirectoryConnectSettingResult']:
        return pulumi.get(self, "connect_settings")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> str:
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter(name="dnsIpAddresses")
    def dns_ip_addresses(self) -> List[str]:
        return pulumi.get(self, "dns_ip_addresses")

    @property
    @pulumi.getter
    def edition(self) -> str:
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> bool:
        return pulumi.get(self, "enable_sso")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> str:
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter
    def size(self) -> str:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcSettings")
    def vpc_settings(self) -> List['outputs.GetDirectoryVpcSettingResult']:
        return pulumi.get(self, "vpc_settings")


class AwaitableGetDirectoryResult(GetDirectoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryResult(
            access_url=self.access_url,
            alias=self.alias,
            connect_settings=self.connect_settings,
            description=self.description,
            directory_id=self.directory_id,
            dns_ip_addresses=self.dns_ip_addresses,
            edition=self.edition,
            enable_sso=self.enable_sso,
            id=self.id,
            name=self.name,
            security_group_id=self.security_group_id,
            short_name=self.short_name,
            size=self.size,
            tags=self.tags,
            type=self.type,
            vpc_settings=self.vpc_settings)


def get_directory(directory_id: Optional[str] = None,
                  tags: Optional[Mapping[str, str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['directoryId'] = directory_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:directoryservice/getDirectory:getDirectory', __args__, opts=opts, typ=GetDirectoryResult).value

    return AwaitableGetDirectoryResult(
        access_url=__ret__.access_url,
        alias=__ret__.alias,
        connect_settings=__ret__.connect_settings,
        description=__ret__.description,
        directory_id=__ret__.directory_id,
        dns_ip_addresses=__ret__.dns_ip_addresses,
        edition=__ret__.edition,
        enable_sso=__ret__.enable_sso,
        id=__ret__.id,
        name=__ret__.name,
        security_group_id=__ret__.security_group_id,
        short_name=__ret__.short_name,
        size=__ret__.size,
        tags=__ret__.tags,
        type=__ret__.type,
        vpc_settings=__ret__.vpc_settings)

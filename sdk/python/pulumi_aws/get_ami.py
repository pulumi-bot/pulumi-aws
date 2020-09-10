# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetAmiResult',
    'AwaitableGetAmiResult',
    'get_ami',
]

@pulumi.output_type
class GetAmiResult:
    """
    A collection of values returned by getAmi.
    """
    def __init__(__self__, architecture=None, arn=None, block_device_mappings=None, creation_date=None, description=None, executable_users=None, filters=None, hypervisor=None, id=None, image_id=None, image_location=None, image_owner_alias=None, image_type=None, kernel_id=None, most_recent=None, name=None, name_regex=None, owner_id=None, owners=None, platform=None, product_codes=None, public=None, ramdisk_id=None, root_device_name=None, root_device_type=None, root_snapshot_id=None, sriov_net_support=None, state=None, state_reason=None, tags=None, virtualization_type=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if block_device_mappings and not isinstance(block_device_mappings, list):
            raise TypeError("Expected argument 'block_device_mappings' to be a list")
        pulumi.set(__self__, "block_device_mappings", block_device_mappings)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if executable_users and not isinstance(executable_users, list):
            raise TypeError("Expected argument 'executable_users' to be a list")
        pulumi.set(__self__, "executable_users", executable_users)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if hypervisor and not isinstance(hypervisor, str):
            raise TypeError("Expected argument 'hypervisor' to be a str")
        pulumi.set(__self__, "hypervisor", hypervisor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if image_location and not isinstance(image_location, str):
            raise TypeError("Expected argument 'image_location' to be a str")
        pulumi.set(__self__, "image_location", image_location)
        if image_owner_alias and not isinstance(image_owner_alias, str):
            raise TypeError("Expected argument 'image_owner_alias' to be a str")
        pulumi.set(__self__, "image_owner_alias", image_owner_alias)
        if image_type and not isinstance(image_type, str):
            raise TypeError("Expected argument 'image_type' to be a str")
        pulumi.set(__self__, "image_type", image_type)
        if kernel_id and not isinstance(kernel_id, str):
            raise TypeError("Expected argument 'kernel_id' to be a str")
        pulumi.set(__self__, "kernel_id", kernel_id)
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        pulumi.set(__self__, "most_recent", most_recent)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        pulumi.set(__self__, "owners", owners)
        if platform and not isinstance(platform, str):
            raise TypeError("Expected argument 'platform' to be a str")
        pulumi.set(__self__, "platform", platform)
        if product_codes and not isinstance(product_codes, list):
            raise TypeError("Expected argument 'product_codes' to be a list")
        pulumi.set(__self__, "product_codes", product_codes)
        if public and not isinstance(public, bool):
            raise TypeError("Expected argument 'public' to be a bool")
        pulumi.set(__self__, "public", public)
        if ramdisk_id and not isinstance(ramdisk_id, str):
            raise TypeError("Expected argument 'ramdisk_id' to be a str")
        pulumi.set(__self__, "ramdisk_id", ramdisk_id)
        if root_device_name and not isinstance(root_device_name, str):
            raise TypeError("Expected argument 'root_device_name' to be a str")
        pulumi.set(__self__, "root_device_name", root_device_name)
        if root_device_type and not isinstance(root_device_type, str):
            raise TypeError("Expected argument 'root_device_type' to be a str")
        pulumi.set(__self__, "root_device_type", root_device_type)
        if root_snapshot_id and not isinstance(root_snapshot_id, str):
            raise TypeError("Expected argument 'root_snapshot_id' to be a str")
        pulumi.set(__self__, "root_snapshot_id", root_snapshot_id)
        if sriov_net_support and not isinstance(sriov_net_support, str):
            raise TypeError("Expected argument 'sriov_net_support' to be a str")
        pulumi.set(__self__, "sriov_net_support", sriov_net_support)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_reason and not isinstance(state_reason, dict):
            raise TypeError("Expected argument 'state_reason' to be a dict")
        pulumi.set(__self__, "state_reason", state_reason)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if virtualization_type and not isinstance(virtualization_type, str):
            raise TypeError("Expected argument 'virtualization_type' to be a str")
        pulumi.set(__self__, "virtualization_type", virtualization_type)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="blockDeviceMappings")
    def block_device_mappings(self) -> List['outputs.GetAmiBlockDeviceMappingResult']:
        return pulumi.get(self, "block_device_mappings")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executableUsers")
    def executable_users(self) -> Optional[List[str]]:
        return pulumi.get(self, "executable_users")

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetAmiFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def hypervisor(self) -> str:
        return pulumi.get(self, "hypervisor")

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
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageLocation")
    def image_location(self) -> str:
        return pulumi.get(self, "image_location")

    @property
    @pulumi.getter(name="imageOwnerAlias")
    def image_owner_alias(self) -> str:
        return pulumi.get(self, "image_owner_alias")

    @property
    @pulumi.getter(name="imageType")
    def image_type(self) -> str:
        return pulumi.get(self, "image_type")

    @property
    @pulumi.getter(name="kernelId")
    def kernel_id(self) -> str:
        return pulumi.get(self, "kernel_id")

    @property
    @pulumi.getter(name="mostRecent")
    def most_recent(self) -> Optional[bool]:
        return pulumi.get(self, "most_recent")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def owners(self) -> List[str]:
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter
    def platform(self) -> str:
        return pulumi.get(self, "platform")

    @property
    @pulumi.getter(name="productCodes")
    def product_codes(self) -> List['outputs.GetAmiProductCodeResult']:
        return pulumi.get(self, "product_codes")

    @property
    @pulumi.getter
    def public(self) -> bool:
        return pulumi.get(self, "public")

    @property
    @pulumi.getter(name="ramdiskId")
    def ramdisk_id(self) -> str:
        return pulumi.get(self, "ramdisk_id")

    @property
    @pulumi.getter(name="rootDeviceName")
    def root_device_name(self) -> str:
        return pulumi.get(self, "root_device_name")

    @property
    @pulumi.getter(name="rootDeviceType")
    def root_device_type(self) -> str:
        return pulumi.get(self, "root_device_type")

    @property
    @pulumi.getter(name="rootSnapshotId")
    def root_snapshot_id(self) -> str:
        return pulumi.get(self, "root_snapshot_id")

    @property
    @pulumi.getter(name="sriovNetSupport")
    def sriov_net_support(self) -> str:
        return pulumi.get(self, "sriov_net_support")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateReason")
    def state_reason(self) -> Mapping[str, str]:
        return pulumi.get(self, "state_reason")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualizationType")
    def virtualization_type(self) -> str:
        return pulumi.get(self, "virtualization_type")


class AwaitableGetAmiResult(GetAmiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAmiResult(
            architecture=self.architecture,
            arn=self.arn,
            block_device_mappings=self.block_device_mappings,
            creation_date=self.creation_date,
            description=self.description,
            executable_users=self.executable_users,
            filters=self.filters,
            hypervisor=self.hypervisor,
            id=self.id,
            image_id=self.image_id,
            image_location=self.image_location,
            image_owner_alias=self.image_owner_alias,
            image_type=self.image_type,
            kernel_id=self.kernel_id,
            most_recent=self.most_recent,
            name=self.name,
            name_regex=self.name_regex,
            owner_id=self.owner_id,
            owners=self.owners,
            platform=self.platform,
            product_codes=self.product_codes,
            public=self.public,
            ramdisk_id=self.ramdisk_id,
            root_device_name=self.root_device_name,
            root_device_type=self.root_device_type,
            root_snapshot_id=self.root_snapshot_id,
            sriov_net_support=self.sriov_net_support,
            state=self.state,
            state_reason=self.state_reason,
            tags=self.tags,
            virtualization_type=self.virtualization_type)


def get_ami(executable_users: Optional[List[str]] = None,
            filters: Optional[List[pulumi.InputType['GetAmiFilterArgs']]] = None,
            most_recent: Optional[bool] = None,
            name_regex: Optional[str] = None,
            owners: Optional[List[str]] = None,
            tags: Optional[Mapping[str, str]] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAmiResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['executableUsers'] = executable_users
    __args__['filters'] = filters
    __args__['mostRecent'] = most_recent
    __args__['nameRegex'] = name_regex
    __args__['owners'] = owners
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getAmi:getAmi', __args__, opts=opts, typ=GetAmiResult).value

    return AwaitableGetAmiResult(
        architecture=__ret__.architecture,
        arn=__ret__.arn,
        block_device_mappings=__ret__.block_device_mappings,
        creation_date=__ret__.creation_date,
        description=__ret__.description,
        executable_users=__ret__.executable_users,
        filters=__ret__.filters,
        hypervisor=__ret__.hypervisor,
        id=__ret__.id,
        image_id=__ret__.image_id,
        image_location=__ret__.image_location,
        image_owner_alias=__ret__.image_owner_alias,
        image_type=__ret__.image_type,
        kernel_id=__ret__.kernel_id,
        most_recent=__ret__.most_recent,
        name=__ret__.name,
        name_regex=__ret__.name_regex,
        owner_id=__ret__.owner_id,
        owners=__ret__.owners,
        platform=__ret__.platform,
        product_codes=__ret__.product_codes,
        public=__ret__.public,
        ramdisk_id=__ret__.ramdisk_id,
        root_device_name=__ret__.root_device_name,
        root_device_type=__ret__.root_device_type,
        root_snapshot_id=__ret__.root_snapshot_id,
        sriov_net_support=__ret__.sriov_net_support,
        state=__ret__.state,
        state_reason=__ret__.state_reason,
        tags=__ret__.tags,
        virtualization_type=__ret__.virtualization_type)

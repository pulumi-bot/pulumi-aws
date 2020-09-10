# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetDelegationSetResult',
    'AwaitableGetDelegationSetResult',
    'get_delegation_set',
]

@pulumi.output_type
class GetDelegationSetResult:
    """
    A collection of values returned by getDelegationSet.
    """
    def __init__(__self__, caller_reference=None, id=None, name_servers=None):
        if caller_reference and not isinstance(caller_reference, str):
            raise TypeError("Expected argument 'caller_reference' to be a str")
        pulumi.set(__self__, "caller_reference", caller_reference)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_servers and not isinstance(name_servers, list):
            raise TypeError("Expected argument 'name_servers' to be a list")
        pulumi.set(__self__, "name_servers", name_servers)

    @property
    @pulumi.getter(name="callerReference")
    def caller_reference(self) -> str:
        return pulumi.get(self, "caller_reference")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> List[str]:
        return pulumi.get(self, "name_servers")


class AwaitableGetDelegationSetResult(GetDelegationSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDelegationSetResult(
            caller_reference=self.caller_reference,
            id=self.id,
            name_servers=self.name_servers)


def get_delegation_set(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDelegationSetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:route53/getDelegationSet:getDelegationSet', __args__, opts=opts, typ=GetDelegationSetResult).value

    return AwaitableGetDelegationSetResult(
        caller_reference=__ret__.caller_reference,
        id=__ret__.id,
        name_servers=__ret__.name_servers)

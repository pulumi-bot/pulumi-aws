# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetEndpointResult',
    'AwaitableGetEndpointResult',
    'get_endpoint',
]

@pulumi.output_type
class GetEndpointResult:
    """
    A collection of values returned by getEndpoint.
    """
    def __init__(__self__, endpoint_address=None, endpoint_type=None, id=None):
        if endpoint_address and not isinstance(endpoint_address, str):
            raise TypeError("Expected argument 'endpoint_address' to be a str")
        pulumi.set(__self__, "endpoint_address", endpoint_address)
        if endpoint_type and not isinstance(endpoint_type, str):
            raise TypeError("Expected argument 'endpoint_type' to be a str")
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="endpointAddress")
    def endpoint_address(self) -> str:
        return pulumi.get(self, "endpoint_address")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[str]:
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetEndpointResult(GetEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointResult(
            endpoint_address=self.endpoint_address,
            endpoint_type=self.endpoint_type,
            id=self.id)


def get_endpoint(endpoint_type: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['endpointType'] = endpoint_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:iot/getEndpoint:getEndpoint', __args__, opts=opts, typ=GetEndpointResult).value

    return AwaitableGetEndpointResult(
        endpoint_address=__ret__.endpoint_address,
        endpoint_type=__ret__.endpoint_type,
        id=__ret__.id)

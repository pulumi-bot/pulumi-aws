# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetResolverEndpointResult',
    'AwaitableGetResolverEndpointResult',
    'get_resolver_endpoint',
]

@pulumi.output_type
class GetResolverEndpointResult:
    """
    A collection of values returned by getResolverEndpoint.
    """
    def __init__(__self__, arn=None, direction=None, filters=None, id=None, ip_addresses=None, name=None, resolver_endpoint_id=None, status=None, vpc_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resolver_endpoint_id and not isinstance(resolver_endpoint_id, str):
            raise TypeError("Expected argument 'resolver_endpoint_id' to be a str")
        pulumi.set(__self__, "resolver_endpoint_id", resolver_endpoint_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def direction(self) -> str:
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetResolverEndpointFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resolverEndpointId")
    def resolver_endpoint_id(self) -> Optional[str]:
        return pulumi.get(self, "resolver_endpoint_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetResolverEndpointResult(GetResolverEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverEndpointResult(
            arn=self.arn,
            direction=self.direction,
            filters=self.filters,
            id=self.id,
            ip_addresses=self.ip_addresses,
            name=self.name,
            resolver_endpoint_id=self.resolver_endpoint_id,
            status=self.status,
            vpc_id=self.vpc_id)


def get_resolver_endpoint(filters: Optional[Sequence[pulumi.InputType['GetResolverEndpointFilterArgs']]] = None,
                          resolver_endpoint_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverEndpointResult:
    """
    `route53.ResolverEndpoint` provides details about a specific Route53 Resolver Endpoint.

    This data source allows to find a list of IPaddresses associated with a specific Route53 Resolver Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.route53.get_resolver_endpoint(resolver_endpoint_id="rslvr-in-1abc2345ef678g91h")
    ```

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.route53.get_resolver_endpoint(filters=[{
        "name": "NAME",
        "values": ["MyResolverExampleName"],
    }])
    ```


    :param Sequence[pulumi.InputType['GetResolverEndpointFilterArgs']] filters: One or more name/value pairs to use as filters. There are
           several valid keys, for a full reference, check out
           [Route53resolver Filter value in the AWS API reference][1].
    :param str resolver_endpoint_id: The ID of the Route53 Resolver Endpoint.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['resolverEndpointId'] = resolver_endpoint_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:route53/getResolverEndpoint:getResolverEndpoint', __args__, opts=opts, typ=GetResolverEndpointResult).value

    return AwaitableGetResolverEndpointResult(
        arn=__ret__.arn,
        direction=__ret__.direction,
        filters=__ret__.filters,
        id=__ret__.id,
        ip_addresses=__ret__.ip_addresses,
        name=__ret__.name,
        resolver_endpoint_id=__ret__.resolver_endpoint_id,
        status=__ret__.status,
        vpc_id=__ret__.vpc_id)

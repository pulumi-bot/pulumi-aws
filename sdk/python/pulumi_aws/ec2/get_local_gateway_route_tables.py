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
    'GetLocalGatewayRouteTablesResult',
    'AwaitableGetLocalGatewayRouteTablesResult',
    'get_local_gateway_route_tables',
]



@pulumi.output_type
class GetLocalGatewayRouteTablesResult:
    """
    A collection of values returned by getLocalGatewayRouteTables.
    """
    def __init__(__self__, filters=None, id=None, ids=None, tags=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetLocalGatewayRouteTablesFilterResult']]:
        ...

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        ...

    @property
    @pulumi.getter
    def ids(self) -> List[str]:
        """
        Set of Local Gateway Route Table identifiers
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        ...



class AwaitableGetLocalGatewayRouteTablesResult(GetLocalGatewayRouteTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocalGatewayRouteTablesResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            tags=self.tags)


def get_local_gateway_route_tables(filters: Optional[List[pulumi.InputType['GetLocalGatewayRouteTablesFilterArgs']]] = None,
                                   tags: Optional[Mapping[str, str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocalGatewayRouteTablesResult:
    """
    Provides information for multiple EC2 Local Gateway Route Tables, such as their identifiers.


    :param List[pulumi.InputType['GetLocalGatewayRouteTablesFilterArgs']] filters: Custom filter block as described below.
    :param Mapping[str, str] tags: A mapping of tags, each pair of which must exactly match
           a pair on the desired local gateway route table.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLocalGatewayRouteTables:getLocalGatewayRouteTables', __args__, opts=opts, typ=GetLocalGatewayRouteTablesResult).value

    return AwaitableGetLocalGatewayRouteTablesResult(
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        tags=__ret__.tags)

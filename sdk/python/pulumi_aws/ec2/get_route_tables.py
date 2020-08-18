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
    'GetRouteTablesResult',
    'AwaitableGetRouteTablesResult',
    'get_route_tables',
]



@pulumi.output_type
class GetRouteTablesResult:
    """
    A collection of values returned by getRouteTables.
    """
    def __init__(__self__, filters=None, id=None, ids=None, tags=None, vpc_id=None):
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
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetRouteTablesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> List[str]:
        """
        A set of all the route table ids found. This data source will fail if none are found.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")



class AwaitableGetRouteTablesResult(GetRouteTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteTablesResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_route_tables(filters: Optional[List[pulumi.InputType['GetRouteTablesFilterArgs']]] = None,
                     tags: Optional[Mapping[str, str]] = None,
                     vpc_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteTablesResult:
    """
    This resource can be useful for getting back a list of route table ids to be referenced elsewhere.

    ## Example Usage

    The following adds a route for a particular cidr block to every (private
    kops) route table in a specified vpc to use a particular vpc peering
    connection.

    ```python
    import pulumi
    import pulumi_aws as aws

    rts = aws.ec2.get_route_tables(vpc_id=var["vpc_id"],
        filters=[{
            "name": "tag:kubernetes.io/kops/role",
            "values": ["private*"],
        }])
    route = []
    for range in [{"value": i} for i in range(0, len(rts.ids))]:
        route.append(aws.ec2.Route(f"route-{range['value']}",
            route_table_id=rts.ids[range["value"]],
            destination_cidr_block="10.0.1.0/22",
            vpc_peering_connection_id="pcx-0e9a7a9ecd137dc54"))
    ```


    :param List[pulumi.InputType['GetRouteTablesFilterArgs']] filters: Custom filter block as described below.
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match
           a pair on the desired route tables.
    :param str vpc_id: The VPC ID that you want to filter from.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getRouteTables:getRouteTables', __args__, opts=opts, typ=GetRouteTablesResult).value

    return AwaitableGetRouteTablesResult(
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)

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
    'GetRouteTableResult',
    'AwaitableGetRouteTableResult',
    'get_route_table',
]


class GetRouteTableResult:
    """
    A collection of values returned by getRouteTable.
    """
    def __init__(__self__, default_association_route_table=None, default_propagation_route_table=None, filters=None, id=None, tags=None, transit_gateway_id=None):
        if default_association_route_table and not isinstance(default_association_route_table, bool):
            raise TypeError("Expected argument 'default_association_route_table' to be a bool")
        __self__.default_association_route_table = default_association_route_table
        """
        Boolean whether this is the default association route table for the EC2 Transit Gateway
        """
        if default_propagation_route_table and not isinstance(default_propagation_route_table, bool):
            raise TypeError("Expected argument 'default_propagation_route_table' to be a bool")
        __self__.default_propagation_route_table = default_propagation_route_table
        """
        Boolean whether this is the default propagation route table for the EC2 Transit Gateway
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        EC2 Transit Gateway Route Table identifier
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Key-value tags for the EC2 Transit Gateway Route Table
        """
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        __self__.transit_gateway_id = transit_gateway_id
        """
        EC2 Transit Gateway identifier
        """


class AwaitableGetRouteTableResult(GetRouteTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteTableResult(
            default_association_route_table=self.default_association_route_table,
            default_propagation_route_table=self.default_propagation_route_table,
            filters=self.filters,
            id=self.id,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id)


def get_route_table(filters: Optional[List[pulumi.InputType['GetRouteTableFilterArgs']]] = None,
                    id: Optional[str] = None,
                    tags: Optional[Mapping[str, str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteTableResult:
    """
    Get information on an EC2 Transit Gateway Route Table.

    ## Example Usage
    ### By Filter

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_route_table(filters=[
        {
            "name": "default-association-route-table",
            "values": ["true"],
        },
        {
            "name": "transit-gateway-id",
            "values": ["tgw-12345678"],
        },
    ])
    ```
    ### By Identifier

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_route_table(id="tgw-rtb-12345678")
    ```


    :param List[pulumi.InputType['GetRouteTableFilterArgs']] filters: One or more configuration blocks containing name-values filters. Detailed below.
    :param str id: Identifier of the EC2 Transit Gateway Route Table.
    :param Mapping[str, str] tags: Key-value tags for the EC2 Transit Gateway Route Table
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getRouteTable:getRouteTable', __args__, opts=opts).value

    return AwaitableGetRouteTableResult(
        default_association_route_table=__ret__.get('defaultAssociationRouteTable'),
        default_propagation_route_table=__ret__.get('defaultPropagationRouteTable'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        tags=__ret__.get('tags'),
        transit_gateway_id=__ret__.get('transitGatewayId'))

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



@pulumi.output_type
class GetRouteTableResult:
    """
    A collection of values returned by getRouteTable.
    """
    def __init__(__self__, associations=None, filters=None, gateway_id=None, id=None, owner_id=None, route_table_id=None, routes=None, subnet_id=None, tags=None, vpc_id=None):
        if associations and not isinstance(associations, list):
            raise TypeError("Expected argument 'associations' to be a list")
        pulumi.set(__self__, "associations", associations)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        pulumi.set(__self__, "gateway_id", gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError("Expected argument 'route_table_id' to be a str")
        pulumi.set(__self__, "route_table_id", route_table_id)
        if routes and not isinstance(routes, list):
            raise TypeError("Expected argument 'routes' to be a list")
        pulumi.set(__self__, "routes", routes)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def associations(self) -> List['outputs.GetRouteTableAssociationResult']:
        return pulumi.get(self, "associations")

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetRouteTableFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> str:
        """
        The Gateway ID. Only set when associated with an Internet Gateway or Virtual Private Gateway.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        The ID of the AWS account that owns the route table
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> str:
        """
        The Route Table ID.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter
    def routes(self) -> List['outputs.GetRouteTableRouteResult']:
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The Subnet ID. Only set when associated with a Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")



class AwaitableGetRouteTableResult(GetRouteTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteTableResult(
            associations=self.associations,
            filters=self.filters,
            gateway_id=self.gateway_id,
            id=self.id,
            owner_id=self.owner_id,
            route_table_id=self.route_table_id,
            routes=self.routes,
            subnet_id=self.subnet_id,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_route_table(filters: Optional[List[pulumi.InputType['GetRouteTableFilterArgs']]] = None,
                    gateway_id: Optional[str] = None,
                    route_table_id: Optional[str] = None,
                    subnet_id: Optional[str] = None,
                    tags: Optional[Mapping[str, str]] = None,
                    vpc_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteTableResult:
    """
    `ec2.RouteTable` provides details about a specific Route Table.

    This resource can prove useful when a module accepts a Subnet id as
    an input variable and needs to, for example, add a route in
    the Route Table.

    ## Example Usage

    The following example shows how one might accept a Route Table id as a variable
    and use this data source to obtain the data necessary to create a route.

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    selected = aws.ec2.get_route_table(subnet_id=subnet_id)
    route = aws.ec2.Route("route",
        route_table_id=selected.id,
        destination_cidr_block="10.0.1.0/22",
        vpc_peering_connection_id="pcx-45ff3dc1")
    ```


    :param List[pulumi.InputType['GetRouteTableFilterArgs']] filters: Custom filter block as described below.
    :param str gateway_id: The id of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
    :param str route_table_id: The id of the specific Route Table to retrieve.
    :param str subnet_id: The id of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match
           a pair on the desired Route Table.
    :param str vpc_id: The id of the VPC that the desired Route Table belongs to.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['gatewayId'] = gateway_id
    __args__['routeTableId'] = route_table_id
    __args__['subnetId'] = subnet_id
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getRouteTable:getRouteTable', __args__, opts=opts, typ=GetRouteTableResult).value

    return AwaitableGetRouteTableResult(
        associations=__ret__.associations,
        filters=__ret__.filters,
        gateway_id=__ret__.gateway_id,
        id=__ret__.id,
        owner_id=__ret__.owner_id,
        route_table_id=__ret__.route_table_id,
        routes=__ret__.routes,
        subnet_id=__ret__.subnet_id,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)

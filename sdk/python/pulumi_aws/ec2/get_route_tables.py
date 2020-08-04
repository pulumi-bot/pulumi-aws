# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetRouteTablesResult:
    """
    A collection of values returned by getRouteTables.
    """
    def __init__(__self__, filters=None, id=None, ids=None, tags=None, vpc_id=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A set of all the route table ids found. This data source will fail if none are found.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id


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


def get_route_tables(filters=None, tags=None, vpc_id=None, opts=None):
    """
    This resource can be useful for getting back a list of route table ids to be referenced elsewhere.

    ## Example Usage

    The following adds a route for a particular cidr block to every (private
    kops) route table in a specified vpc to use a particular vpc peering
    connection.

    ```python
    import pulumi
    import pulumi_aws as aws

    rts = aws.ec2.get_route_tables(filters=[aws.ec2.GetRouteTablesFilterArgs(
            name="tag:kubernetes.io/kops/role",
            values=["private*"],
        )],
        vpc_id=var["vpc_id"])
    route = []
    for range in [{"value": i} for i in range(0, len(rts.ids))]:
        route.append(aws.ec2.Route(f"route-{range['value']}",
            destination_cidr_block="10.0.1.0/22",
            route_table_id=rts.ids[range["value"]],
            vpc_peering_connection_id="pcx-0e9a7a9ecd137dc54"))
    ```


    :param list filters: Custom filter block as described below.
    :param dict tags: A map of tags, each pair of which must exactly match
           a pair on the desired route tables.
    :param str vpc_id: The VPC ID that you want to filter from.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeRouteTables.html).
      * `values` (`list`) - Set of values that are accepted for the given field.
        A Route Table will be selected if any one of the given values matches.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getRouteTables:getRouteTables', __args__, opts=opts).value

    return AwaitableGetRouteTablesResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'))

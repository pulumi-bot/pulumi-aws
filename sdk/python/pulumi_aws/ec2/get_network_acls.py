# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetNetworkAclsResult:
    """
    A collection of values returned by getNetworkAcls.
    """
    def __init__(__self__, filters=None, ids=None, tags=None, vpc_id=None, id=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of all the network ACL ids found. This data source will fail if none are found.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_network_acls(filters=None,tags=None,vpc_id=None,opts=None):
    __args__ = dict()

    __args__['filters'] = filters
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __ret__ = await pulumi.runtime.invoke('aws:ec2/getNetworkAcls:getNetworkAcls', __args__, opts=opts)

    return GetNetworkAclsResult(
        filters=__ret__.get('filters'),
        ids=__ret__.get('ids'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'),
        id=__ret__.get('id'))

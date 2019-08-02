# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetVpcsResult:
    """
    A collection of values returned by getVpcs.
    """
    def __init__(__self__, filters=None, ids=None, tags=None, id=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of all the VPC Ids found. This data source will fail if none are found.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_vpcs(filters=None,tags=None,opts=None):
    """
    This resource can be useful for getting back a list of VPC Ids for a region.
    
    The following example retrieves a list of VPC Ids with a custom tag of `service` set to a value of "production".

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpcs.html.markdown.
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('aws:ec2/getVpcs:getVpcs', __args__, opts=opts)

    return GetVpcsResult(
        filters=__ret__.get('filters'),
        ids=__ret__.get('ids'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))

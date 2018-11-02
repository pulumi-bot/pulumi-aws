# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetVpcsResult(object):
    """
    A collection of values returned by getVpcs.
    """
    def __init__(__self__, ids=None, tags=None, id=None):
        if ids and not isinstance(ids, list):
            raise TypeError('Expected argument ids to be a list')
        __self__.ids = ids
        """
        A list of all the VPC Ids found. This data source will fail if none are found.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_vpcs(filters=None, tags=None):
    """
    This resource can be useful for getting back a list of VPC Ids for a region.
    
    The following example retrieves a list of VPC Ids with a custom tag of `service` set to a value of "production".
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['tags'] = tags
    __ret__ = pulumi.runtime.invoke('aws:ec2/getVpcs:getVpcs', __args__)

    return GetVpcsResult(
        ids=__ret__.get('ids'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))

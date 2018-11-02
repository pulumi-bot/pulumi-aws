# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities

class GetRegionResult(object):
    """
    A collection of values returned by getRegion.
    """
    def __init__(__self__, current=None, description=None, endpoint=None, name=None, id=None):
        if current and not isinstance(current, bool):
            raise TypeError('Expected argument current to be a bool')
        __self__.current = current
        """
        `true` if the selected region is the one configured on the
        provider, or `false` otherwise.
        """
        if description and not isinstance(description, str):
            raise TypeError('Expected argument description to be a str')
        __self__.description = description
        """
        The region's description in this format: "Location (Region name)".
        """
        if endpoint and not isinstance(endpoint, str):
            raise TypeError('Expected argument endpoint to be a str')
        __self__.endpoint = endpoint
        """
        The EC2 endpoint for the selected region.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        """
        The name of the selected region.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_region(current=None, endpoint=None, name=None):
    """
    `aws_region` provides details about a specific AWS region.
    
    As well as validating a given region name this resource can be used to
    discover the name of the region configured within the provider. The latter
    can be useful in a child module which is inheriting an AWS provider
    configuration from its parent module.
    """
    __args__ = dict()

    __args__['current'] = current
    __args__['endpoint'] = endpoint
    __args__['name'] = name
    __ret__ = pulumi.runtime.invoke('aws:index/getRegion:getRegion', __args__)

    return GetRegionResult(
        current=__ret__.get('current'),
        description=__ret__.get('description'),
        endpoint=__ret__.get('endpoint'),
        name=__ret__.get('name'),
        id=__ret__.get('id'))

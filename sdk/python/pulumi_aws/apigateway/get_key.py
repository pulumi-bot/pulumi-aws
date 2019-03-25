# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetKeyResult:
    """
    A collection of values returned by getKey.
    """
    def __init__(__self__, id=None, name=None, value=None):
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        Set to the ID of the API Key.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        """
        Set to the name of the API Key.
        """
        if value and not isinstance(value, str):
            raise TypeError('Expected argument value to be a str')
        __self__.value = value
        """
        Set to the value of the API Key.
        """

async def get_key(id=None,opts=None):
    """
    Use this data source to get the name and value of a pre-existing API Key, for
    example to supply credentials for a dependency microservice.
    """
    __args__ = dict()

    __args__['id'] = id
    __ret__ = await pulumi.runtime.invoke('aws:apigateway/getKey:getKey', __args__, opts=opts)

    return GetKeyResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        value=__ret__.get('value'))

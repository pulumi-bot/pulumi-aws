# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetResourceResult:
    """
    A collection of values returned by getResource.
    """
    def __init__(__self__, parent_id=None, path=None, path_part=None, rest_api_id=None, id=None):
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        __self__.parent_id = parent_id
        """
        Set to the ID of the parent Resource.
        """
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if path_part and not isinstance(path_part, str):
            raise TypeError("Expected argument 'path_part' to be a str")
        __self__.path_part = path_part
        """
        Set to the path relative to the parent Resource.
        """
        if rest_api_id and not isinstance(rest_api_id, str):
            raise TypeError("Expected argument 'rest_api_id' to be a str")
        __self__.rest_api_id = rest_api_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_resource(path=None,rest_api_id=None,opts=None):
    """
    Use this data source to get the id of a Resource in API Gateway. 
    To fetch the Resource, you must provide the REST API id as well as the full path.  
    """
    __args__ = dict()

    __args__['path'] = path
    __args__['restApiId'] = rest_api_id
 .   if opts is None:
         opts = pulumi.ResourceOptions()
     if opts.version is None:
         opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('aws:apigateway/getResource:getResource', __args__, opts=opts)

    return GetResourceResult(
        parent_id=__ret__.get('parentId'),
        path=__ret__.get('path'),
        path_part=__ret__.get('pathPart'),
        rest_api_id=__ret__.get('restApiId'),
        id=__ret__.get('id'))

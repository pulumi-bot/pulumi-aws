# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, arn=None, path=None, permissions_boundary=None, user_id=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) assigned by AWS for this user.
        """
        if path and not isinstance(path, str):
            raise TypeError('Expected argument path to be a str')
        __self__.path = path
        """
        Path in which this user was created.
        """
        if permissions_boundary and not isinstance(permissions_boundary, str):
            raise TypeError('Expected argument permissions_boundary to be a str')
        __self__.permissions_boundary = permissions_boundary
        """
        The ARN of the policy that is used to set the permissions boundary for the user.
        """
        if user_id and not isinstance(user_id, str):
            raise TypeError('Expected argument user_id to be a str')
        __self__.user_id = user_id
        """
        The unique ID assigned by AWS for this user.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_user(user_name=None,opts=None):
    """
    This data source can be used to fetch information about a specific
    IAM user. By using this data source, you can reference IAM user
    properties without having to hard code ARNs or unique IDs as input.
    """
    __args__ = dict()

    __args__['userName'] = user_name
    __ret__ = await pulumi.runtime.invoke('aws:iam/getUser:getUser', __args__, opts=opts)

    return GetUserResult(
        arn=__ret__.get('arn'),
        path=__ret__.get('path'),
        permissions_boundary=__ret__.get('permissionsBoundary'),
        user_id=__ret__.get('userId'),
        id=__ret__.get('id'))

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, arn=None, id=None, path=None, permissions_boundary=None, user_id=None, user_name=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) assigned by AWS for this user.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        """
        Path in which this user was created.
        """
        if permissions_boundary and not isinstance(permissions_boundary, str):
            raise TypeError("Expected argument 'permissions_boundary' to be a str")
        __self__.permissions_boundary = permissions_boundary
        """
        The ARN of the policy that is used to set the permissions boundary for the user.
        """
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        __self__.user_id = user_id
        """
        The unique ID assigned by AWS for this user.
        """
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        __self__.user_name = user_name
        """
        The name associated to this User
        """
class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            arn=self.arn,
            id=self.id,
            path=self.path,
            permissions_boundary=self.permissions_boundary,
            user_id=self.user_id,
            user_name=self.user_name)

def get_user(user_name=None,opts=None):
    """
    This data source can be used to fetch information about a specific
    IAM user. By using this data source, you can reference IAM user
    properties without having to hard code ARNs or unique IDs as input.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.iam.get_user(user_name="an_example_user_name")
    ```


    :param str user_name: The friendly IAM user name to match.
    """
    __args__ = dict()


    __args__['userName'] = user_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:iam/getUser:getUser', __args__, opts=opts).value

    return AwaitableGetUserResult(
        arn=__ret__.get('arn'),
        id=__ret__.get('id'),
        path=__ret__.get('path'),
        permissions_boundary=__ret__.get('permissionsBoundary'),
        user_id=__ret__.get('userId'),
        user_name=__ret__.get('userName'))

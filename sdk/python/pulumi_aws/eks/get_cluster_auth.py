# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetClusterAuthResult:
    """
    A collection of values returned by getClusterAuth.
    """
    def __init__(__self__, name=None, token=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if token and not isinstance(token, str):
            raise TypeError("Expected argument 'token' to be a str")
        __self__.token = token
        """
        The token to use to authenticate with the cluster.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_cluster_auth(name=None,opts=None):
    """
    Get an authentication token to communicate with an EKS cluster.
    
    Uses IAM credentials from the AWS provider to generate a temporary token that is compatible with
    [AWS IAM Authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator) authentication.
    This can be used to authenticate to an EKS cluster or to a cluster that has the AWS IAM Authenticator
    server configured.
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('aws:eks/getClusterAuth:getClusterAuth', __args__, opts=opts)

    return GetClusterAuthResult(
        name=__ret__.get('name'),
        token=__ret__.get('token'),
        id=__ret__.get('id'))

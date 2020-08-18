# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetUserPoolsResult',
    'AwaitableGetUserPoolsResult',
    'get_user_pools',
]


class GetUserPoolsResult:
    """
    A collection of values returned by getUserPools.
    """
    def __init__(__self__, arns=None, id=None, ids=None, name=None):
        if arns and not isinstance(arns, list):
            raise TypeError("Expected argument 'arns' to be a list")
        __self__.arns = arns
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
        The list of cognito user pool ids.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name


class AwaitableGetUserPoolsResult(GetUserPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserPoolsResult(
            arns=self.arns,
            id=self.id,
            ids=self.ids,
            name=self.name)


def get_user_pools(name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserPoolsResult:
    """
    Use this data source to get a list of cognito user pools.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    selected_rest_api = aws.apigateway.get_rest_api(name=var["api_gateway_name"])
    selected_user_pools = aws.cognito.get_user_pools(name=var["cognito_user_pool_name"])
    cognito = aws.apigateway.Authorizer("cognito",
        type="COGNITO_USER_POOLS",
        rest_api=selected_rest_api.id,
        provider_arns=selected_user_pools.arns)
    ```


    :param str name: Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:cognito/getUserPools:getUserPools', __args__, opts=opts).value

    return AwaitableGetUserPoolsResult(
        arns=__ret__.get('arns'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name=__ret__.get('name'))

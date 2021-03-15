# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = [
    'GetCallerIdentityResult',
    'AwaitableGetCallerIdentityResult',
    'get_caller_identity',
]

@pulumi.output_type
class GetCallerIdentityResult:
    """
    A collection of values returned by getCallerIdentity.
    """
    def __init__(__self__, account_id=None, arn=None, id=None, user_id=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The AWS Account ID number of the account that owns or contains the calling entity.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The AWS ARN associated with the calling entity.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The unique identifier of the calling entity.
        """
        return pulumi.get(self, "user_id")


class AwaitableGetCallerIdentityResult(GetCallerIdentityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCallerIdentityResult(
            account_id=self.account_id,
            arn=self.arn,
            id=self.id,
            user_id=self.user_id)


def get_caller_identity(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCallerIdentityResult:
    """
    Use this data source to get the access to the effective Account ID, User ID, and ARN in
    which this provider is authorized.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_caller_identity()
    pulumi.export("accountId", current.account_id)
    pulumi.export("callerArn", current.arn)
    pulumi.export("callerUser", current.user_id)
    ```
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getCallerIdentity:getCallerIdentity', __args__, opts=opts, typ=GetCallerIdentityResult).value

    return AwaitableGetCallerIdentityResult(
        account_id=__ret__.account_id,
        arn=__ret__.arn,
        id=__ret__.id,
        user_id=__ret__.user_id)

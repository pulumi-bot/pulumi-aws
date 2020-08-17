# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetAccountAliasResult',
    'AwaitableGetAccountAliasResult',
    'get_account_alias',
]



@pulumi.output_type
class GetAccountAliasResult:
    """
    A collection of values returned by getAccountAlias.
    """
    def __init__(__self__, account_alias=None, id=None):
        if account_alias and not isinstance(account_alias, str):
            raise TypeError("Expected argument 'account_alias' to be a str")
        pulumi.set(__self__, "account_alias", account_alias)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> str:
        """
        The alias associated with the AWS account.
        """
        ...

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        ...



class AwaitableGetAccountAliasResult(GetAccountAliasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountAliasResult(
            account_alias=self.account_alias,
            id=self.id)


def get_account_alias(                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountAliasResult:
    """
    The IAM Account Alias data source allows access to the account alias
    for the effective account in which this provider is working.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.iam.get_account_alias()
    pulumi.export("accountId", current.account_alias)
    ```
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:iam/getAccountAlias:getAccountAlias', __args__, opts=opts, typ=GetAccountAliasResult).value

    return AwaitableGetAccountAliasResult(
        account_alias=__ret__.account_alias,
        id=__ret__.id)

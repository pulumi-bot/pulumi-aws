# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSiteResult',
    'AwaitableGetSiteResult',
    'get_site',
]

@pulumi.output_type
class GetSiteResult:
    """
    A collection of values returned by getSite.
    """
    def __init__(__self__, account_id=None, description=None, id=None, name=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        AWS Account identifier.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetSiteResult(GetSiteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSiteResult(
            account_id=self.account_id,
            description=self.description,
            id=self.id,
            name=self.name)


def get_site(id: Optional[str] = None,
             name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSiteResult:
    """
    Provides details about an Outposts Site.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.outposts.get_site(name="example")
    ```


    :param str id: Identifier of the Site.
    :param str name: Name of the Site.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:outposts/getSite:getSite', __args__, opts=opts, typ=GetSiteResult).value

    return AwaitableGetSiteResult(
        account_id=__ret__.account_id,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_site)
def get_site_apply(id: Optional[pulumi.Input[str]] = None,
                   name: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSiteResult]:
    ...

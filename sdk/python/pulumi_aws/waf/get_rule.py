# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetRuleResult:
    """
    A collection of values returned by getRule.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name


class AwaitableGetRuleResult(GetRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRuleResult(
            id=self.id,
            name=self.name)


def get_rule(name=None,opts=None):
    """
    `waf.Rule` Retrieves a WAF Rule Resource Id.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.waf.get_rule(name="tfWAFRule")
    ```


    :param str name: The name of the WAF rule.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:waf/getRule:getRule', __args__, opts=opts).value

    return AwaitableGetRuleResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'))

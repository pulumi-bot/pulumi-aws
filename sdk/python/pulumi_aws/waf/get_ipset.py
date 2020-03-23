# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetIpsetResult:
    """
    A collection of values returned by getIpset.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetIpsetResult(GetIpsetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpsetResult(
            id=self.id,
            name=self.name)

def get_ipset(name=None,opts=None):
    """
    `waf.IpSet` Retrieves a WAF IP Set Resource Id.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/waf_ipset.html.markdown.


    :param str name: The name of the WAF IP set.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:waf/getIpset:getIpset', __args__, opts=opts).value

    return AwaitableGetIpsetResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'))

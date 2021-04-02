# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetOriginRequestPolicyResult',
    'AwaitableGetOriginRequestPolicyResult',
    'get_origin_request_policy',
]

@pulumi.output_type
class GetOriginRequestPolicyResult:
    """
    A collection of values returned by getOriginRequestPolicy.
    """
    def __init__(__self__, comment=None, cookies_configs=None, etag=None, headers_configs=None, id=None, name=None, query_strings_configs=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if cookies_configs and not isinstance(cookies_configs, list):
            raise TypeError("Expected argument 'cookies_configs' to be a list")
        pulumi.set(__self__, "cookies_configs", cookies_configs)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if headers_configs and not isinstance(headers_configs, list):
            raise TypeError("Expected argument 'headers_configs' to be a list")
        pulumi.set(__self__, "headers_configs", headers_configs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if query_strings_configs and not isinstance(query_strings_configs, list):
            raise TypeError("Expected argument 'query_strings_configs' to be a list")
        pulumi.set(__self__, "query_strings_configs", query_strings_configs)

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment to describe the origin request policy.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="cookiesConfigs")
    def cookies_configs(self) -> Sequence['outputs.GetOriginRequestPolicyCookiesConfigResult']:
        """
        Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
        """
        return pulumi.get(self, "cookies_configs")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        The current version of the origin request policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="headersConfigs")
    def headers_configs(self) -> Sequence['outputs.GetOriginRequestPolicyHeadersConfigResult']:
        """
        Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
        """
        return pulumi.get(self, "headers_configs")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryStringsConfigs")
    def query_strings_configs(self) -> Sequence['outputs.GetOriginRequestPolicyQueryStringsConfigResult']:
        """
        Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query Strings Config for more information.
        """
        return pulumi.get(self, "query_strings_configs")


class AwaitableGetOriginRequestPolicyResult(GetOriginRequestPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOriginRequestPolicyResult(
            comment=self.comment,
            cookies_configs=self.cookies_configs,
            etag=self.etag,
            headers_configs=self.headers_configs,
            id=self.id,
            name=self.name,
            query_strings_configs=self.query_strings_configs)


def get_origin_request_policy(id: Optional[str] = None,
                              name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOriginRequestPolicyResult:
    """
    ## Example Usage

    The following example below creates a CloudFront origin request policy.

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.cloudfront.get_origin_request_policy(name="example-policy")
    ```


    :param str id: The identifier for the origin request policy.
    :param str name: Unique name to identify the origin request policy.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:cloudfront/getOriginRequestPolicy:getOriginRequestPolicy', __args__, opts=opts, typ=GetOriginRequestPolicyResult).value

    return AwaitableGetOriginRequestPolicyResult(
        comment=__ret__.comment,
        cookies_configs=__ret__.cookies_configs,
        etag=__ret__.etag,
        headers_configs=__ret__.headers_configs,
        id=__ret__.id,
        name=__ret__.name,
        query_strings_configs=__ret__.query_strings_configs)

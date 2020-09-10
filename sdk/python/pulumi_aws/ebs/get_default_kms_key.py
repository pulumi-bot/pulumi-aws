# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetDefaultKmsKeyResult',
    'AwaitableGetDefaultKmsKeyResult',
    'get_default_kms_key',
]

@pulumi.output_type
class GetDefaultKmsKeyResult:
    """
    A collection of values returned by getDefaultKmsKey.
    """
    def __init__(__self__, id=None, key_arn=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_arn and not isinstance(key_arn, str):
            raise TypeError("Expected argument 'key_arn' to be a str")
        pulumi.set(__self__, "key_arn", key_arn)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> str:
        return pulumi.get(self, "key_arn")


class AwaitableGetDefaultKmsKeyResult(GetDefaultKmsKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultKmsKeyResult(
            id=self.id,
            key_arn=self.key_arn)


def get_default_kms_key(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultKmsKeyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getDefaultKmsKey:getDefaultKmsKey', __args__, opts=opts, typ=GetDefaultKmsKeyResult).value

    return AwaitableGetDefaultKmsKeyResult(
        id=__ret__.id,
        key_arn=__ret__.key_arn)

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'GetPartitionResult',
    'AwaitableGetPartitionResult',
    'get_partition',
]


class GetPartitionResult:
    """
    A collection of values returned by getPartition.
    """
    def __init__(__self__, dns_suffix=None, id=None, partition=None):
        if dns_suffix and not isinstance(dns_suffix, str):
            raise TypeError("Expected argument 'dns_suffix' to be a str")
        __self__.dns_suffix = dns_suffix
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        __self__.partition = partition


class AwaitableGetPartitionResult(GetPartitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPartitionResult(
            dns_suffix=self.dns_suffix,
            id=self.id,
            partition=self.partition)


def get_partition(                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPartitionResult:
    """
    Use this data source to lookup current AWS partition in which this provider is working

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_partition()
    s3_policy = aws.iam.get_policy_document(statements=[{
        "actions": ["s3:ListBucket"],
        "resources": [f"arn:{current.partition}:s3:::my-bucket"],
        "sid": "1",
    }])
    ```
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getPartition:getPartition', __args__, opts=opts).value

    return AwaitableGetPartitionResult(
        dns_suffix=__ret__.get('dnsSuffix'),
        id=__ret__.get('id'),
        partition=__ret__.get('partition'))

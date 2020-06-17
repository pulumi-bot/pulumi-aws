# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDefaultKmsKeyResult:
    """
    A collection of values returned by getDefaultKmsKey.
    """
    def __init__(__self__, id=None, key_arn=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if key_arn and not isinstance(key_arn, str):
            raise TypeError("Expected argument 'key_arn' to be a str")
        __self__.key_arn = key_arn
        """
        Amazon Resource Name (ARN) of the default KMS key uses to encrypt an EBS volume in this region when no key is specified in an API call that creates the volume and encryption by default is enabled.
        """
class AwaitableGetDefaultKmsKeyResult(GetDefaultKmsKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultKmsKeyResult(
            id=self.id,
            key_arn=self.key_arn)

def get_default_kms_key(opts=None):
    """
    Use this data source to get the default EBS encryption KMS key in the current region.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.ebs.get_default_kms_key()
    example = aws.ebs.Volume("example",
        availability_zone="us-west-2a",
        encrypted=True,
        kms_key_id=current.key_arn)
    ```
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getDefaultKmsKey:getDefaultKmsKey', __args__, opts=opts).value

    return AwaitableGetDefaultKmsKeyResult(
        id=__ret__.get('id'),
        key_arn=__ret__.get('keyArn'))

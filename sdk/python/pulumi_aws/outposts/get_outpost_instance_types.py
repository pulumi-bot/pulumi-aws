# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetOutpostInstanceTypesResult:
    """
    A collection of values returned by getOutpostInstanceTypes.
    """
    def __init__(__self__, arn=None, id=None, instance_types=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instance_types and not isinstance(instance_types, list):
            raise TypeError("Expected argument 'instance_types' to be a list")
        __self__.instance_types = instance_types
        """
        Set of instance types.
        """


class AwaitableGetOutpostInstanceTypesResult(GetOutpostInstanceTypesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOutpostInstanceTypesResult(
            arn=self.arn,
            id=self.id,
            instance_types=self.instance_types)


def get_outpost_instance_types(arn=None,opts=None):
    """
    Information about Outposts Instance Types.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.outposts.get_outpost_instance_types(arn=data["aws_outposts_outpost"]["example"]["arn"])
    ```


    :param str arn: Outpost Amazon Resource Name (ARN).
    """
    __args__ = dict()

    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:outposts/getOutpostInstanceTypes:getOutpostInstanceTypes', __args__, opts=opts).value

    return AwaitableGetOutpostInstanceTypesResult(
        arn=__ret__.get('arn'),
        id=__ret__.get('id'),
        instance_types=__ret__.get('instanceTypes'))

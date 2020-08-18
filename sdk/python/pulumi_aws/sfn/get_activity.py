# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetActivityResult',
    'AwaitableGetActivityResult',
    'get_activity',
]



@pulumi.output_type
class GetActivityResult:
    """
    A collection of values returned by getActivity.
    """
    def __init__(__self__, arn=None, creation_date=None, id=None, name=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The date the activity was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")



class AwaitableGetActivityResult(GetActivityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActivityResult(
            arn=self.arn,
            creation_date=self.creation_date,
            id=self.id,
            name=self.name)


def get_activity(arn: Optional[str] = None,
                 name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActivityResult:
    """
    Provides a Step Functions Activity data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    sfn_activity = aws.sfn.get_activity(name="my-activity")
    ```


    :param str arn: The Amazon Resource Name (ARN) that identifies the activity.
    :param str name: The name that identifies the activity.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:sfn/getActivity:getActivity', __args__, opts=opts, typ=GetActivityResult).value

    return AwaitableGetActivityResult(
        arn=__ret__.arn,
        creation_date=__ret__.creation_date,
        id=__ret__.id,
        name=__ret__.name)

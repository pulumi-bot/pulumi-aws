# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetEventCategoriesResult',
    'AwaitableGetEventCategoriesResult',
    'get_event_categories',
]

@pulumi.output_type
class GetEventCategoriesResult:
    """
    A collection of values returned by getEventCategories.
    """
    def __init__(__self__, event_categories=None, id=None, source_type=None):
        if event_categories and not isinstance(event_categories, list):
            raise TypeError("Expected argument 'event_categories' to be a list")
        pulumi.set(__self__, "event_categories", event_categories)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> Sequence[str]:
        """
        A list of the event categories.
        """
        return pulumi.get(self, "event_categories")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[str]:
        return pulumi.get(self, "source_type")


class AwaitableGetEventCategoriesResult(GetEventCategoriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventCategoriesResult(
            event_categories=self.event_categories,
            id=self.id,
            source_type=self.source_type)


def get_event_categories(source_type: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventCategoriesResult:
    """
    ## Example Usage

    List the event categories of all the RDS resources.

    ```python
    import pulumi
    import pulumi_aws as aws

    example_event_categories = aws.rds.get_event_categories()
    pulumi.export("example", example_event_categories.event_categories)
    ```

    List the event categories specific to the RDS resource `db-snapshot`.

    ```python
    import pulumi
    import pulumi_aws as aws

    example_event_categories = aws.rds.get_event_categories(source_type="db-snapshot")
    pulumi.export("example", example_event_categories.event_categories)
    ```


    :param str source_type: The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot.
    """
    __args__ = dict()
    __args__['sourceType'] = source_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:rds/getEventCategories:getEventCategories', __args__, opts=opts, typ=GetEventCategoriesResult).value

    return AwaitableGetEventCategoriesResult(
        event_categories=__ret__.event_categories,
        id=__ret__.id,
        source_type=__ret__.source_type)

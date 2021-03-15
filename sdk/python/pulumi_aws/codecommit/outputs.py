# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'TriggerTrigger',
]

@pulumi.output_type
class TriggerTrigger(dict):
    def __init__(__self__, *,
                 destination_arn: str,
                 events: Sequence[str],
                 name: str,
                 branches: Optional[Sequence[str]] = None,
                 custom_data: Optional[str] = None):
        """
        :param str destination_arn: The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
        :param Sequence[str] events: The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
        :param str name: The name of the trigger.
        :param Sequence[str] branches: The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
        :param str custom_data: Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
        """
        pulumi.set(__self__, "destination_arn", destination_arn)
        pulumi.set(__self__, "events", events)
        pulumi.set(__self__, "name", name)
        if branches is not None:
            pulumi.set(__self__, "branches", branches)
        if custom_data is not None:
            pulumi.set(__self__, "custom_data", custom_data)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> str:
        """
        The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def events(self) -> Sequence[str]:
        """
        The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the trigger.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def branches(self) -> Optional[Sequence[str]]:
        """
        The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
        """
        return pulumi.get(self, "branches")

    @property
    @pulumi.getter(name="customData")
    def custom_data(self) -> Optional[str]:
        """
        Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
        """
        return pulumi.get(self, "custom_data")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



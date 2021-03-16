# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'QueueReservationPlanSettings',
]

@pulumi.output_type
class QueueReservationPlanSettings(dict):
    def __init__(__self__, *,
                 commitment: str,
                 renewal_type: str,
                 reserved_slots: int):
        """
        :param str commitment: The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
        :param str renewal_type: Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
        :param int reserved_slots: Specifies the number of reserved transcode slots (RTS) for queue.
        """
        pulumi.set(__self__, "commitment", commitment)
        pulumi.set(__self__, "renewal_type", renewal_type)
        pulumi.set(__self__, "reserved_slots", reserved_slots)

    @property
    @pulumi.getter
    def commitment(self) -> str:
        """
        The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
        """
        return pulumi.get(self, "commitment")

    @property
    @pulumi.getter(name="renewalType")
    def renewal_type(self) -> str:
        """
        Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
        """
        return pulumi.get(self, "renewal_type")

    @property
    @pulumi.getter(name="reservedSlots")
    def reserved_slots(self) -> int:
        """
        Specifies the number of reserved transcode slots (RTS) for queue.
        """
        return pulumi.get(self, "reserved_slots")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ClusterParameterGroupParameterArgs',
]

@pulumi.input_type
class ClusterParameterGroupParameterArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str],
                 apply_method: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of the documentDB parameter.
        :param pulumi.Input[str] value: The value of the documentDB parameter.
        :param pulumi.Input[str] apply_method: Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if apply_method is not None:
            pulumi.set(__self__, "apply_method", apply_method)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the documentDB parameter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the documentDB parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="applyMethod")
    def apply_method(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        """
        return pulumi.get(self, "apply_method")

    @apply_method.setter
    def apply_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apply_method", value)



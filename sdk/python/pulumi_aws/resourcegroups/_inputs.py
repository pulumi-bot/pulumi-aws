# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GroupResourceQueryArgs',
]

@pulumi.input_type
class GroupResourceQueryArgs:
    def __init__(__self__, *,
                 query: pulumi.Input[str],
                 type: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "query", query)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def query(self) -> pulumi.Input[str]:
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: pulumi.Input[str]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)



# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GroupResourceQuery',
]

@pulumi.output_type
class GroupResourceQuery(dict):
    @property
    @pulumi.getter
    def query(self) -> str:
        """
        The resource query as a JSON string.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the resource query. Defaults to `TAG_FILTERS_1_0`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



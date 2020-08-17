# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ServerEndpointDetails',
]

@pulumi.output_type
class ServerEndpointDetails(dict):
    def __init__(__self__, *,
                 vpc_endpoint_id: str):
        """
        :param str vpc_endpoint_id: The ID of the VPC endpoint.
        """
        pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> str:
        """
        The ID of the VPC endpoint.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



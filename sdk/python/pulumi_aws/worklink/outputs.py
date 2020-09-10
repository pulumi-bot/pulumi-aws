# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'FleetIdentityProvider',
    'FleetNetwork',
]

@pulumi.output_type
class FleetIdentityProvider(dict):
    def __init__(__self__, *,
                 saml_metadata: str,
                 type: str):
        pulumi.set(__self__, "saml_metadata", saml_metadata)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="samlMetadata")
    def saml_metadata(self) -> str:
        return pulumi.get(self, "saml_metadata")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FleetNetwork(dict):
    def __init__(__self__, *,
                 security_group_ids: List[str],
                 subnet_ids: List[str],
                 vpc_id: str):
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> List[str]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



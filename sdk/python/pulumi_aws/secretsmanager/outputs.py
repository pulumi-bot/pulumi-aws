# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'SecretRotationRotationRules',
    'SecretRotationRules',
    'GetSecretRotationRotationRuleResult',
    'GetSecretRotationRuleResult',
]

@pulumi.output_type
class SecretRotationRotationRules(dict):
    def __init__(__self__, *,
                 automatically_after_days: float):
        """
        :param float automatically_after_days: Specifies the number of days between automatic scheduled rotations of the secret.
        """
        pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> float:
        """
        Specifies the number of days between automatic scheduled rotations of the secret.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SecretRotationRules(dict):
    def __init__(__self__, *,
                 automatically_after_days: float):
        """
        :param float automatically_after_days: Specifies the number of days between automatic scheduled rotations of the secret.
        """
        pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> float:
        """
        Specifies the number of days between automatic scheduled rotations of the secret.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetSecretRotationRotationRuleResult(dict):
    def __init__(__self__, *,
                 automatically_after_days: float):
        pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> float:
        ...


@pulumi.output_type
class GetSecretRotationRuleResult(dict):
    def __init__(__self__, *,
                 automatically_after_days: float):
        pulumi.set(__self__, "automatically_after_days", automatically_after_days)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> float:
        ...



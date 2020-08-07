# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'BrokerConfiguration',
    'BrokerEncryptionOptions',
    'BrokerInstance',
    'BrokerLogs',
    'BrokerMaintenanceWindowStartTime',
    'BrokerUser',
    'GetBrokerConfiguration',
    'GetBrokerEncryptionOption',
    'GetBrokerInstance',
    'GetBrokerLogs',
    'GetBrokerMaintenanceWindowStartTime',
    'GetBrokerUser',
]

@pulumi.output_type
class BrokerConfiguration(dict):
    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The Configuration ID.
        """
        ...

    @property
    @pulumi.getter
    def revision(self) -> Optional[float]:
        """
        Revision of the Configuration.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BrokerEncryptionOptions(dict):
    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS managed CMKs or customer managed CMKs are in use, this value must be configured.
        """
        ...

    @property
    @pulumi.getter(name="useAwsOwnedKey")
    def use_aws_owned_key(self) -> Optional[bool]:
        """
        Boolean to enable an AWS owned Key Management Service (KMS) Customer Master Key (CMK) that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS managed Customer Master Key (CMK) aliased to `aws/mq` in your account.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BrokerInstance(dict):
    @property
    @pulumi.getter(name="consoleUrl")
    def console_url(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[List[str]]:
        ...

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BrokerLogs(dict):
    @property
    @pulumi.getter
    def audit(self) -> Optional[bool]:
        """
        Enables audit logging. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        """
        ...

    @property
    @pulumi.getter
    def general(self) -> Optional[bool]:
        """
        Enables general logging via CloudWatch. Defaults to `false`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BrokerMaintenanceWindowStartTime(dict):
    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> str:
        """
        The day of the week. e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`
        """
        ...

    @property
    @pulumi.getter(name="timeOfDay")
    def time_of_day(self) -> str:
        """
        The time, in 24-hour format. e.g. `02:00`
        """
        ...

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        """
        The time zone, UTC by default, in either the Country/City format, or the UTC offset format. e.g. `CET`
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BrokerUser(dict):
    @property
    @pulumi.getter(name="consoleAccess")
    def console_access(self) -> Optional[bool]:
        """
        Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
        """
        ...

    @property
    @pulumi.getter
    def groups(self) -> Optional[List[str]]:
        """
        The list of groups (20 maximum) to which the ActiveMQ user belongs.
        """
        ...

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        """
        ...

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username of the user.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBrokerConfiguration(dict):
    @property
    @pulumi.getter
    def id(self) -> str:
        ...

    @property
    @pulumi.getter
    def revision(self) -> float:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBrokerEncryptionOption(dict):
    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        ...

    @property
    @pulumi.getter(name="useAwsOwnedKey")
    def use_aws_owned_key(self) -> bool:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBrokerInstance(dict):
    @property
    @pulumi.getter(name="consoleUrl")
    def console_url(self) -> str:
        ...

    @property
    @pulumi.getter
    def endpoints(self) -> List[str]:
        ...

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBrokerLogs(dict):
    @property
    @pulumi.getter
    def audit(self) -> bool:
        ...

    @property
    @pulumi.getter
    def general(self) -> bool:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBrokerMaintenanceWindowStartTime(dict):
    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> str:
        ...

    @property
    @pulumi.getter(name="timeOfDay")
    def time_of_day(self) -> str:
        ...

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetBrokerUser(dict):
    @property
    @pulumi.getter(name="consoleAccess")
    def console_access(self) -> bool:
        ...

    @property
    @pulumi.getter
    def groups(self) -> List[str]:
        ...

    @property
    @pulumi.getter
    def username(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



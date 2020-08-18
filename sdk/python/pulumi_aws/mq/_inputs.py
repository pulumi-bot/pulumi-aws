# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'BrokerConfigurationArgs',
    'BrokerEncryptionOptionsArgs',
    'BrokerInstanceArgs',
    'BrokerLogsArgs',
    'BrokerMaintenanceWindowStartTimeArgs',
    'BrokerUserArgs',
    'GetBrokerLogsArgs',
]

@pulumi.input_type
class BrokerConfigurationArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 revision: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[str] id: The Configuration ID.
        :param pulumi.Input[float] revision: Revision of the Configuration.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "revision", revision)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The Configuration ID.
        """
        ...

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[float]]:
        """
        Revision of the Configuration.
        """
        ...

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[float]]):
        ...


@pulumi.input_type
class BrokerEncryptionOptionsArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 use_aws_owned_key: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] kms_key_id: Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS managed CMKs or customer managed CMKs are in use, this value must be configured.
        :param pulumi.Input[bool] use_aws_owned_key: Boolean to enable an AWS owned Key Management Service (KMS) Customer Master Key (CMK) that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS managed Customer Master Key (CMK) aliased to `aws/mq` in your account.
        """
        pulumi.set(__self__, "kmsKeyId", kms_key_id)
        pulumi.set(__self__, "useAwsOwnedKey", use_aws_owned_key)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS managed CMKs or customer managed CMKs are in use, this value must be configured.
        """
        ...

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="useAwsOwnedKey")
    def use_aws_owned_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean to enable an AWS owned Key Management Service (KMS) Customer Master Key (CMK) that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS managed Customer Master Key (CMK) aliased to `aws/mq` in your account.
        """
        ...

    @use_aws_owned_key.setter
    def use_aws_owned_key(self, value: Optional[pulumi.Input[bool]]):
        ...


@pulumi.input_type
class BrokerInstanceArgs:
    def __init__(__self__, *,
                 console_url: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "consoleUrl", console_url)
        pulumi.set(__self__, "endpoints", endpoints)
        pulumi.set(__self__, "ipAddress", ip_address)

    @property
    @pulumi.getter(name="consoleUrl")
    def console_url(self) -> Optional[pulumi.Input[str]]:
        ...

    @console_url.setter
    def console_url(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        ...

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        ...

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        ...

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class BrokerLogsArgs:
    def __init__(__self__, *,
                 audit: Optional[pulumi.Input[bool]] = None,
                 general: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] audit: Enables audit logging. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        :param pulumi.Input[bool] general: Enables general logging via CloudWatch. Defaults to `false`.
        """
        pulumi.set(__self__, "audit", audit)
        pulumi.set(__self__, "general", general)

    @property
    @pulumi.getter
    def audit(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables audit logging. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        """
        ...

    @audit.setter
    def audit(self, value: Optional[pulumi.Input[bool]]):
        ...

    @property
    @pulumi.getter
    def general(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables general logging via CloudWatch. Defaults to `false`.
        """
        ...

    @general.setter
    def general(self, value: Optional[pulumi.Input[bool]]):
        ...


@pulumi.input_type
class BrokerMaintenanceWindowStartTimeArgs:
    def __init__(__self__, *,
                 day_of_week: pulumi.Input[str],
                 time_of_day: pulumi.Input[str],
                 time_zone: pulumi.Input[str]):
        """
        :param pulumi.Input[str] day_of_week: The day of the week. e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`
        :param pulumi.Input[str] time_of_day: The time, in 24-hour format. e.g. `02:00`
        :param pulumi.Input[str] time_zone: The time zone, UTC by default, in either the Country/City format, or the UTC offset format. e.g. `CET`
        """
        pulumi.set(__self__, "dayOfWeek", day_of_week)
        pulumi.set(__self__, "timeOfDay", time_of_day)
        pulumi.set(__self__, "timeZone", time_zone)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> pulumi.Input[str]:
        """
        The day of the week. e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`
        """
        ...

    @day_of_week.setter
    def day_of_week(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="timeOfDay")
    def time_of_day(self) -> pulumi.Input[str]:
        """
        The time, in 24-hour format. e.g. `02:00`
        """
        ...

    @time_of_day.setter
    def time_of_day(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Input[str]:
        """
        The time zone, UTC by default, in either the Country/City format, or the UTC offset format. e.g. `CET`
        """
        ...

    @time_zone.setter
    def time_zone(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class BrokerUserArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 console_access: Optional[pulumi.Input[bool]] = None,
                 groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] password: The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        :param pulumi.Input[str] username: The username of the user.
        :param pulumi.Input[bool] console_access: Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
        :param pulumi.Input[List[pulumi.Input[str]]] groups: The list of groups (20 maximum) to which the ActiveMQ user belongs.
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        pulumi.set(__self__, "consoleAccess", console_access)
        pulumi.set(__self__, "groups", groups)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        """
        ...

    @password.setter
    def password(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The username of the user.
        """
        ...

    @username.setter
    def username(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="consoleAccess")
    def console_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
        """
        ...

    @console_access.setter
    def console_access(self, value: Optional[pulumi.Input[bool]]):
        ...

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The list of groups (20 maximum) to which the ActiveMQ user belongs.
        """
        ...

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class GetBrokerLogsArgs:
    def __init__(__self__, *,
                 audit: bool,
                 general: bool):
        pulumi.set(__self__, "audit", audit)
        pulumi.set(__self__, "general", general)

    @property
    @pulumi.getter
    def audit(self) -> bool:
        ...

    @audit.setter
    def audit(self, value: bool):
        ...

    @property
    @pulumi.getter
    def general(self) -> bool:
        ...

    @general.setter
    def general(self, value: bool):
        ...



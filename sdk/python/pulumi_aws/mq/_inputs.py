# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'BrokerConfigurationArgs',
    'BrokerEncryptionOptionsArgs',
    'BrokerInstanceArgs',
    'BrokerLdapServerMetadataArgs',
    'BrokerLogsArgs',
    'BrokerMaintenanceWindowStartTimeArgs',
    'BrokerUserArgs',
    'GetBrokerLogsArgs',
]

@pulumi.input_type
class BrokerConfigurationArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 revision: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] id: The Configuration ID.
        :param pulumi.Input[int] revision: Revision of the Configuration.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The Configuration ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[int]]:
        """
        Revision of the Configuration.
        """
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision", value)


@pulumi.input_type
class BrokerEncryptionOptionsArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 use_aws_owned_key: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] kms_key_id: Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS-managed CMKs or customer-managed CMKs are in use, this value must be configured.
        :param pulumi.Input[bool] use_aws_owned_key: Whether to enable an AWS-owned KMS CMK that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS-managed CMK aliased to `aws/mq` in your account.
        """
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if use_aws_owned_key is not None:
            pulumi.set(__self__, "use_aws_owned_key", use_aws_owned_key)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS-managed CMKs or customer-managed CMKs are in use, this value must be configured.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="useAwsOwnedKey")
    def use_aws_owned_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable an AWS-owned KMS CMK that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS-managed CMK aliased to `aws/mq` in your account.
        """
        return pulumi.get(self, "use_aws_owned_key")

    @use_aws_owned_key.setter
    def use_aws_owned_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_aws_owned_key", value)


@pulumi.input_type
class BrokerInstanceArgs:
    def __init__(__self__, *,
                 console_url: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None):
        if console_url is not None:
            pulumi.set(__self__, "console_url", console_url)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)

    @property
    @pulumi.getter(name="consoleUrl")
    def console_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "console_url")

    @console_url.setter
    def console_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "console_url", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)


@pulumi.input_type
class BrokerLdapServerMetadataArgs:
    def __init__(__self__, *,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_base: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 role_search_matching: Optional[pulumi.Input[str]] = None,
                 role_search_subtree: Optional[pulumi.Input[bool]] = None,
                 service_account_password: Optional[pulumi.Input[str]] = None,
                 service_account_username: Optional[pulumi.Input[str]] = None,
                 user_base: Optional[pulumi.Input[str]] = None,
                 user_role_name: Optional[pulumi.Input[str]] = None,
                 user_search_matching: Optional[pulumi.Input[str]] = None,
                 user_search_subtree: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hosts: List of a fully qualified domain name of the LDAP server and an optional failover server.
        :param pulumi.Input[str] role_base: Fully qualified name of the directory to search for a user’s groups.
        :param pulumi.Input[str] role_name: Specifies the LDAP attribute that identifies the group name attribute in the object returned from the group membership query.
        :param pulumi.Input[str] role_search_matching: Search criteria for groups.
        :param pulumi.Input[bool] role_search_subtree: Whether the directory search scope is the entire sub-tree.
        :param pulumi.Input[str] service_account_password: Service account password.
        :param pulumi.Input[str] service_account_username: Service account username.
        :param pulumi.Input[str] user_base: Fully qualified name of the directory where you want to search for users.
        :param pulumi.Input[str] user_role_name: Specifies the name of the LDAP attribute for the user group membership.
        :param pulumi.Input[str] user_search_matching: Search criteria for users.
        :param pulumi.Input[bool] user_search_subtree: Whether the directory search scope is the entire sub-tree.
        """
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if role_base is not None:
            pulumi.set(__self__, "role_base", role_base)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if role_search_matching is not None:
            pulumi.set(__self__, "role_search_matching", role_search_matching)
        if role_search_subtree is not None:
            pulumi.set(__self__, "role_search_subtree", role_search_subtree)
        if service_account_password is not None:
            pulumi.set(__self__, "service_account_password", service_account_password)
        if service_account_username is not None:
            pulumi.set(__self__, "service_account_username", service_account_username)
        if user_base is not None:
            pulumi.set(__self__, "user_base", user_base)
        if user_role_name is not None:
            pulumi.set(__self__, "user_role_name", user_role_name)
        if user_search_matching is not None:
            pulumi.set(__self__, "user_search_matching", user_search_matching)
        if user_search_subtree is not None:
            pulumi.set(__self__, "user_search_subtree", user_search_subtree)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of a fully qualified domain name of the LDAP server and an optional failover server.
        """
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter(name="roleBase")
    def role_base(self) -> Optional[pulumi.Input[str]]:
        """
        Fully qualified name of the directory to search for a user’s groups.
        """
        return pulumi.get(self, "role_base")

    @role_base.setter
    def role_base(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_base", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the LDAP attribute that identifies the group name attribute in the object returned from the group membership query.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="roleSearchMatching")
    def role_search_matching(self) -> Optional[pulumi.Input[str]]:
        """
        Search criteria for groups.
        """
        return pulumi.get(self, "role_search_matching")

    @role_search_matching.setter
    def role_search_matching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_search_matching", value)

    @property
    @pulumi.getter(name="roleSearchSubtree")
    def role_search_subtree(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the directory search scope is the entire sub-tree.
        """
        return pulumi.get(self, "role_search_subtree")

    @role_search_subtree.setter
    def role_search_subtree(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "role_search_subtree", value)

    @property
    @pulumi.getter(name="serviceAccountPassword")
    def service_account_password(self) -> Optional[pulumi.Input[str]]:
        """
        Service account password.
        """
        return pulumi.get(self, "service_account_password")

    @service_account_password.setter
    def service_account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_password", value)

    @property
    @pulumi.getter(name="serviceAccountUsername")
    def service_account_username(self) -> Optional[pulumi.Input[str]]:
        """
        Service account username.
        """
        return pulumi.get(self, "service_account_username")

    @service_account_username.setter
    def service_account_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_username", value)

    @property
    @pulumi.getter(name="userBase")
    def user_base(self) -> Optional[pulumi.Input[str]]:
        """
        Fully qualified name of the directory where you want to search for users.
        """
        return pulumi.get(self, "user_base")

    @user_base.setter
    def user_base(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_base", value)

    @property
    @pulumi.getter(name="userRoleName")
    def user_role_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the LDAP attribute for the user group membership.
        """
        return pulumi.get(self, "user_role_name")

    @user_role_name.setter
    def user_role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_role_name", value)

    @property
    @pulumi.getter(name="userSearchMatching")
    def user_search_matching(self) -> Optional[pulumi.Input[str]]:
        """
        Search criteria for users.
        """
        return pulumi.get(self, "user_search_matching")

    @user_search_matching.setter
    def user_search_matching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_search_matching", value)

    @property
    @pulumi.getter(name="userSearchSubtree")
    def user_search_subtree(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the directory search scope is the entire sub-tree.
        """
        return pulumi.get(self, "user_search_subtree")

    @user_search_subtree.setter
    def user_search_subtree(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "user_search_subtree", value)


@pulumi.input_type
class BrokerLogsArgs:
    def __init__(__self__, *,
                 audit: Optional[pulumi.Input[bool]] = None,
                 general: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] audit: Enables audit logging. Auditing is only possible for `engine_type` of `ActiveMQ`. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        :param pulumi.Input[bool] general: Enables general logging via CloudWatch. Defaults to `false`.
        """
        if audit is not None:
            pulumi.set(__self__, "audit", audit)
        if general is not None:
            pulumi.set(__self__, "general", general)

    @property
    @pulumi.getter
    def audit(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables audit logging. Auditing is only possible for `engine_type` of `ActiveMQ`. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        """
        return pulumi.get(self, "audit")

    @audit.setter
    def audit(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "audit", value)

    @property
    @pulumi.getter
    def general(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables general logging via CloudWatch. Defaults to `false`.
        """
        return pulumi.get(self, "general")

    @general.setter
    def general(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "general", value)


@pulumi.input_type
class BrokerMaintenanceWindowStartTimeArgs:
    def __init__(__self__, *,
                 day_of_week: pulumi.Input[str],
                 time_of_day: pulumi.Input[str],
                 time_zone: pulumi.Input[str]):
        """
        :param pulumi.Input[str] day_of_week: Day of the week, e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`.
        :param pulumi.Input[str] time_of_day: Time, in 24-hour format, e.g. `02:00`.
        :param pulumi.Input[str] time_zone: Time zone in either the Country/City format or the UTC offset format, e.g. `CET`.
        """
        pulumi.set(__self__, "day_of_week", day_of_week)
        pulumi.set(__self__, "time_of_day", time_of_day)
        pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> pulumi.Input[str]:
        """
        Day of the week, e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`.
        """
        return pulumi.get(self, "day_of_week")

    @day_of_week.setter
    def day_of_week(self, value: pulumi.Input[str]):
        pulumi.set(self, "day_of_week", value)

    @property
    @pulumi.getter(name="timeOfDay")
    def time_of_day(self) -> pulumi.Input[str]:
        """
        Time, in 24-hour format, e.g. `02:00`.
        """
        return pulumi.get(self, "time_of_day")

    @time_of_day.setter
    def time_of_day(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_of_day", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Input[str]:
        """
        Time zone in either the Country/City format or the UTC offset format, e.g. `CET`.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class BrokerUserArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 console_access: Optional[pulumi.Input[bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] password: Password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        :param pulumi.Input[str] username: Username of the user.
        :param pulumi.Input[bool] console_access: Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user. Applies to `engine_type` of `ActiveMQ` only.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to `engine_type` of `ActiveMQ` only.
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if console_access is not None:
            pulumi.set(__self__, "console_access", console_access)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        Username of the user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="consoleAccess")
    def console_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user. Applies to `engine_type` of `ActiveMQ` only.
        """
        return pulumi.get(self, "console_access")

    @console_access.setter
    def console_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "console_access", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to `engine_type` of `ActiveMQ` only.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)


@pulumi.input_type
class GetBrokerLogsArgs:
    def __init__(__self__, *,
                 audit: pulumi.Input[bool],
                 general: pulumi.Input[bool]):
        pulumi.set(__self__, "audit", audit)
        pulumi.set(__self__, "general", general)

    @property
    @pulumi.getter
    def audit(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "audit")

    @audit.setter
    def audit(self, value: pulumi.Input[bool]):
        pulumi.set(self, "audit", value)

    @property
    @pulumi.getter
    def general(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "general")

    @general.setter
    def general(self, value: pulumi.Input[bool]):
        pulumi.set(self, "general", value)



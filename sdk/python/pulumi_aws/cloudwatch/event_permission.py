# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['EventPermissionArgs', 'EventPermission']

@pulumi.input_type
class EventPermissionArgs:
    def __init__(__self__, *,
                 principal: pulumi.Input[str],
                 statement_id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['EventPermissionConditionArgs']] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EventPermission resource.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input['EventPermissionConditionArgs'] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] event_bus_name: The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        """
        pulumi.set(__self__, "principal", principal)
        pulumi.set(__self__, "statement_id", statement_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if event_bus_name is not None:
            pulumi.set(__self__, "event_bus_name", event_bus_name)

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Input[str]:
        """
        The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> pulumi.Input[str]:
        """
        An identifier string for the external account that you are granting permissions to.
        """
        return pulumi.get(self, "statement_id")

    @statement_id.setter
    def statement_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "statement_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['EventPermissionConditionArgs']]:
        """
        Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['EventPermissionConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> Optional[pulumi.Input[str]]:
        """
        The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        """
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_name", value)


@pulumi.input_type
class _EventPermissionState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['EventPermissionConditionArgs']] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EventPermission resources.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input['EventPermissionConditionArgs'] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] event_bus_name: The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if event_bus_name is not None:
            pulumi.set(__self__, "event_bus_name", event_bus_name)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if statement_id is not None:
            pulumi.set(__self__, "statement_id", statement_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['EventPermissionConditionArgs']]:
        """
        Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['EventPermissionConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> Optional[pulumi.Input[str]]:
        """
        The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        """
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_name", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        """
        The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> Optional[pulumi.Input[str]]:
        """
        An identifier string for the external account that you are granting permissions to.
        """
        return pulumi.get(self, "statement_id")

    @statement_id.setter
    def statement_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "statement_id", value)


class EventPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['EventPermissionConditionArgs']]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage
        ### Account Access

        ```python
        import pulumi
        import pulumi_aws as aws

        dev_account_access = aws.cloudwatch.EventPermission("devAccountAccess",
            principal="123456789012",
            statement_id="DevAccountAccess")
        ```
        ### Organization Access

        ```python
        import pulumi
        import pulumi_aws as aws

        organization_access = aws.cloudwatch.EventPermission("organizationAccess",
            principal="*",
            statement_id="OrganizationAccess",
            condition=aws.cloudwatch.EventPermissionConditionArgs(
                key="aws:PrincipalOrgID",
                type="StringEquals",
                value=aws_organizations_organization["example"]["id"],
            ))
        ```

        ## Import

        EventBridge permissions can be imported using the `event_bus_name/statement_id` (if you omit `event_bus_name`, the `default` event bus will be used), e.g.

        ```sh
         $ pulumi import aws:cloudwatch/eventPermission:EventPermission DevAccountAccess example-event-bus/DevAccountAccess
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input[pulumi.InputType['EventPermissionConditionArgs']] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] event_bus_name: The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: EventPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage
        ### Account Access

        ```python
        import pulumi
        import pulumi_aws as aws

        dev_account_access = aws.cloudwatch.EventPermission("devAccountAccess",
            principal="123456789012",
            statement_id="DevAccountAccess")
        ```
        ### Organization Access

        ```python
        import pulumi
        import pulumi_aws as aws

        organization_access = aws.cloudwatch.EventPermission("organizationAccess",
            principal="*",
            statement_id="OrganizationAccess",
            condition=aws.cloudwatch.EventPermissionConditionArgs(
                key="aws:PrincipalOrgID",
                type="StringEquals",
                value=aws_organizations_organization["example"]["id"],
            ))
        ```

        ## Import

        EventBridge permissions can be imported using the `event_bus_name/statement_id` (if you omit `event_bus_name`, the `default` event bus will be used), e.g.

        ```sh
         $ pulumi import aws:cloudwatch/eventPermission:EventPermission DevAccountAccess example-event-bus/DevAccountAccess
        ```

        :param str resource_name_: The name of the resource.
        :param EventPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['EventPermissionConditionArgs']]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventPermissionArgs.__new__(EventPermissionArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["condition"] = condition
            __props__.__dict__["event_bus_name"] = event_bus_name
            if principal is None and not opts.urn:
                raise TypeError("Missing required property 'principal'")
            __props__.__dict__["principal"] = principal
            if statement_id is None and not opts.urn:
                raise TypeError("Missing required property 'statement_id'")
            __props__.__dict__["statement_id"] = statement_id
        super(EventPermission, __self__).__init__(
            'aws:cloudwatch/eventPermission:EventPermission',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['EventPermissionConditionArgs']]] = None,
            event_bus_name: Optional[pulumi.Input[str]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            statement_id: Optional[pulumi.Input[str]] = None) -> 'EventPermission':
        """
        Get an existing EventPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input[pulumi.InputType['EventPermissionConditionArgs']] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] event_bus_name: The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventPermissionState.__new__(_EventPermissionState)

        __props__.__dict__["action"] = action
        __props__.__dict__["condition"] = condition
        __props__.__dict__["event_bus_name"] = event_bus_name
        __props__.__dict__["principal"] = principal
        __props__.__dict__["statement_id"] = statement_id
        return EventPermission(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[str]]:
        """
        The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.EventPermissionCondition']]:
        """
        Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> pulumi.Output[Optional[str]]:
        """
        The event bus to set the permissions on. If you omit this, the permissions are set on the `default` event bus.
        """
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[str]:
        """
        The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> pulumi.Output[str]:
        """
        An identifier string for the external account that you are granting permissions to.
        """
        return pulumi.get(self, "statement_id")


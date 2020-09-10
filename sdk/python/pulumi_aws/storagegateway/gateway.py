# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Gateway']


class Gateway(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activation_key: Optional[pulumi.Input[str]] = None,
                 cloudwatch_log_group_arn: Optional[pulumi.Input[str]] = None,
                 gateway_ip_address: Optional[pulumi.Input[str]] = None,
                 gateway_name: Optional[pulumi.Input[str]] = None,
                 gateway_timezone: Optional[pulumi.Input[str]] = None,
                 gateway_type: Optional[pulumi.Input[str]] = None,
                 gateway_vpc_endpoint: Optional[pulumi.Input[str]] = None,
                 medium_changer_type: Optional[pulumi.Input[str]] = None,
                 smb_active_directory_settings: Optional[pulumi.Input[pulumi.InputType['GatewaySmbActiveDirectorySettingsArgs']]] = None,
                 smb_guest_password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tape_drive_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Gateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['activation_key'] = activation_key
            __props__['cloudwatch_log_group_arn'] = cloudwatch_log_group_arn
            __props__['gateway_ip_address'] = gateway_ip_address
            if gateway_name is None:
                raise TypeError("Missing required property 'gateway_name'")
            __props__['gateway_name'] = gateway_name
            if gateway_timezone is None:
                raise TypeError("Missing required property 'gateway_timezone'")
            __props__['gateway_timezone'] = gateway_timezone
            __props__['gateway_type'] = gateway_type
            __props__['gateway_vpc_endpoint'] = gateway_vpc_endpoint
            __props__['medium_changer_type'] = medium_changer_type
            __props__['smb_active_directory_settings'] = smb_active_directory_settings
            __props__['smb_guest_password'] = smb_guest_password
            __props__['tags'] = tags
            __props__['tape_drive_type'] = tape_drive_type
            __props__['arn'] = None
            __props__['gateway_id'] = None
        super(Gateway, __self__).__init__(
            'aws:storagegateway/gateway:Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activation_key: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cloudwatch_log_group_arn: Optional[pulumi.Input[str]] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            gateway_ip_address: Optional[pulumi.Input[str]] = None,
            gateway_name: Optional[pulumi.Input[str]] = None,
            gateway_timezone: Optional[pulumi.Input[str]] = None,
            gateway_type: Optional[pulumi.Input[str]] = None,
            gateway_vpc_endpoint: Optional[pulumi.Input[str]] = None,
            medium_changer_type: Optional[pulumi.Input[str]] = None,
            smb_active_directory_settings: Optional[pulumi.Input[pulumi.InputType['GatewaySmbActiveDirectorySettingsArgs']]] = None,
            smb_guest_password: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tape_drive_type: Optional[pulumi.Input[str]] = None) -> 'Gateway':
        """
        Get an existing Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["activation_key"] = activation_key
        __props__["arn"] = arn
        __props__["cloudwatch_log_group_arn"] = cloudwatch_log_group_arn
        __props__["gateway_id"] = gateway_id
        __props__["gateway_ip_address"] = gateway_ip_address
        __props__["gateway_name"] = gateway_name
        __props__["gateway_timezone"] = gateway_timezone
        __props__["gateway_type"] = gateway_type
        __props__["gateway_vpc_endpoint"] = gateway_vpc_endpoint
        __props__["medium_changer_type"] = medium_changer_type
        __props__["smb_active_directory_settings"] = smb_active_directory_settings
        __props__["smb_guest_password"] = smb_guest_password
        __props__["tags"] = tags
        __props__["tape_drive_type"] = tape_drive_type
        return Gateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activationKey")
    def activation_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "activation_key")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cloudwatchLogGroupArn")
    def cloudwatch_log_group_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cloudwatch_log_group_arn")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="gatewayIpAddress")
    def gateway_ip_address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway_ip_address")

    @property
    @pulumi.getter(name="gatewayName")
    def gateway_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway_name")

    @property
    @pulumi.getter(name="gatewayTimezone")
    def gateway_timezone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway_timezone")

    @property
    @pulumi.getter(name="gatewayType")
    def gateway_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "gateway_type")

    @property
    @pulumi.getter(name="gatewayVpcEndpoint")
    def gateway_vpc_endpoint(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "gateway_vpc_endpoint")

    @property
    @pulumi.getter(name="mediumChangerType")
    def medium_changer_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "medium_changer_type")

    @property
    @pulumi.getter(name="smbActiveDirectorySettings")
    def smb_active_directory_settings(self) -> pulumi.Output[Optional['outputs.GatewaySmbActiveDirectorySettings']]:
        return pulumi.get(self, "smb_active_directory_settings")

    @property
    @pulumi.getter(name="smbGuestPassword")
    def smb_guest_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "smb_guest_password")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tapeDriveType")
    def tape_drive_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tape_drive_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


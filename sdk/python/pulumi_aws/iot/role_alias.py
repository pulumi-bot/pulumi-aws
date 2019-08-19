# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RoleAlias(pulumi.CustomResource):
    alias: pulumi.Output[str]
    """
    The name of the role alias.
    """
    arn: pulumi.Output[str]
    """
    The ARN assigned by AWS to this role alias.
    """
    credential_duration: pulumi.Output[float]
    """
    The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
    """
    role_arn: pulumi.Output[str]
    """
    The identity of the role to which the alias refers.
    """
    def __init__(__self__, resource_name, opts=None, alias=None, credential_duration=None, role_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an IoT role alias.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The name of the role alias.
        :param pulumi.Input[float] credential_duration: The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
        :param pulumi.Input[str] role_arn: The identity of the role to which the alias refers.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_role_alias.html.markdown.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if alias is None:
                raise TypeError("Missing required property 'alias'")
            __props__['alias'] = alias
            __props__['credential_duration'] = credential_duration
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['arn'] = None
        super(RoleAlias, __self__).__init__(
            'aws:iot/roleAlias:RoleAlias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alias=None, arn=None, credential_duration=None, role_arn=None):
        """
        Get an existing RoleAlias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The name of the role alias.
        :param pulumi.Input[str] arn: The ARN assigned by AWS to this role alias.
        :param pulumi.Input[float] credential_duration: The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
        :param pulumi.Input[str] role_arn: The identity of the role to which the alias refers.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_role_alias.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["alias"] = alias
        __props__["arn"] = arn
        __props__["credential_duration"] = credential_duration
        __props__["role_arn"] = role_arn
        return RoleAlias(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


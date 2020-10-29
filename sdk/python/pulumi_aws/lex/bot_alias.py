# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BotAlias']


class BotAlias(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bot_name: Optional[pulumi.Input[str]] = None,
                 bot_version: Optional[pulumi.Input[str]] = None,
                 conversation_logs: Optional[pulumi.Input[pulumi.InputType['BotAliasConversationLogsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Amazon Lex Bot Alias resource. For more information see
        [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bot_name: The name of the bot.
        :param pulumi.Input[str] bot_version: The name of the bot.
        :param pulumi.Input[pulumi.InputType['BotAliasConversationLogsArgs']] conversation_logs: The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
        :param pulumi.Input[str] description: A description of the alias.
        :param pulumi.Input[str] name: The name of the alias. The name is not case sensitive.
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

            if bot_name is None:
                raise TypeError("Missing required property 'bot_name'")
            __props__['bot_name'] = bot_name
            if bot_version is None:
                raise TypeError("Missing required property 'bot_version'")
            __props__['bot_version'] = bot_version
            __props__['conversation_logs'] = conversation_logs
            __props__['description'] = description
            __props__['name'] = name
            __props__['arn'] = None
            __props__['checksum'] = None
            __props__['created_date'] = None
            __props__['last_updated_date'] = None
        super(BotAlias, __self__).__init__(
            'aws:lex/botAlias:BotAlias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bot_name: Optional[pulumi.Input[str]] = None,
            bot_version: Optional[pulumi.Input[str]] = None,
            checksum: Optional[pulumi.Input[str]] = None,
            conversation_logs: Optional[pulumi.Input[pulumi.InputType['BotAliasConversationLogsArgs']]] = None,
            created_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            last_updated_date: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'BotAlias':
        """
        Get an existing BotAlias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the bot alias.
        :param pulumi.Input[str] bot_name: The name of the bot.
        :param pulumi.Input[str] bot_version: The name of the bot.
        :param pulumi.Input[str] checksum: Checksum of the bot alias.
        :param pulumi.Input[pulumi.InputType['BotAliasConversationLogsArgs']] conversation_logs: The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
        :param pulumi.Input[str] created_date: The date that the bot alias was created.
        :param pulumi.Input[str] description: A description of the alias.
        :param pulumi.Input[str] last_updated_date: The date that the bot alias was updated. When you create a resource, the creation date and the last updated date are the same.
        :param pulumi.Input[str] name: The name of the alias. The name is not case sensitive.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["bot_name"] = bot_name
        __props__["bot_version"] = bot_version
        __props__["checksum"] = checksum
        __props__["conversation_logs"] = conversation_logs
        __props__["created_date"] = created_date
        __props__["description"] = description
        __props__["last_updated_date"] = last_updated_date
        __props__["name"] = name
        return BotAlias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the bot alias.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="botName")
    def bot_name(self) -> pulumi.Output[str]:
        """
        The name of the bot.
        """
        return pulumi.get(self, "bot_name")

    @property
    @pulumi.getter(name="botVersion")
    def bot_version(self) -> pulumi.Output[str]:
        """
        The name of the bot.
        """
        return pulumi.get(self, "bot_version")

    @property
    @pulumi.getter
    def checksum(self) -> pulumi.Output[str]:
        """
        Checksum of the bot alias.
        """
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter(name="conversationLogs")
    def conversation_logs(self) -> pulumi.Output[Optional['outputs.BotAliasConversationLogs']]:
        """
        The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
        """
        return pulumi.get(self, "conversation_logs")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[str]:
        """
        The date that the bot alias was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the alias.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastUpdatedDate")
    def last_updated_date(self) -> pulumi.Output[str]:
        """
        The date that the bot alias was updated. When you create a resource, the creation date and the last updated date are the same.
        """
        return pulumi.get(self, "last_updated_date")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the alias. The name is not case sensitive.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


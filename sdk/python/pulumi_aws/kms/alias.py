# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Alias(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the key alias.
    """
    name: pulumi.Output[str]
    """
    The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
    """
    name_prefix: pulumi.Output[str]
    """
    Creates an unique alias beginning with the specified prefix.
    The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
    """
    target_key_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the target key identifier.
    """
    target_key_id: pulumi.Output[str]
    """
    Identifier for the key for which the alias is for, can be either an ARN or key_id.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, name_prefix=None, target_key_id=None):
        """
        Provides an alias for a KMS customer master key. AWS Console enforces 1-to-1 mapping between aliases & keys,
        but API (hence Terraform too) allows you to create as many aliases as
        the [account limits](http://docs.aws.amazon.com/kms/latest/developerguide/limits.html) allow you.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] name: The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
        :param pulumi.Input[str] name_prefix: Creates an unique alias beginning with the specified prefix.
               The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
        :param pulumi.Input[str] target_key_id: Identifier for the key for which the alias is for, can be either an ARN or key_id.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        if target_key_id is None:
            raise TypeError('Missing required property target_key_id')
        __props__['target_key_id'] = target_key_id

        __props__['arn'] = None
        __props__['target_key_arn'] = None

        super(Alias, __self__).__init__(
            'aws:kms/alias:Alias',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


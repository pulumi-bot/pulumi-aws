# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Vault(pulumi.CustomResource):
    access_policy: pulumi.Output[str]
    """
    The policy document. This is a JSON formatted string.
    The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
    """
    arn: pulumi.Output[str]
    """
    The ARN of the vault.
    """
    location: pulumi.Output[str]
    """
    The URI of the vault that was created.
    """
    name: pulumi.Output[str]
    """
    The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
    """
    notifications: pulumi.Output[list]
    """
    The notifications for the Vault. Fields documented below.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, __name__, __opts__=None, access_policy=None, name=None, notifications=None, tags=None):
        """
        Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality
        
        > **NOTE:** When removing a Glacier Vault, the Vault must be empty.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
               The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        :param pulumi.Input[str] name: The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        :param pulumi.Input[list] notifications: The notifications for the Vault. Fields documented below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['access_policy'] = access_policy

        __props__['name'] = name

        __props__['notifications'] = notifications

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['location'] = None

        super(Vault, __self__).__init__(
            'aws:glacier/vault:Vault',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


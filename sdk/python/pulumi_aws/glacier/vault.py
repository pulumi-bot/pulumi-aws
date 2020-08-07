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

__all__ = ['Vault']


class Vault(pulumi.CustomResource):
    access_policy: pulumi.Output[Optional[str]] = pulumi.property("accessPolicy")
    """
    The policy document. This is a JSON formatted string.
    The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
    """

    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN of the vault.
    """

    location: pulumi.Output[str] = pulumi.property("location")
    """
    The URI of the vault that was created.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
    """

    notifications: pulumi.Output[Optional[List['outputs.VaultNotification']]] = pulumi.property("notifications")
    """
    The notifications for the Vault. Fields documented below.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the resource.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality

        > **NOTE:** When removing a Glacier Vault, the Vault must be empty.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        aws_sns_topic = aws.sns.Topic("awsSnsTopic")
        my_archive = aws.glacier.Vault("myArchive",
            access_policy=\"\"\"{
            "Version":"2012-10-17",
            "Statement":[
               {
                  "Sid": "add-read-only-perm",
                  "Principal": "*",
                  "Effect": "Allow",
                  "Action": [
                     "glacier:InitiateJob",
                     "glacier:GetJobOutput"
                  ],
                  "Resource": "arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive"
               }
            ]
        }

        \"\"\",
            notifications=[{
                "events": [
                    "ArchiveRetrievalCompleted",
                    "InventoryRetrievalCompleted",
                ],
                "sns_topic": aws_sns_topic.arn,
            }],
            tags={
                "Test": "MyArchive",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
               The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        :param pulumi.Input[str] name: The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]] notifications: The notifications for the Vault. Fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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

            __props__['access_policy'] = access_policy
            __props__['name'] = name
            __props__['notifications'] = notifications
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['location'] = None
        super(Vault, __self__).__init__(
            'aws:glacier/vault:Vault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            access_policy: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Vault':
        """
        Get an existing Vault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
               The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        :param pulumi.Input[str] arn: The ARN of the vault.
        :param pulumi.Input[str] location: The URI of the vault that was created.
        :param pulumi.Input[str] name: The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]] notifications: The notifications for the Vault. Fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_policy"] = access_policy
        __props__["arn"] = arn
        __props__["location"] = location
        __props__["name"] = name
        __props__["notifications"] = notifications
        __props__["tags"] = tags
        return Vault(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


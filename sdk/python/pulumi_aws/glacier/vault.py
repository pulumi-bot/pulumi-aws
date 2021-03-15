# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VaultArgs', 'Vault']

@pulumi.input_type
class VaultArgs:
    def __init__(__self__, *,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input['VaultNotificationArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Vault resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
               The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        :param pulumi.Input[str] name: The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        :param pulumi.Input[Sequence[pulumi.Input['VaultNotificationArgs']]] notifications: The notifications for the Vault. Fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        if access_policy is not None:
            pulumi.set(__self__, "access_policy", access_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessPolicy")
    def access_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The policy document. This is a JSON formatted string.
        The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        """
        return pulumi.get(self, "access_policy")

    @access_policy.setter
    def access_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VaultNotificationArgs']]]]:
        """
        The notifications for the Vault. Fields documented below.
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VaultNotificationArgs']]]]):
        pulumi.set(self, "notifications", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Vault(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VaultArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality

        > **NOTE:** When removing a Glacier Vault, the Vault must be empty.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        aws_sns_topic = aws.sns.Topic("awsSnsTopic")
        my_archive = aws.glacier.Vault("myArchive",
            notifications=[aws.glacier.VaultNotificationArgs(
                sns_topic=aws_sns_topic.arn,
                events=[
                    "ArchiveRetrievalCompleted",
                    "InventoryRetrievalCompleted",
                ],
            )],
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
            tags={
                "Test": "MyArchive",
            })
        ```

        ## Import

        Glacier Vaults can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:glacier/vault:Vault archive my_archive
        ```

        :param str resource_name: The name of the resource.
        :param VaultArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]]] = None,
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
            notifications=[aws.glacier.VaultNotificationArgs(
                sns_topic=aws_sns_topic.arn,
                events=[
                    "ArchiveRetrievalCompleted",
                    "InventoryRetrievalCompleted",
                ],
            )],
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
            tags={
                "Test": "MyArchive",
            })
        ```

        ## Import

        Glacier Vaults can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:glacier/vault:Vault archive my_archive
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
               The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        :param pulumi.Input[str] name: The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]] notifications: The notifications for the Vault. Fields documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VaultArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_policy: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Vault':
        """
        Get an existing Vault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
               The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        :param pulumi.Input[str] arn: The ARN of the vault.
        :param pulumi.Input[str] location: The URI of the vault that was created.
        :param pulumi.Input[str] name: The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VaultNotificationArgs']]]] notifications: The notifications for the Vault. Fields documented below.
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

    @property
    @pulumi.getter(name="accessPolicy")
    def access_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The policy document. This is a JSON formatted string.
        The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
        """
        return pulumi.get(self, "access_policy")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the vault.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The URI of the vault that was created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output[Optional[Sequence['outputs.VaultNotification']]]:
        """
        The notifications for the Vault. Fields documented below.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


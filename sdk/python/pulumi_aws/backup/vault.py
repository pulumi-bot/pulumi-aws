# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['VaultArgs', 'Vault']

@pulumi.input_type
class VaultArgs:
    def __init__(__self__, *,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Vault resource.
        :param pulumi.Input[str] kms_key_arn: The server-side encryption key that is used to protect your backups.
        :param pulumi.Input[str] name: Name of the backup vault to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the resources that you create.
        """
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The server-side encryption key that is used to protect your backups.
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the backup vault to create.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Metadata that you can assign to help organize the resources that you create.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Vault(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS Backup vault resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.backup.Vault("example", kms_key_arn=aws_kms_key["example"]["arn"])
        ```

        ## Import

        Backup vault can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:backup/vault:Vault test-vault TestVault
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_arn: The server-side encryption key that is used to protect your backups.
        :param pulumi.Input[str] name: Name of the backup vault to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the resources that you create.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VaultArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Backup vault resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.backup.Vault("example", kms_key_arn=aws_kms_key["example"]["arn"])
        ```

        ## Import

        Backup vault can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:backup/vault:Vault test-vault TestVault
        ```

        :param str resource_name: The name of the resource.
        :param VaultArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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

            __props__['kms_key_arn'] = kms_key_arn
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['recovery_points'] = None
        super(Vault, __self__).__init__(
            'aws:backup/vault:Vault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            kms_key_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            recovery_points: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Vault':
        """
        Get an existing Vault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the vault.
        :param pulumi.Input[str] kms_key_arn: The server-side encryption key that is used to protect your backups.
        :param pulumi.Input[str] name: Name of the backup vault to create.
        :param pulumi.Input[int] recovery_points: The number of recovery points that are stored in a backup vault.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Metadata that you can assign to help organize the resources that you create.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["kms_key_arn"] = kms_key_arn
        __props__["name"] = name
        __props__["recovery_points"] = recovery_points
        __props__["tags"] = tags
        return Vault(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the vault.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[str]:
        """
        The server-side encryption key that is used to protect your backups.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the backup vault to create.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recoveryPoints")
    def recovery_points(self) -> pulumi.Output[int]:
        """
        The number of recovery points that are stored in a backup vault.
        """
        return pulumi.get(self, "recovery_points")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Metadata that you can assign to help organize the resources that you create.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


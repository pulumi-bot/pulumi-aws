# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountAliasArgs', 'AccountAlias']

@pulumi.input_type
class AccountAliasArgs:
    def __init__(__self__, *,
                 account_alias: pulumi.Input[str]):
        """
        The set of arguments for constructing a AccountAlias resource.
        :param pulumi.Input[str] account_alias: The account alias
        """
        pulumi.set(__self__, "account_alias", account_alias)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> pulumi.Input[str]:
        """
        The account alias
        """
        return pulumi.get(self, "account_alias")

    @account_alias.setter
    def account_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_alias", value)


@pulumi.input_type
class _AccountAliasState:
    def __init__(__self__, *,
                 account_alias: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccountAlias resources.
        :param pulumi.Input[str] account_alias: The account alias
        """
        if account_alias is not None:
            pulumi.set(__self__, "account_alias", account_alias)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The account alias
        """
        return pulumi.get(self, "account_alias")

    @account_alias.setter
    def account_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_alias", value)


class AccountAlias(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_alias: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        > **Note:** There is only a single account alias per AWS account.

        Manages the account alias for the AWS Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        alias = aws.iam.AccountAlias("alias", account_alias="my-account-alias")
        ```

        ## Import

        The current Account Alias can be imported using the `account_alias`, e.g.

        ```sh
         $ pulumi import aws:iam/accountAlias:AccountAlias alias my-account-alias
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_alias: The account alias
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: AccountAliasArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > **Note:** There is only a single account alias per AWS account.

        Manages the account alias for the AWS Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        alias = aws.iam.AccountAlias("alias", account_alias="my-account-alias")
        ```

        ## Import

        The current Account Alias can be imported using the `account_alias`, e.g.

        ```sh
         $ pulumi import aws:iam/accountAlias:AccountAlias alias my-account-alias
        ```

        :param str resource_name_: The name of the resource.
        :param AccountAliasArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountAliasArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_alias: Optional[pulumi.Input[str]] = None,
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
            __props__ = AccountAliasArgs.__new__(AccountAliasArgs)

            if account_alias is None and not opts.urn:
                raise TypeError("Missing required property 'account_alias'")
            __props__.__dict__["account_alias"] = account_alias
        super(AccountAlias, __self__).__init__(
            'aws:iam/accountAlias:AccountAlias',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_alias: Optional[pulumi.Input[str]] = None) -> 'AccountAlias':
        """
        Get an existing AccountAlias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_alias: The account alias
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountAliasState.__new__(_AccountAliasState)

        __props__.__dict__["account_alias"] = account_alias
        return AccountAlias(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> pulumi.Output[str]:
        """
        The account alias
        """
        return pulumi.get(self, "account_alias")


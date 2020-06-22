# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AccountPasswordPolicy(pulumi.CustomResource):
    allow_users_to_change_password: pulumi.Output[bool]
    """
    Whether to allow users to change their own password
    """
    expire_passwords: pulumi.Output[bool]
    """
    Indicates whether passwords in the account expire.
    Returns `true` if `max_password_age` contains a value greater than `0`.
    Returns `false` if it is `0` or _not present_.
    """
    hard_expiry: pulumi.Output[bool]
    """
    Whether users are prevented from setting a new password after their password has expired
    (i.e. require administrator reset)
    """
    max_password_age: pulumi.Output[float]
    """
    The number of days that an user password is valid.
    """
    minimum_password_length: pulumi.Output[float]
    """
    Minimum length to require for user passwords.
    """
    password_reuse_prevention: pulumi.Output[float]
    """
    The number of previous passwords that users are prevented from reusing.
    """
    require_lowercase_characters: pulumi.Output[bool]
    """
    Whether to require lowercase characters for user passwords.
    """
    require_numbers: pulumi.Output[bool]
    """
    Whether to require numbers for user passwords.
    """
    require_symbols: pulumi.Output[bool]
    """
    Whether to require symbols for user passwords.
    """
    require_uppercase_characters: pulumi.Output[bool]
    """
    Whether to require uppercase characters for user passwords.
    """
    def __init__(__self__, resource_name, opts=None, allow_users_to_change_password=None, hard_expiry=None, max_password_age=None, minimum_password_length=None, password_reuse_prevention=None, require_lowercase_characters=None, require_numbers=None, require_symbols=None, require_uppercase_characters=None, __props__=None, __name__=None, __opts__=None):
        """
        > **Note:** There is only a single policy allowed per AWS account. An existing policy will be lost when using this resource as an effect of this limitation.

        Manages Password Policy for the AWS Account.
        See more about [Account Password Policy](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html)
        in the official AWS docs.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        strict = aws.iam.AccountPasswordPolicy("strict",
            allow_users_to_change_password=True,
            minimum_password_length=8,
            require_lowercase_characters=True,
            require_numbers=True,
            require_symbols=True,
            require_uppercase_characters=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_users_to_change_password: Whether to allow users to change their own password
        :param pulumi.Input[bool] hard_expiry: Whether users are prevented from setting a new password after their password has expired
               (i.e. require administrator reset)
        :param pulumi.Input[float] max_password_age: The number of days that an user password is valid.
        :param pulumi.Input[float] minimum_password_length: Minimum length to require for user passwords.
        :param pulumi.Input[float] password_reuse_prevention: The number of previous passwords that users are prevented from reusing.
        :param pulumi.Input[bool] require_lowercase_characters: Whether to require lowercase characters for user passwords.
        :param pulumi.Input[bool] require_numbers: Whether to require numbers for user passwords.
        :param pulumi.Input[bool] require_symbols: Whether to require symbols for user passwords.
        :param pulumi.Input[bool] require_uppercase_characters: Whether to require uppercase characters for user passwords.
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

            __props__['allow_users_to_change_password'] = allow_users_to_change_password
            __props__['hard_expiry'] = hard_expiry
            __props__['max_password_age'] = max_password_age
            __props__['minimum_password_length'] = minimum_password_length
            __props__['password_reuse_prevention'] = password_reuse_prevention
            __props__['require_lowercase_characters'] = require_lowercase_characters
            __props__['require_numbers'] = require_numbers
            __props__['require_symbols'] = require_symbols
            __props__['require_uppercase_characters'] = require_uppercase_characters
            __props__['expire_passwords'] = None
        super(AccountPasswordPolicy, __self__).__init__(
            'aws:iam/accountPasswordPolicy:AccountPasswordPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allow_users_to_change_password=None, expire_passwords=None, hard_expiry=None, max_password_age=None, minimum_password_length=None, password_reuse_prevention=None, require_lowercase_characters=None, require_numbers=None, require_symbols=None, require_uppercase_characters=None):
        """
        Get an existing AccountPasswordPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_users_to_change_password: Whether to allow users to change their own password
        :param pulumi.Input[bool] expire_passwords: Indicates whether passwords in the account expire.
               Returns `true` if `max_password_age` contains a value greater than `0`.
               Returns `false` if it is `0` or _not present_.
        :param pulumi.Input[bool] hard_expiry: Whether users are prevented from setting a new password after their password has expired
               (i.e. require administrator reset)
        :param pulumi.Input[float] max_password_age: The number of days that an user password is valid.
        :param pulumi.Input[float] minimum_password_length: Minimum length to require for user passwords.
        :param pulumi.Input[float] password_reuse_prevention: The number of previous passwords that users are prevented from reusing.
        :param pulumi.Input[bool] require_lowercase_characters: Whether to require lowercase characters for user passwords.
        :param pulumi.Input[bool] require_numbers: Whether to require numbers for user passwords.
        :param pulumi.Input[bool] require_symbols: Whether to require symbols for user passwords.
        :param pulumi.Input[bool] require_uppercase_characters: Whether to require uppercase characters for user passwords.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_users_to_change_password"] = allow_users_to_change_password
        __props__["expire_passwords"] = expire_passwords
        __props__["hard_expiry"] = hard_expiry
        __props__["max_password_age"] = max_password_age
        __props__["minimum_password_length"] = minimum_password_length
        __props__["password_reuse_prevention"] = password_reuse_prevention
        __props__["require_lowercase_characters"] = require_lowercase_characters
        __props__["require_numbers"] = require_numbers
        __props__["require_symbols"] = require_symbols
        __props__["require_uppercase_characters"] = require_uppercase_characters
        return AccountPasswordPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['UserLoginProfile']


class UserLoginProfile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 password_length: Optional[pulumi.Input[int]] = None,
                 password_reset_required: Optional[pulumi.Input[bool]] = None,
                 pgp_key: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an IAM User Login Profile with limited support for password creation during this provider resource creation. Uses PGP to encrypt the password for safe transport to the user. PGP keys can be obtained from Keybase.

        > To reset an IAM User login password via this provider, you can use delete and recreate this resource or change any of the arguments.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] password_length: The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
        :param pulumi.Input[bool] password_reset_required: Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
        :param pulumi.Input[str] pgp_key: Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
        :param pulumi.Input[str] user: The IAM user's name.
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

            __props__['password_length'] = password_length
            __props__['password_reset_required'] = password_reset_required
            if pgp_key is None:
                raise TypeError("Missing required property 'pgp_key'")
            __props__['pgp_key'] = pgp_key
            if user is None:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
            __props__['encrypted_password'] = None
            __props__['key_fingerprint'] = None
        super(UserLoginProfile, __self__).__init__(
            'aws:iam/userLoginProfile:UserLoginProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            encrypted_password: Optional[pulumi.Input[str]] = None,
            key_fingerprint: Optional[pulumi.Input[str]] = None,
            password_length: Optional[pulumi.Input[int]] = None,
            password_reset_required: Optional[pulumi.Input[bool]] = None,
            pgp_key: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'UserLoginProfile':
        """
        Get an existing UserLoginProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] encrypted_password: The encrypted password, base64 encoded. Only available if password was handled on this provider resource creation, not import.
        :param pulumi.Input[str] key_fingerprint: The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
        :param pulumi.Input[int] password_length: The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
        :param pulumi.Input[bool] password_reset_required: Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
        :param pulumi.Input[str] pgp_key: Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
        :param pulumi.Input[str] user: The IAM user's name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["encrypted_password"] = encrypted_password
        __props__["key_fingerprint"] = key_fingerprint
        __props__["password_length"] = password_length
        __props__["password_reset_required"] = password_reset_required
        __props__["pgp_key"] = pgp_key
        __props__["user"] = user
        return UserLoginProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="encryptedPassword")
    def encrypted_password(self) -> pulumi.Output[str]:
        """
        The encrypted password, base64 encoded. Only available if password was handled on this provider resource creation, not import.
        """
        return pulumi.get(self, "encrypted_password")

    @property
    @pulumi.getter(name="keyFingerprint")
    def key_fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint of the PGP key used to encrypt the password. Only available if password was handled on this provider resource creation, not import.
        """
        return pulumi.get(self, "key_fingerprint")

    @property
    @pulumi.getter(name="passwordLength")
    def password_length(self) -> pulumi.Output[Optional[int]]:
        """
        The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
        """
        return pulumi.get(self, "password_length")

    @property
    @pulumi.getter(name="passwordResetRequired")
    def password_reset_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument.
        """
        return pulumi.get(self, "password_reset_required")

    @property
    @pulumi.getter(name="pgpKey")
    def pgp_key(self) -> pulumi.Output[str]:
        """
        Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Only applies on resource creation. Drift detection is not possible with this argument.
        """
        return pulumi.get(self, "pgp_key")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The IAM user's name.
        """
        return pulumi.get(self, "user")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


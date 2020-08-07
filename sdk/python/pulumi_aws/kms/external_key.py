# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ExternalKey']


class ExternalKey(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The Amazon Resource Name (ARN) of the key.
    """

    deletion_window_in_days: pulumi.Output[Optional[float]] = pulumi.property("deletionWindowInDays")
    """
    Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    Description of the key.
    """

    enabled: pulumi.Output[bool] = pulumi.property("enabled")
    """
    Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
    """

    expiration_model: pulumi.Output[str] = pulumi.property("expirationModel")
    """
    Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
    """

    key_material_base64: pulumi.Output[Optional[str]] = pulumi.property("keyMaterialBase64")
    """
    Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
    """

    key_state: pulumi.Output[str] = pulumi.property("keyState")
    """
    The state of the CMK.
    """

    key_usage: pulumi.Output[str] = pulumi.property("keyUsage")
    """
    The cryptographic operations for which you can use the CMK.
    """

    policy: pulumi.Output[str] = pulumi.property("policy")
    """
    A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A key-value map of tags to assign to the key.
    """

    valid_to: pulumi.Output[Optional[str]] = pulumi.property("validTo")
    """
    Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_window_in_days: Optional[pulumi.Input[float]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_material_base64: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 valid_to: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a KMS Customer Master Key that uses external key material. To instead manage a KMS Customer Master Key where AWS automatically generates and potentially rotates key material, see the `kms.Key` resource.

        > **Note:** All arguments including the key material will be stored in the raw state as plain-text.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.kms.ExternalKey("example", description="KMS EXTERNAL for AMI encryption")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] deletion_window_in_days: Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
        :param pulumi.Input[str] description: Description of the key.
        :param pulumi.Input[bool] enabled: Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
        :param pulumi.Input[str] key_material_base64: Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
        :param pulumi.Input[str] policy: A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A key-value map of tags to assign to the key.
        :param pulumi.Input[str] valid_to: Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
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

            __props__['deletion_window_in_days'] = deletion_window_in_days
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['key_material_base64'] = key_material_base64
            __props__['policy'] = policy
            __props__['tags'] = tags
            __props__['valid_to'] = valid_to
            __props__['arn'] = None
            __props__['expiration_model'] = None
            __props__['key_state'] = None
            __props__['key_usage'] = None
        super(ExternalKey, __self__).__init__(
            'aws:kms/externalKey:ExternalKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            deletion_window_in_days: Optional[pulumi.Input[float]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            expiration_model: Optional[pulumi.Input[str]] = None,
            key_material_base64: Optional[pulumi.Input[str]] = None,
            key_state: Optional[pulumi.Input[str]] = None,
            key_usage: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            valid_to: Optional[pulumi.Input[str]] = None) -> 'ExternalKey':
        """
        Get an existing ExternalKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the key.
        :param pulumi.Input[float] deletion_window_in_days: Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
        :param pulumi.Input[str] description: Description of the key.
        :param pulumi.Input[bool] enabled: Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
        :param pulumi.Input[str] expiration_model: Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
        :param pulumi.Input[str] key_material_base64: Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
        :param pulumi.Input[str] key_state: The state of the CMK.
        :param pulumi.Input[str] key_usage: The cryptographic operations for which you can use the CMK.
        :param pulumi.Input[str] policy: A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A key-value map of tags to assign to the key.
        :param pulumi.Input[str] valid_to: Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["deletion_window_in_days"] = deletion_window_in_days
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["expiration_model"] = expiration_model
        __props__["key_material_base64"] = key_material_base64
        __props__["key_state"] = key_state
        __props__["key_usage"] = key_usage
        __props__["policy"] = policy
        __props__["tags"] = tags
        __props__["valid_to"] = valid_to
        return ExternalKey(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['NotebookInstance']


class NotebookInstance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_code_repository: Optional[pulumi.Input[str]] = None,
                 direct_internet_access: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 lifecycle_config_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 root_access: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 volume_size: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Sagemaker Notebook Instance resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_code_repository: The Git repository associated with the notebook instance as its default code repository
        :param pulumi.Input[str] direct_internet_access: Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
        :param pulumi.Input[str] instance_type: The name of ML compute instance type.
        :param pulumi.Input[str] kms_key_id: The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        :param pulumi.Input[str] lifecycle_config_name: The name of a lifecycle configuration to associate with the notebook instance.
        :param pulumi.Input[str] name: The name of the notebook instance (must be unique).
        :param pulumi.Input[str] role_arn: The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        :param pulumi.Input[str] root_access: Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: The associated security groups.
        :param pulumi.Input[str] subnet_id: The VPC subnet ID.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[int] volume_size: The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
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

            __props__['default_code_repository'] = default_code_repository
            __props__['direct_internet_access'] = direct_internet_access
            if instance_type is None:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            __props__['kms_key_id'] = kms_key_id
            __props__['lifecycle_config_name'] = lifecycle_config_name
            __props__['name'] = name
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['root_access'] = root_access
            __props__['security_groups'] = security_groups
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['volume_size'] = volume_size
            __props__['arn'] = None
        super(NotebookInstance, __self__).__init__(
            'aws:sagemaker/notebookInstance:NotebookInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            default_code_repository: Optional[pulumi.Input[str]] = None,
            direct_internet_access: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            lifecycle_config_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            root_access: Optional[pulumi.Input[str]] = None,
            security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            volume_size: Optional[pulumi.Input[int]] = None) -> 'NotebookInstance':
        """
        Get an existing NotebookInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
        :param pulumi.Input[str] default_code_repository: The Git repository associated with the notebook instance as its default code repository
        :param pulumi.Input[str] direct_internet_access: Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
        :param pulumi.Input[str] instance_type: The name of ML compute instance type.
        :param pulumi.Input[str] kms_key_id: The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        :param pulumi.Input[str] lifecycle_config_name: The name of a lifecycle configuration to associate with the notebook instance.
        :param pulumi.Input[str] name: The name of the notebook instance (must be unique).
        :param pulumi.Input[str] role_arn: The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        :param pulumi.Input[str] root_access: Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: The associated security groups.
        :param pulumi.Input[str] subnet_id: The VPC subnet ID.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[int] volume_size: The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["default_code_repository"] = default_code_repository
        __props__["direct_internet_access"] = direct_internet_access
        __props__["instance_type"] = instance_type
        __props__["kms_key_id"] = kms_key_id
        __props__["lifecycle_config_name"] = lifecycle_config_name
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        __props__["root_access"] = root_access
        __props__["security_groups"] = security_groups
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["volume_size"] = volume_size
        return NotebookInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="defaultCodeRepository")
    def default_code_repository(self) -> pulumi.Output[Optional[str]]:
        """
        The Git repository associated with the notebook instance as its default code repository
        """
        return pulumi.get(self, "default_code_repository")

    @property
    @pulumi.getter(name="directInternetAccess")
    def direct_internet_access(self) -> pulumi.Output[Optional[str]]:
        """
        Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
        """
        return pulumi.get(self, "direct_internet_access")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The name of ML compute instance type.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="lifecycleConfigName")
    def lifecycle_config_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a lifecycle configuration to associate with the notebook instance.
        """
        return pulumi.get(self, "lifecycle_config_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the notebook instance (must be unique).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="rootAccess")
    def root_access(self) -> pulumi.Output[Optional[str]]:
        """
        Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
        """
        return pulumi.get(self, "root_access")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        The associated security groups.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        The VPC subnet ID.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> pulumi.Output[Optional[int]]:
        """
        The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
        """
        return pulumi.get(self, "volume_size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NotebookInstance(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
    """
    instance_type: pulumi.Output[str]
    """
    The name of ML compute instance type.
    """
    kms_key_id: pulumi.Output[str]
    """
    The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
    """
    lifecycle_config_name: pulumi.Output[str]
    """
    The name of a lifecycle configuration to associate with the notebook instance.
    """
    name: pulumi.Output[str]
    """
    The name of the notebook instance (must be unique).
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
    """
    security_groups: pulumi.Output[list]
    """
    The associated security groups.
    """
    subnet_id: pulumi.Output[str]
    """
    The VPC subnet ID.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, instance_type=None, kms_key_id=None, lifecycle_config_name=None, name=None, role_arn=None, security_groups=None, subnet_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Sagemaker Notebook Instance resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_type: The name of ML compute instance type.
        :param pulumi.Input[str] kms_key_id: The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        :param pulumi.Input[str] lifecycle_config_name: The name of a lifecycle configuration to associate with the notebook instance.
        :param pulumi.Input[str] name: The name of the notebook instance (must be unique).
        :param pulumi.Input[str] role_arn: The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        :param pulumi.Input[list] security_groups: The associated security groups.
        :param pulumi.Input[str] subnet_id: The VPC subnet ID.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_notebook_instance.html.markdown.
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

            if instance_type is None:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            __props__['kms_key_id'] = kms_key_id
            __props__['lifecycle_config_name'] = lifecycle_config_name
            __props__['name'] = name
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['security_groups'] = security_groups
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['arn'] = None
        super(NotebookInstance, __self__).__init__(
            'aws:sagemaker/notebookInstance:NotebookInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, instance_type=None, kms_key_id=None, lifecycle_config_name=None, name=None, role_arn=None, security_groups=None, subnet_id=None, tags=None):
        """
        Get an existing NotebookInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
        :param pulumi.Input[str] instance_type: The name of ML compute instance type.
        :param pulumi.Input[str] kms_key_id: The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        :param pulumi.Input[str] lifecycle_config_name: The name of a lifecycle configuration to associate with the notebook instance.
        :param pulumi.Input[str] name: The name of the notebook instance (must be unique).
        :param pulumi.Input[str] role_arn: The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        :param pulumi.Input[list] security_groups: The associated security groups.
        :param pulumi.Input[str] subnet_id: The VPC subnet ID.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_notebook_instance.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["instance_type"] = instance_type
        __props__["kms_key_id"] = kms_key_id
        __props__["lifecycle_config_name"] = lifecycle_config_name
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        __props__["security_groups"] = security_groups
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        return NotebookInstance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


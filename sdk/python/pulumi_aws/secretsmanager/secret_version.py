# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SecretVersion(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the secret.
    """
    secret_binary: pulumi.Output[str]
    """
    Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
    """
    secret_id: pulumi.Output[str]
    """
    Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
    """
    secret_string: pulumi.Output[str]
    """
    Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
    """
    version_id: pulumi.Output[str]
    """
    The unique identifier of the version of the secret.
    """
    version_stages: pulumi.Output[list]
    """
    Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
    """
    def __init__(__self__, resource_name, opts=None, secret_binary=None, secret_id=None, secret_string=None, version_stages=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage AWS Secrets Manager secret version including its secret value. To manage secret metadata, see the [`secretsmanager.Secret` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret.html).
        
        > **NOTE:** If the `AWSCURRENT` staging label is present on this version during resource deletion, that label cannot be removed and will be skipped to prevent errors when fully deleting the secret. That label will leave this secret version active even after the resource is deleted from this provider unless the secret itself is deleted. Move the `AWSCURRENT` staging label before or after deleting this resource from this provider to fully trigger version deprecation if necessary.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secret_binary: Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
        :param pulumi.Input[str] secret_id: Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        :param pulumi.Input[str] secret_string: Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
        :param pulumi.Input[list] version_stages: Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/secretsmanager_secret_version.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['secret_binary'] = secret_binary
        if secret_id is None:
            raise TypeError("Missing required property 'secret_id'")
        __props__['secret_id'] = secret_id
        __props__['secret_string'] = secret_string
        __props__['version_stages'] = version_stages
        __props__['arn'] = None
        __props__['version_id'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(SecretVersion, __self__).__init__(
            'aws:secretsmanager/secretVersion:SecretVersion',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


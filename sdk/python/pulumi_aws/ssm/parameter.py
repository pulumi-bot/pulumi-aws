# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Parameter(pulumi.CustomResource):
    allowed_pattern: pulumi.Output[str]
    """
    A regular expression used to validate the parameter value.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the parameter.
    """
    description: pulumi.Output[str]
    """
    The description of the parameter.
    """
    key_id: pulumi.Output[str]
    """
    The KMS key id or arn for encrypting a SecureString.
    """
    name: pulumi.Output[str]
    """
    The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
    """
    overwrite: pulumi.Output[bool]
    """
    Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the object.
    """
    tier: pulumi.Output[str]
    """
    The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
    """
    type: pulumi.Output[str]
    """
    The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
    """
    value: pulumi.Output[str]
    """
    The value of the parameter.
    """
    version: pulumi.Output[float]
    """
    The version of the parameter.
    """
    def __init__(__self__, resource_name, opts=None, allowed_pattern=None, arn=None, description=None, key_id=None, name=None, overwrite=None, tags=None, tier=None, type=None, value=None, __name__=None, __opts__=None):
        """
        Provides an SSM Parameter resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_pattern: A regular expression used to validate the parameter value.
        :param pulumi.Input[str] arn: The ARN of the parameter.
        :param pulumi.Input[str] description: The description of the parameter.
        :param pulumi.Input[str] key_id: The KMS key id or arn for encrypting a SecureString.
        :param pulumi.Input[str] name: The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        :param pulumi.Input[bool] overwrite: Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the object.
        :param pulumi.Input[str] tier: The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard` and `Advanced`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        :param pulumi.Input[str] type: The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        :param pulumi.Input[str] value: The value of the parameter.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_parameter.html.markdown.
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

        __props__['allowed_pattern'] = allowed_pattern
        __props__['arn'] = arn
        __props__['description'] = description
        __props__['key_id'] = key_id
        __props__['name'] = name
        __props__['overwrite'] = overwrite
        __props__['tags'] = tags
        __props__['tier'] = tier
        if type is None:
            raise TypeError("Missing required property 'type'")
        __props__['type'] = type
        if value is None:
            raise TypeError("Missing required property 'value'")
        __props__['value'] = value
        __props__['version'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Parameter, __self__).__init__(
            'aws:ssm/parameter:Parameter',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


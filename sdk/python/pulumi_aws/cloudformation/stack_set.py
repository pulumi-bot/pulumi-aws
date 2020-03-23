# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class StackSet(pulumi.CustomResource):
    administration_role_arn: pulumi.Output[str]
    """
    Amazon Resource Number (ARN) of the IAM Role in the administrator account.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the StackSet.
    """
    capabilities: pulumi.Output[list]
    """
    A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
    """
    description: pulumi.Output[str]
    """
    Description of the StackSet.
    """
    execution_role_name: pulumi.Output[str]
    """
    Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
    """
    name: pulumi.Output[str]
    """
    Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
    """
    parameters: pulumi.Output[dict]
    """
    Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
    """
    stack_set_id: pulumi.Output[str]
    """
    Unique identifier of the StackSet.
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
    """
    template_body: pulumi.Output[str]
    """
    String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
    """
    template_url: pulumi.Output[str]
    """
    String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
    """
    def __init__(__self__, resource_name, opts=None, administration_role_arn=None, capabilities=None, description=None, execution_role_name=None, name=None, parameters=None, tags=None, template_body=None, template_url=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a CloudFormation StackSet. StackSets allow CloudFormation templates to be easily deployed across multiple accounts and regions via StackSet Instances ([`cloudformation.StackSetInstance` resource](https://www.terraform.io/docs/providers/aws/r/cloudformation_stack_set_instance.html)). Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).

        > **NOTE:** All template parameters, including those with a `Default`, must be configured or ignored with the `lifecycle` configuration block `ignore_changes` argument.

        > **NOTE:** All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudformation_stack_set.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administration_role_arn: Amazon Resource Number (ARN) of the IAM Role in the administrator account.
        :param pulumi.Input[list] capabilities: A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
        :param pulumi.Input[str] description: Description of the StackSet.
        :param pulumi.Input[str] execution_role_name: Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
        :param pulumi.Input[str] name: Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
        :param pulumi.Input[dict] parameters: Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
        :param pulumi.Input[dict] tags: Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
        :param pulumi.Input[str] template_body: String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
        :param pulumi.Input[str] template_url: String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
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

            if administration_role_arn is None:
                raise TypeError("Missing required property 'administration_role_arn'")
            __props__['administration_role_arn'] = administration_role_arn
            __props__['capabilities'] = capabilities
            __props__['description'] = description
            __props__['execution_role_name'] = execution_role_name
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['tags'] = tags
            __props__['template_body'] = template_body
            __props__['template_url'] = template_url
            __props__['arn'] = None
            __props__['stack_set_id'] = None
        super(StackSet, __self__).__init__(
            'aws:cloudformation/stackSet:StackSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, administration_role_arn=None, arn=None, capabilities=None, description=None, execution_role_name=None, name=None, parameters=None, stack_set_id=None, tags=None, template_body=None, template_url=None):
        """
        Get an existing StackSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administration_role_arn: Amazon Resource Number (ARN) of the IAM Role in the administrator account.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the StackSet.
        :param pulumi.Input[list] capabilities: A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
        :param pulumi.Input[str] description: Description of the StackSet.
        :param pulumi.Input[str] execution_role_name: Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
        :param pulumi.Input[str] name: Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
        :param pulumi.Input[dict] parameters: Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
        :param pulumi.Input[str] stack_set_id: Unique identifier of the StackSet.
        :param pulumi.Input[dict] tags: Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
        :param pulumi.Input[str] template_body: String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
        :param pulumi.Input[str] template_url: String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["administration_role_arn"] = administration_role_arn
        __props__["arn"] = arn
        __props__["capabilities"] = capabilities
        __props__["description"] = description
        __props__["execution_role_name"] = execution_role_name
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["stack_set_id"] = stack_set_id
        __props__["tags"] = tags
        __props__["template_body"] = template_body
        __props__["template_url"] = template_url
        return StackSet(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


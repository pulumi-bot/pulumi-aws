# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ProvisionedConcurrencyConfig(pulumi.CustomResource):
    function_name: pulumi.Output[str]
    """
    Name or Amazon Resource Name (ARN) of the Lambda Function.
    """
    provisioned_concurrent_executions: pulumi.Output[float]
    """
    Amount of capacity to allocate. Must be greater than or equal to `1`.
    """
    qualifier: pulumi.Output[str]
    """
    Lambda Function version or Lambda Alias name.
    """
    def __init__(__self__, resource_name, opts=None, function_name=None, provisioned_concurrent_executions=None, qualifier=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Lambda Provisioned Concurrency Configuration.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_provisioned_concurrency_config.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_name: Name or Amazon Resource Name (ARN) of the Lambda Function.
        :param pulumi.Input[float] provisioned_concurrent_executions: Amount of capacity to allocate. Must be greater than or equal to `1`.
        :param pulumi.Input[str] qualifier: Lambda Function version or Lambda Alias name.
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

            if function_name is None:
                raise TypeError("Missing required property 'function_name'")
            __props__['function_name'] = function_name
            if provisioned_concurrent_executions is None:
                raise TypeError("Missing required property 'provisioned_concurrent_executions'")
            __props__['provisioned_concurrent_executions'] = provisioned_concurrent_executions
            if qualifier is None:
                raise TypeError("Missing required property 'qualifier'")
            __props__['qualifier'] = qualifier
        super(ProvisionedConcurrencyConfig, __self__).__init__(
            'aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, function_name=None, provisioned_concurrent_executions=None, qualifier=None):
        """
        Get an existing ProvisionedConcurrencyConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_name: Name or Amazon Resource Name (ARN) of the Lambda Function.
        :param pulumi.Input[float] provisioned_concurrent_executions: Amount of capacity to allocate. Must be greater than or equal to `1`.
        :param pulumi.Input[str] qualifier: Lambda Function version or Lambda Alias name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["function_name"] = function_name
        __props__["provisioned_concurrent_executions"] = provisioned_concurrent_executions
        __props__["qualifier"] = qualifier
        return ProvisionedConcurrencyConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


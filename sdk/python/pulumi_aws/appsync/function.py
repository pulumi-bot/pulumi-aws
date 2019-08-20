# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Function(pulumi.CustomResource):
    api_id: pulumi.Output[str]
    """
    The ID of the associated AppSync API.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the Function object.
    """
    data_source: pulumi.Output[str]
    """
    The Function DataSource name.
    """
    description: pulumi.Output[str]
    """
    The Function description.
    """
    function_id: pulumi.Output[str]
    """
    A unique ID representing the Function object.
    """
    function_version: pulumi.Output[str]
    """
    The version of the request mapping template. Currently the supported value is `2018-05-29`.
    """
    name: pulumi.Output[str]
    """
    The Function name. The function name does not have to be unique.
    """
    request_mapping_template: pulumi.Output[str]
    """
    The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
    """
    response_mapping_template: pulumi.Output[str]
    """
    The Function response mapping template.
    """
    def __init__(__self__, resource_name, opts=None, api_id=None, data_source=None, description=None, function_version=None, name=None, request_mapping_template=None, response_mapping_template=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AppSync Function.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The ID of the associated AppSync API.
        :param pulumi.Input[str] data_source: The Function DataSource name.
        :param pulumi.Input[str] description: The Function description.
        :param pulumi.Input[str] function_version: The version of the request mapping template. Currently the supported value is `2018-05-29`.
        :param pulumi.Input[str] name: The Function name. The function name does not have to be unique.
        :param pulumi.Input[str] request_mapping_template: The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        :param pulumi.Input[str] response_mapping_template: The Function response mapping template.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_function.html.markdown.
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

            if api_id is None:
                raise TypeError("Missing required property 'api_id'")
            __props__['api_id'] = api_id
            if data_source is None:
                raise TypeError("Missing required property 'data_source'")
            __props__['data_source'] = data_source
            __props__['description'] = description
            __props__['function_version'] = function_version
            __props__['name'] = name
            if request_mapping_template is None:
                raise TypeError("Missing required property 'request_mapping_template'")
            __props__['request_mapping_template'] = request_mapping_template
            if response_mapping_template is None:
                raise TypeError("Missing required property 'response_mapping_template'")
            __props__['response_mapping_template'] = response_mapping_template
            __props__['arn'] = None
            __props__['function_id'] = None
        super(Function, __self__).__init__(
            'aws:appsync/function:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_id=None, arn=None, data_source=None, description=None, function_id=None, function_version=None, name=None, request_mapping_template=None, response_mapping_template=None):
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The ID of the associated AppSync API.
        :param pulumi.Input[str] arn: The ARN of the Function object.
        :param pulumi.Input[str] data_source: The Function DataSource name.
        :param pulumi.Input[str] description: The Function description.
        :param pulumi.Input[str] function_id: A unique ID representing the Function object.
        :param pulumi.Input[str] function_version: The version of the request mapping template. Currently the supported value is `2018-05-29`.
        :param pulumi.Input[str] name: The Function name. The function name does not have to be unique.
        :param pulumi.Input[str] request_mapping_template: The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        :param pulumi.Input[str] response_mapping_template: The Function response mapping template.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_function.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["api_id"] = api_id
        __props__["arn"] = arn
        __props__["data_source"] = data_source
        __props__["description"] = description
        __props__["function_id"] = function_id
        __props__["function_version"] = function_version
        __props__["name"] = name
        __props__["request_mapping_template"] = request_mapping_template
        __props__["response_mapping_template"] = response_mapping_template
        return Function(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Resolver(pulumi.CustomResource):
    api_id: pulumi.Output[str]
    """
    The API ID for the GraphQL API.
    """
    arn: pulumi.Output[str]
    """
    The ARN
    """
    data_source: pulumi.Output[str]
    """
    The DataSource name.
    """
    field: pulumi.Output[str]
    """
    The field name from the schema defined in the GraphQL API.
    """
    kind: pulumi.Output[str]
    """
    The resolver type. Valid values are `UNIT` and `PIPELINE`.
    """
    pipeline_config: pulumi.Output[dict]
    """
    The PipelineConfig. A `pipeline_config` block is documented below.
    """
    request_template: pulumi.Output[str]
    """
    The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
    """
    response_template: pulumi.Output[str]
    """
    The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
    """
    type: pulumi.Output[str]
    """
    The type name from the schema defined in the GraphQL API.
    """
    def __init__(__self__, resource_name, opts=None, api_id=None, data_source=None, field=None, kind=None, pipeline_config=None, request_template=None, response_template=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AppSync Resolver.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API ID for the GraphQL API.
        :param pulumi.Input[str] data_source: The DataSource name.
        :param pulumi.Input[str] field: The field name from the schema defined in the GraphQL API.
        :param pulumi.Input[str] kind: The resolver type. Valid values are `UNIT` and `PIPELINE`.
        :param pulumi.Input[dict] pipeline_config: The PipelineConfig. A `pipeline_config` block is documented below.
        :param pulumi.Input[str] request_template: The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
        :param pulumi.Input[str] response_template: The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
        :param pulumi.Input[str] type: The type name from the schema defined in the GraphQL API.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_resolver.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            if api_id is None:
                raise TypeError("Missing required property 'api_id'")
            __props__['api_id'] = api_id
            __props__['data_source'] = data_source
            if field is None:
                raise TypeError("Missing required property 'field'")
            __props__['field'] = field
            __props__['kind'] = kind
            __props__['pipeline_config'] = pipeline_config
            if request_template is None:
                raise TypeError("Missing required property 'request_template'")
            __props__['request_template'] = request_template
            if response_template is None:
                raise TypeError("Missing required property 'response_template'")
            __props__['response_template'] = response_template
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['arn'] = None
        super(Resolver, __self__).__init__(
            'aws:appsync/resolver:Resolver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_id=None, arn=None, data_source=None, field=None, kind=None, pipeline_config=None, request_template=None, response_template=None, type=None):
        """
        Get an existing Resolver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API ID for the GraphQL API.
        :param pulumi.Input[str] arn: The ARN
        :param pulumi.Input[str] data_source: The DataSource name.
        :param pulumi.Input[str] field: The field name from the schema defined in the GraphQL API.
        :param pulumi.Input[str] kind: The resolver type. Valid values are `UNIT` and `PIPELINE`.
        :param pulumi.Input[dict] pipeline_config: The PipelineConfig. A `pipeline_config` block is documented below.
        :param pulumi.Input[str] request_template: The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
        :param pulumi.Input[str] response_template: The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
        :param pulumi.Input[str] type: The type name from the schema defined in the GraphQL API.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_resolver.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["api_id"] = api_id
        __props__["arn"] = arn
        __props__["data_source"] = data_source
        __props__["field"] = field
        __props__["kind"] = kind
        __props__["pipeline_config"] = pipeline_config
        __props__["request_template"] = request_template
        __props__["response_template"] = response_template
        __props__["type"] = type
        return Resolver(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


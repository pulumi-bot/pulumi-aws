# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NamedQuery(pulumi.CustomResource):
    database: pulumi.Output[str]
    """
    The database to which the query belongs.
    """
    description: pulumi.Output[str]
    """
    A brief explanation of the query. Maximum length of 1024.
    """
    name: pulumi.Output[str]
    """
    The plain language name for the query. Maximum length of 128.
    """
    query: pulumi.Output[str]
    """
    The text of the query itself. In other words, all query statements. Maximum length of 262144.
    """
    workgroup: pulumi.Output[str]
    """
    The workgroup to which the query belongs. Defaults to `primary`
    """
    def __init__(__self__, resource_name, opts=None, database=None, description=None, name=None, query=None, workgroup=None, __name__=None, __opts__=None):
        """
        Provides an Athena Named Query resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to which the query belongs.
        :param pulumi.Input[str] description: A brief explanation of the query. Maximum length of 1024.
        :param pulumi.Input[str] name: The plain language name for the query. Maximum length of 128.
        :param pulumi.Input[str] query: The text of the query itself. In other words, all query statements. Maximum length of 262144.
        :param pulumi.Input[str] workgroup: The workgroup to which the query belongs. Defaults to `primary`

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/athena_named_query.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if database is None:
            raise TypeError("Missing required property 'database'")
        __props__['database'] = database

        __props__['description'] = description

        __props__['name'] = name

        if query is None:
            raise TypeError("Missing required property 'query'")
        __props__['query'] = query

        __props__['workgroup'] = workgroup

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(NamedQuery, __self__).__init__(
            'aws:athena/namedQuery:NamedQuery',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


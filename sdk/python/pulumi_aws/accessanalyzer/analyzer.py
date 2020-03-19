# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Analyzer(pulumi.CustomResource):
    analyzer_name: pulumi.Output[str]
    """
    Name of the Analyzer.
    """
    arn: pulumi.Output[str]
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags.
    """
    type: pulumi.Output[str]
    """
    Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
    """
    def __init__(__self__, resource_name, opts=None, analyzer_name=None, tags=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/accessanalyzer_analyzer.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] analyzer_name: Name of the Analyzer.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags.
        :param pulumi.Input[str] type: Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
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

            if analyzer_name is None:
                raise TypeError("Missing required property 'analyzer_name'")
            __props__['analyzer_name'] = analyzer_name
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['arn'] = None
        super(Analyzer, __self__).__init__(
            'aws:accessanalyzer/analyzer:Analyzer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, analyzer_name=None, arn=None, tags=None, type=None):
        """
        Get an existing Analyzer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] analyzer_name: Name of the Analyzer.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags.
        :param pulumi.Input[str] type: Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["analyzer_name"] = analyzer_name
        __props__["arn"] = arn
        __props__["tags"] = tags
        __props__["type"] = type
        return Analyzer(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


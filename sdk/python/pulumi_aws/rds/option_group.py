# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class OptionGroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the db option group.
    """
    engine_name: pulumi.Output[str]
    """
    Specifies the name of the engine that this option group should be associated with.
    """
    major_engine_version: pulumi.Output[str]
    """
    Specifies the major version of the engine that this option group should be associated with.
    """
    name: pulumi.Output[str]
    """
    The Name of the setting.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
    """
    options: pulumi.Output[list]
    """
    A list of Options to apply.
    """
    option_group_description: pulumi.Output[str]
    """
    The description of the option group. Defaults to "Managed by Pulumi".
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, engine_name=None, major_engine_version=None, name=None, name_prefix=None, options=None, option_group_description=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an RDS DB option group resource. Documentation of the available options for various RDS engines can be found at:
        * [MariaDB Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Options.html)
        * [Microsoft SQL Server Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.html)
        * [MySQL Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.html)
        * [Oracle Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.Options.html)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] engine_name: Specifies the name of the engine that this option group should be associated with.
        :param pulumi.Input[str] major_engine_version: Specifies the major version of the engine that this option group should be associated with.
        :param pulumi.Input[str] name: The Name of the setting.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
        :param pulumi.Input[list] options: A list of Options to apply.
        :param pulumi.Input[str] option_group_description: The description of the option group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_option_group.html.markdown.
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

            if engine_name is None:
                raise TypeError("Missing required property 'engine_name'")
            __props__['engine_name'] = engine_name
            if major_engine_version is None:
                raise TypeError("Missing required property 'major_engine_version'")
            __props__['major_engine_version'] = major_engine_version
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['options'] = options
            if option_group_description is None:
                option_group_description = 'Managed by Pulumi'
            __props__['option_group_description'] = option_group_description
            __props__['tags'] = tags
            __props__['arn'] = None
        super(OptionGroup, __self__).__init__(
            'aws:rds/optionGroup:OptionGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, engine_name=None, major_engine_version=None, name=None, name_prefix=None, options=None, option_group_description=None, tags=None):
        """
        Get an existing OptionGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the db option group.
        :param pulumi.Input[str] engine_name: Specifies the name of the engine that this option group should be associated with.
        :param pulumi.Input[str] major_engine_version: Specifies the major version of the engine that this option group should be associated with.
        :param pulumi.Input[str] name: The Name of the setting.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
        :param pulumi.Input[list] options: A list of Options to apply.
        :param pulumi.Input[str] option_group_description: The description of the option group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_option_group.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["engine_name"] = engine_name
        __props__["major_engine_version"] = major_engine_version
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["options"] = options
        __props__["option_group_description"] = option_group_description
        __props__["tags"] = tags
        return OptionGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Plan(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the backup plan.
    """
    name: pulumi.Output[str]
    """
    The display name of a backup plan.
    """
    rules: pulumi.Output[list]
    """
    A rule object that specifies a scheduled task that is used to back up a selection of resources.
    """
    tags: pulumi.Output[dict]
    """
    Metadata that you can assign to help organize the plans you create.
    """
    version: pulumi.Output[str]
    """
    Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
    """
    def __init__(__self__, resource_name, opts=None, name=None, rules=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Backup plan resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The display name of a backup plan.
        :param pulumi.Input[list] rules: A rule object that specifies a scheduled task that is used to back up a selection of resources.
        :param pulumi.Input[dict] tags: Metadata that you can assign to help organize the plans you create.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_plan.html.markdown.
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

            __props__['name'] = name
            if rules is None:
                raise TypeError("Missing required property 'rules'")
            __props__['rules'] = rules
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['version'] = None
        super(Plan, __self__).__init__(
            'aws:backup/plan:Plan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, name=None, rules=None, tags=None, version=None):
        """
        Get an existing Plan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the backup plan.
        :param pulumi.Input[str] name: The display name of a backup plan.
        :param pulumi.Input[list] rules: A rule object that specifies a scheduled task that is used to back up a selection of resources.
        :param pulumi.Input[dict] tags: Metadata that you can assign to help organize the plans you create.
        :param pulumi.Input[str] version: Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_plan.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["name"] = name
        __props__["rules"] = rules
        __props__["tags"] = tags
        __props__["version"] = version
        return Plan(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


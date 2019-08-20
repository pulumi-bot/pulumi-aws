# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GlobalTable(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the DynamoDB Global Table
    """
    name: pulumi.Output[str]
    """
    The name of the global table. Must match underlying DynamoDB Table names in all regions.
    """
    replicas: pulumi.Output[list]
    """
    Underlying DynamoDB Table. At least 1 replica must be defined. See below.
    """
    def __init__(__self__, resource_name, opts=None, name=None, replicas=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage a DynamoDB Global Table. These are layered on top of existing DynamoDB Tables.
        
        > Note: There are many restrictions before you can properly create DynamoDB Global Tables in multiple regions. See the [AWS DynamoDB Global Table Requirements](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables_reqs_bestpractices.html) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the global table. Must match underlying DynamoDB Table names in all regions.
        :param pulumi.Input[list] replicas: Underlying DynamoDB Table. At least 1 replica must be defined. See below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dynamodb_global_table.html.markdown.
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
            if replicas is None:
                raise TypeError("Missing required property 'replicas'")
            __props__['replicas'] = replicas
            __props__['arn'] = None
        super(GlobalTable, __self__).__init__(
            'aws:dynamodb/globalTable:GlobalTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, name=None, replicas=None):
        """
        Get an existing GlobalTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the DynamoDB Global Table
        :param pulumi.Input[str] name: The name of the global table. Must match underlying DynamoDB Table names in all regions.
        :param pulumi.Input[list] replicas: Underlying DynamoDB Table. At least 1 replica must be defined. See below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dynamodb_global_table.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["name"] = name
        __props__["replicas"] = replicas
        return GlobalTable(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


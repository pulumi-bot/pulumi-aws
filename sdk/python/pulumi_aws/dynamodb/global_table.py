# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
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
    def __init__(__self__, __name__, __opts__=None, name=None, replicas=None):
        """
        Provides a resource to manage a DynamoDB Global Table. These are layered on top of existing DynamoDB Tables.
        
        > Note: There are many restrictions before you can properly create DynamoDB Global Tables in multiple regions. See the [AWS DynamoDB Global Table Requirements](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables_reqs_bestpractices.html) for more information.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] name: The name of the global table. Must match underlying DynamoDB Table names in all regions.
        :param pulumi.Input[list] replicas: Underlying DynamoDB Table. At least 1 replica must be defined. See below.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        if replicas is None:
            raise TypeError('Missing required property replicas')
        __props__['replicas'] = replicas

        __props__['arn'] = None

        super(GlobalTable, __self__).__init__(
            'aws:dynamodb/globalTable:GlobalTable',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


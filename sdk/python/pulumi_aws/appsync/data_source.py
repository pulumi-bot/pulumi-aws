# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class DataSource(pulumi.CustomResource):
    """
    Provides an AppSync DataSource.
    """
    def __init__(__self__, __name__, __opts__=None, api_id=None, description=None, dynamodb_config=None, elasticsearch_config=None, http_config=None, lambda_config=None, name=None, service_role_arn=None, type=None):
        """Create a DataSource resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not api_id:
            raise TypeError('Missing required property api_id')
        __props__['apiId'] = api_id

        __props__['description'] = description

        __props__['dynamodbConfig'] = dynamodb_config

        __props__['elasticsearchConfig'] = elasticsearch_config

        __props__['httpConfig'] = http_config

        __props__['lambdaConfig'] = lambda_config

        __props__['name'] = name

        __props__['serviceRoleArn'] = service_role_arn

        if not type:
            raise TypeError('Missing required property type')
        __props__['type'] = type

        __props__['arn'] = None

        super(DataSource, __self__).__init__(
            'aws:appsync/dataSource:DataSource',
            __name__,
            __props__,
            __opts__)


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ResourceDataSync(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Name for the configuration.
    """
    s3_destination: pulumi.Output[dict]
    """
    Amazon S3 configuration details for the sync.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, s3_destination=None):
        """
        Provides a SSM resource data sync.
        
        ## s3_destination
        
        `s3_destination` supports the following:
        
        * `bucket_name` - (Required) Name of S3 bucket where the aggregated data is stored.
        * `region` - (Required) Region with the bucket targeted by the Resource Data Sync.
        * `kms_key_arn` - (Optional) ARN of an encryption key for a destination in Amazon S3.
        * `prefix` - (Optional) Prefix for the bucket.
        * `sync_format` - (Optional) A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] name: Name for the configuration.
        :param pulumi.Input[dict] s3_destination: Amazon S3 configuration details for the sync.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        if not s3_destination:
            raise TypeError('Missing required property s3_destination')
        __props__['s3_destination'] = s3_destination

        super(ResourceDataSync, __self__).__init__(
            'aws:ssm/resourceDataSync:ResourceDataSync',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProductSubscription(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of a resource that represents your subscription to the product that generates the findings that you want to import into Security Hub.
    """
    product_arn: pulumi.Output[str]
    """
    The ARN of the product that generates findings that you want to import into Security Hub - see below.
    """
    def __init__(__self__, resource_name, opts=None, product_arn=None, __name__=None, __opts__=None):
        """
        Subscribes to a Security Hub product.
        
        > **NOTE:** This AWS service is in Preview and may change before General Availability release. Backwards compatibility is not guaranteed between Terraform AWS Provider releases.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] product_arn: The ARN of the product that generates findings that you want to import into Security Hub - see below.
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

        if product_arn is None:
            raise TypeError("Missing required property 'product_arn'")
        __props__['product_arn'] = product_arn

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ProductSubscription, __self__).__init__(
            'aws:securityhub/productSubscription:ProductSubscription',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


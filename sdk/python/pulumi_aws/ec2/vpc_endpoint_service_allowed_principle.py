# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpcEndpointServiceAllowedPrinciple(pulumi.CustomResource):
    """
    Provides a resource to allow a principal to discover a VPC endpoint service.
    
    > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** Terraform provides
    both a standalone VPC Endpoint Service Allowed Principal resource
    and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
    a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
    and will overwrite the association.
    """
    def __init__(__self__, __name__, __opts__=None, principal_arn=None, vpc_endpoint_service_id=None):
        """Create a VpcEndpointServiceAllowedPrinciple resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not principal_arn:
            raise TypeError('Missing required property principal_arn')
        __props__['principal_arn'] = principal_arn

        if not vpc_endpoint_service_id:
            raise TypeError('Missing required property vpc_endpoint_service_id')
        __props__['vpc_endpoint_service_id'] = vpc_endpoint_service_id

        super(VpcEndpointServiceAllowedPrinciple, __self__).__init__(
            'aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


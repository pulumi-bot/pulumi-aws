# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ZoneAssociation(pulumi.CustomResource):
    vpc_id: pulumi.Output[str]
    """
    The VPC to associate with the private hosted zone.
    """
    vpc_region: pulumi.Output[str]
    """
    The VPC's region. Defaults to the region of the AWS provider.
    """
    zone_id: pulumi.Output[str]
    """
    The private hosted zone to associate.
    """
    def __init__(__self__, __name__, __opts__=None, vpc_id=None, vpc_region=None, zone_id=None):
        """
        Manages a Route53 Hosted Zone VPC association. VPC associations can only be made on private zones.
        
        > **NOTE:** Unless explicit association ordering is required (e.g. a separate cross-account association authorization), usage of this resource is not recommended. Use the `vpc` configuration blocks available within the [`aws_route53_zone` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone.html) instead.
        
        > **NOTE:** Terraform provides both this standalone Zone VPC Association resource and exclusive VPC associations defined in-line in the [`aws_route53_zone` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone.html) via `vpc` configuration blocks. At this time, you cannot use those in-line VPC associations in conjunction with this resource and the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use the generic Terraform resource [lifecycle configuration block](https://www.terraform.io/docs/configuration/resources.html#lifecycle) with `ignore_changes` in the `aws_route53_zone` resource to manage additional associations via this resource.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] vpc_id: The VPC to associate with the private hosted zone.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        :param pulumi.Input[str] zone_id: The private hosted zone to associate.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not vpc_id:
            raise TypeError('Missing required property vpc_id')
        __props__['vpc_id'] = vpc_id

        __props__['vpc_region'] = vpc_region

        if not zone_id:
            raise TypeError('Missing required property zone_id')
        __props__['zone_id'] = zone_id

        super(ZoneAssociation, __self__).__init__(
            'aws:route53/zoneAssociation:ZoneAssociation',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


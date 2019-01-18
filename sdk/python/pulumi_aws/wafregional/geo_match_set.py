# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GeoMatchSet(pulumi.CustomResource):
    geo_match_constraints: pulumi.Output[list]
    """
    The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
    """
    name: pulumi.Output[str]
    """
    The name or description of the Geo Match Set.
    """
    def __init__(__self__, __name__, __opts__=None, geo_match_constraints=None, name=None):
        """
        Provides a WAF Regional Geo Match Set Resource
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] geo_match_constraints: The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
        :param pulumi.Input[str] name: The name or description of the Geo Match Set.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['geo_match_constraints'] = geo_match_constraints

        __props__['name'] = name

        super(GeoMatchSet, __self__).__init__(
            'aws:wafregional/geoMatchSet:GeoMatchSet',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


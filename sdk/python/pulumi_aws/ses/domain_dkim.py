# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class DomainDkim(pulumi.CustomResource):
    """
    Provides an SES domain DKIM generation resource.
    
    Domain ownership needs to be confirmed first using [ses_domain_identity Resource](https://www.terraform.io/docs/providers/aws/r/ses_domain_identity.html)
    """
    def __init__(__self__, __name__, __opts__=None, domain=None):
        """Create a DomainDkim resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not domain:
            raise TypeError('Missing required property domain')
        __props__['domain'] = domain

        __props__['dkim_tokens'] = None

        super(DomainDkim, __self__).__init__(
            'aws:ses/domainDkim:DomainDkim',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


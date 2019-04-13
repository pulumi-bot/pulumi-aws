# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SamlProvider(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN assigned by AWS for this provider.
    """
    name: pulumi.Output[str]
    """
    The name of the provider to create.
    """
    saml_metadata_document: pulumi.Output[str]
    """
    An XML document generated by an identity provider that supports SAML 2.0.
    """
    valid_until: pulumi.Output[str]
    """
    The expiration date and time for the SAML provider in RFC1123 format, e.g. `Mon, 02 Jan 2006 15:04:05 MST`.
    """
    def __init__(__self__, resource_name, opts=None, name=None, saml_metadata_document=None, __name__=None, __opts__=None):
        """
        Provides an IAM SAML provider.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the provider to create.
        :param pulumi.Input[str] saml_metadata_document: An XML document generated by an identity provider that supports SAML 2.0.
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

        __props__['name'] = name

        if saml_metadata_document is None:
            raise TypeError("Missing required property 'saml_metadata_document'")
        __props__['saml_metadata_document'] = saml_metadata_document

        __props__['arn'] = None
        __props__['valid_until'] = None

        super(SamlProvider, __self__).__init__(
            'aws:iam/samlProvider:SamlProvider',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


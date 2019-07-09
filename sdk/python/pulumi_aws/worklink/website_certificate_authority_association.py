# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class WebsiteCertificateAuthorityAssociation(pulumi.CustomResource):
    certificate: pulumi.Output[str]
    """
    The root certificate of the Certificate Authority.
    """
    display_name: pulumi.Output[str]
    """
    The certificate name to display.
    """
    fleet_arn: pulumi.Output[str]
    """
    The ARN of the fleet.
    """
    website_ca_id: pulumi.Output[str]
    """
    A unique identifier for the Certificate Authority.
    """
    def __init__(__self__, resource_name, opts=None, certificate=None, display_name=None, fleet_arn=None, __name__=None, __opts__=None):
        """
        Create a WebsiteCertificateAuthorityAssociation resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: The root certificate of the Certificate Authority.
        :param pulumi.Input[str] display_name: The certificate name to display.
        :param pulumi.Input[str] fleet_arn: The ARN of the fleet.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/worklink_website_certificate_authority_association.html.markdown.
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

        if certificate is None:
            raise TypeError("Missing required property 'certificate'")
        __props__['certificate'] = certificate

        __props__['display_name'] = display_name

        if fleet_arn is None:
            raise TypeError("Missing required property 'fleet_arn'")
        __props__['fleet_arn'] = fleet_arn

        __props__['website_ca_id'] = None

        super(WebsiteCertificateAuthorityAssociation, __self__).__init__(
            'aws:worklink/websiteCertificateAuthorityAssociation:WebsiteCertificateAuthorityAssociation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


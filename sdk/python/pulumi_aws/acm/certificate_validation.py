# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CertificateValidation(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The ARN of the certificate that is being validated.
    """
    validation_record_fqdns: pulumi.Output[list]
    """
    List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
    """
    def __init__(__self__, resource_name, opts=None, certificate_arn=None, validation_record_fqdns=None, __name__=None, __opts__=None):
        """
        This resource represents a successful validation of an ACM certificate in concert
        with other resources.
        
        Most commonly, this resource is used together with `aws_route53_record` and
        `aws_acm_certificate` to request a DNS validated certificate,
        deploy the required validation records and wait for validation to complete.
        
        > **WARNING:** This resource implements a part of the validation workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate that is being validated.
        :param pulumi.Input[list] validation_record_fqdns: List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acm_certificate_validation.html.markdown.
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

        if certificate_arn is None:
            raise TypeError("Missing required property 'certificate_arn'")
        __props__['certificate_arn'] = certificate_arn

        __props__['validation_record_fqdns'] = validation_record_fqdns

        super(CertificateValidation, __self__).__init__(
            'aws:acm/certificateValidation:CertificateValidation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


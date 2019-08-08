# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ListenerCertificate(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The ARN of the certificate to attach to the listener.
    """
    listener_arn: pulumi.Output[str]
    """
    The ARN of the listener to which to attach the certificate.
    """
    def __init__(__self__, resource_name, opts=None, certificate_arn=None, listener_arn=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer Listener Certificate resource.
        
        This resource is for additional certificates and does not replace the default certificate on the listener.
        
        > **Note:** `alb.ListenerCertificate` is known as `lb.ListenerCertificate`. The functionality is identical.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate to attach to the listener.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the certificate.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener_certificate.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if certificate_arn is None:
            raise TypeError("Missing required property 'certificate_arn'")
        __props__['certificate_arn'] = certificate_arn
        if listener_arn is None:
            raise TypeError("Missing required property 'listener_arn'")
        __props__['listener_arn'] = listener_arn
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate")])
        opts = alias_opts if opts is None else opts.merge(alias_opts)
        super(ListenerCertificate, __self__).__init__(
            'aws:lb/listenerCertificate:ListenerCertificate',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Fleet(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the created WorkLink Fleet.
    """
    audit_stream_arn: pulumi.Output[str]
    """
    The ARN of the Amazon Kinesis data stream that receives the audit events.
    """
    company_code: pulumi.Output[str]
    """
    The identifier used by users to sign in to the Amazon WorkLink app.
    """
    created_time: pulumi.Output[str]
    """
    The time that the fleet was created.
    """
    device_ca_certificate: pulumi.Output[str]
    """
    The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
    """
    display_name: pulumi.Output[str]
    """
    The name of the fleet.
    """
    identity_provider: pulumi.Output[dict]
    """
    Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
    """
    last_updated_time: pulumi.Output[str]
    """
    The time that the fleet was last updated.
    """
    name: pulumi.Output[str]
    """
    A region-unique name for the AMI.
    """
    network: pulumi.Output[dict]
    """
    Provide this to allow manage the company network configuration for the fleet. Fields documented below.
    """
    optimize_for_end_user_location: pulumi.Output[bool]
    """
    The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
    """
    def __init__(__self__, resource_name, opts=None, audit_stream_arn=None, device_ca_certificate=None, display_name=None, identity_provider=None, name=None, network=None, optimize_for_end_user_location=None, __name__=None, __opts__=None):
        """
        Create a Fleet resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] audit_stream_arn: The ARN of the Amazon Kinesis data stream that receives the audit events.
        :param pulumi.Input[str] device_ca_certificate: The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        :param pulumi.Input[str] display_name: The name of the fleet.
        :param pulumi.Input[dict] identity_provider: Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        :param pulumi.Input[str] name: A region-unique name for the AMI.
        :param pulumi.Input[dict] network: Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        :param pulumi.Input[bool] optimize_for_end_user_location: The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/worklink_fleet.html.markdown.
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

        __props__['audit_stream_arn'] = audit_stream_arn

        __props__['device_ca_certificate'] = device_ca_certificate

        __props__['display_name'] = display_name

        __props__['identity_provider'] = identity_provider

        __props__['name'] = name

        __props__['network'] = network

        __props__['optimize_for_end_user_location'] = optimize_for_end_user_location

        __props__['arn'] = None
        __props__['company_code'] = None
        __props__['created_time'] = None
        __props__['last_updated_time'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Fleet, __self__).__init__(
            'aws:worklink/fleet:Fleet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


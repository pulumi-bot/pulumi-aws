# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class PublicDnsNamespace(pulumi.CustomResource):
    """
    Provides a Service Discovery Public DNS Namespace resource.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, name=None):
        """Create a PublicDnsNamespace resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['name'] = name

        __props__['arn'] = None
        __props__['hosted_zone'] = None

        super(PublicDnsNamespace, __self__).__init__(
            'aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace',
            __name__,
            __props__,
            __opts__)


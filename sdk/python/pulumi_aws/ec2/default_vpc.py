# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class DefaultVpc(pulumi.CustomResource):
    """
    Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
    in the current region.
    
    For AWS accounts created after 2013-12-04, each region comes with a Default VPC.
    **This is an advanced resource**, and has special caveats to be aware of when
    using it. Please read this document in its entirety before using this resource.
    
    The `aws_default_vpc` behaves differently from normal resources, in that
    Terraform does not _create_ this resource, but instead "adopts" it
    into management. 
    """
    def __init__(__self__, __name__, __opts__=None, enable_classiclink=None, enable_classiclink_dns_support=None, enable_dns_hostnames=None, enable_dns_support=None, tags=None):
        """Create a DefaultVpc resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['enableClassiclink'] = enable_classiclink

        __props__['enableClassiclinkDnsSupport'] = enable_classiclink_dns_support

        __props__['enableDnsHostnames'] = enable_dns_hostnames

        __props__['enableDnsSupport'] = enable_dns_support

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['assign_generated_ipv6_cidr_block'] = None
        __props__['cidr_block'] = None
        __props__['default_network_acl_id'] = None
        __props__['default_route_table_id'] = None
        __props__['default_security_group_id'] = None
        __props__['dhcp_options_id'] = None
        __props__['instance_tenancy'] = None
        __props__['ipv6_association_id'] = None
        __props__['ipv6_cidr_block'] = None
        __props__['main_route_table_id'] = None

        super(DefaultVpc, __self__).__init__(
            'aws:ec2/defaultVpc:DefaultVpc',
            __name__,
            __props__,
            __opts__)


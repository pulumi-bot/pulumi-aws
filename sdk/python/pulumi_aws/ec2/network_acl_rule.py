# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NetworkAclRule(pulumi.CustomResource):
    """
    Creates an entry (a rule) in a network ACL with the specified rule number.
    
    > **NOTE on Network ACLs and Network ACL Rules:** Terraform currently
    provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
    defined in-line. At this time you cannot use a Network ACL with in-line rules
    in conjunction with any Network ACL Rule resources. Doing so will cause
    a conflict of rule settings and will overwrite rules.
    """
    def __init__(__self__, __name__, __opts__=None, cidr_block=None, egress=None, from_port=None, icmp_code=None, icmp_type=None, ipv6_cidr_block=None, network_acl_id=None, protocol=None, rule_action=None, rule_number=None, to_port=None):
        """Create a NetworkAclRule resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cidr_block'] = cidr_block

        __props__['egress'] = egress

        __props__['from_port'] = from_port

        __props__['icmp_code'] = icmp_code

        __props__['icmp_type'] = icmp_type

        __props__['ipv6_cidr_block'] = ipv6_cidr_block

        if not network_acl_id:
            raise TypeError('Missing required property network_acl_id')
        __props__['network_acl_id'] = network_acl_id

        if not protocol:
            raise TypeError('Missing required property protocol')
        __props__['protocol'] = protocol

        if not rule_action:
            raise TypeError('Missing required property rule_action')
        __props__['rule_action'] = rule_action

        if not rule_number:
            raise TypeError('Missing required property rule_number')
        __props__['rule_number'] = rule_number

        __props__['to_port'] = to_port

        super(NetworkAclRule, __self__).__init__(
            'aws:ec2/networkAclRule:NetworkAclRule',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NetworkAclRule(pulumi.CustomResource):
    cidr_block: pulumi.Output[str]
    """
    The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
    """
    egress: pulumi.Output[bool]
    """
    Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
    """
    from_port: pulumi.Output[float]
    """
    The from port to match.
    """
    icmp_code: pulumi.Output[str]
    """
    ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
    """
    icmp_type: pulumi.Output[str]
    """
    ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
    """
    ipv6_cidr_block: pulumi.Output[str]
    """
    The IPv6 CIDR block to allow or deny.
    """
    network_acl_id: pulumi.Output[str]
    """
    The ID of the network ACL.
    """
    protocol: pulumi.Output[str]
    """
    The protocol. A value of -1 means all protocols.
    """
    rule_action: pulumi.Output[str]
    """
    Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
    """
    rule_number: pulumi.Output[float]
    """
    The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
    """
    to_port: pulumi.Output[float]
    """
    The to port to match.
    """
    def __init__(__self__, resource_name, opts=None, cidr_block=None, egress=None, from_port=None, icmp_code=None, icmp_type=None, ipv6_cidr_block=None, network_acl_id=None, protocol=None, rule_action=None, rule_number=None, to_port=None, __name__=None, __opts__=None):
        """
        Create a NetworkAclRule resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        :param pulumi.Input[bool] egress: Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        :param pulumi.Input[float] from_port: The from port to match.
        :param pulumi.Input[str] icmp_code: ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] icmp_type: ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block to allow or deny.
        :param pulumi.Input[str] network_acl_id: The ID of the network ACL.
        :param pulumi.Input[str] protocol: The protocol. A value of -1 means all protocols.
        :param pulumi.Input[str] rule_action: Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        :param pulumi.Input[float] rule_number: The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        :param pulumi.Input[float] to_port: The to port to match.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/network_acl_rule.html.markdown.
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

        __props__['cidr_block'] = cidr_block

        __props__['egress'] = egress

        __props__['from_port'] = from_port

        __props__['icmp_code'] = icmp_code

        __props__['icmp_type'] = icmp_type

        __props__['ipv6_cidr_block'] = ipv6_cidr_block

        if network_acl_id is None:
            raise TypeError("Missing required property 'network_acl_id'")
        __props__['network_acl_id'] = network_acl_id

        if protocol is None:
            raise TypeError("Missing required property 'protocol'")
        __props__['protocol'] = protocol

        if rule_action is None:
            raise TypeError("Missing required property 'rule_action'")
        __props__['rule_action'] = rule_action

        if rule_number is None:
            raise TypeError("Missing required property 'rule_number'")
        __props__['rule_number'] = rule_number

        __props__['to_port'] = to_port

        super(NetworkAclRule, __self__).__init__(
            'aws:ec2/networkAclRule:NetworkAclRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


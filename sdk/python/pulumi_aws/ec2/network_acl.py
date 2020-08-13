# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class NetworkAcl(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the network ACL
    """
    egress: pulumi.Output[list]
    """
    Specifies an egress rule. Parameters defined below.

      * `action` (`str`) - The action to take.
      * `cidr_block` (`str`) - The CIDR block to match. This must be a
        valid network mask.
      * `from_port` (`float`) - The from port to match.
      * `icmp_code` (`float`) - The ICMP type code to be used. Default 0.
      * `icmp_type` (`float`) - The ICMP type to be used. Default 0.
      * `ipv6_cidr_block` (`str`) - The IPv6 CIDR block.
      * `protocol` (`str`) - The protocol to match. If using the -1 'all'
        protocol, you must specify a from and to port of 0.
      * `ruleNo` (`float`) - The rule number. Used for ordering.
      * `to_port` (`float`) - The to port to match.
    """
    ingress: pulumi.Output[list]
    """
    Specifies an ingress rule. Parameters defined below.

      * `action` (`str`) - The action to take.
      * `cidr_block` (`str`) - The CIDR block to match. This must be a
        valid network mask.
      * `from_port` (`float`) - The from port to match.
      * `icmp_code` (`float`) - The ICMP type code to be used. Default 0.
      * `icmp_type` (`float`) - The ICMP type to be used. Default 0.
      * `ipv6_cidr_block` (`str`) - The IPv6 CIDR block.
      * `protocol` (`str`) - The protocol to match. If using the -1 'all'
        protocol, you must specify a from and to port of 0.
      * `ruleNo` (`float`) - The rule number. Used for ordering.
      * `to_port` (`float`) - The to port to match.
    """
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the network ACL.
    """
    subnet_ids: pulumi.Output[list]
    """
    A list of Subnet IDs to apply the ACL to
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the associated VPC.
    """
    def __init__(__self__, resource_name, opts=None, egress=None, ingress=None, subnet_ids=None, tags=None, vpc_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an network ACL resource. You might set up network ACLs with rules similar
        to your security groups in order to add an additional layer of security to your VPC.

        > **NOTE on Network ACLs and Network ACL Rules:** This provider currently
        provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
        defined in-line. At this time you cannot use a Network ACL with in-line rules
        in conjunction with any Network ACL Rule resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.NetworkAcl("main",
            vpc_id=aws_vpc["main"]["id"],
            egress=[{
                "protocol": "tcp",
                "ruleNo": 200,
                "action": "allow",
                "cidr_block": "10.3.0.0/18",
                "from_port": 443,
                "to_port": 443,
            }],
            ingress=[{
                "protocol": "tcp",
                "ruleNo": 100,
                "action": "allow",
                "cidr_block": "10.3.0.0/18",
                "from_port": 80,
                "to_port": 80,
            }],
            tags={
                "Name": "main",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] egress: Specifies an egress rule. Parameters defined below.
        :param pulumi.Input[list] ingress: Specifies an ingress rule. Parameters defined below.
        :param pulumi.Input[list] subnet_ids: A list of Subnet IDs to apply the ACL to
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the associated VPC.

        The **egress** object supports the following:

          * `action` (`pulumi.Input[str]`) - The action to take.
          * `cidr_block` (`pulumi.Input[str]`) - The CIDR block to match. This must be a
            valid network mask.
          * `from_port` (`pulumi.Input[float]`) - The from port to match.
          * `icmp_code` (`pulumi.Input[float]`) - The ICMP type code to be used. Default 0.
          * `icmp_type` (`pulumi.Input[float]`) - The ICMP type to be used. Default 0.
          * `ipv6_cidr_block` (`pulumi.Input[str]`) - The IPv6 CIDR block.
          * `protocol` (`pulumi.Input[str]`) - The protocol to match. If using the -1 'all'
            protocol, you must specify a from and to port of 0.
          * `ruleNo` (`pulumi.Input[float]`) - The rule number. Used for ordering.
          * `to_port` (`pulumi.Input[float]`) - The to port to match.

        The **ingress** object supports the following:

          * `action` (`pulumi.Input[str]`) - The action to take.
          * `cidr_block` (`pulumi.Input[str]`) - The CIDR block to match. This must be a
            valid network mask.
          * `from_port` (`pulumi.Input[float]`) - The from port to match.
          * `icmp_code` (`pulumi.Input[float]`) - The ICMP type code to be used. Default 0.
          * `icmp_type` (`pulumi.Input[float]`) - The ICMP type to be used. Default 0.
          * `ipv6_cidr_block` (`pulumi.Input[str]`) - The IPv6 CIDR block.
          * `protocol` (`pulumi.Input[str]`) - The protocol to match. If using the -1 'all'
            protocol, you must specify a from and to port of 0.
          * `ruleNo` (`pulumi.Input[float]`) - The rule number. Used for ordering.
          * `to_port` (`pulumi.Input[float]`) - The to port to match.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['egress'] = egress
            __props__['ingress'] = ingress
            __props__['subnetIds'] = subnet_ids
            __props__['tags'] = tags
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpcId'] = vpc_id
            __props__['arn'] = None
            __props__['owner_id'] = None
        super(NetworkAcl, __self__).__init__(
            'aws:ec2/networkAcl:NetworkAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, egress=None, ingress=None, owner_id=None, subnet_ids=None, tags=None, vpc_id=None):
        """
        Get an existing NetworkAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the network ACL
        :param pulumi.Input[list] egress: Specifies an egress rule. Parameters defined below.
        :param pulumi.Input[list] ingress: Specifies an ingress rule. Parameters defined below.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the network ACL.
        :param pulumi.Input[list] subnet_ids: A list of Subnet IDs to apply the ACL to
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the associated VPC.

        The **egress** object supports the following:

          * `action` (`pulumi.Input[str]`) - The action to take.
          * `cidr_block` (`pulumi.Input[str]`) - The CIDR block to match. This must be a
            valid network mask.
          * `from_port` (`pulumi.Input[float]`) - The from port to match.
          * `icmp_code` (`pulumi.Input[float]`) - The ICMP type code to be used. Default 0.
          * `icmp_type` (`pulumi.Input[float]`) - The ICMP type to be used. Default 0.
          * `ipv6_cidr_block` (`pulumi.Input[str]`) - The IPv6 CIDR block.
          * `protocol` (`pulumi.Input[str]`) - The protocol to match. If using the -1 'all'
            protocol, you must specify a from and to port of 0.
          * `ruleNo` (`pulumi.Input[float]`) - The rule number. Used for ordering.
          * `to_port` (`pulumi.Input[float]`) - The to port to match.

        The **ingress** object supports the following:

          * `action` (`pulumi.Input[str]`) - The action to take.
          * `cidr_block` (`pulumi.Input[str]`) - The CIDR block to match. This must be a
            valid network mask.
          * `from_port` (`pulumi.Input[float]`) - The from port to match.
          * `icmp_code` (`pulumi.Input[float]`) - The ICMP type code to be used. Default 0.
          * `icmp_type` (`pulumi.Input[float]`) - The ICMP type to be used. Default 0.
          * `ipv6_cidr_block` (`pulumi.Input[str]`) - The IPv6 CIDR block.
          * `protocol` (`pulumi.Input[str]`) - The protocol to match. If using the -1 'all'
            protocol, you must specify a from and to port of 0.
          * `ruleNo` (`pulumi.Input[float]`) - The rule number. Used for ordering.
          * `to_port` (`pulumi.Input[float]`) - The to port to match.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["egress"] = egress
        __props__["ingress"] = ingress
        __props__["owner_id"] = owner_id
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return NetworkAcl(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

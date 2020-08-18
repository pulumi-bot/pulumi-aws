# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['NetworkAcl']


class NetworkAcl(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN of the network ACL
    """

    egress: pulumi.Output[List['outputs.NetworkAclEgress']] = pulumi.property("egress")
    """
    Specifies an egress rule. Parameters defined below.
    """

    ingress: pulumi.Output[List['outputs.NetworkAclIngress']] = pulumi.property("ingress")
    """
    Specifies an ingress rule. Parameters defined below.
    """

    owner_id: pulumi.Output[str] = pulumi.property("ownerId")
    """
    The ID of the AWS account that owns the network ACL.
    """

    subnet_ids: pulumi.Output[List[str]] = pulumi.property("subnetIds")
    """
    A list of Subnet IDs to apply the ACL to
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    vpc_id: pulumi.Output[str] = pulumi.property("vpcId")
    """
    The ID of the associated VPC.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclEgressArgs']]]]] = None,
                 ingress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclIngressArgs']]]]] = None,
                 subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclEgressArgs']]]] egress: Specifies an egress rule. Parameters defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclIngressArgs']]]] ingress: Specifies an ingress rule. Parameters defined below.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: A list of Subnet IDs to apply the ACL to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the associated VPC.
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
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
            __props__['owner_id'] = None
        super(NetworkAcl, __self__).__init__(
            'aws:ec2/networkAcl:NetworkAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            egress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclEgressArgs']]]]] = None,
            ingress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclIngressArgs']]]]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'NetworkAcl':
        """
        Get an existing NetworkAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the network ACL
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclEgressArgs']]]] egress: Specifies an egress rule. Parameters defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclIngressArgs']]]] ingress: Specifies an ingress rule. Parameters defined below.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the network ACL.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: A list of Subnet IDs to apply the ACL to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the associated VPC.
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


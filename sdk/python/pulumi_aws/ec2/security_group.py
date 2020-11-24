# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SecurityGroup']


class SecurityGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupEgressArgs']]]]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a security group resource.

        > **NOTE on Security Groups and Security Group Rules:** This provider currently
        provides both a standalone Security Group Rule resource (a single `ingress` or
        `egress` rule), and a Security Group resource with `ingress` and `egress` rules
        defined in-line. At this time you cannot use a Security Group with in-line rules
        in conjunction with any Security Group Rule resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.

        > **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).

        > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.

        ## Example Usage

        Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        allow_tls = aws.ec2.SecurityGroup("allowTls",
            description="Allow TLS inbound traffic",
            vpc_id=aws_vpc["main"]["id"],
            ingress=[{
                "description": "TLS from VPC",
                "from_port": 443,
                "to_port": 443,
                "protocol": "tcp",
                "cidr_blocks": [aws_vpc["main"]["cidr_block"]],
            }],
            egress=[{
                "from_port": 0,
                "to_port": 0,
                "protocol": "-1",
                "cidr_blocks": ["0.0.0.0/0"],
            }],
            tags={
                "Name": "allow_tls",
            })
        ```
        ## Usage with prefix list IDs

        Prefix list IDs are managed by AWS internally. Prefix list IDs
        are associated with a prefix list name, or service name, that is linked to a specific region.
        Prefix list IDs are exported on VPC Endpoints, so you can use this format:

        ```python
        import pulumi
        import pulumi_aws as aws

        my_endpoint = aws.ec2.VpcEndpoint("myEndpoint")
        # ... other configuration ...
        # ... other configuration ...
        example = aws.ec2.SecurityGroup("example", egress=[{
            "from_port": 0,
            "to_port": 0,
            "protocol": "-1",
            "prefix_list_ids": [my_endpoint.prefix_list_id],
        }])
        ```

        ## Import

        Security Groups can be imported using the `security group id`, e.g.

        ```sh
         $ pulumi import aws:ec2/securityGroup:SecurityGroup elb_sg sg-903004f8
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of this egress rule.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupEgressArgs']]]] egress: Can be specified multiple times for each
               egress rule. Each egress block supports fields documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]] ingress: Can be specified multiple times for each
               ingress rule. Each ingress block supports fields documented below.
        :param pulumi.Input[str] name: The name of the security group. If omitted, this provider will
               assign a random, unique name
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[bool] revoke_rules_on_delete: Instruct this provider to revoke all of the
               Security Groups attached ingress and egress rules before deleting the rule
               itself. This is normally not needed, however certain AWS services such as
               Elastic Map Reduce may automatically add required rules to security groups used
               with the service, and those rules may contain a cyclic dependency that prevent
               the security groups from being destroyed without removing the dependency first.
               Default `false`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            __props__['egress'] = egress
            __props__['ingress'] = ingress
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['revoke_rules_on_delete'] = revoke_rules_on_delete
            __props__['tags'] = tags
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
            __props__['owner_id'] = None
        super(SecurityGroup, __self__).__init__(
            'aws:ec2/securityGroup:SecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            egress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupEgressArgs']]]]] = None,
            ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            revoke_rules_on_delete: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'SecurityGroup':
        """
        Get an existing SecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the security group
        :param pulumi.Input[str] description: Description of this egress rule.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupEgressArgs']]]] egress: Can be specified multiple times for each
               egress rule. Each egress block supports fields documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]] ingress: Can be specified multiple times for each
               ingress rule. Each ingress block supports fields documented below.
        :param pulumi.Input[str] name: The name of the security group. If omitted, this provider will
               assign a random, unique name
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[str] owner_id: The owner ID.
        :param pulumi.Input[bool] revoke_rules_on_delete: Instruct this provider to revoke all of the
               Security Groups attached ingress and egress rules before deleting the rule
               itself. This is normally not needed, however certain AWS services such as
               Elastic Map Reduce may automatically add required rules to security groups used
               with the service, and those rules may contain a cyclic dependency that prevent
               the security groups from being destroyed without removing the dependency first.
               Default `false`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["egress"] = egress
        __props__["ingress"] = ingress
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["owner_id"] = owner_id
        __props__["revoke_rules_on_delete"] = revoke_rules_on_delete
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return SecurityGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the security group
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of this egress rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def egress(self) -> pulumi.Output[Sequence['outputs.SecurityGroupEgress']]:
        """
        Can be specified multiple times for each
        egress rule. Each egress block supports fields documented below.
        """
        return pulumi.get(self, "egress")

    @property
    @pulumi.getter
    def ingress(self) -> pulumi.Output[Sequence['outputs.SecurityGroupIngress']]:
        """
        Can be specified multiple times for each
        ingress rule. Each ingress block supports fields documented below.
        """
        return pulumi.get(self, "ingress")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the security group. If omitted, this provider will
        assign a random, unique name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The owner ID.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="revokeRulesOnDelete")
    def revoke_rules_on_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Instruct this provider to revoke all of the
        Security Groups attached ingress and egress rules before deleting the rule
        itself. This is normally not needed, however certain AWS services such as
        Elastic Map Reduce may automatically add required rules to security groups used
        with the service, and those rules may contain a cyclic dependency that prevent
        the security groups from being destroyed without removing the dependency first.
        Default `false`
        """
        return pulumi.get(self, "revoke_rules_on_delete")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


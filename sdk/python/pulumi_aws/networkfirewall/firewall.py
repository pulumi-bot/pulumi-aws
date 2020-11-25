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

__all__ = ['Firewall']


class Firewall(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 firewall_policy_arn: Optional[pulumi.Input[str]] = None,
                 firewall_policy_change_protection: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_change_protection: Optional[pulumi.Input[bool]] = None,
                 subnet_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallSubnetMappingArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS Network Firewall Firewall Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.networkfirewall.Firewall("example",
            firewall_policy_arn=aws_networkfirewall_firewall_policy["example"]["arn"],
            vpc_id=aws_vpc["example"]["id"],
            subnet_mappings=[aws.networkfirewall.FirewallSubnetMappingArgs(
                subnet_id=aws_subnet["example"]["id"],
            )],
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            })
        ```

        ## Import

        Network Firewall Firewalls can be imported using their `ARN`.

        ```sh
         $ pulumi import aws:networkfirewall/firewall:Firewall example arn:aws:network-firewall:us-west-1:123456789012:firewall/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_protection: A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
        :param pulumi.Input[str] description: A friendly description of the firewall.
        :param pulumi.Input[str] firewall_policy_arn: The Amazon Resource Name (ARN) of the VPC Firewall policy.
        :param pulumi.Input[bool] firewall_policy_change_protection: A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
        :param pulumi.Input[str] name: A friendly name of the firewall.
        :param pulumi.Input[bool] subnet_change_protection: A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallSubnetMappingArgs']]]] subnet_mappings: Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The key:value pairs to associate with the resource.
        :param pulumi.Input[str] vpc_id: The unique identifier of the VPC where AWS Network Firewall should create the firewall.
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

            __props__['delete_protection'] = delete_protection
            __props__['description'] = description
            if firewall_policy_arn is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_policy_arn'")
            __props__['firewall_policy_arn'] = firewall_policy_arn
            __props__['firewall_policy_change_protection'] = firewall_policy_change_protection
            __props__['name'] = name
            __props__['subnet_change_protection'] = subnet_change_protection
            if subnet_mappings is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_mappings'")
            __props__['subnet_mappings'] = subnet_mappings
            __props__['tags'] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
            __props__['update_token'] = None
        super(Firewall, __self__).__init__(
            'aws:networkfirewall/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            delete_protection: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            firewall_policy_arn: Optional[pulumi.Input[str]] = None,
            firewall_policy_change_protection: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            subnet_change_protection: Optional[pulumi.Input[bool]] = None,
            subnet_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallSubnetMappingArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            update_token: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Firewall':
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the firewall.
        :param pulumi.Input[bool] delete_protection: A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
        :param pulumi.Input[str] description: A friendly description of the firewall.
        :param pulumi.Input[str] firewall_policy_arn: The Amazon Resource Name (ARN) of the VPC Firewall policy.
        :param pulumi.Input[bool] firewall_policy_change_protection: A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
        :param pulumi.Input[str] name: A friendly name of the firewall.
        :param pulumi.Input[bool] subnet_change_protection: A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallSubnetMappingArgs']]]] subnet_mappings: Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The key:value pairs to associate with the resource.
        :param pulumi.Input[str] update_token: A string token used when updating a firewall.
        :param pulumi.Input[str] vpc_id: The unique identifier of the VPC where AWS Network Firewall should create the firewall.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["delete_protection"] = delete_protection
        __props__["description"] = description
        __props__["firewall_policy_arn"] = firewall_policy_arn
        __props__["firewall_policy_change_protection"] = firewall_policy_change_protection
        __props__["name"] = name
        __props__["subnet_change_protection"] = subnet_change_protection
        __props__["subnet_mappings"] = subnet_mappings
        __props__["tags"] = tags
        __props__["update_token"] = update_token
        __props__["vpc_id"] = vpc_id
        return Firewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) that identifies the firewall.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
        """
        return pulumi.get(self, "delete_protection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly description of the firewall.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="firewallPolicyArn")
    def firewall_policy_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the VPC Firewall policy.
        """
        return pulumi.get(self, "firewall_policy_arn")

    @property
    @pulumi.getter(name="firewallPolicyChangeProtection")
    def firewall_policy_change_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
        """
        return pulumi.get(self, "firewall_policy_change_protection")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the firewall.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetChangeProtection")
    def subnet_change_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
        """
        return pulumi.get(self, "subnet_change_protection")

    @property
    @pulumi.getter(name="subnetMappings")
    def subnet_mappings(self) -> pulumi.Output[Sequence['outputs.FirewallSubnetMapping']]:
        """
        Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
        """
        return pulumi.get(self, "subnet_mappings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateToken")
    def update_token(self) -> pulumi.Output[str]:
        """
        A string token used when updating a firewall.
        """
        return pulumi.get(self, "update_token")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the VPC where AWS Network Firewall should create the firewall.
        """
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


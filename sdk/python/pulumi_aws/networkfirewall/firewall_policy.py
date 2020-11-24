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

__all__ = ['FirewallPolicy']


class FirewallPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[pulumi.InputType['FirewallPolicyFirewallPolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS Network Firewall Firewall Policy Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.networkfirewall.FirewallPolicy("example",
            firewall_policy={
                "statelessDefaultActions": ["aws:pass"],
                "statelessFragmentDefaultActions": ["aws:drop"],
                "statelessRuleGroupReferences": [{
                    "priority": 1,
                    "resource_arn": aws_networkfirewall_rule_group["example"]["arn"],
                }],
            },
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            })
        ```
        ## Policy with a Custom Action for Stateless Inspection

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.networkfirewall.FirewallPolicy("test", firewall_policy={
            "statelessCustomActions": [{
                "actionDefinition": {
                    "publishMetricAction": {
                        "dimension": [{
                            "value": "1",
                        }],
                    },
                },
                "actionName": "ExampleCustomAction",
            }],
            "statelessDefaultActions": [
                "aws:pass",
                "ExampleCustomAction",
            ],
            "statelessFragmentDefaultActions": ["aws:drop"],
        })
        ```

        ## Import

        Network Firewall Policies can be imported using their `ARN`.

        ```sh
         $ pulumi import aws:networkfirewall/firewallPolicy:FirewallPolicy example arn:aws:network-firewall:us-west-1:123456789012:firewall-policy/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A friendly description of the firewall policy.
        :param pulumi.Input[pulumi.InputType['FirewallPolicyFirewallPolicyArgs']] firewall_policy: A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
        :param pulumi.Input[str] name: A friendly name of the firewall policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
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

            __props__['description'] = description
            if firewall_policy is None:
                raise TypeError("Missing required property 'firewall_policy'")
            __props__['firewall_policy'] = firewall_policy
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['update_token'] = None
        super(FirewallPolicy, __self__).__init__(
            'aws:networkfirewall/firewallPolicy:FirewallPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            firewall_policy: Optional[pulumi.Input[pulumi.InputType['FirewallPolicyFirewallPolicyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            update_token: Optional[pulumi.Input[str]] = None) -> 'FirewallPolicy':
        """
        Get an existing FirewallPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the firewall policy.
        :param pulumi.Input[str] description: A friendly description of the firewall policy.
        :param pulumi.Input[pulumi.InputType['FirewallPolicyFirewallPolicyArgs']] firewall_policy: A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
        :param pulumi.Input[str] name: A friendly name of the firewall policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
        :param pulumi.Input[str] update_token: A string token used when updating a firewall policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["firewall_policy"] = firewall_policy
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["update_token"] = update_token
        return FirewallPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) that identifies the firewall policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly description of the firewall policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Output['outputs.FirewallPolicyFirewallPolicy']:
        """
        A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
        """
        return pulumi.get(self, "firewall_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the firewall policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An array of key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateToken")
    def update_token(self) -> pulumi.Output[str]:
        """
        A string token used when updating a firewall policy.
        """
        return pulumi.get(self, "update_token")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


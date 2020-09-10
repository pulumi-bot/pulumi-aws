# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['OrganizationManagedRule']


class OrganizationManagedRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 input_parameters: Optional[pulumi.Input[str]] = None,
                 maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id_scope: Optional[pulumi.Input[str]] = None,
                 resource_types_scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 rule_identifier: Optional[pulumi.Input[str]] = None,
                 tag_key_scope: Optional[pulumi.Input[str]] = None,
                 tag_value_scope: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a OrganizationManagedRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
            __props__['excluded_accounts'] = excluded_accounts
            __props__['input_parameters'] = input_parameters
            __props__['maximum_execution_frequency'] = maximum_execution_frequency
            __props__['name'] = name
            __props__['resource_id_scope'] = resource_id_scope
            __props__['resource_types_scopes'] = resource_types_scopes
            if rule_identifier is None:
                raise TypeError("Missing required property 'rule_identifier'")
            __props__['rule_identifier'] = rule_identifier
            __props__['tag_key_scope'] = tag_key_scope
            __props__['tag_value_scope'] = tag_value_scope
            __props__['arn'] = None
        super(OrganizationManagedRule, __self__).__init__(
            'aws:cfg/organizationManagedRule:OrganizationManagedRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            excluded_accounts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            input_parameters: Optional[pulumi.Input[str]] = None,
            maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_id_scope: Optional[pulumi.Input[str]] = None,
            resource_types_scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            rule_identifier: Optional[pulumi.Input[str]] = None,
            tag_key_scope: Optional[pulumi.Input[str]] = None,
            tag_value_scope: Optional[pulumi.Input[str]] = None) -> 'OrganizationManagedRule':
        """
        Get an existing OrganizationManagedRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["excluded_accounts"] = excluded_accounts
        __props__["input_parameters"] = input_parameters
        __props__["maximum_execution_frequency"] = maximum_execution_frequency
        __props__["name"] = name
        __props__["resource_id_scope"] = resource_id_scope
        __props__["resource_types_scopes"] = resource_types_scopes
        __props__["rule_identifier"] = rule_identifier
        __props__["tag_key_scope"] = tag_key_scope
        __props__["tag_value_scope"] = tag_value_scope
        return OrganizationManagedRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="excludedAccounts")
    def excluded_accounts(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "excluded_accounts")

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "input_parameters")

    @property
    @pulumi.getter(name="maximumExecutionFrequency")
    def maximum_execution_frequency(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "maximum_execution_frequency")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceIdScope")
    def resource_id_scope(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "resource_id_scope")

    @property
    @pulumi.getter(name="resourceTypesScopes")
    def resource_types_scopes(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "resource_types_scopes")

    @property
    @pulumi.getter(name="ruleIdentifier")
    def rule_identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "rule_identifier")

    @property
    @pulumi.getter(name="tagKeyScope")
    def tag_key_scope(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tag_key_scope")

    @property
    @pulumi.getter(name="tagValueScope")
    def tag_value_scope(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tag_value_scope")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


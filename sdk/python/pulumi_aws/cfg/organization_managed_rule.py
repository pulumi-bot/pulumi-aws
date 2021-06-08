# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrganizationManagedRuleArgs', 'OrganizationManagedRule']

@pulumi.input_type
class OrganizationManagedRuleArgs:
    def __init__(__self__, *,
                 rule_identifier: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 input_parameters: Optional[pulumi.Input[str]] = None,
                 maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id_scope: Optional[pulumi.Input[str]] = None,
                 resource_types_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag_key_scope: Optional[pulumi.Input[str]] = None,
                 tag_value_scope: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationManagedRule resource.
        :param pulumi.Input[str] rule_identifier: Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        """
        pulumi.set(__self__, "rule_identifier", rule_identifier)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if excluded_accounts is not None:
            pulumi.set(__self__, "excluded_accounts", excluded_accounts)
        if input_parameters is not None:
            pulumi.set(__self__, "input_parameters", input_parameters)
        if maximum_execution_frequency is not None:
            pulumi.set(__self__, "maximum_execution_frequency", maximum_execution_frequency)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_id_scope is not None:
            pulumi.set(__self__, "resource_id_scope", resource_id_scope)
        if resource_types_scopes is not None:
            pulumi.set(__self__, "resource_types_scopes", resource_types_scopes)
        if tag_key_scope is not None:
            pulumi.set(__self__, "tag_key_scope", tag_key_scope)
        if tag_value_scope is not None:
            pulumi.set(__self__, "tag_value_scope", tag_value_scope)

    @property
    @pulumi.getter(name="ruleIdentifier")
    def rule_identifier(self) -> pulumi.Input[str]:
        """
        Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        """
        return pulumi.get(self, "rule_identifier")

    @rule_identifier.setter
    def rule_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_identifier", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="excludedAccounts")
    def excluded_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of AWS account identifiers to exclude from the rule
        """
        return pulumi.get(self, "excluded_accounts")

    @excluded_accounts.setter
    def excluded_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_accounts", value)

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> Optional[pulumi.Input[str]]:
        """
        A string in JSON format that is passed to the AWS Config Rule Lambda Function
        """
        return pulumi.get(self, "input_parameters")

    @input_parameters.setter
    def input_parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_parameters", value)

    @property
    @pulumi.getter(name="maximumExecutionFrequency")
    def maximum_execution_frequency(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        """
        return pulumi.get(self, "maximum_execution_frequency")

    @maximum_execution_frequency.setter
    def maximum_execution_frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maximum_execution_frequency", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceIdScope")
    def resource_id_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the AWS resource to evaluate
        """
        return pulumi.get(self, "resource_id_scope")

    @resource_id_scope.setter
    def resource_id_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id_scope", value)

    @property
    @pulumi.getter(name="resourceTypesScopes")
    def resource_types_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of types of AWS resources to evaluate
        """
        return pulumi.get(self, "resource_types_scopes")

    @resource_types_scopes.setter
    def resource_types_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_types_scopes", value)

    @property
    @pulumi.getter(name="tagKeyScope")
    def tag_key_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Tag key of AWS resources to evaluate
        """
        return pulumi.get(self, "tag_key_scope")

    @tag_key_scope.setter
    def tag_key_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_key_scope", value)

    @property
    @pulumi.getter(name="tagValueScope")
    def tag_value_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Tag value of AWS resources to evaluate
        """
        return pulumi.get(self, "tag_value_scope")

    @tag_value_scope.setter
    def tag_value_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_value_scope", value)


@pulumi.input_type
class _OrganizationManagedRuleState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 input_parameters: Optional[pulumi.Input[str]] = None,
                 maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id_scope: Optional[pulumi.Input[str]] = None,
                 resource_types_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rule_identifier: Optional[pulumi.Input[str]] = None,
                 tag_key_scope: Optional[pulumi.Input[str]] = None,
                 tag_value_scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationManagedRule resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the rule
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] rule_identifier: Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if excluded_accounts is not None:
            pulumi.set(__self__, "excluded_accounts", excluded_accounts)
        if input_parameters is not None:
            pulumi.set(__self__, "input_parameters", input_parameters)
        if maximum_execution_frequency is not None:
            pulumi.set(__self__, "maximum_execution_frequency", maximum_execution_frequency)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_id_scope is not None:
            pulumi.set(__self__, "resource_id_scope", resource_id_scope)
        if resource_types_scopes is not None:
            pulumi.set(__self__, "resource_types_scopes", resource_types_scopes)
        if rule_identifier is not None:
            pulumi.set(__self__, "rule_identifier", rule_identifier)
        if tag_key_scope is not None:
            pulumi.set(__self__, "tag_key_scope", tag_key_scope)
        if tag_value_scope is not None:
            pulumi.set(__self__, "tag_value_scope", tag_value_scope)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the rule
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="excludedAccounts")
    def excluded_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of AWS account identifiers to exclude from the rule
        """
        return pulumi.get(self, "excluded_accounts")

    @excluded_accounts.setter
    def excluded_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_accounts", value)

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> Optional[pulumi.Input[str]]:
        """
        A string in JSON format that is passed to the AWS Config Rule Lambda Function
        """
        return pulumi.get(self, "input_parameters")

    @input_parameters.setter
    def input_parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_parameters", value)

    @property
    @pulumi.getter(name="maximumExecutionFrequency")
    def maximum_execution_frequency(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        """
        return pulumi.get(self, "maximum_execution_frequency")

    @maximum_execution_frequency.setter
    def maximum_execution_frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maximum_execution_frequency", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceIdScope")
    def resource_id_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the AWS resource to evaluate
        """
        return pulumi.get(self, "resource_id_scope")

    @resource_id_scope.setter
    def resource_id_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id_scope", value)

    @property
    @pulumi.getter(name="resourceTypesScopes")
    def resource_types_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of types of AWS resources to evaluate
        """
        return pulumi.get(self, "resource_types_scopes")

    @resource_types_scopes.setter
    def resource_types_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_types_scopes", value)

    @property
    @pulumi.getter(name="ruleIdentifier")
    def rule_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        """
        return pulumi.get(self, "rule_identifier")

    @rule_identifier.setter
    def rule_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_identifier", value)

    @property
    @pulumi.getter(name="tagKeyScope")
    def tag_key_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Tag key of AWS resources to evaluate
        """
        return pulumi.get(self, "tag_key_scope")

    @tag_key_scope.setter
    def tag_key_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_key_scope", value)

    @property
    @pulumi.getter(name="tagValueScope")
    def tag_value_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Tag value of AWS resources to evaluate
        """
        return pulumi.get(self, "tag_value_scope")

    @tag_value_scope.setter
    def tag_value_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_value_scope", value)


class OrganizationManagedRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 input_parameters: Optional[pulumi.Input[str]] = None,
                 maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id_scope: Optional[pulumi.Input[str]] = None,
                 resource_types_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rule_identifier: Optional[pulumi.Input[str]] = None,
                 tag_key_scope: Optional[pulumi.Input[str]] = None,
                 tag_value_scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Config Organization Managed Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Custom Rules (those invoking a custom Lambda Function), see the `cfg.OrganizationCustomRule` resource.

        > **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

        > **NOTE:** Every Organization account except those configured in the `excluded_accounts` argument must have a Configuration Recorder with proper IAM permissions before the rule will successfully create or update. See also the `cfg.Recorder` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_organization = aws.organizations.Organization("exampleOrganization",
            aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
            feature_set="ALL")
        example_organization_managed_rule = aws.cfg.OrganizationManagedRule("exampleOrganizationManagedRule", rule_identifier="IAM_PASSWORD_POLICY",
        opts=pulumi.ResourceOptions(depends_on=[example_organization]))
        ```

        ## Import

        Config Organization Managed Rules can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:cfg/organizationManagedRule:OrganizationManagedRule example example
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] rule_identifier: Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: OrganizationManagedRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Config Organization Managed Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Custom Rules (those invoking a custom Lambda Function), see the `cfg.OrganizationCustomRule` resource.

        > **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

        > **NOTE:** Every Organization account except those configured in the `excluded_accounts` argument must have a Configuration Recorder with proper IAM permissions before the rule will successfully create or update. See also the `cfg.Recorder` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_organization = aws.organizations.Organization("exampleOrganization",
            aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
            feature_set="ALL")
        example_organization_managed_rule = aws.cfg.OrganizationManagedRule("exampleOrganizationManagedRule", rule_identifier="IAM_PASSWORD_POLICY",
        opts=pulumi.ResourceOptions(depends_on=[example_organization]))
        ```

        ## Import

        Config Organization Managed Rules can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:cfg/organizationManagedRule:OrganizationManagedRule example example
        ```

        :param str resource_name_: The name of the resource.
        :param OrganizationManagedRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationManagedRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 input_parameters: Optional[pulumi.Input[str]] = None,
                 maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id_scope: Optional[pulumi.Input[str]] = None,
                 resource_types_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rule_identifier: Optional[pulumi.Input[str]] = None,
                 tag_key_scope: Optional[pulumi.Input[str]] = None,
                 tag_value_scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationManagedRuleArgs.__new__(OrganizationManagedRuleArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["excluded_accounts"] = excluded_accounts
            __props__.__dict__["input_parameters"] = input_parameters
            __props__.__dict__["maximum_execution_frequency"] = maximum_execution_frequency
            __props__.__dict__["name"] = name
            __props__.__dict__["resource_id_scope"] = resource_id_scope
            __props__.__dict__["resource_types_scopes"] = resource_types_scopes
            if rule_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'rule_identifier'")
            __props__.__dict__["rule_identifier"] = rule_identifier
            __props__.__dict__["tag_key_scope"] = tag_key_scope
            __props__.__dict__["tag_value_scope"] = tag_value_scope
            __props__.__dict__["arn"] = None
        super(OrganizationManagedRule, __self__).__init__(
            'aws:cfg/organizationManagedRule:OrganizationManagedRule',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            excluded_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            input_parameters: Optional[pulumi.Input[str]] = None,
            maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_id_scope: Optional[pulumi.Input[str]] = None,
            resource_types_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            rule_identifier: Optional[pulumi.Input[str]] = None,
            tag_key_scope: Optional[pulumi.Input[str]] = None,
            tag_value_scope: Optional[pulumi.Input[str]] = None) -> 'OrganizationManagedRule':
        """
        Get an existing OrganizationManagedRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the rule
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] rule_identifier: Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationManagedRuleState.__new__(_OrganizationManagedRuleState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["excluded_accounts"] = excluded_accounts
        __props__.__dict__["input_parameters"] = input_parameters
        __props__.__dict__["maximum_execution_frequency"] = maximum_execution_frequency
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_id_scope"] = resource_id_scope
        __props__.__dict__["resource_types_scopes"] = resource_types_scopes
        __props__.__dict__["rule_identifier"] = rule_identifier
        __props__.__dict__["tag_key_scope"] = tag_key_scope
        __props__.__dict__["tag_value_scope"] = tag_value_scope
        return OrganizationManagedRule(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the rule
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the rule
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="excludedAccounts")
    def excluded_accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of AWS account identifiers to exclude from the rule
        """
        return pulumi.get(self, "excluded_accounts")

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> pulumi.Output[Optional[str]]:
        """
        A string in JSON format that is passed to the AWS Config Rule Lambda Function
        """
        return pulumi.get(self, "input_parameters")

    @property
    @pulumi.getter(name="maximumExecutionFrequency")
    def maximum_execution_frequency(self) -> pulumi.Output[Optional[str]]:
        """
        The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        """
        return pulumi.get(self, "maximum_execution_frequency")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the rule
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceIdScope")
    def resource_id_scope(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of the AWS resource to evaluate
        """
        return pulumi.get(self, "resource_id_scope")

    @property
    @pulumi.getter(name="resourceTypesScopes")
    def resource_types_scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of types of AWS resources to evaluate
        """
        return pulumi.get(self, "resource_types_scopes")

    @property
    @pulumi.getter(name="ruleIdentifier")
    def rule_identifier(self) -> pulumi.Output[str]:
        """
        Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        """
        return pulumi.get(self, "rule_identifier")

    @property
    @pulumi.getter(name="tagKeyScope")
    def tag_key_scope(self) -> pulumi.Output[Optional[str]]:
        """
        Tag key of AWS resources to evaluate
        """
        return pulumi.get(self, "tag_key_scope")

    @property
    @pulumi.getter(name="tagValueScope")
    def tag_value_scope(self) -> pulumi.Output[Optional[str]]:
        """
        Tag value of AWS resources to evaluate
        """
        return pulumi.get(self, "tag_value_scope")


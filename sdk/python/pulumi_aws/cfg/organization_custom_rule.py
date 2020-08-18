# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['OrganizationCustomRule']


class OrganizationCustomRule(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    Amazon Resource Name (ARN) of the rule
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    Description of the rule
    """

    excluded_accounts: pulumi.Output[Optional[List[str]]] = pulumi.property("excludedAccounts")
    """
    List of AWS account identifiers to exclude from the rule
    """

    input_parameters: pulumi.Output[Optional[str]] = pulumi.property("inputParameters")
    """
    A string in JSON format that is passed to the AWS Config Rule Lambda Function
    """

    lambda_function_arn: pulumi.Output[str] = pulumi.property("lambdaFunctionArn")
    """
    Amazon Resource Name (ARN) of the rule Lambda Function
    """

    maximum_execution_frequency: pulumi.Output[Optional[str]] = pulumi.property("maximumExecutionFrequency")
    """
    The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the rule
    """

    resource_id_scope: pulumi.Output[Optional[str]] = pulumi.property("resourceIdScope")
    """
    Identifier of the AWS resource to evaluate
    """

    resource_types_scopes: pulumi.Output[Optional[List[str]]] = pulumi.property("resourceTypesScopes")
    """
    List of types of AWS resources to evaluate
    """

    tag_key_scope: pulumi.Output[Optional[str]] = pulumi.property("tagKeyScope")
    """
    Tag key of AWS resources to evaluate
    """

    tag_value_scope: pulumi.Output[Optional[str]] = pulumi.property("tagValueScope")
    """
    Tag value of AWS resources to evaluate
    """

    trigger_types: pulumi.Output[List[str]] = pulumi.property("triggerTypes")
    """
    List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 input_parameters: Optional[pulumi.Input[str]] = None,
                 lambda_function_arn: Optional[pulumi.Input[str]] = None,
                 maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id_scope: Optional[pulumi.Input[str]] = None,
                 resource_types_scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tag_key_scope: Optional[pulumi.Input[str]] = None,
                 tag_value_scope: Optional[pulumi.Input[str]] = None,
                 trigger_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Config Organization Custom Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Managed Rules (those invoking an AWS managed rule), see the `aws_config_organization_managed__rule` resource.

        > **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

        > **NOTE:** The proper Lambda permission to allow the AWS Config service invoke the Lambda Function must be in place before the rule will successfully create or update. See also the `lambda.Permission` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_permission = aws.lambda_.Permission("examplePermission",
            action="lambda:InvokeFunction",
            function=aws_lambda_function["example"]["arn"],
            principal="config.amazonaws.com")
        example_organization = aws.organizations.Organization("exampleOrganization",
            aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
            feature_set="ALL")
        example_organization_custom_rule = aws.cfg.OrganizationCustomRule("exampleOrganizationCustomRule",
            lambda_function_arn=aws_lambda_function["example"]["arn"],
            trigger_types=["ConfigurationItemChangeNotification"],
            opts=ResourceOptions(depends_on=[
                    example_permission,
                    example_organization,
                ]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[List[pulumi.Input[str]]] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] lambda_function_arn: Amazon Resource Name (ARN) of the rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[List[pulumi.Input[str]]] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        :param pulumi.Input[List[pulumi.Input[str]]] trigger_types: List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
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
            if lambda_function_arn is None:
                raise TypeError("Missing required property 'lambda_function_arn'")
            __props__['lambda_function_arn'] = lambda_function_arn
            __props__['maximum_execution_frequency'] = maximum_execution_frequency
            __props__['name'] = name
            __props__['resource_id_scope'] = resource_id_scope
            __props__['resource_types_scopes'] = resource_types_scopes
            __props__['tag_key_scope'] = tag_key_scope
            __props__['tag_value_scope'] = tag_value_scope
            if trigger_types is None:
                raise TypeError("Missing required property 'trigger_types'")
            __props__['trigger_types'] = trigger_types
            __props__['arn'] = None
        super(OrganizationCustomRule, __self__).__init__(
            'aws:cfg/organizationCustomRule:OrganizationCustomRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            excluded_accounts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            input_parameters: Optional[pulumi.Input[str]] = None,
            lambda_function_arn: Optional[pulumi.Input[str]] = None,
            maximum_execution_frequency: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_id_scope: Optional[pulumi.Input[str]] = None,
            resource_types_scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tag_key_scope: Optional[pulumi.Input[str]] = None,
            tag_value_scope: Optional[pulumi.Input[str]] = None,
            trigger_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'OrganizationCustomRule':
        """
        Get an existing OrganizationCustomRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the rule
        :param pulumi.Input[str] description: Description of the rule
        :param pulumi.Input[List[pulumi.Input[str]]] excluded_accounts: List of AWS account identifiers to exclude from the rule
        :param pulumi.Input[str] input_parameters: A string in JSON format that is passed to the AWS Config Rule Lambda Function
        :param pulumi.Input[str] lambda_function_arn: Amazon Resource Name (ARN) of the rule Lambda Function
        :param pulumi.Input[str] maximum_execution_frequency: The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[str] resource_id_scope: Identifier of the AWS resource to evaluate
        :param pulumi.Input[List[pulumi.Input[str]]] resource_types_scopes: List of types of AWS resources to evaluate
        :param pulumi.Input[str] tag_key_scope: Tag key of AWS resources to evaluate
        :param pulumi.Input[str] tag_value_scope: Tag value of AWS resources to evaluate
        :param pulumi.Input[List[pulumi.Input[str]]] trigger_types: List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["excluded_accounts"] = excluded_accounts
        __props__["input_parameters"] = input_parameters
        __props__["lambda_function_arn"] = lambda_function_arn
        __props__["maximum_execution_frequency"] = maximum_execution_frequency
        __props__["name"] = name
        __props__["resource_id_scope"] = resource_id_scope
        __props__["resource_types_scopes"] = resource_types_scopes
        __props__["tag_key_scope"] = tag_key_scope
        __props__["tag_value_scope"] = tag_value_scope
        __props__["trigger_types"] = trigger_types
        return OrganizationCustomRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


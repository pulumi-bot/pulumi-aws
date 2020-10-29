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

__all__ = ['ReceiptRule']


class ReceiptRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_header_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleAddHeaderActionArgs']]]]] = None,
                 after: Optional[pulumi.Input[str]] = None,
                 bounce_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleBounceActionArgs']]]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 lambda_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleLambdaActionArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rule_set_name: Optional[pulumi.Input[str]] = None,
                 s3_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleS3ActionArgs']]]]] = None,
                 scan_enabled: Optional[pulumi.Input[bool]] = None,
                 sns_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleSnsActionArgs']]]]] = None,
                 stop_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleStopActionArgs']]]]] = None,
                 tls_policy: Optional[pulumi.Input[str]] = None,
                 workmail_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleWorkmailActionArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an SES receipt rule resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleAddHeaderActionArgs']]]] add_header_actions: A list of Add Header Action blocks. Documented below.
        :param pulumi.Input[str] after: The name of the rule to place this rule after
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleBounceActionArgs']]]] bounce_actions: A list of Bounce Action blocks. Documented below.
        :param pulumi.Input[bool] enabled: If true, the rule will be enabled
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleLambdaActionArgs']]]] lambda_actions: A list of Lambda Action blocks. Documented below.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: A list of email addresses
        :param pulumi.Input[str] rule_set_name: The name of the rule set
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleS3ActionArgs']]]] s3_actions: A list of S3 Action blocks. Documented below.
        :param pulumi.Input[bool] scan_enabled: If true, incoming emails will be scanned for spam and viruses
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleSnsActionArgs']]]] sns_actions: A list of SNS Action blocks. Documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleStopActionArgs']]]] stop_actions: A list of Stop Action blocks. Documented below.
        :param pulumi.Input[str] tls_policy: Require or Optional
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleWorkmailActionArgs']]]] workmail_actions: A list of WorkMail Action blocks. Documented below.
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

            __props__['add_header_actions'] = add_header_actions
            __props__['after'] = after
            __props__['bounce_actions'] = bounce_actions
            __props__['enabled'] = enabled
            __props__['lambda_actions'] = lambda_actions
            __props__['name'] = name
            __props__['recipients'] = recipients
            if rule_set_name is None:
                raise TypeError("Missing required property 'rule_set_name'")
            __props__['rule_set_name'] = rule_set_name
            __props__['s3_actions'] = s3_actions
            __props__['scan_enabled'] = scan_enabled
            __props__['sns_actions'] = sns_actions
            __props__['stop_actions'] = stop_actions
            __props__['tls_policy'] = tls_policy
            __props__['workmail_actions'] = workmail_actions
        super(ReceiptRule, __self__).__init__(
            'aws:ses/receiptRule:ReceiptRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_header_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleAddHeaderActionArgs']]]]] = None,
            after: Optional[pulumi.Input[str]] = None,
            bounce_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleBounceActionArgs']]]]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            lambda_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleLambdaActionArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            rule_set_name: Optional[pulumi.Input[str]] = None,
            s3_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleS3ActionArgs']]]]] = None,
            scan_enabled: Optional[pulumi.Input[bool]] = None,
            sns_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleSnsActionArgs']]]]] = None,
            stop_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleStopActionArgs']]]]] = None,
            tls_policy: Optional[pulumi.Input[str]] = None,
            workmail_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleWorkmailActionArgs']]]]] = None) -> 'ReceiptRule':
        """
        Get an existing ReceiptRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleAddHeaderActionArgs']]]] add_header_actions: A list of Add Header Action blocks. Documented below.
        :param pulumi.Input[str] after: The name of the rule to place this rule after
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleBounceActionArgs']]]] bounce_actions: A list of Bounce Action blocks. Documented below.
        :param pulumi.Input[bool] enabled: If true, the rule will be enabled
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleLambdaActionArgs']]]] lambda_actions: A list of Lambda Action blocks. Documented below.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: A list of email addresses
        :param pulumi.Input[str] rule_set_name: The name of the rule set
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleS3ActionArgs']]]] s3_actions: A list of S3 Action blocks. Documented below.
        :param pulumi.Input[bool] scan_enabled: If true, incoming emails will be scanned for spam and viruses
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleSnsActionArgs']]]] sns_actions: A list of SNS Action blocks. Documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleStopActionArgs']]]] stop_actions: A list of Stop Action blocks. Documented below.
        :param pulumi.Input[str] tls_policy: Require or Optional
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReceiptRuleWorkmailActionArgs']]]] workmail_actions: A list of WorkMail Action blocks. Documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_header_actions"] = add_header_actions
        __props__["after"] = after
        __props__["bounce_actions"] = bounce_actions
        __props__["enabled"] = enabled
        __props__["lambda_actions"] = lambda_actions
        __props__["name"] = name
        __props__["recipients"] = recipients
        __props__["rule_set_name"] = rule_set_name
        __props__["s3_actions"] = s3_actions
        __props__["scan_enabled"] = scan_enabled
        __props__["sns_actions"] = sns_actions
        __props__["stop_actions"] = stop_actions
        __props__["tls_policy"] = tls_policy
        __props__["workmail_actions"] = workmail_actions
        return ReceiptRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addHeaderActions")
    def add_header_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleAddHeaderAction']]]:
        """
        A list of Add Header Action blocks. Documented below.
        """
        return pulumi.get(self, "add_header_actions")

    @property
    @pulumi.getter
    def after(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the rule to place this rule after
        """
        return pulumi.get(self, "after")

    @property
    @pulumi.getter(name="bounceActions")
    def bounce_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleBounceAction']]]:
        """
        A list of Bounce Action blocks. Documented below.
        """
        return pulumi.get(self, "bounce_actions")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        If true, the rule will be enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="lambdaActions")
    def lambda_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleLambdaAction']]]:
        """
        A list of Lambda Action blocks. Documented below.
        """
        return pulumi.get(self, "lambda_actions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the rule
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of email addresses
        """
        return pulumi.get(self, "recipients")

    @property
    @pulumi.getter(name="ruleSetName")
    def rule_set_name(self) -> pulumi.Output[str]:
        """
        The name of the rule set
        """
        return pulumi.get(self, "rule_set_name")

    @property
    @pulumi.getter(name="s3Actions")
    def s3_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleS3Action']]]:
        """
        A list of S3 Action blocks. Documented below.
        """
        return pulumi.get(self, "s3_actions")

    @property
    @pulumi.getter(name="scanEnabled")
    def scan_enabled(self) -> pulumi.Output[bool]:
        """
        If true, incoming emails will be scanned for spam and viruses
        """
        return pulumi.get(self, "scan_enabled")

    @property
    @pulumi.getter(name="snsActions")
    def sns_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleSnsAction']]]:
        """
        A list of SNS Action blocks. Documented below.
        """
        return pulumi.get(self, "sns_actions")

    @property
    @pulumi.getter(name="stopActions")
    def stop_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleStopAction']]]:
        """
        A list of Stop Action blocks. Documented below.
        """
        return pulumi.get(self, "stop_actions")

    @property
    @pulumi.getter(name="tlsPolicy")
    def tls_policy(self) -> pulumi.Output[str]:
        """
        Require or Optional
        """
        return pulumi.get(self, "tls_policy")

    @property
    @pulumi.getter(name="workmailActions")
    def workmail_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ReceiptRuleWorkmailAction']]]:
        """
        A list of WorkMail Action blocks. Documented below.
        """
        return pulumi.get(self, "workmail_actions")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


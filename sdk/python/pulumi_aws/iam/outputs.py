# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetGroupUserResult',
    'GetPolicyDocumentStatementResult',
    'GetPolicyDocumentStatementConditionResult',
    'GetPolicyDocumentStatementNotPrincipalResult',
    'GetPolicyDocumentStatementPrincipalResult',
]

@pulumi.output_type
class GetGroupUserResult(dict):
    def __init__(__self__, *,
                 arn: str,
                 path: str,
                 user_id: str,
                 user_name: str):
        """
        :param str arn: The Amazon Resource Name (ARN) specifying the iam user.
        :param str path: The path to the iam user.
        :param str user_id: The stable and unique string identifying the iam user.
        :param str user_name: The name of the iam user.
        """
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) specifying the iam user.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The path to the iam user.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The stable and unique string identifying the iam user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The name of the iam user.
        """
        return pulumi.get(self, "user_name")


@pulumi.output_type
class GetPolicyDocumentStatementResult(dict):
    def __init__(__self__, *,
                 actions: Optional[Sequence[str]] = None,
                 conditions: Optional[Sequence['outputs.GetPolicyDocumentStatementConditionResult']] = None,
                 effect: Optional[str] = None,
                 not_actions: Optional[Sequence[str]] = None,
                 not_principals: Optional[Sequence['outputs.GetPolicyDocumentStatementNotPrincipalResult']] = None,
                 not_resources: Optional[Sequence[str]] = None,
                 principals: Optional[Sequence['outputs.GetPolicyDocumentStatementPrincipalResult']] = None,
                 resources: Optional[Sequence[str]] = None,
                 sid: Optional[str] = None):
        """
        :param Sequence[str] actions: A list of actions that this statement either allows
               or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
        :param Sequence['GetPolicyDocumentStatementConditionArgs'] conditions: A nested configuration block (described below)
               that defines a further, possibly-service-specific condition that constrains
               whether this statement applies.
        :param str effect: Either "Allow" or "Deny", to specify whether this
               statement allows or denies the given actions. The default is "Allow".
        :param Sequence[str] not_actions: A list of actions that this statement does *not*
               apply to. Used to apply a policy statement to all actions *except* those
               listed.
        :param Sequence['GetPolicyDocumentStatementNotPrincipalArgs'] not_principals: Like `principals` except gives principals that
               the statement does *not* apply to.
        :param Sequence[str] not_resources: A list of resource ARNs that this statement
               does *not* apply to. Used to apply a policy statement to all resources
               *except* those listed.
        :param Sequence['GetPolicyDocumentStatementPrincipalArgs'] principals: A nested configuration block (described below)
               specifying a principal (or principal pattern) to which this statement applies.
        :param Sequence[str] resources: A list of resource ARNs that this statement applies
               to. This is required by AWS if used for an IAM policy.
        :param str sid: An ID for the policy statement.
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if effect is not None:
            pulumi.set(__self__, "effect", effect)
        if not_actions is not None:
            pulumi.set(__self__, "not_actions", not_actions)
        if not_principals is not None:
            pulumi.set(__self__, "not_principals", not_principals)
        if not_resources is not None:
            pulumi.set(__self__, "not_resources", not_resources)
        if principals is not None:
            pulumi.set(__self__, "principals", principals)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if sid is not None:
            pulumi.set(__self__, "sid", sid)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence[str]]:
        """
        A list of actions that this statement either allows
        or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence['outputs.GetPolicyDocumentStatementConditionResult']]:
        """
        A nested configuration block (described below)
        that defines a further, possibly-service-specific condition that constrains
        whether this statement applies.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def effect(self) -> Optional[str]:
        """
        Either "Allow" or "Deny", to specify whether this
        statement allows or denies the given actions. The default is "Allow".
        """
        return pulumi.get(self, "effect")

    @property
    @pulumi.getter(name="notActions")
    def not_actions(self) -> Optional[Sequence[str]]:
        """
        A list of actions that this statement does *not*
        apply to. Used to apply a policy statement to all actions *except* those
        listed.
        """
        return pulumi.get(self, "not_actions")

    @property
    @pulumi.getter(name="notPrincipals")
    def not_principals(self) -> Optional[Sequence['outputs.GetPolicyDocumentStatementNotPrincipalResult']]:
        """
        Like `principals` except gives principals that
        the statement does *not* apply to.
        """
        return pulumi.get(self, "not_principals")

    @property
    @pulumi.getter(name="notResources")
    def not_resources(self) -> Optional[Sequence[str]]:
        """
        A list of resource ARNs that this statement
        does *not* apply to. Used to apply a policy statement to all resources
        *except* those listed.
        """
        return pulumi.get(self, "not_resources")

    @property
    @pulumi.getter
    def principals(self) -> Optional[Sequence['outputs.GetPolicyDocumentStatementPrincipalResult']]:
        """
        A nested configuration block (described below)
        specifying a principal (or principal pattern) to which this statement applies.
        """
        return pulumi.get(self, "principals")

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence[str]]:
        """
        A list of resource ARNs that this statement applies
        to. This is required by AWS if used for an IAM policy.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def sid(self) -> Optional[str]:
        """
        An ID for the policy statement.
        """
        return pulumi.get(self, "sid")


@pulumi.output_type
class GetPolicyDocumentStatementConditionResult(dict):
    def __init__(__self__, *,
                 test: str,
                 values: Sequence[str],
                 variable: str):
        """
        :param str test: The name of the
               [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
               to evaluate.
        :param Sequence[str] values: The values to evaluate the condition against. If multiple
               values are provided, the condition matches if at least one of them applies.
               (That is, the tests are combined with the "OR" boolean operation.)
        :param str variable: The name of a
               [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
               to apply the condition to. Context variables may either be standard AWS
               variables starting with `aws:`, or service-specific variables prefixed with
               the service name.
        """
        pulumi.set(__self__, "test", test)
        pulumi.set(__self__, "values", values)
        pulumi.set(__self__, "variable", variable)

    @property
    @pulumi.getter
    def test(self) -> str:
        """
        The name of the
        [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
        to evaluate.
        """
        return pulumi.get(self, "test")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        The values to evaluate the condition against. If multiple
        values are provided, the condition matches if at least one of them applies.
        (That is, the tests are combined with the "OR" boolean operation.)
        """
        return pulumi.get(self, "values")

    @property
    @pulumi.getter
    def variable(self) -> str:
        """
        The name of a
        [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
        to apply the condition to. Context variables may either be standard AWS
        variables starting with `aws:`, or service-specific variables prefixed with
        the service name.
        """
        return pulumi.get(self, "variable")


@pulumi.output_type
class GetPolicyDocumentStatementNotPrincipalResult(dict):
    def __init__(__self__, *,
                 identifiers: Sequence[str],
                 type: str):
        """
        :param Sequence[str] identifiers: List of identifiers for principals. When `type`
               is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.
        :param str type: The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
        """
        pulumi.set(__self__, "identifiers", identifiers)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def identifiers(self) -> Sequence[str]:
        """
        List of identifiers for principals. When `type`
        is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.
        """
        return pulumi.get(self, "identifiers")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetPolicyDocumentStatementPrincipalResult(dict):
    def __init__(__self__, *,
                 identifiers: Sequence[str],
                 type: str):
        """
        :param Sequence[str] identifiers: List of identifiers for principals. When `type`
               is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.
        :param str type: The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
        """
        pulumi.set(__self__, "identifiers", identifiers)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def identifiers(self) -> Sequence[str]:
        """
        List of identifiers for principals. When `type`
        is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.
        """
        return pulumi.get(self, "identifiers")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
        """
        return pulumi.get(self, "type")



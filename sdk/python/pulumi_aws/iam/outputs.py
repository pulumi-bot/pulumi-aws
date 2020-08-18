# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetGroupUser',
    'GetPolicyDocumentStatement',
    'GetPolicyDocumentStatementCondition',
    'GetPolicyDocumentStatementNotPrincipal',
    'GetPolicyDocumentStatementPrincipal',
]

@pulumi.output_type
class GetGroupUser(dict):
    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) specifying the iam user.
        """
        ...

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The path to the iam user.
        """
        ...

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The stable and unique string identifying the iam user.
        """
        ...

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The name of the iam user.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPolicyDocumentStatement(dict):
    @property
    @pulumi.getter
    def actions(self) -> Optional[List[str]]:
        """
        A list of actions that this statement either allows
        or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
        """
        ...

    @property
    @pulumi.getter
    def conditions(self) -> Optional[List['outputs.GetPolicyDocumentStatementCondition']]:
        """
        A nested configuration block (described below)
        that defines a further, possibly-service-specific condition that constrains
        whether this statement applies.
        """
        ...

    @property
    @pulumi.getter
    def effect(self) -> Optional[str]:
        """
        Either "Allow" or "Deny", to specify whether this
        statement allows or denies the given actions. The default is "Allow".
        """
        ...

    @property
    @pulumi.getter(name="notActions")
    def not_actions(self) -> Optional[List[str]]:
        """
        A list of actions that this statement does *not*
        apply to. Used to apply a policy statement to all actions *except* those
        listed.
        """
        ...

    @property
    @pulumi.getter(name="notPrincipals")
    def not_principals(self) -> Optional[List['outputs.GetPolicyDocumentStatementNotPrincipal']]:
        """
        Like `principals` except gives principals that
        the statement does *not* apply to.
        """
        ...

    @property
    @pulumi.getter(name="notResources")
    def not_resources(self) -> Optional[List[str]]:
        """
        A list of resource ARNs that this statement
        does *not* apply to. Used to apply a policy statement to all resources
        *except* those listed.
        """
        ...

    @property
    @pulumi.getter
    def principals(self) -> Optional[List['outputs.GetPolicyDocumentStatementPrincipal']]:
        """
        A nested configuration block (described below)
        specifying a principal (or principal pattern) to which this statement applies.
        """
        ...

    @property
    @pulumi.getter
    def resources(self) -> Optional[List[str]]:
        """
        A list of resource ARNs that this statement applies
        to. This is required by AWS if used for an IAM policy.
        """
        ...

    @property
    @pulumi.getter
    def sid(self) -> Optional[str]:
        """
        An ID for the policy statement.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPolicyDocumentStatementCondition(dict):
    @property
    @pulumi.getter
    def test(self) -> str:
        """
        The name of the
        [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
        to evaluate.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        The values to evaluate the condition against. If multiple
        values are provided, the condition matches if at least one of them applies.
        (That is, the tests are combined with the "OR" boolean operation.)
        """
        ...

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
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPolicyDocumentStatementNotPrincipal(dict):
    @property
    @pulumi.getter
    def identifiers(self) -> List[str]:
        """
        List of identifiers for principals. When `type`
        is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPolicyDocumentStatementPrincipal(dict):
    @property
    @pulumi.getter
    def identifiers(self) -> List[str]:
        """
        List of identifiers for principals. When `type`
        is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop



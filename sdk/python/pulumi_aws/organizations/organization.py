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

__all__ = ['Organization']


class Organization(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_service_access_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enabled_policy_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 feature_set: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to create an organization.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] aws_service_access_principals: List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_policy_types: List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
        :param pulumi.Input[str] feature_set: Specify "ALL" (default) or "CONSOLIDATED_BILLING".
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

            __props__['aws_service_access_principals'] = aws_service_access_principals
            __props__['enabled_policy_types'] = enabled_policy_types
            __props__['feature_set'] = feature_set
            __props__['accounts'] = None
            __props__['arn'] = None
            __props__['master_account_arn'] = None
            __props__['master_account_email'] = None
            __props__['master_account_id'] = None
            __props__['non_master_accounts'] = None
            __props__['roots'] = None
        super(Organization, __self__).__init__(
            'aws:organizations/organization:Organization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationAccountArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_service_access_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enabled_policy_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            feature_set: Optional[pulumi.Input[str]] = None,
            master_account_arn: Optional[pulumi.Input[str]] = None,
            master_account_email: Optional[pulumi.Input[str]] = None,
            master_account_id: Optional[pulumi.Input[str]] = None,
            non_master_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationNonMasterAccountArgs']]]]] = None,
            roots: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationRootArgs']]]]] = None) -> 'Organization':
        """
        Get an existing Organization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationAccountArgs']]]] accounts: List of organization accounts including the master account. For a list excluding the master account, see the `non_master_accounts` attribute. All elements have these attributes:
        :param pulumi.Input[str] arn: ARN of the root
        :param pulumi.Input[Sequence[pulumi.Input[str]]] aws_service_access_principals: List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_policy_types: List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
        :param pulumi.Input[str] feature_set: Specify "ALL" (default) or "CONSOLIDATED_BILLING".
        :param pulumi.Input[str] master_account_arn: ARN of the master account
        :param pulumi.Input[str] master_account_email: Email address of the master account
        :param pulumi.Input[str] master_account_id: Identifier of the master account
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationNonMasterAccountArgs']]]] non_master_accounts: List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationRootArgs']]]] roots: List of organization roots. All elements have these attributes:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accounts"] = accounts
        __props__["arn"] = arn
        __props__["aws_service_access_principals"] = aws_service_access_principals
        __props__["enabled_policy_types"] = enabled_policy_types
        __props__["feature_set"] = feature_set
        __props__["master_account_arn"] = master_account_arn
        __props__["master_account_email"] = master_account_email
        __props__["master_account_id"] = master_account_id
        __props__["non_master_accounts"] = non_master_accounts
        __props__["roots"] = roots
        return Organization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accounts(self) -> pulumi.Output[Sequence['outputs.OrganizationAccount']]:
        """
        List of organization accounts including the master account. For a list excluding the master account, see the `non_master_accounts` attribute. All elements have these attributes:
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the root
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsServiceAccessPrincipals")
    def aws_service_access_principals(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
        """
        return pulumi.get(self, "aws_service_access_principals")

    @property
    @pulumi.getter(name="enabledPolicyTypes")
    def enabled_policy_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
        """
        return pulumi.get(self, "enabled_policy_types")

    @property
    @pulumi.getter(name="featureSet")
    def feature_set(self) -> pulumi.Output[Optional[str]]:
        """
        Specify "ALL" (default) or "CONSOLIDATED_BILLING".
        """
        return pulumi.get(self, "feature_set")

    @property
    @pulumi.getter(name="masterAccountArn")
    def master_account_arn(self) -> pulumi.Output[str]:
        """
        ARN of the master account
        """
        return pulumi.get(self, "master_account_arn")

    @property
    @pulumi.getter(name="masterAccountEmail")
    def master_account_email(self) -> pulumi.Output[str]:
        """
        Email address of the master account
        """
        return pulumi.get(self, "master_account_email")

    @property
    @pulumi.getter(name="masterAccountId")
    def master_account_id(self) -> pulumi.Output[str]:
        """
        Identifier of the master account
        """
        return pulumi.get(self, "master_account_id")

    @property
    @pulumi.getter(name="nonMasterAccounts")
    def non_master_accounts(self) -> pulumi.Output[Sequence['outputs.OrganizationNonMasterAccount']]:
        """
        List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
        """
        return pulumi.get(self, "non_master_accounts")

    @property
    @pulumi.getter
    def roots(self) -> pulumi.Output[Sequence['outputs.OrganizationRoot']]:
        """
        List of organization roots. All elements have these attributes:
        """
        return pulumi.get(self, "roots")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


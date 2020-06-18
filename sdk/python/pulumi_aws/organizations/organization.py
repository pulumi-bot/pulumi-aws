# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Organization(pulumi.CustomResource):
    accounts: pulumi.Output[list]
    """
    List of organization accounts including the master account. For a list excluding the master account, see the `non_master_accounts` attribute. All elements have these attributes:

      * `arn` (`str`) - ARN of the root
      * `email` (`str`) - Email of the account
      * `id` (`str`) - Identifier of the root
      * `name` (`str`) - The name of the policy type
      * `status` (`str`) - The status of the policy type as it relates to the associated root
    """
    arn: pulumi.Output[str]
    """
    ARN of the root
    """
    aws_service_access_principals: pulumi.Output[list]
    """
    List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
    """
    enabled_policy_types: pulumi.Output[list]
    """
    List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY` and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
    """
    feature_set: pulumi.Output[str]
    """
    Specify "ALL" (default) or "CONSOLIDATED_BILLING".
    """
    master_account_arn: pulumi.Output[str]
    """
    ARN of the master account
    """
    master_account_email: pulumi.Output[str]
    """
    Email address of the master account
    """
    master_account_id: pulumi.Output[str]
    """
    Identifier of the master account
    """
    non_master_accounts: pulumi.Output[list]
    """
    List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:

      * `arn` (`str`) - ARN of the root
      * `email` (`str`) - Email of the account
      * `id` (`str`) - Identifier of the root
      * `name` (`str`) - The name of the policy type
      * `status` (`str`) - The status of the policy type as it relates to the associated root
    """
    roots: pulumi.Output[list]
    """
    List of organization roots. All elements have these attributes:

      * `arn` (`str`) - ARN of the root
      * `id` (`str`) - Identifier of the root
      * `name` (`str`) - The name of the policy type
      * `policyTypes` (`list`) - List of policy types enabled for this root. All elements have these attributes:
        * `status` (`str`) - The status of the policy type as it relates to the associated root
        * `type` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, aws_service_access_principals=None, enabled_policy_types=None, feature_set=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        org = aws.organizations.Organization("org",
            aws_service_access_principals=[
                "cloudtrail.amazonaws.com",
                "config.amazonaws.com",
            ],
            feature_set="ALL")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] aws_service_access_principals: List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
        :param pulumi.Input[list] enabled_policy_types: List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY` and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, accounts=None, arn=None, aws_service_access_principals=None, enabled_policy_types=None, feature_set=None, master_account_arn=None, master_account_email=None, master_account_id=None, non_master_accounts=None, roots=None):
        """
        Get an existing Organization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] accounts: List of organization accounts including the master account. For a list excluding the master account, see the `non_master_accounts` attribute. All elements have these attributes:
        :param pulumi.Input[str] arn: ARN of the root
        :param pulumi.Input[list] aws_service_access_principals: List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
        :param pulumi.Input[list] enabled_policy_types: List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY` and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
        :param pulumi.Input[str] feature_set: Specify "ALL" (default) or "CONSOLIDATED_BILLING".
        :param pulumi.Input[str] master_account_arn: ARN of the master account
        :param pulumi.Input[str] master_account_email: Email address of the master account
        :param pulumi.Input[str] master_account_id: Identifier of the master account
        :param pulumi.Input[list] non_master_accounts: List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
        :param pulumi.Input[list] roots: List of organization roots. All elements have these attributes:

        The **accounts** object supports the following:

          * `arn` (`pulumi.Input[str]`) - ARN of the root
          * `email` (`pulumi.Input[str]`) - Email of the account
          * `id` (`pulumi.Input[str]`) - Identifier of the root
          * `name` (`pulumi.Input[str]`) - The name of the policy type
          * `status` (`pulumi.Input[str]`) - The status of the policy type as it relates to the associated root

        The **non_master_accounts** object supports the following:

          * `arn` (`pulumi.Input[str]`) - ARN of the root
          * `email` (`pulumi.Input[str]`) - Email of the account
          * `id` (`pulumi.Input[str]`) - Identifier of the root
          * `name` (`pulumi.Input[str]`) - The name of the policy type
          * `status` (`pulumi.Input[str]`) - The status of the policy type as it relates to the associated root

        The **roots** object supports the following:

          * `arn` (`pulumi.Input[str]`) - ARN of the root
          * `id` (`pulumi.Input[str]`) - Identifier of the root
          * `name` (`pulumi.Input[str]`) - The name of the policy type
          * `policyTypes` (`pulumi.Input[list]`) - List of policy types enabled for this root. All elements have these attributes:
            * `status` (`pulumi.Input[str]`) - The status of the policy type as it relates to the associated root
            * `type` (`pulumi.Input[str]`)
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

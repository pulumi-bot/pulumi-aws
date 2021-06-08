# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ConfigurationAggregatorArgs', 'ConfigurationAggregator']

@pulumi.input_type
class ConfigurationAggregatorArgs:
    def __init__(__self__, *,
                 account_aggregation_source: Optional[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ConfigurationAggregator resource.
        :param pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs'] account_aggregation_source: The account(s) to aggregate config data from as documented below.
        :param pulumi.Input[str] name: The name of the configuration aggregator.
        :param pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs'] organization_aggregation_source: The organization to aggregate config data from as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if account_aggregation_source is not None:
            pulumi.set(__self__, "account_aggregation_source", account_aggregation_source)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_aggregation_source is not None:
            pulumi.set(__self__, "organization_aggregation_source", organization_aggregation_source)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="accountAggregationSource")
    def account_aggregation_source(self) -> Optional[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]:
        """
        The account(s) to aggregate config data from as documented below.
        """
        return pulumi.get(self, "account_aggregation_source")

    @account_aggregation_source.setter
    def account_aggregation_source(self, value: Optional[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]):
        pulumi.set(self, "account_aggregation_source", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the configuration aggregator.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationAggregationSource")
    def organization_aggregation_source(self) -> Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']]:
        """
        The organization to aggregate config data from as documented below.
        """
        return pulumi.get(self, "organization_aggregation_source")

    @organization_aggregation_source.setter
    def organization_aggregation_source(self, value: Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']]):
        pulumi.set(self, "organization_aggregation_source", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _ConfigurationAggregatorState:
    def __init__(__self__, *,
                 account_aggregation_source: Optional[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ConfigurationAggregator resources.
        :param pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs'] account_aggregation_source: The account(s) to aggregate config data from as documented below.
        :param pulumi.Input[str] arn: The ARN of the aggregator
        :param pulumi.Input[str] name: The name of the configuration aggregator.
        :param pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs'] organization_aggregation_source: The organization to aggregate config data from as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if account_aggregation_source is not None:
            pulumi.set(__self__, "account_aggregation_source", account_aggregation_source)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_aggregation_source is not None:
            pulumi.set(__self__, "organization_aggregation_source", organization_aggregation_source)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="accountAggregationSource")
    def account_aggregation_source(self) -> Optional[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]:
        """
        The account(s) to aggregate config data from as documented below.
        """
        return pulumi.get(self, "account_aggregation_source")

    @account_aggregation_source.setter
    def account_aggregation_source(self, value: Optional[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]):
        pulumi.set(self, "account_aggregation_source", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the aggregator
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the configuration aggregator.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationAggregationSource")
    def organization_aggregation_source(self) -> Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']]:
        """
        The organization to aggregate config data from as documented below.
        """
        return pulumi.get(self, "organization_aggregation_source")

    @organization_aggregation_source.setter
    def organization_aggregation_source(self, value: Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']]):
        pulumi.set(self, "organization_aggregation_source", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class ConfigurationAggregator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_aggregation_source: Optional[pulumi.Input[pulumi.InputType['ConfigurationAggregatorAccountAggregationSourceArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input[pulumi.InputType['ConfigurationAggregatorOrganizationAggregationSourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an AWS Config Configuration Aggregator

        ## Example Usage
        ### Account Based Aggregation

        ```python
        import pulumi
        import pulumi_aws as aws

        account = aws.cfg.ConfigurationAggregator("account", account_aggregation_source=aws.cfg.ConfigurationAggregatorAccountAggregationSourceArgs(
            account_ids=["123456789012"],
            regions=["us-west-2"],
        ))
        ```
        ### Organization Based Aggregation

        ```python
        import pulumi
        import pulumi_aws as aws

        organization_role = aws.iam.Role("organizationRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "config.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        organization_role_policy_attachment = aws.iam.RolePolicyAttachment("organizationRolePolicyAttachment",
            role=organization_role.name,
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations")
        organization_configuration_aggregator = aws.cfg.ConfigurationAggregator("organizationConfigurationAggregator", organization_aggregation_source=aws.cfg.ConfigurationAggregatorOrganizationAggregationSourceArgs(
            all_regions=True,
            role_arn=organization_role.arn,
        ),
        opts=pulumi.ResourceOptions(depends_on=[organization_role_policy_attachment]))
        ```

        ## Import

        Configuration Aggregators can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:cfg/configurationAggregator:ConfigurationAggregator example foo
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConfigurationAggregatorAccountAggregationSourceArgs']] account_aggregation_source: The account(s) to aggregate config data from as documented below.
        :param pulumi.Input[str] name: The name of the configuration aggregator.
        :param pulumi.Input[pulumi.InputType['ConfigurationAggregatorOrganizationAggregationSourceArgs']] organization_aggregation_source: The organization to aggregate config data from as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: Optional[ConfigurationAggregatorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an AWS Config Configuration Aggregator

        ## Example Usage
        ### Account Based Aggregation

        ```python
        import pulumi
        import pulumi_aws as aws

        account = aws.cfg.ConfigurationAggregator("account", account_aggregation_source=aws.cfg.ConfigurationAggregatorAccountAggregationSourceArgs(
            account_ids=["123456789012"],
            regions=["us-west-2"],
        ))
        ```
        ### Organization Based Aggregation

        ```python
        import pulumi
        import pulumi_aws as aws

        organization_role = aws.iam.Role("organizationRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "config.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        organization_role_policy_attachment = aws.iam.RolePolicyAttachment("organizationRolePolicyAttachment",
            role=organization_role.name,
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations")
        organization_configuration_aggregator = aws.cfg.ConfigurationAggregator("organizationConfigurationAggregator", organization_aggregation_source=aws.cfg.ConfigurationAggregatorOrganizationAggregationSourceArgs(
            all_regions=True,
            role_arn=organization_role.arn,
        ),
        opts=pulumi.ResourceOptions(depends_on=[organization_role_policy_attachment]))
        ```

        ## Import

        Configuration Aggregators can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:cfg/configurationAggregator:ConfigurationAggregator example foo
        ```

        :param str resource_name_: The name of the resource.
        :param ConfigurationAggregatorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationAggregatorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_aggregation_source: Optional[pulumi.Input[pulumi.InputType['ConfigurationAggregatorAccountAggregationSourceArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input[pulumi.InputType['ConfigurationAggregatorOrganizationAggregationSourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = ConfigurationAggregatorArgs.__new__(ConfigurationAggregatorArgs)

            __props__.__dict__["account_aggregation_source"] = account_aggregation_source
            __props__.__dict__["name"] = name
            __props__.__dict__["organization_aggregation_source"] = organization_aggregation_source
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["arn"] = None
        super(ConfigurationAggregator, __self__).__init__(
            'aws:cfg/configurationAggregator:ConfigurationAggregator',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_aggregation_source: Optional[pulumi.Input[pulumi.InputType['ConfigurationAggregatorAccountAggregationSourceArgs']]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_aggregation_source: Optional[pulumi.Input[pulumi.InputType['ConfigurationAggregatorOrganizationAggregationSourceArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ConfigurationAggregator':
        """
        Get an existing ConfigurationAggregator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConfigurationAggregatorAccountAggregationSourceArgs']] account_aggregation_source: The account(s) to aggregate config data from as documented below.
        :param pulumi.Input[str] arn: The ARN of the aggregator
        :param pulumi.Input[str] name: The name of the configuration aggregator.
        :param pulumi.Input[pulumi.InputType['ConfigurationAggregatorOrganizationAggregationSourceArgs']] organization_aggregation_source: The organization to aggregate config data from as documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationAggregatorState.__new__(_ConfigurationAggregatorState)

        __props__.__dict__["account_aggregation_source"] = account_aggregation_source
        __props__.__dict__["arn"] = arn
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_aggregation_source"] = organization_aggregation_source
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return ConfigurationAggregator(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountAggregationSource")
    def account_aggregation_source(self) -> pulumi.Output[Optional['outputs.ConfigurationAggregatorAccountAggregationSource']]:
        """
        The account(s) to aggregate config data from as documented below.
        """
        return pulumi.get(self, "account_aggregation_source")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the aggregator
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the configuration aggregator.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationAggregationSource")
    def organization_aggregation_source(self) -> pulumi.Output[Optional['outputs.ConfigurationAggregatorOrganizationAggregationSource']]:
        """
        The organization to aggregate config data from as documented below.
        """
        return pulumi.get(self, "organization_aggregation_source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")


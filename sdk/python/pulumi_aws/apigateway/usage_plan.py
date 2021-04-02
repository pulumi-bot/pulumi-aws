# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['UsagePlanArgs', 'UsagePlan']

@pulumi.input_type
class UsagePlanArgs:
    def __init__(__self__, *,
                 api_stages: Optional[pulumi.Input[Sequence[pulumi.Input['UsagePlanApiStageArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 product_code: Optional[pulumi.Input[str]] = None,
                 quota_settings: Optional[pulumi.Input['UsagePlanQuotaSettingsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 throttle_settings: Optional[pulumi.Input['UsagePlanThrottleSettingsArgs']] = None):
        """
        The set of arguments for constructing a UsagePlan resource.
        :param pulumi.Input[Sequence[pulumi.Input['UsagePlanApiStageArgs']]] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input['UsagePlanQuotaSettingsArgs'] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input['UsagePlanThrottleSettingsArgs'] throttle_settings: The throttling limits of the usage plan.
        """
        if api_stages is not None:
            pulumi.set(__self__, "api_stages", api_stages)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if product_code is not None:
            pulumi.set(__self__, "product_code", product_code)
        if quota_settings is not None:
            pulumi.set(__self__, "quota_settings", quota_settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if throttle_settings is not None:
            pulumi.set(__self__, "throttle_settings", throttle_settings)

    @property
    @pulumi.getter(name="apiStages")
    def api_stages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UsagePlanApiStageArgs']]]]:
        """
        The associated API stages of the usage plan.
        """
        return pulumi.get(self, "api_stages")

    @api_stages.setter
    def api_stages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UsagePlanApiStageArgs']]]]):
        pulumi.set(self, "api_stages", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of a usage plan.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the usage plan.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="productCode")
    def product_code(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        """
        return pulumi.get(self, "product_code")

    @product_code.setter
    def product_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_code", value)

    @property
    @pulumi.getter(name="quotaSettings")
    def quota_settings(self) -> Optional[pulumi.Input['UsagePlanQuotaSettingsArgs']]:
        """
        The quota settings of the usage plan.
        """
        return pulumi.get(self, "quota_settings")

    @quota_settings.setter
    def quota_settings(self, value: Optional[pulumi.Input['UsagePlanQuotaSettingsArgs']]):
        pulumi.set(self, "quota_settings", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="throttleSettings")
    def throttle_settings(self) -> Optional[pulumi.Input['UsagePlanThrottleSettingsArgs']]:
        """
        The throttling limits of the usage plan.
        """
        return pulumi.get(self, "throttle_settings")

    @throttle_settings.setter
    def throttle_settings(self, value: Optional[pulumi.Input['UsagePlanThrottleSettingsArgs']]):
        pulumi.set(self, "throttle_settings", value)


class UsagePlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_stages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UsagePlanApiStageArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 product_code: Optional[pulumi.Input[str]] = None,
                 quota_settings: Optional[pulumi.Input[pulumi.InputType['UsagePlanQuotaSettingsArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 throttle_settings: Optional[pulumi.Input[pulumi.InputType['UsagePlanThrottleSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an API Gateway Usage Plan.

        ## Import

        AWS API Gateway Usage Plan can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:apigateway/usagePlan:UsagePlan myusageplan <usage_plan_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UsagePlanApiStageArgs']]]] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input[pulumi.InputType['UsagePlanQuotaSettingsArgs']] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[pulumi.InputType['UsagePlanThrottleSettingsArgs']] throttle_settings: The throttling limits of the usage plan.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UsagePlanArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an API Gateway Usage Plan.

        ## Import

        AWS API Gateway Usage Plan can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:apigateway/usagePlan:UsagePlan myusageplan <usage_plan_id>
        ```

        :param str resource_name: The name of the resource.
        :param UsagePlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UsagePlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_stages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UsagePlanApiStageArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 product_code: Optional[pulumi.Input[str]] = None,
                 quota_settings: Optional[pulumi.Input[pulumi.InputType['UsagePlanQuotaSettingsArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 throttle_settings: Optional[pulumi.Input[pulumi.InputType['UsagePlanThrottleSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['api_stages'] = api_stages
            __props__['description'] = description
            __props__['name'] = name
            __props__['product_code'] = product_code
            __props__['quota_settings'] = quota_settings
            __props__['tags'] = tags
            __props__['throttle_settings'] = throttle_settings
            __props__['arn'] = None
        super(UsagePlan, __self__).__init__(
            'aws:apigateway/usagePlan:UsagePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_stages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UsagePlanApiStageArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            product_code: Optional[pulumi.Input[str]] = None,
            quota_settings: Optional[pulumi.Input[pulumi.InputType['UsagePlanQuotaSettingsArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            throttle_settings: Optional[pulumi.Input[pulumi.InputType['UsagePlanThrottleSettingsArgs']]] = None) -> 'UsagePlan':
        """
        Get an existing UsagePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UsagePlanApiStageArgs']]]] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input[pulumi.InputType['UsagePlanQuotaSettingsArgs']] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[pulumi.InputType['UsagePlanThrottleSettingsArgs']] throttle_settings: The throttling limits of the usage plan.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_stages"] = api_stages
        __props__["arn"] = arn
        __props__["description"] = description
        __props__["name"] = name
        __props__["product_code"] = product_code
        __props__["quota_settings"] = quota_settings
        __props__["tags"] = tags
        __props__["throttle_settings"] = throttle_settings
        return UsagePlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiStages")
    def api_stages(self) -> pulumi.Output[Optional[Sequence['outputs.UsagePlanApiStage']]]:
        """
        The associated API stages of the usage plan.
        """
        return pulumi.get(self, "api_stages")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of a usage plan.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the usage plan.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="productCode")
    def product_code(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        """
        return pulumi.get(self, "product_code")

    @property
    @pulumi.getter(name="quotaSettings")
    def quota_settings(self) -> pulumi.Output[Optional['outputs.UsagePlanQuotaSettings']]:
        """
        The quota settings of the usage plan.
        """
        return pulumi.get(self, "quota_settings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="throttleSettings")
    def throttle_settings(self) -> pulumi.Output[Optional['outputs.UsagePlanThrottleSettings']]:
        """
        The throttling limits of the usage plan.
        """
        return pulumi.get(self, "throttle_settings")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


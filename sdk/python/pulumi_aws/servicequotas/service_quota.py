# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServiceQuota(pulumi.CustomResource):
    adjustable: pulumi.Output[bool]
    """
    Whether the service quota can be increased.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the service quota.
    """
    default_value: pulumi.Output[float]
    """
    Default value of the service quota.
    """
    quota_code: pulumi.Output[str]
    """
    Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
    """
    quota_name: pulumi.Output[str]
    """
    Name of the quota.
    """
    request_id: pulumi.Output[str]
    request_status: pulumi.Output[str]
    service_code: pulumi.Output[str]
    """
    Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
    """
    service_name: pulumi.Output[str]
    """
    Name of the service.
    """
    value: pulumi.Output[float]
    """
    Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
    """
    def __init__(__self__, resource_name, opts=None, quota_code=None, service_code=None, value=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an individual Service Quota.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] quota_code: Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        :param pulumi.Input[str] service_code: Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
        :param pulumi.Input[float] value: Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/servicequotas_service_quota.html.markdown.
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

            if quota_code is None:
                raise TypeError("Missing required property 'quota_code'")
            __props__['quota_code'] = quota_code
            if service_code is None:
                raise TypeError("Missing required property 'service_code'")
            __props__['service_code'] = service_code
            if value is None:
                raise TypeError("Missing required property 'value'")
            __props__['value'] = value
            __props__['adjustable'] = None
            __props__['arn'] = None
            __props__['default_value'] = None
            __props__['quota_name'] = None
            __props__['request_id'] = None
            __props__['request_status'] = None
            __props__['service_name'] = None
        super(ServiceQuota, __self__).__init__(
            'aws:servicequotas/serviceQuota:ServiceQuota',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, adjustable=None, arn=None, default_value=None, quota_code=None, quota_name=None, request_id=None, request_status=None, service_code=None, service_name=None, value=None):
        """
        Get an existing ServiceQuota resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] adjustable: Whether the service quota can be increased.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the service quota.
        :param pulumi.Input[float] default_value: Default value of the service quota.
        :param pulumi.Input[str] quota_code: Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        :param pulumi.Input[str] quota_name: Name of the quota.
        :param pulumi.Input[str] service_code: Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
        :param pulumi.Input[str] service_name: Name of the service.
        :param pulumi.Input[float] value: Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/servicequotas_service_quota.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["adjustable"] = adjustable
        __props__["arn"] = arn
        __props__["default_value"] = default_value
        __props__["quota_code"] = quota_code
        __props__["quota_name"] = quota_name
        __props__["request_id"] = request_id
        __props__["request_status"] = request_status
        __props__["service_code"] = service_code
        __props__["service_name"] = service_name
        __props__["value"] = value
        return ServiceQuota(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


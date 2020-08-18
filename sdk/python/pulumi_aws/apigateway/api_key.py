# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ApiKey']


class ApiKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an API Gateway API Key.

        > **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api_key = aws.apigateway.ApiKey("myDemoApiKey")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The API key description. Defaults to "Managed by Pulumi".
        :param pulumi.Input[bool] enabled: Specifies whether the API key can be used by callers. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the API key
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] value: The value of the API key. If not specified, it will be automatically generated by AWS on creation.
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['value'] = value
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['last_updated_date'] = None
        super(ApiKey, __self__).__init__(
            'aws:apigateway/apiKey:ApiKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            created_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            last_updated_date: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ApiKey':
        """
        Get an existing ApiKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] created_date: The creation date of the API key
        :param pulumi.Input[str] description: The API key description. Defaults to "Managed by Pulumi".
        :param pulumi.Input[bool] enabled: Specifies whether the API key can be used by callers. Defaults to `true`.
        :param pulumi.Input[str] last_updated_date: The last update date of the API key
        :param pulumi.Input[str] name: The name of the API key
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] value: The value of the API key. If not specified, it will be automatically generated by AWS on creation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["created_date"] = created_date
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["last_updated_date"] = last_updated_date
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["value"] = value
        return ApiKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The creation date of the API key
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The API key description. Defaults to "Managed by Pulumi".
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Specifies whether the API key can be used by callers. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="lastUpdatedDate")
    def last_updated_date(self) -> str:
        """
        The last update date of the API key
        """
        return pulumi.get(self, "last_updated_date")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the API key
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the API key. If not specified, it will be automatically generated by AWS on creation.
        """
        return pulumi.get(self, "value")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


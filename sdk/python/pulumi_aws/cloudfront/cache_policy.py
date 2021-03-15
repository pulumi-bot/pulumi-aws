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

__all__ = ['CachePolicyArgs', 'CachePolicy']

@pulumi.input_type
class CachePolicyArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_ttl: Optional[pulumi.Input[int]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 min_ttl: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters_in_cache_key_and_forwarded_to_origin: Optional[pulumi.Input['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']] = None):
        """
        The set of arguments for constructing a CachePolicy resource.
        :param pulumi.Input[str] comment: A comment to describe the cache policy.
        :param pulumi.Input[int] default_ttl: The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[str] etag: The current version of the cache policy.
        :param pulumi.Input[int] max_ttl: The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[int] min_ttl: The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[str] name: A unique name to identify the cache policy.
        :param pulumi.Input['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs'] parameters_in_cache_key_and_forwarded_to_origin: The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_ttl is not None:
            pulumi.set(__self__, "default_ttl", default_ttl)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if min_ttl is not None:
            pulumi.set(__self__, "min_ttl", min_ttl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters_in_cache_key_and_forwarded_to_origin is not None:
            pulumi.set(__self__, "parameters_in_cache_key_and_forwarded_to_origin", parameters_in_cache_key_and_forwarded_to_origin)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        A comment to describe the cache policy.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultTtl")
    def default_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        """
        return pulumi.get(self, "default_ttl")

    @default_ttl.setter
    def default_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_ttl", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        The current version of the cache policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_ttl", value)

    @property
    @pulumi.getter(name="minTtl")
    def min_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        """
        return pulumi.get(self, "min_ttl")

    @min_ttl.setter
    def min_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ttl", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name to identify the cache policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parametersInCacheKeyAndForwardedToOrigin")
    def parameters_in_cache_key_and_forwarded_to_origin(self) -> Optional[pulumi.Input['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']]:
        """
        The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
        """
        return pulumi.get(self, "parameters_in_cache_key_and_forwarded_to_origin")

    @parameters_in_cache_key_and_forwarded_to_origin.setter
    def parameters_in_cache_key_and_forwarded_to_origin(self, value: Optional[pulumi.Input['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']]):
        pulumi.set(self, "parameters_in_cache_key_and_forwarded_to_origin", value)


class CachePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CachePolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        The following example below creates a CloudFront cache policy.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudfront.CachePolicy("example",
            comment="test comment",
            default_ttl=50,
            max_ttl=100,
            min_ttl=1,
            parameters_in_cache_key_and_forwarded_to_origin=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs(
                cookies_config={
                    "cookieBehavior": "whitelist",
                    "cookies": {
                        "items": ["example"],
                    },
                },
                headers_config={
                    "headerBehavior": "whitelist",
                    "headers": {
                        "items": ["example"],
                    },
                },
                query_strings_config={
                    "queryStringBehavior": "whitelist",
                    "queryStrings": {
                        "items": ["example"],
                    },
                },
            ))
        ```

        :param str resource_name: The name of the resource.
        :param CachePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_ttl: Optional[pulumi.Input[int]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 min_ttl: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters_in_cache_key_and_forwarded_to_origin: Optional[pulumi.Input[pulumi.InputType['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        The following example below creates a CloudFront cache policy.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudfront.CachePolicy("example",
            comment="test comment",
            default_ttl=50,
            max_ttl=100,
            min_ttl=1,
            parameters_in_cache_key_and_forwarded_to_origin=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs(
                cookies_config={
                    "cookieBehavior": "whitelist",
                    "cookies": {
                        "items": ["example"],
                    },
                },
                headers_config={
                    "headerBehavior": "whitelist",
                    "headers": {
                        "items": ["example"],
                    },
                },
                query_strings_config={
                    "queryStringBehavior": "whitelist",
                    "queryStrings": {
                        "items": ["example"],
                    },
                },
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: A comment to describe the cache policy.
        :param pulumi.Input[int] default_ttl: The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[str] etag: The current version of the cache policy.
        :param pulumi.Input[int] max_ttl: The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[int] min_ttl: The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[str] name: A unique name to identify the cache policy.
        :param pulumi.Input[pulumi.InputType['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']] parameters_in_cache_key_and_forwarded_to_origin: The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CachePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_ttl: Optional[pulumi.Input[int]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 min_ttl: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters_in_cache_key_and_forwarded_to_origin: Optional[pulumi.Input[pulumi.InputType['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']]] = None,
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

            __props__['comment'] = comment
            __props__['default_ttl'] = default_ttl
            __props__['etag'] = etag
            __props__['max_ttl'] = max_ttl
            __props__['min_ttl'] = min_ttl
            __props__['name'] = name
            __props__['parameters_in_cache_key_and_forwarded_to_origin'] = parameters_in_cache_key_and_forwarded_to_origin
        super(CachePolicy, __self__).__init__(
            'aws:cloudfront/cachePolicy:CachePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            default_ttl: Optional[pulumi.Input[int]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            max_ttl: Optional[pulumi.Input[int]] = None,
            min_ttl: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters_in_cache_key_and_forwarded_to_origin: Optional[pulumi.Input[pulumi.InputType['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']]] = None) -> 'CachePolicy':
        """
        Get an existing CachePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: A comment to describe the cache policy.
        :param pulumi.Input[int] default_ttl: The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[str] etag: The current version of the cache policy.
        :param pulumi.Input[int] max_ttl: The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[int] min_ttl: The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        :param pulumi.Input[str] name: A unique name to identify the cache policy.
        :param pulumi.Input[pulumi.InputType['CachePolicyParametersInCacheKeyAndForwardedToOriginArgs']] parameters_in_cache_key_and_forwarded_to_origin: The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["comment"] = comment
        __props__["default_ttl"] = default_ttl
        __props__["etag"] = etag
        __props__["max_ttl"] = max_ttl
        __props__["min_ttl"] = min_ttl
        __props__["name"] = name
        __props__["parameters_in_cache_key_and_forwarded_to_origin"] = parameters_in_cache_key_and_forwarded_to_origin
        return CachePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        A comment to describe the cache policy.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="defaultTtl")
    def default_ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        """
        return pulumi.get(self, "default_ttl")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        The current version of the cache policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter(name="minTtl")
    def min_ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        """
        return pulumi.get(self, "min_ttl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name to identify the cache policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parametersInCacheKeyAndForwardedToOrigin")
    def parameters_in_cache_key_and_forwarded_to_origin(self) -> pulumi.Output[Optional['outputs.CachePolicyParametersInCacheKeyAndForwardedToOrigin']]:
        """
        The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
        """
        return pulumi.get(self, "parameters_in_cache_key_and_forwarded_to_origin")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


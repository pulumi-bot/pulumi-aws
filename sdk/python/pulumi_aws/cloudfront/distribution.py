# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Distribution']


class Distribution(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aliases: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 custom_error_responses: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionCustomErrorResponseArgs']]]]] = None,
                 default_cache_behavior: Optional[pulumi.Input[pulumi.InputType['DistributionDefaultCacheBehaviorArgs']]] = None,
                 default_root_object: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 http_version: Optional[pulumi.Input[str]] = None,
                 is_ipv6_enabled: Optional[pulumi.Input[bool]] = None,
                 logging_config: Optional[pulumi.Input[pulumi.InputType['DistributionLoggingConfigArgs']]] = None,
                 ordered_cache_behaviors: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionOrderedCacheBehaviorArgs']]]]] = None,
                 origin_groups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionOriginGroupArgs']]]]] = None,
                 origins: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionOriginArgs']]]]] = None,
                 price_class: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['DistributionRestrictionsArgs']]] = None,
                 retain_on_delete: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 viewer_certificate: Optional[pulumi.Input[pulumi.InputType['DistributionViewerCertificateArgs']]] = None,
                 wait_for_deployment: Optional[pulumi.Input[bool]] = None,
                 web_acl_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Distribution resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            __props__['aliases'] = aliases
            __props__['comment'] = comment
            __props__['custom_error_responses'] = custom_error_responses
            if default_cache_behavior is None:
                raise TypeError("Missing required property 'default_cache_behavior'")
            __props__['default_cache_behavior'] = default_cache_behavior
            __props__['default_root_object'] = default_root_object
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            __props__['http_version'] = http_version
            __props__['is_ipv6_enabled'] = is_ipv6_enabled
            __props__['logging_config'] = logging_config
            __props__['ordered_cache_behaviors'] = ordered_cache_behaviors
            __props__['origin_groups'] = origin_groups
            if origins is None:
                raise TypeError("Missing required property 'origins'")
            __props__['origins'] = origins
            __props__['price_class'] = price_class
            if restrictions is None:
                raise TypeError("Missing required property 'restrictions'")
            __props__['restrictions'] = restrictions
            __props__['retain_on_delete'] = retain_on_delete
            __props__['tags'] = tags
            if viewer_certificate is None:
                raise TypeError("Missing required property 'viewer_certificate'")
            __props__['viewer_certificate'] = viewer_certificate
            __props__['wait_for_deployment'] = wait_for_deployment
            __props__['web_acl_id'] = web_acl_id
            __props__['arn'] = None
            __props__['caller_reference'] = None
            __props__['domain_name'] = None
            __props__['etag'] = None
            __props__['hosted_zone_id'] = None
            __props__['in_progress_validation_batches'] = None
            __props__['last_modified_time'] = None
            __props__['status'] = None
            __props__['trusted_signers'] = None
        super(Distribution, __self__).__init__(
            'aws:cloudfront/distribution:Distribution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aliases: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            caller_reference: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            custom_error_responses: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionCustomErrorResponseArgs']]]]] = None,
            default_cache_behavior: Optional[pulumi.Input[pulumi.InputType['DistributionDefaultCacheBehaviorArgs']]] = None,
            default_root_object: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            hosted_zone_id: Optional[pulumi.Input[str]] = None,
            http_version: Optional[pulumi.Input[str]] = None,
            in_progress_validation_batches: Optional[pulumi.Input[float]] = None,
            is_ipv6_enabled: Optional[pulumi.Input[bool]] = None,
            last_modified_time: Optional[pulumi.Input[str]] = None,
            logging_config: Optional[pulumi.Input[pulumi.InputType['DistributionLoggingConfigArgs']]] = None,
            ordered_cache_behaviors: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionOrderedCacheBehaviorArgs']]]]] = None,
            origin_groups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionOriginGroupArgs']]]]] = None,
            origins: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionOriginArgs']]]]] = None,
            price_class: Optional[pulumi.Input[str]] = None,
            restrictions: Optional[pulumi.Input[pulumi.InputType['DistributionRestrictionsArgs']]] = None,
            retain_on_delete: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            trusted_signers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DistributionTrustedSignerArgs']]]]] = None,
            viewer_certificate: Optional[pulumi.Input[pulumi.InputType['DistributionViewerCertificateArgs']]] = None,
            wait_for_deployment: Optional[pulumi.Input[bool]] = None,
            web_acl_id: Optional[pulumi.Input[str]] = None) -> 'Distribution':
        """
        Get an existing Distribution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["aliases"] = aliases
        __props__["arn"] = arn
        __props__["caller_reference"] = caller_reference
        __props__["comment"] = comment
        __props__["custom_error_responses"] = custom_error_responses
        __props__["default_cache_behavior"] = default_cache_behavior
        __props__["default_root_object"] = default_root_object
        __props__["domain_name"] = domain_name
        __props__["enabled"] = enabled
        __props__["etag"] = etag
        __props__["hosted_zone_id"] = hosted_zone_id
        __props__["http_version"] = http_version
        __props__["in_progress_validation_batches"] = in_progress_validation_batches
        __props__["is_ipv6_enabled"] = is_ipv6_enabled
        __props__["last_modified_time"] = last_modified_time
        __props__["logging_config"] = logging_config
        __props__["ordered_cache_behaviors"] = ordered_cache_behaviors
        __props__["origin_groups"] = origin_groups
        __props__["origins"] = origins
        __props__["price_class"] = price_class
        __props__["restrictions"] = restrictions
        __props__["retain_on_delete"] = retain_on_delete
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["trusted_signers"] = trusted_signers
        __props__["viewer_certificate"] = viewer_certificate
        __props__["wait_for_deployment"] = wait_for_deployment
        __props__["web_acl_id"] = web_acl_id
        return Distribution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def aliases(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "aliases")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="callerReference")
    def caller_reference(self) -> pulumi.Output[str]:
        return pulumi.get(self, "caller_reference")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="customErrorResponses")
    def custom_error_responses(self) -> pulumi.Output[Optional[List['outputs.DistributionCustomErrorResponse']]]:
        return pulumi.get(self, "custom_error_responses")

    @property
    @pulumi.getter(name="defaultCacheBehavior")
    def default_cache_behavior(self) -> pulumi.Output['outputs.DistributionDefaultCacheBehavior']:
        return pulumi.get(self, "default_cache_behavior")

    @property
    @pulumi.getter(name="defaultRootObject")
    def default_root_object(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "default_root_object")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter(name="httpVersion")
    def http_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "http_version")

    @property
    @pulumi.getter(name="inProgressValidationBatches")
    def in_progress_validation_batches(self) -> pulumi.Output[float]:
        return pulumi.get(self, "in_progress_validation_batches")

    @property
    @pulumi.getter(name="isIpv6Enabled")
    def is_ipv6_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "is_ipv6_enabled")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="loggingConfig")
    def logging_config(self) -> pulumi.Output[Optional['outputs.DistributionLoggingConfig']]:
        return pulumi.get(self, "logging_config")

    @property
    @pulumi.getter(name="orderedCacheBehaviors")
    def ordered_cache_behaviors(self) -> pulumi.Output[Optional[List['outputs.DistributionOrderedCacheBehavior']]]:
        return pulumi.get(self, "ordered_cache_behaviors")

    @property
    @pulumi.getter(name="originGroups")
    def origin_groups(self) -> pulumi.Output[Optional[List['outputs.DistributionOriginGroup']]]:
        return pulumi.get(self, "origin_groups")

    @property
    @pulumi.getter
    def origins(self) -> pulumi.Output[List['outputs.DistributionOrigin']]:
        return pulumi.get(self, "origins")

    @property
    @pulumi.getter(name="priceClass")
    def price_class(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "price_class")

    @property
    @pulumi.getter
    def restrictions(self) -> pulumi.Output['outputs.DistributionRestrictions']:
        return pulumi.get(self, "restrictions")

    @property
    @pulumi.getter(name="retainOnDelete")
    def retain_on_delete(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "retain_on_delete")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trustedSigners")
    def trusted_signers(self) -> pulumi.Output[List['outputs.DistributionTrustedSigner']]:
        return pulumi.get(self, "trusted_signers")

    @property
    @pulumi.getter(name="viewerCertificate")
    def viewer_certificate(self) -> pulumi.Output['outputs.DistributionViewerCertificate']:
        return pulumi.get(self, "viewer_certificate")

    @property
    @pulumi.getter(name="waitForDeployment")
    def wait_for_deployment(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "wait_for_deployment")

    @property
    @pulumi.getter(name="webAclId")
    def web_acl_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "web_acl_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


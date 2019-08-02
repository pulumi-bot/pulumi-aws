# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Distribution(pulumi.CustomResource):
    active_trusted_signers: pulumi.Output[dict]
    """
    The key pair IDs that CloudFront is aware of for
    each trusted signer, if the distribution is set up to serve private content
    with signed URLs.
    """
    aliases: pulumi.Output[list]
    """
    Extra CNAMEs (alternate domain names), if any, for
    this distribution.
    """
    arn: pulumi.Output[str]
    """
    The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
    """
    caller_reference: pulumi.Output[str]
    """
    Internal value used by CloudFront to allow future
    updates to the distribution configuration.
    """
    comment: pulumi.Output[str]
    """
    Any comments you want to include about the
    distribution.
    """
    custom_error_responses: pulumi.Output[list]
    """
    One or more custom error response elements (multiples allowed).
    """
    default_cache_behavior: pulumi.Output[dict]
    """
    The default cache behavior for this distribution (maximum
    one).
    """
    default_root_object: pulumi.Output[str]
    """
    The object that you want CloudFront to
    return (for example, index.html) when an end user requests the root URL.
    """
    domain_name: pulumi.Output[str]
    """
    The DNS domain name of either the S3 bucket, or
    web site of your custom origin.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the distribution is enabled to accept end
    user requests for content.
    """
    etag: pulumi.Output[str]
    """
    The current version of the distribution's information. For example:
    `E2QWRUHAPOMQZL`.
    """
    hosted_zone_id: pulumi.Output[str]
    """
    The CloudFront Route 53 zone ID that can be used to
    route an [Alias Resource Record Set][7] to. This attribute is simply an
    alias for the zone ID `Z2FDTNDATAQYW2`.
    """
    http_version: pulumi.Output[str]
    """
    The maximum HTTP version to support on the
    distribution. Allowed values are `http1.1` and `http2`. The default is
    `http2`.
    """
    in_progress_validation_batches: pulumi.Output[float]
    """
    The number of invalidation batches
    currently in progress.
    """
    is_ipv6_enabled: pulumi.Output[bool]
    """
    Whether the IPv6 is enabled for the distribution.
    """
    last_modified_time: pulumi.Output[str]
    """
    The date and time the distribution was last modified.
    """
    logging_config: pulumi.Output[dict]
    """
    The logging
    configuration that controls how logs are written
    to your distribution (maximum one).
    """
    ordered_cache_behaviors: pulumi.Output[list]
    """
    An ordered list of cache behaviors
    resource for this distribution. List from top to bottom
    in order of precedence. The topmost cache behavior will have precedence 0.
    """
    origins: pulumi.Output[list]
    """
    One or more origins for this
    distribution (multiples allowed).
    """
    origin_groups: pulumi.Output[list]
    """
    One or more origin_group for this
    distribution (multiples allowed).
    """
    price_class: pulumi.Output[str]
    """
    The price class for this distribution. One of
    `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
    """
    restrictions: pulumi.Output[dict]
    """
    The restriction
    configuration for this distribution (maximum one).
    """
    retain_on_delete: pulumi.Output[bool]
    """
    Disables the distribution instead of
    deleting it when destroying the resource. If this is set,
    the distribution needs to be deleted manually afterwards. Default: `false`.
    """
    status: pulumi.Output[str]
    """
    The current status of the distribution. `Deployed` if the
    distribution's information is fully propagated throughout the Amazon
    CloudFront system.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    viewer_certificate: pulumi.Output[dict]
    """
    The SSL
    configuration for this distribution (maximum
    one).
    """
    wait_for_deployment: pulumi.Output[bool]
    """
    If enabled, the resource will wait for
    the distribution status to change from `InProgress` to `Deployed`. Setting
    this to`false` will skip the process. Default: `true`.
    """
    web_acl_id: pulumi.Output[str]
    """
    If you're using AWS WAF to filter CloudFront
    requests, the Id of the AWS WAF web ACL that is associated with the
    distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
    region and the credentials configuring this argument must have
    `waf:GetWebACL` permissions assigned.
    """
    def __init__(__self__, resource_name, opts=None, aliases=None, comment=None, custom_error_responses=None, default_cache_behavior=None, default_root_object=None, enabled=None, http_version=None, is_ipv6_enabled=None, logging_config=None, ordered_cache_behaviors=None, origins=None, origin_groups=None, price_class=None, restrictions=None, retain_on_delete=None, tags=None, viewer_certificate=None, wait_for_deployment=None, web_acl_id=None, __name__=None, __opts__=None):
        """
        Creates an Amazon CloudFront web distribution.
        
        For information about CloudFront distributions, see the
        [Amazon CloudFront Developer Guide][1]. For specific information about creating
        CloudFront web distributions, see the [POST Distribution][2] page in the Amazon
        CloudFront API Reference.
        
        > **NOTE:** CloudFront distributions take about 15 minutes to a deployed state
        after creation or modification. During this time, deletes to resources will be
        blocked. If you need to delete a distribution that is enabled and you do not
        want to wait, you need to use the `retain_on_delete` flag.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] aliases: Extra CNAMEs (alternate domain names), if any, for
               this distribution.
        :param pulumi.Input[str] comment: Any comments you want to include about the
               distribution.
        :param pulumi.Input[list] custom_error_responses: One or more custom error response elements (multiples allowed).
        :param pulumi.Input[dict] default_cache_behavior: The default cache behavior for this distribution (maximum
               one).
        :param pulumi.Input[str] default_root_object: The object that you want CloudFront to
               return (for example, index.html) when an end user requests the root URL.
        :param pulumi.Input[bool] enabled: Whether the distribution is enabled to accept end
               user requests for content.
        :param pulumi.Input[str] http_version: The maximum HTTP version to support on the
               distribution. Allowed values are `http1.1` and `http2`. The default is
               `http2`.
        :param pulumi.Input[bool] is_ipv6_enabled: Whether the IPv6 is enabled for the distribution.
        :param pulumi.Input[dict] logging_config: The logging
               configuration that controls how logs are written
               to your distribution (maximum one).
        :param pulumi.Input[list] ordered_cache_behaviors: An ordered list of cache behaviors
               resource for this distribution. List from top to bottom
               in order of precedence. The topmost cache behavior will have precedence 0.
        :param pulumi.Input[list] origins: One or more origins for this
               distribution (multiples allowed).
        :param pulumi.Input[list] origin_groups: One or more origin_group for this
               distribution (multiples allowed).
        :param pulumi.Input[str] price_class: The price class for this distribution. One of
               `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
        :param pulumi.Input[dict] restrictions: The restriction
               configuration for this distribution (maximum one).
        :param pulumi.Input[bool] retain_on_delete: Disables the distribution instead of
               deleting it when destroying the resource. If this is set,
               the distribution needs to be deleted manually afterwards. Default: `false`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] viewer_certificate: The SSL
               configuration for this distribution (maximum
               one).
        :param pulumi.Input[bool] wait_for_deployment: If enabled, the resource will wait for
               the distribution status to change from `InProgress` to `Deployed`. Setting
               this to`false` will skip the process. Default: `true`.
        :param pulumi.Input[str] web_acl_id: If you're using AWS WAF to filter CloudFront
               requests, the Id of the AWS WAF web ACL that is associated with the
               distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
               region and the credentials configuring this argument must have
               `waf:GetWebACL` permissions assigned.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudfront_distribution.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

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

        if origins is None:
            raise TypeError("Missing required property 'origins'")
        __props__['origins'] = origins

        __props__['origin_groups'] = origin_groups

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

        __props__['active_trusted_signers'] = None
        __props__['arn'] = None
        __props__['caller_reference'] = None
        __props__['domain_name'] = None
        __props__['etag'] = None
        __props__['hosted_zone_id'] = None
        __props__['in_progress_validation_batches'] = None
        __props__['last_modified_time'] = None
        __props__['status'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Distribution, __self__).__init__(
            'aws:cloudfront/distribution:Distribution',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


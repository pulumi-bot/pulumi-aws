# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Service(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the service.
    """
    description: pulumi.Output[str]
    """
    The description of the service.
    """
    dns_config: pulumi.Output[dict]
    """
    A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
    """
    health_check_config: pulumi.Output[dict]
    """
    A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
    """
    health_check_custom_config: pulumi.Output[dict]
    """
    A complex type that contains settings for ECS managed health checks.
    """
    name: pulumi.Output[str]
    """
    The name of the service.
    """
    namespace_id: pulumi.Output[str]
    """
    The ID of the namespace to use for DNS configuration.
    """
    def __init__(__self__, resource_name, opts=None, description=None, dns_config=None, health_check_config=None, health_check_custom_config=None, name=None, namespace_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Service Discovery Service resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the service.
        :param pulumi.Input[dict] dns_config: A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
        :param pulumi.Input[dict] health_check_config: A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
        :param pulumi.Input[dict] health_check_custom_config: A complex type that contains settings for ECS managed health checks.
        :param pulumi.Input[str] name: The name of the service.
        :param pulumi.Input[str] namespace_id: The ID of the namespace to use for DNS configuration.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/service_discovery_service.html.markdown.
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

            __props__['description'] = description
            __props__['dns_config'] = dns_config
            __props__['health_check_config'] = health_check_config
            __props__['health_check_custom_config'] = health_check_custom_config
            __props__['name'] = name
            __props__['namespace_id'] = namespace_id
            __props__['arn'] = None
        super(Service, __self__).__init__(
            'aws:servicediscovery/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, dns_config=None, health_check_config=None, health_check_custom_config=None, name=None, namespace_id=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the service.
        :param pulumi.Input[str] description: The description of the service.
        :param pulumi.Input[dict] dns_config: A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
        :param pulumi.Input[dict] health_check_config: A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
        :param pulumi.Input[dict] health_check_custom_config: A complex type that contains settings for ECS managed health checks.
        :param pulumi.Input[str] name: The name of the service.
        :param pulumi.Input[str] namespace_id: The ID of the namespace to use for DNS configuration.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/service_discovery_service.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["description"] = description
        __props__["dns_config"] = dns_config
        __props__["health_check_config"] = health_check_config
        __props__["health_check_custom_config"] = health_check_custom_config
        __props__["name"] = name
        __props__["namespace_id"] = namespace_id
        return Service(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


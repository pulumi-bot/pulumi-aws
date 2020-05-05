# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Agent(pulumi.CustomResource):
    activation_key: pulumi.Output[str]
    """
    DataSync Agent activation key during resource creation. Conflicts with `ip_address`. If an `ip_address` is provided instead, the provider will retrieve the `activation_key` as part of the resource creation.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the DataSync Agent.
    """
    ip_address: pulumi.Output[str]
    """
    DataSync Agent IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. DataSync Agent must be accessible on port 80 from where the provider is running.
    """
    name: pulumi.Output[str]
    """
    Name of the DataSync Agent.
    """
    tags: pulumi.Output[dict]
    """
    Key-value pairs of resource tags to assign to the DataSync Agent.
    """
    def __init__(__self__, resource_name, opts=None, activation_key=None, ip_address=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an AWS DataSync Agent deployed on premises.

        > **NOTE:** One of `activation_key` or `ip_address` must be provided for resource creation (agent activation). Neither is required for resource import. If using `ip_address`, this provider must be able to make an HTTP (port 80) GET request to the specified IP address from where it is running. The agent will turn off that HTTP server after activation.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.datasync.Agent("example", ip_address="1.2.3.4")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] activation_key: DataSync Agent activation key during resource creation. Conflicts with `ip_address`. If an `ip_address` is provided instead, the provider will retrieve the `activation_key` as part of the resource creation.
        :param pulumi.Input[str] ip_address: DataSync Agent IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. DataSync Agent must be accessible on port 80 from where the provider is running.
        :param pulumi.Input[str] name: Name of the DataSync Agent.
        :param pulumi.Input[dict] tags: Key-value pairs of resource tags to assign to the DataSync Agent.
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

            __props__['activation_key'] = activation_key
            __props__['ip_address'] = ip_address
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Agent, __self__).__init__(
            'aws:datasync/agent:Agent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, activation_key=None, arn=None, ip_address=None, name=None, tags=None):
        """
        Get an existing Agent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] activation_key: DataSync Agent activation key during resource creation. Conflicts with `ip_address`. If an `ip_address` is provided instead, the provider will retrieve the `activation_key` as part of the resource creation.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the DataSync Agent.
        :param pulumi.Input[str] ip_address: DataSync Agent IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. DataSync Agent must be accessible on port 80 from where the provider is running.
        :param pulumi.Input[str] name: Name of the DataSync Agent.
        :param pulumi.Input[dict] tags: Key-value pairs of resource tags to assign to the DataSync Agent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["activation_key"] = activation_key
        __props__["arn"] = arn
        __props__["ip_address"] = ip_address
        __props__["name"] = name
        __props__["tags"] = tags
        return Agent(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


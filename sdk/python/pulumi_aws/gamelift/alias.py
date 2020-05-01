# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Alias(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Alias ARN.
    """
    description: pulumi.Output[str]
    """
    Description of the alias.
    """
    name: pulumi.Output[str]
    """
    Name of the alias.
    """
    routing_strategy: pulumi.Output[dict]
    """
    Specifies the fleet and/or routing type to use for the alias.

      * `fleetId` (`str`) - ID of the Gamelift Fleet to point the alias to.
      * `message` (`str`) - Message text to be used with the `TERMINAL` routing strategy.
      * `type` (`str`) - Type of routing strategy. e.g. `SIMPLE` or `TERMINAL`
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, routing_strategy=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Gamelift Alias resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.gamelift.Alias("example",
            description="Example Description",
            routing_strategy={
                "message": "Example Message",
                "type": "TERMINAL",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the alias.
        :param pulumi.Input[str] name: Name of the alias.
        :param pulumi.Input[dict] routing_strategy: Specifies the fleet and/or routing type to use for the alias.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **routing_strategy** object supports the following:

          * `fleetId` (`pulumi.Input[str]`) - ID of the Gamelift Fleet to point the alias to.
          * `message` (`pulumi.Input[str]`) - Message text to be used with the `TERMINAL` routing strategy.
          * `type` (`pulumi.Input[str]`) - Type of routing strategy. e.g. `SIMPLE` or `TERMINAL`
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
            __props__['name'] = name
            if routing_strategy is None:
                raise TypeError("Missing required property 'routing_strategy'")
            __props__['routing_strategy'] = routing_strategy
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Alias, __self__).__init__(
            'aws:gamelift/alias:Alias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, name=None, routing_strategy=None, tags=None):
        """
        Get an existing Alias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Alias ARN.
        :param pulumi.Input[str] description: Description of the alias.
        :param pulumi.Input[str] name: Name of the alias.
        :param pulumi.Input[dict] routing_strategy: Specifies the fleet and/or routing type to use for the alias.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **routing_strategy** object supports the following:

          * `fleetId` (`pulumi.Input[str]`) - ID of the Gamelift Fleet to point the alias to.
          * `message` (`pulumi.Input[str]`) - Message text to be used with the `TERMINAL` routing strategy.
          * `type` (`pulumi.Input[str]`) - Type of routing strategy. e.g. `SIMPLE` or `TERMINAL`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["name"] = name
        __props__["routing_strategy"] = routing_strategy
        __props__["tags"] = tags
        return Alias(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


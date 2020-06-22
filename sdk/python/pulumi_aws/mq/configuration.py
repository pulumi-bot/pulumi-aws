# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Configuration(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the configuration.
    """
    data: pulumi.Output[str]
    """
    The broker configuration in XML format.
    See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
    for supported parameters and format of the XML.
    """
    description: pulumi.Output[str]
    """
    The description of the configuration.
    """
    engine_type: pulumi.Output[str]
    """
    The type of broker engine.
    """
    engine_version: pulumi.Output[str]
    """
    The version of the broker engine.
    """
    latest_revision: pulumi.Output[float]
    """
    The latest revision of the configuration.
    """
    name: pulumi.Output[str]
    """
    The name of the configuration
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, data=None, description=None, engine_type=None, engine_version=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an MQ Configuration Resource.

        For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.mq.Configuration("example",
            data=\"\"\"<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
        <broker xmlns="http://activemq.apache.org/schema/core">
          <plugins>
            <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
            <statisticsBrokerPlugin/>
            <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
          </plugins>
        </broker>

        \"\"\",
            description="Example Configuration",
            engine_type="ActiveMQ",
            engine_version="5.15.0")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: The broker configuration in XML format.
               See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
               for supported parameters and format of the XML.
        :param pulumi.Input[str] description: The description of the configuration.
        :param pulumi.Input[str] engine_type: The type of broker engine.
        :param pulumi.Input[str] engine_version: The version of the broker engine.
        :param pulumi.Input[str] name: The name of the configuration
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
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

            if data is None:
                raise TypeError("Missing required property 'data'")
            __props__['data'] = data
            __props__['description'] = description
            if engine_type is None:
                raise TypeError("Missing required property 'engine_type'")
            __props__['engine_type'] = engine_type
            if engine_version is None:
                raise TypeError("Missing required property 'engine_version'")
            __props__['engine_version'] = engine_version
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['latest_revision'] = None
        super(Configuration, __self__).__init__(
            'aws:mq/configuration:Configuration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, data=None, description=None, engine_type=None, engine_version=None, latest_revision=None, name=None, tags=None):
        """
        Get an existing Configuration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the configuration.
        :param pulumi.Input[str] data: The broker configuration in XML format.
               See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
               for supported parameters and format of the XML.
        :param pulumi.Input[str] description: The description of the configuration.
        :param pulumi.Input[str] engine_type: The type of broker engine.
        :param pulumi.Input[str] engine_version: The version of the broker engine.
        :param pulumi.Input[float] latest_revision: The latest revision of the configuration.
        :param pulumi.Input[str] name: The name of the configuration
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["data"] = data
        __props__["description"] = description
        __props__["engine_type"] = engine_type
        __props__["engine_version"] = engine_version
        __props__["latest_revision"] = latest_revision
        __props__["name"] = name
        __props__["tags"] = tags
        return Configuration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

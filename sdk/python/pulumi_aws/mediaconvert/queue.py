# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Queue(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Arn of the queue
    """
    description: pulumi.Output[str]
    """
    A description of the queue
    """
    name: pulumi.Output[str]
    """
    A unique identifier describing the queue
    """
    pricing_plan: pulumi.Output[str]
    """
    Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
    """
    reservation_plan_settings: pulumi.Output[dict]
    """
    A detail pricing plan of the  reserved queue. See below.

      * `commitment` (`str`) - The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
      * `renewalType` (`str`) - Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
      * `reservedSlots` (`float`) - Specifies the number of reserved transcode slots (RTS) for queue.
    """
    status: pulumi.Output[str]
    """
    A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, pricing_plan=None, reservation_plan_settings=None, status=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Elemental MediaConvert Queue.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.mediaconvert.Queue("test")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the queue
        :param pulumi.Input[str] name: A unique identifier describing the queue
        :param pulumi.Input[str] pricing_plan: Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
        :param pulumi.Input[dict] reservation_plan_settings: A detail pricing plan of the  reserved queue. See below.
        :param pulumi.Input[str] status: A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **reservation_plan_settings** object supports the following:

          * `commitment` (`pulumi.Input[str]`) - The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
          * `renewalType` (`pulumi.Input[str]`) - Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
          * `reservedSlots` (`pulumi.Input[float]`) - Specifies the number of reserved transcode slots (RTS) for queue.
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
            __props__['pricing_plan'] = pricing_plan
            __props__['reservation_plan_settings'] = reservation_plan_settings
            __props__['status'] = status
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Queue, __self__).__init__(
            'aws:mediaconvert/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, name=None, pricing_plan=None, reservation_plan_settings=None, status=None, tags=None):
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Arn of the queue
        :param pulumi.Input[str] description: A description of the queue
        :param pulumi.Input[str] name: A unique identifier describing the queue
        :param pulumi.Input[str] pricing_plan: Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
        :param pulumi.Input[dict] reservation_plan_settings: A detail pricing plan of the  reserved queue. See below.
        :param pulumi.Input[str] status: A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **reservation_plan_settings** object supports the following:

          * `commitment` (`pulumi.Input[str]`) - The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
          * `renewalType` (`pulumi.Input[str]`) - Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
          * `reservedSlots` (`pulumi.Input[float]`) - Specifies the number of reserved transcode slots (RTS) for queue.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["name"] = name
        __props__["pricing_plan"] = pricing_plan
        __props__["reservation_plan_settings"] = reservation_plan_settings
        __props__["status"] = status
        __props__["tags"] = tags
        return Queue(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

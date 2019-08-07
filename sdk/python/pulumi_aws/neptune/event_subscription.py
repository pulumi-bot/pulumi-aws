# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventSubscription(pulumi.CustomResource):
    arn: pulumi.Output[str]
    customer_aws_id: pulumi.Output[str]
    enabled: pulumi.Output[bool]
    """
    A boolean flag to enable/disable the subscription. Defaults to true.
    """
    event_categories: pulumi.Output[list]
    """
    A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
    """
    name: pulumi.Output[str]
    """
    The name of the Neptune event subscription. By default generated by this provider.
    """
    name_prefix: pulumi.Output[str]
    """
    The name of the Neptune event subscription. Conflicts with `name`.
    """
    sns_topic_arn: pulumi.Output[str]
    """
    The ARN of the SNS topic to send events to.
    """
    source_ids: pulumi.Output[list]
    """
    A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
    """
    source_type: pulumi.Output[str]
    """
    The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, enabled=None, event_categories=None, name=None, name_prefix=None, sns_topic_arn=None, source_ids=None, source_type=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        ## Attributes
        
        The following additional atttributes are provided:
        
        * `id` - The name of the Neptune event notification subscription.
        * `arn` - The Amazon Resource Name of the Neptune event notification subscription.
        * `customer_aws_id` - The AWS customer account associated with the Neptune event notification subscription.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[list] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[list] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/neptune_event_subscription.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['enabled'] = enabled
            __props__['event_categories'] = event_categories
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            if sns_topic_arn is None:
                raise TypeError("Missing required property 'sns_topic_arn'")
            __props__['sns_topic_arn'] = sns_topic_arn
            __props__['source_ids'] = source_ids
            __props__['source_type'] = source_type
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['customer_aws_id'] = None
        super(EventSubscription, __self__).__init__(
            'aws:neptune/eventSubscription:EventSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, customer_aws_id=None, enabled=None, event_categories=None, name=None, name_prefix=None, sns_topic_arn=None, source_ids=None, source_type=None, tags=None):
        """
        Get an existing EventSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[list] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[list] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/neptune_event_subscription.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["customer_aws_id"] = customer_aws_id
        __props__["enabled"] = enabled
        __props__["event_categories"] = event_categories
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["sns_topic_arn"] = sns_topic_arn
        __props__["source_ids"] = source_ids
        __props__["source_type"] = source_type
        __props__["tags"] = tags
        return EventSubscription(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Environment(pulumi.CustomResource):
    all_settings: pulumi.Output[list]
    """
    List of all option settings configured in this Environment. These
    are a combination of default settings and their overrides from `setting` in
    the configuration.

      * `name` (`str`) - A unique name for this Environment. This name is used
        in the application URL
      * `namespace` (`str`)
      * `resource` (`str`)
      * `value` (`str`)
    """
    application: pulumi.Output[str]
    """
    Name of the application that contains the version
    to be deployed
    """
    arn: pulumi.Output[str]
    autoscaling_groups: pulumi.Output[list]
    """
    The autoscaling groups used by this Environment.
    """
    cname: pulumi.Output[str]
    """
    Fully qualified DNS name for this Environment.
    """
    cname_prefix: pulumi.Output[str]
    """
    Prefix to use for the fully qualified DNS name of
    the Environment.
    """
    description: pulumi.Output[str]
    """
    Short description of the Environment
    """
    endpoint_url: pulumi.Output[str]
    """
    The URL to the Load Balancer for this Environment
    """
    instances: pulumi.Output[list]
    """
    Instances used by this Environment.
    """
    launch_configurations: pulumi.Output[list]
    """
    Launch configurations in use by this Environment.
    """
    load_balancers: pulumi.Output[list]
    """
    Elastic load balancers in use by this Environment.
    """
    name: pulumi.Output[str]
    """
    A unique name for this Environment. This name is used
    in the application URL
    """
    platform_arn: pulumi.Output[str]
    """
    The [ARN][2] of the Elastic Beanstalk [Platform][3]
    to use in deployment
    """
    poll_interval: pulumi.Output[str]
    """
    The time between polling the AWS API to
    check if changes have been applied. Use this to adjust the rate of API calls
    for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
    use the default behavior, which is an exponential backoff
    """
    queues: pulumi.Output[list]
    """
    SQS queues in use by this Environment.
    """
    settings: pulumi.Output[list]
    """
    Option settings to configure the new Environment. These
    override specific values that are set as defaults. The format is detailed
    below in Option Settings

      * `name` (`str`) - A unique name for this Environment. This name is used
        in the application URL
      * `namespace` (`str`)
      * `resource` (`str`)
      * `value` (`str`)
    """
    solution_stack_name: pulumi.Output[str]
    """
    A solution stack to base your environment
    off of. Example stacks can be found in the [Amazon API documentation][1]
    """
    tags: pulumi.Output[dict]
    """
    A set of tags to apply to the Environment.
    """
    template_name: pulumi.Output[str]
    """
    The name of the Elastic Beanstalk Configuration
    template to use in deployment
    """
    tier: pulumi.Output[str]
    """
    Elastic Beanstalk Environment tier. Valid values are `Worker`
    or `WebServer`. If tier is left blank `WebServer` will be used.
    """
    triggers: pulumi.Output[list]
    """
    Autoscaling triggers in use by this Environment.
    """
    version: pulumi.Output[str]
    """
    The name of the Elastic Beanstalk Application Version
    to use in deployment.
    """
    wait_for_ready_timeout: pulumi.Output[str]
    """
    The maximum
    [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
    wait for an Elastic Beanstalk Environment to be in a ready state before timing
    out.
    """
    def __init__(__self__, resource_name, opts=None, application=None, cname_prefix=None, description=None, name=None, platform_arn=None, poll_interval=None, settings=None, solution_stack_name=None, tags=None, template_name=None, tier=None, version=None, wait_for_ready_timeout=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Elastic Beanstalk Environment Resource. Elastic Beanstalk allows
        you to deploy and manage applications in the AWS cloud without worrying about
        the infrastructure that runs those applications.

        Environments are often things such as `development`, `integration`, or
        `production`.

        ## Option Settings

        {{% examples %}}

        Some options can be stack-specific, check [AWS Docs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html)
        for supported options and examples.

        The `setting` and `all_settings` mappings support the following format:

        * `namespace` - unique namespace identifying the option's associated AWS resource
        * `name` - name of the configuration option
        * `value` - value for the configuration option
        * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)

        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastic_beanstalk_environment.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] application: Name of the application that contains the version
               to be deployed
        :param pulumi.Input[str] cname_prefix: Prefix to use for the fully qualified DNS name of
               the Environment.
        :param pulumi.Input[str] description: Short description of the Environment
        :param pulumi.Input[str] name: A unique name for this Environment. This name is used
               in the application URL
        :param pulumi.Input[str] platform_arn: The [ARN][2] of the Elastic Beanstalk [Platform][3]
               to use in deployment
        :param pulumi.Input[str] poll_interval: The time between polling the AWS API to
               check if changes have been applied. Use this to adjust the rate of API calls
               for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
               use the default behavior, which is an exponential backoff
        :param pulumi.Input[list] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your environment
               off of. Example stacks can be found in the [Amazon API documentation][1]
        :param pulumi.Input[dict] tags: A set of tags to apply to the Environment.
        :param pulumi.Input[str] template_name: The name of the Elastic Beanstalk Configuration
               template to use in deployment
        :param pulumi.Input[str] tier: Elastic Beanstalk Environment tier. Valid values are `Worker`
               or `WebServer`. If tier is left blank `WebServer` will be used.
        :param pulumi.Input[str] version: The name of the Elastic Beanstalk Application Version
               to use in deployment.
        :param pulumi.Input[str] wait_for_ready_timeout: The maximum
               [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
               wait for an Elastic Beanstalk Environment to be in a ready state before timing
               out.

        The **settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - A unique name for this Environment. This name is used
            in the application URL
          * `namespace` (`pulumi.Input[str]`)
          * `resource` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
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

            if application is None:
                raise TypeError("Missing required property 'application'")
            __props__['application'] = application
            __props__['cname_prefix'] = cname_prefix
            __props__['description'] = description
            __props__['name'] = name
            __props__['platform_arn'] = platform_arn
            __props__['poll_interval'] = poll_interval
            __props__['settings'] = settings
            __props__['solution_stack_name'] = solution_stack_name
            __props__['tags'] = tags
            __props__['template_name'] = template_name
            __props__['tier'] = tier
            __props__['version'] = version
            __props__['wait_for_ready_timeout'] = wait_for_ready_timeout
            __props__['all_settings'] = None
            __props__['arn'] = None
            __props__['autoscaling_groups'] = None
            __props__['cname'] = None
            __props__['endpoint_url'] = None
            __props__['instances'] = None
            __props__['launch_configurations'] = None
            __props__['load_balancers'] = None
            __props__['queues'] = None
            __props__['triggers'] = None
        super(Environment, __self__).__init__(
            'aws:elasticbeanstalk/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, all_settings=None, application=None, arn=None, autoscaling_groups=None, cname=None, cname_prefix=None, description=None, endpoint_url=None, instances=None, launch_configurations=None, load_balancers=None, name=None, platform_arn=None, poll_interval=None, queues=None, settings=None, solution_stack_name=None, tags=None, template_name=None, tier=None, triggers=None, version=None, wait_for_ready_timeout=None):
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] all_settings: List of all option settings configured in this Environment. These
               are a combination of default settings and their overrides from `setting` in
               the configuration.
        :param pulumi.Input[dict] application: Name of the application that contains the version
               to be deployed
        :param pulumi.Input[list] autoscaling_groups: The autoscaling groups used by this Environment.
        :param pulumi.Input[str] cname: Fully qualified DNS name for this Environment.
        :param pulumi.Input[str] cname_prefix: Prefix to use for the fully qualified DNS name of
               the Environment.
        :param pulumi.Input[str] description: Short description of the Environment
        :param pulumi.Input[str] endpoint_url: The URL to the Load Balancer for this Environment
        :param pulumi.Input[list] instances: Instances used by this Environment.
        :param pulumi.Input[list] launch_configurations: Launch configurations in use by this Environment.
        :param pulumi.Input[list] load_balancers: Elastic load balancers in use by this Environment.
        :param pulumi.Input[str] name: A unique name for this Environment. This name is used
               in the application URL
        :param pulumi.Input[str] platform_arn: The [ARN][2] of the Elastic Beanstalk [Platform][3]
               to use in deployment
        :param pulumi.Input[str] poll_interval: The time between polling the AWS API to
               check if changes have been applied. Use this to adjust the rate of API calls
               for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
               use the default behavior, which is an exponential backoff
        :param pulumi.Input[list] queues: SQS queues in use by this Environment.
        :param pulumi.Input[list] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your environment
               off of. Example stacks can be found in the [Amazon API documentation][1]
        :param pulumi.Input[dict] tags: A set of tags to apply to the Environment.
        :param pulumi.Input[str] template_name: The name of the Elastic Beanstalk Configuration
               template to use in deployment
        :param pulumi.Input[str] tier: Elastic Beanstalk Environment tier. Valid values are `Worker`
               or `WebServer`. If tier is left blank `WebServer` will be used.
        :param pulumi.Input[list] triggers: Autoscaling triggers in use by this Environment.
        :param pulumi.Input[str] version: The name of the Elastic Beanstalk Application Version
               to use in deployment.
        :param pulumi.Input[str] wait_for_ready_timeout: The maximum
               [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
               wait for an Elastic Beanstalk Environment to be in a ready state before timing
               out.

        The **all_settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - A unique name for this Environment. This name is used
            in the application URL
          * `namespace` (`pulumi.Input[str]`)
          * `resource` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - A unique name for this Environment. This name is used
            in the application URL
          * `namespace` (`pulumi.Input[str]`)
          * `resource` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["all_settings"] = all_settings
        __props__["application"] = application
        __props__["arn"] = arn
        __props__["autoscaling_groups"] = autoscaling_groups
        __props__["cname"] = cname
        __props__["cname_prefix"] = cname_prefix
        __props__["description"] = description
        __props__["endpoint_url"] = endpoint_url
        __props__["instances"] = instances
        __props__["launch_configurations"] = launch_configurations
        __props__["load_balancers"] = load_balancers
        __props__["name"] = name
        __props__["platform_arn"] = platform_arn
        __props__["poll_interval"] = poll_interval
        __props__["queues"] = queues
        __props__["settings"] = settings
        __props__["solution_stack_name"] = solution_stack_name
        __props__["tags"] = tags
        __props__["template_name"] = template_name
        __props__["tier"] = tier
        __props__["triggers"] = triggers
        __props__["version"] = version
        __props__["wait_for_ready_timeout"] = wait_for_ready_timeout
        return Environment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


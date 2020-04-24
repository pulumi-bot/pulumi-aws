# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ConfigurationTemplate(pulumi.CustomResource):
    application: pulumi.Output[str]
    """
    name of the application to associate with this configuration template
    """
    description: pulumi.Output[str]
    """
    Short description of the Template
    """
    environment_id: pulumi.Output[str]
    """
    The ID of the environment used with this configuration template
    """
    name: pulumi.Output[str]
    """
    A unique name for this Template.
    """
    settings: pulumi.Output[list]
    """
    Option settings to configure the new Environment. These
    override specific values that are set as defaults. The format is detailed
    below in Option Settings

      * `name` (`str`) - A unique name for this Template.
      * `namespace` (`str`)
      * `resource` (`str`)
      * `value` (`str`)
    """
    solution_stack_name: pulumi.Output[str]
    """
    A solution stack to base your Template
    off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
    """
    def __init__(__self__, resource_name, opts=None, application=None, description=None, environment_id=None, name=None, settings=None, solution_stack_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Elastic Beanstalk Configuration Template, which are associated with
        a specific application and are used to deploy different versions of the
        application with the same configuration settings.


        ## Option Settings

        The `setting` field supports the following format:

        * `namespace` - unique namespace identifying the option's associated AWS resource
        * `name` - name of the configuration option
        * `value` - value for the configuration option
        * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application: name of the application to associate with this configuration template
        :param pulumi.Input[str] description: Short description of the Template
        :param pulumi.Input[str] environment_id: The ID of the environment used with this configuration template
        :param pulumi.Input[str] name: A unique name for this Template.
        :param pulumi.Input[list] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your Template
               off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)

        The **settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - A unique name for this Template.
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
            __props__['description'] = description
            __props__['environment_id'] = environment_id
            __props__['name'] = name
            __props__['settings'] = settings
            __props__['solution_stack_name'] = solution_stack_name
        super(ConfigurationTemplate, __self__).__init__(
            'aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, application=None, description=None, environment_id=None, name=None, settings=None, solution_stack_name=None):
        """
        Get an existing ConfigurationTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application: name of the application to associate with this configuration template
        :param pulumi.Input[str] description: Short description of the Template
        :param pulumi.Input[str] environment_id: The ID of the environment used with this configuration template
        :param pulumi.Input[str] name: A unique name for this Template.
        :param pulumi.Input[list] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your Template
               off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)

        The **settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - A unique name for this Template.
          * `namespace` (`pulumi.Input[str]`)
          * `resource` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application"] = application
        __props__["description"] = description
        __props__["environment_id"] = environment_id
        __props__["name"] = name
        __props__["settings"] = settings
        __props__["solution_stack_name"] = solution_stack_name
        return ConfigurationTemplate(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class NotebookInstanceLifecycleConfiguration(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
    """
    name: pulumi.Output[str]
    """
    The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
    """
    on_create: pulumi.Output[str]
    """
    A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
    """
    on_start: pulumi.Output[str]
    """
    A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
    """
    def __init__(__self__, resource_name, opts=None, name=None, on_create=None, on_start=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a lifecycle configuration for SageMaker Notebook Instances.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_notebook_instance_lifecycle_configuration.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[str] on_create: A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
        :param pulumi.Input[str] on_start: A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
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

            __props__['name'] = name
            __props__['on_create'] = on_create
            __props__['on_start'] = on_start
            __props__['arn'] = None
        super(NotebookInstanceLifecycleConfiguration, __self__).__init__(
            'aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, name=None, on_create=None, on_start=None):
        """
        Get an existing NotebookInstanceLifecycleConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
        :param pulumi.Input[str] name: The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[str] on_create: A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
        :param pulumi.Input[str] on_start: A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["on_create"] = on_create
        __props__["on_start"] = on_start
        return NotebookInstanceLifecycleConfiguration(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


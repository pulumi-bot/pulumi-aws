# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class JobDefinition(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name of the job definition.
    """
    container_properties: pulumi.Output[str]
    """
    A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
    provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the job definition.
    """
    parameters: pulumi.Output[dict]
    """
    Specifies the parameter substitution placeholders to set in the job definition.
    """
    retry_strategy: pulumi.Output[dict]
    """
    Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
    Maximum number of `retry_strategy` is `1`.  Defined below.

      * `attempts` (`float`) - The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
    """
    revision: pulumi.Output[float]
    """
    The revision of the job definition.
    """
    timeout: pulumi.Output[dict]
    """
    Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.

      * `attemptDurationSeconds` (`float`) - The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
    """
    type: pulumi.Output[str]
    """
    The type of job definition.  Must be `container`
    """
    def __init__(__self__, resource_name, opts=None, container_properties=None, name=None, parameters=None, retry_strategy=None, timeout=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Batch Job Definition resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.batch.JobDefinition("test",
            container_properties=\"\"\"{
        	"command": ["ls", "-la"],
        	"image": "busybox",
        	"memory": 1024,
        	"vcpus": 1,
        	"volumes": [
              {
                "host": {
                  "sourcePath": "/tmp"
                },
                "name": "tmp"
              }
            ],
        	"environment": [
        		{"name": "VARNAME", "value": "VARVAL"}
        	],
        	"mountPoints": [
        		{
                  "sourceVolume": "tmp",
                  "containerPath": "/tmp",
                  "readOnly": false
                }
        	],
            "ulimits": [
              {
                "hardLimit": 1024,
                "name": "nofile",
                "softLimit": 1024
              }
            ]
        }

        \"\"\",
            type="container")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_properties: A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
               provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
        :param pulumi.Input[str] name: Specifies the name of the job definition.
        :param pulumi.Input[dict] parameters: Specifies the parameter substitution placeholders to set in the job definition.
        :param pulumi.Input[dict] retry_strategy: Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
               Maximum number of `retry_strategy` is `1`.  Defined below.
        :param pulumi.Input[dict] timeout: Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        :param pulumi.Input[str] type: The type of job definition.  Must be `container`

        The **retry_strategy** object supports the following:

          * `attempts` (`pulumi.Input[float]`) - The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.

        The **timeout** object supports the following:

          * `attemptDurationSeconds` (`pulumi.Input[float]`) - The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
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

            __props__['container_properties'] = container_properties
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['retry_strategy'] = retry_strategy
            __props__['timeout'] = timeout
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['arn'] = None
            __props__['revision'] = None
        super(JobDefinition, __self__).__init__(
            'aws:batch/jobDefinition:JobDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, container_properties=None, name=None, parameters=None, retry_strategy=None, revision=None, timeout=None, type=None):
        """
        Get an existing JobDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name of the job definition.
        :param pulumi.Input[str] container_properties: A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
               provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
        :param pulumi.Input[str] name: Specifies the name of the job definition.
        :param pulumi.Input[dict] parameters: Specifies the parameter substitution placeholders to set in the job definition.
        :param pulumi.Input[dict] retry_strategy: Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
               Maximum number of `retry_strategy` is `1`.  Defined below.
        :param pulumi.Input[float] revision: The revision of the job definition.
        :param pulumi.Input[dict] timeout: Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        :param pulumi.Input[str] type: The type of job definition.  Must be `container`

        The **retry_strategy** object supports the following:

          * `attempts` (`pulumi.Input[float]`) - The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.

        The **timeout** object supports the following:

          * `attemptDurationSeconds` (`pulumi.Input[float]`) - The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["container_properties"] = container_properties
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["retry_strategy"] = retry_strategy
        __props__["revision"] = revision
        __props__["timeout"] = timeout
        __props__["type"] = type
        return JobDefinition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

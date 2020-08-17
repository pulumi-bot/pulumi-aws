# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['JobDefinition']


class JobDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_properties: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 retry_strategy: Optional[pulumi.Input[pulumi.InputType['JobDefinitionRetryStrategyArgs']]] = None,
                 timeout: Optional[pulumi.Input[pulumi.InputType['JobDefinitionTimeoutArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Specifies the parameter substitution placeholders to set in the job definition.
        :param pulumi.Input[pulumi.InputType['JobDefinitionRetryStrategyArgs']] retry_strategy: Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
               Maximum number of `retry_strategy` is `1`.  Defined below.
        :param pulumi.Input[pulumi.InputType['JobDefinitionTimeoutArgs']] timeout: Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        :param pulumi.Input[str] type: The type of job definition.  Must be `container`
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
            opts.version = _utilities.get_version()
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
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            container_properties: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            retry_strategy: Optional[pulumi.Input[pulumi.InputType['JobDefinitionRetryStrategyArgs']]] = None,
            revision: Optional[pulumi.Input[float]] = None,
            timeout: Optional[pulumi.Input[pulumi.InputType['JobDefinitionTimeoutArgs']]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'JobDefinition':
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
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Specifies the parameter substitution placeholders to set in the job definition.
        :param pulumi.Input[pulumi.InputType['JobDefinitionRetryStrategyArgs']] retry_strategy: Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
               Maximum number of `retry_strategy` is `1`.  Defined below.
        :param pulumi.Input[float] revision: The revision of the job definition.
        :param pulumi.Input[pulumi.InputType['JobDefinitionTimeoutArgs']] timeout: Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        :param pulumi.Input[str] type: The type of job definition.  Must be `container`
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

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name of the job definition.
        """
        ...

    @property
    @pulumi.getter(name="containerProperties")
    def container_properties(self) -> Optional[str]:
        """
        A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
        provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the job definition.
        """
        ...

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, str]]:
        """
        Specifies the parameter substitution placeholders to set in the job definition.
        """
        ...

    @property
    @pulumi.getter(name="retryStrategy")
    def retry_strategy(self) -> Optional['outputs.JobDefinitionRetryStrategy']:
        """
        Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
        Maximum number of `retry_strategy` is `1`.  Defined below.
        """
        ...

    @property
    @pulumi.getter
    def revision(self) -> float:
        """
        The revision of the job definition.
        """
        ...

    @property
    @pulumi.getter
    def timeout(self) -> Optional['outputs.JobDefinitionTimeout']:
        """
        Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of job definition.  Must be `container`
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


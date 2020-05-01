# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Job(pulumi.CustomResource):
    allocated_capacity: pulumi.Output[float]
    """
    **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of Glue Job
    """
    command: pulumi.Output[dict]
    """
    The command of the job. Defined below.

      * `name` (`str`) - The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `max_capacity` needs to be set if `pythonshell` is chosen.
      * `pythonVersion` (`str`) - The Python version being used to execute a Python shell job. Allowed values are 2 or 3.
      * `scriptLocation` (`str`) - Specifies the S3 path to a script that executes a job.
    """
    connections: pulumi.Output[list]
    """
    The list of connections used for this job.
    """
    default_arguments: pulumi.Output[dict]
    """
    The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
    """
    description: pulumi.Output[str]
    """
    Description of the job.
    """
    execution_property: pulumi.Output[dict]
    """
    Execution property of the job. Defined below.

      * `maxConcurrentRuns` (`float`) - The maximum number of concurrent runs allowed for a job. The default is 1.
    """
    glue_version: pulumi.Output[str]
    """
    The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
    """
    max_capacity: pulumi.Output[float]
    """
    The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`.
    """
    max_retries: pulumi.Output[float]
    """
    The maximum number of times to retry this job if it fails.
    """
    name: pulumi.Output[str]
    """
    The name you assign to this job. It must be unique in your account.
    """
    notification_property: pulumi.Output[dict]
    """
    Notification property of the job. Defined below.

      * `notifyDelayAfter` (`float`) - After a job run starts, the number of minutes to wait before sending a job run delay notification.
    """
    number_of_workers: pulumi.Output[float]
    """
    The number of workers of a defined workerType that are allocated when a job runs.
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of the IAM role associated with this job.
    """
    security_configuration: pulumi.Output[str]
    """
    The name of the Security Configuration to be associated with the job.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    timeout: pulumi.Output[float]
    """
    The job timeout in minutes. The default is 2880 minutes (48 hours).
    """
    worker_type: pulumi.Output[str]
    """
    The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
    """
    def __init__(__self__, resource_name, opts=None, allocated_capacity=None, command=None, connections=None, default_arguments=None, description=None, execution_property=None, glue_version=None, max_capacity=None, max_retries=None, name=None, notification_property=None, number_of_workers=None, role_arn=None, security_configuration=None, tags=None, timeout=None, worker_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Glue Job resource.

        > Glue functionality, such as monitoring and logging of jobs, is typically managed with the `default_arguments` argument. See the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the Glue developer guide for additional information.

        ## Example Usage

        ### Python Job

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Job("example",
            command={
                "scriptLocation": f"s3://{aws_s3_bucket['example']['bucket']}/example.py",
            },
            role_arn=aws_iam_role["example"]["arn"])
        ```

        ### Scala Job

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Job("example",
            command={
                "scriptLocation": f"s3://{aws_s3_bucket['example']['bucket']}/example.scala",
            },
            default_arguments={
                "--job-language": "scala",
            },
            role_arn=aws_iam_role["example"]["arn"])
        ```

        ### Enabling CloudWatch Logs and Metrics

        ```python
        import pulumi
        import pulumi_aws as aws

        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=14)
        example_job = aws.glue.Job("exampleJob", default_arguments={
            "--continuous-log-logGroup": example_log_group.name,
            "--enable-continuous-cloudwatch-log": "true",
            "--enable-continuous-log-filter": "true",
            "--enable-metrics": "",
        })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] allocated_capacity: **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
        :param pulumi.Input[dict] command: The command of the job. Defined below.
        :param pulumi.Input[list] connections: The list of connections used for this job.
        :param pulumi.Input[dict] default_arguments: The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
        :param pulumi.Input[str] description: Description of the job.
        :param pulumi.Input[dict] execution_property: Execution property of the job. Defined below.
        :param pulumi.Input[str] glue_version: The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        :param pulumi.Input[float] max_capacity: The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`.
        :param pulumi.Input[float] max_retries: The maximum number of times to retry this job if it fails.
        :param pulumi.Input[str] name: The name you assign to this job. It must be unique in your account.
        :param pulumi.Input[dict] notification_property: Notification property of the job. Defined below.
        :param pulumi.Input[float] number_of_workers: The number of workers of a defined workerType that are allocated when a job runs.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role associated with this job.
        :param pulumi.Input[str] security_configuration: The name of the Security Configuration to be associated with the job.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        :param pulumi.Input[float] timeout: The job timeout in minutes. The default is 2880 minutes (48 hours).
        :param pulumi.Input[str] worker_type: The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.

        The **command** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `max_capacity` needs to be set if `pythonshell` is chosen.
          * `pythonVersion` (`pulumi.Input[str]`) - The Python version being used to execute a Python shell job. Allowed values are 2 or 3.
          * `scriptLocation` (`pulumi.Input[str]`) - Specifies the S3 path to a script that executes a job.

        The **execution_property** object supports the following:

          * `maxConcurrentRuns` (`pulumi.Input[float]`) - The maximum number of concurrent runs allowed for a job. The default is 1.

        The **notification_property** object supports the following:

          * `notifyDelayAfter` (`pulumi.Input[float]`) - After a job run starts, the number of minutes to wait before sending a job run delay notification.
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

            __props__['allocated_capacity'] = allocated_capacity
            if command is None:
                raise TypeError("Missing required property 'command'")
            __props__['command'] = command
            __props__['connections'] = connections
            __props__['default_arguments'] = default_arguments
            __props__['description'] = description
            __props__['execution_property'] = execution_property
            __props__['glue_version'] = glue_version
            __props__['max_capacity'] = max_capacity
            __props__['max_retries'] = max_retries
            __props__['name'] = name
            __props__['notification_property'] = notification_property
            __props__['number_of_workers'] = number_of_workers
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['security_configuration'] = security_configuration
            __props__['tags'] = tags
            __props__['timeout'] = timeout
            __props__['worker_type'] = worker_type
            __props__['arn'] = None
        super(Job, __self__).__init__(
            'aws:glue/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allocated_capacity=None, arn=None, command=None, connections=None, default_arguments=None, description=None, execution_property=None, glue_version=None, max_capacity=None, max_retries=None, name=None, notification_property=None, number_of_workers=None, role_arn=None, security_configuration=None, tags=None, timeout=None, worker_type=None):
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] allocated_capacity: **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of Glue Job
        :param pulumi.Input[dict] command: The command of the job. Defined below.
        :param pulumi.Input[list] connections: The list of connections used for this job.
        :param pulumi.Input[dict] default_arguments: The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
        :param pulumi.Input[str] description: Description of the job.
        :param pulumi.Input[dict] execution_property: Execution property of the job. Defined below.
        :param pulumi.Input[str] glue_version: The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        :param pulumi.Input[float] max_capacity: The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`.
        :param pulumi.Input[float] max_retries: The maximum number of times to retry this job if it fails.
        :param pulumi.Input[str] name: The name you assign to this job. It must be unique in your account.
        :param pulumi.Input[dict] notification_property: Notification property of the job. Defined below.
        :param pulumi.Input[float] number_of_workers: The number of workers of a defined workerType that are allocated when a job runs.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role associated with this job.
        :param pulumi.Input[str] security_configuration: The name of the Security Configuration to be associated with the job.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        :param pulumi.Input[float] timeout: The job timeout in minutes. The default is 2880 minutes (48 hours).
        :param pulumi.Input[str] worker_type: The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.

        The **command** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `max_capacity` needs to be set if `pythonshell` is chosen.
          * `pythonVersion` (`pulumi.Input[str]`) - The Python version being used to execute a Python shell job. Allowed values are 2 or 3.
          * `scriptLocation` (`pulumi.Input[str]`) - Specifies the S3 path to a script that executes a job.

        The **execution_property** object supports the following:

          * `maxConcurrentRuns` (`pulumi.Input[float]`) - The maximum number of concurrent runs allowed for a job. The default is 1.

        The **notification_property** object supports the following:

          * `notifyDelayAfter` (`pulumi.Input[float]`) - After a job run starts, the number of minutes to wait before sending a job run delay notification.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocated_capacity"] = allocated_capacity
        __props__["arn"] = arn
        __props__["command"] = command
        __props__["connections"] = connections
        __props__["default_arguments"] = default_arguments
        __props__["description"] = description
        __props__["execution_property"] = execution_property
        __props__["glue_version"] = glue_version
        __props__["max_capacity"] = max_capacity
        __props__["max_retries"] = max_retries
        __props__["name"] = name
        __props__["notification_property"] = notification_property
        __props__["number_of_workers"] = number_of_workers
        __props__["role_arn"] = role_arn
        __props__["security_configuration"] = security_configuration
        __props__["tags"] = tags
        __props__["timeout"] = timeout
        __props__["worker_type"] = worker_type
        return Job(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


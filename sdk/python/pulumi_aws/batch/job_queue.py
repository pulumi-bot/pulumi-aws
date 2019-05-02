# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class JobQueue(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name of the job queue.
    """
    compute_environments: pulumi.Output[list]
    """
    Specifies the set of compute environments
    mapped to a job queue and their order.  The position of the compute environments
    in the list will dictate the order. You can associate up to 3 compute environments
    with a job queue.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the job queue.
    """
    priority: pulumi.Output[float]
    """
    The priority of the job queue. Job queues with a higher priority
    are evaluated first when associated with the same compute environment.
    """
    state: pulumi.Output[str]
    """
    The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
    """
    def __init__(__self__, resource_name, opts=None, compute_environments=None, name=None, priority=None, state=None, __name__=None, __opts__=None):
        """
        Provides a Batch Job Queue resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] compute_environments: Specifies the set of compute environments
               mapped to a job queue and their order.  The position of the compute environments
               in the list will dictate the order. You can associate up to 3 compute environments
               with a job queue.
        :param pulumi.Input[str] name: Specifies the name of the job queue.
        :param pulumi.Input[float] priority: The priority of the job queue. Job queues with a higher priority
               are evaluated first when associated with the same compute environment.
        :param pulumi.Input[str] state: The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if compute_environments is None:
            raise TypeError("Missing required property 'compute_environments'")
        __props__['compute_environments'] = compute_environments

        __props__['name'] = name

        if priority is None:
            raise TypeError("Missing required property 'priority'")
        __props__['priority'] = priority

        if state is None:
            raise TypeError("Missing required property 'state'")
        __props__['state'] = state

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(JobQueue, __self__).__init__(
            'aws:batch/jobQueue:JobQueue',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


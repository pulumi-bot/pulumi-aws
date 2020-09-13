# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ComputeEnvironment']


class ComputeEnvironment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_environment_name: Optional[pulumi.Input[str]] = None,
                 compute_environment_name_prefix: Optional[pulumi.Input[str]] = None,
                 compute_resources: Optional[pulumi.Input[pulumi.InputType['ComputeEnvironmentComputeResourcesArgs']]] = None,
                 service_role: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a AWS Batch compute environment. Compute environments contain the Amazon ECS container instances that are used to run containerized batch jobs.

        For information about AWS Batch, see [What is AWS Batch?](http://docs.aws.amazon.com/batch/latest/userguide/what-is-batch.html) .
        For information about compute environment, see [Compute Environments](http://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) .

        > **Note:** To prevent a race condition during environment deletion, make sure to set `depends_on` to the related `iam.RolePolicyAttachment`;
        otherwise, the policy may be destroyed too soon and the compute environment will then get stuck in the `DELETING` state, see [Troubleshooting AWS Batch](http://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html) .

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        ecs_instance_role_role = aws.iam.Role("ecsInstanceRoleRole", assume_role_policy=\"\"\"{
            "Version": "2012-10-17",
            "Statement": [
        	{
        	    "Action": "sts:AssumeRole",
        	    "Effect": "Allow",
        	    "Principal": {
        		"Service": "ec2.amazonaws.com"
        	    }
        	}
            ]
        }
        \"\"\")
        ecs_instance_role_role_policy_attachment = aws.iam.RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment",
            role=ecs_instance_role_role.name,
            policy_arn="arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role")
        ecs_instance_role_instance_profile = aws.iam.InstanceProfile("ecsInstanceRoleInstanceProfile", role=ecs_instance_role_role.name)
        aws_batch_service_role_role = aws.iam.Role("awsBatchServiceRoleRole", assume_role_policy=\"\"\"{
            "Version": "2012-10-17",
            "Statement": [
        	{
        	    "Action": "sts:AssumeRole",
        	    "Effect": "Allow",
        	    "Principal": {
        		"Service": "batch.amazonaws.com"
        	    }
        	}
            ]
        }
        \"\"\")
        aws_batch_service_role_role_policy_attachment = aws.iam.RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment",
            role=aws_batch_service_role_role.name,
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole")
        sample_security_group = aws.ec2.SecurityGroup("sampleSecurityGroup", egress=[aws.ec2.SecurityGroupEgressArgs(
            from_port=0,
            to_port=0,
            protocol="-1",
            cidr_blocks=["0.0.0.0/0"],
        )])
        sample_vpc = aws.ec2.Vpc("sampleVpc", cidr_block="10.1.0.0/16")
        sample_subnet = aws.ec2.Subnet("sampleSubnet",
            vpc_id=sample_vpc.id,
            cidr_block="10.1.1.0/24")
        sample_compute_environment = aws.batch.ComputeEnvironment("sampleComputeEnvironment",
            compute_environment_name="sample",
            compute_resources=aws.batch.ComputeEnvironmentComputeResourcesArgs(
                instance_role=ecs_instance_role_instance_profile.arn,
                instance_types=["c4.large"],
                max_vcpus=16,
                min_vcpus=0,
                security_group_ids=[sample_security_group.id],
                subnets=[sample_subnet.id],
                type="EC2",
            ),
            service_role=aws_batch_service_role_role.arn,
            type="MANAGED",
            opts=ResourceOptions(depends_on=[aws_batch_service_role_role_policy_attachment]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_environment_name: The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[str] compute_environment_name_prefix: Creates a unique compute environment name beginning with the specified prefix. Conflicts with `compute_environment_name`.
        :param pulumi.Input[pulumi.InputType['ComputeEnvironmentComputeResourcesArgs']] compute_resources: Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
        :param pulumi.Input[str] service_role: The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
        :param pulumi.Input[str] state: The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
        :param pulumi.Input[str] type: The type of compute environment. Valid items are `EC2` or `SPOT`.
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

            __props__['compute_environment_name'] = compute_environment_name
            __props__['compute_environment_name_prefix'] = compute_environment_name_prefix
            __props__['compute_resources'] = compute_resources
            if service_role is None:
                raise TypeError("Missing required property 'service_role'")
            __props__['service_role'] = service_role
            __props__['state'] = state
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['arn'] = None
            __props__['ecs_cluster_arn'] = None
            __props__['status'] = None
            __props__['status_reason'] = None
        super(ComputeEnvironment, __self__).__init__(
            'aws:batch/computeEnvironment:ComputeEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            compute_environment_name: Optional[pulumi.Input[str]] = None,
            compute_environment_name_prefix: Optional[pulumi.Input[str]] = None,
            compute_resources: Optional[pulumi.Input[pulumi.InputType['ComputeEnvironmentComputeResourcesArgs']]] = None,
            ecs_cluster_arn: Optional[pulumi.Input[str]] = None,
            service_role: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            status_reason: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ComputeEnvironment':
        """
        Get an existing ComputeEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the compute environment.
        :param pulumi.Input[str] compute_environment_name: The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[str] compute_environment_name_prefix: Creates a unique compute environment name beginning with the specified prefix. Conflicts with `compute_environment_name`.
        :param pulumi.Input[pulumi.InputType['ComputeEnvironmentComputeResourcesArgs']] compute_resources: Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
        :param pulumi.Input[str] ecs_cluster_arn: The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
        :param pulumi.Input[str] service_role: The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
        :param pulumi.Input[str] state: The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
        :param pulumi.Input[str] status: The current status of the compute environment (for example, CREATING or VALID).
        :param pulumi.Input[str] status_reason: A short, human-readable string to provide additional details about the current status of the compute environment.
        :param pulumi.Input[str] type: The type of compute environment. Valid items are `EC2` or `SPOT`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["compute_environment_name"] = compute_environment_name
        __props__["compute_environment_name_prefix"] = compute_environment_name_prefix
        __props__["compute_resources"] = compute_resources
        __props__["ecs_cluster_arn"] = ecs_cluster_arn
        __props__["service_role"] = service_role
        __props__["state"] = state
        __props__["status"] = status
        __props__["status_reason"] = status_reason
        __props__["type"] = type
        return ComputeEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the compute environment.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="computeEnvironmentName")
    def compute_environment_name(self) -> pulumi.Output[str]:
        """
        The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, this provider will assign a random, unique name.
        """
        return pulumi.get(self, "compute_environment_name")

    @property
    @pulumi.getter(name="computeEnvironmentNamePrefix")
    def compute_environment_name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Creates a unique compute environment name beginning with the specified prefix. Conflicts with `compute_environment_name`.
        """
        return pulumi.get(self, "compute_environment_name_prefix")

    @property
    @pulumi.getter(name="computeResources")
    def compute_resources(self) -> pulumi.Output[Optional['outputs.ComputeEnvironmentComputeResources']]:
        """
        Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
        """
        return pulumi.get(self, "compute_resources")

    @property
    @pulumi.getter(name="ecsClusterArn")
    def ecs_cluster_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
        """
        return pulumi.get(self, "ecs_cluster_arn")

    @property
    @pulumi.getter(name="serviceRole")
    def service_role(self) -> pulumi.Output[str]:
        """
        The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
        """
        return pulumi.get(self, "service_role")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current status of the compute environment (for example, CREATING or VALID).
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> pulumi.Output[str]:
        """
        A short, human-readable string to provide additional details about the current status of the compute environment.
        """
        return pulumi.get(self, "status_reason")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of compute environment. Valid items are `EC2` or `SPOT`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


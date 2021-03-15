# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['DevEndpointArgs', 'DevEndpoint']

@pulumi.input_type
class DevEndpointArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 arguments: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 extra_jars_s3_path: Optional[pulumi.Input[str]] = None,
                 extra_python_libs_s3_path: Optional[pulumi.Input[str]] = None,
                 glue_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number_of_nodes: Optional[pulumi.Input[int]] = None,
                 number_of_workers: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 public_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_configuration: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 worker_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DevEndpoint resource.
        :param pulumi.Input[str] role_arn: The IAM role for this endpoint.
        :param pulumi.Input[Mapping[str, Any]] arguments: A map of arguments used to configure the endpoint.
        :param pulumi.Input[str] extra_jars_s3_path: Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
        :param pulumi.Input[str] extra_python_libs_s3_path: Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
        :param pulumi.Input[str] glue_version: -  Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
        :param pulumi.Input[str] name: The name of this endpoint. It must be unique in your account.
        :param pulumi.Input[int] number_of_nodes: The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `worker_type`.
        :param pulumi.Input[int] number_of_workers: The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
        :param pulumi.Input[str] public_key: The public key to be used by this endpoint for authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_keys: A list of public keys to be used by this endpoint for authentication.
        :param pulumi.Input[str] security_configuration: The name of the Security Configuration structure to be used with this endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Security group IDs for the security groups to be used by this endpoint.
        :param pulumi.Input[str] subnet_id: The subnet ID for the new endpoint to use.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        :param pulumi.Input[str] worker_type: The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
        """
        pulumi.set(__self__, "role_arn", role_arn)
        if arguments is not None:
            pulumi.set(__self__, "arguments", arguments)
        if extra_jars_s3_path is not None:
            pulumi.set(__self__, "extra_jars_s3_path", extra_jars_s3_path)
        if extra_python_libs_s3_path is not None:
            pulumi.set(__self__, "extra_python_libs_s3_path", extra_python_libs_s3_path)
        if glue_version is not None:
            pulumi.set(__self__, "glue_version", glue_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if number_of_nodes is not None:
            pulumi.set(__self__, "number_of_nodes", number_of_nodes)
        if number_of_workers is not None:
            pulumi.set(__self__, "number_of_workers", number_of_workers)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if public_keys is not None:
            pulumi.set(__self__, "public_keys", public_keys)
        if security_configuration is not None:
            pulumi.set(__self__, "security_configuration", security_configuration)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if worker_type is not None:
            pulumi.set(__self__, "worker_type", worker_type)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The IAM role for this endpoint.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def arguments(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map of arguments used to configure the endpoint.
        """
        return pulumi.get(self, "arguments")

    @arguments.setter
    def arguments(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "arguments", value)

    @property
    @pulumi.getter(name="extraJarsS3Path")
    def extra_jars_s3_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
        """
        return pulumi.get(self, "extra_jars_s3_path")

    @extra_jars_s3_path.setter
    def extra_jars_s3_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extra_jars_s3_path", value)

    @property
    @pulumi.getter(name="extraPythonLibsS3Path")
    def extra_python_libs_s3_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
        """
        return pulumi.get(self, "extra_python_libs_s3_path")

    @extra_python_libs_s3_path.setter
    def extra_python_libs_s3_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extra_python_libs_s3_path", value)

    @property
    @pulumi.getter(name="glueVersion")
    def glue_version(self) -> Optional[pulumi.Input[str]]:
        """
        -  Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
        """
        return pulumi.get(self, "glue_version")

    @glue_version.setter
    def glue_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "glue_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this endpoint. It must be unique in your account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="numberOfNodes")
    def number_of_nodes(self) -> Optional[pulumi.Input[int]]:
        """
        The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `worker_type`.
        """
        return pulumi.get(self, "number_of_nodes")

    @number_of_nodes.setter
    def number_of_nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "number_of_nodes", value)

    @property
    @pulumi.getter(name="numberOfWorkers")
    def number_of_workers(self) -> Optional[pulumi.Input[int]]:
        """
        The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
        """
        return pulumi.get(self, "number_of_workers")

    @number_of_workers.setter
    def number_of_workers(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "number_of_workers", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public key to be used by this endpoint for authentication.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of public keys to be used by this endpoint for authentication.
        """
        return pulumi.get(self, "public_keys")

    @public_keys.setter
    def public_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "public_keys", value)

    @property
    @pulumi.getter(name="securityConfiguration")
    def security_configuration(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Security Configuration structure to be used with this endpoint.
        """
        return pulumi.get(self, "security_configuration")

    @security_configuration.setter
    def security_configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_configuration", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Security group IDs for the security groups to be used by this endpoint.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet ID for the new endpoint to use.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="workerType")
    def worker_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
        """
        return pulumi.get(self, "worker_type")

    @worker_type.setter
    def worker_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "worker_type", value)


class DevEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DevEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Glue Development Endpoint resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=["sts:AssumeRole"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="Service",
                identifiers=["glue.amazonaws.com"],
            )],
        )])
        example_role = aws.iam.Role("exampleRole", assume_role_policy=example_policy_document.json)
        example_dev_endpoint = aws.glue.DevEndpoint("exampleDevEndpoint", role_arn=example_role.arn)
        example__aws_glue_service_role = aws.iam.RolePolicyAttachment("example-AWSGlueServiceRole",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole",
            role=example_role.name)
        ```

        ## Import

        A Glue Development Endpoint can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:glue/devEndpoint:DevEndpoint example foo
        ```

        :param str resource_name: The name of the resource.
        :param DevEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arguments: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 extra_jars_s3_path: Optional[pulumi.Input[str]] = None,
                 extra_python_libs_s3_path: Optional[pulumi.Input[str]] = None,
                 glue_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number_of_nodes: Optional[pulumi.Input[int]] = None,
                 number_of_workers: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 public_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_configuration: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 worker_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Glue Development Endpoint resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=["sts:AssumeRole"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="Service",
                identifiers=["glue.amazonaws.com"],
            )],
        )])
        example_role = aws.iam.Role("exampleRole", assume_role_policy=example_policy_document.json)
        example_dev_endpoint = aws.glue.DevEndpoint("exampleDevEndpoint", role_arn=example_role.arn)
        example__aws_glue_service_role = aws.iam.RolePolicyAttachment("example-AWSGlueServiceRole",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole",
            role=example_role.name)
        ```

        ## Import

        A Glue Development Endpoint can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:glue/devEndpoint:DevEndpoint example foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] arguments: A map of arguments used to configure the endpoint.
        :param pulumi.Input[str] extra_jars_s3_path: Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
        :param pulumi.Input[str] extra_python_libs_s3_path: Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
        :param pulumi.Input[str] glue_version: -  Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
        :param pulumi.Input[str] name: The name of this endpoint. It must be unique in your account.
        :param pulumi.Input[int] number_of_nodes: The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `worker_type`.
        :param pulumi.Input[int] number_of_workers: The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
        :param pulumi.Input[str] public_key: The public key to be used by this endpoint for authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_keys: A list of public keys to be used by this endpoint for authentication.
        :param pulumi.Input[str] role_arn: The IAM role for this endpoint.
        :param pulumi.Input[str] security_configuration: The name of the Security Configuration structure to be used with this endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Security group IDs for the security groups to be used by this endpoint.
        :param pulumi.Input[str] subnet_id: The subnet ID for the new endpoint to use.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        :param pulumi.Input[str] worker_type: The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DevEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arguments: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 extra_jars_s3_path: Optional[pulumi.Input[str]] = None,
                 extra_python_libs_s3_path: Optional[pulumi.Input[str]] = None,
                 glue_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number_of_nodes: Optional[pulumi.Input[int]] = None,
                 number_of_workers: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 public_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_configuration: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 worker_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['arguments'] = arguments
            __props__['extra_jars_s3_path'] = extra_jars_s3_path
            __props__['extra_python_libs_s3_path'] = extra_python_libs_s3_path
            __props__['glue_version'] = glue_version
            __props__['name'] = name
            __props__['number_of_nodes'] = number_of_nodes
            __props__['number_of_workers'] = number_of_workers
            __props__['public_key'] = public_key
            __props__['public_keys'] = public_keys
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['security_configuration'] = security_configuration
            __props__['security_group_ids'] = security_group_ids
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['worker_type'] = worker_type
            __props__['arn'] = None
            __props__['availability_zone'] = None
            __props__['failure_reason'] = None
            __props__['private_address'] = None
            __props__['public_address'] = None
            __props__['status'] = None
            __props__['vpc_id'] = None
            __props__['yarn_endpoint_address'] = None
            __props__['zeppelin_remote_spark_interpreter_port'] = None
        super(DevEndpoint, __self__).__init__(
            'aws:glue/devEndpoint:DevEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arguments: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            extra_jars_s3_path: Optional[pulumi.Input[str]] = None,
            extra_python_libs_s3_path: Optional[pulumi.Input[str]] = None,
            failure_reason: Optional[pulumi.Input[str]] = None,
            glue_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            number_of_nodes: Optional[pulumi.Input[int]] = None,
            number_of_workers: Optional[pulumi.Input[int]] = None,
            private_address: Optional[pulumi.Input[str]] = None,
            public_address: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            public_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            security_configuration: Optional[pulumi.Input[str]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            worker_type: Optional[pulumi.Input[str]] = None,
            yarn_endpoint_address: Optional[pulumi.Input[str]] = None,
            zeppelin_remote_spark_interpreter_port: Optional[pulumi.Input[int]] = None) -> 'DevEndpoint':
        """
        Get an existing DevEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] arguments: A map of arguments used to configure the endpoint.
        :param pulumi.Input[str] arn: The ARN of the endpoint.
        :param pulumi.Input[str] availability_zone: The AWS availability zone where this endpoint is located.
        :param pulumi.Input[str] extra_jars_s3_path: Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
        :param pulumi.Input[str] extra_python_libs_s3_path: Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
        :param pulumi.Input[str] failure_reason: The reason for a current failure in this endpoint.
        :param pulumi.Input[str] glue_version: -  Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
        :param pulumi.Input[str] name: The name of this endpoint. It must be unique in your account.
        :param pulumi.Input[int] number_of_nodes: The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `worker_type`.
        :param pulumi.Input[int] number_of_workers: The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
        :param pulumi.Input[str] private_address: A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
        :param pulumi.Input[str] public_address: The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
        :param pulumi.Input[str] public_key: The public key to be used by this endpoint for authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_keys: A list of public keys to be used by this endpoint for authentication.
        :param pulumi.Input[str] role_arn: The IAM role for this endpoint.
        :param pulumi.Input[str] security_configuration: The name of the Security Configuration structure to be used with this endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Security group IDs for the security groups to be used by this endpoint.
        :param pulumi.Input[str] status: The current status of this endpoint.
        :param pulumi.Input[str] subnet_id: The subnet ID for the new endpoint to use.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        :param pulumi.Input[str] vpc_id: he ID of the VPC used by this endpoint.
        :param pulumi.Input[str] worker_type: The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
        :param pulumi.Input[str] yarn_endpoint_address: The YARN endpoint address used by this endpoint.
        :param pulumi.Input[int] zeppelin_remote_spark_interpreter_port: The Apache Zeppelin port for the remote Apache Spark interpreter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arguments"] = arguments
        __props__["arn"] = arn
        __props__["availability_zone"] = availability_zone
        __props__["extra_jars_s3_path"] = extra_jars_s3_path
        __props__["extra_python_libs_s3_path"] = extra_python_libs_s3_path
        __props__["failure_reason"] = failure_reason
        __props__["glue_version"] = glue_version
        __props__["name"] = name
        __props__["number_of_nodes"] = number_of_nodes
        __props__["number_of_workers"] = number_of_workers
        __props__["private_address"] = private_address
        __props__["public_address"] = public_address
        __props__["public_key"] = public_key
        __props__["public_keys"] = public_keys
        __props__["role_arn"] = role_arn
        __props__["security_configuration"] = security_configuration
        __props__["security_group_ids"] = security_group_ids
        __props__["status"] = status
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        __props__["worker_type"] = worker_type
        __props__["yarn_endpoint_address"] = yarn_endpoint_address
        __props__["zeppelin_remote_spark_interpreter_port"] = zeppelin_remote_spark_interpreter_port
        return DevEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arguments(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A map of arguments used to configure the endpoint.
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the endpoint.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The AWS availability zone where this endpoint is located.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="extraJarsS3Path")
    def extra_jars_s3_path(self) -> pulumi.Output[Optional[str]]:
        """
        Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
        """
        return pulumi.get(self, "extra_jars_s3_path")

    @property
    @pulumi.getter(name="extraPythonLibsS3Path")
    def extra_python_libs_s3_path(self) -> pulumi.Output[Optional[str]]:
        """
        Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
        """
        return pulumi.get(self, "extra_python_libs_s3_path")

    @property
    @pulumi.getter(name="failureReason")
    def failure_reason(self) -> pulumi.Output[str]:
        """
        The reason for a current failure in this endpoint.
        """
        return pulumi.get(self, "failure_reason")

    @property
    @pulumi.getter(name="glueVersion")
    def glue_version(self) -> pulumi.Output[Optional[str]]:
        """
        -  Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
        """
        return pulumi.get(self, "glue_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this endpoint. It must be unique in your account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfNodes")
    def number_of_nodes(self) -> pulumi.Output[Optional[int]]:
        """
        The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `worker_type`.
        """
        return pulumi.get(self, "number_of_nodes")

    @property
    @pulumi.getter(name="numberOfWorkers")
    def number_of_workers(self) -> pulumi.Output[Optional[int]]:
        """
        The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
        """
        return pulumi.get(self, "number_of_workers")

    @property
    @pulumi.getter(name="privateAddress")
    def private_address(self) -> pulumi.Output[str]:
        """
        A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
        """
        return pulumi.get(self, "private_address")

    @property
    @pulumi.getter(name="publicAddress")
    def public_address(self) -> pulumi.Output[str]:
        """
        The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
        """
        return pulumi.get(self, "public_address")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[Optional[str]]:
        """
        The public key to be used by this endpoint for authentication.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of public keys to be used by this endpoint for authentication.
        """
        return pulumi.get(self, "public_keys")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The IAM role for this endpoint.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="securityConfiguration")
    def security_configuration(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Security Configuration structure to be used with this endpoint.
        """
        return pulumi.get(self, "security_configuration")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Security group IDs for the security groups to be used by this endpoint.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current status of this endpoint.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        The subnet ID for the new endpoint to use.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        he ID of the VPC used by this endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="workerType")
    def worker_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
        """
        return pulumi.get(self, "worker_type")

    @property
    @pulumi.getter(name="yarnEndpointAddress")
    def yarn_endpoint_address(self) -> pulumi.Output[str]:
        """
        The YARN endpoint address used by this endpoint.
        """
        return pulumi.get(self, "yarn_endpoint_address")

    @property
    @pulumi.getter(name="zeppelinRemoteSparkInterpreterPort")
    def zeppelin_remote_spark_interpreter_port(self) -> pulumi.Output[int]:
        """
        The Apache Zeppelin port for the remote Apache Spark interpreter.
        """
        return pulumi.get(self, "zeppelin_remote_spark_interpreter_port")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


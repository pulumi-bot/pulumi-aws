# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['FileSystemPolicyArgs', 'FileSystemPolicy']

@pulumi.input_type
class FileSystemPolicyArgs:
    def __init__(__self__, *,
                 file_system_id: pulumi.Input[str],
                 policy: pulumi.Input[str]):
        """
        The set of arguments for constructing a FileSystemPolicy resource.
        :param pulumi.Input[str] file_system_id: The ID of the EFS file system.
        :param pulumi.Input[str] policy: The JSON formatted file system policy for the EFS file system. see [Docs](https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies) for more info.
        """
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        The ID of the EFS file system.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The JSON formatted file system policy for the EFS file system. see [Docs](https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies) for more info.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)


class FileSystemPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Elastic File System (EFS) File System Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        fs = aws.efs.FileSystem("fs")
        policy = aws.efs.FileSystemPolicy("policy",
            file_system_id=fs.id,
            policy=f\"\"\"{{
            "Version": "2012-10-17",
            "Id": "ExamplePolicy01",
            "Statement": [
                {{
                    "Sid": "ExampleStatement01",
                    "Effect": "Allow",
                    "Principal": {{
                        "AWS": "*"
                    }},
                    "Resource": "{aws_efs_file_system["test"]["arn"]}",
                    "Action": [
                        "elasticfilesystem:ClientMount",
                        "elasticfilesystem:ClientWrite"
                    ],
                    "Condition": {{
                        "Bool": {{
                            "aws:SecureTransport": "true"
                        }}
                    }}
                }}
            ]
        }}
        \"\"\")
        ```

        ## Import

        The EFS file system policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:efs/fileSystemPolicy:FileSystemPolicy foo fs-6fa144c6
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_system_id: The ID of the EFS file system.
        :param pulumi.Input[str] policy: The JSON formatted file system policy for the EFS file system. see [Docs](https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies) for more info.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileSystemPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Elastic File System (EFS) File System Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        fs = aws.efs.FileSystem("fs")
        policy = aws.efs.FileSystemPolicy("policy",
            file_system_id=fs.id,
            policy=f\"\"\"{{
            "Version": "2012-10-17",
            "Id": "ExamplePolicy01",
            "Statement": [
                {{
                    "Sid": "ExampleStatement01",
                    "Effect": "Allow",
                    "Principal": {{
                        "AWS": "*"
                    }},
                    "Resource": "{aws_efs_file_system["test"]["arn"]}",
                    "Action": [
                        "elasticfilesystem:ClientMount",
                        "elasticfilesystem:ClientWrite"
                    ],
                    "Condition": {{
                        "Bool": {{
                            "aws:SecureTransport": "true"
                        }}
                    }}
                }}
            ]
        }}
        \"\"\")
        ```

        ## Import

        The EFS file system policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:efs/fileSystemPolicy:FileSystemPolicy foo fs-6fa144c6
        ```

        :param str resource_name: The name of the resource.
        :param FileSystemPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileSystemPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
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

            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__['file_system_id'] = file_system_id
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
        super(FileSystemPolicy, __self__).__init__(
            'aws:efs/fileSystemPolicy:FileSystemPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'FileSystemPolicy':
        """
        Get an existing FileSystemPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_system_id: The ID of the EFS file system.
        :param pulumi.Input[str] policy: The JSON formatted file system policy for the EFS file system. see [Docs](https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies) for more info.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["file_system_id"] = file_system_id
        __props__["policy"] = policy
        return FileSystemPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        The ID of the EFS file system.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The JSON formatted file system policy for the EFS file system. see [Docs](https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies) for more info.
        """
        return pulumi.get(self, "policy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


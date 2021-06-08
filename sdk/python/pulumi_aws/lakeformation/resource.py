# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourceArgs', 'Resource']

@pulumi.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str],
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Resource resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the resource, an S3 path.
        :param pulumi.Input[str] role_arn: Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        pulumi.set(__self__, "arn", arn)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the resource, an S3 path.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class _ResourceState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 last_modified: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Resource resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the resource, an S3 path.
        :param pulumi.Input[str] last_modified: (Optional) The date and time the resource was last modified in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        :param pulumi.Input[str] role_arn: Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if last_modified is not None:
            pulumi.set(__self__, "last_modified", last_modified)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the resource, an S3 path.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) The date and time the resource was last modified in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        """
        return pulumi.get(self, "last_modified")

    @last_modified.setter
    def last_modified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modified", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class Resource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Registers a Lake Formation resource (e.g. S3 bucket) as managed by the Data Catalog. In other words, the S3 path is added to the data lake.

        Choose a role that has read/write access to the chosen Amazon S3 path or use the service-linked role. When you register the S3 path, the service-linked role and a new inline policy are created on your behalf. Lake Formation adds the first path to the inline policy and attaches it to the service-linked role. When you register subsequent paths, Lake Formation adds the path to the existing policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.get_bucket(bucket="an-example-bucket")
        example_resource = aws.lakeformation.Resource("exampleResource", arn=example_bucket.arn)
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the resource, an S3 path.
        :param pulumi.Input[str] role_arn: Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: ResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Registers a Lake Formation resource (e.g. S3 bucket) as managed by the Data Catalog. In other words, the S3 path is added to the data lake.

        Choose a role that has read/write access to the chosen Amazon S3 path or use the service-linked role. When you register the S3 path, the service-linked role and a new inline policy are created on your behalf. Lake Formation adds the first path to the inline policy and attaches it to the service-linked role. When you register subsequent paths, Lake Formation adds the path to the existing policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.get_bucket(bucket="an-example-bucket")
        example_resource = aws.lakeformation.Resource("exampleResource", arn=example_bucket.arn)
        ```

        :param str resource_name_: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceArgs.__new__(ResourceArgs)

            if arn is None and not opts.urn:
                raise TypeError("Missing required property 'arn'")
            __props__.__dict__["arn"] = arn
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["last_modified"] = None
        super(Resource, __self__).__init__(
            'aws:lakeformation/resource:Resource',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            last_modified: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the resource, an S3 path.
        :param pulumi.Input[str] last_modified: (Optional) The date and time the resource was last modified in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        :param pulumi.Input[str] role_arn: Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceState.__new__(_ResourceState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["last_modified"] = last_modified
        __props__.__dict__["role_arn"] = role_arn
        return Resource(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the resource, an S3 path.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        (Optional) The date and time the resource was last modified in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        Role that has read/write access to the resource. If not provided, the Lake Formation service-linked role must exist and is used.
        """
        return pulumi.get(self, "role_arn")


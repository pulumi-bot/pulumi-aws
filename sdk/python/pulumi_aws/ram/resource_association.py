# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ResourceAssociationArgs', 'ResourceAssociation']

@pulumi.input_type
class ResourceAssociationArgs:
    def __init__(__self__, *,
                 resource_arn: pulumi.Input[str],
                 resource_share_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a ResourceAssociation resource.
        :param pulumi.Input[str] resource_arn: Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        :param pulumi.Input[str] resource_share_arn: Amazon Resource Name (ARN) of the RAM Resource Share.
        """
        pulumi.set(__self__, "resource_arn", resource_arn)
        pulumi.set(__self__, "resource_share_arn", resource_share_arn)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        """
        return pulumi.get(self, "resource_arn")

    @resource_arn.setter
    def resource_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_arn", value)

    @property
    @pulumi.getter(name="resourceShareArn")
    def resource_share_arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the RAM Resource Share.
        """
        return pulumi.get(self, "resource_share_arn")

    @resource_share_arn.setter
    def resource_share_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_share_arn", value)


class ResourceAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Resource Access Manager (RAM) Resource Association.

        > *NOTE:* Certain AWS resources (e.g. EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ram.ResourceAssociation("example",
            resource_arn=aws_subnet["example"]["arn"],
            resource_share_arn=aws_ram_resource_share["example"]["arn"])
        ```

        ## Import

        RAM Resource Associations can be imported using their Resource Share ARN and Resource ARN separated by a comma, e.g.

        ```sh
         $ pulumi import aws:ram/resourceAssociation:ResourceAssociation example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,arn:aws:ec2:eu-west-1:123456789012:subnet/subnet-12345678
        ```

        :param str resource_name: The name of the resource.
        :param ResourceAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 resource_share_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Resource Access Manager (RAM) Resource Association.

        > *NOTE:* Certain AWS resources (e.g. EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ram.ResourceAssociation("example",
            resource_arn=aws_subnet["example"]["arn"],
            resource_share_arn=aws_ram_resource_share["example"]["arn"])
        ```

        ## Import

        RAM Resource Associations can be imported using their Resource Share ARN and Resource ARN separated by a comma, e.g.

        ```sh
         $ pulumi import aws:ram/resourceAssociation:ResourceAssociation example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,arn:aws:ec2:eu-west-1:123456789012:subnet/subnet-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        :param pulumi.Input[str] resource_share_arn: Amazon Resource Name (ARN) of the RAM Resource Share.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 resource_share_arn: Optional[pulumi.Input[str]] = None,
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

            if resource_arn is None and not opts.urn:
                raise TypeError("Missing required property 'resource_arn'")
            __props__['resource_arn'] = resource_arn
            if resource_share_arn is None and not opts.urn:
                raise TypeError("Missing required property 'resource_share_arn'")
            __props__['resource_share_arn'] = resource_share_arn
        super(ResourceAssociation, __self__).__init__(
            'aws:ram/resourceAssociation:ResourceAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            resource_arn: Optional[pulumi.Input[str]] = None,
            resource_share_arn: Optional[pulumi.Input[str]] = None) -> 'ResourceAssociation':
        """
        Get an existing ResourceAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        :param pulumi.Input[str] resource_share_arn: Amazon Resource Name (ARN) of the RAM Resource Share.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["resource_arn"] = resource_arn
        __props__["resource_share_arn"] = resource_share_arn
        return ResourceAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="resourceShareArn")
    def resource_share_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the RAM Resource Share.
        """
        return pulumi.get(self, "resource_share_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


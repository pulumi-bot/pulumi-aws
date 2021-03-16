# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SecurityGroupArgs', 'SecurityGroup']

@pulumi.input_type
class SecurityGroupArgs:
    def __init__(__self__, *,
                 ingress: pulumi.Input[Sequence[pulumi.Input['SecurityGroupIngressArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecurityGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['SecurityGroupIngressArgs']]] ingress: A list of ingress rules.
        :param pulumi.Input[str] description: The description of the DB security group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[str] name: The name of the DB security group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        pulumi.set(__self__, "ingress", ingress)
        if description is None:
            description = 'Managed by Pulumi'
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def ingress(self) -> pulumi.Input[Sequence[pulumi.Input['SecurityGroupIngressArgs']]]:
        """
        A list of ingress rules.
        """
        return pulumi.get(self, "ingress")

    @ingress.setter
    def ingress(self, value: pulumi.Input[Sequence[pulumi.Input['SecurityGroupIngressArgs']]]):
        pulumi.set(self, "ingress", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the DB security group. Defaults to "Managed by Pulumi".
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the DB security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class SecurityGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an RDS security group resource. This is only for DB instances in the
        EC2-Classic Platform. For instances inside a VPC, use the
        `aws_db_instance.vpc_security_group_ids`
        attribute instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.rds.SecurityGroup("default", ingress=[aws.rds.SecurityGroupIngressArgs(
            cidr="10.0.0.0/24",
        )])
        ```

        ## Import

        DB Security groups can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:rds/securityGroup:SecurityGroup default aws_rds_sg-1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the DB security group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]] ingress: A list of ingress rules.
        :param pulumi.Input[str] name: The name of the DB security group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an RDS security group resource. This is only for DB instances in the
        EC2-Classic Platform. For instances inside a VPC, use the
        `aws_db_instance.vpc_security_group_ids`
        attribute instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.rds.SecurityGroup("default", ingress=[aws.rds.SecurityGroupIngressArgs(
            cidr="10.0.0.0/24",
        )])
        ```

        ## Import

        DB Security groups can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:rds/securityGroup:SecurityGroup default aws_rds_sg-1
        ```

        :param str resource_name: The name of the resource.
        :param SecurityGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            if ingress is None and not opts.urn:
                raise TypeError("Missing required property 'ingress'")
            __props__['ingress'] = ingress
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(SecurityGroup, __self__).__init__(
            'aws:rds/securityGroup:SecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ingress: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SecurityGroup':
        """
        Get an existing SecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The arn of the DB security group.
        :param pulumi.Input[str] description: The description of the DB security group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupIngressArgs']]]] ingress: A list of ingress rules.
        :param pulumi.Input[str] name: The name of the DB security group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["ingress"] = ingress
        __props__["name"] = name
        __props__["tags"] = tags
        return SecurityGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The arn of the DB security group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the DB security group. Defaults to "Managed by Pulumi".
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ingress(self) -> pulumi.Output[Sequence['outputs.SecurityGroupIngress']]:
        """
        A list of ingress rules.
        """
        return pulumi.get(self, "ingress")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the DB security group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcLinkArgs', 'VpcLink']

@pulumi.input_type
class VpcLinkArgs:
    def __init__(__self__, *,
                 target_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VpcLink resource.
        :param pulumi.Input[str] target_arn: The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        :param pulumi.Input[str] description: The description of the VPC link.
        :param pulumi.Input[str] name: The name used to label and identify the VPC link.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        pulumi.set(__self__, "target_arn", target_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Input[str]:
        """
        The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        """
        return pulumi.get(self, "target_arn")

    @target_arn.setter
    def target_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the VPC link.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name used to label and identify the VPC link.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _VpcLinkState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcLink resources.
        :param pulumi.Input[str] description: The description of the VPC link.
        :param pulumi.Input[str] name: The name used to label and identify the VPC link.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] target_arn: The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if target_arn is not None:
            pulumi.set(__self__, "target_arn", target_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the VPC link.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name used to label and identify the VPC link.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        """
        return pulumi.get(self, "target_arn")

    @target_arn.setter
    def target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_arn", value)


class VpcLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an API Gateway VPC Link.

        > **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
        To enable private integration for HTTP APIs, use the `Amazon API Gateway Version 2 VPC Link` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_load_balancer = aws.lb.LoadBalancer("exampleLoadBalancer",
            internal=True,
            load_balancer_type="network",
            subnet_mappings=[aws.lb.LoadBalancerSubnetMappingArgs(
                subnet_id="12345",
            )])
        example_vpc_link = aws.apigateway.VpcLink("exampleVpcLink",
            description="example description",
            target_arn=[example_load_balancer.arn])
        ```

        ## Import

        API Gateway VPC Link can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:apigateway/vpcLink:VpcLink example <vpc_link_id>
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the VPC link.
        :param pulumi.Input[str] name: The name used to label and identify the VPC link.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] target_arn: The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: VpcLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an API Gateway VPC Link.

        > **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
        To enable private integration for HTTP APIs, use the `Amazon API Gateway Version 2 VPC Link` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_load_balancer = aws.lb.LoadBalancer("exampleLoadBalancer",
            internal=True,
            load_balancer_type="network",
            subnet_mappings=[aws.lb.LoadBalancerSubnetMappingArgs(
                subnet_id="12345",
            )])
        example_vpc_link = aws.apigateway.VpcLink("exampleVpcLink",
            description="example description",
            target_arn=[example_load_balancer.arn])
        ```

        ## Import

        API Gateway VPC Link can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:apigateway/vpcLink:VpcLink example <vpc_link_id>
        ```

        :param str resource_name_: The name of the resource.
        :param VpcLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcLinkArgs.__new__(VpcLinkArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            if target_arn is None and not opts.urn:
                raise TypeError("Missing required property 'target_arn'")
            __props__.__dict__["target_arn"] = target_arn
            __props__.__dict__["arn"] = None
        super(VpcLink, __self__).__init__(
            'aws:apigateway/vpcLink:VpcLink',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_arn: Optional[pulumi.Input[str]] = None) -> 'VpcLink':
        """
        Get an existing VpcLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the VPC link.
        :param pulumi.Input[str] name: The name used to label and identify the VPC link.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] target_arn: The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcLinkState.__new__(_VpcLinkState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["target_arn"] = target_arn
        return VpcLink(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the VPC link.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name used to label and identify the VPC link.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Output[str]:
        """
        The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
        """
        return pulumi.get(self, "target_arn")


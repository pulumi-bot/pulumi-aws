# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ContainerPolicy']


class ContainerPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a MediaStore Container Policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        current_region = aws.get_region()
        current_caller_identity = aws.get_caller_identity()
        example_container = aws.mediastore.Container("exampleContainer")
        example_container_policy = aws.mediastore.ContainerPolicy("exampleContainerPolicy",
            container_name=example_container.name,
            policy=example_container.name.apply(lambda name: f\"\"\"{{
        	"Version": "2012-10-17",
        	"Statement": [{{
        		"Sid": "MediaStoreFullAccess",
        		"Action": [ "mediastore:*" ],
        		"Principal": {{"AWS" : "arn:aws:iam::{current_caller_identity.account_id}:root"}},
        		"Effect": "Allow",
        		"Resource": "arn:aws:mediastore:{current_region.name}:{current_caller_identity.account_id}:container/{name}/*",
        		"Condition": {{
        			"Bool": {{ "aws:SecureTransport": "true" }}
        		}}
        	}}]
        }}
        \"\"\"))
        ```

        ## Import

        MediaStore Container Policy can be imported using the MediaStore Container Name, e.g.

        ```sh
         $ pulumi import aws:mediastore/containerPolicy:ContainerPolicy example example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_name: The name of the container.
        :param pulumi.Input[str] policy: The contents of the policy.
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

            if container_name is None and not opts.urn:
                raise TypeError("Missing required property 'container_name'")
            __props__['container_name'] = container_name
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
        super(ContainerPolicy, __self__).__init__(
            'aws:mediastore/containerPolicy:ContainerPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_name: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'ContainerPolicy':
        """
        Get an existing ContainerPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_name: The name of the container.
        :param pulumi.Input[str] policy: The contents of the policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["container_name"] = container_name
        __props__["policy"] = policy
        return ContainerPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Output[str]:
        """
        The name of the container.
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The contents of the policy.
        """
        return pulumi.get(self, "policy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


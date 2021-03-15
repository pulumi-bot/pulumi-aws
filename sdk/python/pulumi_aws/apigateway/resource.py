# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ResourceArgs', 'Resource']

@pulumi.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 parent_id: pulumi.Input[str],
                 path_part: pulumi.Input[str],
                 rest_api: pulumi.Input[str]):
        """
        The set of arguments for constructing a Resource resource.
        :param pulumi.Input[str] parent_id: The ID of the parent API resource
        :param pulumi.Input[str] path_part: The last path segment of this API resource.
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        """
        pulumi.set(__self__, "parent_id", parent_id)
        pulumi.set(__self__, "path_part", path_part)
        pulumi.set(__self__, "rest_api", rest_api)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Input[str]:
        """
        The ID of the parent API resource
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter(name="pathPart")
    def path_part(self) -> pulumi.Input[str]:
        """
        The last path segment of this API resource.
        """
        return pulumi.get(self, "path_part")

    @path_part.setter
    def path_part(self, value: pulumi.Input[str]):
        pulumi.set(self, "path_part", value)

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Input[str]:
        """
        The ID of the associated REST API
        """
        return pulumi.get(self, "rest_api")

    @rest_api.setter
    def rest_api(self, value: pulumi.Input[str]):
        pulumi.set(self, "rest_api", value)


class Resource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an API Gateway Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
        my_demo_resource = aws.apigateway.Resource("myDemoResource",
            rest_api=my_demo_api.id,
            parent_id=my_demo_api.root_resource_id,
            path_part="mydemoresource")
        ```

        ## Import

        `aws_api_gateway_resource` can be imported using `REST-API-ID/RESOURCE-ID`, e.g.

        ```sh
         $ pulumi import aws:apigateway/resource:Resource example 12345abcde/67890fghij
        ```

        :param str resource_name: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 path_part: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an API Gateway Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
        my_demo_resource = aws.apigateway.Resource("myDemoResource",
            rest_api=my_demo_api.id,
            parent_id=my_demo_api.root_resource_id,
            path_part="mydemoresource")
        ```

        ## Import

        `aws_api_gateway_resource` can be imported using `REST-API-ID/RESOURCE-ID`, e.g.

        ```sh
         $ pulumi import aws:apigateway/resource:Resource example 12345abcde/67890fghij
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent_id: The ID of the parent API resource
        :param pulumi.Input[str] path_part: The last path segment of this API resource.
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 path_part: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
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

            if parent_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_id'")
            __props__['parent_id'] = parent_id
            if path_part is None and not opts.urn:
                raise TypeError("Missing required property 'path_part'")
            __props__['path_part'] = path_part
            if rest_api is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            __props__['path'] = None
        super(Resource, __self__).__init__(
            'aws:apigateway/resource:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            parent_id: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            path_part: Optional[pulumi.Input[str]] = None,
            rest_api: Optional[pulumi.Input[str]] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent_id: The ID of the parent API resource
        :param pulumi.Input[str] path: The complete path for this API resource, including all parent paths.
        :param pulumi.Input[str] path_part: The last path segment of this API resource.
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["parent_id"] = parent_id
        __props__["path"] = path
        __props__["path_part"] = path_part
        __props__["rest_api"] = rest_api
        return Resource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        The ID of the parent API resource
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The complete path for this API resource, including all parent paths.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="pathPart")
    def path_part(self) -> pulumi.Output[str]:
        """
        The last path segment of this API resource.
        """
        return pulumi.get(self, "path_part")

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Output[str]:
        """
        The ID of the associated REST API
        """
        return pulumi.get(self, "rest_api")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


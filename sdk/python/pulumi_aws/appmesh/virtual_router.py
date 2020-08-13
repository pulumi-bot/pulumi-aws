# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class VirtualRouter(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the virtual router.
    """
    created_date: pulumi.Output[str]
    """
    The creation date of the virtual router.
    """
    last_updated_date: pulumi.Output[str]
    """
    The last update date of the virtual router.
    """
    mesh_name: pulumi.Output[str]
    """
    The name of the service mesh in which to create the virtual router.
    """
    name: pulumi.Output[str]
    """
    The name to use for the virtual router.
    """
    spec: pulumi.Output[dict]
    """
    The virtual router specification to apply.

      * `listener` (`dict`) - The listeners that the virtual router is expected to receive inbound traffic from.
        Currently only one listener is supported per virtual router.
        * `portMapping` (`dict`) - The port mapping information for the listener.
          * `port` (`float`) - The port used for the port mapping.
          * `protocol` (`str`) - The protocol used for the port mapping. Valid values are `http` and `tcp`.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, mesh_name=None, name=None, spec=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS App Mesh virtual router resource.

        ## Breaking Changes

        Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:

        * Remove service `service_names` from the `spec` argument.
        AWS has created a `appmesh.VirtualService` resource for each of service names.
        These resource can be imported using `import`.

        * Add a `listener` configuration block to the `spec` argument.

        The state associated with existing resources will automatically be migrated.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        serviceb = aws.appmesh.VirtualRouter("serviceb",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "listener": {
                    "portMapping": {
                        "port": 8080,
                        "protocol": "http",
                    },
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual router.
        :param pulumi.Input[str] name: The name to use for the virtual router.
        :param pulumi.Input[dict] spec: The virtual router specification to apply.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **spec** object supports the following:

          * `listener` (`pulumi.Input[dict]`) - The listeners that the virtual router is expected to receive inbound traffic from.
            Currently only one listener is supported per virtual router.
            * `portMapping` (`pulumi.Input[dict]`) - The port mapping information for the listener.
              * `port` (`pulumi.Input[float]`) - The port used for the port mapping.
              * `protocol` (`pulumi.Input[str]`) - The protocol used for the port mapping. Valid values are `http` and `tcp`.
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

            if mesh_name is None:
                raise TypeError("Missing required property 'mesh_name'")
            __props__['meshName'] = mesh_name
            __props__['name'] = name
            if spec is None:
                raise TypeError("Missing required property 'spec'")
            __props__['spec'] = spec
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['last_updated_date'] = None
        super(VirtualRouter, __self__).__init__(
            'aws:appmesh/virtualRouter:VirtualRouter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, created_date=None, last_updated_date=None, mesh_name=None, name=None, spec=None, tags=None):
        """
        Get an existing VirtualRouter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the virtual router.
        :param pulumi.Input[str] created_date: The creation date of the virtual router.
        :param pulumi.Input[str] last_updated_date: The last update date of the virtual router.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual router.
        :param pulumi.Input[str] name: The name to use for the virtual router.
        :param pulumi.Input[dict] spec: The virtual router specification to apply.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **spec** object supports the following:

          * `listener` (`pulumi.Input[dict]`) - The listeners that the virtual router is expected to receive inbound traffic from.
            Currently only one listener is supported per virtual router.
            * `portMapping` (`pulumi.Input[dict]`) - The port mapping information for the listener.
              * `port` (`pulumi.Input[float]`) - The port used for the port mapping.
              * `protocol` (`pulumi.Input[str]`) - The protocol used for the port mapping. Valid values are `http` and `tcp`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["created_date"] = created_date
        __props__["last_updated_date"] = last_updated_date
        __props__["mesh_name"] = mesh_name
        __props__["name"] = name
        __props__["spec"] = spec
        __props__["tags"] = tags
        return VirtualRouter(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

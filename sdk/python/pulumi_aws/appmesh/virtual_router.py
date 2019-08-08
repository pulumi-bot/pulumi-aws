# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

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
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, mesh_name=None, name=None, spec=None, tags=None, __name__=None, __opts__=None):
        """
        Provides an AWS App Mesh virtual router resource.
        
        ## Breaking Changes
        
        Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
        
        * Remove service `service_names` from the `spec` argument.
        AWS has created a `appmesh.VirtualService` resource for each of service names.
        These resource can be imported using `import`.
        
        * Add a `listener` configuration block to the `spec` argument.
        
        The state associated with existing resources will automatically be migrated.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual router.
        :param pulumi.Input[str] name: The name to use for the virtual router.
        :param pulumi.Input[dict] spec: The virtual router specification to apply.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_router.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if mesh_name is None:
            raise TypeError("Missing required property 'mesh_name'")
        __props__['mesh_name'] = mesh_name
        __props__['name'] = name
        if spec is None:
            raise TypeError("Missing required property 'spec'")
        __props__['spec'] = spec
        __props__['tags'] = tags
        __props__['arn'] = None
        __props__['created_date'] = None
        __props__['last_updated_date'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(VirtualRouter, __self__).__init__(
            'aws:appmesh/virtualRouter:VirtualRouter',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


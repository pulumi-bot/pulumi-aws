# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VirtualNode']


class VirtualNode(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN of the virtual node.
    """

    created_date: pulumi.Output[str] = pulumi.property("createdDate")
    """
    The creation date of the virtual node.
    """

    last_updated_date: pulumi.Output[str] = pulumi.property("lastUpdatedDate")
    """
    The last update date of the virtual node.
    """

    mesh_name: pulumi.Output[str] = pulumi.property("meshName")
    """
    The name of the service mesh in which to create the virtual node.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name to use for the virtual node.
    """

    spec: pulumi.Output['outputs.VirtualNodeSpec'] = pulumi.property("spec")
    """
    The virtual node specification to apply.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the resource.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mesh_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['VirtualNodeSpecArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS App Mesh virtual node resource.

        ## Breaking Changes

        Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92)), `appmesh.VirtualNode` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:

        * Rename the `service_name` attribute of the `dns` object to `hostname`.

        * Replace the `backends` attribute of the `spec` object with one or more `backend` configuration blocks,
        setting `virtual_service_name` to the name of the service.

        The state associated with existing resources will automatically be migrated.

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        serviceb1 = aws.appmesh.VirtualNode("serviceb1",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "backends": [{
                    "virtualService": {
                        "virtualServiceName": "servicea.simpleapp.local",
                    },
                }],
                "listener": {
                    "portMapping": {
                        "port": 8080,
                        "protocol": "http",
                    },
                },
                "serviceDiscovery": {
                    "dns": {
                        "hostname": "serviceb.simpleapp.local",
                    },
                },
            })
        ```
        ### AWS Cloud Map Service Discovery

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.servicediscovery.HttpNamespace("example")
        serviceb1 = aws.appmesh.VirtualNode("serviceb1",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "backends": [{
                    "virtualService": {
                        "virtualServiceName": "servicea.simpleapp.local",
                    },
                }],
                "listener": {
                    "portMapping": {
                        "port": 8080,
                        "protocol": "http",
                    },
                },
                "serviceDiscovery": {
                    "awsCloudMap": {
                        "attributes": {
                            "stack": "blue",
                        },
                        "service_name": "serviceb1",
                        "namespaceName": example.name,
                    },
                },
            })
        ```
        ### Listener Health Check

        ```python
        import pulumi
        import pulumi_aws as aws

        serviceb1 = aws.appmesh.VirtualNode("serviceb1",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "backends": [{
                    "virtualService": {
                        "virtualServiceName": "servicea.simpleapp.local",
                    },
                }],
                "listener": {
                    "portMapping": {
                        "port": 8080,
                        "protocol": "http",
                    },
                    "health_check": {
                        "protocol": "http",
                        "path": "/ping",
                        "healthyThreshold": 2,
                        "unhealthyThreshold": 2,
                        "timeoutMillis": 2000,
                        "intervalMillis": 5000,
                    },
                },
                "serviceDiscovery": {
                    "dns": {
                        "hostname": "serviceb.simpleapp.local",
                    },
                },
            })
        ```
        ### Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        serviceb1 = aws.appmesh.VirtualNode("serviceb1",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "backends": [{
                    "virtualService": {
                        "virtualServiceName": "servicea.simpleapp.local",
                    },
                }],
                "listener": {
                    "portMapping": {
                        "port": 8080,
                        "protocol": "http",
                    },
                },
                "serviceDiscovery": {
                    "dns": {
                        "hostname": "serviceb.simpleapp.local",
                    },
                },
                "logging": {
                    "accessLog": {
                        "file": {
                            "path": "/dev/stdout",
                        },
                    },
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual node.
        :param pulumi.Input[str] name: The name to use for the virtual node.
        :param pulumi.Input[pulumi.InputType['VirtualNodeSpecArgs']] spec: The virtual node specification to apply.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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
            __props__['mesh_name'] = mesh_name
            __props__['name'] = name
            if spec is None:
                raise TypeError("Missing required property 'spec'")
            __props__['spec'] = spec
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['last_updated_date'] = None
        super(VirtualNode, __self__).__init__(
            'aws:appmesh/virtualNode:VirtualNode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            created_date: Optional[pulumi.Input[str]] = None,
            last_updated_date: Optional[pulumi.Input[str]] = None,
            mesh_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            spec: Optional[pulumi.Input[pulumi.InputType['VirtualNodeSpecArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'VirtualNode':
        """
        Get an existing VirtualNode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the virtual node.
        :param pulumi.Input[str] created_date: The creation date of the virtual node.
        :param pulumi.Input[str] last_updated_date: The last update date of the virtual node.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual node.
        :param pulumi.Input[str] name: The name to use for the virtual node.
        :param pulumi.Input[pulumi.InputType['VirtualNodeSpecArgs']] spec: The virtual node specification to apply.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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
        return VirtualNode(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


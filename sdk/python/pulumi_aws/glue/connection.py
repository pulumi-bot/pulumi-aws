# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Connection']


class Connection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 connection_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 match_criterias: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_connection_requirements: Optional[pulumi.Input[pulumi.InputType['ConnectionPhysicalConnectionRequirementsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Glue Connection resource.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] connection_properties: A map of key-value pairs used as parameters for this connection.
        :param pulumi.Input[str] connection_type: The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
        :param pulumi.Input[str] description: Description of the connection.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] match_criterias: A list of criteria that can be used in selecting this connection.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[pulumi.InputType['ConnectionPhysicalConnectionRequirementsArgs']] physical_connection_requirements: A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
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

            __props__['catalog_id'] = catalog_id
            if connection_properties is None:
                raise TypeError("Missing required property 'connection_properties'")
            __props__['connection_properties'] = connection_properties
            __props__['connection_type'] = connection_type
            __props__['description'] = description
            __props__['match_criterias'] = match_criterias
            __props__['name'] = name
            __props__['physical_connection_requirements'] = physical_connection_requirements
            __props__['arn'] = None
        super(Connection, __self__).__init__(
            'aws:glue/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            connection_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            connection_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            match_criterias: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            physical_connection_requirements: Optional[pulumi.Input[pulumi.InputType['ConnectionPhysicalConnectionRequirementsArgs']]] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Glue Connection.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] connection_properties: A map of key-value pairs used as parameters for this connection.
        :param pulumi.Input[str] connection_type: The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
        :param pulumi.Input[str] description: Description of the connection.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] match_criterias: A list of criteria that can be used in selecting this connection.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[pulumi.InputType['ConnectionPhysicalConnectionRequirementsArgs']] physical_connection_requirements: A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["catalog_id"] = catalog_id
        __props__["connection_properties"] = connection_properties
        __props__["connection_type"] = connection_type
        __props__["description"] = description
        __props__["match_criterias"] = match_criterias
        __props__["name"] = name
        __props__["physical_connection_requirements"] = physical_connection_requirements
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Glue Connection.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="connectionProperties")
    def connection_properties(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of key-value pairs used as parameters for this connection.
        """
        return pulumi.get(self, "connection_properties")

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
        """
        return pulumi.get(self, "connection_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="matchCriterias")
    def match_criterias(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of criteria that can be used in selecting this connection.
        """
        return pulumi.get(self, "match_criterias")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the connection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="physicalConnectionRequirements")
    def physical_connection_requirements(self) -> pulumi.Output[Optional['outputs.ConnectionPhysicalConnectionRequirements']]:
        """
        A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
        """
        return pulumi.get(self, "physical_connection_requirements")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


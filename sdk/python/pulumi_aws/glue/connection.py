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

__all__ = ['ConnectionArgs', 'Connection']

@pulumi.input_type
class ConnectionArgs:
    def __init__(__self__, *,
                 connection_properties: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 match_criterias: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_connection_requirements: Optional[pulumi.Input['ConnectionPhysicalConnectionRequirementsArgs']] = None):
        """
        The set of arguments for constructing a Connection resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] connection_properties: A map of key-value pairs used as parameters for this connection.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        :param pulumi.Input[str] connection_type: The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
        :param pulumi.Input[str] description: Description of the connection.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] match_criterias: A list of criteria that can be used in selecting this connection.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input['ConnectionPhysicalConnectionRequirementsArgs'] physical_connection_requirements: A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
        """
        pulumi.set(__self__, "connection_properties", connection_properties)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if connection_type is not None:
            pulumi.set(__self__, "connection_type", connection_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if match_criterias is not None:
            pulumi.set(__self__, "match_criterias", match_criterias)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if physical_connection_requirements is not None:
            pulumi.set(__self__, "physical_connection_requirements", physical_connection_requirements)

    @property
    @pulumi.getter(name="connectionProperties")
    def connection_properties(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        A map of key-value pairs used as parameters for this connection.
        """
        return pulumi.get(self, "connection_properties")

    @connection_properties.setter
    def connection_properties(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "connection_properties", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
        """
        return pulumi.get(self, "connection_type")

    @connection_type.setter
    def connection_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the connection.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="matchCriterias")
    def match_criterias(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of criteria that can be used in selecting this connection.
        """
        return pulumi.get(self, "match_criterias")

    @match_criterias.setter
    def match_criterias(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "match_criterias", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the connection.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="physicalConnectionRequirements")
    def physical_connection_requirements(self) -> Optional[pulumi.Input['ConnectionPhysicalConnectionRequirementsArgs']]:
        """
        A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
        """
        return pulumi.get(self, "physical_connection_requirements")

    @physical_connection_requirements.setter
    def physical_connection_requirements(self, value: Optional[pulumi.Input['ConnectionPhysicalConnectionRequirementsArgs']]):
        pulumi.set(self, "physical_connection_requirements", value)


class Connection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Glue Connection resource.

        ## Example Usage
        ### Non-VPC Connection

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Connection("example", connection_properties={
            "JDBC_CONNECTION_URL": "jdbc:mysql://example.com/exampledatabase",
            "PASSWORD": "examplepassword",
            "USERNAME": "exampleusername",
        })
        ```
        ### VPC Connection

        For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html#connection-JDBC-VPC).

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Connection("example",
            connection_properties={
                "JDBC_CONNECTION_URL": f"jdbc:mysql://{aws_rds_cluster['example']['endpoint']}/exampledatabase",
                "PASSWORD": "examplepassword",
                "USERNAME": "exampleusername",
            },
            physical_connection_requirements=aws.glue.ConnectionPhysicalConnectionRequirementsArgs(
                availability_zone=aws_subnet["example"]["availability_zone"],
                security_group_id_lists=[aws_security_group["example"]["id"]],
                subnet_id=aws_subnet["example"]["id"],
            ))
        ```

        ## Import

        Glue Connections can be imported using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`, e.g.

        ```sh
         $ pulumi import aws:glue/connection:Connection MyConnection 123456789012:MyConnection
        ```

        :param str resource_name: The name of the resource.
        :param ConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
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
        ### Non-VPC Connection

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Connection("example", connection_properties={
            "JDBC_CONNECTION_URL": "jdbc:mysql://example.com/exampledatabase",
            "PASSWORD": "examplepassword",
            "USERNAME": "exampleusername",
        })
        ```
        ### VPC Connection

        For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html#connection-JDBC-VPC).

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Connection("example",
            connection_properties={
                "JDBC_CONNECTION_URL": f"jdbc:mysql://{aws_rds_cluster['example']['endpoint']}/exampledatabase",
                "PASSWORD": "examplepassword",
                "USERNAME": "exampleusername",
            },
            physical_connection_requirements=aws.glue.ConnectionPhysicalConnectionRequirementsArgs(
                availability_zone=aws_subnet["example"]["availability_zone"],
                security_group_id_lists=[aws_security_group["example"]["id"]],
                subnet_id=aws_subnet["example"]["id"],
            ))
        ```

        ## Import

        Glue Connections can be imported using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`, e.g.

        ```sh
         $ pulumi import aws:glue/connection:Connection MyConnection 123456789012:MyConnection
        ```

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
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
            if connection_properties is None and not opts.urn:
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


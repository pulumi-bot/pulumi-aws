# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Connection(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the Glue Connection.
    """
    catalog_id: pulumi.Output[str]
    """
    The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
    """
    connection_properties: pulumi.Output[dict]
    """
    A map of key-value pairs used as parameters for this connection.
    """
    connection_type: pulumi.Output[str]
    """
    The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`. Defaults to `JBDC`.
    """
    description: pulumi.Output[str]
    """
    Description of the connection.
    """
    match_criterias: pulumi.Output[list]
    """
    A list of criteria that can be used in selecting this connection.
    """
    name: pulumi.Output[str]
    """
    The name of the connection.
    """
    physical_connection_requirements: pulumi.Output[dict]
    """
    A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.

      * `availability_zone` (`str`) - The availability zone of the connection. This field is redundant and implied by `subnet_id`, but is currently an api requirement.
      * `securityGroupIdLists` (`list`) - The security group ID list used by the connection.
      * `subnet_id` (`str`) - The subnet ID used by the connection.
    """
    def __init__(__self__, resource_name, opts=None, catalog_id=None, connection_properties=None, connection_type=None, description=None, match_criterias=None, name=None, physical_connection_requirements=None, __props__=None, __name__=None, __opts__=None):
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

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Connection("example",
            connection_properties={
                "JDBC_CONNECTION_URL": f"jdbc:mysql://{aws_rds_cluster['example']['endpoint']}/exampledatabase",
                "PASSWORD": "examplepassword",
                "USERNAME": "exampleusername",
            },
            physical_connection_requirements={
                "availability_zone": aws_subnet["example"]["availability_zone"],
                "securityGroupIdList": [aws_security_group["example"]["id"]],
                "subnet_id": aws_subnet["example"]["id"],
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        :param pulumi.Input[dict] connection_properties: A map of key-value pairs used as parameters for this connection.
        :param pulumi.Input[str] connection_type: The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`. Defaults to `JBDC`.
        :param pulumi.Input[str] description: Description of the connection.
        :param pulumi.Input[list] match_criterias: A list of criteria that can be used in selecting this connection.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[dict] physical_connection_requirements: A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.

        The **physical_connection_requirements** object supports the following:

          * `availability_zone` (`pulumi.Input[str]`) - The availability zone of the connection. This field is redundant and implied by `subnet_id`, but is currently an api requirement.
          * `securityGroupIdLists` (`pulumi.Input[list]`) - The security group ID list used by the connection.
          * `subnet_id` (`pulumi.Input[str]`) - The subnet ID used by the connection.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, arn=None, catalog_id=None, connection_properties=None, connection_type=None, description=None, match_criterias=None, name=None, physical_connection_requirements=None):
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Glue Connection.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
        :param pulumi.Input[dict] connection_properties: A map of key-value pairs used as parameters for this connection.
        :param pulumi.Input[str] connection_type: The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`. Defaults to `JBDC`.
        :param pulumi.Input[str] description: Description of the connection.
        :param pulumi.Input[list] match_criterias: A list of criteria that can be used in selecting this connection.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[dict] physical_connection_requirements: A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.

        The **physical_connection_requirements** object supports the following:

          * `availability_zone` (`pulumi.Input[str]`) - The availability zone of the connection. This field is redundant and implied by `subnet_id`, but is currently an api requirement.
          * `securityGroupIdLists` (`pulumi.Input[list]`) - The security group ID list used by the connection.
          * `subnet_id` (`pulumi.Input[str]`) - The subnet ID used by the connection.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


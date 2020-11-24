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

__all__ = ['GlobalTable']


class GlobalTable(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replicas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableReplicaArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html). These are layered on top of existing DynamoDB Tables.

        > **NOTE:** To instead manage [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html), use the `dynamodb.Table` resource `replica` configuration block.

        > Note: There are many restrictions before you can properly create DynamoDB Global Tables in multiple regions. See the [AWS DynamoDB Global Table Requirements](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables_reqs_bestpractices.html) for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        us_east_1 = pulumi.providers.Aws("us-east-1", region="us-east-1")
        us_west_2 = pulumi.providers.Aws("us-west-2", region="us-west-2")
        us_east_1_table = aws.dynamodb.Table("us-east-1Table",
            hash_key="myAttribute",
            stream_enabled=True,
            stream_view_type="NEW_AND_OLD_IMAGES",
            read_capacity=1,
            write_capacity=1,
            attributes=[{
                "name": "myAttribute",
                "type": "S",
            }],
            opts=ResourceOptions(provider=aws["us-east-1"]))
        us_west_2_table = aws.dynamodb.Table("us-west-2Table",
            hash_key="myAttribute",
            stream_enabled=True,
            stream_view_type="NEW_AND_OLD_IMAGES",
            read_capacity=1,
            write_capacity=1,
            attributes=[{
                "name": "myAttribute",
                "type": "S",
            }],
            opts=ResourceOptions(provider=aws["us-west-2"]))
        my_table = aws.dynamodb.GlobalTable("myTable", replicas=[
            {
                "regionName": "us-east-1",
            },
            {
                "regionName": "us-west-2",
            },
        ],
        opts=ResourceOptions(provider=aws["us-east-1"],
            depends_on=[
                us_east_1_table,
                us_west_2_table,
            ]))
        ```

        ## Import

        DynamoDB Global Tables can be imported using the global table name, e.g.

        ```sh
         $ pulumi import aws:dynamodb/globalTable:GlobalTable MyTable MyTable
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the global table. Must match underlying DynamoDB Table names in all regions.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableReplicaArgs']]]] replicas: Underlying DynamoDB Table. At least 1 replica must be defined. See below.
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

            __props__['name'] = name
            if replicas is None:
                raise TypeError("Missing required property 'replicas'")
            __props__['replicas'] = replicas
            __props__['arn'] = None
        super(GlobalTable, __self__).__init__(
            'aws:dynamodb/globalTable:GlobalTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            replicas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableReplicaArgs']]]]] = None) -> 'GlobalTable':
        """
        Get an existing GlobalTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the DynamoDB Global Table
        :param pulumi.Input[str] name: The name of the global table. Must match underlying DynamoDB Table names in all regions.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GlobalTableReplicaArgs']]]] replicas: Underlying DynamoDB Table. At least 1 replica must be defined. See below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["replicas"] = replicas
        return GlobalTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the DynamoDB Global Table
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the global table. Must match underlying DynamoDB Table names in all regions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def replicas(self) -> pulumi.Output[Sequence['outputs.GlobalTableReplica']]:
        """
        Underlying DynamoDB Table. At least 1 replica must be defined. See below.
        """
        return pulumi.get(self, "replicas")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


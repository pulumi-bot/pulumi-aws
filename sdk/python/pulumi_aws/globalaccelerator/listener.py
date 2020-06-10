# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Listener(pulumi.CustomResource):
    accelerator_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of your accelerator.
    """
    client_affinity: pulumi.Output[str]
    """
    Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
    """
    port_ranges: pulumi.Output[list]
    """
    The list of port ranges for the connections from clients to the accelerator. Fields documented below.

      * `from_port` (`float`) - The first port in the range of ports, inclusive.
      * `to_port` (`float`) - The last port in the range of ports, inclusive.
    """
    protocol: pulumi.Output[str]
    """
    The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
    """
    def __init__(__self__, resource_name, opts=None, accelerator_arn=None, client_affinity=None, port_ranges=None, protocol=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Global Accelerator listener.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        example_accelerator = aws.globalaccelerator.Accelerator("exampleAccelerator",
            attributes={
                "flowLogsEnabled": True,
                "flowLogsS3Bucket": "example-bucket",
                "flowLogsS3Prefix": "flow-logs/",
            },
            enabled=True,
            ip_address_type="IPV4")
        example_listener = aws.globalaccelerator.Listener("exampleListener",
            accelerator_arn=example_accelerator.id,
            client_affinity="SOURCE_IP",
            port_ranges=[{
                "from_port": 80,
                "to_port": 80,
            }],
            protocol="TCP")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_arn: The Amazon Resource Name (ARN) of your accelerator.
        :param pulumi.Input[str] client_affinity: Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
        :param pulumi.Input[list] port_ranges: The list of port ranges for the connections from clients to the accelerator. Fields documented below.
        :param pulumi.Input[str] protocol: The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.

        The **port_ranges** object supports the following:

          * `from_port` (`pulumi.Input[float]`) - The first port in the range of ports, inclusive.
          * `to_port` (`pulumi.Input[float]`) - The last port in the range of ports, inclusive.
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

            if accelerator_arn is None:
                raise TypeError("Missing required property 'accelerator_arn'")
            __props__['accelerator_arn'] = accelerator_arn
            __props__['client_affinity'] = client_affinity
            if port_ranges is None:
                raise TypeError("Missing required property 'port_ranges'")
            __props__['port_ranges'] = port_ranges
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
        super(Listener, __self__).__init__(
            'aws:globalaccelerator/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accelerator_arn=None, client_affinity=None, port_ranges=None, protocol=None):
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_arn: The Amazon Resource Name (ARN) of your accelerator.
        :param pulumi.Input[str] client_affinity: Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
        :param pulumi.Input[list] port_ranges: The list of port ranges for the connections from clients to the accelerator. Fields documented below.
        :param pulumi.Input[str] protocol: The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.

        The **port_ranges** object supports the following:

          * `from_port` (`pulumi.Input[float]`) - The first port in the range of ports, inclusive.
          * `to_port` (`pulumi.Input[float]`) - The last port in the range of ports, inclusive.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accelerator_arn"] = accelerator_arn
        __props__["client_affinity"] = client_affinity
        __props__["port_ranges"] = port_ranges
        __props__["protocol"] = protocol
        return Listener(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

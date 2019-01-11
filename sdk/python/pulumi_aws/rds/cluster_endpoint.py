# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ClusterEndpoint(pulumi.CustomResource):
    """
    Manages a RDS Aurora Cluster Endpoint.
    You can refer to the [User Guide][1].
    
    """
    def __init__(__self__, __name__, __opts__=None, cluster_endpoint_identifier=None, cluster_identifier=None, custom_endpoint_type=None, excluded_members=None, static_members=None):
        """Create a ClusterEndpoint resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not cluster_endpoint_identifier:
            raise TypeError('Missing required property cluster_endpoint_identifier')
        __props__['cluster_endpoint_identifier'] = cluster_endpoint_identifier

        if not cluster_identifier:
            raise TypeError('Missing required property cluster_identifier')
        __props__['cluster_identifier'] = cluster_identifier

        if not custom_endpoint_type:
            raise TypeError('Missing required property custom_endpoint_type')
        __props__['custom_endpoint_type'] = custom_endpoint_type

        __props__['excluded_members'] = excluded_members

        __props__['static_members'] = static_members

        __props__['arn'] = None
        __props__['endpoint'] = None

        super(ClusterEndpoint, __self__).__init__(
            'aws:rds/clusterEndpoint:ClusterEndpoint',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


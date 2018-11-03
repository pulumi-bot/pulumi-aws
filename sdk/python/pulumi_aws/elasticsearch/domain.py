# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Domain(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, access_policies=None, advanced_options=None, cluster_config=None, cognito_options=None, domain_name=None, ebs_options=None, elasticsearch_version=None, encrypt_at_rest=None, log_publishing_options=None, node_to_node_encryption=None, snapshot_options=None, tags=None, vpc_options=None):
        """Create a Domain resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['accessPolicies'] = access_policies

        __props__['advancedOptions'] = advanced_options

        __props__['clusterConfig'] = cluster_config

        __props__['cognitoOptions'] = cognito_options

        __props__['domainName'] = domain_name

        __props__['ebsOptions'] = ebs_options

        __props__['elasticsearchVersion'] = elasticsearch_version

        __props__['encryptAtRest'] = encrypt_at_rest

        __props__['logPublishingOptions'] = log_publishing_options

        __props__['nodeToNodeEncryption'] = node_to_node_encryption

        __props__['snapshotOptions'] = snapshot_options

        __props__['tags'] = tags

        __props__['vpcOptions'] = vpc_options

        __props__['arn'] = None
        __props__['domain_id'] = None
        __props__['endpoint'] = None
        __props__['kibana_endpoint'] = None

        super(Domain, __self__).__init__(
            'aws:elasticsearch/domain:Domain',
            __name__,
            __props__,
            __opts__)


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDomainResult:
    """
    A collection of values returned by getDomain.
    """
    def __init__(__self__, access_policies=None, advanced_options=None, arn=None, cluster_configs=None, cognito_options=None, created=None, deleted=None, domain_id=None, domain_name=None, ebs_options=None, elasticsearch_version=None, encryption_at_rests=None, endpoint=None, id=None, kibana_endpoint=None, log_publishing_options=None, node_to_node_encryptions=None, processing=None, snapshot_options=None, tags=None, vpc_options=None):
        if access_policies and not isinstance(access_policies, str):
            raise TypeError("Expected argument 'access_policies' to be a str")
        __self__.access_policies = access_policies
        """
        The policy document attached to the domain.
        """
        if advanced_options and not isinstance(advanced_options, dict):
            raise TypeError("Expected argument 'advanced_options' to be a dict")
        __self__.advanced_options = advanced_options
        """
        Key-value string pairs to specify advanced configuration options.
        """
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the domain.
        """
        if cluster_configs and not isinstance(cluster_configs, list):
            raise TypeError("Expected argument 'cluster_configs' to be a list")
        __self__.cluster_configs = cluster_configs
        """
        Cluster configuration of the domain.
        """
        if cognito_options and not isinstance(cognito_options, list):
            raise TypeError("Expected argument 'cognito_options' to be a list")
        __self__.cognito_options = cognito_options
        """
        Domain Amazon Cognito Authentication options for Kibana.
        """
        if created and not isinstance(created, bool):
            raise TypeError("Expected argument 'created' to be a bool")
        __self__.created = created
        """
        Status of the creation of the domain.
        """
        if deleted and not isinstance(deleted, bool):
            raise TypeError("Expected argument 'deleted' to be a bool")
        __self__.deleted = deleted
        """
        Status of the deletion of the domain.
        """
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        __self__.domain_id = domain_id
        """
        Unique identifier for the domain.
        """
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        __self__.domain_name = domain_name
        if ebs_options and not isinstance(ebs_options, list):
            raise TypeError("Expected argument 'ebs_options' to be a list")
        __self__.ebs_options = ebs_options
        """
        EBS Options for the instances in the domain.
        """
        if elasticsearch_version and not isinstance(elasticsearch_version, str):
            raise TypeError("Expected argument 'elasticsearch_version' to be a str")
        __self__.elasticsearch_version = elasticsearch_version
        """
        ElasticSearch version for the domain.
        """
        if encryption_at_rests and not isinstance(encryption_at_rests, list):
            raise TypeError("Expected argument 'encryption_at_rests' to be a list")
        __self__.encryption_at_rests = encryption_at_rests
        """
        Domain encryption at rest related options.
        """
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        __self__.endpoint = endpoint
        """
        Domain-specific endpoint used to submit index, search, and data upload requests.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if kibana_endpoint and not isinstance(kibana_endpoint, str):
            raise TypeError("Expected argument 'kibana_endpoint' to be a str")
        __self__.kibana_endpoint = kibana_endpoint
        """
        Domain-specific endpoint used to access the Kibana application.
        """
        if log_publishing_options and not isinstance(log_publishing_options, list):
            raise TypeError("Expected argument 'log_publishing_options' to be a list")
        __self__.log_publishing_options = log_publishing_options
        """
        Domain log publishing related options.
        """
        if node_to_node_encryptions and not isinstance(node_to_node_encryptions, list):
            raise TypeError("Expected argument 'node_to_node_encryptions' to be a list")
        __self__.node_to_node_encryptions = node_to_node_encryptions
        """
        Domain in transit encryption related options.
        """
        if processing and not isinstance(processing, str):
            raise TypeError("Expected argument 'processing' to be a str")
        __self__.processing = processing
        """
        Status of a configuration change in the domain.
        * `snapshot_options` – Domain snapshot related options.
        """
        if snapshot_options and not isinstance(snapshot_options, list):
            raise TypeError("Expected argument 'snapshot_options' to be a list")
        __self__.snapshot_options = snapshot_options
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        The tags assigned to the domain.
        """
        if vpc_options and not isinstance(vpc_options, list):
            raise TypeError("Expected argument 'vpc_options' to be a list")
        __self__.vpc_options = vpc_options
        """
        VPC Options for private Elasticsearch domains.
        """
class AwaitableGetDomainResult(GetDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainResult(
            access_policies=self.access_policies,
            advanced_options=self.advanced_options,
            arn=self.arn,
            cluster_configs=self.cluster_configs,
            cognito_options=self.cognito_options,
            created=self.created,
            deleted=self.deleted,
            domain_id=self.domain_id,
            domain_name=self.domain_name,
            ebs_options=self.ebs_options,
            elasticsearch_version=self.elasticsearch_version,
            encryption_at_rests=self.encryption_at_rests,
            endpoint=self.endpoint,
            id=self.id,
            kibana_endpoint=self.kibana_endpoint,
            log_publishing_options=self.log_publishing_options,
            node_to_node_encryptions=self.node_to_node_encryptions,
            processing=self.processing,
            snapshot_options=self.snapshot_options,
            tags=self.tags,
            vpc_options=self.vpc_options)

def get_domain(domain_name=None,tags=None,opts=None):
    """
    Use this data source to get information about an Elasticsearch Domain




    :param str domain_name: Name of the domain.
    :param dict tags: The tags assigned to the domain.
    """
    __args__ = dict()


    __args__['domainName'] = domain_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticsearch/getDomain:getDomain', __args__, opts=opts).value

    return AwaitableGetDomainResult(
        access_policies=__ret__.get('accessPolicies'),
        advanced_options=__ret__.get('advancedOptions'),
        arn=__ret__.get('arn'),
        cluster_configs=__ret__.get('clusterConfigs'),
        cognito_options=__ret__.get('cognitoOptions'),
        created=__ret__.get('created'),
        deleted=__ret__.get('deleted'),
        domain_id=__ret__.get('domainId'),
        domain_name=__ret__.get('domainName'),
        ebs_options=__ret__.get('ebsOptions'),
        elasticsearch_version=__ret__.get('elasticsearchVersion'),
        encryption_at_rests=__ret__.get('encryptionAtRests'),
        endpoint=__ret__.get('endpoint'),
        id=__ret__.get('id'),
        kibana_endpoint=__ret__.get('kibanaEndpoint'),
        log_publishing_options=__ret__.get('logPublishingOptions'),
        node_to_node_encryptions=__ret__.get('nodeToNodeEncryptions'),
        processing=__ret__.get('processing'),
        snapshot_options=__ret__.get('snapshotOptions'),
        tags=__ret__.get('tags'),
        vpc_options=__ret__.get('vpcOptions'))

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, cluster_certificates=None, cluster_id=None, cluster_state=None, id=None, security_group_id=None, subnet_ids=None, vpc_id=None):
        if cluster_certificates and not isinstance(cluster_certificates, dict):
            raise TypeError("Expected argument 'cluster_certificates' to be a dict")
        __self__.cluster_certificates = cluster_certificates
        """
        The list of cluster certificates.
        * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
        * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state.
        * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
        * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
        * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
        The number of available cluster certificates may vary depending on state of the cluster.
        """
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        __self__.cluster_id = cluster_id
        if cluster_state and not isinstance(cluster_state, str):
            raise TypeError("Expected argument 'cluster_state' to be a str")
        __self__.cluster_state = cluster_state
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        __self__.security_group_id = security_group_id
        """
        The ID of the security group associated with the CloudHSM cluster.
        """
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        __self__.subnet_ids = subnet_ids
        """
        The IDs of subnets in which cluster operates.
        """
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        The id of the VPC that the CloudHSM cluster resides in.
        """
class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            cluster_certificates=self.cluster_certificates,
            cluster_id=self.cluster_id,
            cluster_state=self.cluster_state,
            id=self.id,
            security_group_id=self.security_group_id,
            subnet_ids=self.subnet_ids,
            vpc_id=self.vpc_id)

def get_cluster(cluster_id=None,cluster_state=None,opts=None):
    """
    Use this data source to get information about a CloudHSM v2 cluster

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudhsm_v2_cluster.html.markdown.


    :param str cluster_id: The id of Cloud HSM v2 cluster.
    :param str cluster_state: The state of the cluster to be found.
    """
    __args__ = dict()


    __args__['clusterId'] = cluster_id
    __args__['clusterState'] = cluster_state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:cloudhsmv2/getCluster:getCluster', __args__, opts=opts).value

    return AwaitableGetClusterResult(
        cluster_certificates=__ret__.get('clusterCertificates'),
        cluster_id=__ret__.get('clusterId'),
        cluster_state=__ret__.get('clusterState'),
        id=__ret__.get('id'),
        security_group_id=__ret__.get('securityGroupId'),
        subnet_ids=__ret__.get('subnetIds'),
        vpc_id=__ret__.get('vpcId'))

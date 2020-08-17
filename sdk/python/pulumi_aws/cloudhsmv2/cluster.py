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

__all__ = ['Cluster']


class Cluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hsm_type: Optional[pulumi.Input[str]] = None,
                 source_backup_identifier: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Amazon CloudHSM v2 cluster.

        For information about CloudHSM v2, see the
        [AWS CloudHSM User Guide](https://docs.aws.amazon.com/cloudhsm/latest/userguide/introduction.html) and the [Amazon
        CloudHSM API Reference][2].

        > **NOTE:** A CloudHSM Cluster can take several minutes to set up.
        Practically no single attribute can be updated, except for `tags`.
        If you need to delete a cluster, you have to remove its HSM modules first.
        To initialize cluster, you have to add an HSM instance to the cluster, then sign CSR and upload it.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hsm_type: The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
        :param pulumi.Input[str] source_backup_identifier: The id of Cloud HSM v2 cluster backup to be restored.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: The IDs of subnets in which cluster will operate.
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

            if hsm_type is None:
                raise TypeError("Missing required property 'hsm_type'")
            __props__['hsm_type'] = hsm_type
            __props__['source_backup_identifier'] = source_backup_identifier
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['cluster_certificates'] = None
            __props__['cluster_id'] = None
            __props__['cluster_state'] = None
            __props__['security_group_id'] = None
            __props__['vpc_id'] = None
        super(Cluster, __self__).__init__(
            'aws:cloudhsmv2/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ClusterClusterCertificateArgs']]]]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_state: Optional[pulumi.Input[str]] = None,
            hsm_type: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            source_backup_identifier: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ClusterClusterCertificateArgs']]]] cluster_certificates: The list of cluster certificates.
               * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
               * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in `UNINITIALIZED` state after an HSM instance is added to the cluster.
               * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
               * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
               * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
        :param pulumi.Input[str] cluster_id: The id of the CloudHSM cluster.
        :param pulumi.Input[str] cluster_state: The state of the CloudHSM cluster.
        :param pulumi.Input[str] hsm_type: The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
        :param pulumi.Input[str] security_group_id: The ID of the security group associated with the CloudHSM cluster.
        :param pulumi.Input[str] source_backup_identifier: The id of Cloud HSM v2 cluster backup to be restored.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: The IDs of subnets in which cluster will operate.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC that the CloudHSM cluster resides in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_certificates"] = cluster_certificates
        __props__["cluster_id"] = cluster_id
        __props__["cluster_state"] = cluster_state
        __props__["hsm_type"] = hsm_type
        __props__["security_group_id"] = security_group_id
        __props__["source_backup_identifier"] = source_backup_identifier
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterCertificates")
    def cluster_certificates(self) -> List['outputs.ClusterClusterCertificate']:
        """
        The list of cluster certificates.
        * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
        * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in `UNINITIALIZED` state after an HSM instance is added to the cluster.
        * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
        * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
        * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
        """
        ...

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The id of the CloudHSM cluster.
        """
        ...

    @property
    @pulumi.getter(name="clusterState")
    def cluster_state(self) -> str:
        """
        The state of the CloudHSM cluster.
        """
        ...

    @property
    @pulumi.getter(name="hsmType")
    def hsm_type(self) -> str:
        """
        The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
        """
        ...

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        """
        The ID of the security group associated with the CloudHSM cluster.
        """
        ...

    @property
    @pulumi.getter(name="sourceBackupIdentifier")
    def source_backup_identifier(self) -> Optional[str]:
        """
        The id of Cloud HSM v2 cluster backup to be restored.
        """
        ...

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        """
        The IDs of subnets in which cluster will operate.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource.
        """
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The id of the VPC that the CloudHSM cluster resides in.
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


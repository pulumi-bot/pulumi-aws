# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    cluster_certificates: pulumi.Output[dict]
    """
    The list of cluster certificates.
    * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
    * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state after an hsm instance is added to the cluster.
    * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
    * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
    * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
    """
    cluster_id: pulumi.Output[str]
    """
    The id of the CloudHSM cluster.
    """
    cluster_state: pulumi.Output[str]
    """
    The state of the cluster.
    """
    hsm_type: pulumi.Output[str]
    """
    The type of HSM module in the cluster. Currently, only hsm1.medium is supported.
    """
    security_group_id: pulumi.Output[str]
    """
    The ID of the security group associated with the CloudHSM cluster.
    """
    source_backup_identifier: pulumi.Output[str]
    """
    The id of Cloud HSM v2 cluster backup to be restored.
    """
    subnet_ids: pulumi.Output[list]
    """
    The IDs of subnets in which cluster will operate.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The id of the VPC that the CloudHSM cluster resides in.
    """
    def __init__(__self__, resource_name, opts=None, hsm_type=None, source_backup_identifier=None, subnet_ids=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Amazon CloudHSM v2 cluster.
        
        For information about CloudHSM v2, see the
        [AWS CloudHSM User Guide][1] and the [Amazon
        CloudHSM API Reference][2].
        
        > **NOTE:** CloudHSM can take up to several minutes to be set up.
        Practically no single attribute can be updated except TAGS.
        If you need to delete a cluster, you have to remove its HSM modules first.
        To initialize cluster, you have to add an hsm instance to the cluster then sign CSR and upload it.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hsm_type: The type of HSM module in the cluster. Currently, only hsm1.medium is supported.
        :param pulumi.Input[str] source_backup_identifier: The id of Cloud HSM v2 cluster backup to be restored.
        :param pulumi.Input[list] subnet_ids: The IDs of subnets in which cluster will operate.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudhsm_v2_cluster.html.markdown.
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
    def get(resource_name, id, opts=None, cluster_certificates=None, cluster_id=None, cluster_state=None, hsm_type=None, security_group_id=None, source_backup_identifier=None, subnet_ids=None, tags=None, vpc_id=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] cluster_certificates: The list of cluster certificates.
               * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
               * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state after an hsm instance is added to the cluster.
               * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
               * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
               * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
        :param pulumi.Input[str] cluster_id: The id of the CloudHSM cluster.
        :param pulumi.Input[str] cluster_state: The state of the cluster.
        :param pulumi.Input[str] hsm_type: The type of HSM module in the cluster. Currently, only hsm1.medium is supported.
        :param pulumi.Input[str] security_group_id: The ID of the security group associated with the CloudHSM cluster.
        :param pulumi.Input[str] source_backup_identifier: The id of Cloud HSM v2 cluster backup to be restored.
        :param pulumi.Input[list] subnet_ids: The IDs of subnets in which cluster will operate.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC that the CloudHSM cluster resides in.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudhsm_v2_cluster.html.markdown.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


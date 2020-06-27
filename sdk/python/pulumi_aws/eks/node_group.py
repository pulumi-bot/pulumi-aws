# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class NodeGroup(pulumi.CustomResource):
    ami_type: pulumi.Output[str]
    """
    Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`. This provider will only perform drift detection if a configuration value is provided.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the EKS Node Group.
    """
    cluster_name: pulumi.Output[str]
    """
    Name of the EKS Cluster.
    """
    disk_size: pulumi.Output[float]
    """
    Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
    """
    force_update_version: pulumi.Output[bool]
    """
    Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
    """
    instance_types: pulumi.Output[str]
    """
    Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided. Currently, the EKS API only accepts a single value in the set.
    """
    labels: pulumi.Output[dict]
    """
    Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
    """
    node_group_name: pulumi.Output[str]
    """
    Name of the EKS Node Group.
    """
    node_role_arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
    """
    release_version: pulumi.Output[str]
    """
    AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
    """
    remote_access: pulumi.Output[dict]
    """
    Configuration block with remote access settings. Detailed below.

      * `ec2SshKey` (`str`) - EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `source_security_group_ids` when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
      * `sourceSecurityGroupIds` (`list`) - Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2_ssh_key`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
    """
    resources: pulumi.Output[list]
    """
    List of objects containing information about underlying resources.

      * `autoscaling_groups` (`list`) - List of objects containing information about AutoScaling Groups.
        * `name` (`str`) - Name of the AutoScaling Group.

      * `remoteAccessSecurityGroupId` (`str`) - Identifier of the remote access EC2 Security Group.
    """
    scaling_config: pulumi.Output[dict]
    """
    Configuration block with scaling settings. Detailed below.

      * `desiredSize` (`float`) - Desired number of worker nodes.
      * `max_size` (`float`) - Maximum number of worker nodes.
      * `min_size` (`float`) - Minimum number of worker nodes.
    """
    status: pulumi.Output[str]
    """
    Status of the EKS Node Group.
    """
    subnet_ids: pulumi.Output[list]
    """
    Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags.
    """
    version: pulumi.Output[str]
    """
    Kubernetes version. Defaults to EKS Cluster Kubernetes version. This provider will only perform drift detection if a configuration value is provided.
    """
    def __init__(__self__, resource_name, opts=None, ami_type=None, cluster_name=None, disk_size=None, force_update_version=None, instance_types=None, labels=None, node_group_name=None, node_role_arn=None, release_version=None, remote_access=None, scaling_config=None, subnet_ids=None, tags=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.eks.NodeGroup("example",
            cluster_name=aws_eks_cluster["example"]["name"],
            node_role_arn=aws_iam_role["example"]["arn"],
            subnet_ids=[__item["id"] for __item in aws_subnet["example"]],
            scaling_config={
                "desiredSize": 1,
                "max_size": 1,
                "min_size": 1,
            })
        ```
        ### Ignoring Changes to Desired Size

        You can utilize [ignoreChanges](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) create an EKS Node Group with an initial size of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configurations ...
        example = aws.eks.NodeGroup("example", scaling_config={
            "desiredSize": 2,
        })
        ```
        ### Example IAM Role for EKS Node Group

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example = aws.iam.Role("example", assume_role_policy=json.dumps({
            "Statement": [{
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Principal": {
                    "Service": "ec2.amazonaws.com",
                },
            }],
            "Version": "2012-10-17",
        }))
        example__amazon_eks_worker_node_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy",
            policy_arn="arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
            role=example.name)
        example__amazon_ekscni_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy",
            policy_arn="arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
            role=example.name)
        example__amazon_ec2_container_registry_read_only = aws.iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly",
            policy_arn="arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
            role=example.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ami_type: Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`. This provider will only perform drift detection if a configuration value is provided.
        :param pulumi.Input[str] cluster_name: Name of the EKS Cluster.
        :param pulumi.Input[float] disk_size: Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        :param pulumi.Input[bool] force_update_version: Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        :param pulumi.Input[str] instance_types: Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided. Currently, the EKS API only accepts a single value in the set.
        :param pulumi.Input[dict] labels: Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        :param pulumi.Input[str] node_group_name: Name of the EKS Node Group.
        :param pulumi.Input[str] node_role_arn: Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        :param pulumi.Input[str] release_version: AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        :param pulumi.Input[dict] remote_access: Configuration block with remote access settings. Detailed below.
        :param pulumi.Input[dict] scaling_config: Configuration block with scaling settings. Detailed below.
        :param pulumi.Input[list] subnet_ids: Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags.
        :param pulumi.Input[str] version: Kubernetes version. Defaults to EKS Cluster Kubernetes version. This provider will only perform drift detection if a configuration value is provided.

        The **remote_access** object supports the following:

          * `ec2SshKey` (`pulumi.Input[str]`) - EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `source_security_group_ids` when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
          * `sourceSecurityGroupIds` (`pulumi.Input[list]`) - Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2_ssh_key`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).

        The **scaling_config** object supports the following:

          * `desiredSize` (`pulumi.Input[float]`) - Desired number of worker nodes.
          * `max_size` (`pulumi.Input[float]`) - Maximum number of worker nodes.
          * `min_size` (`pulumi.Input[float]`) - Minimum number of worker nodes.
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

            __props__['ami_type'] = ami_type
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['disk_size'] = disk_size
            __props__['force_update_version'] = force_update_version
            __props__['instance_types'] = instance_types
            __props__['labels'] = labels
            __props__['node_group_name'] = node_group_name
            if node_role_arn is None:
                raise TypeError("Missing required property 'node_role_arn'")
            __props__['node_role_arn'] = node_role_arn
            __props__['release_version'] = release_version
            __props__['remote_access'] = remote_access
            if scaling_config is None:
                raise TypeError("Missing required property 'scaling_config'")
            __props__['scaling_config'] = scaling_config
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['version'] = version
            __props__['arn'] = None
            __props__['resources'] = None
            __props__['status'] = None
        super(NodeGroup, __self__).__init__(
            'aws:eks/nodeGroup:NodeGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ami_type=None, arn=None, cluster_name=None, disk_size=None, force_update_version=None, instance_types=None, labels=None, node_group_name=None, node_role_arn=None, release_version=None, remote_access=None, resources=None, scaling_config=None, status=None, subnet_ids=None, tags=None, version=None):
        """
        Get an existing NodeGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ami_type: Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`. This provider will only perform drift detection if a configuration value is provided.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the EKS Node Group.
        :param pulumi.Input[str] cluster_name: Name of the EKS Cluster.
        :param pulumi.Input[float] disk_size: Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        :param pulumi.Input[bool] force_update_version: Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        :param pulumi.Input[str] instance_types: Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided. Currently, the EKS API only accepts a single value in the set.
        :param pulumi.Input[dict] labels: Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        :param pulumi.Input[str] node_group_name: Name of the EKS Node Group.
        :param pulumi.Input[str] node_role_arn: Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        :param pulumi.Input[str] release_version: AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        :param pulumi.Input[dict] remote_access: Configuration block with remote access settings. Detailed below.
        :param pulumi.Input[list] resources: List of objects containing information about underlying resources.
        :param pulumi.Input[dict] scaling_config: Configuration block with scaling settings. Detailed below.
        :param pulumi.Input[str] status: Status of the EKS Node Group.
        :param pulumi.Input[list] subnet_ids: Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags.
        :param pulumi.Input[str] version: Kubernetes version. Defaults to EKS Cluster Kubernetes version. This provider will only perform drift detection if a configuration value is provided.

        The **remote_access** object supports the following:

          * `ec2SshKey` (`pulumi.Input[str]`) - EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `source_security_group_ids` when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
          * `sourceSecurityGroupIds` (`pulumi.Input[list]`) - Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2_ssh_key`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).

        The **resources** object supports the following:

          * `autoscaling_groups` (`pulumi.Input[list]`) - List of objects containing information about AutoScaling Groups.
            * `name` (`pulumi.Input[str]`) - Name of the AutoScaling Group.

          * `remoteAccessSecurityGroupId` (`pulumi.Input[str]`) - Identifier of the remote access EC2 Security Group.

        The **scaling_config** object supports the following:

          * `desiredSize` (`pulumi.Input[float]`) - Desired number of worker nodes.
          * `max_size` (`pulumi.Input[float]`) - Maximum number of worker nodes.
          * `min_size` (`pulumi.Input[float]`) - Minimum number of worker nodes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ami_type"] = ami_type
        __props__["arn"] = arn
        __props__["cluster_name"] = cluster_name
        __props__["disk_size"] = disk_size
        __props__["force_update_version"] = force_update_version
        __props__["instance_types"] = instance_types
        __props__["labels"] = labels
        __props__["node_group_name"] = node_group_name
        __props__["node_role_arn"] = node_role_arn
        __props__["release_version"] = release_version
        __props__["remote_access"] = remote_access
        __props__["resources"] = resources
        __props__["scaling_config"] = scaling_config
        __props__["status"] = status
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["version"] = version
        return NodeGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

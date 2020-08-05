# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Cluster(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the cluster.
    """
    certificate_authority: pulumi.Output[dict]
    """
    Nested attribute containing `certificate-authority-data` for your cluster.

      * `data` (`str`) - The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
    """
    created_at: pulumi.Output[str]
    enabled_cluster_log_types: pulumi.Output[list]
    """
    A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
    """
    encryption_config: pulumi.Output[dict]
    """
    Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.

      * `provider` (`dict`) - Configuration block with provider for encryption. Detailed below.
        * `key_arn` (`str`) - Amazon Resource Name (ARN) of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see [Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html).

      * `resources` (`list`) - List of strings with resources to be encrypted. Valid values: `secrets`
    """
    endpoint: pulumi.Output[str]
    """
    The endpoint for your Kubernetes API server.
    """
    identities: pulumi.Output[list]
    """
    Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.

      * `oidcs` (`list`) - Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster.
        * `issuer` (`str`) - Issuer URL for the OpenID Connect identity provider.
    """
    name: pulumi.Output[str]
    """
    Name of the cluster.
    """
    platform_version: pulumi.Output[str]
    """
    The platform version for the cluster.
    """
    role_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
    """
    status: pulumi.Output[str]
    """
    The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of resource tags.
    """
    version: pulumi.Output[str]
    """
    Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
    """
    vpc_config: pulumi.Output[dict]
    """
    Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.

      * `clusterSecurityGroupId` (`str`) - The cluster security group that was created by Amazon EKS for the cluster.
      * `endpointPrivateAccess` (`bool`) - Indicates whether or not the Amazon EKS private API server endpoint is enabled. Default is `false`.
      * `endpointPublicAccess` (`bool`) - Indicates whether or not the Amazon EKS public API server endpoint is enabled. Default is `true`.
      * `publicAccessCidrs` (`list`) - List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. This provider will only perform drift detection of its value when present in a configuration.
      * `security_group_ids` (`list`) - List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
      * `subnet_ids` (`list`) - List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
      * `vpc_id` (`str`) - The VPC associated with your cluster.
    """
    def __init__(__self__, resource_name, opts=None, enabled_cluster_log_types=None, encryption_config=None, name=None, role_arn=None, tags=None, version=None, vpc_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an EKS Cluster.

        ## Example Usage
        ### Example IAM Role for EKS Cluster

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.iam.Role("example", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": "eks.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }

        \"\"\")
        example__amazon_eks_cluster_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy",
            policy_arn="arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
            role=example.name)
        example__amazon_eks_service_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSServicePolicy",
            policy_arn="arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
            role=example.name)
        ```
        ### Enabling Control Plane Logging

        [EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabled_cluster_log_types` argument. To manage the CloudWatch Log Group retention period, the `cloudwatch.LogGroup` resource can be used.

        > The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.

        ```python
        import pulumi
        import pulumi_aws as aws

        config = pulumi.Config()
        cluster_name = config.get("clusterName")
        if cluster_name is None:
            cluster_name = "example"
        example_cluster = aws.eks.Cluster("exampleCluster", enabled_cluster_log_types=[
            "api",
            "audit",
        ],
        opts=ResourceOptions(depends_on=["aws_cloudwatch_log_group.example"]))
        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=7)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] enabled_cluster_log_types: A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
        :param pulumi.Input[dict] encryption_config: Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
        :param pulumi.Input[str] name: Name of the cluster.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
        :param pulumi.Input[dict] tags: Key-value map of resource tags.
        :param pulumi.Input[str] version: Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
        :param pulumi.Input[dict] vpc_config: Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.

        The **encryption_config** object supports the following:

          * `provider` (`pulumi.Input[dict]`) - Configuration block with provider for encryption. Detailed below.
            * `key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see [Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html).

          * `resources` (`pulumi.Input[list]`) - List of strings with resources to be encrypted. Valid values: `secrets`

        The **vpc_config** object supports the following:

          * `clusterSecurityGroupId` (`pulumi.Input[str]`) - The cluster security group that was created by Amazon EKS for the cluster.
          * `endpointPrivateAccess` (`pulumi.Input[bool]`) - Indicates whether or not the Amazon EKS private API server endpoint is enabled. Default is `false`.
          * `endpointPublicAccess` (`pulumi.Input[bool]`) - Indicates whether or not the Amazon EKS public API server endpoint is enabled. Default is `true`.
          * `publicAccessCidrs` (`pulumi.Input[list]`) - List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. This provider will only perform drift detection of its value when present in a configuration.
          * `security_group_ids` (`pulumi.Input[list]`) - List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
          * `subnet_ids` (`pulumi.Input[list]`) - List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
          * `vpc_id` (`pulumi.Input[str]`) - The VPC associated with your cluster.
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

            __props__['enabled_cluster_log_types'] = enabled_cluster_log_types
            __props__['encryption_config'] = encryption_config
            __props__['name'] = name
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['tags'] = tags
            __props__['version'] = version
            if vpc_config is None:
                raise TypeError("Missing required property 'vpc_config'")
            __props__['vpc_config'] = vpc_config
            __props__['arn'] = None
            __props__['certificate_authority'] = None
            __props__['created_at'] = None
            __props__['endpoint'] = None
            __props__['identities'] = None
            __props__['platform_version'] = None
            __props__['status'] = None
        super(Cluster, __self__).__init__(
            'aws:eks/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, certificate_authority=None, created_at=None, enabled_cluster_log_types=None, encryption_config=None, endpoint=None, identities=None, name=None, platform_version=None, role_arn=None, status=None, tags=None, version=None, vpc_config=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the cluster.
        :param pulumi.Input[dict] certificate_authority: Nested attribute containing `certificate-authority-data` for your cluster.
        :param pulumi.Input[list] enabled_cluster_log_types: A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
        :param pulumi.Input[dict] encryption_config: Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
        :param pulumi.Input[str] endpoint: The endpoint for your Kubernetes API server.
        :param pulumi.Input[list] identities: Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
        :param pulumi.Input[str] name: Name of the cluster.
        :param pulumi.Input[str] platform_version: The platform version for the cluster.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
        :param pulumi.Input[str] status: The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
        :param pulumi.Input[dict] tags: Key-value map of resource tags.
        :param pulumi.Input[str] version: Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
        :param pulumi.Input[dict] vpc_config: Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.

        The **certificate_authority** object supports the following:

          * `data` (`pulumi.Input[str]`) - The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.

        The **encryption_config** object supports the following:

          * `provider` (`pulumi.Input[dict]`) - Configuration block with provider for encryption. Detailed below.
            * `key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see [Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html).

          * `resources` (`pulumi.Input[list]`) - List of strings with resources to be encrypted. Valid values: `secrets`

        The **identities** object supports the following:

          * `oidcs` (`pulumi.Input[list]`) - Nested attribute containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster.
            * `issuer` (`pulumi.Input[str]`) - Issuer URL for the OpenID Connect identity provider.

        The **vpc_config** object supports the following:

          * `clusterSecurityGroupId` (`pulumi.Input[str]`) - The cluster security group that was created by Amazon EKS for the cluster.
          * `endpointPrivateAccess` (`pulumi.Input[bool]`) - Indicates whether or not the Amazon EKS private API server endpoint is enabled. Default is `false`.
          * `endpointPublicAccess` (`pulumi.Input[bool]`) - Indicates whether or not the Amazon EKS public API server endpoint is enabled. Default is `true`.
          * `publicAccessCidrs` (`pulumi.Input[list]`) - List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. This provider will only perform drift detection of its value when present in a configuration.
          * `security_group_ids` (`pulumi.Input[list]`) - List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
          * `subnet_ids` (`pulumi.Input[list]`) - List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
          * `vpc_id` (`pulumi.Input[str]`) - The VPC associated with your cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["certificate_authority"] = certificate_authority
        __props__["created_at"] = created_at
        __props__["enabled_cluster_log_types"] = enabled_cluster_log_types
        __props__["encryption_config"] = encryption_config
        __props__["endpoint"] = endpoint
        __props__["identities"] = identities
        __props__["name"] = name
        __props__["platform_version"] = platform_version
        __props__["role_arn"] = role_arn
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["version"] = version
        __props__["vpc_config"] = vpc_config
        return Cluster(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EKS Cluster.
//
// ## Example Usage
// ### Example IAM Role for EKS Cluster
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"eks.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "example_AmazonEKSClusterPolicy", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
// 			Role:      example.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "example_AmazonEKSServicePolicy", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSServicePolicy"),
// 			Role:      example.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enabling Control Plane Logging
//
// [EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabledClusterLogTypes` argument. To manage the CloudWatch Log Group retention period, the `cloudwatch.LogGroup` resource can be used.
//
// > The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/eks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = eks.NewCluster(ctx, "exampleCluster", &eks.ClusterArgs{
// 			EnabledClusterLogTypes: pulumi.StringArray{
// 				pulumi.String("api"),
// 				pulumi.String("audit"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
// 			RetentionInDays: pulumi.Int(7),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Nested attribute containing `certificate-authority-data` for your cluster.
	CertificateAuthority ClusterCertificateAuthorityOutput `pulumi:"certificateAuthority"`
	CreatedAt            pulumi.StringOutput               `pulumi:"createdAt"`
	// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	EnabledClusterLogTypes pulumi.StringArrayOutput `pulumi:"enabledClusterLogTypes"`
	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	EncryptionConfig ClusterEncryptionConfigPtrOutput `pulumi:"encryptionConfig"`
	// The endpoint for your Kubernetes API server.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
	Identities ClusterIdentityArrayOutput `pulumi:"identities"`
	// Name of the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The platform version for the cluster.
	PlatformVersion pulumi.StringOutput `pulumi:"platformVersion"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	Version pulumi.StringOutput `pulumi:"version"`
	// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
	VpcConfig ClusterVpcConfigOutput `pulumi:"vpcConfig"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil || args.VpcConfig == nil {
		return nil, errors.New("missing required argument 'VpcConfig'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:eks/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws:eks/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The Amazon Resource Name (ARN) of the cluster.
	Arn *string `pulumi:"arn"`
	// Nested attribute containing `certificate-authority-data` for your cluster.
	CertificateAuthority *ClusterCertificateAuthority `pulumi:"certificateAuthority"`
	CreatedAt            *string                      `pulumi:"createdAt"`
	// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	EnabledClusterLogTypes []string `pulumi:"enabledClusterLogTypes"`
	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	EncryptionConfig *ClusterEncryptionConfig `pulumi:"encryptionConfig"`
	// The endpoint for your Kubernetes API server.
	Endpoint *string `pulumi:"endpoint"`
	// Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
	Identities []ClusterIdentity `pulumi:"identities"`
	// Name of the cluster.
	Name *string `pulumi:"name"`
	// The platform version for the cluster.
	PlatformVersion *string `pulumi:"platformVersion"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	RoleArn *string `pulumi:"roleArn"`
	// The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
	Status *string `pulumi:"status"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	Version *string `pulumi:"version"`
	// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
	VpcConfig *ClusterVpcConfig `pulumi:"vpcConfig"`
}

type ClusterState struct {
	// The Amazon Resource Name (ARN) of the cluster.
	Arn pulumi.StringPtrInput
	// Nested attribute containing `certificate-authority-data` for your cluster.
	CertificateAuthority ClusterCertificateAuthorityPtrInput
	CreatedAt            pulumi.StringPtrInput
	// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	EnabledClusterLogTypes pulumi.StringArrayInput
	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	EncryptionConfig ClusterEncryptionConfigPtrInput
	// The endpoint for your Kubernetes API server.
	Endpoint pulumi.StringPtrInput
	// Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
	Identities ClusterIdentityArrayInput
	// Name of the cluster.
	Name pulumi.StringPtrInput
	// The platform version for the cluster.
	PlatformVersion pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	RoleArn pulumi.StringPtrInput
	// The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
	Status pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	Version pulumi.StringPtrInput
	// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
	VpcConfig ClusterVpcConfigPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	EnabledClusterLogTypes []string `pulumi:"enabledClusterLogTypes"`
	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	EncryptionConfig *ClusterEncryptionConfig `pulumi:"encryptionConfig"`
	// Name of the cluster.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	RoleArn string `pulumi:"roleArn"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	Version *string `pulumi:"version"`
	// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
	VpcConfig ClusterVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	EnabledClusterLogTypes pulumi.StringArrayInput
	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	EncryptionConfig ClusterEncryptionConfigPtrInput
	// Name of the cluster.
	Name pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `iam.RolePolicy` resource) or `iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	RoleArn pulumi.StringInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	Version pulumi.StringPtrInput
	// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
	VpcConfig ClusterVpcConfigInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

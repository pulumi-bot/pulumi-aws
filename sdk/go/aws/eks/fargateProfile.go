// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EKS Fargate Profile.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/eks_fargate_profile.html.markdown.
type FargateProfile struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the EKS Cluster.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Name of the EKS Fargate Profile.
	FargateProfileName pulumi.StringOutput `pulumi:"fargateProfileName"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn pulumi.StringOutput `pulumi:"podExecutionRoleArn"`
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors FargateProfileSelectorArrayOutput `pulumi:"selectors"`
	// Status of the EKS Fargate Profile.
	Status pulumi.StringOutput `pulumi:"status"`
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Key-value mapping of resource tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewFargateProfile registers a new resource with the given unique name, arguments, and options.
func NewFargateProfile(ctx *pulumi.Context,
	name string, args *FargateProfileArgs, opts ...pulumi.ResourceOption) (*FargateProfile, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.PodExecutionRoleArn == nil {
		return nil, errors.New("missing required argument 'PodExecutionRoleArn'")
	}
	if args == nil || args.Selectors == nil {
		return nil, errors.New("missing required argument 'Selectors'")
	}
	if args == nil {
		args = &FargateProfileArgs{}
	}
	var resource FargateProfile
	err := ctx.RegisterResource("aws:eks/fargateProfile:FargateProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFargateProfile gets an existing FargateProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFargateProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FargateProfileState, opts ...pulumi.ResourceOption) (*FargateProfile, error) {
	var resource FargateProfile
	err := ctx.ReadResource("aws:eks/fargateProfile:FargateProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FargateProfile resources.
type fargateProfileState struct {
	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn *string `pulumi:"arn"`
	// Name of the EKS Cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Name of the EKS Fargate Profile.
	FargateProfileName *string `pulumi:"fargateProfileName"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn *string `pulumi:"podExecutionRoleArn"`
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors []FargateProfileSelector `pulumi:"selectors"`
	// Status of the EKS Fargate Profile.
	Status *string `pulumi:"status"`
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

type FargateProfileState struct {
	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn pulumi.StringPtrInput
	// Name of the EKS Cluster.
	ClusterName pulumi.StringPtrInput
	// Name of the EKS Fargate Profile.
	FargateProfileName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn pulumi.StringPtrInput
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors FargateProfileSelectorArrayInput
	// Status of the EKS Fargate Profile.
	Status pulumi.StringPtrInput
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds pulumi.StringArrayInput
	// Key-value mapping of resource tags.
	Tags pulumi.MapInput
}

func (FargateProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateProfileState)(nil)).Elem()
}

type fargateProfileArgs struct {
	// Name of the EKS Cluster.
	ClusterName string `pulumi:"clusterName"`
	// Name of the EKS Fargate Profile.
	FargateProfileName *string `pulumi:"fargateProfileName"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn string `pulumi:"podExecutionRoleArn"`
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors []FargateProfileSelector `pulumi:"selectors"`
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a FargateProfile resource.
type FargateProfileArgs struct {
	// Name of the EKS Cluster.
	ClusterName pulumi.StringInput
	// Name of the EKS Fargate Profile.
	FargateProfileName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn pulumi.StringInput
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors FargateProfileSelectorArrayInput
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds pulumi.StringArrayInput
	// Key-value mapping of resource tags.
	Tags pulumi.MapInput
}

func (FargateProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateProfileArgs)(nil)).Elem()
}


// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sagemaker Notebook Instance resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
// 			RoleArn:      pulumi.Any(aws_iam_role.Role.Arn),
// 			InstanceType: pulumi.String("ml.t2.medium"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("foo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Code repository usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := sagemaker.NewCodeRepository(ctx, "example", &sagemaker.CodeRepositoryArgs{
// 			CodeRepositoryName: pulumi.String("my-notebook-instance-code-repo"),
// 			GitConfig: &sagemaker.CodeRepositoryGitConfigArgs{
// 				RepositoryUrl: pulumi.String("https://github.com/terraform-providers/terraform-provider-aws.git"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
// 			RoleArn:               pulumi.Any(aws_iam_role.Role.Arn),
// 			InstanceType:          pulumi.String("ml.t2.medium"),
// 			DefaultCodeRepository: example.CodeRepositoryName,
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("foo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NotebookInstance struct {
	pulumi.CustomResourceState

	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories pulumi.StringArrayOutput `pulumi:"additionalCodeRepositories"`
	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository pulumi.StringPtrOutput `pulumi:"defaultCodeRepository"`
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess pulumi.StringPtrOutput `pulumi:"directInternetAccess"`
	// The name of ML compute instance type.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName pulumi.StringPtrOutput `pulumi:"lifecycleConfigName"`
	// The name of the notebook instance (must be unique).
	Name pulumi.StringOutput `pulumi:"name"`
	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess pulumi.StringPtrOutput `pulumi:"rootAccess"`
	// The associated security groups.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The VPC subnet ID.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	Url pulumi.StringOutput `pulumi:"url"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize pulumi.IntPtrOutput `pulumi:"volumeSize"`
}

// NewNotebookInstance registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstance(ctx *pulumi.Context,
	name string, args *NotebookInstanceArgs, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource NotebookInstance
	err := ctx.RegisterResource("aws:sagemaker/notebookInstance:NotebookInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstance gets an existing NotebookInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceState, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	var resource NotebookInstance
	err := ctx.ReadResource("aws:sagemaker/notebookInstance:NotebookInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstance resources.
type notebookInstanceState struct {
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories []string `pulumi:"additionalCodeRepositories"`
	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn *string `pulumi:"arn"`
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository *string `pulumi:"defaultCodeRepository"`
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess *string `pulumi:"directInternetAccess"`
	// The name of ML compute instance type.
	InstanceType *string `pulumi:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName *string `pulumi:"lifecycleConfigName"`
	// The name of the notebook instance (must be unique).
	Name *string `pulumi:"name"`
	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess *string `pulumi:"rootAccess"`
	// The associated security groups.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The VPC subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	Url *string `pulumi:"url"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize *int `pulumi:"volumeSize"`
}

type NotebookInstanceState struct {
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn pulumi.StringPtrInput
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository pulumi.StringPtrInput
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess pulumi.StringPtrInput
	// The name of ML compute instance type.
	InstanceType pulumi.StringPtrInput
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId pulumi.StringPtrInput
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName pulumi.StringPtrInput
	// The name of the notebook instance (must be unique).
	Name pulumi.StringPtrInput
	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn pulumi.StringPtrInput
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess pulumi.StringPtrInput
	// The associated security groups.
	SecurityGroups pulumi.StringArrayInput
	// The VPC subnet ID.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	Url pulumi.StringPtrInput
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize pulumi.IntPtrInput
}

func (NotebookInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceState)(nil)).Elem()
}

type notebookInstanceArgs struct {
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories []string `pulumi:"additionalCodeRepositories"`
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository *string `pulumi:"defaultCodeRepository"`
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess *string `pulumi:"directInternetAccess"`
	// The name of ML compute instance type.
	InstanceType string `pulumi:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName *string `pulumi:"lifecycleConfigName"`
	// The name of the notebook instance (must be unique).
	Name *string `pulumi:"name"`
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess *string `pulumi:"rootAccess"`
	// The associated security groups.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The VPC subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize *int `pulumi:"volumeSize"`
}

// The set of arguments for constructing a NotebookInstance resource.
type NotebookInstanceArgs struct {
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories pulumi.StringArrayInput
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository pulumi.StringPtrInput
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess pulumi.StringPtrInput
	// The name of ML compute instance type.
	InstanceType pulumi.StringInput
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId pulumi.StringPtrInput
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName pulumi.StringPtrInput
	// The name of the notebook instance (must be unique).
	Name pulumi.StringPtrInput
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn pulumi.StringInput
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess pulumi.StringPtrInput
	// The associated security groups.
	SecurityGroups pulumi.StringArrayInput
	// The VPC subnet ID.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize pulumi.IntPtrInput
}

func (NotebookInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceArgs)(nil)).Elem()
}

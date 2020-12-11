// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic Beanstalk Application Version Resource. Elastic Beanstalk allows
// you to deploy and manage applications in the AWS cloud without worrying about
// the infrastructure that runs those applications.
//
// This resource creates a Beanstalk Application Version that can be deployed to a Beanstalk
// Environment.
//
// > **NOTE on Application Version Resource:**  When using the Application Version resource with multiple
// Elastic Beanstalk Environments it is possible that an error may be returned
// when attempting to delete an Application Version while it is still in use by a different environment.
// To work around this you can either create each environment in a separate AWS account or create your `elasticbeanstalk.ApplicationVersion` resources with a unique names in your Elastic Beanstalk Application. For example &lt;revision&gt;-&lt;environment&gt;.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticbeanstalk"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultBucket, err := s3.NewBucket(ctx, "defaultBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultBucketObject, err := s3.NewBucketObject(ctx, "defaultBucketObject", &s3.BucketObjectArgs{
// 			Bucket: defaultBucket.ID(),
// 			Key:    pulumi.String("beanstalk/go-v1.zip"),
// 			Source: pulumi.NewFileAsset("go-v1.zip"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticbeanstalk.NewApplication(ctx, "defaultApplication", &elasticbeanstalk.ApplicationArgs{
// 			Description: pulumi.String("tf-test-desc"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticbeanstalk.NewApplicationVersion(ctx, "defaultApplicationVersion", &elasticbeanstalk.ApplicationVersionArgs{
// 			Application: pulumi.String("tf-test-name"),
// 			Description: pulumi.String("application version"),
// 			Bucket:      defaultBucket.ID(),
// 			Key:         defaultBucketObject.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ApplicationVersion struct {
	pulumi.CustomResourceState

	// Name of the Beanstalk Application the version is associated with.
	Application pulumi.StringOutput `pulumi:"application"`
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// S3 bucket that contains the Application Version source bundle.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Short description of the Application Version.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// S3 object that is the Application Version source bundle.
	Key pulumi.StringOutput `pulumi:"key"`
	// A unique name for the this Application Version.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of tags for the Elastic Beanstalk Application Version.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewApplicationVersion(ctx *pulumi.Context,
	name string, args *ApplicationVersionArgs, opts ...pulumi.ResourceOption) (*ApplicationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Application == nil {
		return nil, errors.New("invalid value for required argument 'Application'")
	}
	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource ApplicationVersion
	err := ctx.RegisterResource("aws:elasticbeanstalk/applicationVersion:ApplicationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationVersion gets an existing ApplicationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationVersionState, opts ...pulumi.ResourceOption) (*ApplicationVersion, error) {
	var resource ApplicationVersion
	err := ctx.ReadResource("aws:elasticbeanstalk/applicationVersion:ApplicationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationVersion resources.
type applicationVersionState struct {
	// Name of the Beanstalk Application the version is associated with.
	Application *string `pulumi:"application"`
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn *string `pulumi:"arn"`
	// S3 bucket that contains the Application Version source bundle.
	Bucket *string `pulumi:"bucket"`
	// Short description of the Application Version.
	Description *string `pulumi:"description"`
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete *bool `pulumi:"forceDelete"`
	// S3 object that is the Application Version source bundle.
	Key *string `pulumi:"key"`
	// A unique name for the this Application Version.
	Name *string `pulumi:"name"`
	// Key-value map of tags for the Elastic Beanstalk Application Version.
	Tags map[string]string `pulumi:"tags"`
}

type ApplicationVersionState struct {
	// Name of the Beanstalk Application the version is associated with.
	Application pulumi.StringPtrInput
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn pulumi.StringPtrInput
	// S3 bucket that contains the Application Version source bundle.
	Bucket pulumi.StringPtrInput
	// Short description of the Application Version.
	Description pulumi.StringPtrInput
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete pulumi.BoolPtrInput
	// S3 object that is the Application Version source bundle.
	Key pulumi.StringPtrInput
	// A unique name for the this Application Version.
	Name pulumi.StringPtrInput
	// Key-value map of tags for the Elastic Beanstalk Application Version.
	Tags pulumi.StringMapInput
}

func (ApplicationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationVersionState)(nil)).Elem()
}

type applicationVersionArgs struct {
	// Name of the Beanstalk Application the version is associated with.
	Application interface{} `pulumi:"application"`
	// S3 bucket that contains the Application Version source bundle.
	Bucket interface{} `pulumi:"bucket"`
	// Short description of the Application Version.
	Description *string `pulumi:"description"`
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete *bool `pulumi:"forceDelete"`
	// S3 object that is the Application Version source bundle.
	Key string `pulumi:"key"`
	// A unique name for the this Application Version.
	Name *string `pulumi:"name"`
	// Key-value map of tags for the Elastic Beanstalk Application Version.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationVersion resource.
type ApplicationVersionArgs struct {
	// Name of the Beanstalk Application the version is associated with.
	Application pulumi.Input
	// S3 bucket that contains the Application Version source bundle.
	Bucket pulumi.Input
	// Short description of the Application Version.
	Description pulumi.StringPtrInput
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete pulumi.BoolPtrInput
	// S3 object that is the Application Version source bundle.
	Key pulumi.StringInput
	// A unique name for the this Application Version.
	Name pulumi.StringPtrInput
	// Key-value map of tags for the Elastic Beanstalk Application Version.
	Tags pulumi.StringMapInput
}

func (ApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationVersionArgs)(nil)).Elem()
}

type ApplicationVersionInput interface {
	pulumi.Input

	ToApplicationVersionOutput() ApplicationVersionOutput
	ToApplicationVersionOutputWithContext(ctx context.Context) ApplicationVersionOutput
}

type ApplicationVersionPtrInput interface {
	pulumi.Input

	ToApplicationVersionPtrOutput() ApplicationVersionPtrOutput
	ToApplicationVersionPtrOutputWithContext(ctx context.Context) ApplicationVersionPtrOutput
}

func (ApplicationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationVersion)(nil)).Elem()
}

func (i ApplicationVersion) ToApplicationVersionOutput() ApplicationVersionOutput {
	return i.ToApplicationVersionOutputWithContext(context.Background())
}

func (i ApplicationVersion) ToApplicationVersionOutputWithContext(ctx context.Context) ApplicationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationVersionOutput)
}

func (i ApplicationVersion) ToApplicationVersionPtrOutput() ApplicationVersionPtrOutput {
	return i.ToApplicationVersionPtrOutputWithContext(context.Background())
}

func (i ApplicationVersion) ToApplicationVersionPtrOutputWithContext(ctx context.Context) ApplicationVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationVersionPtrOutput)
}

type ApplicationVersionOutput struct {
	*pulumi.OutputState
}

func (ApplicationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationVersionOutput)(nil)).Elem()
}

func (o ApplicationVersionOutput) ToApplicationVersionOutput() ApplicationVersionOutput {
	return o
}

func (o ApplicationVersionOutput) ToApplicationVersionOutputWithContext(ctx context.Context) ApplicationVersionOutput {
	return o
}

type ApplicationVersionPtrOutput struct {
	*pulumi.OutputState
}

func (ApplicationVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationVersion)(nil)).Elem()
}

func (o ApplicationVersionPtrOutput) ToApplicationVersionPtrOutput() ApplicationVersionPtrOutput {
	return o
}

func (o ApplicationVersionPtrOutput) ToApplicationVersionPtrOutputWithContext(ctx context.Context) ApplicationVersionPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationVersionOutput{})
	pulumi.RegisterOutputType(ApplicationVersionPtrOutput{})
}

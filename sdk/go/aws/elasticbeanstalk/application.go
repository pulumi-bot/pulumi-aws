// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Beanstalk Application Resource. Elastic Beanstalk allows
// you to deploy and manage applications in the AWS cloud without worrying about
// the infrastructure that runs those applications.
//
// This resource creates an application that has one configuration template named
// `default`, and no application versions
type Application struct {
	pulumi.CustomResourceState

	AppversionLifecycle ApplicationAppversionLifecyclePtrOutput `pulumi:"appversionLifecycle"`
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Short description of the application
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the application, must be unique within your account
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("aws:elasticbeanstalk/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws:elasticbeanstalk/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	AppversionLifecycle *ApplicationAppversionLifecycle `pulumi:"appversionLifecycle"`
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn *string `pulumi:"arn"`
	// Short description of the application
	Description *string `pulumi:"description"`
	// The name of the application, must be unique within your account
	Name *string `pulumi:"name"`
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ApplicationState struct {
	AppversionLifecycle ApplicationAppversionLifecyclePtrInput
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn pulumi.StringPtrInput
	// Short description of the application
	Description pulumi.StringPtrInput
	// The name of the application, must be unique within your account
	Name pulumi.StringPtrInput
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags pulumi.MapInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	AppversionLifecycle *ApplicationAppversionLifecycleArgs `pulumi:"appversionLifecycle"`
	// Short description of the application
	Description *string `pulumi:"description"`
	// The name of the application, must be unique within your account
	Name *string `pulumi:"name"`
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	AppversionLifecycle ApplicationAppversionLifecycleArgsPtrInput
	// Short description of the application
	Description pulumi.StringPtrInput
	// The name of the application, must be unique within your account
	Name pulumi.StringPtrInput
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags pulumi.MapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

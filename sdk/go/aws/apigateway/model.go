// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Model for a REST API Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewModel(ctx, "myDemoModel", &apigateway.ModelArgs{
// 			RestApi:     myDemoAPI.ID(),
// 			Description: pulumi.String("a JSON schema"),
// 			ContentType: pulumi.String("application/json"),
// 			Schema:      pulumi.String(fmt.Sprintf("%v%v%v", "{\n", "  \"type\": \"object\"\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Model struct {
	pulumi.CustomResourceState

	// The content type of the model
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// The description of the model
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the model
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The schema of the model in a JSON form
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	var resource Model
	err := ctx.RegisterResource("aws:apigateway/model:Model", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelState, opts ...pulumi.ResourceOption) (*Model, error) {
	var resource Model
	err := ctx.ReadResource("aws:apigateway/model:Model", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Model resources.
type modelState struct {
	// The content type of the model
	ContentType *string `pulumi:"contentType"`
	// The description of the model
	Description *string `pulumi:"description"`
	// The name of the model
	Name *string `pulumi:"name"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The schema of the model in a JSON form
	Schema *string `pulumi:"schema"`
}

type ModelState struct {
	// The content type of the model
	ContentType pulumi.StringPtrInput
	// The description of the model
	Description pulumi.StringPtrInput
	// The name of the model
	Name pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The schema of the model in a JSON form
	Schema pulumi.StringPtrInput
}

func (ModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelState)(nil)).Elem()
}

type modelArgs struct {
	// The content type of the model
	ContentType string `pulumi:"contentType"`
	// The description of the model
	Description *string `pulumi:"description"`
	// The name of the model
	Name *string `pulumi:"name"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The schema of the model in a JSON form
	Schema *string `pulumi:"schema"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// The content type of the model
	ContentType pulumi.StringInput
	// The description of the model
	Description pulumi.StringPtrInput
	// The name of the model
	Name pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The schema of the model in a JSON form
	Schema pulumi.StringPtrInput
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelArgs)(nil)).Elem()
}

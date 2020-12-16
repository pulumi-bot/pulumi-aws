// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a settings of an API Gateway Documentation Part.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewDocumentationPart(ctx, "exampleDocumentationPart", &apigateway.DocumentationPartArgs{
// 			Location: &apigateway.DocumentationPartLocationArgs{
// 				Type:   pulumi.String("METHOD"),
// 				Method: pulumi.String("GET"),
// 				Path:   pulumi.String("/example"),
// 			},
// 			Properties: pulumi.String("{\"description\":\"Example description\"}"),
// 			RestApiId:  exampleRestApi.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// API Gateway documentation_parts can be imported using `REST-API-ID/DOC-PART-ID`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/documentationPart:DocumentationPart example 5i4e1ko720/3oyy3t
// ```
type DocumentationPart struct {
	pulumi.CustomResourceState

	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocationOutput `pulumi:"location"`
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties pulumi.StringOutput `pulumi:"properties"`
	// The ID of the associated Rest API
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
}

// NewDocumentationPart registers a new resource with the given unique name, arguments, and options.
func NewDocumentationPart(ctx *pulumi.Context,
	name string, args *DocumentationPartArgs, opts ...pulumi.ResourceOption) (*DocumentationPart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource DocumentationPart
	err := ctx.RegisterResource("aws:apigateway/documentationPart:DocumentationPart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentationPart gets an existing DocumentationPart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentationPart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentationPartState, opts ...pulumi.ResourceOption) (*DocumentationPart, error) {
	var resource DocumentationPart
	err := ctx.ReadResource("aws:apigateway/documentationPart:DocumentationPart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentationPart resources.
type documentationPartState struct {
	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location *DocumentationPartLocation `pulumi:"location"`
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties *string `pulumi:"properties"`
	// The ID of the associated Rest API
	RestApiId *string `pulumi:"restApiId"`
}

type DocumentationPartState struct {
	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocationPtrInput
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties pulumi.StringPtrInput
	// The ID of the associated Rest API
	RestApiId pulumi.StringPtrInput
}

func (DocumentationPartState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentationPartState)(nil)).Elem()
}

type documentationPartArgs struct {
	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocation `pulumi:"location"`
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties string `pulumi:"properties"`
	// The ID of the associated Rest API
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a DocumentationPart resource.
type DocumentationPartArgs struct {
	// The location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocationInput
	// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties pulumi.StringInput
	// The ID of the associated Rest API
	RestApiId pulumi.StringInput
}

func (DocumentationPartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentationPartArgs)(nil)).Elem()
}

type DocumentationPartInput interface {
	pulumi.Input

	ToDocumentationPartOutput() DocumentationPartOutput
	ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput
}

func (*DocumentationPart) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentationPart)(nil))
}

func (i *DocumentationPart) ToDocumentationPartOutput() DocumentationPartOutput {
	return i.ToDocumentationPartOutputWithContext(context.Background())
}

func (i *DocumentationPart) ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentationPartOutput)
}

type DocumentationPartOutput struct {
	*pulumi.OutputState
}

func (DocumentationPartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentationPart)(nil))
}

func (o DocumentationPartOutput) ToDocumentationPartOutput() DocumentationPartOutput {
	return o
}

func (o DocumentationPartOutput) ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DocumentationPartOutput{})
}

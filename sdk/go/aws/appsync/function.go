// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AppSync Function.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("API_KEY"),
// 			Schema:             pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "type Mutation {\n", "  putPost(id: ID!, title: String!): Post\n", "}\n", "\n", "type Post {\n", "  id: ID!\n", "  title: String!\n", "}\n", "\n", "type Query {\n", "  singlePost(id: ID!): Post\n", "}\n", "\n", "schema {\n", "  query: Query\n", "  mutation: Mutation\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataSource, err := appsync.NewDataSource(ctx, "exampleDataSource", &appsync.DataSourceArgs{
// 			ApiId: exampleGraphQLApi.ID(),
// 			Name:  pulumi.String("example"),
// 			Type:  pulumi.String("HTTP"),
// 			HttpConfig: &appsync.DataSourceHttpConfigArgs{
// 				Endpoint: pulumi.String("http://example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appsync.NewFunction(ctx, "exampleFunction", &appsync.FunctionArgs{
// 			ApiId:                   exampleGraphQLApi.ID(),
// 			DataSource:              exampleDataSource.Name,
// 			Name:                    pulumi.String("example"),
// 			RequestMappingTemplate:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"version\": \"2018-05-29\",\n", "    \"method\": \"GET\",\n", "    \"resourcePath\": \"/\",\n", "    \"params\":{\n", "        \"headers\": ", "$", "utils.http.copyheaders(", "$", "ctx.request.headers)\n", "    }\n", "}\n")),
// 			ResponseMappingTemplate: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "#if(", "$", "ctx.result.statusCode == 200)\n", "    ", "$", "ctx.result.body\n", "#else\n", "    ", "$", "utils.appendError(", "$", "ctx.result.body, ", "$", "ctx.result.statusCode)\n", "#end\n")),
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
// `aws_appsync_function` can be imported using the AppSync API ID and Function ID separated by `-`, e.g.
//
// ```sh
//  $ pulumi import aws:appsync/function:Function example xxxxx-yyyyy
// ```
type Function struct {
	pulumi.CustomResourceState

	// The ID of the associated AppSync API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The ARN of the Function object.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Function DataSource name.
	DataSource pulumi.StringOutput `pulumi:"dataSource"`
	// The Function description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique ID representing the Function object.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion pulumi.StringPtrOutput `pulumi:"functionVersion"`
	// The Function name. The function name does not have to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringOutput `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate pulumi.StringOutput `pulumi:"responseMappingTemplate"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DataSource == nil {
		return nil, errors.New("invalid value for required argument 'DataSource'")
	}
	if args.RequestMappingTemplate == nil {
		return nil, errors.New("invalid value for required argument 'RequestMappingTemplate'")
	}
	if args.ResponseMappingTemplate == nil {
		return nil, errors.New("invalid value for required argument 'ResponseMappingTemplate'")
	}
	var resource Function
	err := ctx.RegisterResource("aws:appsync/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("aws:appsync/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// The ID of the associated AppSync API.
	ApiId *string `pulumi:"apiId"`
	// The ARN of the Function object.
	Arn *string `pulumi:"arn"`
	// The Function DataSource name.
	DataSource *string `pulumi:"dataSource"`
	// The Function description.
	Description *string `pulumi:"description"`
	// A unique ID representing the Function object.
	FunctionId *string `pulumi:"functionId"`
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// The Function name. The function name does not have to be unique.
	Name *string `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate *string `pulumi:"responseMappingTemplate"`
}

type FunctionState struct {
	// The ID of the associated AppSync API.
	ApiId pulumi.StringPtrInput
	// The ARN of the Function object.
	Arn pulumi.StringPtrInput
	// The Function DataSource name.
	DataSource pulumi.StringPtrInput
	// The Function description.
	Description pulumi.StringPtrInput
	// A unique ID representing the Function object.
	FunctionId pulumi.StringPtrInput
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion pulumi.StringPtrInput
	// The Function name. The function name does not have to be unique.
	Name pulumi.StringPtrInput
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringPtrInput
	// The Function response mapping template.
	ResponseMappingTemplate pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// The ID of the associated AppSync API.
	ApiId string `pulumi:"apiId"`
	// The Function DataSource name.
	DataSource string `pulumi:"dataSource"`
	// The Function description.
	Description *string `pulumi:"description"`
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// The Function name. The function name does not have to be unique.
	Name *string `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate string `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate string `pulumi:"responseMappingTemplate"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// The ID of the associated AppSync API.
	ApiId pulumi.StringInput
	// The Function DataSource name.
	DataSource pulumi.StringInput
	// The Function description.
	Description pulumi.StringPtrInput
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion pulumi.StringPtrInput
	// The Function name. The function name does not have to be unique.
	Name pulumi.StringPtrInput
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringInput
	// The Function response mapping template.
	ResponseMappingTemplate pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((*Function)(nil))
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

func (i *Function) ToFunctionPtrOutput() FunctionPtrOutput {
	return i.ToFunctionPtrOutputWithContext(context.Background())
}

func (i *Function) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionPtrOutput)
}

type FunctionPtrInput interface {
	pulumi.Input

	ToFunctionPtrOutput() FunctionPtrOutput
	ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput
}

type functionPtrType FunctionArgs

func (*functionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil))
}

func (i *functionPtrType) ToFunctionPtrOutput() FunctionPtrOutput {
	return i.ToFunctionPtrOutputWithContext(context.Background())
}

func (i *functionPtrType) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionPtrOutput)
}

// FunctionArrayInput is an input type that accepts FunctionArray and FunctionArrayOutput values.
// You can construct a concrete instance of `FunctionArrayInput` via:
//
//          FunctionArray{ FunctionArgs{...} }
type FunctionArrayInput interface {
	pulumi.Input

	ToFunctionArrayOutput() FunctionArrayOutput
	ToFunctionArrayOutputWithContext(context.Context) FunctionArrayOutput
}

type FunctionArray []FunctionInput

func (FunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Function)(nil))
}

func (i FunctionArray) ToFunctionArrayOutput() FunctionArrayOutput {
	return i.ToFunctionArrayOutputWithContext(context.Background())
}

func (i FunctionArray) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionArrayOutput)
}

// FunctionMapInput is an input type that accepts FunctionMap and FunctionMapOutput values.
// You can construct a concrete instance of `FunctionMapInput` via:
//
//          FunctionMap{ "key": FunctionArgs{...} }
type FunctionMapInput interface {
	pulumi.Input

	ToFunctionMapOutput() FunctionMapOutput
	ToFunctionMapOutputWithContext(context.Context) FunctionMapOutput
}

type FunctionMap map[string]FunctionInput

func (FunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Function)(nil))
}

func (i FunctionMap) ToFunctionMapOutput() FunctionMapOutput {
	return i.ToFunctionMapOutputWithContext(context.Background())
}

func (i FunctionMap) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionMapOutput)
}

type FunctionOutput struct {
	*pulumi.OutputState
}

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Function)(nil))
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionPtrOutput() FunctionPtrOutput {
	return o.ToFunctionPtrOutputWithContext(context.Background())
}

func (o FunctionOutput) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return o.ApplyT(func(v Function) *Function {
		return &v
	}).(FunctionPtrOutput)
}

type FunctionPtrOutput struct {
	*pulumi.OutputState
}

func (FunctionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil))
}

func (o FunctionPtrOutput) ToFunctionPtrOutput() FunctionPtrOutput {
	return o
}

func (o FunctionPtrOutput) ToFunctionPtrOutputWithContext(ctx context.Context) FunctionPtrOutput {
	return o
}

type FunctionArrayOutput struct{ *pulumi.OutputState }

func (FunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Function)(nil))
}

func (o FunctionArrayOutput) ToFunctionArrayOutput() FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) Index(i pulumi.IntInput) FunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Function {
		return vs[0].([]Function)[vs[1].(int)]
	}).(FunctionOutput)
}

type FunctionMapOutput struct{ *pulumi.OutputState }

func (FunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Function)(nil))
}

func (o FunctionMapOutput) ToFunctionMapOutput() FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) MapIndex(k pulumi.StringInput) FunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Function {
		return vs[0].(map[string]Function)[vs[1].(string)]
	}).(FunctionOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionOutput{})
	pulumi.RegisterOutputType(FunctionPtrOutput{})
	pulumi.RegisterOutputType(FunctionArrayOutput{})
	pulumi.RegisterOutputType(FunctionMapOutput{})
}

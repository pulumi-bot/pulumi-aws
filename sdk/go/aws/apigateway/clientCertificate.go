// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Client Certificate.
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
// 		_, err := apigateway.NewClientCertificate(ctx, "demo", &apigateway.ClientCertificateArgs{
// 			Description: pulumi.String("My client certificate"),
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
// API Gateway Client Certificates can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/clientCertificate:ClientCertificate demo ab1cqe
// ```
type ClientCertificate struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date when the client certificate was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the client certificate.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The date when the client certificate will expire.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// The PEM-encoded public key of the client certificate.
	PemEncodedCertificate pulumi.StringOutput `pulumi:"pemEncodedCertificate"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewClientCertificate registers a new resource with the given unique name, arguments, and options.
func NewClientCertificate(ctx *pulumi.Context,
	name string, args *ClientCertificateArgs, opts ...pulumi.ResourceOption) (*ClientCertificate, error) {
	if args == nil {
		args = &ClientCertificateArgs{}
	}

	var resource ClientCertificate
	err := ctx.RegisterResource("aws:apigateway/clientCertificate:ClientCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientCertificate gets an existing ClientCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientCertificateState, opts ...pulumi.ResourceOption) (*ClientCertificate, error) {
	var resource ClientCertificate
	err := ctx.ReadResource("aws:apigateway/clientCertificate:ClientCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientCertificate resources.
type clientCertificateState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The date when the client certificate was created.
	CreatedDate *string `pulumi:"createdDate"`
	// The description of the client certificate.
	Description *string `pulumi:"description"`
	// The date when the client certificate will expire.
	ExpirationDate *string `pulumi:"expirationDate"`
	// The PEM-encoded public key of the client certificate.
	PemEncodedCertificate *string `pulumi:"pemEncodedCertificate"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type ClientCertificateState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The date when the client certificate was created.
	CreatedDate pulumi.StringPtrInput
	// The description of the client certificate.
	Description pulumi.StringPtrInput
	// The date when the client certificate will expire.
	ExpirationDate pulumi.StringPtrInput
	// The PEM-encoded public key of the client certificate.
	PemEncodedCertificate pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (ClientCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientCertificateState)(nil)).Elem()
}

type clientCertificateArgs struct {
	// The description of the client certificate.
	Description *string `pulumi:"description"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClientCertificate resource.
type ClientCertificateArgs struct {
	// The description of the client certificate.
	Description pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (ClientCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientCertificateArgs)(nil)).Elem()
}

type ClientCertificateInput interface {
	pulumi.Input

	ToClientCertificateOutput() ClientCertificateOutput
	ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput
}

type ClientCertificatePtrInput interface {
	pulumi.Input

	ToClientCertificatePtrOutput() ClientCertificatePtrOutput
	ToClientCertificatePtrOutputWithContext(ctx context.Context) ClientCertificatePtrOutput
}

func (ClientCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (i ClientCertificate) ToClientCertificateOutput() ClientCertificateOutput {
	return i.ToClientCertificateOutputWithContext(context.Background())
}

func (i ClientCertificate) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateOutput)
}

func (i ClientCertificate) ToClientCertificatePtrOutput() ClientCertificatePtrOutput {
	return i.ToClientCertificatePtrOutputWithContext(context.Background())
}

func (i ClientCertificate) ToClientCertificatePtrOutputWithContext(ctx context.Context) ClientCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificatePtrOutput)
}

type ClientCertificateOutput struct {
	*pulumi.OutputState
}

func (ClientCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateOutput)(nil)).Elem()
}

func (o ClientCertificateOutput) ToClientCertificateOutput() ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return o
}

type ClientCertificatePtrOutput struct {
	*pulumi.OutputState
}

func (ClientCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientCertificate)(nil)).Elem()
}

func (o ClientCertificatePtrOutput) ToClientCertificatePtrOutput() ClientCertificatePtrOutput {
	return o
}

func (o ClientCertificatePtrOutput) ToClientCertificatePtrOutputWithContext(ctx context.Context) ClientCertificatePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClientCertificateOutput{})
	pulumi.RegisterOutputType(ClientCertificatePtrOutput{})
}

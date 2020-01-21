// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package clientCertificate

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway Client Certificate.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_client_certificate.html.markdown.
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
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
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
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
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
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (ClientCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientCertificateState)(nil)).Elem()
}

type clientCertificateArgs struct {
	// The description of the client certificate.
	Description *string `pulumi:"description"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ClientCertificate resource.
type ClientCertificateArgs struct {
	// The description of the client certificate.
	Description pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
}

func (ClientCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientCertificateArgs)(nil)).Elem()
}


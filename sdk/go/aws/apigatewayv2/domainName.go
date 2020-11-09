// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 domain name.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
//
// > **Note:** This resource establishes ownership of and the TLS settings for
// a particular domain name. An API stage can be associated with the domain name using the `apigatewayv2.ApiMapping` resource.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewDomainName(ctx, "example", &apigatewayv2.DomainNameArgs{
// 			DomainName: pulumi.String("ws-api.example.com"),
// 			DomainNameConfiguration: &apigatewayv2.DomainNameDomainNameConfigurationArgs{
// 				CertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
// 				EndpointType:   pulumi.String("REGIONAL"),
// 				SecurityPolicy: pulumi.String("TLS_1_2"),
// 			},
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
// `aws_apigatewayv2_domain_name` can be imported by using the domain name, e.g.
//
// ```sh
//  $ pulumi import aws:apigatewayv2/domainName:DomainName example ws-api.example.com
// ```
type DomainName struct {
	pulumi.CustomResourceState

	// The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
	ApiMappingSelectionExpression pulumi.StringOutput `pulumi:"apiMappingSelectionExpression"`
	// The ARN of the domain name.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfigurationOutput `pulumi:"domainNameConfiguration"`
	// A map of tags to assign to the domain name.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDomainName registers a new resource with the given unique name, arguments, and options.
func NewDomainName(ctx *pulumi.Context,
	name string, args *DomainNameArgs, opts ...pulumi.ResourceOption) (*DomainName, error) {
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	if args == nil || args.DomainNameConfiguration == nil {
		return nil, errors.New("missing required argument 'DomainNameConfiguration'")
	}
	if args == nil {
		args = &DomainNameArgs{}
	}
	var resource DomainName
	err := ctx.RegisterResource("aws:apigatewayv2/domainName:DomainName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainName gets an existing DomainName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainNameState, opts ...pulumi.ResourceOption) (*DomainName, error) {
	var resource DomainName
	err := ctx.ReadResource("aws:apigatewayv2/domainName:DomainName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainName resources.
type domainNameState struct {
	// The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
	ApiMappingSelectionExpression *string `pulumi:"apiMappingSelectionExpression"`
	// The ARN of the domain name.
	Arn *string `pulumi:"arn"`
	// The domain name.
	DomainName *string `pulumi:"domainName"`
	// The domain name configuration.
	DomainNameConfiguration *DomainNameDomainNameConfiguration `pulumi:"domainNameConfiguration"`
	// A map of tags to assign to the domain name.
	Tags map[string]string `pulumi:"tags"`
}

type DomainNameState struct {
	// The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
	ApiMappingSelectionExpression pulumi.StringPtrInput
	// The ARN of the domain name.
	Arn pulumi.StringPtrInput
	// The domain name.
	DomainName pulumi.StringPtrInput
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfigurationPtrInput
	// A map of tags to assign to the domain name.
	Tags pulumi.StringMapInput
}

func (DomainNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameState)(nil)).Elem()
}

type domainNameArgs struct {
	// The domain name.
	DomainName string `pulumi:"domainName"`
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfiguration `pulumi:"domainNameConfiguration"`
	// A map of tags to assign to the domain name.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DomainName resource.
type DomainNameArgs struct {
	// The domain name.
	DomainName pulumi.StringInput
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfigurationInput
	// A map of tags to assign to the domain name.
	Tags pulumi.StringMapInput
}

func (DomainNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameArgs)(nil)).Elem()
}

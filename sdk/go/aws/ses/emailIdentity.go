// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SES email identity resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewEmailIdentity(ctx, "example", &ses.EmailIdentityArgs{
// 			Email: pulumi.String("email@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EmailIdentity struct {
	pulumi.CustomResourceState

	// The ARN of the email identity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The email address to assign to SES
	Email pulumi.StringOutput `pulumi:"email"`
}

// NewEmailIdentity registers a new resource with the given unique name, arguments, and options.
func NewEmailIdentity(ctx *pulumi.Context,
	name string, args *EmailIdentityArgs, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	var resource EmailIdentity
	err := ctx.RegisterResource("aws:ses/emailIdentity:EmailIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailIdentity gets an existing EmailIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailIdentityState, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	var resource EmailIdentity
	err := ctx.ReadResource("aws:ses/emailIdentity:EmailIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailIdentity resources.
type emailIdentityState struct {
	// The ARN of the email identity.
	Arn *string `pulumi:"arn"`
	// The email address to assign to SES
	Email *string `pulumi:"email"`
}

type EmailIdentityState struct {
	// The ARN of the email identity.
	Arn pulumi.StringPtrInput
	// The email address to assign to SES
	Email pulumi.StringPtrInput
}

func (EmailIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityState)(nil)).Elem()
}

type emailIdentityArgs struct {
	// The email address to assign to SES
	Email string `pulumi:"email"`
}

// The set of arguments for constructing a EmailIdentity resource.
type EmailIdentityArgs struct {
	// The email address to assign to SES
	Email pulumi.StringInput
}

func (EmailIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityArgs)(nil)).Elem()
}

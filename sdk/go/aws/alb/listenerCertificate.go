// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Load Balancer Listener Certificate resource.
//
// This resource is for additional certificates and does not replace the default certificate on the listener.
//
// > **Note:** `alb.ListenerCertificate` is known as `lb.ListenerCertificate`. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acm"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleCertificate, err := acm.NewCertificate(ctx, "exampleCertificate", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerCertificate(ctx, "exampleListenerCertificate", &lb.ListenerCertificateArgs{
// 			CertificateArn: exampleCertificate.Arn,
// 			ListenerArn:    frontEndListener.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ListenerCertificate struct {
	pulumi.CustomResourceState

	// The ARN of the certificate to attach to the listener.
	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`
	// The ARN of the listener to which to attach the certificate.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
}

// NewListenerCertificate registers a new resource with the given unique name, arguments, and options.
func NewListenerCertificate(ctx *pulumi.Context,
	name string, args *ListenerCertificateArgs, opts ...pulumi.ResourceOption) (*ListenerCertificate, error) {
	if args == nil || args.CertificateArn == nil {
		return nil, errors.New("missing required argument 'CertificateArn'")
	}
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	if args == nil {
		args = &ListenerCertificateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:applicationloadbalancing/listenerCertificate:ListenerCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource ListenerCertificate
	err := ctx.RegisterResource("aws:alb/listenerCertificate:ListenerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerCertificate gets an existing ListenerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerCertificateState, opts ...pulumi.ResourceOption) (*ListenerCertificate, error) {
	var resource ListenerCertificate
	err := ctx.ReadResource("aws:alb/listenerCertificate:ListenerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerCertificate resources.
type listenerCertificateState struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn *string `pulumi:"certificateArn"`
	// The ARN of the listener to which to attach the certificate.
	ListenerArn *string `pulumi:"listenerArn"`
}

type ListenerCertificateState struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn pulumi.StringPtrInput
	// The ARN of the listener to which to attach the certificate.
	ListenerArn pulumi.StringPtrInput
}

func (ListenerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerCertificateState)(nil)).Elem()
}

type listenerCertificateArgs struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn string `pulumi:"certificateArn"`
	// The ARN of the listener to which to attach the certificate.
	ListenerArn string `pulumi:"listenerArn"`
}

// The set of arguments for constructing a ListenerCertificate resource.
type ListenerCertificateArgs struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn pulumi.StringInput
	// The ARN of the listener to which to attach the certificate.
	ListenerArn pulumi.StringInput
}

func (ListenerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerCertificateArgs)(nil)).Elem()
}

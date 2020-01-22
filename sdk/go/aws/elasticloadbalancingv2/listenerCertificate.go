// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancingv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer Listener Certificate resource.
// 
// This resource is for additional certificates and does not replace the default certificate on the listener.
// 
// > **Note:** `alb.ListenerCertificate` is known as `lb.ListenerCertificate`. The functionality is identical.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener_certificate_legacy.html.markdown.
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
	var resource ListenerCertificate
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate", name, id, state, &resource, opts...)
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


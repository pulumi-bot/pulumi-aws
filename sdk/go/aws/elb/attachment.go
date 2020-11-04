// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches an EC2 instance to an Elastic Load Balancer (ELB). For attaching resources with Application Load Balancer (ALB) or Network Load Balancer (NLB), see the `lb.TargetGroupAttachment` resource.
//
// > **NOTE on ELB Instances and ELB Attachments:** This provider currently provides
// both a standalone ELB Attachment resource (describing an instance attached to
// an ELB), and an Elastic Load Balancer resource with
// `instances` defined in-line. At this time you cannot use an ELB with in-line
// instances in conjunction with an ELB Attachment resource. Doing so will cause a
// conflict and will overwrite attachments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elb.NewAttachment(ctx, "baz", &elb.AttachmentArgs{
// 			Elb:      pulumi.Any(aws_elb.Bar.Id),
// 			Instance: pulumi.Any(aws_instance.Foo.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Attachment struct {
	pulumi.CustomResourceState

	// The name of the ELB.
	Elb pulumi.StringOutput `pulumi:"elb"`
	// Instance ID to place in the ELB pool.
	Instance pulumi.StringOutput `pulumi:"instance"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Elb == nil {
		return nil, errors.New("invalid value for required argument 'Elb'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancing/attachment:Attachment"),
		},
	})
	opts = append(opts, aliases)
	var resource Attachment
	err := ctx.RegisterResource("aws:elb/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("aws:elb/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	// The name of the ELB.
	Elb *string `pulumi:"elb"`
	// Instance ID to place in the ELB pool.
	Instance *string `pulumi:"instance"`
}

type AttachmentState struct {
	// The name of the ELB.
	Elb pulumi.StringPtrInput
	// Instance ID to place in the ELB pool.
	Instance pulumi.StringPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	// The name of the ELB.
	Elb string `pulumi:"elb"`
	// Instance ID to place in the ELB pool.
	Instance string `pulumi:"instance"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The name of the ELB.
	Elb pulumi.StringInput
	// Instance ID to place in the ELB pool.
	Instance pulumi.StringInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}

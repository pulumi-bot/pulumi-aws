// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the `elb.Attachment` resource.
//
// > **Note:** `alb.TargetGroupAttachment` is known as `lb.TargetGroupAttachment`. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testTargetGroup, err := lb.NewTargetGroup(ctx, "testTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testInstance, err := ec2.NewInstance(ctx, "testInstance", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewTargetGroupAttachment(ctx, "testTargetGroupAttachment", &lb.TargetGroupAttachmentArgs{
// 			Port:           pulumi.Int(80),
// 			TargetGroupArn: testTargetGroup.Arn,
// 			TargetId:       testInstance.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Usage with lambda
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testTargetGroup, err := lb.NewTargetGroup(ctx, "testTargetGroup", &lb.TargetGroupArgs{
// 			TargetType: pulumi.String("lambda"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testFunction, err := lambda.NewFunction(ctx, "testFunction", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lambda.NewPermission(ctx, "withLb", &lambda.PermissionArgs{
// 			Action:    pulumi.String("lambda:InvokeFunction"),
// 			Function:  testFunction.Arn,
// 			Principal: pulumi.String("elasticloadbalancing.amazonaws.com"),
// 			SourceArn: testTargetGroup.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewTargetGroupAttachment(ctx, "testTargetGroupAttachment", &lb.TargetGroupAttachmentArgs{
// 			TargetGroupArn: testTargetGroup.Arn,
// 			TargetId:       testFunction.Arn,
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			"aws_lambda_permission.with_lb",
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.elasticloadbalancingv2.TargetGroupAttachment has been deprecated in favor of aws.lb.TargetGroupAttachment
type TargetGroupAttachment struct {
	pulumi.CustomResourceState

	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// The port on which targets receive traffic.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The ARN of the target group with which to register targets
	TargetGroupArn pulumi.StringOutput `pulumi:"targetGroupArn"`
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
}

// NewTargetGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewTargetGroupAttachment(ctx *pulumi.Context,
	name string, args *TargetGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*TargetGroupAttachment, error) {
	if args == nil || args.TargetGroupArn == nil {
		return nil, errors.New("missing required argument 'TargetGroupArn'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	if args == nil {
		args = &TargetGroupAttachmentArgs{}
	}
	var resource TargetGroupAttachment
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGroupAttachment gets an existing TargetGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGroupAttachmentState, opts ...pulumi.ResourceOption) (*TargetGroupAttachment, error) {
	var resource TargetGroupAttachment
	err := ctx.ReadResource("aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGroupAttachment resources.
type targetGroupAttachmentState struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The port on which targets receive traffic.
	Port *int `pulumi:"port"`
	// The ARN of the target group with which to register targets
	TargetGroupArn *string `pulumi:"targetGroupArn"`
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId *string `pulumi:"targetId"`
}

type TargetGroupAttachmentState struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone pulumi.StringPtrInput
	// The port on which targets receive traffic.
	Port pulumi.IntPtrInput
	// The ARN of the target group with which to register targets
	TargetGroupArn pulumi.StringPtrInput
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId pulumi.StringPtrInput
}

func (TargetGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupAttachmentState)(nil)).Elem()
}

type targetGroupAttachmentArgs struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The port on which targets receive traffic.
	Port *int `pulumi:"port"`
	// The ARN of the target group with which to register targets
	TargetGroupArn string `pulumi:"targetGroupArn"`
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId string `pulumi:"targetId"`
}

// The set of arguments for constructing a TargetGroupAttachment resource.
type TargetGroupAttachmentArgs struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone pulumi.StringPtrInput
	// The port on which targets receive traffic.
	Port pulumi.IntPtrInput
	// The ARN of the target group with which to register targets
	TargetGroupArn pulumi.StringInput
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId pulumi.StringInput
}

func (TargetGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupAttachmentArgs)(nil)).Elem()
}

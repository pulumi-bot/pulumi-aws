// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables AWS Shield Advanced for a specific AWS resource.
// The resource can be an Amazon CloudFront distribution, Elastic Load Balancing load balancer, AWS Global Accelerator accelerator, Elastic IP Address, or an Amazon Route 53 hosted zone.
//
// ## Example Usage
// ### Create protection
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/shield"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := aws.LookupAvailabilityZones(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentRegion, err := aws.LookupRegion(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentCallerIdentity, err := aws.LookupCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooEip, err := ec2.NewEip(ctx, "fooEip", &ec2.EipArgs{
// 			Vpc: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = shield.NewProtection(ctx, "fooProtection", &shield.ProtectionArgs{
// 			ResourceArn: fooEip.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:ec2:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":eip-allocation/", id), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Protection struct {
	pulumi.CustomResourceState

	// A friendly name for the Protection you are creating.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewProtection registers a new resource with the given unique name, arguments, and options.
func NewProtection(ctx *pulumi.Context,
	name string, args *ProtectionArgs, opts ...pulumi.ResourceOption) (*Protection, error) {
	if args == nil || args.ResourceArn == nil {
		return nil, errors.New("missing required argument 'ResourceArn'")
	}
	if args == nil {
		args = &ProtectionArgs{}
	}
	var resource Protection
	err := ctx.RegisterResource("aws:shield/protection:Protection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtection gets an existing Protection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionState, opts ...pulumi.ResourceOption) (*Protection, error) {
	var resource Protection
	err := ctx.ReadResource("aws:shield/protection:Protection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Protection resources.
type protectionState struct {
	// A friendly name for the Protection you are creating.
	Name *string `pulumi:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn *string `pulumi:"resourceArn"`
}

type ProtectionState struct {
	// A friendly name for the Protection you are creating.
	Name pulumi.StringPtrInput
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn pulumi.StringPtrInput
}

func (ProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionState)(nil)).Elem()
}

type protectionArgs struct {
	// A friendly name for the Protection you are creating.
	Name *string `pulumi:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a Protection resource.
type ProtectionArgs struct {
	// A friendly name for the Protection you are creating.
	Name pulumi.StringPtrInput
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn pulumi.StringInput
}

func (ProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionArgs)(nil)).Elem()
}

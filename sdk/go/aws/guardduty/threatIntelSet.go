// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage a GuardDuty ThreatIntelSet.
//
// > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage ThreatIntelSets. ThreatIntelSets that are uploaded by the primary account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-threat-intel-set.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/guardduty"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := guardduty.NewDetector(ctx, "primary", &guardduty.DetectorArgs{
// 			Enable: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myThreatIntelSetBucketObject, err := s3.NewBucketObject(ctx, "myThreatIntelSetBucketObject", &s3.BucketObjectArgs{
// 			Acl:     pulumi.String("public-read"),
// 			Content: pulumi.String("10.0.0.0/8\n"),
// 			Bucket:  bucket.ID(),
// 			Key:     pulumi.String("MyThreatIntelSet"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = guardduty.NewThreatIntelSet(ctx, "myThreatIntelSetThreatIntelSet", &guardduty.ThreatIntelSetArgs{
// 			Activate:   pulumi.Bool(true),
// 			DetectorId: primary.ID(),
// 			Format:     pulumi.String("TXT"),
// 			Location: pulumi.All(myThreatIntelSetBucketObject.Bucket, myThreatIntelSetBucketObject.Key).ApplyT(func(_args []interface{}) (string, error) {
// 				bucket := _args[0].(string)
// 				key := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v", "https://s3.amazonaws.com/", bucket, "/", key), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ThreatIntelSet struct {
	pulumi.CustomResourceState

	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate pulumi.BoolOutput `pulumi:"activate"`
	// Amazon Resource Name (ARN) of the GuardDuty ThreatIntelSet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringOutput `pulumi:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location pulumi.StringOutput `pulumi:"location"`
	// The friendly name to identify the ThreatIntelSet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewThreatIntelSet registers a new resource with the given unique name, arguments, and options.
func NewThreatIntelSet(ctx *pulumi.Context,
	name string, args *ThreatIntelSetArgs, opts ...pulumi.ResourceOption) (*ThreatIntelSet, error) {
	if args == nil || args.Activate == nil {
		return nil, errors.New("missing required argument 'Activate'")
	}
	if args == nil || args.DetectorId == nil {
		return nil, errors.New("missing required argument 'DetectorId'")
	}
	if args == nil || args.Format == nil {
		return nil, errors.New("missing required argument 'Format'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &ThreatIntelSetArgs{}
	}
	var resource ThreatIntelSet
	err := ctx.RegisterResource("aws:guardduty/threatIntelSet:ThreatIntelSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThreatIntelSet gets an existing ThreatIntelSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreatIntelSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelSetState, opts ...pulumi.ResourceOption) (*ThreatIntelSet, error) {
	var resource ThreatIntelSet
	err := ctx.ReadResource("aws:guardduty/threatIntelSet:ThreatIntelSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThreatIntelSet resources.
type threatIntelSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate *bool `pulumi:"activate"`
	// Amazon Resource Name (ARN) of the GuardDuty ThreatIntelSet.
	Arn *string `pulumi:"arn"`
	// The detector ID of the GuardDuty.
	DetectorId *string `pulumi:"detectorId"`
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format *string `pulumi:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location *string `pulumi:"location"`
	// The friendly name to identify the ThreatIntelSet.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type ThreatIntelSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the GuardDuty ThreatIntelSet.
	Arn pulumi.StringPtrInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringPtrInput
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringPtrInput
	// The URI of the file that contains the ThreatIntelSet.
	Location pulumi.StringPtrInput
	// The friendly name to identify the ThreatIntelSet.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (ThreatIntelSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelSetState)(nil)).Elem()
}

type threatIntelSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate bool `pulumi:"activate"`
	// The detector ID of the GuardDuty.
	DetectorId string `pulumi:"detectorId"`
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format string `pulumi:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location string `pulumi:"location"`
	// The friendly name to identify the ThreatIntelSet.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ThreatIntelSet resource.
type ThreatIntelSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate pulumi.BoolInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringInput
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringInput
	// The URI of the file that contains the ThreatIntelSet.
	Location pulumi.StringInput
	// The friendly name to identify the ThreatIntelSet.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (ThreatIntelSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelSetArgs)(nil)).Elem()
}

type ThreatIntelSetInput interface {
	pulumi.Input

	ToThreatIntelSetOutput() ThreatIntelSetOutput
	ToThreatIntelSetOutputWithContext(ctx context.Context) ThreatIntelSetOutput
}

func (ThreatIntelSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelSet)(nil)).Elem()
}

func (i ThreatIntelSet) ToThreatIntelSetOutput() ThreatIntelSetOutput {
	return i.ToThreatIntelSetOutputWithContext(context.Background())
}

func (i ThreatIntelSet) ToThreatIntelSetOutputWithContext(ctx context.Context) ThreatIntelSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelSetOutput)
}

type ThreatIntelSetOutput struct {
	*pulumi.OutputState
}

func (ThreatIntelSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelSetOutput)(nil)).Elem()
}

func (o ThreatIntelSetOutput) ToThreatIntelSetOutput() ThreatIntelSetOutput {
	return o
}

func (o ThreatIntelSetOutput) ToThreatIntelSetOutputWithContext(ctx context.Context) ThreatIntelSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelSetOutput{})
}

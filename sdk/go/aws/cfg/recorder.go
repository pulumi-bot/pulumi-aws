// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Config Configuration Recorder. Please note that this resource **does not start** the created recorder automatically.
//
// > **Note:** _Starting_ the Configuration Recorder requires a `delivery channel` (while delivery channel creation requires Configuration Recorder). This is why `cfg.RecorderStatus` is a separate resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
// 			RoleArn: role.Arn,
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
// Configuration Recorder can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import aws:cfg/recorder:Recorder foo example
// ```
type Recorder struct {
	pulumi.CustomResourceState

	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Recording group - see below.
	RecordingGroup RecorderRecordingGroupOutput `pulumi:"recordingGroup"`
	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewRecorder registers a new resource with the given unique name, arguments, and options.
func NewRecorder(ctx *pulumi.Context,
	name string, args *RecorderArgs, opts ...pulumi.ResourceOption) (*Recorder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Recorder
	err := ctx.RegisterResource("aws:cfg/recorder:Recorder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecorder gets an existing Recorder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecorder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecorderState, opts ...pulumi.ResourceOption) (*Recorder, error) {
	var resource Recorder
	err := ctx.ReadResource("aws:cfg/recorder:Recorder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Recorder resources.
type recorderState struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name *string `pulumi:"name"`
	// Recording group - see below.
	RecordingGroup *RecorderRecordingGroup `pulumi:"recordingGroup"`
	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn *string `pulumi:"roleArn"`
}

type RecorderState struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringPtrInput
	// Recording group - see below.
	RecordingGroup RecorderRecordingGroupPtrInput
	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn pulumi.StringPtrInput
}

func (RecorderState) ElementType() reflect.Type {
	return reflect.TypeOf((*recorderState)(nil)).Elem()
}

type recorderArgs struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name *string `pulumi:"name"`
	// Recording group - see below.
	RecordingGroup *RecorderRecordingGroup `pulumi:"recordingGroup"`
	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a Recorder resource.
type RecorderArgs struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringPtrInput
	// Recording group - see below.
	RecordingGroup RecorderRecordingGroupPtrInput
	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn pulumi.StringInput
}

func (RecorderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recorderArgs)(nil)).Elem()
}

type RecorderInput interface {
	pulumi.Input

	ToRecorderOutput() RecorderOutput
	ToRecorderOutputWithContext(ctx context.Context) RecorderOutput
}

func (*Recorder) ElementType() reflect.Type {
	return reflect.TypeOf((*Recorder)(nil))
}

func (i *Recorder) ToRecorderOutput() RecorderOutput {
	return i.ToRecorderOutputWithContext(context.Background())
}

func (i *Recorder) ToRecorderOutputWithContext(ctx context.Context) RecorderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecorderOutput)
}

func (i *Recorder) ToRecorderPtrOutput() RecorderPtrOutput {
	return i.ToRecorderPtrOutputWithContext(context.Background())
}

func (i *Recorder) ToRecorderPtrOutputWithContext(ctx context.Context) RecorderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecorderPtrOutput)
}

type RecorderPtrInput interface {
	pulumi.Input

	ToRecorderPtrOutput() RecorderPtrOutput
	ToRecorderPtrOutputWithContext(ctx context.Context) RecorderPtrOutput
}

type recorderPtrType RecorderArgs

func (*recorderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Recorder)(nil))
}

func (i *recorderPtrType) ToRecorderPtrOutput() RecorderPtrOutput {
	return i.ToRecorderPtrOutputWithContext(context.Background())
}

func (i *recorderPtrType) ToRecorderPtrOutputWithContext(ctx context.Context) RecorderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecorderOutput).ToRecorderPtrOutput()
}

// RecorderArrayInput is an input type that accepts RecorderArray and RecorderArrayOutput values.
// You can construct a concrete instance of `RecorderArrayInput` via:
//
//          RecorderArray{ RecorderArgs{...} }
type RecorderArrayInput interface {
	pulumi.Input

	ToRecorderArrayOutput() RecorderArrayOutput
	ToRecorderArrayOutputWithContext(context.Context) RecorderArrayOutput
}

type RecorderArray []RecorderInput

func (RecorderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Recorder)(nil))
}

func (i RecorderArray) ToRecorderArrayOutput() RecorderArrayOutput {
	return i.ToRecorderArrayOutputWithContext(context.Background())
}

func (i RecorderArray) ToRecorderArrayOutputWithContext(ctx context.Context) RecorderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecorderArrayOutput)
}

// RecorderMapInput is an input type that accepts RecorderMap and RecorderMapOutput values.
// You can construct a concrete instance of `RecorderMapInput` via:
//
//          RecorderMap{ "key": RecorderArgs{...} }
type RecorderMapInput interface {
	pulumi.Input

	ToRecorderMapOutput() RecorderMapOutput
	ToRecorderMapOutputWithContext(context.Context) RecorderMapOutput
}

type RecorderMap map[string]RecorderInput

func (RecorderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Recorder)(nil))
}

func (i RecorderMap) ToRecorderMapOutput() RecorderMapOutput {
	return i.ToRecorderMapOutputWithContext(context.Background())
}

func (i RecorderMap) ToRecorderMapOutputWithContext(ctx context.Context) RecorderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecorderMapOutput)
}

type RecorderOutput struct {
	*pulumi.OutputState
}

func (RecorderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Recorder)(nil))
}

func (o RecorderOutput) ToRecorderOutput() RecorderOutput {
	return o
}

func (o RecorderOutput) ToRecorderOutputWithContext(ctx context.Context) RecorderOutput {
	return o
}

func (o RecorderOutput) ToRecorderPtrOutput() RecorderPtrOutput {
	return o.ToRecorderPtrOutputWithContext(context.Background())
}

func (o RecorderOutput) ToRecorderPtrOutputWithContext(ctx context.Context) RecorderPtrOutput {
	return o.ApplyT(func(v Recorder) *Recorder {
		return &v
	}).(RecorderPtrOutput)
}

type RecorderPtrOutput struct {
	*pulumi.OutputState
}

func (RecorderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Recorder)(nil))
}

func (o RecorderPtrOutput) ToRecorderPtrOutput() RecorderPtrOutput {
	return o
}

func (o RecorderPtrOutput) ToRecorderPtrOutputWithContext(ctx context.Context) RecorderPtrOutput {
	return o
}

type RecorderArrayOutput struct{ *pulumi.OutputState }

func (RecorderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Recorder)(nil))
}

func (o RecorderArrayOutput) ToRecorderArrayOutput() RecorderArrayOutput {
	return o
}

func (o RecorderArrayOutput) ToRecorderArrayOutputWithContext(ctx context.Context) RecorderArrayOutput {
	return o
}

func (o RecorderArrayOutput) Index(i pulumi.IntInput) RecorderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Recorder {
		return vs[0].([]Recorder)[vs[1].(int)]
	}).(RecorderOutput)
}

type RecorderMapOutput struct{ *pulumi.OutputState }

func (RecorderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Recorder)(nil))
}

func (o RecorderMapOutput) ToRecorderMapOutput() RecorderMapOutput {
	return o
}

func (o RecorderMapOutput) ToRecorderMapOutputWithContext(ctx context.Context) RecorderMapOutput {
	return o
}

func (o RecorderMapOutput) MapIndex(k pulumi.StringInput) RecorderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Recorder {
		return vs[0].(map[string]Recorder)[vs[1].(string)]
	}).(RecorderOutput)
}

func init() {
	pulumi.RegisterOutputType(RecorderOutput{})
	pulumi.RegisterOutputType(RecorderPtrOutput{})
	pulumi.RegisterOutputType(RecorderArrayOutput{})
	pulumi.RegisterOutputType(RecorderMapOutput{})
}

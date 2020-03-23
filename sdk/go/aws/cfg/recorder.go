// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cfg

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Config Configuration Recorder. Please note that this resource **does not start** the created recorder automatically.
//
// > **Note:** _Starting_ the Configuration Recorder requires a [delivery channel](https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html) (while delivery channel creation requires Configuration Recorder). This is why [`cfg.RecorderStatus`](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html) is a separate resource.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_configuration_recorder.html.markdown.
type Recorder struct {
	pulumi.CustomResourceState

	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Recording group - see below.
	RecordingGroup RecorderRecordingGroupOutput `pulumi:"recordingGroup"`
	// Amazon Resource Name (ARN) of the IAM role.
	// used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
	// See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewRecorder registers a new resource with the given unique name, arguments, and options.
func NewRecorder(ctx *pulumi.Context,
	name string, args *RecorderArgs, opts ...pulumi.ResourceOption) (*Recorder, error) {
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil {
		args = &RecorderArgs{}
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
	// Amazon Resource Name (ARN) of the IAM role.
	// used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
	// See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn *string `pulumi:"roleArn"`
}

type RecorderState struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringPtrInput
	// Recording group - see below.
	RecordingGroup RecorderRecordingGroupPtrInput
	// Amazon Resource Name (ARN) of the IAM role.
	// used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
	// See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
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
	// Amazon Resource Name (ARN) of the IAM role.
	// used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
	// See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a Recorder resource.
type RecorderArgs struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringPtrInput
	// Recording group - see below.
	RecordingGroup RecorderRecordingGroupPtrInput
	// Amazon Resource Name (ARN) of the IAM role.
	// used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
	// See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn pulumi.StringInput
}

func (RecorderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recorderArgs)(nil)).Elem()
}


// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Kinesis Video Stream resource. Amazon Kinesis Video Streams makes it easy to securely stream video from connected devices to AWS for analytics, machine learning (ML), playback, and other processing.
//
// For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
type VideoStream struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A time stamp that indicates when the stream was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours pulumi.IntPtrOutput `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName pulumi.StringPtrOutput `pulumi:"deviceName"`
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType pulumi.StringPtrOutput `pulumi:"mediaType"`
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The version of the stream.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewVideoStream registers a new resource with the given unique name, arguments, and options.
func NewVideoStream(ctx *pulumi.Context,
	name string, args *VideoStreamArgs, opts ...pulumi.ResourceOption) (*VideoStream, error) {
	if args == nil {
		args = &VideoStreamArgs{}
	}
	var resource VideoStream
	err := ctx.RegisterResource("aws:kinesis/videoStream:VideoStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVideoStream gets an existing VideoStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVideoStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VideoStreamState, opts ...pulumi.ResourceOption) (*VideoStream, error) {
	var resource VideoStream
	err := ctx.ReadResource("aws:kinesis/videoStream:VideoStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VideoStream resources.
type videoStreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn *string `pulumi:"arn"`
	// A time stamp that indicates when the stream was created.
	CreationTime *string `pulumi:"creationTime"`
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours *int `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName *string `pulumi:"deviceName"`
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType *string `pulumi:"mediaType"`
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version of the stream.
	Version *string `pulumi:"version"`
}

type VideoStreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringPtrInput
	// A time stamp that indicates when the stream was created.
	CreationTime pulumi.StringPtrInput
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours pulumi.IntPtrInput
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName pulumi.StringPtrInput
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId pulumi.StringPtrInput
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The version of the stream.
	Version pulumi.StringPtrInput
}

func (VideoStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*videoStreamState)(nil)).Elem()
}

type videoStreamArgs struct {
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours *int `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName *string `pulumi:"deviceName"`
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType *string `pulumi:"mediaType"`
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VideoStream resource.
type VideoStreamArgs struct {
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours pulumi.IntPtrInput
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName pulumi.StringPtrInput
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId pulumi.StringPtrInput
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VideoStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*videoStreamArgs)(nil)).Elem()
}

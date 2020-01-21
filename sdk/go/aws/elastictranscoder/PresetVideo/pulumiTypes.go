// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package PresetVideo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type PresetVideo struct {
	// The display aspect ratio of the video in the output file. Valid values are: `auto`, `1:1`, `4:3`, `3:2`, `16:9`. (Note; to better control resolution and aspect ratio of output videos, we recommend that you use the values `maxWidth`, `maxHeight`, `sizingPolicy`, `paddingPolicy`, and `displayAspectRatio` instead of `resolution` and `aspectRatio`.)
	AspectRatio *string `pulumi:"aspectRatio"`
	// The bit rate of the video stream in the output file, in kilobits/second. You can configure variable bit rate or constant bit rate encoding.
	BitRate *string `pulumi:"bitRate"`
	// The video codec for the output file. Valid values are `gif`, `H.264`, `mpeg2`, `vp8`, and `vp9`.
	Codec *string `pulumi:"codec"`
	// The value that Elastic Transcoder adds to the metadata in the output file. If you set DisplayAspectRatio to auto, Elastic Transcoder chooses an aspect ratio that ensures square pixels. If you specify another option, Elastic Transcoder sets that value in the output file.
	DisplayAspectRatio *string `pulumi:"displayAspectRatio"`
	// Whether to use a fixed value for Video:FixedGOP. Not applicable for containers of type gif. Valid values are true and false. Also known as, Fixed Number of Frames Between Keyframes.
	FixedGop *string `pulumi:"fixedGop"`
	// The frames per second for the video stream in the output file. The following values are valid: `auto`, `10`, `15`, `23.97`, `24`, `25`, `29.97`, `30`, `50`, `60`.
	FrameRate *string `pulumi:"frameRate"`
	// The maximum number of frames between key frames. Not applicable for containers of type gif.
	KeyframesMaxDist *string `pulumi:"keyframesMaxDist"`
	// If you specify auto for FrameRate, Elastic Transcoder uses the frame rate of the input video for the frame rate of the output video, up to the maximum frame rate. If you do not specify a MaxFrameRate, Elastic Transcoder will use a default of 30.
	MaxFrameRate *string `pulumi:"maxFrameRate"`
	// The maximum height of the watermark.
	MaxHeight *string `pulumi:"maxHeight"`
	// The maximum width of the watermark.
	MaxWidth *string `pulumi:"maxWidth"`
	// When you set PaddingPolicy to Pad, Elastic Transcoder might add black bars to the top and bottom and/or left and right sides of the output video to make the total size of the output video match the values that you specified for `maxWidth` and `maxHeight`.
	PaddingPolicy *string `pulumi:"paddingPolicy"`
	// The width and height of the video in the output file, in pixels. Valid values are `auto` and `widthxheight`. (see note for `aspectRatio`)
	Resolution *string `pulumi:"resolution"`
	// A value that controls scaling of the watermark. Valid values are: `Fit`, `Stretch`, `ShrinkToFit`
	SizingPolicy *string `pulumi:"sizingPolicy"`
}

type PresetVideoInput interface {
	pulumi.Input

	ToPresetVideoOutput() PresetVideoOutput
	ToPresetVideoOutputWithContext(context.Context) PresetVideoOutput
}

type PresetVideoArgs struct {
	// The display aspect ratio of the video in the output file. Valid values are: `auto`, `1:1`, `4:3`, `3:2`, `16:9`. (Note; to better control resolution and aspect ratio of output videos, we recommend that you use the values `maxWidth`, `maxHeight`, `sizingPolicy`, `paddingPolicy`, and `displayAspectRatio` instead of `resolution` and `aspectRatio`.)
	AspectRatio pulumi.StringPtrInput `pulumi:"aspectRatio"`
	// The bit rate of the video stream in the output file, in kilobits/second. You can configure variable bit rate or constant bit rate encoding.
	BitRate pulumi.StringPtrInput `pulumi:"bitRate"`
	// The video codec for the output file. Valid values are `gif`, `H.264`, `mpeg2`, `vp8`, and `vp9`.
	Codec pulumi.StringPtrInput `pulumi:"codec"`
	// The value that Elastic Transcoder adds to the metadata in the output file. If you set DisplayAspectRatio to auto, Elastic Transcoder chooses an aspect ratio that ensures square pixels. If you specify another option, Elastic Transcoder sets that value in the output file.
	DisplayAspectRatio pulumi.StringPtrInput `pulumi:"displayAspectRatio"`
	// Whether to use a fixed value for Video:FixedGOP. Not applicable for containers of type gif. Valid values are true and false. Also known as, Fixed Number of Frames Between Keyframes.
	FixedGop pulumi.StringPtrInput `pulumi:"fixedGop"`
	// The frames per second for the video stream in the output file. The following values are valid: `auto`, `10`, `15`, `23.97`, `24`, `25`, `29.97`, `30`, `50`, `60`.
	FrameRate pulumi.StringPtrInput `pulumi:"frameRate"`
	// The maximum number of frames between key frames. Not applicable for containers of type gif.
	KeyframesMaxDist pulumi.StringPtrInput `pulumi:"keyframesMaxDist"`
	// If you specify auto for FrameRate, Elastic Transcoder uses the frame rate of the input video for the frame rate of the output video, up to the maximum frame rate. If you do not specify a MaxFrameRate, Elastic Transcoder will use a default of 30.
	MaxFrameRate pulumi.StringPtrInput `pulumi:"maxFrameRate"`
	// The maximum height of the watermark.
	MaxHeight pulumi.StringPtrInput `pulumi:"maxHeight"`
	// The maximum width of the watermark.
	MaxWidth pulumi.StringPtrInput `pulumi:"maxWidth"`
	// When you set PaddingPolicy to Pad, Elastic Transcoder might add black bars to the top and bottom and/or left and right sides of the output video to make the total size of the output video match the values that you specified for `maxWidth` and `maxHeight`.
	PaddingPolicy pulumi.StringPtrInput `pulumi:"paddingPolicy"`
	// The width and height of the video in the output file, in pixels. Valid values are `auto` and `widthxheight`. (see note for `aspectRatio`)
	Resolution pulumi.StringPtrInput `pulumi:"resolution"`
	// A value that controls scaling of the watermark. Valid values are: `Fit`, `Stretch`, `ShrinkToFit`
	SizingPolicy pulumi.StringPtrInput `pulumi:"sizingPolicy"`
}

func (PresetVideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PresetVideo)(nil)).Elem()
}

func (i PresetVideoArgs) ToPresetVideoOutput() PresetVideoOutput {
	return i.ToPresetVideoOutputWithContext(context.Background())
}

func (i PresetVideoArgs) ToPresetVideoOutputWithContext(ctx context.Context) PresetVideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetVideoOutput)
}

func (i PresetVideoArgs) ToPresetVideoPtrOutput() PresetVideoPtrOutput {
	return i.ToPresetVideoPtrOutputWithContext(context.Background())
}

func (i PresetVideoArgs) ToPresetVideoPtrOutputWithContext(ctx context.Context) PresetVideoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetVideoOutput).ToPresetVideoPtrOutputWithContext(ctx)
}

type PresetVideoPtrInput interface {
	pulumi.Input

	ToPresetVideoPtrOutput() PresetVideoPtrOutput
	ToPresetVideoPtrOutputWithContext(context.Context) PresetVideoPtrOutput
}

type presetVideoPtrType PresetVideoArgs

func PresetVideoPtr(v *PresetVideoArgs) PresetVideoPtrInput {	return (*presetVideoPtrType)(v)
}

func (*presetVideoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PresetVideo)(nil)).Elem()
}

func (i *presetVideoPtrType) ToPresetVideoPtrOutput() PresetVideoPtrOutput {
	return i.ToPresetVideoPtrOutputWithContext(context.Background())
}

func (i *presetVideoPtrType) ToPresetVideoPtrOutputWithContext(ctx context.Context) PresetVideoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetVideoPtrOutput)
}

type PresetVideoOutput struct { *pulumi.OutputState }

func (PresetVideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PresetVideo)(nil)).Elem()
}

func (o PresetVideoOutput) ToPresetVideoOutput() PresetVideoOutput {
	return o
}

func (o PresetVideoOutput) ToPresetVideoOutputWithContext(ctx context.Context) PresetVideoOutput {
	return o
}

func (o PresetVideoOutput) ToPresetVideoPtrOutput() PresetVideoPtrOutput {
	return o.ToPresetVideoPtrOutputWithContext(context.Background())
}

func (o PresetVideoOutput) ToPresetVideoPtrOutputWithContext(ctx context.Context) PresetVideoPtrOutput {
	return o.ApplyT(func(v PresetVideo) *PresetVideo {
		return &v
	}).(PresetVideoPtrOutput)
}
// The display aspect ratio of the video in the output file. Valid values are: `auto`, `1:1`, `4:3`, `3:2`, `16:9`. (Note; to better control resolution and aspect ratio of output videos, we recommend that you use the values `maxWidth`, `maxHeight`, `sizingPolicy`, `paddingPolicy`, and `displayAspectRatio` instead of `resolution` and `aspectRatio`.)
func (o PresetVideoOutput) AspectRatio() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.AspectRatio }).(pulumi.StringPtrOutput)
}

// The bit rate of the video stream in the output file, in kilobits/second. You can configure variable bit rate or constant bit rate encoding.
func (o PresetVideoOutput) BitRate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.BitRate }).(pulumi.StringPtrOutput)
}

// The video codec for the output file. Valid values are `gif`, `H.264`, `mpeg2`, `vp8`, and `vp9`.
func (o PresetVideoOutput) Codec() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.Codec }).(pulumi.StringPtrOutput)
}

// The value that Elastic Transcoder adds to the metadata in the output file. If you set DisplayAspectRatio to auto, Elastic Transcoder chooses an aspect ratio that ensures square pixels. If you specify another option, Elastic Transcoder sets that value in the output file.
func (o PresetVideoOutput) DisplayAspectRatio() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.DisplayAspectRatio }).(pulumi.StringPtrOutput)
}

// Whether to use a fixed value for Video:FixedGOP. Not applicable for containers of type gif. Valid values are true and false. Also known as, Fixed Number of Frames Between Keyframes.
func (o PresetVideoOutput) FixedGop() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.FixedGop }).(pulumi.StringPtrOutput)
}

// The frames per second for the video stream in the output file. The following values are valid: `auto`, `10`, `15`, `23.97`, `24`, `25`, `29.97`, `30`, `50`, `60`.
func (o PresetVideoOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

// The maximum number of frames between key frames. Not applicable for containers of type gif.
func (o PresetVideoOutput) KeyframesMaxDist() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.KeyframesMaxDist }).(pulumi.StringPtrOutput)
}

// If you specify auto for FrameRate, Elastic Transcoder uses the frame rate of the input video for the frame rate of the output video, up to the maximum frame rate. If you do not specify a MaxFrameRate, Elastic Transcoder will use a default of 30.
func (o PresetVideoOutput) MaxFrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.MaxFrameRate }).(pulumi.StringPtrOutput)
}

// The maximum height of the watermark.
func (o PresetVideoOutput) MaxHeight() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.MaxHeight }).(pulumi.StringPtrOutput)
}

// The maximum width of the watermark.
func (o PresetVideoOutput) MaxWidth() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.MaxWidth }).(pulumi.StringPtrOutput)
}

// When you set PaddingPolicy to Pad, Elastic Transcoder might add black bars to the top and bottom and/or left and right sides of the output video to make the total size of the output video match the values that you specified for `maxWidth` and `maxHeight`.
func (o PresetVideoOutput) PaddingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.PaddingPolicy }).(pulumi.StringPtrOutput)
}

// The width and height of the video in the output file, in pixels. Valid values are `auto` and `widthxheight`. (see note for `aspectRatio`)
func (o PresetVideoOutput) Resolution() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.Resolution }).(pulumi.StringPtrOutput)
}

// A value that controls scaling of the watermark. Valid values are: `Fit`, `Stretch`, `ShrinkToFit`
func (o PresetVideoOutput) SizingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.SizingPolicy }).(pulumi.StringPtrOutput)
}

type PresetVideoPtrOutput struct { *pulumi.OutputState}

func (PresetVideoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PresetVideo)(nil)).Elem()
}

func (o PresetVideoPtrOutput) ToPresetVideoPtrOutput() PresetVideoPtrOutput {
	return o
}

func (o PresetVideoPtrOutput) ToPresetVideoPtrOutputWithContext(ctx context.Context) PresetVideoPtrOutput {
	return o
}

func (o PresetVideoPtrOutput) Elem() PresetVideoOutput {
	return o.ApplyT(func (v *PresetVideo) PresetVideo { return *v }).(PresetVideoOutput)
}

// The display aspect ratio of the video in the output file. Valid values are: `auto`, `1:1`, `4:3`, `3:2`, `16:9`. (Note; to better control resolution and aspect ratio of output videos, we recommend that you use the values `maxWidth`, `maxHeight`, `sizingPolicy`, `paddingPolicy`, and `displayAspectRatio` instead of `resolution` and `aspectRatio`.)
func (o PresetVideoPtrOutput) AspectRatio() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.AspectRatio }).(pulumi.StringPtrOutput)
}

// The bit rate of the video stream in the output file, in kilobits/second. You can configure variable bit rate or constant bit rate encoding.
func (o PresetVideoPtrOutput) BitRate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.BitRate }).(pulumi.StringPtrOutput)
}

// The video codec for the output file. Valid values are `gif`, `H.264`, `mpeg2`, `vp8`, and `vp9`.
func (o PresetVideoPtrOutput) Codec() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.Codec }).(pulumi.StringPtrOutput)
}

// The value that Elastic Transcoder adds to the metadata in the output file. If you set DisplayAspectRatio to auto, Elastic Transcoder chooses an aspect ratio that ensures square pixels. If you specify another option, Elastic Transcoder sets that value in the output file.
func (o PresetVideoPtrOutput) DisplayAspectRatio() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.DisplayAspectRatio }).(pulumi.StringPtrOutput)
}

// Whether to use a fixed value for Video:FixedGOP. Not applicable for containers of type gif. Valid values are true and false. Also known as, Fixed Number of Frames Between Keyframes.
func (o PresetVideoPtrOutput) FixedGop() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.FixedGop }).(pulumi.StringPtrOutput)
}

// The frames per second for the video stream in the output file. The following values are valid: `auto`, `10`, `15`, `23.97`, `24`, `25`, `29.97`, `30`, `50`, `60`.
func (o PresetVideoPtrOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

// The maximum number of frames between key frames. Not applicable for containers of type gif.
func (o PresetVideoPtrOutput) KeyframesMaxDist() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.KeyframesMaxDist }).(pulumi.StringPtrOutput)
}

// If you specify auto for FrameRate, Elastic Transcoder uses the frame rate of the input video for the frame rate of the output video, up to the maximum frame rate. If you do not specify a MaxFrameRate, Elastic Transcoder will use a default of 30.
func (o PresetVideoPtrOutput) MaxFrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.MaxFrameRate }).(pulumi.StringPtrOutput)
}

// The maximum height of the watermark.
func (o PresetVideoPtrOutput) MaxHeight() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.MaxHeight }).(pulumi.StringPtrOutput)
}

// The maximum width of the watermark.
func (o PresetVideoPtrOutput) MaxWidth() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.MaxWidth }).(pulumi.StringPtrOutput)
}

// When you set PaddingPolicy to Pad, Elastic Transcoder might add black bars to the top and bottom and/or left and right sides of the output video to make the total size of the output video match the values that you specified for `maxWidth` and `maxHeight`.
func (o PresetVideoPtrOutput) PaddingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.PaddingPolicy }).(pulumi.StringPtrOutput)
}

// The width and height of the video in the output file, in pixels. Valid values are `auto` and `widthxheight`. (see note for `aspectRatio`)
func (o PresetVideoPtrOutput) Resolution() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.Resolution }).(pulumi.StringPtrOutput)
}

// A value that controls scaling of the watermark. Valid values are: `Fit`, `Stretch`, `ShrinkToFit`
func (o PresetVideoPtrOutput) SizingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PresetVideo) *string { return v.SizingPolicy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PresetVideoOutput{})
	pulumi.RegisterOutputType(PresetVideoPtrOutput{})
}

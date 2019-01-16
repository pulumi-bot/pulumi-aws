# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Preset(pulumi.CustomResource):
    """
    Provides an Elastic Transcoder preset resource.
    """
    def __init__(__self__, __name__, __opts__=None, audio=None, audio_codec_options=None, container=None, description=None, name=None, thumbnails=None, type=None, video=None, video_codec_options=None, video_watermarks=None):
        """Create a Preset resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['audio'] = audio

        __props__['audio_codec_options'] = audio_codec_options

        if not container:
            raise TypeError('Missing required property container')
        __props__['container'] = container

        __props__['description'] = description

        __props__['name'] = name

        __props__['thumbnails'] = thumbnails

        __props__['type'] = type

        __props__['video'] = video

        __props__['video_codec_options'] = video_codec_options

        __props__['video_watermarks'] = video_watermarks

        __props__['arn'] = None

        super(Preset, __self__).__init__(
            'aws:elastictranscoder/preset:Preset',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


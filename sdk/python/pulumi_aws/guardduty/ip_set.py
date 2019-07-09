# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IPSet(pulumi.CustomResource):
    activate: pulumi.Output[bool]
    """
    Specifies whether GuardDuty is to start using the uploaded IPSet.
    """
    detector_id: pulumi.Output[str]
    """
    The detector ID of the GuardDuty.
    """
    format: pulumi.Output[str]
    """
    The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
    """
    location: pulumi.Output[str]
    """
    The URI of the file that contains the IPSet.
    """
    name: pulumi.Output[str]
    """
    The friendly name to identify the IPSet.
    """
    def __init__(__self__, resource_name, opts=None, activate=None, detector_id=None, format=None, location=None, name=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage a GuardDuty IPSet.
        
        > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage IPSets. IPSets that are uploaded by the master account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-ip-set.html)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] activate: Specifies whether GuardDuty is to start using the uploaded IPSet.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty.
        :param pulumi.Input[str] format: The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
        :param pulumi.Input[str] location: The URI of the file that contains the IPSet.
        :param pulumi.Input[str] name: The friendly name to identify the IPSet.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/guardduty_ipset.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if activate is None:
            raise TypeError("Missing required property 'activate'")
        __props__['activate'] = activate

        if detector_id is None:
            raise TypeError("Missing required property 'detector_id'")
        __props__['detector_id'] = detector_id

        if format is None:
            raise TypeError("Missing required property 'format'")
        __props__['format'] = format

        if location is None:
            raise TypeError("Missing required property 'location'")
        __props__['location'] = location

        __props__['name'] = name

        super(IPSet, __self__).__init__(
            'aws:guardduty/iPSet:IPSet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


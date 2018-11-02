# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class MailFrom(pulumi.CustomResource):
    """
    Provides an SES domain MAIL FROM resource.
    
    ~> **NOTE:** For the MAIL FROM domain to be fully usable, this resource should be paired with the [aws_ses_domain_identity resource](https://www.terraform.io/docs/providers/aws/r/ses_domain_identity.html). To validate the MAIL FROM domain, a DNS MX record is required. To pass SPF checks, a DNS TXT record may also be required. See the [Amazon SES MAIL FROM documentation](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-set.html) for more information.
    """
    def __init__(__self__, __name__, __opts__=None, behavior_on_mx_failure=None, domain=None, mail_from_domain=None):
        """Create a MailFrom resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['behaviorOnMxFailure'] = behavior_on_mx_failure

        if not domain:
            raise TypeError('Missing required property domain')
        __props__['domain'] = domain

        if not mail_from_domain:
            raise TypeError('Missing required property mail_from_domain')
        __props__['mailFromDomain'] = mail_from_domain

        super(MailFrom, __self__).__init__(
            'aws:ses/mailFrom:MailFrom',
            __name__,
            __props__,
            __opts__)


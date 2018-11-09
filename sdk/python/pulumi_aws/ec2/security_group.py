# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class SecurityGroup(pulumi.CustomResource):
    """
    Provides a security group resource.
    
    ~> **NOTE on Security Groups and Security Group Rules:** Terraform currently
    provides both a standalone Security Group Rule resource (a single `ingress` or
    `egress` rule), and a Security Group resource with `ingress` and `egress` rules
    defined in-line. At this time you cannot use a Security Group with in-line rules
    in conjunction with any Security Group Rule resources. Doing so will cause
    a conflict of rule settings and will overwrite rules.
    
    ~> **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
    """
    def __init__(__self__, __name__, __opts__=None, description=None, egress=None, ingress=None, name=None, name_prefix=None, revoke_rules_on_delete=None, tags=None, vpc_id=None):
        """Create a SecurityGroup resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        description = 'Managed by Pulumi'
        __props__['description'] = description

        __props__['egress'] = egress

        __props__['ingress'] = ingress

        __props__['name'] = name

        __props__['namePrefix'] = name_prefix

        __props__['revokeRulesOnDelete'] = revoke_rules_on_delete

        __props__['tags'] = tags

        __props__['vpcId'] = vpc_id

        __props__['arn'] = None
        __props__['owner_id'] = None

        super(SecurityGroup, __self__).__init__(
            'aws:ec2/securityGroup:SecurityGroup',
            __name__,
            __props__,
            __opts__)


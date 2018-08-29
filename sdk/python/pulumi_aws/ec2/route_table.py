# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class RouteTable(pulumi.CustomResource):
    """
    Provides a resource to create a VPC routing table.
    
    ~> **NOTE on Route Tables and Routes:** Terraform currently
    provides both a standalone Route resource and a Route Table resource with routes
    defined in-line. At this time you cannot use a Route Table with in-line routes
    in conjunction with any Route resources. Doing so will cause
    a conflict of rule settings and will overwrite rules.
    
    ~> **NOTE on `gateway_id` and `nat_gateway_id`:** The AWS API is very forgiving with these two
    attributes and the `aws_route_table` resource can be created with a NAT ID specified as a Gateway ID attribute.
    This _will_ lead to a permanent diff between your configuration and statefile, as the API returns the correct
    parameters in the returned route table. If you're experiencing constant diffs in your `aws_route_table` resources,
    the first thing to check is whether or not you're specifying a NAT ID instead of a Gateway ID, or vice-versa.
    
    ~> **NOTE on `propagating_vgws` and the `aws_vpn_gateway_route_propagation` resource:**
    If the `propagating_vgws` argument is present, it's not supported to _also_
    define route propagations using `aws_vpn_gateway_route_propagation`, since
    this resource will delete any propagating gateways not explicitly listed in
    `propagating_vgws`. Omit this argument when defining route propagation using
    the separate resource.
    """
    def __init__(__self__, __name__, __opts__=None, propagating_vgws=None, routes=None, tags=None, vpc_id=None):
        """Create a RouteTable resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if propagating_vgws and not isinstance(propagating_vgws, list):
            raise TypeError('Expected property propagating_vgws to be a list')
        __self__.propagating_vgws = propagating_vgws
        """
        A list of virtual gateways for propagation.
        """
        __props__['propagatingVgws'] = propagating_vgws

        if routes and not isinstance(routes, list):
            raise TypeError('Expected property routes to be a list')
        __self__.routes = routes
        """
        A list of route objects. Their keys are documented below.
        """
        __props__['routes'] = routes

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        if not vpc_id:
            raise TypeError('Missing required property vpc_id')
        elif not isinstance(vpc_id, basestring):
            raise TypeError('Expected property vpc_id to be a basestring')
        __self__.vpc_id = vpc_id
        """
        The VPC ID.
        """
        __props__['vpcId'] = vpc_id

        super(RouteTable, __self__).__init__(
            'aws:ec2/routeTable:RouteTable',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'propagatingVgws' in outs:
            self.propagating_vgws = outs['propagatingVgws']
        if 'routes' in outs:
            self.routes = outs['routes']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'vpcId' in outs:
            self.vpc_id = outs['vpcId']

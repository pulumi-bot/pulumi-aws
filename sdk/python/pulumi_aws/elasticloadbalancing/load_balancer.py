# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("aws.elasticloadbalancing.LoadBalancer has been deprecated in favour of aws.elb.LoadBalancer", DeprecationWarning)
class LoadBalancer(pulumi.CustomResource):
    access_logs: pulumi.Output[dict]
    """
    An Access Logs block. Access Logs documented below.

      * `bucket` (`str`) - The S3 bucket name to store the logs in.
      * `bucket_prefix` (`str`) - The S3 bucket prefix. Logs are stored in the root if not configured.
      * `enabled` (`bool`) - Boolean to enable / disable `access_logs`. Default is `true`
      * `interval` (`float`) - The publishing interval in minutes. Default: 60 minutes.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the ELB
    """
    availability_zones: pulumi.Output[list]
    """
    The AZ's to serve traffic in.
    """
    connection_draining: pulumi.Output[bool]
    """
    Boolean to enable connection draining. Default: `false`
    """
    connection_draining_timeout: pulumi.Output[float]
    """
    The time in seconds to allow for connections to drain. Default: `300`
    """
    cross_zone_load_balancing: pulumi.Output[bool]
    """
    Enable cross-zone load balancing. Default: `true`
    """
    dns_name: pulumi.Output[str]
    """
    The DNS name of the ELB
    """
    health_check: pulumi.Output[dict]
    """
    A health_check block. Health Check documented below.

      * `healthyThreshold` (`float`) - The number of checks before the instance is declared healthy.
      * `interval` (`float`) - The interval between checks.
      * `target` (`str`) - The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
        values are:
        * `HTTP`, `HTTPS` - PORT and PATH are required
        * `TCP`, `SSL` - PORT is required, PATH is not supported
      * `timeout` (`float`) - The length of time before the check times out.
      * `unhealthyThreshold` (`float`) - The number of checks before the instance is declared unhealthy.
    """
    idle_timeout: pulumi.Output[float]
    """
    The time in seconds that the connection is allowed to be idle. Default: `60`
    """
    instances: pulumi.Output[list]
    """
    A list of instance ids to place in the ELB pool.
    """
    internal: pulumi.Output[bool]
    """
    If true, ELB will be an internal ELB.
    """
    listeners: pulumi.Output[list]
    """
    A list of listener blocks. Listeners documented below.

      * `instance_port` (`float`) - The port on the instance to route to
      * `instanceProtocol` (`str`) - The protocol to use to the instance. Valid
        values are `HTTP`, `HTTPS`, `TCP`, or `SSL`
      * `lb_port` (`float`) - The port to listen on for the load balancer
      * `lbProtocol` (`str`) - The protocol to listen on. Valid values are `HTTP`,
        `HTTPS`, `TCP`, or `SSL`
      * `sslCertificateId` (`str`) - The ARN of an SSL certificate you have
        uploaded to AWS IAM. **Note ECDSA-specific restrictions below.  Only valid when `lb_protocol` is either HTTPS or SSL**
    """
    name: pulumi.Output[str]
    """
    The name of the ELB. By default generated by this provider.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified
    prefix. Conflicts with `name`.
    """
    security_groups: pulumi.Output[list]
    """
    A list of security group IDs to assign to the ELB.
    Only valid if creating an ELB within a VPC
    """
    source_security_group: pulumi.Output[str]
    """
    The name of the security group that you can use as
    part of your inbound rules for your load balancer's back-end application
    instances. Use this for Classic or Default VPC only.
    """
    source_security_group_id: pulumi.Output[str]
    """
    The ID of the security group that you can use as
    part of your inbound rules for your load balancer's back-end application
    instances. Only available on ELBs launched in a VPC.
    """
    subnets: pulumi.Output[list]
    """
    A list of subnet IDs to attach to the ELB.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    zone_id: pulumi.Output[str]
    """
    The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
    """
    warnings.warn("aws.elasticloadbalancing.LoadBalancer has been deprecated in favour of aws.elb.LoadBalancer", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, access_logs=None, availability_zones=None, connection_draining=None, connection_draining_timeout=None, cross_zone_load_balancing=None, health_check=None, idle_timeout=None, instances=None, internal=None, listeners=None, name=None, name_prefix=None, security_groups=None, source_security_group=None, subnets=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Elastic Load Balancer resource, also known as a "Classic
        Load Balancer" after the release of
        [Application/Network Load Balancers](https://www.terraform.io/docs/providers/aws/r/lb.html).

        > **NOTE on ELB Instances and ELB Attachments:** This provider currently
        provides both a standalone ELB Attachment resource
        (describing an instance attached to an ELB), and an ELB resource with
        `instances` defined in-line. At this time you cannot use an ELB with in-line
        instances in conjunction with a ELB Attachment resources. Doing so will cause a
        conflict and will overwrite attachments.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new load balancer
        bar = aws.elb.LoadBalancer("bar",
            access_logs={
                "bucket": "foo",
                "bucket_prefix": "bar",
                "interval": 60,
            },
            availability_zones=[
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            ],
            connection_draining=True,
            connection_draining_timeout=400,
            cross_zone_load_balancing=True,
            health_check={
                "healthyThreshold": 2,
                "interval": 30,
                "target": "HTTP:8000/",
                "timeout": 3,
                "unhealthyThreshold": 2,
            },
            idle_timeout=400,
            instances=[aws_instance["foo"]["id"]],
            listeners=[
                {
                    "instance_port": 8000,
                    "instanceProtocol": "http",
                    "lb_port": 80,
                    "lbProtocol": "http",
                },
                {
                    "instance_port": 8000,
                    "instanceProtocol": "http",
                    "lb_port": 443,
                    "lbProtocol": "https",
                    "sslCertificateId": "arn:aws:iam::123456789012:server-certificate/certName",
                },
            ],
            tags={
                "Name": "foobar-elb",
            })
        ```

        ## Note on ECDSA Key Algorithm

        If the ARN of the `ssl_certificate_id` that is pointed to references a
        certificate that was signed by an ECDSA key, note that ELB only supports the
        P256 and P384 curves.  Using a certificate signed by a key using a different
        curve could produce the error `ERR_SSL_VERSION_OR_CIPHER_MISMATCH` in your
        browser.

        Deprecated: aws.elasticloadbalancing.LoadBalancer has been deprecated in favour of aws.elb.LoadBalancer

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_logs: An Access Logs block. Access Logs documented below.
        :param pulumi.Input[list] availability_zones: The AZ's to serve traffic in.
        :param pulumi.Input[bool] connection_draining: Boolean to enable connection draining. Default: `false`
        :param pulumi.Input[float] connection_draining_timeout: The time in seconds to allow for connections to drain. Default: `300`
        :param pulumi.Input[bool] cross_zone_load_balancing: Enable cross-zone load balancing. Default: `true`
        :param pulumi.Input[dict] health_check: A health_check block. Health Check documented below.
        :param pulumi.Input[float] idle_timeout: The time in seconds that the connection is allowed to be idle. Default: `60`
        :param pulumi.Input[list] instances: A list of instance ids to place in the ELB pool.
        :param pulumi.Input[bool] internal: If true, ELB will be an internal ELB.
        :param pulumi.Input[list] listeners: A list of listener blocks. Listeners documented below.
        :param pulumi.Input[str] name: The name of the ELB. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[list] security_groups: A list of security group IDs to assign to the ELB.
               Only valid if creating an ELB within a VPC
        :param pulumi.Input[str] source_security_group: The name of the security group that you can use as
               part of your inbound rules for your load balancer's back-end application
               instances. Use this for Classic or Default VPC only.
        :param pulumi.Input[list] subnets: A list of subnet IDs to attach to the ELB.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **access_logs** object supports the following:

          * `bucket` (`pulumi.Input[str]`) - The S3 bucket name to store the logs in.
          * `bucket_prefix` (`pulumi.Input[str]`) - The S3 bucket prefix. Logs are stored in the root if not configured.
          * `enabled` (`pulumi.Input[bool]`) - Boolean to enable / disable `access_logs`. Default is `true`
          * `interval` (`pulumi.Input[float]`) - The publishing interval in minutes. Default: 60 minutes.

        The **health_check** object supports the following:

          * `healthyThreshold` (`pulumi.Input[float]`) - The number of checks before the instance is declared healthy.
          * `interval` (`pulumi.Input[float]`) - The interval between checks.
          * `target` (`pulumi.Input[str]`) - The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
            values are:
            * `HTTP`, `HTTPS` - PORT and PATH are required
            * `TCP`, `SSL` - PORT is required, PATH is not supported
          * `timeout` (`pulumi.Input[float]`) - The length of time before the check times out.
          * `unhealthyThreshold` (`pulumi.Input[float]`) - The number of checks before the instance is declared unhealthy.

        The **listeners** object supports the following:

          * `instance_port` (`pulumi.Input[float]`) - The port on the instance to route to
          * `instanceProtocol` (`pulumi.Input[str]`) - The protocol to use to the instance. Valid
            values are `HTTP`, `HTTPS`, `TCP`, or `SSL`
          * `lb_port` (`pulumi.Input[float]`) - The port to listen on for the load balancer
          * `lbProtocol` (`pulumi.Input[str]`) - The protocol to listen on. Valid values are `HTTP`,
            `HTTPS`, `TCP`, or `SSL`
          * `sslCertificateId` (`pulumi.Input[str]`) - The ARN of an SSL certificate you have
            uploaded to AWS IAM. **Note ECDSA-specific restrictions below.  Only valid when `lb_protocol` is either HTTPS or SSL**
        """
        pulumi.log.warn("LoadBalancer is deprecated: aws.elasticloadbalancing.LoadBalancer has been deprecated in favour of aws.elb.LoadBalancer")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['access_logs'] = access_logs
            __props__['availability_zones'] = availability_zones
            __props__['connection_draining'] = connection_draining
            __props__['connection_draining_timeout'] = connection_draining_timeout
            __props__['cross_zone_load_balancing'] = cross_zone_load_balancing
            __props__['health_check'] = health_check
            __props__['idle_timeout'] = idle_timeout
            __props__['instances'] = instances
            __props__['internal'] = internal
            if listeners is None:
                raise TypeError("Missing required property 'listeners'")
            __props__['listeners'] = listeners
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['security_groups'] = security_groups
            __props__['source_security_group'] = source_security_group
            __props__['subnets'] = subnets
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['dns_name'] = None
            __props__['source_security_group_id'] = None
            __props__['zone_id'] = None
        super(LoadBalancer, __self__).__init__(
            'aws:elasticloadbalancing/loadBalancer:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_logs=None, arn=None, availability_zones=None, connection_draining=None, connection_draining_timeout=None, cross_zone_load_balancing=None, dns_name=None, health_check=None, idle_timeout=None, instances=None, internal=None, listeners=None, name=None, name_prefix=None, security_groups=None, source_security_group=None, source_security_group_id=None, subnets=None, tags=None, zone_id=None):
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_logs: An Access Logs block. Access Logs documented below.
        :param pulumi.Input[str] arn: The ARN of the ELB
        :param pulumi.Input[list] availability_zones: The AZ's to serve traffic in.
        :param pulumi.Input[bool] connection_draining: Boolean to enable connection draining. Default: `false`
        :param pulumi.Input[float] connection_draining_timeout: The time in seconds to allow for connections to drain. Default: `300`
        :param pulumi.Input[bool] cross_zone_load_balancing: Enable cross-zone load balancing. Default: `true`
        :param pulumi.Input[str] dns_name: The DNS name of the ELB
        :param pulumi.Input[dict] health_check: A health_check block. Health Check documented below.
        :param pulumi.Input[float] idle_timeout: The time in seconds that the connection is allowed to be idle. Default: `60`
        :param pulumi.Input[list] instances: A list of instance ids to place in the ELB pool.
        :param pulumi.Input[bool] internal: If true, ELB will be an internal ELB.
        :param pulumi.Input[list] listeners: A list of listener blocks. Listeners documented below.
        :param pulumi.Input[str] name: The name of the ELB. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[list] security_groups: A list of security group IDs to assign to the ELB.
               Only valid if creating an ELB within a VPC
        :param pulumi.Input[str] source_security_group: The name of the security group that you can use as
               part of your inbound rules for your load balancer's back-end application
               instances. Use this for Classic or Default VPC only.
        :param pulumi.Input[str] source_security_group_id: The ID of the security group that you can use as
               part of your inbound rules for your load balancer's back-end application
               instances. Only available on ELBs launched in a VPC.
        :param pulumi.Input[list] subnets: A list of subnet IDs to attach to the ELB.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] zone_id: The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)

        The **access_logs** object supports the following:

          * `bucket` (`pulumi.Input[str]`) - The S3 bucket name to store the logs in.
          * `bucket_prefix` (`pulumi.Input[str]`) - The S3 bucket prefix. Logs are stored in the root if not configured.
          * `enabled` (`pulumi.Input[bool]`) - Boolean to enable / disable `access_logs`. Default is `true`
          * `interval` (`pulumi.Input[float]`) - The publishing interval in minutes. Default: 60 minutes.

        The **health_check** object supports the following:

          * `healthyThreshold` (`pulumi.Input[float]`) - The number of checks before the instance is declared healthy.
          * `interval` (`pulumi.Input[float]`) - The interval between checks.
          * `target` (`pulumi.Input[str]`) - The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
            values are:
            * `HTTP`, `HTTPS` - PORT and PATH are required
            * `TCP`, `SSL` - PORT is required, PATH is not supported
          * `timeout` (`pulumi.Input[float]`) - The length of time before the check times out.
          * `unhealthyThreshold` (`pulumi.Input[float]`) - The number of checks before the instance is declared unhealthy.

        The **listeners** object supports the following:

          * `instance_port` (`pulumi.Input[float]`) - The port on the instance to route to
          * `instanceProtocol` (`pulumi.Input[str]`) - The protocol to use to the instance. Valid
            values are `HTTP`, `HTTPS`, `TCP`, or `SSL`
          * `lb_port` (`pulumi.Input[float]`) - The port to listen on for the load balancer
          * `lbProtocol` (`pulumi.Input[str]`) - The protocol to listen on. Valid values are `HTTP`,
            `HTTPS`, `TCP`, or `SSL`
          * `sslCertificateId` (`pulumi.Input[str]`) - The ARN of an SSL certificate you have
            uploaded to AWS IAM. **Note ECDSA-specific restrictions below.  Only valid when `lb_protocol` is either HTTPS or SSL**
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_logs"] = access_logs
        __props__["arn"] = arn
        __props__["availability_zones"] = availability_zones
        __props__["connection_draining"] = connection_draining
        __props__["connection_draining_timeout"] = connection_draining_timeout
        __props__["cross_zone_load_balancing"] = cross_zone_load_balancing
        __props__["dns_name"] = dns_name
        __props__["health_check"] = health_check
        __props__["idle_timeout"] = idle_timeout
        __props__["instances"] = instances
        __props__["internal"] = internal
        __props__["listeners"] = listeners
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["security_groups"] = security_groups
        __props__["source_security_group"] = source_security_group
        __props__["source_security_group_id"] = source_security_group_id
        __props__["subnets"] = subnets
        __props__["tags"] = tags
        __props__["zone_id"] = zone_id
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


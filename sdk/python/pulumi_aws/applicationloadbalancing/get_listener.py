# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("aws.applicationloadbalancing.getListener has been deprecated in favor of aws.alb.getListener", DeprecationWarning)

class GetListenerResult:
    """
    A collection of values returned by getListener.
    """
    def __init__(__self__, arn=None, certificate_arn=None, default_actions=None, id=None, load_balancer_arn=None, port=None, protocol=None, ssl_policy=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if certificate_arn and not isinstance(certificate_arn, str):
            raise TypeError("Expected argument 'certificate_arn' to be a str")
        __self__.certificate_arn = certificate_arn
        if default_actions and not isinstance(default_actions, list):
            raise TypeError("Expected argument 'default_actions' to be a list")
        __self__.default_actions = default_actions
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if load_balancer_arn and not isinstance(load_balancer_arn, str):
            raise TypeError("Expected argument 'load_balancer_arn' to be a str")
        __self__.load_balancer_arn = load_balancer_arn
        if port and not isinstance(port, float):
            raise TypeError("Expected argument 'port' to be a float")
        __self__.port = port
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        __self__.protocol = protocol
        if ssl_policy and not isinstance(ssl_policy, str):
            raise TypeError("Expected argument 'ssl_policy' to be a str")
        __self__.ssl_policy = ssl_policy


class AwaitableGetListenerResult(GetListenerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetListenerResult(
            arn=self.arn,
            certificate_arn=self.certificate_arn,
            default_actions=self.default_actions,
            id=self.id,
            load_balancer_arn=self.load_balancer_arn,
            port=self.port,
            protocol=self.protocol,
            ssl_policy=self.ssl_policy)


def get_listener(arn=None,load_balancer_arn=None,port=None,opts=None):
    """
    > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.

    Provides information about a Load Balancer Listener.

    This data source can prove useful when a module accepts an LB Listener as an
    input variable and needs to know the LB it is attached to, or other
    information specific to the listener in question.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    listener_arn = config.require_object("listenerArn")
    listener = aws.lb.get_listener(arn=listener_arn)
    selected = aws.lb.get_load_balancer(name="default-public")
    selected443 = aws.lb.get_listener(load_balancer_arn=selected.arn,
        port=443)
    ```


    :param str arn: The arn of the listener. Required if `load_balancer_arn` and `port` is not set.
    :param str load_balancer_arn: The arn of the load balancer. Required if `arn` is not set.
    :param float port: The port of the listener. Required if `arn` is not set.
    """
    pulumi.log.warn("get_listener is deprecated: aws.applicationloadbalancing.getListener has been deprecated in favor of aws.alb.getListener")
    __args__ = dict()

    __args__['arn'] = arn
    __args__['loadBalancerArn'] = load_balancer_arn
    __args__['port'] = port
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:applicationloadbalancing/getListener:getListener', __args__, opts=opts).value

    return AwaitableGetListenerResult(
        arn=__ret__.get('arn'),
        certificate_arn=__ret__.get('certificateArn'),
        default_actions=__ret__.get('defaultActions'),
        id=__ret__.get('id'),
        load_balancer_arn=__ret__.get('loadBalancerArn'),
        port=__ret__.get('port'),
        protocol=__ret__.get('protocol'),
        ssl_policy=__ret__.get('sslPolicy'))

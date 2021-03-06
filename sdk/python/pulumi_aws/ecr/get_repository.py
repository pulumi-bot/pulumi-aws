# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetRepositoryResult(object):
    """
    A collection of values returned by getRepository.
    """
    def __init__(__self__, arn=None, registry_id=None, repository_url=None, id=None):
        if arn and not isinstance(arn, basestring):
            raise TypeError('Expected argument arn to be a basestring')
        __self__.arn = arn
        """
        Full ARN of the repository.
        """
        if registry_id and not isinstance(registry_id, basestring):
            raise TypeError('Expected argument registry_id to be a basestring')
        __self__.registry_id = registry_id
        """
        The registry ID where the repository was created.
        """
        if repository_url and not isinstance(repository_url, basestring):
            raise TypeError('Expected argument repository_url to be a basestring')
        __self__.repository_url = repository_url
        """
        The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_repository(name=None):
    """
    The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = pulumi.runtime.invoke('aws:ecr/getRepository:getRepository', __args__)

    return GetRepositoryResult(
        arn=__ret__.get('arn'),
        registry_id=__ret__.get('registryId'),
        repository_url=__ret__.get('repositoryUrl'),
        id=__ret__.get('id'))

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 external_connections: Optional[pulumi.Input['RepositoryExternalConnectionsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstreams: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[str] domain: The domain that contains the created repository.
        :param pulumi.Input[str] repository: The name of the repository to create.
        :param pulumi.Input[str] description: The description of the repository.
        :param pulumi.Input[str] domain_owner: The account number of the AWS account that owns the domain.
        :param pulumi.Input['RepositoryExternalConnectionsArgs'] external_connections: An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]] upstreams: A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "repository", repository)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_owner is not None:
            pulumi.set(__self__, "domain_owner", domain_owner)
        if external_connections is not None:
            pulumi.set(__self__, "external_connections", external_connections)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if upstreams is not None:
            pulumi.set(__self__, "upstreams", upstreams)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain that contains the created repository.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The name of the repository to create.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the repository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainOwner")
    def domain_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The account number of the AWS account that owns the domain.
        """
        return pulumi.get(self, "domain_owner")

    @domain_owner.setter
    def domain_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_owner", value)

    @property
    @pulumi.getter(name="externalConnections")
    def external_connections(self) -> Optional[pulumi.Input['RepositoryExternalConnectionsArgs']]:
        """
        An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        """
        return pulumi.get(self, "external_connections")

    @external_connections.setter
    def external_connections(self, value: Optional[pulumi.Input['RepositoryExternalConnectionsArgs']]):
        pulumi.set(self, "external_connections", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter
    def upstreams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]]]:
        """
        A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        return pulumi.get(self, "upstreams")

    @upstreams.setter
    def upstreams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]]]):
        pulumi.set(self, "upstreams", value)


@pulumi.input_type
class _RepositoryState:
    def __init__(__self__, *,
                 administrator_account: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 external_connections: Optional[pulumi.Input['RepositoryExternalConnectionsArgs']] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstreams: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]]] = None):
        """
        Input properties used for looking up and filtering Repository resources.
        :param pulumi.Input[str] administrator_account: The account number of the AWS account that manages the repository.
        :param pulumi.Input[str] arn: The ARN of the repository.
        :param pulumi.Input[str] description: The description of the repository.
        :param pulumi.Input[str] domain: The domain that contains the created repository.
        :param pulumi.Input[str] domain_owner: The account number of the AWS account that owns the domain.
        :param pulumi.Input['RepositoryExternalConnectionsArgs'] external_connections: An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        :param pulumi.Input[str] repository: The name of the repository to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]] upstreams: A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        if administrator_account is not None:
            pulumi.set(__self__, "administrator_account", administrator_account)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if domain_owner is not None:
            pulumi.set(__self__, "domain_owner", domain_owner)
        if external_connections is not None:
            pulumi.set(__self__, "external_connections", external_connections)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if upstreams is not None:
            pulumi.set(__self__, "upstreams", upstreams)

    @property
    @pulumi.getter(name="administratorAccount")
    def administrator_account(self) -> Optional[pulumi.Input[str]]:
        """
        The account number of the AWS account that manages the repository.
        """
        return pulumi.get(self, "administrator_account")

    @administrator_account.setter
    def administrator_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrator_account", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the repository.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the repository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain that contains the created repository.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="domainOwner")
    def domain_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The account number of the AWS account that owns the domain.
        """
        return pulumi.get(self, "domain_owner")

    @domain_owner.setter
    def domain_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_owner", value)

    @property
    @pulumi.getter(name="externalConnections")
    def external_connections(self) -> Optional[pulumi.Input['RepositoryExternalConnectionsArgs']]:
        """
        An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        """
        return pulumi.get(self, "external_connections")

    @external_connections.setter
    def external_connections(self, value: Optional[pulumi.Input['RepositoryExternalConnectionsArgs']]):
        pulumi.set(self, "external_connections", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the repository to create.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter
    def upstreams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]]]:
        """
        A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        return pulumi.get(self, "upstreams")

    @upstreams.setter
    def upstreams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryUpstreamArgs']]]]):
        pulumi.set(self, "upstreams", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 external_connections: Optional[pulumi.Input[pulumi.InputType['RepositoryExternalConnectionsArgs']]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstreams: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryUpstreamArgs']]]]] = None,
                 __props__=None):
        """
        Provides a CodeArtifact Repository Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey", description="domain key")
        example_domain = aws.codeartifact.Domain("exampleDomain",
            domain="example",
            encryption_key=example_key.arn)
        test = aws.codeartifact.Repository("test",
            repository="example",
            domain=example_domain.domain)
        ```
        ### With Upstream Repository

        ```python
        import pulumi
        import pulumi_aws as aws

        upstream = aws.codeartifact.Repository("upstream",
            repository="upstream",
            domain=aws_codeartifact_domain["test"]["domain"])
        test = aws.codeartifact.Repository("test",
            repository="example",
            domain=aws_codeartifact_domain["example"]["domain"],
            upstreams=[aws.codeartifact.RepositoryUpstreamArgs(
                repository_name=upstream.repository,
            )])
        ```
        ### With External Connection

        ```python
        import pulumi
        import pulumi_aws as aws

        upstream = aws.codeartifact.Repository("upstream",
            repository="upstream",
            domain=aws_codeartifact_domain["test"]["domain"])
        test = aws.codeartifact.Repository("test",
            repository="example",
            domain=aws_codeartifact_domain["example"]["domain"],
            external_connections=aws.codeartifact.RepositoryExternalConnectionsArgs(
                external_connection_name="public:npmjs",
            ))
        ```

        ## Import

        CodeArtifact Repository can be imported using the CodeArtifact Repository ARN, e.g.

        ```sh
         $ pulumi import aws:codeartifact/repository:Repository example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the repository.
        :param pulumi.Input[str] domain: The domain that contains the created repository.
        :param pulumi.Input[str] domain_owner: The account number of the AWS account that owns the domain.
        :param pulumi.Input[pulumi.InputType['RepositoryExternalConnectionsArgs']] external_connections: An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        :param pulumi.Input[str] repository: The name of the repository to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryUpstreamArgs']]]] upstreams: A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: RepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CodeArtifact Repository Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey", description="domain key")
        example_domain = aws.codeartifact.Domain("exampleDomain",
            domain="example",
            encryption_key=example_key.arn)
        test = aws.codeartifact.Repository("test",
            repository="example",
            domain=example_domain.domain)
        ```
        ### With Upstream Repository

        ```python
        import pulumi
        import pulumi_aws as aws

        upstream = aws.codeartifact.Repository("upstream",
            repository="upstream",
            domain=aws_codeartifact_domain["test"]["domain"])
        test = aws.codeartifact.Repository("test",
            repository="example",
            domain=aws_codeartifact_domain["example"]["domain"],
            upstreams=[aws.codeartifact.RepositoryUpstreamArgs(
                repository_name=upstream.repository,
            )])
        ```
        ### With External Connection

        ```python
        import pulumi
        import pulumi_aws as aws

        upstream = aws.codeartifact.Repository("upstream",
            repository="upstream",
            domain=aws_codeartifact_domain["test"]["domain"])
        test = aws.codeartifact.Repository("test",
            repository="example",
            domain=aws_codeartifact_domain["example"]["domain"],
            external_connections=aws.codeartifact.RepositoryExternalConnectionsArgs(
                external_connection_name="public:npmjs",
            ))
        ```

        ## Import

        CodeArtifact Repository can be imported using the CodeArtifact Repository ARN, e.g.

        ```sh
         $ pulumi import aws:codeartifact/repository:Repository example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
        ```

        :param str resource_name_: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 external_connections: Optional[pulumi.Input[pulumi.InputType['RepositoryExternalConnectionsArgs']]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstreams: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryUpstreamArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            __props__.__dict__["description"] = description
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["domain_owner"] = domain_owner
            __props__.__dict__["external_connections"] = external_connections
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["upstreams"] = upstreams
            __props__.__dict__["administrator_account"] = None
            __props__.__dict__["arn"] = None
        super(Repository, __self__).__init__(
            'aws:codeartifact/repository:Repository',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            administrator_account: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            domain_owner: Optional[pulumi.Input[str]] = None,
            external_connections: Optional[pulumi.Input[pulumi.InputType['RepositoryExternalConnectionsArgs']]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            upstreams: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryUpstreamArgs']]]]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_account: The account number of the AWS account that manages the repository.
        :param pulumi.Input[str] arn: The ARN of the repository.
        :param pulumi.Input[str] description: The description of the repository.
        :param pulumi.Input[str] domain: The domain that contains the created repository.
        :param pulumi.Input[str] domain_owner: The account number of the AWS account that owns the domain.
        :param pulumi.Input[pulumi.InputType['RepositoryExternalConnectionsArgs']] external_connections: An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        :param pulumi.Input[str] repository: The name of the repository to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryUpstreamArgs']]]] upstreams: A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryState.__new__(_RepositoryState)

        __props__.__dict__["administrator_account"] = administrator_account
        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["domain"] = domain
        __props__.__dict__["domain_owner"] = domain_owner
        __props__.__dict__["external_connections"] = external_connections
        __props__.__dict__["repository"] = repository
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["upstreams"] = upstreams
        return Repository(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administratorAccount")
    def administrator_account(self) -> pulumi.Output[str]:
        """
        The account number of the AWS account that manages the repository.
        """
        return pulumi.get(self, "administrator_account")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the repository.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the repository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain that contains the created repository.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainOwner")
    def domain_owner(self) -> pulumi.Output[str]:
        """
        The account number of the AWS account that owns the domain.
        """
        return pulumi.get(self, "domain_owner")

    @property
    @pulumi.getter(name="externalConnections")
    def external_connections(self) -> pulumi.Output[Optional['outputs.RepositoryExternalConnections']]:
        """
        An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
        """
        return pulumi.get(self, "external_connections")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The name of the repository to create.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter
    def upstreams(self) -> pulumi.Output[Optional[Sequence['outputs.RepositoryUpstream']]]:
        """
        A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
        """
        return pulumi.get(self, "upstreams")


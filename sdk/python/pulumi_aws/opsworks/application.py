# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 stack_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 app_sources: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationAppSourceArgs']]]] = None,
                 auto_bundle_on_deploy: Optional[pulumi.Input[str]] = None,
                 aws_flow_ruby_settings: Optional[pulumi.Input[str]] = None,
                 data_source_arn: Optional[pulumi.Input[str]] = None,
                 data_source_database_name: Optional[pulumi.Input[str]] = None,
                 data_source_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document_root: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationEnvironmentArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rails_env: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 ssl_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationSslConfigurationArgs']]]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] stack_id: The id of the stack the application will belong to.
        :param pulumi.Input[str] type: Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationAppSourceArgs']]] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationEnvironmentArgs']]] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationSslConfigurationArgs']]] ssl_configurations: The SSL configuration of the app. Object is described below.
        """
        pulumi.set(__self__, "stack_id", stack_id)
        pulumi.set(__self__, "type", type)
        if app_sources is not None:
            pulumi.set(__self__, "app_sources", app_sources)
        if auto_bundle_on_deploy is not None:
            pulumi.set(__self__, "auto_bundle_on_deploy", auto_bundle_on_deploy)
        if aws_flow_ruby_settings is not None:
            pulumi.set(__self__, "aws_flow_ruby_settings", aws_flow_ruby_settings)
        if data_source_arn is not None:
            pulumi.set(__self__, "data_source_arn", data_source_arn)
        if data_source_database_name is not None:
            pulumi.set(__self__, "data_source_database_name", data_source_database_name)
        if data_source_type is not None:
            pulumi.set(__self__, "data_source_type", data_source_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if document_root is not None:
            pulumi.set(__self__, "document_root", document_root)
        if domains is not None:
            pulumi.set(__self__, "domains", domains)
        if enable_ssl is not None:
            pulumi.set(__self__, "enable_ssl", enable_ssl)
        if environments is not None:
            pulumi.set(__self__, "environments", environments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rails_env is not None:
            pulumi.set(__self__, "rails_env", rails_env)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)
        if ssl_configurations is not None:
            pulumi.set(__self__, "ssl_configurations", ssl_configurations)

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Input[str]:
        """
        The id of the stack the application will belong to.
        """
        return pulumi.get(self, "stack_id")

    @stack_id.setter
    def stack_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="appSources")
    def app_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationAppSourceArgs']]]]:
        """
        SCM configuration of the app as described below.
        """
        return pulumi.get(self, "app_sources")

    @app_sources.setter
    def app_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationAppSourceArgs']]]]):
        pulumi.set(self, "app_sources", value)

    @property
    @pulumi.getter(name="autoBundleOnDeploy")
    def auto_bundle_on_deploy(self) -> Optional[pulumi.Input[str]]:
        """
        Run bundle install when deploying for application of type `rails`.
        """
        return pulumi.get(self, "auto_bundle_on_deploy")

    @auto_bundle_on_deploy.setter
    def auto_bundle_on_deploy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_bundle_on_deploy", value)

    @property
    @pulumi.getter(name="awsFlowRubySettings")
    def aws_flow_ruby_settings(self) -> Optional[pulumi.Input[str]]:
        """
        Specify activity and workflow workers for your app using the aws-flow gem.
        """
        return pulumi.get(self, "aws_flow_ruby_settings")

    @aws_flow_ruby_settings.setter
    def aws_flow_ruby_settings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_flow_ruby_settings", value)

    @property
    @pulumi.getter(name="dataSourceArn")
    def data_source_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The data source's ARN.
        """
        return pulumi.get(self, "data_source_arn")

    @data_source_arn.setter
    def data_source_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source_arn", value)

    @property
    @pulumi.getter(name="dataSourceDatabaseName")
    def data_source_database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The database name.
        """
        return pulumi.get(self, "data_source_database_name")

    @data_source_database_name.setter
    def data_source_database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source_database_name", value)

    @property
    @pulumi.getter(name="dataSourceType")
    def data_source_type(self) -> Optional[pulumi.Input[str]]:
        """
        The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        """
        return pulumi.get(self, "data_source_type")

    @data_source_type.setter
    def data_source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the app.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="documentRoot")
    def document_root(self) -> Optional[pulumi.Input[str]]:
        """
        Subfolder for the document root for application of type `rails`.
        """
        return pulumi.get(self, "document_root")

    @document_root.setter
    def document_root(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document_root", value)

    @property
    @pulumi.getter
    def domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of virtual host alias.
        """
        return pulumi.get(self, "domains")

    @domains.setter
    def domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domains", value)

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        """
        return pulumi.get(self, "enable_ssl")

    @enable_ssl.setter
    def enable_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ssl", value)

    @property
    @pulumi.getter
    def environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationEnvironmentArgs']]]]:
        """
        Object to define environment variables.  Object is described below.
        """
        return pulumi.get(self, "environments")

    @environments.setter
    def environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationEnvironmentArgs']]]]):
        pulumi.set(self, "environments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable name for the application.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="railsEnv")
    def rails_env(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Rails environment for application of type `rails`.
        """
        return pulumi.get(self, "rails_env")

    @rails_env.setter
    def rails_env(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rails_env", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)

    @property
    @pulumi.getter(name="sslConfigurations")
    def ssl_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationSslConfigurationArgs']]]]:
        """
        The SSL configuration of the app. Object is described below.
        """
        return pulumi.get(self, "ssl_configurations")

    @ssl_configurations.setter
    def ssl_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationSslConfigurationArgs']]]]):
        pulumi.set(self, "ssl_configurations", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an OpsWorks application resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_app = aws.opsworks.Application("foo-app",
            short_name="foobar",
            stack_id=aws_opsworks_stack["main"]["id"],
            type="rails",
            description="This is a Rails application",
            domains=[
                "example.com",
                "sub.example.com",
            ],
            environments=[aws.opsworks.ApplicationEnvironmentArgs(
                key="key",
                value="value",
                secure=False,
            )],
            app_sources=[aws.opsworks.ApplicationAppSourceArgs(
                type="git",
                revision="master",
                url="https://github.com/example.git",
            )],
            enable_ssl=True,
            ssl_configurations=[aws.opsworks.ApplicationSslConfigurationArgs(
                private_key=(lambda path: open(path).read())("./foobar.key"),
                certificate=(lambda path: open(path).read())("./foobar.crt"),
            )],
            document_root="public",
            auto_bundle_on_deploy="true",
            rails_env="staging")
        ```

        ## Import

        Opsworks Application can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:opsworks/application:Application test <id>
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]]] = None,
                 auto_bundle_on_deploy: Optional[pulumi.Input[str]] = None,
                 aws_flow_ruby_settings: Optional[pulumi.Input[str]] = None,
                 data_source_arn: Optional[pulumi.Input[str]] = None,
                 data_source_database_name: Optional[pulumi.Input[str]] = None,
                 data_source_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document_root: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rails_env: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 ssl_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an OpsWorks application resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_app = aws.opsworks.Application("foo-app",
            short_name="foobar",
            stack_id=aws_opsworks_stack["main"]["id"],
            type="rails",
            description="This is a Rails application",
            domains=[
                "example.com",
                "sub.example.com",
            ],
            environments=[aws.opsworks.ApplicationEnvironmentArgs(
                key="key",
                value="value",
                secure=False,
            )],
            app_sources=[aws.opsworks.ApplicationAppSourceArgs(
                type="git",
                revision="master",
                url="https://github.com/example.git",
            )],
            enable_ssl=True,
            ssl_configurations=[aws.opsworks.ApplicationSslConfigurationArgs(
                private_key=(lambda path: open(path).read())("./foobar.key"),
                certificate=(lambda path: open(path).read())("./foobar.crt"),
            )],
            document_root="public",
            auto_bundle_on_deploy="true",
            rails_env="staging")
        ```

        ## Import

        Opsworks Application can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:opsworks/application:Application test <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]] ssl_configurations: The SSL configuration of the app. Object is described below.
        :param pulumi.Input[str] stack_id: The id of the stack the application will belong to.
        :param pulumi.Input[str] type: Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]]] = None,
                 auto_bundle_on_deploy: Optional[pulumi.Input[str]] = None,
                 aws_flow_ruby_settings: Optional[pulumi.Input[str]] = None,
                 data_source_arn: Optional[pulumi.Input[str]] = None,
                 data_source_database_name: Optional[pulumi.Input[str]] = None,
                 data_source_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document_root: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rails_env: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 ssl_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['app_sources'] = app_sources
            __props__['auto_bundle_on_deploy'] = auto_bundle_on_deploy
            __props__['aws_flow_ruby_settings'] = aws_flow_ruby_settings
            __props__['data_source_arn'] = data_source_arn
            __props__['data_source_database_name'] = data_source_database_name
            __props__['data_source_type'] = data_source_type
            __props__['description'] = description
            __props__['document_root'] = document_root
            __props__['domains'] = domains
            __props__['enable_ssl'] = enable_ssl
            __props__['environments'] = environments
            __props__['name'] = name
            __props__['rails_env'] = rails_env
            __props__['short_name'] = short_name
            __props__['ssl_configurations'] = ssl_configurations
            if stack_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_id'")
            __props__['stack_id'] = stack_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(Application, __self__).__init__(
            'aws:opsworks/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]]] = None,
            auto_bundle_on_deploy: Optional[pulumi.Input[str]] = None,
            aws_flow_ruby_settings: Optional[pulumi.Input[str]] = None,
            data_source_arn: Optional[pulumi.Input[str]] = None,
            data_source_database_name: Optional[pulumi.Input[str]] = None,
            data_source_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            document_root: Optional[pulumi.Input[str]] = None,
            domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enable_ssl: Optional[pulumi.Input[bool]] = None,
            environments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rails_env: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None,
            ssl_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]] ssl_configurations: The SSL configuration of the app. Object is described below.
        :param pulumi.Input[str] stack_id: The id of the stack the application will belong to.
        :param pulumi.Input[str] type: Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_sources"] = app_sources
        __props__["auto_bundle_on_deploy"] = auto_bundle_on_deploy
        __props__["aws_flow_ruby_settings"] = aws_flow_ruby_settings
        __props__["data_source_arn"] = data_source_arn
        __props__["data_source_database_name"] = data_source_database_name
        __props__["data_source_type"] = data_source_type
        __props__["description"] = description
        __props__["document_root"] = document_root
        __props__["domains"] = domains
        __props__["enable_ssl"] = enable_ssl
        __props__["environments"] = environments
        __props__["name"] = name
        __props__["rails_env"] = rails_env
        __props__["short_name"] = short_name
        __props__["ssl_configurations"] = ssl_configurations
        __props__["stack_id"] = stack_id
        __props__["type"] = type
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appSources")
    def app_sources(self) -> pulumi.Output[Sequence['outputs.ApplicationAppSource']]:
        """
        SCM configuration of the app as described below.
        """
        return pulumi.get(self, "app_sources")

    @property
    @pulumi.getter(name="autoBundleOnDeploy")
    def auto_bundle_on_deploy(self) -> pulumi.Output[Optional[str]]:
        """
        Run bundle install when deploying for application of type `rails`.
        """
        return pulumi.get(self, "auto_bundle_on_deploy")

    @property
    @pulumi.getter(name="awsFlowRubySettings")
    def aws_flow_ruby_settings(self) -> pulumi.Output[Optional[str]]:
        """
        Specify activity and workflow workers for your app using the aws-flow gem.
        """
        return pulumi.get(self, "aws_flow_ruby_settings")

    @property
    @pulumi.getter(name="dataSourceArn")
    def data_source_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The data source's ARN.
        """
        return pulumi.get(self, "data_source_arn")

    @property
    @pulumi.getter(name="dataSourceDatabaseName")
    def data_source_database_name(self) -> pulumi.Output[Optional[str]]:
        """
        The database name.
        """
        return pulumi.get(self, "data_source_database_name")

    @property
    @pulumi.getter(name="dataSourceType")
    def data_source_type(self) -> pulumi.Output[Optional[str]]:
        """
        The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        """
        return pulumi.get(self, "data_source_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the app.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="documentRoot")
    def document_root(self) -> pulumi.Output[Optional[str]]:
        """
        Subfolder for the document root for application of type `rails`.
        """
        return pulumi.get(self, "document_root")

    @property
    @pulumi.getter
    def domains(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of virtual host alias.
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        """
        return pulumi.get(self, "enable_ssl")

    @property
    @pulumi.getter
    def environments(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationEnvironment']]]:
        """
        Object to define environment variables.  Object is described below.
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A human-readable name for the application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="railsEnv")
    def rails_env(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Rails environment for application of type `rails`.
        """
        return pulumi.get(self, "rails_env")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        """
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter(name="sslConfigurations")
    def ssl_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationSslConfiguration']]]:
        """
        The SSL configuration of the app. Object is described below.
        """
        return pulumi.get(self, "ssl_configurations")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        The id of the stack the application will belong to.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


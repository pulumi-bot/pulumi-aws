# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Application']


class Application(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_sources: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]]] = None,
                 auto_bundle_on_deploy: Optional[pulumi.Input[str]] = None,
                 aws_flow_ruby_settings: Optional[pulumi.Input[str]] = None,
                 data_source_arn: Optional[pulumi.Input[str]] = None,
                 data_source_database_name: Optional[pulumi.Input[str]] = None,
                 data_source_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document_root: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 enable_ssl: Optional[pulumi.Input[bool]] = None,
                 environments: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rails_env: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 ssl_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]]] = None,
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
            environments=[{
                "key": "key",
                "value": "value",
                "secure": False,
            }],
            app_sources=[{
                "type": "git",
                "revision": "master",
                "url": "https://github.com/example.git",
            }],
            enable_ssl=True,
            ssl_configurations=[{
                "private_key": (lambda path: open(path).read())("./foobar.key"),
                "certificate": (lambda path: open(path).read())("./foobar.crt"),
            }],
            document_root="public",
            auto_bundle_on_deploy="true",
            rails_env="staging")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[List[pulumi.Input[str]]] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]] ssl_configurations: The SSL configuration of the app. Object is described below.
        :param pulumi.Input[str] stack_id: The id of the stack the application will belong to.
        :param pulumi.Input[str] type: Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        """
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
            if stack_id is None:
                raise TypeError("Missing required property 'stack_id'")
            __props__['stack_id'] = stack_id
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(Application, __self__).__init__(
            'aws:opsworks/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            app_sources: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]]] = None,
            auto_bundle_on_deploy: Optional[pulumi.Input[str]] = None,
            aws_flow_ruby_settings: Optional[pulumi.Input[str]] = None,
            data_source_arn: Optional[pulumi.Input[str]] = None,
            data_source_database_name: Optional[pulumi.Input[str]] = None,
            data_source_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            document_root: Optional[pulumi.Input[str]] = None,
            domains: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            enable_ssl: Optional[pulumi.Input[bool]] = None,
            environments: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rails_env: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None,
            ssl_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationAppSourceArgs']]]] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[List[pulumi.Input[str]]] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationEnvironmentArgs']]]] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationSslConfigurationArgs']]]] ssl_configurations: The SSL configuration of the app. Object is described below.
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
    def app_sources(self) -> List['outputs.ApplicationAppSource']:
        """
        SCM configuration of the app as described below.
        """
        return pulumi.get(self, "app_sources")

    @property
    @pulumi.getter(name="autoBundleOnDeploy")
    def auto_bundle_on_deploy(self) -> Optional[str]:
        """
        Run bundle install when deploying for application of type `rails`.
        """
        return pulumi.get(self, "auto_bundle_on_deploy")

    @property
    @pulumi.getter(name="awsFlowRubySettings")
    def aws_flow_ruby_settings(self) -> Optional[str]:
        """
        Specify activity and workflow workers for your app using the aws-flow gem.
        """
        return pulumi.get(self, "aws_flow_ruby_settings")

    @property
    @pulumi.getter(name="dataSourceArn")
    def data_source_arn(self) -> Optional[str]:
        """
        The data source's ARN.
        """
        return pulumi.get(self, "data_source_arn")

    @property
    @pulumi.getter(name="dataSourceDatabaseName")
    def data_source_database_name(self) -> Optional[str]:
        """
        The database name.
        """
        return pulumi.get(self, "data_source_database_name")

    @property
    @pulumi.getter(name="dataSourceType")
    def data_source_type(self) -> Optional[str]:
        """
        The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        """
        return pulumi.get(self, "data_source_type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of the app.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="documentRoot")
    def document_root(self) -> Optional[str]:
        """
        Subfolder for the document root for application of type `rails`.
        """
        return pulumi.get(self, "document_root")

    @property
    @pulumi.getter
    def domains(self) -> Optional[List[str]]:
        """
        A list of virtual host alias.
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> Optional[bool]:
        """
        Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        """
        return pulumi.get(self, "enable_ssl")

    @property
    @pulumi.getter
    def environments(self) -> Optional[List['outputs.ApplicationEnvironment']]:
        """
        Object to define environment variables.  Object is described below.
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-readable name for the application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="railsEnv")
    def rails_env(self) -> Optional[str]:
        """
        The name of the Rails environment for application of type `rails`.
        """
        return pulumi.get(self, "rails_env")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> str:
        """
        A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        """
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter(name="sslConfigurations")
    def ssl_configurations(self) -> Optional[List['outputs.ApplicationSslConfiguration']]:
        """
        The SSL configuration of the app. Object is described below.
        """
        return pulumi.get(self, "ssl_configurations")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        The id of the stack the application will belong to.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


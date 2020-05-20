# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Application(pulumi.CustomResource):
    app_sources: pulumi.Output[list]
    """
    SCM configuration of the app as described below.

      * `password` (`str`) - Password to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
      * `revision` (`str`) - For sources that are version-aware, the revision to use.
      * `sshKey` (`str`) - SSH key to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
      * `type` (`str`) - The type of source to use. For example, "archive".
      * `url` (`str`) - The URL where the app resource can be found.
      * `username` (`str`) - Username to use when authenticating to the source.
    """
    auto_bundle_on_deploy: pulumi.Output[str]
    """
    Run bundle install when deploying for application of type `rails`.
    """
    aws_flow_ruby_settings: pulumi.Output[str]
    """
    Specify activity and workflow workers for your app using the aws-flow gem.
    """
    data_source_arn: pulumi.Output[str]
    """
    The data source's ARN.
    """
    data_source_database_name: pulumi.Output[str]
    """
    The database name.
    """
    data_source_type: pulumi.Output[str]
    """
    The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
    """
    description: pulumi.Output[str]
    """
    A description of the app.
    """
    document_root: pulumi.Output[str]
    """
    Subfolder for the document root for application of type `rails`.
    """
    domains: pulumi.Output[list]
    """
    A list of virtual host alias.
    """
    enable_ssl: pulumi.Output[bool]
    """
    Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
    """
    environments: pulumi.Output[list]
    """
    Object to define environment variables.  Object is described below.

      * `key` (`str`) - Variable name.
      * `secure` (`bool`) - Set visibility of the variable value to `true` or `false`.
      * `value` (`str`) - Variable value.
    """
    name: pulumi.Output[str]
    """
    A human-readable name for the application.
    """
    rails_env: pulumi.Output[str]
    """
    The name of the Rails environment for application of type `rails`.
    """
    short_name: pulumi.Output[str]
    """
    A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
    """
    ssl_configurations: pulumi.Output[list]
    """
    The SSL configuration of the app. Object is described below.

      * `certificate` (`str`) - The contents of the certificate's domain.crt file.
      * `chain` (`str`) - Can be used to specify an intermediate certificate authority key or client authentication.
      * `private_key` (`str`) - The private key; the contents of the certificate's domain.key file.
    """
    stack_id: pulumi.Output[str]
    """
    The id of the stack the application will belong to.
    """
    type: pulumi.Output[str]
    """
    Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
    """
    def __init__(__self__, resource_name, opts=None, app_sources=None, auto_bundle_on_deploy=None, aws_flow_ruby_settings=None, data_source_arn=None, data_source_database_name=None, data_source_type=None, description=None, document_root=None, domains=None, enable_ssl=None, environments=None, name=None, rails_env=None, short_name=None, ssl_configurations=None, stack_id=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an OpsWorks application resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        foo_app = aws.opsworks.Application("foo-app",
            app_sources=[{
                "revision": "master",
                "type": "git",
                "url": "https://github.com/example.git",
            }],
            auto_bundle_on_deploy=True,
            description="This is a Rails application",
            document_root="public",
            domains=[
                "example.com",
                "sub.example.com",
            ],
            enable_ssl=True,
            environments=[{
                "key": "key",
                "secure": False,
                "value": "value",
            }],
            rails_env="staging",
            short_name="foobar",
            ssl_configurations=[{
                "certificate": (lambda path: open(path).read())("./foobar.crt"),
                "private_key": (lambda path: open(path).read())("./foobar.key"),
            }],
            stack_id=aws_opsworks_stack["main"]["id"],
            type="rails")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[list] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[list] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[list] ssl_configurations: The SSL configuration of the app. Object is described below.
        :param pulumi.Input[str] stack_id: The id of the stack the application will belong to.
        :param pulumi.Input[str] type: Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.

        The **app_sources** object supports the following:

          * `password` (`pulumi.Input[str]`) - Password to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
          * `revision` (`pulumi.Input[str]`) - For sources that are version-aware, the revision to use.
          * `sshKey` (`pulumi.Input[str]`) - SSH key to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
          * `type` (`pulumi.Input[str]`) - The type of source to use. For example, "archive".
          * `url` (`pulumi.Input[str]`) - The URL where the app resource can be found.
          * `username` (`pulumi.Input[str]`) - Username to use when authenticating to the source.

        The **environments** object supports the following:

          * `key` (`pulumi.Input[str]`) - Variable name.
          * `secure` (`pulumi.Input[bool]`) - Set visibility of the variable value to `true` or `false`.
          * `value` (`pulumi.Input[str]`) - Variable value.

        The **ssl_configurations** object supports the following:

          * `certificate` (`pulumi.Input[str]`) - The contents of the certificate's domain.crt file.
          * `chain` (`pulumi.Input[str]`) - Can be used to specify an intermediate certificate authority key or client authentication.
          * `private_key` (`pulumi.Input[str]`) - The private key; the contents of the certificate's domain.key file.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, app_sources=None, auto_bundle_on_deploy=None, aws_flow_ruby_settings=None, data_source_arn=None, data_source_database_name=None, data_source_type=None, description=None, document_root=None, domains=None, enable_ssl=None, environments=None, name=None, rails_env=None, short_name=None, ssl_configurations=None, stack_id=None, type=None):
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_sources: SCM configuration of the app as described below.
        :param pulumi.Input[str] auto_bundle_on_deploy: Run bundle install when deploying for application of type `rails`.
        :param pulumi.Input[str] aws_flow_ruby_settings: Specify activity and workflow workers for your app using the aws-flow gem.
        :param pulumi.Input[str] data_source_arn: The data source's ARN.
        :param pulumi.Input[str] data_source_database_name: The database name.
        :param pulumi.Input[str] data_source_type: The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
        :param pulumi.Input[str] description: A description of the app.
        :param pulumi.Input[str] document_root: Subfolder for the document root for application of type `rails`.
        :param pulumi.Input[list] domains: A list of virtual host alias.
        :param pulumi.Input[bool] enable_ssl: Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
        :param pulumi.Input[list] environments: Object to define environment variables.  Object is described below.
        :param pulumi.Input[str] name: A human-readable name for the application.
        :param pulumi.Input[str] rails_env: The name of the Rails environment for application of type `rails`.
        :param pulumi.Input[str] short_name: A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
        :param pulumi.Input[list] ssl_configurations: The SSL configuration of the app. Object is described below.
        :param pulumi.Input[str] stack_id: The id of the stack the application will belong to.
        :param pulumi.Input[str] type: Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.

        The **app_sources** object supports the following:

          * `password` (`pulumi.Input[str]`) - Password to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
          * `revision` (`pulumi.Input[str]`) - For sources that are version-aware, the revision to use.
          * `sshKey` (`pulumi.Input[str]`) - SSH key to use when authenticating to the source. This provider cannot perform drift detection of this configuration.
          * `type` (`pulumi.Input[str]`) - The type of source to use. For example, "archive".
          * `url` (`pulumi.Input[str]`) - The URL where the app resource can be found.
          * `username` (`pulumi.Input[str]`) - Username to use when authenticating to the source.

        The **environments** object supports the following:

          * `key` (`pulumi.Input[str]`) - Variable name.
          * `secure` (`pulumi.Input[bool]`) - Set visibility of the variable value to `true` or `false`.
          * `value` (`pulumi.Input[str]`) - Variable value.

        The **ssl_configurations** object supports the following:

          * `certificate` (`pulumi.Input[str]`) - The contents of the certificate's domain.crt file.
          * `chain` (`pulumi.Input[str]`) - Can be used to specify an intermediate certificate authority key or client authentication.
          * `private_key` (`pulumi.Input[str]`) - The private key; the contents of the certificate's domain.key file.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


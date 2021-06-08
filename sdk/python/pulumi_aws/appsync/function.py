# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FunctionArgs', 'Function']

@pulumi.input_type
class FunctionArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 data_source: pulumi.Input[str],
                 request_mapping_template: pulumi.Input[str],
                 response_mapping_template: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Function resource.
        :param pulumi.Input[str] api_id: The ID of the associated AppSync API.
        :param pulumi.Input[str] data_source: The Function DataSource name.
        :param pulumi.Input[str] request_mapping_template: The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        :param pulumi.Input[str] response_mapping_template: The Function response mapping template.
        :param pulumi.Input[str] description: The Function description.
        :param pulumi.Input[str] function_version: The version of the request mapping template. Currently the supported value is `2018-05-29`.
        :param pulumi.Input[str] name: The Function name. The function name does not have to be unique.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "data_source", data_source)
        pulumi.set(__self__, "request_mapping_template", request_mapping_template)
        pulumi.set(__self__, "response_mapping_template", response_mapping_template)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if function_version is not None:
            pulumi.set(__self__, "function_version", function_version)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        The ID of the associated AppSync API.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> pulumi.Input[str]:
        """
        The Function DataSource name.
        """
        return pulumi.get(self, "data_source")

    @data_source.setter
    def data_source(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_source", value)

    @property
    @pulumi.getter(name="requestMappingTemplate")
    def request_mapping_template(self) -> pulumi.Input[str]:
        """
        The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        """
        return pulumi.get(self, "request_mapping_template")

    @request_mapping_template.setter
    def request_mapping_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "request_mapping_template", value)

    @property
    @pulumi.getter(name="responseMappingTemplate")
    def response_mapping_template(self) -> pulumi.Input[str]:
        """
        The Function response mapping template.
        """
        return pulumi.get(self, "response_mapping_template")

    @response_mapping_template.setter
    def response_mapping_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "response_mapping_template", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Function description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the request mapping template. Currently the supported value is `2018-05-29`.
        """
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Function name. The function name does not have to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FunctionState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_mapping_template: Optional[pulumi.Input[str]] = None,
                 response_mapping_template: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Function resources.
        :param pulumi.Input[str] api_id: The ID of the associated AppSync API.
        :param pulumi.Input[str] arn: The ARN of the Function object.
        :param pulumi.Input[str] data_source: The Function DataSource name.
        :param pulumi.Input[str] description: The Function description.
        :param pulumi.Input[str] function_id: A unique ID representing the Function object.
        :param pulumi.Input[str] function_version: The version of the request mapping template. Currently the supported value is `2018-05-29`.
        :param pulumi.Input[str] name: The Function name. The function name does not have to be unique.
        :param pulumi.Input[str] request_mapping_template: The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        :param pulumi.Input[str] response_mapping_template: The Function response mapping template.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if data_source is not None:
            pulumi.set(__self__, "data_source", data_source)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if function_id is not None:
            pulumi.set(__self__, "function_id", function_id)
        if function_version is not None:
            pulumi.set(__self__, "function_version", function_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_mapping_template is not None:
            pulumi.set(__self__, "request_mapping_template", request_mapping_template)
        if response_mapping_template is not None:
            pulumi.set(__self__, "response_mapping_template", response_mapping_template)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the associated AppSync API.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Function object.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> Optional[pulumi.Input[str]]:
        """
        The Function DataSource name.
        """
        return pulumi.get(self, "data_source")

    @data_source.setter
    def data_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Function description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID representing the Function object.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the request mapping template. Currently the supported value is `2018-05-29`.
        """
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Function name. The function name does not have to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requestMappingTemplate")
    def request_mapping_template(self) -> Optional[pulumi.Input[str]]:
        """
        The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        """
        return pulumi.get(self, "request_mapping_template")

    @request_mapping_template.setter
    def request_mapping_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_mapping_template", value)

    @property
    @pulumi.getter(name="responseMappingTemplate")
    def response_mapping_template(self) -> Optional[pulumi.Input[str]]:
        """
        The Function response mapping template.
        """
        return pulumi.get(self, "response_mapping_template")

    @response_mapping_template.setter
    def response_mapping_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_mapping_template", value)


class Function(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_mapping_template: Optional[pulumi.Input[str]] = None,
                 response_mapping_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AppSync Function.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi",
            authentication_type="API_KEY",
            schema=\"\"\"type Mutation {
          putPost(id: ID!, title: String!): Post
        }

        type Post {
          id: ID!
          title: String!
        }

        type Query {
          singlePost(id: ID!): Post
        }

        schema {
          query: Query
          mutation: Mutation
        }
        \"\"\")
        example_data_source = aws.appsync.DataSource("exampleDataSource",
            api_id=example_graph_ql_api.id,
            name="example",
            type="HTTP",
            http_config=aws.appsync.DataSourceHttpConfigArgs(
                endpoint="http://example.com",
            ))
        example_function = aws.appsync.Function("exampleFunction",
            api_id=example_graph_ql_api.id,
            data_source=example_data_source.name,
            name="example",
            request_mapping_template=\"\"\"{
            "version": "2018-05-29",
            "method": "GET",
            "resourcePath": "/",
            "params":{
                "headers": $utils.http.copyheaders($ctx.request.headers)
            }
        }
        \"\"\",
            response_mapping_template=\"\"\"#if($ctx.result.statusCode == 200)
            $ctx.result.body
        #else
            $utils.appendError($ctx.result.body, $ctx.result.statusCode)
        #end
        \"\"\")
        ```

        ## Import

        `aws_appsync_function` can be imported using the AppSync API ID and Function ID separated by `-`, e.g.

        ```sh
         $ pulumi import aws:appsync/function:Function example xxxxx-yyyyy
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The ID of the associated AppSync API.
        :param pulumi.Input[str] data_source: The Function DataSource name.
        :param pulumi.Input[str] description: The Function description.
        :param pulumi.Input[str] function_version: The version of the request mapping template. Currently the supported value is `2018-05-29`.
        :param pulumi.Input[str] name: The Function name. The function name does not have to be unique.
        :param pulumi.Input[str] request_mapping_template: The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        :param pulumi.Input[str] response_mapping_template: The Function response mapping template.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: FunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AppSync Function.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi",
            authentication_type="API_KEY",
            schema=\"\"\"type Mutation {
          putPost(id: ID!, title: String!): Post
        }

        type Post {
          id: ID!
          title: String!
        }

        type Query {
          singlePost(id: ID!): Post
        }

        schema {
          query: Query
          mutation: Mutation
        }
        \"\"\")
        example_data_source = aws.appsync.DataSource("exampleDataSource",
            api_id=example_graph_ql_api.id,
            name="example",
            type="HTTP",
            http_config=aws.appsync.DataSourceHttpConfigArgs(
                endpoint="http://example.com",
            ))
        example_function = aws.appsync.Function("exampleFunction",
            api_id=example_graph_ql_api.id,
            data_source=example_data_source.name,
            name="example",
            request_mapping_template=\"\"\"{
            "version": "2018-05-29",
            "method": "GET",
            "resourcePath": "/",
            "params":{
                "headers": $utils.http.copyheaders($ctx.request.headers)
            }
        }
        \"\"\",
            response_mapping_template=\"\"\"#if($ctx.result.statusCode == 200)
            $ctx.result.body
        #else
            $utils.appendError($ctx.result.body, $ctx.result.statusCode)
        #end
        \"\"\")
        ```

        ## Import

        `aws_appsync_function` can be imported using the AppSync API ID and Function ID separated by `-`, e.g.

        ```sh
         $ pulumi import aws:appsync/function:Function example xxxxx-yyyyy
        ```

        :param str resource_name_: The name of the resource.
        :param FunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_mapping_template: Optional[pulumi.Input[str]] = None,
                 response_mapping_template: Optional[pulumi.Input[str]] = None,
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
            __props__ = FunctionArgs.__new__(FunctionArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            if data_source is None and not opts.urn:
                raise TypeError("Missing required property 'data_source'")
            __props__.__dict__["data_source"] = data_source
            __props__.__dict__["description"] = description
            __props__.__dict__["function_version"] = function_version
            __props__.__dict__["name"] = name
            if request_mapping_template is None and not opts.urn:
                raise TypeError("Missing required property 'request_mapping_template'")
            __props__.__dict__["request_mapping_template"] = request_mapping_template
            if response_mapping_template is None and not opts.urn:
                raise TypeError("Missing required property 'response_mapping_template'")
            __props__.__dict__["response_mapping_template"] = response_mapping_template
            __props__.__dict__["arn"] = None
            __props__.__dict__["function_id"] = None
        super(Function, __self__).__init__(
            'aws:appsync/function:Function',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            data_source: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            function_id: Optional[pulumi.Input[str]] = None,
            function_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            request_mapping_template: Optional[pulumi.Input[str]] = None,
            response_mapping_template: Optional[pulumi.Input[str]] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The ID of the associated AppSync API.
        :param pulumi.Input[str] arn: The ARN of the Function object.
        :param pulumi.Input[str] data_source: The Function DataSource name.
        :param pulumi.Input[str] description: The Function description.
        :param pulumi.Input[str] function_id: A unique ID representing the Function object.
        :param pulumi.Input[str] function_version: The version of the request mapping template. Currently the supported value is `2018-05-29`.
        :param pulumi.Input[str] name: The Function name. The function name does not have to be unique.
        :param pulumi.Input[str] request_mapping_template: The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        :param pulumi.Input[str] response_mapping_template: The Function response mapping template.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionState.__new__(_FunctionState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["data_source"] = data_source
        __props__.__dict__["description"] = description
        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["function_version"] = function_version
        __props__.__dict__["name"] = name
        __props__.__dict__["request_mapping_template"] = request_mapping_template
        __props__.__dict__["response_mapping_template"] = response_mapping_template
        return Function(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        The ID of the associated AppSync API.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Function object.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> pulumi.Output[str]:
        """
        The Function DataSource name.
        """
        return pulumi.get(self, "data_source")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Function description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[str]:
        """
        A unique ID representing the Function object.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of the request mapping template. Currently the supported value is `2018-05-29`.
        """
        return pulumi.get(self, "function_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Function name. The function name does not have to be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="requestMappingTemplate")
    def request_mapping_template(self) -> pulumi.Output[str]:
        """
        The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        """
        return pulumi.get(self, "request_mapping_template")

    @property
    @pulumi.getter(name="responseMappingTemplate")
    def response_mapping_template(self) -> pulumi.Output[str]:
        """
        The Function response mapping template.
        """
        return pulumi.get(self, "response_mapping_template")


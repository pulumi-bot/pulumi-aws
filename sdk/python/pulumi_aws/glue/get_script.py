# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetScriptResult:
    """
    A collection of values returned by getScript.
    """
    def __init__(__self__, dag_edges=None, dag_nodes=None, id=None, language=None, python_script=None, scala_code=None):
        if dag_edges and not isinstance(dag_edges, list):
            raise TypeError("Expected argument 'dag_edges' to be a list")
        __self__.dag_edges = dag_edges
        if dag_nodes and not isinstance(dag_nodes, list):
            raise TypeError("Expected argument 'dag_nodes' to be a list")
        __self__.dag_nodes = dag_nodes
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        __self__.language = language
        if python_script and not isinstance(python_script, str):
            raise TypeError("Expected argument 'python_script' to be a str")
        __self__.python_script = python_script
        """
        The Python script generated from the DAG when the `language` argument is set to `PYTHON`.
        """
        if scala_code and not isinstance(scala_code, str):
            raise TypeError("Expected argument 'scala_code' to be a str")
        __self__.scala_code = scala_code
        """
        The Scala code generated from the DAG when the `language` argument is set to `SCALA`.
        """


class AwaitableGetScriptResult(GetScriptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScriptResult(
            dag_edges=self.dag_edges,
            dag_nodes=self.dag_nodes,
            id=self.id,
            language=self.language,
            python_script=self.python_script,
            scala_code=self.scala_code)


def get_script(dag_edges=None, dag_nodes=None, language=None, opts=None):
    """
    Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).

    ## Example Usage
    ### Generate Python Script

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.glue.get_script(dag_edges=[
            aws.glue.GetScriptDagEdgeArgs(
                source="datasource0",
                target="applymapping1",
            ),
            aws.glue.GetScriptDagEdgeArgs(
                source="applymapping1",
                target="selectfields2",
            ),
            aws.glue.GetScriptDagEdgeArgs(
                source="selectfields2",
                target="resolvechoice3",
            ),
            aws.glue.GetScriptDagEdgeArgs(
                source="resolvechoice3",
                target="datasink4",
            ),
        ],
        dag_nodes=[
            aws.glue.GetScriptDagNodeArgs(
                args=[
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="database",
                        value=f"\"{aws_glue_catalog_database['source']['name']}\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="table_name",
                        value=f"\"{aws_glue_catalog_table['source']['name']}\"",
                    ),
                ],
                id="datasource0",
                node_type="DataSource",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[aws.glue.GetScriptDagNodeArgArgs(
                    name="mapping",
                    value="[(\"column1\", \"string\", \"column1\", \"string\")]",
                )],
                id="applymapping1",
                node_type="ApplyMapping",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[aws.glue.GetScriptDagNodeArgArgs(
                    name="paths",
                    value="[\"column1\"]",
                )],
                id="selectfields2",
                node_type="SelectFields",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="choice",
                        value="\"MATCH_CATALOG\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="database",
                        value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="table_name",
                        value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                    ),
                ],
                id="resolvechoice3",
                node_type="ResolveChoice",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="database",
                        value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="table_name",
                        value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                    ),
                ],
                id="datasink4",
                node_type="DataSink",
            ),
        ],
        language="PYTHON")
    pulumi.export("pythonScript", example.python_script)
    ```
    ### Generate Scala Code

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.glue.get_script(dag_edges=[
            aws.glue.GetScriptDagEdgeArgs(
                source="datasource0",
                target="applymapping1",
            ),
            aws.glue.GetScriptDagEdgeArgs(
                source="applymapping1",
                target="selectfields2",
            ),
            aws.glue.GetScriptDagEdgeArgs(
                source="selectfields2",
                target="resolvechoice3",
            ),
            aws.glue.GetScriptDagEdgeArgs(
                source="resolvechoice3",
                target="datasink4",
            ),
        ],
        dag_nodes=[
            aws.glue.GetScriptDagNodeArgs(
                args=[
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="database",
                        value=f"\"{aws_glue_catalog_database['source']['name']}\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="table_name",
                        value=f"\"{aws_glue_catalog_table['source']['name']}\"",
                    ),
                ],
                id="datasource0",
                node_type="DataSource",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[aws.glue.GetScriptDagNodeArgArgs(
                    name="mappings",
                    value="[(\"column1\", \"string\", \"column1\", \"string\")]",
                )],
                id="applymapping1",
                node_type="ApplyMapping",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[aws.glue.GetScriptDagNodeArgArgs(
                    name="paths",
                    value="[\"column1\"]",
                )],
                id="selectfields2",
                node_type="SelectFields",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="choice",
                        value="\"MATCH_CATALOG\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="database",
                        value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="table_name",
                        value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                    ),
                ],
                id="resolvechoice3",
                node_type="ResolveChoice",
            ),
            aws.glue.GetScriptDagNodeArgs(
                args=[
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="database",
                        value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                    ),
                    aws.glue.GetScriptDagNodeArgArgs(
                        name="table_name",
                        value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                    ),
                ],
                id="datasink4",
                node_type="DataSink",
            ),
        ],
        language="SCALA")
    pulumi.export("scalaCode", example.scala_code)
    ```


    :param list dag_edges: A list of the edges in the DAG. Defined below.
    :param list dag_nodes: A list of the nodes in the DAG. Defined below.
    :param str language: The programming language of the resulting code from the DAG. Defaults to `PYTHON`. Valid values are `PYTHON` and `SCALA`.

    The **dag_edges** object supports the following:

      * `source` (`str`) - The ID of the node at which the edge starts.
      * `target` (`str`) - The ID of the node at which the edge ends.
      * `targetParameter` (`str`) - The target of the edge.

    The **dag_nodes** object supports the following:

      * `args` (`list`) - Nested configuration an argument or property of a node. Defined below.
        * `name` (`str`) - The name of the argument or property.
        * `param` (`bool`) - Boolean if the value is used as a parameter. Defaults to `false`.
        * `value` (`str`) - The value of the argument or property.

      * `id` (`str`) - A node identifier that is unique within the node's graph.
      * `lineNumber` (`float`) - The line number of the node.
      * `node_type` (`str`) - The type of node this is.
    """
    __args__ = dict()
    __args__['dagEdges'] = dag_edges
    __args__['dagNodes'] = dag_nodes
    __args__['language'] = language
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:glue/getScript:getScript', __args__, opts=opts).value

    return AwaitableGetScriptResult(
        dag_edges=__ret__.get('dagEdges'),
        dag_nodes=__ret__.get('dagNodes'),
        id=__ret__.get('id'),
        language=__ret__.get('language'),
        python_script=__ret__.get('pythonScript'),
        scala_code=__ret__.get('scalaCode'))

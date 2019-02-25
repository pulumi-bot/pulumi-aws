# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetScriptResult:
    """
    A collection of values returned by getScript.
    """
    def __init__(__self__, python_script=None, scala_code=None, id=None):
        if python_script and not isinstance(python_script, str):
            raise TypeError('Expected argument python_script to be a str')
        __self__.python_script = python_script
        """
        The Python script generated from the DAG when the `language` argument is set to `PYTHON`.
        """
        if scala_code and not isinstance(scala_code, str):
            raise TypeError('Expected argument scala_code to be a str')
        __self__.scala_code = scala_code
        """
        The Scala code generated from the DAG when the `language` argument is set to `SCALA`.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_script(dag_edges=None,dag_nodes=None,language=None,opts=None):
    """
    Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).
    """
    __args__ = dict()

    __args__['dagEdges'] = dag_edges
    __args__['dagNodes'] = dag_nodes
    __args__['language'] = language
    __ret__ = await pulumi.runtime.invoke('aws:glue/getScript:getScript', __args__, opts=opts)

    return GetScriptResult(
        python_script=__ret__.get('pythonScript'),
        scala_code=__ret__.get('scalaCode'),
        id=__ret__.get('id'))

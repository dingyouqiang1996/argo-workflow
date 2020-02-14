# coding: utf-8

"""
    Argo

    Workflow Service API performs CRUD actions against application resources  # noqa: E501

    The version of the OpenAPI document: latest
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.io.argoproj.argo.client.model.workflow_workflow_terminate_request import WorkflowWorkflowTerminateRequest  # noqa: E501
from openapi_client.rest import ApiException

class TestWorkflowWorkflowTerminateRequest(unittest.TestCase):
    """WorkflowWorkflowTerminateRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test WorkflowWorkflowTerminateRequest
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.workflow_workflow_terminate_request.WorkflowWorkflowTerminateRequest()  # noqa: E501
        if include_optional :
            return WorkflowWorkflowTerminateRequest(
                name = '0', 
                namespace = '0'
            )
        else :
            return WorkflowWorkflowTerminateRequest(
        )

    def testWorkflowWorkflowTerminateRequest(self):
        """Test WorkflowWorkflowTerminateRequest"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

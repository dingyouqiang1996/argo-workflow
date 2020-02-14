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
from openapi_client.io.argoproj.argo.client.model.workflowv1alpha1_object_reference import Workflowv1alpha1ObjectReference  # noqa: E501
from openapi_client.rest import ApiException

class TestWorkflowv1alpha1ObjectReference(unittest.TestCase):
    """Workflowv1alpha1ObjectReference unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test Workflowv1alpha1ObjectReference
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.workflowv1alpha1_object_reference.Workflowv1alpha1ObjectReference()  # noqa: E501
        if include_optional :
            return Workflowv1alpha1ObjectReference(
                api_version = '0', 
                field_path = '0', 
                kind = '0', 
                name = '0', 
                namespace = '0', 
                resource_version = '0', 
                uid = '0'
            )
        else :
            return Workflowv1alpha1ObjectReference(
        )

    def testWorkflowv1alpha1ObjectReference(self):
        """Test Workflowv1alpha1ObjectReference"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

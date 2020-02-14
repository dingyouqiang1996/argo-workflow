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
from openapi_client.io.argoproj.argo.client.model.workflowv1alpha1_secret_volume_source import Workflowv1alpha1SecretVolumeSource  # noqa: E501
from openapi_client.rest import ApiException

class TestWorkflowv1alpha1SecretVolumeSource(unittest.TestCase):
    """Workflowv1alpha1SecretVolumeSource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test Workflowv1alpha1SecretVolumeSource
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.workflowv1alpha1_secret_volume_source.Workflowv1alpha1SecretVolumeSource()  # noqa: E501
        if include_optional :
            return Workflowv1alpha1SecretVolumeSource(
                default_mode = 56, 
                items = [
                    openapi_client.models.workflowv1alpha1_key_to_path.workflowv1alpha1KeyToPath(
                        key = '0', 
                        mode = 56, 
                        path = '0', )
                    ], 
                optional = True, 
                secret_name = '0'
            )
        else :
            return Workflowv1alpha1SecretVolumeSource(
        )

    def testWorkflowv1alpha1SecretVolumeSource(self):
        """Test Workflowv1alpha1SecretVolumeSource"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

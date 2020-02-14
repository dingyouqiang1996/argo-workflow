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
from openapi_client.io.argoproj.argo.client.model.v1_create_options import V1CreateOptions  # noqa: E501
from openapi_client.rest import ApiException

class TestV1CreateOptions(unittest.TestCase):
    """V1CreateOptions unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1CreateOptions
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1_create_options.V1CreateOptions()  # noqa: E501
        if include_optional :
            return V1CreateOptions(
                dry_run = [
                    '0'
                    ], 
                field_manager = '0', 
                type_meta = openapi_client.models.metav1_type_meta.metav1TypeMeta(
                    api_version = '0', 
                    kind = '0', )
            )
        else :
            return V1CreateOptions(
        )

    def testV1CreateOptions(self):
        """Test V1CreateOptions"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

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
from openapi_client.io.argoproj.argo.client.model.v1alpha1_ttl_strategy import V1alpha1TTLStrategy  # noqa: E501
from openapi_client.rest import ApiException

class TestV1alpha1TTLStrategy(unittest.TestCase):
    """V1alpha1TTLStrategy unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1alpha1TTLStrategy
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1alpha1_ttl_strategy.V1alpha1TTLStrategy()  # noqa: E501
        if include_optional :
            return V1alpha1TTLStrategy(
                seconds_after_completion = 56, 
                seconds_after_failure = 56, 
                seconds_after_success = 56
            )
        else :
            return V1alpha1TTLStrategy(
        )

    def testV1alpha1TTLStrategy(self):
        """Test V1alpha1TTLStrategy"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

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
from openapi_client.io.argoproj.argo.client.model.v1alpha1_flex_volume_source import V1alpha1FlexVolumeSource  # noqa: E501
from openapi_client.rest import ApiException

class TestV1alpha1FlexVolumeSource(unittest.TestCase):
    """V1alpha1FlexVolumeSource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1alpha1FlexVolumeSource
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1alpha1_flex_volume_source.V1alpha1FlexVolumeSource()  # noqa: E501
        if include_optional :
            return V1alpha1FlexVolumeSource(
                driver = '0', 
                fs_type = '0', 
                options = {
                    'key' : '0'
                    }, 
                read_only = True, 
                secret_ref = openapi_client.models.v1alpha1_local_object_reference.v1alpha1LocalObjectReference(
                    name = '0', )
            )
        else :
            return V1alpha1FlexVolumeSource(
        )

    def testV1alpha1FlexVolumeSource(self):
        """Test V1alpha1FlexVolumeSource"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

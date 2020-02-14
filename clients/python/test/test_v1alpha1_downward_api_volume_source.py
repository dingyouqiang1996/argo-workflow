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
from openapi_client.io.argoproj.argo.client.model.v1alpha1_downward_api_volume_source import V1alpha1DownwardAPIVolumeSource  # noqa: E501
from openapi_client.rest import ApiException

class TestV1alpha1DownwardAPIVolumeSource(unittest.TestCase):
    """V1alpha1DownwardAPIVolumeSource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1alpha1DownwardAPIVolumeSource
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1alpha1_downward_api_volume_source.V1alpha1DownwardAPIVolumeSource()  # noqa: E501
        if include_optional :
            return V1alpha1DownwardAPIVolumeSource(
                default_mode = 56, 
                items = [
                    openapi_client.models.v1alpha1_downward_api_volume_file.v1alpha1DownwardAPIVolumeFile(
                        field_ref = openapi_client.models.v1alpha1_object_field_selector.v1alpha1ObjectFieldSelector(
                            api_version = '0', 
                            field_path = '0', ), 
                        mode = 56, 
                        path = '0', 
                        resource_field_ref = openapi_client.models.v1alpha1_resource_field_selector.v1alpha1ResourceFieldSelector(
                            container_name = '0', 
                            divisor = openapi_client.models.resource_quantity.resourceQuantity(
                                string = '0', ), 
                            resource = '0', ), )
                    ]
            )
        else :
            return V1alpha1DownwardAPIVolumeSource(
        )

    def testV1alpha1DownwardAPIVolumeSource(self):
        """Test V1alpha1DownwardAPIVolumeSource"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

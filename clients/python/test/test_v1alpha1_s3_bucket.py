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
from openapi_client.io.argoproj.argo.client.model.v1alpha1_s3_bucket import V1alpha1S3Bucket  # noqa: E501
from openapi_client.rest import ApiException

class TestV1alpha1S3Bucket(unittest.TestCase):
    """V1alpha1S3Bucket unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1alpha1S3Bucket
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1alpha1_s3_bucket.V1alpha1S3Bucket()  # noqa: E501
        if include_optional :
            return V1alpha1S3Bucket(
                access_key_secret = openapi_client.models.workflowv1alpha1_secret_key_selector.workflowv1alpha1SecretKeySelector(
                    key = '0', 
                    local_object_reference = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    optional = True, ), 
                bucket = '0', 
                endpoint = '0', 
                insecure = True, 
                region = '0', 
                role_arn = '0', 
                secret_key_secret = openapi_client.models.workflowv1alpha1_secret_key_selector.workflowv1alpha1SecretKeySelector(
                    key = '0', 
                    local_object_reference = openapi_client.models.workflowv1alpha1_local_object_reference.workflowv1alpha1LocalObjectReference(
                        name = '0', ), 
                    optional = True, )
            )
        else :
            return V1alpha1S3Bucket(
        )

    def testV1alpha1S3Bucket(self):
        """Test V1alpha1S3Bucket"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

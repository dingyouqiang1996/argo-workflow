"""
    Argo Server API

    You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`  # noqa: E501

    The version of the OpenAPI document: VERSION
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from io.argoproj.workflow.apis.info_service_api import InfoServiceApi  # noqa: E501


class TestInfoServiceApi(unittest.TestCase):
    """InfoServiceApi unit test stubs"""

    def setUp(self):
        self.api = InfoServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_info_service_get_info(self):
        """Test case for info_service_get_info

        """
        pass

    def test_info_service_get_user_info(self):
        """Test case for info_service_get_user_info

        """
        pass

    def test_info_service_get_version(self):
        """Test case for info_service_get_version

        """
        pass


if __name__ == '__main__':
    unittest.main()

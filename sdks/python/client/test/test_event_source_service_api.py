"""
    Argo Server API

    You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`  # noqa: E501

    The version of the OpenAPI document: VERSION
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from io.argoproj.workflow.apis.event_source_service_api import EventSourceServiceApi  # noqa: E501


class TestEventSourceServiceApi(unittest.TestCase):
    """EventSourceServiceApi unit test stubs"""

    def setUp(self):
        self.api = EventSourceServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_event_source_service_create_event_source(self):
        """Test case for event_source_service_create_event_source

        """
        pass

    def test_event_source_service_delete_event_source(self):
        """Test case for event_source_service_delete_event_source

        """
        pass

    def test_event_source_service_event_sources_logs(self):
        """Test case for event_source_service_event_sources_logs

        """
        pass

    def test_event_source_service_get_event_source(self):
        """Test case for event_source_service_get_event_source

        """
        pass

    def test_event_source_service_list_event_sources(self):
        """Test case for event_source_service_list_event_sources

        """
        pass

    def test_event_source_service_update_event_source(self):
        """Test case for event_source_service_update_event_source

        """
        pass

    def test_event_source_service_watch_event_sources(self):
        """Test case for event_source_service_watch_event_sources

        """
        pass


if __name__ == '__main__':
    unittest.main()

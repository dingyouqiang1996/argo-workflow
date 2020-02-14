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
from openapi_client.io.argoproj.argo.client.model.v1_lifecycle import V1Lifecycle  # noqa: E501
from openapi_client.rest import ApiException

class TestV1Lifecycle(unittest.TestCase):
    """V1Lifecycle unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1Lifecycle
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1_lifecycle.V1Lifecycle()  # noqa: E501
        if include_optional :
            return V1Lifecycle(
                post_start = openapi_client.models.v1_handler.v1Handler(
                    exec = openapi_client.models.v1_exec_action.v1ExecAction(
                        command = [
                            '0'
                            ], ), 
                    http_get = openapi_client.models.v1_http_get_action.v1HTTPGetAction(
                        host = '0', 
                        http_headers = [
                            openapi_client.models.http_header_describes_a_custom_header_to_be_used_in_http_probes.HTTPHeader describes a custom header to be used in HTTP probes(
                                name = '0', 
                                value = '0', )
                            ], 
                        path = '0', 
                        port = openapi_client.models.int_or_string_is_a_type_that_can_hold_an_int32_or_a_string/__when_used_in
json_or_yaml_marshalling_and_unmarshalling,_it_produces_or_consumes_the
inner_type/__this_allows_you_to_have,_for_example,_a_json_field_that_can
accept_a_name_or_number/
todo:_rename_to_int32_or_string.IntOrString is a type that can hold an int32 or a string.  When used in
JSON or YAML marshalling and unmarshalling, it produces or consumes the
inner type.  This allows you to have, for example, a JSON field that can
accept a name or number.
TODO: Rename to Int32OrString(
                            int_val = 56, 
                            str_val = '0', 
                            type = '0', ), 
                        scheme = '0', ), 
                    tcp_socket = openapi_client.models.tcp_socket_action_describes_an_action_based_on_opening_a_socket.TCPSocketAction describes an action based on opening a socket(
                        host = '0', ), ), 
                pre_stop = openapi_client.models.v1_handler.v1Handler(
                    exec = openapi_client.models.v1_exec_action.v1ExecAction(
                        command = [
                            '0'
                            ], ), 
                    http_get = openapi_client.models.v1_http_get_action.v1HTTPGetAction(
                        host = '0', 
                        http_headers = [
                            openapi_client.models.http_header_describes_a_custom_header_to_be_used_in_http_probes.HTTPHeader describes a custom header to be used in HTTP probes(
                                name = '0', 
                                value = '0', )
                            ], 
                        path = '0', 
                        port = openapi_client.models.int_or_string_is_a_type_that_can_hold_an_int32_or_a_string/__when_used_in
json_or_yaml_marshalling_and_unmarshalling,_it_produces_or_consumes_the
inner_type/__this_allows_you_to_have,_for_example,_a_json_field_that_can
accept_a_name_or_number/
todo:_rename_to_int32_or_string.IntOrString is a type that can hold an int32 or a string.  When used in
JSON or YAML marshalling and unmarshalling, it produces or consumes the
inner type.  This allows you to have, for example, a JSON field that can
accept a name or number.
TODO: Rename to Int32OrString(
                            int_val = 56, 
                            str_val = '0', 
                            type = '0', ), 
                        scheme = '0', ), 
                    tcp_socket = openapi_client.models.tcp_socket_action_describes_an_action_based_on_opening_a_socket.TCPSocketAction describes an action based on opening a socket(
                        host = '0', ), )
            )
        else :
            return V1Lifecycle(
        )

    def testV1Lifecycle(self):
        """Test V1Lifecycle"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

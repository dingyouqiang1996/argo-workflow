# coding: utf-8

"""
    Argo Server API

    You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`  # noqa: E501

    The version of the OpenAPI document: VERSION
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import re  # noqa: F401

# python 2 and python 3 compatibility library
import six

from argo_workflows.api_client import ApiClient
from argo_workflows.exceptions import (  # noqa: F401
    ApiTypeError,
    ApiValueError
)


class ArtifactServiceApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

    def get_input_artifact(self, namespace, name, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an input artifact.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_input_artifact(namespace, name, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str namespace: (required)
        :param str name: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        return self.get_input_artifact_with_http_info(namespace, name, pod_name, artifact_name, **kwargs)  # noqa: E501

    def get_input_artifact_with_http_info(self, namespace, name, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an input artifact.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_input_artifact_with_http_info(namespace, name, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str namespace: (required)
        :param str name: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _return_http_data_only: response data without head status code
                                       and headers
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """

        local_var_params = locals()

        all_params = [
            'namespace',
            'name',
            'pod_name',
            'artifact_name'
        ]
        all_params.extend(
            [
                'async_req',
                '_return_http_data_only',
                '_preload_content',
                '_request_timeout'
            ]
        )

        for key, val in six.iteritems(local_var_params['kwargs']):
            if key not in all_params:
                raise ApiTypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_input_artifact" % key
                )
            local_var_params[key] = val
        del local_var_params['kwargs']
        # verify the required parameter 'namespace' is set
        if self.api_client.client_side_validation and ('namespace' not in local_var_params or  # noqa: E501
                                                        local_var_params['namespace'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `namespace` when calling `get_input_artifact`")  # noqa: E501
        # verify the required parameter 'name' is set
        if self.api_client.client_side_validation and ('name' not in local_var_params or  # noqa: E501
                                                        local_var_params['name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `name` when calling `get_input_artifact`")  # noqa: E501
        # verify the required parameter 'pod_name' is set
        if self.api_client.client_side_validation and ('pod_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['pod_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `pod_name` when calling `get_input_artifact`")  # noqa: E501
        # verify the required parameter 'artifact_name' is set
        if self.api_client.client_side_validation and ('artifact_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['artifact_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `artifact_name` when calling `get_input_artifact`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'namespace' in local_var_params:
            path_params['namespace'] = local_var_params['namespace']  # noqa: E501
        if 'name' in local_var_params:
            path_params['name'] = local_var_params['name']  # noqa: E501
        if 'pod_name' in local_var_params:
            path_params['podName'] = local_var_params['pod_name']  # noqa: E501
        if 'artifact_name' in local_var_params:
            path_params['artifactName'] = local_var_params['artifact_name']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/input-artifacts/{namespace}/{name}/{podName}/{artifactName}', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type=None,  # noqa: E501
            auth_settings=auth_settings,
            async_req=local_var_params.get('async_req'),
            _return_http_data_only=local_var_params.get('_return_http_data_only'),  # noqa: E501
            _preload_content=local_var_params.get('_preload_content', True),
            _request_timeout=local_var_params.get('_request_timeout'),
            collection_formats=collection_formats)

    def get_input_artifact_by_uid(self, namespace, uid, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an input artifact by UID.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_input_artifact_by_uid(namespace, uid, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str namespace: (required)
        :param str uid: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        return self.get_input_artifact_by_uid_with_http_info(namespace, uid, pod_name, artifact_name, **kwargs)  # noqa: E501

    def get_input_artifact_by_uid_with_http_info(self, namespace, uid, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an input artifact by UID.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_input_artifact_by_uid_with_http_info(namespace, uid, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str namespace: (required)
        :param str uid: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _return_http_data_only: response data without head status code
                                       and headers
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """

        local_var_params = locals()

        all_params = [
            'namespace',
            'uid',
            'pod_name',
            'artifact_name'
        ]
        all_params.extend(
            [
                'async_req',
                '_return_http_data_only',
                '_preload_content',
                '_request_timeout'
            ]
        )

        for key, val in six.iteritems(local_var_params['kwargs']):
            if key not in all_params:
                raise ApiTypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_input_artifact_by_uid" % key
                )
            local_var_params[key] = val
        del local_var_params['kwargs']
        # verify the required parameter 'namespace' is set
        if self.api_client.client_side_validation and ('namespace' not in local_var_params or  # noqa: E501
                                                        local_var_params['namespace'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `namespace` when calling `get_input_artifact_by_uid`")  # noqa: E501
        # verify the required parameter 'uid' is set
        if self.api_client.client_side_validation and ('uid' not in local_var_params or  # noqa: E501
                                                        local_var_params['uid'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `uid` when calling `get_input_artifact_by_uid`")  # noqa: E501
        # verify the required parameter 'pod_name' is set
        if self.api_client.client_side_validation and ('pod_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['pod_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `pod_name` when calling `get_input_artifact_by_uid`")  # noqa: E501
        # verify the required parameter 'artifact_name' is set
        if self.api_client.client_side_validation and ('artifact_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['artifact_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `artifact_name` when calling `get_input_artifact_by_uid`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'namespace' in local_var_params:
            path_params['namespace'] = local_var_params['namespace']  # noqa: E501
        if 'uid' in local_var_params:
            path_params['uid'] = local_var_params['uid']  # noqa: E501
        if 'pod_name' in local_var_params:
            path_params['podName'] = local_var_params['pod_name']  # noqa: E501
        if 'artifact_name' in local_var_params:
            path_params['artifactName'] = local_var_params['artifact_name']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/input-artifacts-by-uid/{uid}/{podName}/{artifactName}', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type=None,  # noqa: E501
            auth_settings=auth_settings,
            async_req=local_var_params.get('async_req'),
            _return_http_data_only=local_var_params.get('_return_http_data_only'),  # noqa: E501
            _preload_content=local_var_params.get('_preload_content', True),
            _request_timeout=local_var_params.get('_request_timeout'),
            collection_formats=collection_formats)

    def get_output_artifact(self, namespace, name, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an output artifact.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_output_artifact(namespace, name, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str namespace: (required)
        :param str name: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        return self.get_output_artifact_with_http_info(namespace, name, pod_name, artifact_name, **kwargs)  # noqa: E501

    def get_output_artifact_with_http_info(self, namespace, name, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an output artifact.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_output_artifact_with_http_info(namespace, name, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str namespace: (required)
        :param str name: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _return_http_data_only: response data without head status code
                                       and headers
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """

        local_var_params = locals()

        all_params = [
            'namespace',
            'name',
            'pod_name',
            'artifact_name'
        ]
        all_params.extend(
            [
                'async_req',
                '_return_http_data_only',
                '_preload_content',
                '_request_timeout'
            ]
        )

        for key, val in six.iteritems(local_var_params['kwargs']):
            if key not in all_params:
                raise ApiTypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_output_artifact" % key
                )
            local_var_params[key] = val
        del local_var_params['kwargs']
        # verify the required parameter 'namespace' is set
        if self.api_client.client_side_validation and ('namespace' not in local_var_params or  # noqa: E501
                                                        local_var_params['namespace'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `namespace` when calling `get_output_artifact`")  # noqa: E501
        # verify the required parameter 'name' is set
        if self.api_client.client_side_validation and ('name' not in local_var_params or  # noqa: E501
                                                        local_var_params['name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `name` when calling `get_output_artifact`")  # noqa: E501
        # verify the required parameter 'pod_name' is set
        if self.api_client.client_side_validation and ('pod_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['pod_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `pod_name` when calling `get_output_artifact`")  # noqa: E501
        # verify the required parameter 'artifact_name' is set
        if self.api_client.client_side_validation and ('artifact_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['artifact_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `artifact_name` when calling `get_output_artifact`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'namespace' in local_var_params:
            path_params['namespace'] = local_var_params['namespace']  # noqa: E501
        if 'name' in local_var_params:
            path_params['name'] = local_var_params['name']  # noqa: E501
        if 'pod_name' in local_var_params:
            path_params['podName'] = local_var_params['pod_name']  # noqa: E501
        if 'artifact_name' in local_var_params:
            path_params['artifactName'] = local_var_params['artifact_name']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/artifacts/{namespace}/{name}/{podName}/{artifactName}', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type=None,  # noqa: E501
            auth_settings=auth_settings,
            async_req=local_var_params.get('async_req'),
            _return_http_data_only=local_var_params.get('_return_http_data_only'),  # noqa: E501
            _preload_content=local_var_params.get('_preload_content', True),
            _request_timeout=local_var_params.get('_request_timeout'),
            collection_formats=collection_formats)

    def get_output_artifact_by_uid(self, uid, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an output artifact by UID.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_output_artifact_by_uid(uid, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str uid: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        return self.get_output_artifact_by_uid_with_http_info(uid, pod_name, artifact_name, **kwargs)  # noqa: E501

    def get_output_artifact_by_uid_with_http_info(self, uid, pod_name, artifact_name, **kwargs):  # noqa: E501
        """Get an output artifact by UID.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_output_artifact_by_uid_with_http_info(uid, pod_name, artifact_name, async_req=True)
        >>> result = thread.get()

        :param async_req bool: execute request asynchronously
        :param str uid: (required)
        :param str pod_name: (required)
        :param str artifact_name: (required)
        :param _return_http_data_only: response data without head status code
                                       and headers
        :param _preload_content: if False, the urllib3.HTTPResponse object will
                                 be returned without reading/decoding response
                                 data. Default is True.
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: None
                 If the method is called asynchronously,
                 returns the request thread.
        """

        local_var_params = locals()

        all_params = [
            'uid',
            'pod_name',
            'artifact_name'
        ]
        all_params.extend(
            [
                'async_req',
                '_return_http_data_only',
                '_preload_content',
                '_request_timeout'
            ]
        )

        for key, val in six.iteritems(local_var_params['kwargs']):
            if key not in all_params:
                raise ApiTypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_output_artifact_by_uid" % key
                )
            local_var_params[key] = val
        del local_var_params['kwargs']
        # verify the required parameter 'uid' is set
        if self.api_client.client_side_validation and ('uid' not in local_var_params or  # noqa: E501
                                                        local_var_params['uid'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `uid` when calling `get_output_artifact_by_uid`")  # noqa: E501
        # verify the required parameter 'pod_name' is set
        if self.api_client.client_side_validation and ('pod_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['pod_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `pod_name` when calling `get_output_artifact_by_uid`")  # noqa: E501
        # verify the required parameter 'artifact_name' is set
        if self.api_client.client_side_validation and ('artifact_name' not in local_var_params or  # noqa: E501
                                                        local_var_params['artifact_name'] is None):  # noqa: E501
            raise ApiValueError("Missing the required parameter `artifact_name` when calling `get_output_artifact_by_uid`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'uid' in local_var_params:
            path_params['uid'] = local_var_params['uid']  # noqa: E501
        if 'pod_name' in local_var_params:
            path_params['podName'] = local_var_params['pod_name']  # noqa: E501
        if 'artifact_name' in local_var_params:
            path_params['artifactName'] = local_var_params['artifact_name']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/artifacts-by-uid/{uid}/{podName}/{artifactName}', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type=None,  # noqa: E501
            auth_settings=auth_settings,
            async_req=local_var_params.get('async_req'),
            _return_http_data_only=local_var_params.get('_return_http_data_only'),  # noqa: E501
            _preload_content=local_var_params.get('_preload_content', True),
            _request_timeout=local_var_params.get('_request_timeout'),
            collection_formats=collection_formats)

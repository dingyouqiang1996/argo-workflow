# coding: utf-8

"""
    Argo Server API

    You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`  # noqa: E501

    The version of the OpenAPI document: VERSION
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from argo_workflows.configuration import Configuration


class GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'metadata': 'ObjectMeta',
        'spec': 'GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec',
        'status': 'GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus'
    }

    attribute_map = {
        'metadata': 'metadata',
        'spec': 'spec',
        'status': 'status'
    }

    def __init__(self, metadata=None, spec=None, status=None, local_vars_configuration=None):  # noqa: E501
        """GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._metadata = None
        self._spec = None
        self._status = None
        self.discriminator = None

        if metadata is not None:
            self.metadata = metadata
        if spec is not None:
            self.spec = spec
        if status is not None:
            self.status = status

    @property
    def metadata(self):
        """Gets the metadata of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501


        :return: The metadata of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501
        :rtype: ObjectMeta
        """
        return self._metadata

    @metadata.setter
    def metadata(self, metadata):
        """Sets the metadata of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.


        :param metadata: The metadata of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501
        :type: ObjectMeta
        """

        self._metadata = metadata

    @property
    def spec(self):
        """Gets the spec of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501


        :return: The spec of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501
        :rtype: GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec
        """
        return self._spec

    @spec.setter
    def spec(self, spec):
        """Sets the spec of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.


        :param spec: The spec of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501
        :type: GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec
        """

        self._spec = spec

    @property
    def status(self):
        """Gets the status of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501


        :return: The status of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501
        :rtype: GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus
        """
        return self._status

    @status.setter
    def status(self, status):
        """Sets the status of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.


        :param status: The status of this GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline.  # noqa: E501
        :type: GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus
        """

        self._status = status

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, GithubComArgoprojLabsArgoDataflowApiV1alpha1Pipeline):
            return True

        return self.to_dict() != other.to_dict()

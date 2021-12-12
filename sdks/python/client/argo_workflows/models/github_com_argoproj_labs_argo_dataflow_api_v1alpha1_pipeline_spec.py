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


class GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec(object):
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
        'deletion_delay': 'Duration',
        'steps': 'list[GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec]'
    }

    attribute_map = {
        'deletion_delay': 'deletionDelay',
        'steps': 'steps'
    }

    def __init__(self, deletion_delay=None, steps=None, local_vars_configuration=None):  # noqa: E501
        """GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._deletion_delay = None
        self._steps = None
        self.discriminator = None

        if deletion_delay is not None:
            self.deletion_delay = deletion_delay
        if steps is not None:
            self.steps = steps

    @property
    def deletion_delay(self):
        """Gets the deletion_delay of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.  # noqa: E501


        :return: The deletion_delay of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.  # noqa: E501
        :rtype: Duration
        """
        return self._deletion_delay

    @deletion_delay.setter
    def deletion_delay(self, deletion_delay):
        """Sets the deletion_delay of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.


        :param deletion_delay: The deletion_delay of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.  # noqa: E501
        :type: Duration
        """

        self._deletion_delay = deletion_delay

    @property
    def steps(self):
        """Gets the steps of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.  # noqa: E501


        :return: The steps of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.  # noqa: E501
        :rtype: list[GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec]
        """
        return self._steps

    @steps.setter
    def steps(self, steps):
        """Sets the steps of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.


        :param steps: The steps of this GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec.  # noqa: E501
        :type: list[GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec]
        """

        self._steps = steps

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
        if not isinstance(other, GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec):
            return True

        return self.to_dict() != other.to_dict()

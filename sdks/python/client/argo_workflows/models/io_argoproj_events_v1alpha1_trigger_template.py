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


class IoArgoprojEventsV1alpha1TriggerTemplate(object):
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
        'argo_workflow': 'IoArgoprojEventsV1alpha1ArgoWorkflowTrigger',
        'aws_lambda': 'IoArgoprojEventsV1alpha1AWSLambdaTrigger',
        'azure_event_hubs': 'IoArgoprojEventsV1alpha1AzureEventHubsTrigger',
        'conditions': 'str',
        'custom': 'IoArgoprojEventsV1alpha1CustomTrigger',
        'http': 'IoArgoprojEventsV1alpha1HTTPTrigger',
        'k8s': 'IoArgoprojEventsV1alpha1StandardK8STrigger',
        'kafka': 'IoArgoprojEventsV1alpha1KafkaTrigger',
        'log': 'IoArgoprojEventsV1alpha1LogTrigger',
        'name': 'str',
        'nats': 'IoArgoprojEventsV1alpha1NATSTrigger',
        'open_whisk': 'IoArgoprojEventsV1alpha1OpenWhiskTrigger',
        'pulsar': 'IoArgoprojEventsV1alpha1PulsarTrigger',
        'slack': 'IoArgoprojEventsV1alpha1SlackTrigger'
    }

    attribute_map = {
        'argo_workflow': 'argoWorkflow',
        'aws_lambda': 'awsLambda',
        'azure_event_hubs': 'azureEventHubs',
        'conditions': 'conditions',
        'custom': 'custom',
        'http': 'http',
        'k8s': 'k8s',
        'kafka': 'kafka',
        'log': 'log',
        'name': 'name',
        'nats': 'nats',
        'open_whisk': 'openWhisk',
        'pulsar': 'pulsar',
        'slack': 'slack'
    }

    def __init__(self, argo_workflow=None, aws_lambda=None, azure_event_hubs=None, conditions=None, custom=None, http=None, k8s=None, kafka=None, log=None, name=None, nats=None, open_whisk=None, pulsar=None, slack=None, local_vars_configuration=None):  # noqa: E501
        """IoArgoprojEventsV1alpha1TriggerTemplate - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._argo_workflow = None
        self._aws_lambda = None
        self._azure_event_hubs = None
        self._conditions = None
        self._custom = None
        self._http = None
        self._k8s = None
        self._kafka = None
        self._log = None
        self._name = None
        self._nats = None
        self._open_whisk = None
        self._pulsar = None
        self._slack = None
        self.discriminator = None

        if argo_workflow is not None:
            self.argo_workflow = argo_workflow
        if aws_lambda is not None:
            self.aws_lambda = aws_lambda
        if azure_event_hubs is not None:
            self.azure_event_hubs = azure_event_hubs
        if conditions is not None:
            self.conditions = conditions
        if custom is not None:
            self.custom = custom
        if http is not None:
            self.http = http
        if k8s is not None:
            self.k8s = k8s
        if kafka is not None:
            self.kafka = kafka
        if log is not None:
            self.log = log
        if name is not None:
            self.name = name
        if nats is not None:
            self.nats = nats
        if open_whisk is not None:
            self.open_whisk = open_whisk
        if pulsar is not None:
            self.pulsar = pulsar
        if slack is not None:
            self.slack = slack

    @property
    def argo_workflow(self):
        """Gets the argo_workflow of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The argo_workflow of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1ArgoWorkflowTrigger
        """
        return self._argo_workflow

    @argo_workflow.setter
    def argo_workflow(self, argo_workflow):
        """Sets the argo_workflow of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param argo_workflow: The argo_workflow of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1ArgoWorkflowTrigger
        """

        self._argo_workflow = argo_workflow

    @property
    def aws_lambda(self):
        """Gets the aws_lambda of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The aws_lambda of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1AWSLambdaTrigger
        """
        return self._aws_lambda

    @aws_lambda.setter
    def aws_lambda(self, aws_lambda):
        """Sets the aws_lambda of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param aws_lambda: The aws_lambda of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1AWSLambdaTrigger
        """

        self._aws_lambda = aws_lambda

    @property
    def azure_event_hubs(self):
        """Gets the azure_event_hubs of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The azure_event_hubs of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1AzureEventHubsTrigger
        """
        return self._azure_event_hubs

    @azure_event_hubs.setter
    def azure_event_hubs(self, azure_event_hubs):
        """Sets the azure_event_hubs of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param azure_event_hubs: The azure_event_hubs of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1AzureEventHubsTrigger
        """

        self._azure_event_hubs = azure_event_hubs

    @property
    def conditions(self):
        """Gets the conditions of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The conditions of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: str
        """
        return self._conditions

    @conditions.setter
    def conditions(self, conditions):
        """Sets the conditions of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param conditions: The conditions of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: str
        """

        self._conditions = conditions

    @property
    def custom(self):
        """Gets the custom of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The custom of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1CustomTrigger
        """
        return self._custom

    @custom.setter
    def custom(self, custom):
        """Sets the custom of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param custom: The custom of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1CustomTrigger
        """

        self._custom = custom

    @property
    def http(self):
        """Gets the http of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The http of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1HTTPTrigger
        """
        return self._http

    @http.setter
    def http(self, http):
        """Sets the http of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param http: The http of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1HTTPTrigger
        """

        self._http = http

    @property
    def k8s(self):
        """Gets the k8s of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The k8s of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1StandardK8STrigger
        """
        return self._k8s

    @k8s.setter
    def k8s(self, k8s):
        """Sets the k8s of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param k8s: The k8s of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1StandardK8STrigger
        """

        self._k8s = k8s

    @property
    def kafka(self):
        """Gets the kafka of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The kafka of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1KafkaTrigger
        """
        return self._kafka

    @kafka.setter
    def kafka(self, kafka):
        """Sets the kafka of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param kafka: The kafka of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1KafkaTrigger
        """

        self._kafka = kafka

    @property
    def log(self):
        """Gets the log of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The log of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1LogTrigger
        """
        return self._log

    @log.setter
    def log(self, log):
        """Sets the log of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param log: The log of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1LogTrigger
        """

        self._log = log

    @property
    def name(self):
        """Gets the name of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501

        Name is a unique name of the action to take.  # noqa: E501

        :return: The name of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this IoArgoprojEventsV1alpha1TriggerTemplate.

        Name is a unique name of the action to take.  # noqa: E501

        :param name: The name of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def nats(self):
        """Gets the nats of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The nats of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1NATSTrigger
        """
        return self._nats

    @nats.setter
    def nats(self, nats):
        """Sets the nats of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param nats: The nats of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1NATSTrigger
        """

        self._nats = nats

    @property
    def open_whisk(self):
        """Gets the open_whisk of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The open_whisk of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1OpenWhiskTrigger
        """
        return self._open_whisk

    @open_whisk.setter
    def open_whisk(self, open_whisk):
        """Sets the open_whisk of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param open_whisk: The open_whisk of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1OpenWhiskTrigger
        """

        self._open_whisk = open_whisk

    @property
    def pulsar(self):
        """Gets the pulsar of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The pulsar of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1PulsarTrigger
        """
        return self._pulsar

    @pulsar.setter
    def pulsar(self, pulsar):
        """Sets the pulsar of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param pulsar: The pulsar of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1PulsarTrigger
        """

        self._pulsar = pulsar

    @property
    def slack(self):
        """Gets the slack of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501


        :return: The slack of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :rtype: IoArgoprojEventsV1alpha1SlackTrigger
        """
        return self._slack

    @slack.setter
    def slack(self, slack):
        """Sets the slack of this IoArgoprojEventsV1alpha1TriggerTemplate.


        :param slack: The slack of this IoArgoprojEventsV1alpha1TriggerTemplate.  # noqa: E501
        :type: IoArgoprojEventsV1alpha1SlackTrigger
        """

        self._slack = slack

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
        if not isinstance(other, IoArgoprojEventsV1alpha1TriggerTemplate):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, IoArgoprojEventsV1alpha1TriggerTemplate):
            return True

        return self.to_dict() != other.to_dict()

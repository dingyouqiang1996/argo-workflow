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
from openapi_client.io.argoproj.argo.client.model.v1_pod_affinity import V1PodAffinity  # noqa: E501
from openapi_client.rest import ApiException

class TestV1PodAffinity(unittest.TestCase):
    """V1PodAffinity unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1PodAffinity
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.v1_pod_affinity.V1PodAffinity()  # noqa: E501
        if include_optional :
            return V1PodAffinity(
                preferred_during_scheduling_ignored_during_execution = [
                    openapi_client.models.the_weights_of_all_of_the_matched_weighted_pod_affinity_term_fields_are_added_per_node_to_find_the_most_preferred_node(s).The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)(
                        pod_affinity_term = openapi_client.models.defines_a_set_of_pods_(namely_those_matching_the_label_selector
relative_to_the_given_namespace(s))_that_this_pod_should_be
co_located_(affinity)_or_not_co_located_(anti_affinity)_with,
where_co_located_is_defined_as_running_on_a_node_whose_value_of
the_label_with_key_<topology_key>_matches_that_of_any_node_on_which
a_pod_of_the_set_of_pods_is_running.Defines a set of pods (namely those matching the labelSelector
relative to the given namespace(s)) that this pod should be
co-located (affinity) or not co-located (anti-affinity) with,
where co-located is defined as running on a node whose value of
the label with key <topologyKey> matches that of any node on which
a pod of the set of pods is running(
                            label_selector = openapi_client.models.v1_label_selector.v1LabelSelector(
                                match_expressions = [
                                    openapi_client.models.v1_label_selector_requirement.v1LabelSelectorRequirement(
                                        key = '0', 
                                        operator = '0', 
                                        values = [
                                            '0'
                                            ], )
                                    ], 
                                match_labels = {
                                    'key' : '0'
                                    }, ), 
                            namespaces = [
                                '0'
                                ], 
                            topology_key = '0', ), 
                        weight = 56, )
                    ], 
                required_during_scheduling_ignored_during_execution = [
                    openapi_client.models.defines_a_set_of_pods_(namely_those_matching_the_label_selector
relative_to_the_given_namespace(s))_that_this_pod_should_be
co_located_(affinity)_or_not_co_located_(anti_affinity)_with,
where_co_located_is_defined_as_running_on_a_node_whose_value_of
the_label_with_key_<topology_key>_matches_that_of_any_node_on_which
a_pod_of_the_set_of_pods_is_running.Defines a set of pods (namely those matching the labelSelector
relative to the given namespace(s)) that this pod should be
co-located (affinity) or not co-located (anti-affinity) with,
where co-located is defined as running on a node whose value of
the label with key <topologyKey> matches that of any node on which
a pod of the set of pods is running(
                        label_selector = openapi_client.models.v1_label_selector.v1LabelSelector(
                            match_expressions = [
                                openapi_client.models.v1_label_selector_requirement.v1LabelSelectorRequirement(
                                    key = '0', 
                                    operator = '0', 
                                    values = [
                                        '0'
                                        ], )
                                ], 
                            match_labels = {
                                'key' : '0'
                                }, ), 
                        namespaces = [
                            '0'
                            ], 
                        topology_key = '0', )
                    ]
            )
        else :
            return V1PodAffinity(
        )

    def testV1PodAffinity(self):
        """Test V1PodAffinity"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()

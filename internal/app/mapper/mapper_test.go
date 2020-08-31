package mapper

import (
	"fmt"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/xds-relay/internal/app/transport"
	"github.com/envoyproxy/xds-relay/internal/pkg/stats"
	aggregationv1 "github.com/envoyproxy/xds-relay/pkg/api/aggregation/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

type KeyerConfiguration = aggregationv1.KeyerConfiguration
type Fragment = aggregationv1.KeyerConfiguration_Fragment
type FragmentRule = aggregationv1.KeyerConfiguration_Fragment_Rule
type MatchPredicate = aggregationv1.MatchPredicate
type ResultPredicate = aggregationv1.ResultPredicate

const (
	clusterTypeURL   = "type.googleapis.com/envoy.api.v2.Cluster"
	listenerTypeURL  = "type.googleapis.com/envoy.api.v2.Listener"
	nodeid           = "nodeid"
	nodecluster      = "cluster"
	noderegion       = "region"
	nodezone         = "zone"
	nodesubzone      = "subzone"
	resource1        = "resource1"
	resource2        = "resource2"
	stringFragment   = "stringFragment"
	nodeIDField      = aggregationv1.NodeFieldType_NODE_ID
	nodeClusterField = aggregationv1.NodeFieldType_NODE_CLUSTER
	nodeRegionField  = aggregationv1.NodeFieldType_NODE_LOCALITY_REGION
	nodeZoneField    = aggregationv1.NodeFieldType_NODE_LOCALITY_ZONE
	nodeSubZoneField = aggregationv1.NodeFieldType_NODE_LOCALITY_SUBZONE
)

var positiveTests = []TableEntry{
	{
		Description: "AnyMatch returns StringFragment",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestTypeMatch Matches with a single typeurl",
		Parameters: []interface{}{
			getRequestTypeMatch([]string{clusterTypeURL}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestTypeMatch Matches with first among multiple typeurl",
		Parameters: []interface{}{
			getRequestTypeMatch([]string{clusterTypeURL, listenerTypeURL}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestTypeMatch Matches with second among multiple typeurl",
		Parameters: []interface{}{
			getRequestTypeMatch([]string{listenerTypeURL, clusterTypeURL}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestNodeMatch with node id exact match",
		Parameters: []interface{}{
			getRequestNodeIDExactMatch(nodeid),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestNodeMatch with node cluster exact match",
		Parameters: []interface{}{
			getRequestNodeClusterExactMatch(nodecluster),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestNodeMatch with node locality match",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion, nodezone, nodesubzone),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestNodeMatch with node id regex match",
		Parameters: []interface{}{
			getRequestNodeIDRegexMatch("n....d"),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "RequestNodeMatch with node cluster regex match",
		Parameters: []interface{}{
			getRequestNodeClusterRegexMatch("c*s*r"),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "AndMatch RequestNodeMatch",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeIDExactMatch(nodeid),
					getRequestNodeClusterExactMatch(nodecluster),
				}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "AndMatch recursive",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeIDExactMatch(nodeid),
						getRequestNodeClusterExactMatch(nodecluster),
					}),
					getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeLocality(noderegion, nodezone, ""),
					}),
				}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "OrMatch RequestNodeMatch first predicate",
		Parameters: []interface{}{
			getRequestNodeOrMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeIDExactMatch(""),
					getRequestNodeClusterExactMatch(nodecluster),
				}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "OrMatch RequestNodeMatch second predicate",
		Parameters: []interface{}{
			getRequestNodeOrMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeIDExactMatch(nodeid),
					getRequestNodeClusterExactMatch(""),
				}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "OrMatch recursive",
		Parameters: []interface{}{
			getRequestNodeOrMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeIDExactMatch(""),
						getRequestNodeClusterExactMatch(nodecluster),
					}),
					getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeLocality(noderegion, "", ""),
					}),
				}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "AndMatch OrMatch recursive combined",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeIDExactMatch(""),
						getRequestNodeClusterExactMatch(nodecluster),
					}),
					getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeLocality(noderegion, nodezone, ""),
					}),
				}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "NotMatch RequestType",
		Parameters: []interface{}{
			getRequestTypeNotMatch([]string{listenerTypeURL}),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node id regex",
		Parameters: []interface{}{
			getRequestNodeRegexNotMatch(nodeIDField),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node cluster regex",
		Parameters: []interface{}{
			getRequestNodeRegexNotMatch(nodeClusterField),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node locality",
		Parameters: []interface{}{
			getRequestNodeLocalityNotMatch(),
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "Not Match recursive",
		Parameters: []interface{}{
			&MatchPredicate{
				Type: &aggregationv1.MatchPredicate_NotMatch{
					NotMatch: getRequestNodeAndMatch(
						[]*aggregationv1.MatchPredicate{
							getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
								getRequestNodeIDExactMatch(""),
								getRequestNodeClusterExactMatch(""),
							}),
							getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
								getRequestNodeLocalityNotMatch(),
							}),
						}),
				},
			},
			getResultStringFragment(),
			clusterTypeURL,
			stringFragment,
		},
	},
	{
		Description: "AnyMatch With exact Node Id result",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeIDFragment(getExactAction()),
			clusterTypeURL,
			nodeid,
		},
	},
	{
		Description: "AnyMatch With regex Node Id result",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeIDFragment(getRegexAction("n....d", "replace")),
			clusterTypeURL,
			"replace",
		},
	},
	{
		Description: "AnyMatch With exact Node Cluster match",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeClusterFragment(getExactAction()),
			clusterTypeURL,
			nodecluster,
		},
	},
	{
		Description: "AnyMatch With regex Node Cluster match",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeClusterFragment(getRegexAction("c.*r", "replace")),
			clusterTypeURL,
			"replace",
		},
	},
	// {
	// 	Description: "AnyMatch With Exact Node Region match",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeRegionField, getExactAction()),
	// 		clusterTypeURL,
	// 		noderegion,
	// 	},
	// },
	// {
	// 	Description: "AnyMatch With regex Node Region match",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeRegionField, getRegexAction("r(egion)", "p$1")),
	// 		clusterTypeURL,
	// 		"pegion",
	// 	},
	// },
	// {
	// 	Description: "AnyMatch With exact Node Zone match",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeZoneField, getExactAction()),
	// 		clusterTypeURL,
	// 		nodezone,
	// 	},
	// },
	// {
	// 	Description: "AnyMatch With regex Node Zone match",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeZoneField, getRegexAction("z..e$", "newzone")),
	// 		clusterTypeURL,
	// 		"newzone",
	// 	},
	// },
	// {
	// 	Description: "AnyMatch With exact Node Subzone match",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeSubZoneField, getExactAction()),
	// 		clusterTypeURL,
	// 		nodesubzone,
	// 	},
	// },
	// {
	// 	Description: "AnyMatch With regex Node Subzone match",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeSubZoneField, getRegexAction("[^0-9](u|v)b{1}...e", "zone")),
	// 		clusterTypeURL,
	// 		"zone",
	// 	},
	// },
	{
		Description: "AnyMatch With result concatenation",
		Parameters: []interface{}{
			getAnyMatch(true),
			getRepeatedResultPredicate1(),
			clusterTypeURL,
			nodeid + nodecluster,
		},
	},
	// {
	// 	Description: "AnyMatch With result concatenation recursive",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getRepeatedResultPredicate2(),
	// 		clusterTypeURL,
	// 		"str" + noderegion + nodezone + "nTid" + nodecluster,
	// 	},
	// },
	{
		Description: "AnyMatch With resource names fragment element 0",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResourceNameFragment(0, getExactAction()),
			clusterTypeURL,
			resource1,
		},
	},
	{
		Description: "AnyMatch With resource names fragment element 1",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResourceNameFragment(1, getExactAction()),
			clusterTypeURL,
			resource2,
		},
	},
	{
		Description: "AnyMatch With resource names fragment regex",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResourceNameFragment(0, getRegexAction("^"+resource1+"$", "rsrc")),
			clusterTypeURL,
			"rsrc",
		},
	},
	{
		Description: "AnyMatch With result concatenation of resource name fragment",
		Parameters: []interface{}{
			getAnyMatch(true),
			getRepeatedResultPredicate3(),
			clusterTypeURL,
			resource1 + resource2,
		},
	},
}

var multiFragmentPositiveTests = []TableEntry{
	{
		Description: "both fragments match",
		Parameters: []interface{}{
			getAnyMatch(true),
			getAnyMatch(true),
			getResultStringFragment(),
			getResultStringFragment(),
			stringFragment + "_" + stringFragment,
		},
	},
	{
		Description: "first fragment match",
		Parameters: []interface{}{
			getAnyMatch(true),
			getAnyMatch(false),
			getResultStringFragment(),
			getResultStringFragment(),
			stringFragment,
		},
	},
	{
		Description: "second fragment match",
		Parameters: []interface{}{
			getAnyMatch(false),
			getAnyMatch(true),
			getResultStringFragment(),
			getResultStringFragment(),
			stringFragment,
		},
	},
}

var negativeTests = []TableEntry{
	{
		Description: "AnyMatch returns empty String",
		Parameters: []interface{}{
			getAnyMatch(false),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestTypeMatch does not match with unmatched typeurl",
		Parameters: []interface{}{
			getRequestTypeMatch([]string{""}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node id does not match",
		Parameters: []interface{}{
			getRequestNodeIDExactMatch(nodeid + "{5}"),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node cluster does not match",
		Parameters: []interface{}{
			getRequestNodeClusterExactMatch("[^a-z]odecluster"),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node region does not match",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion+"\\d", nodezone, nodesubzone),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node zone does not match",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion, "zon[A-Z]", nodesubzone),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node subzone does not match",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion, nodezone, nodesubzone+"+"),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node id regex does not match",
		Parameters: []interface{}{
			getRequestNodeIDRegexMatch(nodeid + "{5}"),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with node cluster regex does not match",
		Parameters: []interface{}{
			getRequestNodeClusterRegexMatch("[^a-z]odecluster"),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "RequestNodeMatch with exact match request node id mismatch",
		Parameters: []interface{}{
			getRequestNodeIDExactMatch(nodeid),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode("mismatch", nodecluster, noderegion, nodezone, nodesubzone)),
		},
	},
	{
		Description: "RequestNodeMatch with regex match request node id mismatch",
		Parameters: []interface{}{
			getRequestNodeIDRegexMatch(nodeid),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode("mismatch", nodecluster, noderegion, nodezone, nodesubzone)),
		},
	},
	{
		Description: "RequestNodeMatch with exact match request node cluster mismatch",
		Parameters: []interface{}{
			getRequestNodeClusterExactMatch(nodecluster),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode(nodeid, "mismatch", noderegion, nodezone, nodesubzone)),
		},
	},
	{
		Description: "RequestNodeMatch with regex match request node cluster mismatch",
		Parameters: []interface{}{
			getRequestNodeClusterRegexMatch(nodecluster),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode(nodeid, "mismatch", noderegion, nodezone, nodesubzone)),
		},
	},
	{
		Description: "RequestNodeMatch with exact match request node region mismatch",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion, nodezone, nodesubzone),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, "mismatch", nodezone, nodesubzone)),
		},
	},
	{
		Description: "RequestNodeMatch with exact match request node zone mismatch",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion, nodezone, nodesubzone),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, noderegion, "mismatch", nodesubzone)),
		},
	},
	{
		Description: "RequestNodeMatch with exact match request node subzone mismatch",
		Parameters: []interface{}{
			getRequestNodeLocality(noderegion, nodezone, nodesubzone),
			getResultStringFragment(),
			getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, noderegion, nodezone, "mismatch")),
		},
	},
	{
		Description: "AndMatch RequestNodeMatch does not match first predicate",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeIDExactMatch("nonmatchingnode"),
					getRequestNodeClusterExactMatch(nodecluster)}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "AndMatch RequestNodeMatch does not match second predicate",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeIDExactMatch(nodeid),
					getRequestNodeClusterExactMatch("nomatch")}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "AndMatch recursive",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeIDExactMatch("nonmatchingnode"),
						getRequestNodeClusterExactMatch(nodecluster),
					}),
					getRequestNodeClusterExactMatch(nodecluster)}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "OrMatch RequestNodeMatch does not match",
		Parameters: []interface{}{
			getRequestNodeOrMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeIDExactMatch(""),
					getRequestNodeClusterExactMatch("")}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "OrMatch recursive",
		Parameters: []interface{}{
			getRequestNodeOrMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeIDExactMatch(""),
						getRequestNodeClusterExactMatch(""),
					}),
					getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeLocality("nomatch", "", ""),
					}),
				}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "AndMatch OrMatch recursive combined",
		Parameters: []interface{}{
			getRequestNodeAndMatch(
				[]*aggregationv1.MatchPredicate{
					getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeIDExactMatch(""),
						getRequestNodeClusterExactMatch(""),
					}),
					getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
						getRequestNodeLocality("", "", ""),
					}),
				}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "NotMatch RequestType",
		Parameters: []interface{}{
			getRequestTypeNotMatch([]string{clusterTypeURL}),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node id regex",
		Parameters: []interface{}{
			getRequestNodeExactNotMatch(nodeIDField, nodeid),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node cluster regex",
		Parameters: []interface{}{
			getRequestNodeExactNotMatch(nodeClusterField, nodecluster),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node region regex",
		Parameters: []interface{}{
			getRequestNodeExactNotMatch(nodeRegionField, noderegion),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node zone regex",
		Parameters: []interface{}{
			getRequestNodeExactNotMatch(nodeZoneField, nodezone),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "Not Match RequestNodeMatch with node subzone regex",
		Parameters: []interface{}{
			getRequestNodeExactNotMatch(nodeSubZoneField, nodesubzone),
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
	{
		Description: "Not Match recursive",
		Parameters: []interface{}{
			&MatchPredicate{
				Type: &aggregationv1.MatchPredicate_NotMatch{
					NotMatch: getRequestNodeAndMatch(
						[]*aggregationv1.MatchPredicate{
							getRequestNodeOrMatch([]*aggregationv1.MatchPredicate{
								getRequestNodeIDExactMatch(nodeid),
								getRequestNodeClusterExactMatch(nodecluster),
							}),
							getRequestNodeAndMatch([]*aggregationv1.MatchPredicate{
								getRequestNodeLocality(noderegion, nodezone, ""),
							}),
						}),
				},
			},
			getResultStringFragment(),
			getDiscoveryRequest(),
		},
	},
}

var multiFragmentNegativeTests = []TableEntry{
	{
		Description: "no fragments match",
		Parameters: []interface{}{
			getAnyMatch(false),
			getAnyMatch(false),
			getResultStringFragment(),
			getResultStringFragment(),
		},
	},
}

var regexpErrorCases = []TableEntry{
	{
		Description: "Regex compilation failure in Nodefragment should return error",
		Parameters: []interface{}{
			getRequestNodeIDRegexMatch("\xbd\xb2"),
			getResultStringFragment(),
		},
	},
	{
		Description: "Regex compilation failure in RegexAction pattern should return error",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeIDFragment(getRegexAction("\xbd\xb2", "")),
		},
	},
}

var regexpErrorCasesMultipleFragments = []TableEntry{
	{
		Description: "Regex parse failure in first request predicate",
		Parameters: []interface{}{
			getRequestNodeIDRegexMatch("\xbd\xb2"),
			getAnyMatch(true),
			getResultStringFragment(),
			getResultStringFragment(),
		},
	},
	{
		Description: "Regex parse failure in second request predicate",
		Parameters: []interface{}{
			getAnyMatch(true),
			getRequestNodeIDRegexMatch("\xbd\xb2"),
			getResultStringFragment(),
			getResultStringFragment(),
		},
	},
	{
		Description: "Regex parse failure in first response predicate",
		Parameters: []interface{}{
			getAnyMatch(true),
			getAnyMatch(true),
			getResultRequestNodeIDFragment(getRegexAction("\xbd\xb2", "")),
			getResultStringFragment(),
		},
	},
	{
		Description: "Regex parse failure in second response predicate",
		Parameters: []interface{}{
			getAnyMatch(true),
			getAnyMatch(true),
			getResultStringFragment(),
			getResultRequestNodeIDFragment(getRegexAction("\xbd\xb2", "")),
		},
	},
	{
		Description: "Regex parse failure in resource name fragment",
		Parameters: []interface{}{
			getAnyMatch(true),
			getAnyMatch(true),
			getResultStringFragment(),
			getResourceNameFragment(0, getRegexAction("\xbd\xb2", "")),
		},
	},
}

var emptyFragmentErrorCases = []TableEntry{
	{
		Description: "empty node id in request predicate",
		Parameters: []interface{}{
			getRequestNodeIDExactMatch(nodeid),
			getResultRequestNodeIDFragment(getExactAction()),
			getDiscoveryRequestWithNode(getNode("", nodecluster, noderegion, nodezone, nodesubzone)),
			"MatchPredicate Node field cannot be empty",
		},
	},
	{
		Description: "empty node cluster in request predicate",
		Parameters: []interface{}{
			getRequestNodeClusterExactMatch(nodecluster),
			getResultRequestNodeIDFragment(getExactAction()),
			getDiscoveryRequestWithNode(getNode(nodeid, "", noderegion, nodezone, nodesubzone)),
			"MatchPredicate Node field cannot be empty",
		},
	},
	{
		Description: "empty node id in response action",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeIDFragment(getExactAction()),
			getDiscoveryRequestWithNode(getNode("", nodecluster, noderegion, nodezone, nodesubzone)),
			"RequestNodeFragment exact match resulted in an empty fragment",
		},
	},
	{
		Description: "empty node cluster in response action",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeClusterFragment(getExactAction()),
			getDiscoveryRequestWithNode(getNode(nodeid, "", noderegion, nodezone, nodesubzone)),
			"RequestNodeFragment exact match resulted in an empty fragment",
		},
	},
	// {
	// 	Description: "empty node region in response action",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeRegionField, getExactAction()),
	// 		getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, "", nodezone, nodesubzone)),
	// 		"RequestNodeFragment exact match resulted in an empty fragment",
	// 	},
	// },
	// {
	// 	Description: "empty node zone in response action",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeZoneField, getExactAction()),
	// 		getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, noderegion, "", nodesubzone)),
	// 		"RequestNodeFragment exact match resulted in an empty fragment",
	// 	},
	// },
	// {
	// 	Description: "empty node subzone in response action",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeSubZoneField, getExactAction()),
	// 		getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, noderegion, nodezone, "")),
	// 		"RequestNodeFragment exact match resulted in an empty fragment",
	// 	},
	// },
	{
		Description: "empty node id in regex response action",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeIDFragment(getRegexAction(nodeid, "")),
			getDiscoveryRequest(),
			"RequestNodeFragment regex match resulted in an empty fragment",
		},
	},
	{
		Description: "empty node cluster in regex response action",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResultRequestNodeClusterFragment(getRegexAction(nodecluster, "")),
			getDiscoveryRequest(),
			"RequestNodeFragment regex match resulted in an empty fragment",
		},
	},
	// {
	// 	Description: "empty node region in regex response action",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeRegionField, getRegexAction(noderegion, "")),
	// 		getDiscoveryRequest(),
	// 		"RequestNodeFragment regex match resulted in an empty fragment",
	// 	},
	// },
	// {
	// 	Description: "empty node zone in regex response action",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeZoneField, getRegexAction(nodezone, "")),
	// 		getDiscoveryRequest(),
	// 		"RequestNodeFragment regex match resulted in an empty fragment",
	// 	},
	// },
	// {
	// 	Description: "empty node subzone in regex response action",
	// 	Parameters: []interface{}{
	// 		getAnyMatch(true),
	// 		getResultRequestNodeFragment(nodeSubZoneField, getRegexAction(nodesubzone, "")),
	// 		getDiscoveryRequest(),
	// 		"RequestNodeFragment regex match resulted in an empty fragment",
	// 	},
	// },
	{
		Description: "resource fragment is negative",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResourceNameFragment(-1, getExactAction()),
			getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, "", nodezone, nodesubzone)),
			"ResourceNamesFragment.Element cannot be negative or larger than length",
		},
	},
	{
		Description: "resource fragment is above range",
		Parameters: []interface{}{
			getAnyMatch(true),
			getResourceNameFragment(10, getExactAction()),
			getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, "", nodezone, nodesubzone)),
			"ResourceNamesFragment.Element cannot be negative or larger than length",
		},
	},
}

var _ = Describe("GetKey", func() {
	DescribeTable("should be able to return fragment for",
		func(match *MatchPredicate, result *ResultPredicate, typeurl string, assert string) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match,
								Result: result,
							},
						},
					},
				},
			}
			mockScope := stats.NewMockScope("mock")
			mapper := New(&protoConfig, mockScope)
			request := getDiscoveryRequest()
			request.TypeUrl = typeurl
			key, err := mapper.GetKey(transport.NewRequestV2(&request))
			Expect(mockScope.Snapshot().Counters()["mock.mapper.success+"].Value()).To(Equal(int64(1)))
			Expect(key).To(Equal(assert))
			Expect(err).Should(BeNil())
		}, positiveTests...)

	DescribeTable("should be able to return error",
		func(match *MatchPredicate, result *ResultPredicate, request v2.DiscoveryRequest) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match,
								Result: result,
							},
						},
					},
				},
			}
			mockScope := stats.NewMockScope("mock")
			mapper := New(&protoConfig, mockScope)
			key, err := mapper.GetKey(transport.NewRequestV2(&request))
			Expect(mockScope.Snapshot().Counters()["mock.mapper.error+"].Value()).To(Equal(int64(1)))
			Expect(key).To(Equal(""))
			Expect(err).Should(Equal(fmt.Errorf("Cannot map the input to a key")))
		},
		negativeTests...)

	DescribeTable("should be able to join multiple fragments",
		func(match1 *MatchPredicate,
			match2 *MatchPredicate,
			result1 *ResultPredicate,
			result2 *ResultPredicate,
			expectedKey string) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match1,
								Result: result1,
							},
						},
					},
					{
						Rules: []*FragmentRule{
							{
								Match:  match2,
								Result: result2,
							},
						},
					},
				},
			}
			mapper := New(&protoConfig, stats.NewMockScope(""))
			req := getDiscoveryRequest()
			key, err := mapper.GetKey(transport.NewRequestV2(&req))
			Expect(expectedKey).To(Equal(key))
			Expect(err).Should(BeNil())
		},
		multiFragmentPositiveTests...)

	DescribeTable("should be able to return error for non matching multiple fragments",
		func(match1 *MatchPredicate, match2 *MatchPredicate, result1 *ResultPredicate, result2 *ResultPredicate) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match1,
								Result: result1,
							},
						},
					},
					{
						Rules: []*FragmentRule{
							{
								Match:  match2,
								Result: result2,
							},
						},
					},
				},
			}
			mapper := New(&protoConfig, stats.NewMockScope(""))
			req := getDiscoveryRequest()
			key, err := mapper.GetKey(transport.NewRequestV2(&req))
			Expect(key).To(Equal(""))
			Expect(err).Should(Equal(fmt.Errorf("Cannot map the input to a key")))
		},
		multiFragmentNegativeTests...)

	DescribeTable("should be able to return error for regex evaluation failures",
		func(match *MatchPredicate, result *ResultPredicate) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match,
								Result: result,
							},
						},
					},
				},
			}
			mapper := New(&protoConfig, stats.NewMockScope(""))
			req := getDiscoveryRequest()
			key, err := mapper.GetKey(transport.NewRequestV2(&req))
			Expect(key).To(Equal(""))
			Expect(err.Error()).Should(Equal("error parsing regexp: invalid UTF-8: `\xbd\xb2`"))
		},
		regexpErrorCases...)

	DescribeTable("should return error for multiple fragments regex failure",
		func(match1 *MatchPredicate, match2 *MatchPredicate, result1 *ResultPredicate, result2 *ResultPredicate) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match1,
								Result: result1,
							},
						},
					},
					{
						Rules: []*FragmentRule{
							{
								Match:  match2,
								Result: result2,
							},
						},
					},
				},
			}
			mapper := New(&protoConfig, stats.NewMockScope(""))
			req := getDiscoveryRequest()
			key, err := mapper.GetKey(transport.NewRequestV2(&req))
			Expect(key).To(Equal(""))
			Expect(err.Error()).Should(Equal("error parsing regexp: invalid UTF-8: `\xbd\xb2`"))
		},
		regexpErrorCasesMultipleFragments...)

	DescribeTable("should return error for empty fragments",
		func(match *MatchPredicate, result *ResultPredicate, request v2.DiscoveryRequest, assert string) {
			protoConfig := KeyerConfiguration{
				Fragments: []*Fragment{
					{
						Rules: []*FragmentRule{
							{
								Match:  match,
								Result: result,
							},
						},
					},
				},
			}
			mapper := New(&protoConfig, stats.NewMockScope(""))
			key, err := mapper.GetKey(transport.NewRequestV2(&request))
			Expect(key).To(Equal(""))
			Expect(err).Should(Equal(fmt.Errorf(assert)))
		},
		emptyFragmentErrorCases...)

	It("TypeUrl should not be empty", func() {
		mapper := New(&KeyerConfiguration{}, stats.NewMockScope(""))
		request := getDiscoveryRequest()
		request.TypeUrl = ""
		key, err := mapper.GetKey(transport.NewRequestV2(&request))
		Expect(key).To(Equal(""))
		Expect(err).Should(Equal(fmt.Errorf("typeURL is empty")))
	})
})

func getAnyMatch(any bool) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_AnyMatch{
			AnyMatch: any,
		},
	}
}

func getRequestTypeMatch(typeurls []string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_RequestTypeMatch_{
			RequestTypeMatch: &aggregationv1.MatchPredicate_RequestTypeMatch{
				Types: typeurls,
			},
		},
	}
}

func getRequestNodeIDExactMatch(exact string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_RequestNodeMatch_{
			RequestNodeMatch: &aggregationv1.MatchPredicate_RequestNodeMatch{
				Type: &aggregationv1.MatchPredicate_RequestNodeMatch_IdMatch{
					IdMatch: &aggregationv1.NodeStringMatch{
						Type: &aggregationv1.NodeStringMatch_ExactMatch{
							ExactMatch: exact,
						},
					},
				},
			},
		},
	}
}

func getRequestNodeIDRegexMatch(regex string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_RequestNodeMatch_{
			RequestNodeMatch: &aggregationv1.MatchPredicate_RequestNodeMatch{
				Type: &aggregationv1.MatchPredicate_RequestNodeMatch_IdMatch{
					IdMatch: &aggregationv1.NodeStringMatch{
						Type: &aggregationv1.NodeStringMatch_RegexMatch{
							RegexMatch: regex,
						},
					},
				},
			},
		},
	}
}

func getRequestNodeClusterExactMatch(exact string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_RequestNodeMatch_{
			RequestNodeMatch: &aggregationv1.MatchPredicate_RequestNodeMatch{
				Type: &aggregationv1.MatchPredicate_RequestNodeMatch_ClusterMatch{
					ClusterMatch: &aggregationv1.NodeStringMatch{
						Type: &aggregationv1.NodeStringMatch_ExactMatch{
							ExactMatch: exact,
						},
					},
				},
			},
		},
	}
}

func getRequestNodeClusterRegexMatch(regex string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_RequestNodeMatch_{
			RequestNodeMatch: &aggregationv1.MatchPredicate_RequestNodeMatch{
				Type: &aggregationv1.MatchPredicate_RequestNodeMatch_ClusterMatch{
					ClusterMatch: &aggregationv1.NodeStringMatch{
						Type: &aggregationv1.NodeStringMatch_RegexMatch{
							RegexMatch: regex,
						},
					},
				},
			},
		},
	}
}

func getRequestNodeLocality(region string, zone string, subZone string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_RequestNodeMatch_{
			RequestNodeMatch: &aggregationv1.MatchPredicate_RequestNodeMatch{
				Type: &aggregationv1.MatchPredicate_RequestNodeMatch_LocalityMatch{
					LocalityMatch: &aggregationv1.NodeLocalityMatch{
						Region:  region,
						Zone:    zone,
						SubZone: subZone,
					},
				},
			},
		},
	}
}

func getRequestNodeAndMatch(predicates []*MatchPredicate) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_AndMatch{
			AndMatch: &aggregationv1.MatchPredicate_MatchSet{
				Rules: predicates,
			},
		},
	}
}

func getRequestNodeOrMatch(predicates []*MatchPredicate) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_OrMatch{
			OrMatch: &aggregationv1.MatchPredicate_MatchSet{
				Rules: predicates,
			},
		},
	}
}

func getRequestTypeNotMatch(typeurls []string) *MatchPredicate {
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_NotMatch{
			NotMatch: getRequestTypeMatch(typeurls),
		},
	}
}

func getRequestNodeLocalityNotMatch() *MatchPredicate {
	return &matchPredicate{
		Type: &aggregationv1.MatchPredicate_NotMatch{
			NotMatch: getRequestNodeLocality("r1", "z2", "sz3"),
		},
	}
}

// TODO: this is the indication we don't need that enum in the protos, i.e. this should
// be a test-only construct.
func getRequestNodeRegexNotMatch(field aggregationv1.NodeFieldType) *MatchPredicate {
	var matchPredicate *aggregationv1.MatchPredicate
	notMatchRegex := "notmatchregex"
	switch field {
	case aggregationv1.NodeFieldType_NODE_ID:
		matchPredicate = getRequestNodeIDRegexMatch(notMatchRegex)
	case aggregationv1.NodeFieldType_NODE_CLUSTER:
		matchPredicate = getRequestNodeClusterRegexMatch(notMatchRegex)
	}
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_NotMatch{
			NotMatch: matchPredicate,
		},
	}
}

// TODO: ditto
func getRequestNodeExactNotMatch(field aggregationv1.NodeFieldType, exact string) *MatchPredicate {
	var matchPredicate *aggregationv1.MatchPredicate
	switch field {
	case aggregationv1.NodeFieldType_NODE_ID:
		matchPredicate = getRequestNodeIDExactMatch(exact)
	case aggregationv1.NodeFieldType_NODE_CLUSTER:
		matchPredicate = getRequestNodeClusterExactMatch(exact)
	}
	return &MatchPredicate{
		Type: &aggregationv1.MatchPredicate_NotMatch{
			NotMatch: matchPredicate,
		},
	}
}

func getResultStringFragment() *ResultPredicate {
	return &ResultPredicate{
		Type: &aggregationv1.ResultPredicate_StringFragment{
			StringFragment: stringFragment,
		},
	}
}

func getRepeatedResultPredicate1() *ResultPredicate {
	return &ResultPredicate{
		Type: &aggregationv1.ResultPredicate_AndResult_{
			AndResult: &aggregationv1.ResultPredicate_AndResult{
				ResultPredicates: []*aggregationv1.ResultPredicate{
					getResultRequestNodeIDFragment(getExactAction()),
					getResultRequestNodeClusterFragment(getExactAction()),
				},
			},
		},
	}
}

// func getRepeatedResultPredicate2() *ResultPredicate {
// 	return &ResultPredicate{
// 		Type: &aggregationv1.ResultPredicate_AndResult_{
// 			AndResult: &aggregationv1.ResultPredicate_AndResult{
// 				ResultPredicates: []*aggregationv1.ResultPredicate{
// 					{
// 						Type: &aggregationv1.ResultPredicate_AndResult_{
// 							AndResult: &aggregationv1.ResultPredicate_AndResult{
// 								ResultPredicates: []*aggregationv1.ResultPredicate{
// 									{
// 										Type: &aggregationv1.ResultPredicate_StringFragment{
// 											StringFragment: "str",
// 										},
// 									},
// 									getResultRequestNodeFragment(nodeRegionField, getExactAction()),
// 									getResultRequestNodeFragment(nodeZoneField, getExactAction()),
// 								},
// 							},
// 						},
// 					},
// 					getResultRequestNodeIDFragment(getRegexAction("ode", "T")),
// 					getResultRequestNodeClusterFragment(getExactAction()),
// 				},
// 			},
// 		},
// 	}
// }

func getRepeatedResultPredicate3() *ResultPredicate {
	return &ResultPredicate{
		Type: &aggregationv1.ResultPredicate_AndResult_{
			AndResult: &aggregationv1.ResultPredicate_AndResult{
				ResultPredicates: []*aggregationv1.ResultPredicate{
					getResourceNameFragment(0, getExactAction()),
					getResourceNameFragment(1, getExactAction()),
				},
			},
		},
	}
}

func getResourceNameFragment(element int, action *aggregationv1.ResultPredicate_ResultAction) *ResultPredicate {
	return &ResultPredicate{
		Type: &aggregationv1.ResultPredicate_ResourceNamesFragment_{
			ResourceNamesFragment: &aggregationv1.ResultPredicate_ResourceNamesFragment{
				Element: int32(element),
				Action:  action,
			},
		},
	}
}

func getResultRequestNodeIDFragment(
	action *aggregationv1.ResultPredicate_ResultAction) *ResultPredicate {
	return &ResultPredicate{
		Type: &aggregationv1.ResultPredicate_RequestNodeFragment_{
			RequestNodeFragment: &aggregationv1.ResultPredicate_RequestNodeFragment{
				Action: &aggregationv1.ResultPredicate_RequestNodeFragment_IdAction{
					IdAction: action,
				},
			},
		},
	}
}

func getResultRequestNodeClusterFragment(
	action *aggregationv1.ResultPredicate_ResultAction) *ResultPredicate {
	return &ResultPredicate{
		Type: &aggregationv1.ResultPredicate_RequestNodeFragment_{
			RequestNodeFragment: &aggregationv1.ResultPredicate_RequestNodeFragment{
				Action: &aggregationv1.ResultPredicate_RequestNodeFragment_ClusterAction{
					ClusterAction: action,
				},
			},
		},
	}
}

func getExactAction() *aggregationv1.ResultPredicate_ResultAction {
	return &aggregationv1.ResultPredicate_ResultAction{
		Action: &aggregationv1.ResultPredicate_ResultAction_Exact{
			Exact: true,
		},
	}
}

func getRegexAction(pattern string, replace string) *aggregationv1.ResultPredicate_ResultAction {
	return &aggregationv1.ResultPredicate_ResultAction{
		Action: &aggregationv1.ResultPredicate_ResultAction_RegexAction_{
			RegexAction: &aggregationv1.ResultPredicate_ResultAction_RegexAction{
				Pattern: pattern,
				Replace: replace,
			},
		},
	}
}

func getDiscoveryRequest() v2.DiscoveryRequest {
	return getDiscoveryRequestWithNode(getNode(nodeid, nodecluster, noderegion, nodezone, nodesubzone))
}

func getDiscoveryRequestWithNode(node *core.Node) v2.DiscoveryRequest {
	return v2.DiscoveryRequest{
		Node:          node,
		VersionInfo:   "version",
		ResourceNames: []string{resource1, resource2},
		TypeUrl:       clusterTypeURL,
	}
}

func getNode(id string, cluster string, region string, zone string, subzone string) *core.Node {
	return &core.Node{
		Id:      id,
		Cluster: cluster,
		Locality: &core.Locality{
			Region:  region,
			Zone:    zone,
			SubZone: subzone,
		},
	}
}

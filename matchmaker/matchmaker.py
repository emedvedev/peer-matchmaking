import sys

from peermap import PeerMap


if __name__ == '__main__':

    peer_map = PeerMap("sample.json")

    pairs = [
        [sys.argv[1], sys.argv[2]]
        # ["8.0.0.1", "9.0.0.1"],
        # ["8.0.0.1", "10.0.0.1"]
    ]

    for pair in pairs:
        nodes = [peer_map.find_subnet(pair[0]), peer_map.find_subnet(pair[1])]
        possible_peering = peer_map.common_ixps(*nodes)
        if possible_peering:
            print "%s (%s) could connect to %s (%s) through %s." % (
                nodes[0]['label'], pair[0], nodes[1]['label'], pair[1],
                ", ".join([ixp['label'] for ixp in possible_peering])
            )
        else:
            print "There is no IXP that could connect %s (%s) to %s (%s)." % (
                nodes[0]['label'], pair[0], nodes[1]['label'], pair[1]
            )

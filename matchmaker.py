import sys

from peermap import PeerMap


if __name__ == '__main__':

    peer_map = PeerMap("sample.json")

    pairs = [
        ["Matt's node", "Ed's subnet"],
        ["Matt's node", "Netflix"]
    ]

    for pair in pairs:
        possible_peering = peer_map.common_ixps(*pair)
        if possible_peering:
            print "%s could connect to %s through %s." % (
                pair[0], pair[1],
                ", ".join([ixp['label'] for ixp in possible_peering])
            )
        else:
            print "There is no IXP that could connect %s to %s." % (
                pair[0], pair[1]
            )

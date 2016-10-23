#!/usr/bin/python

import json

from netaddr import IPNetwork, IPAddress


class PeerMap:
    def __init__(self, path):
        self.path = path
        self._load_map()

    def _load_map(self):
        self.map = None
        self._lookup_table_id = {}
        self._lookup_table_label = {}
        self._lookup_table_subnet = {}
        with open(self.path, 'r') as map_file:
            self.map = json.load(map_file)
        for node in self.map:
            self._lookup_table_id[node['id']] = node
            self._lookup_table_label[node['label']] = node
            if 'subnet' in node:
                self._lookup_table_subnet[node['subnet']] = node

    def lookup(self, node):
        if type(node) is int or type(node) is str and node.isdigit():
            return self._lookup_table_id.get(node)
        elif type(node) is str:
            return self._lookup_table_label.get(node)
        elif node in self.map:
            return node
        else:
            return None

    def find_subnet(self, ip_address):
        for subnet, node in self._lookup_table_subnet.items():
            if IPAddress(ip_address) in IPNetwork(subnet):
                return node
        return None

    def find_peers(self, node, peer_type, depth):
        node = self.lookup(node)
        peers = []
        if node['type'] == peer_type:
            peers.append(node['id'])
        elif 'peering' in node and depth:
            for peer in node['peering']:
                peers += self.find_peers(peer, peer_type, depth-1)
        return set(peers)

    def common_ixps(self, source, destination):
        source_ixps = self.find_peers(source, 'ixp', depth=1)
        destination_ixps = self.find_peers(destination, 'ixp', depth=2)
        return [self.lookup(node) for node in source_ixps & destination_ixps]

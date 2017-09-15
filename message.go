package main

import "net"

// ISISLsp represents the format of a Cisco-IOS-XR-clns-isis-oper
// :isis/instances/instance/levels/level/detailed-lsps/detailed-lsp
// message.
type ISISLsp struct {
	Rows []struct {
		Content struct {
			LspBody       string `json:"lsp_body,omitempty"`
			LspHeaderData struct {
				LocalLspFlag                   bool   `json:"local_lsp_flag,omitempty"`
				LspActiveFlag                  bool   `json:"lsp_active_flag,omitempty"`
				LspAttachedFlag                bool   `json:"lsp_attached_flag,omitempty"`
				LspChecksum                    int64  `json:"lsp_checksum,omitempty"`
				LspHoldtime                    int64  `json:"lsp_holdtime,omitempty"`
				LspID                          string `json:"lsp_id,omitempty"`
				LspLength                      int64  `json:"lsp_length,omitempty"`
				LspLevel                       string `json:"lsp_level,omitempty"`
				LspNonV1AFlag                  int64  `json:"lsp_non_v1_a_flag,omitempty"`
				LspOverloadedFlag              bool   `json:"lsp_overloaded_flag,omitempty"`
				LspParitionRepairSupportedFlag bool   `json:"lsp_parition_repair_supported_flag,omitempty"`
				LspSequenceNumber              int64  `json:"lsp_sequence_number,omitempty"`
			} `json:"lsp_header_data,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InstanceName string `json:"instance_name,omitempty"`
			Level        int64  `json:"level,string,omitempty"`
			LspID        string `json:"lsp_id,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp int64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   int64  `json:"collection_end_time,omitempty"`
		CollectionID        int64  `json:"collection_id,omitempty"`
		CollectionStartTime int64  `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        int64  `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

// ISISInt represents the format of a Cisco-IOS-XR-clns-isis-oper
// :isis/instances/instance/interfaces/interface
// message.
type ISISInt struct {
	Rows []struct {
		Content struct {
			ConfiguredCircuitType string `json:"configured_circuit_type,omitempty"`
			ConfiguredStatus      struct {
				AdjacencyFormStatus bool `json:"adjacency_form_status,omitempty"`
				AdvPrefixStatus     bool `json:"adv_prefix_status,omitempty"`
			} `json:"configured_status,omitempty"`
			InterfaceStatusAndData struct {
				Enabled struct {
					AdjacencyFormStatus struct {
						Disabled struct {
							ReasonCode string `json:"reason_code,omitempty"`
						} `json:"disabled,omitempty"`
						Status string `json:"status,omitempty"`
					} `json:"adjacency_form_status,omitempty"`
					AdvPrefixStatus struct {
						Disabled struct {
							ReasonCode string `json:"reason_code,omitempty"`
						} `json:"disabled,omitempty"`
						Status string `json:"status,omitempty"`
					} `json:"adv_prefix_status,omitempty"`
					BfdData struct {
						Enabled     bool  `json:"enabled,omitempty"`
						Interval    int64 `json:"interval,omitempty"`
						Ipv6Enabled bool  `json:"ipv6_enabled,omitempty"`
						Multiplier  int64 `json:"multiplier,omitempty"`
					} `json:"bfd_data,omitempty"`
					ClnsData struct {
						ClnsStatus struct {
							ClnsDownInfo struct {
								ReasonCode string `json:"reason_code,omitempty"`
							} `json:"clns_down_info,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"clns_status,omitempty"`
						LastLspIDTransmitted     string `json:"last_lsp_id_transmitted,omitempty"`
						LastLspLevelTransmitted  string `json:"last_lsp_level_transmitted,omitempty"`
						LspTransmitRequestedFlag bool   `json:"lsp_transmit_requested_flag,omitempty"`
						LspTransmittedFlag       bool   `json:"lsp_transmitted_flag,omitempty"`
						LspTxmitdB2BLimit        int64  `json:"lsp_txmitd_b2_b_limit,omitempty"`
						LspTxmtB2BMsecs          int64  `json:"lsp_txmt_b2_b_msecs,omitempty"`
						MediaSpecificState       struct {
							ClnsMediaType string `json:"clns_media_type,omitempty"`
							ClnsP2PData   struct {
								P2POverLanData struct {
									MulticastStatus struct {
										Status string `json:"status,omitempty"`
									} `json:"multicast_status,omitempty"`
								} `json:"p2_p_over_lan_data,omitempty"`
							} `json:"clns_p2_p_data,omitempty"`
						} `json:"media_specific_state,omitempty"`
						MtuInfo struct {
							Mtu    int64  `json:"mtu,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"mtu_info,omitempty"`
						SnpaState struct {
							Known struct {
								Snpa string `json:"snpa,omitempty"`
							} `json:"known,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"snpa_state,omitempty"`
						TimeUntilNextLsp int64 `json:"time_until_next_lsp,omitempty"`
					} `json:"clns_data,omitempty"`
					EnabledCircuitType   string `json:"enabled_circuit_type,omitempty"`
					InterfaceMediaType   string `json:"interface_media_type,omitempty"`
					LocalCircuitNumber   int64  `json:"local_circuit_number,omitempty"`
					PerAddressFamilyData []struct {
						AfName   string `json:"af_name,omitempty"`
						AfStatus struct {
							AfData struct {
								ForwardingAddressStatus struct {
									ForwardingAddressData struct {
										ForwardingAddress []struct {
											AfName string `json:"af_name,omitempty"`
											Ipv4   string `json:"ipv4,omitempty"`
											Ipv6   struct {
												Value net.IP `json:"value,omitempty"`
											} `json:"ipv6,omitempty"`
										} `json:"forwarding_address,omitempty"`
										UnnumberedInterfaceName string `json:"unnumbered_interface_name,omitempty"`
									} `json:"forwarding_address_data,omitempty"`
									Status string `json:"status,omitempty"`
								} `json:"forwarding_address_status,omitempty"`
								PrefixStatus struct {
									PrefixData struct {
										IsUnnumbered bool `json:"is_unnumbered,omitempty"`
										Prefix       []struct {
											AfName string `json:"af_name,omitempty"`
											Ipv6   struct {
												Prefix       net.IP `json:"prefix,omitempty"`
												PrefixLength uint8  `json:"prefix_length,omitempty"`
											} `json:"ipv6,omitempty"`
										} `json:"prefix,omitempty"`
									} `json:"prefix_data,omitempty"`
									Status string `json:"status,omitempty"`
								} `json:"prefix_status,omitempty"`
								ProtocolStatus struct {
									Disabled struct {
										ReasonCode string `json:"reason_code,omitempty"`
									} `json:"disabled,omitempty"`
									Status string `json:"status,omitempty"`
								} `json:"protocol_status,omitempty"`
							} `json:"af_data,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"af_status,omitempty"`
					} `json:"per_address_family_data,omitempty"`
					PerTopologyData []struct {
						Status struct {
							Enabled struct {
								AdjacencyFormStatus struct {
									Disabled struct {
										ReasonCode string `json:"reason_code,omitempty"`
									} `json:"disabled,omitempty"`
									Status string `json:"status,omitempty"`
								} `json:"adjacency_form_status,omitempty"`
								AdvPrefixStatus struct {
									Disabled struct {
										ReasonCode string `json:"reason_code,omitempty"`
									} `json:"disabled,omitempty"`
									Status string `json:"status,omitempty"`
								} `json:"adv_prefix_status,omitempty"`
								LdPv6SyncStatus       bool  `json:"ld_pv6_sync_status,omitempty"`
								LdpSyncStatus         bool  `json:"ldp_sync_status,omitempty"`
								Level1LdpSyncEnabled  bool  `json:"level1_ldp_sync_enabled,omitempty"`
								Level1LkgpConfigured  bool  `json:"level1_lkgp_configured,omitempty"`
								Level1Metric          int64 `json:"level1_metric,omitempty"`
								Level1OffsetMetric    int64 `json:"level1_offset_metric,omitempty"`
								Level1PpConfigured    bool  `json:"level1_pp_configured,omitempty"`
								Level1PpMetric        int64 `json:"level1_pp_metric,omitempty"`
								Level1Weight          int64 `json:"level1_weight,omitempty"`
								Level2LdpSyncEnabled  bool  `json:"level2_ldp_sync_enabled,omitempty"`
								Level2LkgpConfigured  bool  `json:"level2_lkgp_configured,omitempty"`
								Level2Metric          int64 `json:"level2_metric,omitempty"`
								Level2OffsetMetric    int64 `json:"level2_offset_metric,omitempty"`
								Level2PpConfigured    bool  `json:"level2_pp_configured,omitempty"`
								Level2PpMetric        int64 `json:"level2_pp_metric,omitempty"`
								Level2Weight          int64 `json:"level2_weight,omitempty"`
								MaxBkpLabelSupported  int64 `json:"max_bkp_label_supported,omitempty"`
								MaxLabelSupported     int64 `json:"max_label_supported,omitempty"`
								MaxSrteLabelSupported int64 `json:"max_srte_label_supported,omitempty"`
							} `json:"enabled,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"status,omitempty"`
						TopologyID struct {
							AfName       string `json:"af_name,omitempty"`
							SafName      string `json:"saf_name,omitempty"`
							TopologyName string `json:"topology_name,omitempty"`
							VrfName      string `json:"vrf_name,omitempty"`
						} `json:"topology_id,omitempty"`
					} `json:"per_topology_data,omitempty"`
					RsiSrlgRegistered bool `json:"rsi_srlg_registered,omitempty"`
				} `json:"enabled,omitempty"`
				Status string `json:"status,omitempty"`
			} `json:"interface_status_and_data,omitempty"`
			IsType    string `json:"is_type,omitempty"`
			MeshGroup int64  `json:"mesh_group,omitempty"`
			NsrIntf   bool   `json:"nsr_intf,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InstanceName  string `json:"instance_name,omitempty"`
			InterfaceName string `json:"interface_name,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp int64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   int64  `json:"collection_end_time,omitempty"`
		CollectionID        int64  `json:"collection_id,omitempty"`
		CollectionStartTime int64  `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        int64  `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

// ISISNbr represents the format of a Cisco-IOS-XR-clns-isis-oper
// :isis/instances/instance/neighbors/neighbor
// message.
type ISISNbr struct {
	Rows []struct {
		Content struct {
			LocalInterface              string `json:"local_interface,omitempty"`
			NeighborActiveAreaAddresses []struct {
				Value string `json:"value,omitempty"`
			} `json:"neighbor_active_area_addresses,omitempty"`
			NeighborCircuitType          string `json:"neighbor_circuit_type,omitempty"`
			NeighborHoldtime             int64  `json:"neighbor_holdtime,omitempty"`
			NeighborIetfNsfCapableFlag   int64  `json:"neighbor_ietf_nsf_capable_flag,omitempty"`
			NeighborMediaType            string `json:"neighbor_media_type,omitempty"`
			NeighborPerAddressFamilyData []struct {
				AfName string `json:"af_name,omitempty"`
				Ipv6   struct {
					InterfaceAddresses []struct {
						Value net.IP `json:"value,omitempty"`
					} `json:"interface_addresses,omitempty"`
					NextHop net.IP `json:"next_hop,omitempty"`
				} `json:"ipv6,omitempty"`
			} `json:"neighbor_per_address_family_data,omitempty"`
			NeighborSnpa            string `json:"neighbor_snpa,omitempty"`
			NeighborState           string `json:"neighbor_state,omitempty"`
			NeighborSystemID        string `json:"neighbor_system_id,omitempty"`
			NeighborUptime          int64  `json:"neighbor_uptime,omitempty"`
			NeighborUptimeValidFlag bool   `json:"neighbor_uptime_valid_flag,omitempty"`
			NsrStandby              bool   `json:"nsr_standby,omitempty"`
			TopologiesSupported     []struct {
				AfName       string `json:"af_name,omitempty"`
				SafName      string `json:"saf_name,omitempty"`
				TopologyName string `json:"topology_name,omitempty"`
				VrfName      string `json:"vrf_name,omitempty"`
			} `json:"topologies_supported,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InstanceName  string `json:"instance_name,omitempty"`
			InterfaceName string `json:"interface_name,omitempty"`
			SystemID      string `json:"system_id,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp int64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   int64  `json:"collection_end_time,omitempty"`
		CollectionID        int64  `json:"collection_id,omitempty"`
		CollectionStartTime int64  `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        int64  `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

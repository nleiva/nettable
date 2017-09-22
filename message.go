package nettable

import "net"

// ISISLSP represents the format of a Cisco-IOS-XR-clns-isis-oper
// :isis/instances/instance/levels/level/detailed-lsps/detailed-lsp
// message.
type ISISLSP struct {
	Rows []struct {
		Content struct {
			LspBody       string `json:"lsp_body,omitempty"`
			LspHeaderData struct {
				LocalLspFlag                   bool   `json:"local_lsp_flag,omitempty"`
				LspActiveFlag                  bool   `json:"lsp_active_flag,omitempty"`
				LspAttachedFlag                bool   `json:"lsp_attached_flag,omitempty"`
				LspChecksum                    uint32 `json:"lsp_checksum,omitempty"`
				LspHoldtime                    uint32 `json:"lsp_holdtime,omitempty"`
				LspID                          string `json:"lsp_id,omitempty"`
				LspLength                      uint32 `json:"lsp_length,omitempty"`
				LspLevel                       string `json:"lsp_level,omitempty"`
				LspNonV1AFlag                  uint32 `json:"lsp_non_v1_a_flag,omitempty"`
				LspOverloadedFlag              bool   `json:"lsp_overloaded_flag,omitempty"`
				LspParitionRepairSupportedFlag bool   `json:"lsp_parition_repair_supported_flag,omitempty"`
				LspSequenceNumber              uint32 `json:"lsp_sequence_number,omitempty"`
			} `json:"lsp_header_data,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InstanceName string `json:"instance_name,omitempty"`
			Level        string `json:"level,omitempty"`
			LspID        string `json:"lsp_id,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp uint64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
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
						Enabled     bool   `json:"enabled,omitempty"`
						Interval    uint32 `json:"interval,omitempty"`
						Ipv6Enabled bool   `json:"ipv6_enabled,omitempty"`
						Multiplier  uint32 `json:"multiplier,omitempty"`
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
						LspTxmitdB2BLimit        uint32 `json:"lsp_txmitd_b2_b_limit,omitempty"`
						LspTxmtB2BMsecs          uint32 `json:"lsp_txmt_b2_b_msecs,omitempty"`
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
							Mtu    uint32 `json:"mtu,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"mtu_info,omitempty"`
						SnpaState struct {
							Known struct {
								Snpa string `json:"snpa,omitempty"`
							} `json:"known,omitempty"`
							Status string `json:"status,omitempty"`
						} `json:"snpa_state,omitempty"`
						TimeUntilNextLsp uint32 `json:"time_until_next_lsp,omitempty"`
					} `json:"clns_data,omitempty"`
					EnabledCircuitType   string `json:"enabled_circuit_type,omitempty"`
					InterfaceMediaType   string `json:"interface_media_type,omitempty"`
					LocalCircuitNumber   uint32 `json:"local_circuit_number,omitempty"`
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
								LdPv6SyncStatus       bool   `json:"ld_pv6_sync_status,omitempty"`
								LdpSyncStatus         bool   `json:"ldp_sync_status,omitempty"`
								Level1LdpSyncEnabled  bool   `json:"level1_ldp_sync_enabled,omitempty"`
								Level1LkgpConfigured  bool   `json:"level1_lkgp_configured,omitempty"`
								Level1Metric          uint32 `json:"level1_metric,omitempty"`
								Level1OffsetMetric    uint32 `json:"level1_offset_metric,omitempty"`
								Level1PpConfigured    bool   `json:"level1_pp_configured,omitempty"`
								Level1PpMetric        uint32 `json:"level1_pp_metric,omitempty"`
								Level1Weight          uint32 `json:"level1_weight,omitempty"`
								Level2LdpSyncEnabled  bool   `json:"level2_ldp_sync_enabled,omitempty"`
								Level2LkgpConfigured  bool   `json:"level2_lkgp_configured,omitempty"`
								Level2Metric          uint32 `json:"level2_metric,omitempty"`
								Level2OffsetMetric    uint32 `json:"level2_offset_metric,omitempty"`
								Level2PpConfigured    bool   `json:"level2_pp_configured,omitempty"`
								Level2PpMetric        uint32 `json:"level2_pp_metric,omitempty"`
								Level2Weight          uint32 `json:"level2_weight,omitempty"`
								MaxBkpLabelSupported  uint32 `json:"max_bkp_label_supported,omitempty"`
								MaxLabelSupported     uint32 `json:"max_label_supported,omitempty"`
								MaxSrteLabelSupported uint32 `json:"max_srte_label_supported,omitempty"`
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
			MeshGroup uint32 `json:"mesh_group,omitempty"`
			NsrIntf   bool   `json:"nsr_intf,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InstanceName  string `json:"instance_name,omitempty"`
			InterfaceName string `json:"interface_name,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp uint64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
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
			NeighborHoldtime             uint32 `json:"neighbor_holdtime,omitempty"`
			NeighborIetfNsfCapableFlag   uint32 `json:"neighbor_ietf_nsf_capable_flag,omitempty"`
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
			NeighborUptime          uint32 `json:"neighbor_uptime,omitempty"`
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
		Timestamp uint64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

// IntCount represents the format of a Cisco-IOS-XR-infra-statsd-oper
// :infra-statistics/interfaces/interface/latest/generic-counters
// message.
type IntCount struct {
	Rows []struct {
		Content struct {
			Applique                       uint32 `json:"applique,omitempty"`
			AvailabilityFlag               uint32 `json:"availability_flag,omitempty"`
			BroadcastPacketsReceived       uint64 `json:"broadcast_packets_received,omitempty"`
			BroadcastPacketsSent           uint64 `json:"broadcast_packets_sent,omitempty"`
			BytesReceived                  uint64 `json:"bytes_received,omitempty"`
			BytesSent                      uint64 `json:"bytes_sent,omitempty"`
			CarrierTransitions             uint32 `json:"carrier_transitions,omitempty"`
			CrcErrors                      uint32 `json:"crc_errors,omitempty"`
			FramingErrorsReceived          uint32 `json:"framing_errors_received,omitempty"`
			GiantPacketsReceived           uint32 `json:"giant_packets_received,omitempty"`
			InputAborts                    uint32 `json:"input_aborts,omitempty"`
			InputDrops                     uint32 `json:"input_drops,omitempty"`
			InputErrors                    uint32 `json:"input_errors,omitempty"`
			InputIgnoredPackets            uint32 `json:"input_ignored_packets,omitempty"`
			InputOverruns                  uint32 `json:"input_overruns,omitempty"`
			InputQueueDrops                uint32 `json:"input_queue_drops,omitempty"`
			LastDataTime                   uint32 `json:"last_data_time,omitempty"`
			LastDiscontinuityTime          uint32 `json:"last_discontinuity_time,omitempty"`
			MulticastPacketsReceived       uint64 `json:"multicast_packets_received,omitempty"`
			MulticastPacketsSent           uint64 `json:"multicast_packets_sent,omitempty"`
			OutputBufferFailures           uint32 `json:"output_buffer_failures,omitempty"`
			OutputBuffersSwappedOut        uint32 `json:"output_buffers_swapped_out,omitempty"`
			OutputDrops                    uint32 `json:"output_drops,omitempty"`
			OutputErrors                   uint32 `json:"output_errors,omitempty"`
			OutputQueueDrops               uint32 `json:"output_queue_drops,omitempty"`
			OutputUnderruns                uint32 `json:"output_underruns,omitempty"`
			PacketsReceived                uint64 `json:"packets_received,omitempty"`
			PacketsSent                    uint64 `json:"packets_sent,omitempty"`
			ParityPacketsReceived          uint32 `json:"parity_packets_received,omitempty"`
			Resets                         uint32 `json:"resets,omitempty"`
			RuntPacketsReceived            uint32 `json:"runt_packets_received,omitempty"`
			SecondsSinceLastClearCounters  uint32 `json:"seconds_since_last_clear_counters,omitempty"`
			SecondsSincePacketReceived     uint32 `json:"seconds_since_packet_received,omitempty"`
			SecondsSincePacketSent         uint32 `json:"seconds_since_packet_sent,omitempty"`
			ThrottledPacketsReceived       uint32 `json:"throttled_packets_received,omitempty"`
			UnknownProtocolPacketsReceived uint32 `json:"unknown_protocol_packets_received,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InterfaceName string `json:"interface_name,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp uint64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

// IntRate represents the format of a Cisco-IOS-XR-infra-statsd-oper
// :infra-statistics/interfaces/interface/data-rate
// message.
type IntRate struct {
	Rows []struct {
		Content struct {
			Bandwidth            uint32 `json:"bandwidth,omitempty"`
			InputDataRate        uint64 `json:"input_data_rate,omitempty"`
			InputLoad            uint32 `json:"input_load,omitempty"`
			InputPacketRate      uint64 `json:"input_packet_rate,omitempty"`
			LoadInterval         uint32 `json:"load_interval,omitempty"`
			OutputDataRate       uint64 `json:"output_data_rate,omitempty"`
			OutputLoad           uint32 `json:"output_load,omitempty"`
			OutputPacketRate     uint64 `json:"output_packet_rate,omitempty"`
			PeakInputDataRate    uint64 `json:"peak_input_data_rate,omitempty"`
			PeakInputPacketRate  uint64 `json:"peak_input_packet_rate,omitempty"`
			PeakOutputDataRate   uint64 `json:"peak_output_data_rate,omitempty"`
			PeakOutputPacketRate uint64 `json:"peak_output_packet_rate,omitempty"`
			Reliability          uint32 `json:"reliability,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			InterfaceName string `json:"interface_name,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp uint64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

// RIBIPv6 represents the format of a Cisco-IOS-XR-ip-rib-ipv6-oper
// :ipv6-rib/vrfs/vrf/afs/af/safs/saf/ip-rib-route-table-names/ip-rib-route-table-name/routes/route
// message.
type RIBIPv6 struct {
	Rows []struct {
		Content struct {
			Active             bool   `json:"active,omitempty"`
			AttributeIDentity  uint32 `json:"attribute_identity,omitempty"`
			ClientID           uint32 `json:"client_id,omitempty"`
			Distance           uint32 `json:"distance,omitempty"`
			Diversion          bool   `json:"diversion,omitempty"`
			DiversionDistance  uint32 `json:"diversion_distance,omitempty"`
			DiversionProtoName string `json:"diversion_proto_name,omitempty"`
			ExtendedFlags      uint64 `json:"extended_flags,omitempty"`
			Flags              uint32 `json:"flags,omitempty"`
			FlowTag            uint32 `json:"flow_tag,omitempty"`
			FwdClass           uint32 `json:"fwd_class,omitempty"`
			Instance           string `json:"instance,omitempty"`
			Metric             uint32 `json:"metric,omitempty"`
			PathsCount         uint32 `json:"paths_count,omitempty"`
			PicCount           uint32 `json:"pic_count,omitempty"`
			Prefix             net.IP `json:"prefix,omitempty"`
			PrefixLength       uint32 `json:"prefix_length,omitempty"`
			Priority           uint32 `json:"priority,omitempty"`
			ProtocolID         uint32 `json:"protocol_id,omitempty"`
			ProtocolName       string `json:"protocol_name,omitempty"`
			QosGroup           uint32 `json:"qos_group,omitempty"`
			RouteAge           uint32 `json:"route_age,omitempty"`
			RouteLabel         uint32 `json:"route_label,omitempty"`
			RouteModifyTime    uint64 `json:"route_modify_time,omitempty"`
			RoutePath          struct {
				Ipv6RibEdmPath []struct {
					Address                     net.IP `json:"address,omitempty"`
					BackupPathid                uint32 `json:"backup_pathid,omitempty"`
					BindingLabel                uint32 `json:"binding_label,omitempty"`
					Flags                       uint32 `json:"flags,omitempty"`
					Flags64                     uint64 `json:"flags64,omitempty"`
					HasLabelstk                 bool   `json:"has_labelstk,omitempty"`
					InformationSource           net.IP `json:"information_source,omitempty"`
					InterfaceName               net.IP `json:"interface_name,omitempty"`
					LoadMetric                  uint32 `json:"load_metric,omitempty"`
					Looped                      bool   `json:"looped,omitempty"`
					Metric                      uint32 `json:"metric,omitempty"`
					MplsFeid                    uint64 `json:"mpls_feid,omitempty"`
					MvpnPresent                 bool   `json:"mvpn_present,omitempty"`
					NextHopAfi                  int64  `json:"next_hop_afi,omitempty"`
					NextHopID                   uint32 `json:"next_hop_id,omitempty"`
					NextHopIDRefcount           uint32 `json:"next_hop_id_refcount,omitempty"`
					NextHopSafi                 uint32 `json:"next_hop_safi,omitempty"`
					NextHopTableID              uint32 `json:"next_hop_table_id,omitempty"`
					NextHopTableName            string `json:"next_hop_table_name,omitempty"`
					NextHopVrfName              string `json:"next_hop_vrf_name,omitempty"`
					NhidFeid                    uint64 `json:"nhid_feid,omitempty"`
					NumLabels                   uint32 `json:"num_labels,omitempty"`
					NumberOfExtendedCommunities uint32 `json:"number_of_extended_communities,omitempty"`
					OspfAreaID                  string `json:"ospf_area_id,omitempty"`
					Pathid                      uint32 `json:"pathid,omitempty"`
					PathrtPresent               bool   `json:"pathrt_present,omitempty"`
					PrivateFlags                uint32 `json:"private_flags,omitempty"`
					RefCntOfBackup              uint32 `json:"ref_cnt_of_backup,omitempty"`
					RouteLabel                  uint32 `json:"route_label,omitempty"`
					SegmentedNexthopPresent     bool   `json:"segmented_nexthop_present,omitempty"`
					SourceasrtPresent           bool   `json:"sourceasrt_present,omitempty"`
					SourcerdPresent             bool   `json:"sourcerd_present,omitempty"`
					TunnelID                    uint32 `json:"tunnel_id,omitempty"`
					V6Nexthop                   net.IP `json:"v6_nexthop,omitempty"`
					VrfimportrtPresent          bool   `json:"vrfimportrt_present,omitempty"`
				} `json:"ipv6_rib_edm_path,omitempty"`
			} `json:"route_path,omitempty"`
			RoutePrecedence uint32 `json:"route_precedence,omitempty"`
			RouteType       uint32 `json:"route_type,omitempty"`
			RouteVersion    uint32 `json:"route_version,omitempty"`
			SvdType         uint32 `json:"svd_type,omitempty"`
			Tag             uint32 `json:"tag,omitempty"`
			TblVersion      uint32 `json:"tbl_version,omitempty"`
			TrafficIndex    uint32 `json:"traffic_index,omitempty"`
			Version         uint32 `json:"version,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			Address        string `json:"address,omitempty"`
			AfName         string `json:"af_name,omitempty"`
			InterfaceName  string `json:"interface_name,omitempty"`
			NextHopAddress string `json:"next_hop_address,omitempty"`
			PrefixLength   uint32 `json:"prefix_length,omitempty"`
			RouteTableName string `json:"route_table_name,omitempty"`
			SafName        string `json:"saf_name,omitempty"`
			VrfName        string `json:"vrf_name,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp int64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

// BGPNbr represents the format of a Cisco-IOS-XR-ipv4-bgp-oper
// :bgp/instances/instance/instance-active/default-vrf/afs/af/neighbor-af-table/neighbor
// message.
type BGPNbr struct {
	Rows []struct {
		Content struct {
			ActiveBmpServers                uint32 `json:"active_bmp_servers,omitempty"`
			ActualLastWriteBytes            uint32 `json:"actual_last_write_bytes,omitempty"`
			ActualLastWriteResetBytes       uint32 `json:"actual_last_write_reset_bytes,omitempty"`
			ActualSecondLastWriteBytes      uint32 `json:"actual_second_last_write_bytes,omitempty"`
			ActualSecondLastWriteResetBytes uint32 `json:"actual_second_last_write_reset_bytes,omitempty"`
			AfData                          []struct {
				Value struct {
					AcceptOwnEnabled                            bool   `json:"accept_own_enabled,omitempty"`
					AckedVersion                                uint32 `json:"acked_version,omitempty"`
					AddressFamilyLongLivedTime                  uint32 `json:"address_family_long_lived_time,omitempty"`
					AdvertiseAfi                                bool   `json:"advertise_afi,omitempty"`
					AdvertiseAfiDefVrfImpDisable                bool   `json:"advertise_afi_def_vrf_imp_disable,omitempty"`
					AdvertiseAfiDisable                         bool   `json:"advertise_afi_disable,omitempty"`
					AdvertiseAfiEoRReady                        bool   `json:"advertise_afi_eo_r_ready,omitempty"`
					AdvertiseAfiReorg                           bool   `json:"advertise_afi_reorg,omitempty"`
					AdvertiseAfiVrfReImpDisable                 bool   `json:"advertise_afi_vrf_re_imp_disable,omitempty"`
					AdvertiseEvpNv4AfiDefVrfImpDisable          bool   `json:"advertise_evp_nv4_afi_def_vrf_imp_disable,omitempty"`
					AdvertiseEvpNv4AfiVrfReImpDisable           bool   `json:"advertise_evp_nv4_afi_vrf_re_imp_disable,omitempty"`
					AdvertiseEvpNv6AfiDefVrfImpDisable          bool   `json:"advertise_evp_nv6_afi_def_vrf_imp_disable,omitempty"`
					AdvertiseEvpNv6AfiVrfReImpDisable           bool   `json:"advertise_evp_nv6_afi_vrf_re_imp_disable,omitempty"`
					AdvertiseLocalLabeledRouteUnicast           bool   `json:"advertise_local_labeled_route_unicast,omitempty"`
					AdvertiseRtType                             uint32 `json:"advertise_rt_type,omitempty"`
					AdvertiseV4Flags                            uint32 `json:"advertise_v4_flags,omitempty"`
					AdvertiseV6Flags                            uint32 `json:"advertise_v6_flags,omitempty"`
					AfName                                      string `json:"af_name,omitempty"`
					AfrpkiAllowInvalid                          bool   `json:"afrpki_allow_invalid,omitempty"`
					AfrpkiDisable                               bool   `json:"afrpki_disable,omitempty"`
					AfrpkiSignalIbgp                            bool   `json:"afrpki_signal_ibgp,omitempty"`
					AfrpkiUseValidity                           bool   `json:"afrpki_use_validity,omitempty"`
					AllowAsInCount                              uint32 `json:"allow_as_in_count,omitempty"`
					AlwaysUseNextHopLocal                       bool   `json:"always_use_next_hop_local,omitempty"`
					EncapsulationType                           uint32 `json:"encapsulation_type,omitempty"`
					EoRReceivedInReadOnly                       bool   `json:"eo_r_received_in_read_only,omitempty"`
					FilterGroupIndex                            uint32 `json:"filter_group_index,omitempty"`
					FlowspecRedirectValidationDIsable           bool   `json:"flowspec_redirect_validation_d_isable,omitempty"`
					FlowspecValidationDIsable                   bool   `json:"flowspec_validation_d_isable,omitempty"`
					ImportReoriginate                           bool   `json:"import_reoriginate,omitempty"`
					ImportReoriginateStitching                  bool   `json:"import_reoriginate_stitching,omitempty"`
					ImportStitching                             bool   `json:"import_stitching,omitempty"`
					IsAddPathReceiveCapabilityAdvertised        bool   `json:"is_add_path_receive_capability_advertised,omitempty"`
					IsAddPathReceiveCapabilityReceived          bool   `json:"is_add_path_receive_capability_received,omitempty"`
					IsAddPathSendCapabilityAdvertised           bool   `json:"is_add_path_send_capability_advertised,omitempty"`
					IsAddPathSendCapabilityReceived             bool   `json:"is_add_path_send_capability_received,omitempty"`
					IsAddpathReceiveOperational                 bool   `json:"is_addpath_receive_operational,omitempty"`
					IsAddpathSendOperational                    bool   `json:"is_addpath_send_operational,omitempty"`
					IsAdvertisePermanentNetwork                 bool   `json:"is_advertise_permanent_network,omitempty"`
					IsAdvertisedGracefulRestart                 bool   `json:"is_advertised_graceful_restart,omitempty"`
					IsAdvertisedOrfReceive                      bool   `json:"is_advertised_orf_receive,omitempty"`
					IsAdvertisedOrfSend                         bool   `json:"is_advertised_orf_send,omitempty"`
					IsAigpSet                                   bool   `json:"is_aigp_set,omitempty"`
					IsAllowAsInSet                              bool   `json:"is_allow_as_in_set,omitempty"`
					IsAsOverrideSet                             bool   `json:"is_as_override_set,omitempty"`
					IsDefaultOriginateSent                      bool   `json:"is_default_originate_sent,omitempty"`
					IsExtNhEncodingCapabilityReceived           bool   `json:"is_ext_nh_encoding_capability_received,omitempty"`
					IsExtNhEncodingCapabilitySent               bool   `json:"is_ext_nh_encoding_capability_sent,omitempty"`
					IsGracefulRestartStateFlag                  bool   `json:"is_graceful_restart_state_flag,omitempty"`
					IsLegacyPeRt                                bool   `json:"is_legacy_pe_rt,omitempty"`
					IsNeighborAfCapable                         bool   `json:"is_neighbor_af_capable,omitempty"`
					IsNeighborEbgpWithoutInboundPolicy          bool   `json:"is_neighbor_ebgp_without_inbound_policy,omitempty"`
					IsNeighborEbgpWithoutOutboundPolicy         bool   `json:"is_neighbor_ebgp_without_outbound_policy,omitempty"`
					IsNeighborRouteReflectorClient              bool   `json:"is_neighbor_route_reflector_client,omitempty"`
					IsOrfSendScheduled                          bool   `json:"is_orf_send_scheduled,omitempty"`
					IsOrfSent                                   bool   `json:"is_orf_sent,omitempty"`
					IsOrrRootAddressConfigured                  bool   `json:"is_orr_root_address_configured,omitempty"`
					IsPeerOrfCapable                            bool   `json:"is_peer_orf_capable,omitempty"`
					IsPrefixOrfPresent                          bool   `json:"is_prefix_orf_present,omitempty"`
					IsReceivedGracefulRestartCapable            bool   `json:"is_received_graceful_restart_capable,omitempty"`
					IsReceivedOrfReceiveCapable                 bool   `json:"is_received_orf_receive_capable,omitempty"`
					IsReceivedOrfSendCapable                    bool   `json:"is_received_orf_send_capable,omitempty"`
					IsRtPresent                                 bool   `json:"is_rt_present,omitempty"`
					IsRtPresentStandby                          bool   `json:"is_rt_present_standby,omitempty"`
					IsSendMcastAttr                             bool   `json:"is_send_mcast_attr,omitempty"`
					IsSoftReconfigurationInboundAllowed         bool   `json:"is_soft_reconfiguration_inbound_allowed,omitempty"`
					IsUpdateDeferred                            bool   `json:"is_update_deferred,omitempty"`
					IsUpdateLeaving                             bool   `json:"is_update_leaving,omitempty"`
					IsUpdateThrottled                           bool   `json:"is_update_throttled,omitempty"`
					IsUseSoftReconfigurationAlwaysOn            bool   `json:"is_use_soft_reconfiguration_always_on,omitempty"`
					LocalRestartTime                            uint32 `json:"local_restart_time,omitempty"`
					LongLivedGracefulRestartCapabilityReceived  bool   `json:"long_lived_graceful_restart_capability_received,omitempty"`
					LongLivedGracefulRestartStaleTimeAccept     uint32 `json:"long_lived_graceful_restart_stale_time_accept,omitempty"`
					LongLivedGracefulRestartStaleTimeConfigured bool   `json:"long_lived_graceful_restart_stale_time_configured,omitempty"`
					LongLivedGracefulRestartStaleTimeReceived   uint32 `json:"long_lived_graceful_restart_stale_time_received,omitempty"`
					LongLivedGracefulRestartStaleTimeSent       uint32 `json:"long_lived_graceful_restart_stale_time_sent,omitempty"`
					MaxPrefixDiscardExtraPaths                  bool   `json:"max_prefix_discard_extra_paths,omitempty"`
					MaxPrefixExceedDiscardPaths                 bool   `json:"max_prefix_exceed_discard_paths,omitempty"`
					MaxPrefixLimit                              uint32 `json:"max_prefix_limit,omitempty"`
					MaxPrefixRestartTime                        uint32 `json:"max_prefix_restart_time,omitempty"`
					MaxPrefixThresholdPercent                   uint32 `json:"max_prefix_threshold_percent,omitempty"`
					NeighborAfPerformanceStatistics             struct {
						ProcessedMessages            uint32 `json:"processed_messages,omitempty"`
						SentMessages                 uint32 `json:"sent_messages,omitempty"`
						SplitHorizonUpdateBlocked    uint32 `json:"split_horizon_update_blocked,omitempty"`
						SplitHorizonUpdateTransmit   uint32 `json:"split_horizon_update_transmit,omitempty"`
						SplitHorizonWithdrawBlocked  uint32 `json:"split_horizon_withdraw_blocked,omitempty"`
						SplitHorizonWithdrawTransmit uint32 `json:"split_horizon_withdraw_transmit,omitempty"`
						SubGroupPendingMessageCount  uint32 `json:"sub_group_pending_message_count,omitempty"`
					} `json:"neighbor_af_performance_statistics,omitempty"`
					NeighborDefaultOriginate                      bool   `json:"neighbor_default_originate,omitempty"`
					NeighborLongLivedGracefulRestartCapable       bool   `json:"neighbor_long_lived_graceful_restart_capable,omitempty"`
					NeighborLongLivedGracefulRestartTimeRemaining uint32 `json:"neighbor_long_lived_graceful_restart_time_remaining,omitempty"`
					NeighborPreservedForwardingState              bool   `json:"neighbor_preserved_forwarding_state,omitempty"`
					NeighborPreservedLongLivedForwardingState     bool   `json:"neighbor_preserved_long_lived_forwarding_state,omitempty"`
					NeighborVersion                               uint32 `json:"neighbor_version,omitempty"`
					NumberOfBestExternalpaths                     uint32 `json:"number_of_best_externalpaths,omitempty"`
					NumberOfBestpaths                             uint32 `json:"number_of_bestpaths,omitempty"`
					OrfEntriesReceived                            uint32 `json:"orf_entries_received,omitempty"`
					OrrGroupIndex                                 uint32 `json:"orr_group_index,omitempty"`
					OrrGroupName                                  string `json:"orr_group_name,omitempty"`
					OutstandingVersion                            uint32 `json:"outstanding_version,omitempty"`
					OutstandingVersionMax                         uint32 `json:"outstanding_version_max,omitempty"`
					PrefixesAccepted                              uint32 `json:"prefixes_accepted,omitempty"`
					PrefixesAdvertised                            uint32 `json:"prefixes_advertised,omitempty"`
					PrefixesBeAdvertised                          uint32 `json:"prefixes_be_advertised,omitempty"`
					PrefixesDenied                                uint32 `json:"prefixes_denied,omitempty"`
					PrefixesDeniedNoPolicy                        uint32 `json:"prefixes_denied_no_policy,omitempty"`
					PrefixesDeniedNonCumulative                   uint32 `json:"prefixes_denied_non_cumulative,omitempty"`
					PrefixesDeniedOrfPolicy                       uint32 `json:"prefixes_denied_orf_policy,omitempty"`
					PrefixesDeniedPolicy                          uint32 `json:"prefixes_denied_policy,omitempty"`
					PrefixesDeniedRtPermit                        uint32 `json:"prefixes_denied_rt_permit,omitempty"`
					PrefixesSuppressed                            uint32 `json:"prefixes_suppressed,omitempty"`
					PrefixesSynced                                uint32 `json:"prefixes_synced,omitempty"`
					PrefixesWithdrawn                             uint32 `json:"prefixes_withdrawn,omitempty"`
					PrefixesWithdrawnNotFound                     uint32 `json:"prefixes_withdrawn_not_found,omitempty"`
					RefreshAckedVersion                           uint32 `json:"refresh_acked_version,omitempty"`
					RefreshTargetVersion                          uint32 `json:"refresh_target_version,omitempty"`
					RefreshVersion                                uint32 `json:"refresh_version,omitempty"`
					RemovePrivateAsEntireAspathFromInboundUpdates bool   `json:"remove_private_as_entire_aspath_from_inbound_updates,omitempty"`
					RemovePrivateAsEntireAspathFromUpdates        bool   `json:"remove_private_as_entire_aspath_from_updates,omitempty"`
					RemovePrivateAsFromInboundUpdates             bool   `json:"remove_private_as_from_inbound_updates,omitempty"`
					RemovePrivateAsFromUpdates                    bool   `json:"remove_private_as_from_updates,omitempty"`
					RestartTime                                   uint32 `json:"restart_time,omitempty"`
					RibPurgeTimeoutValue                          uint32 `json:"rib_purge_timeout_value,omitempty"`
					RoutePolicyDefaultOriginate                   string `json:"route_policy_default_originate,omitempty"`
					RoutePolicyIn                                 string `json:"route_policy_in,omitempty"`
					RoutePolicyOut                                string `json:"route_policy_out,omitempty"`
					RoutePolicyPrefixOrf                          string `json:"route_policy_prefix_orf,omitempty"`
					RouteRefreshesReceived                        uint32 `json:"route_refreshes_received,omitempty"`
					RouteRefreshesSent                            uint32 `json:"route_refreshes_sent,omitempty"`
					SelectiveMultipathEligible                    bool   `json:"selective_multipath_eligible,omitempty"`
					SentCommunityToNeighbor                       bool   `json:"sent_community_to_neighbor,omitempty"`
					SentExtendedCommunityToNeighbor               bool   `json:"sent_extended_community_to_neighbor,omitempty"`
					SentGshutCommunityToNeighbor                  bool   `json:"sent_gshut_community_to_neighbor,omitempty"`
					StalePathTimeout                              uint32 `json:"stale_path_timeout,omitempty"`
					SyncedAckedVersion                            uint32 `json:"synced_acked_version,omitempty"`
					UpdateGroupNumber                             uint32 `json:"update_group_number,omitempty"`
					UseMaxPrefixWarningOnly                       bool   `json:"use_max_prefix_warning_only,omitempty"`
					VpnUpdateGenEnabled                           bool   `json:"vpn_update_gen_enabled,omitempty"`
					VpnUpdateGenTriggerEnabled                    bool   `json:"vpn_update_gen_trigger_enabled,omitempty"`
					Weight                                        uint32 `json:"weight,omitempty"`
				} `json:"value,omitempty"`
			} `json:"af_data,omitempty"`
			AttemptedLastWriteBytes            uint32 `json:"attempted_last_write_bytes,omitempty"`
			AttemptedLastWriteResetBytes       uint32 `json:"attempted_last_write_reset_bytes,omitempty"`
			AttemptedSecondLastWriteBytes      uint32 `json:"attempted_second_last_write_bytes,omitempty"`
			AttemptedSecondLastWriteResetBytes uint32 `json:"attempted_second_last_write_reset_bytes,omitempty"`
			BfdMinintervalval                  uint32 `json:"bfd_minintervalval,omitempty"`
			BfdMultiplierval                   uint32 `json:"bfd_multiplierval,omitempty"`
			BfdSessionCreatedState             string `json:"bfd_session_created_state,omitempty"`
			BfdSessionEnableMode               string `json:"bfd_session_enable_mode,omitempty"`
			BfdSessionState                    string `json:"bfd_session_state,omitempty"`
			BfdStateTs                         uint64 `json:"bfd_state_ts,omitempty"`
			BytesRead                          uint32 `json:"bytes_read,omitempty"`
			BytesWritten                       uint32 `json:"bytes_written,omitempty"`
			ConfiguredHoldTime                 uint32 `json:"configured_hold_time,omitempty"`
			ConfiguredKeepalive                uint32 `json:"configured_keepalive,omitempty"`
			ConfiguredMinAccHoldTime           uint32 `json:"configured_min_acc_hold_time,omitempty"`
			ConnectRetryInterval               uint32 `json:"connect_retry_interval,omitempty"`
			ConnectionAdminStatus              uint32 `json:"connection_admin_status,omitempty"`
			ConnectionDownCount                uint32 `json:"connection_down_count,omitempty"`
			ConnectionEstablishedTime          uint32 `json:"connection_established_time,omitempty"`
			ConnectionLocalAddress             struct {
				Afi                 string `json:"afi,omitempty"`
				Ipv4Address         string `json:"ipv4_address,omitempty"`
				Ipv4LabelAddress    string `json:"ipv4_label_address,omitempty"`
				Ipv4McastAddress    string `json:"ipv4_mcast_address,omitempty"`
				Ipv4VpnAddress      string `json:"ipv4_vpn_address,omitempty"`
				Ipv4VpnaMcastddress string `json:"ipv4_vpna_mcastddress,omitempty"`
				Ipv6Address         struct {
					Value net.IP `json:"value,omitempty"`
				} `json:"ipv6_address,omitempty"`
			} `json:"connection_local_address,omitempty"`
			ConnectionLocalPort     uint32 `json:"connection_local_port,omitempty"`
			ConnectionRemoteAddress struct {
				Afi                 string `json:"afi,omitempty"`
				Ipv4Address         string `json:"ipv4_address,omitempty"`
				Ipv4LabelAddress    string `json:"ipv4_label_address,omitempty"`
				Ipv4McastAddress    string `json:"ipv4_mcast_address,omitempty"`
				Ipv4VpnAddress      string `json:"ipv4_vpn_address,omitempty"`
				Ipv4VpnaMcastddress string `json:"ipv4_vpna_mcastddress,omitempty"`
				Ipv6Address         struct {
					Value net.IP `json:"value,omitempty"`
				} `json:"ipv6_address,omitempty"`
			} `json:"connection_remote_address,omitempty"`
			ConnectionRemotePort                    uint32 `json:"connection_remote_port,omitempty"`
			ConnectionState                         string `json:"connection_state,omitempty"`
			ConnectionUpCount                       uint32 `json:"connection_up_count,omitempty"`
			CountLastWrite                          uint32 `json:"count_last_write,omitempty"`
			Description                             string `json:"description,omitempty"`
			DiscardAs4Path                          uint32 `json:"discard_as4_path,omitempty"`
			DiscardDataBytes                        uint32 `json:"discard_data_bytes,omitempty"`
			DmzLinkBandwidth                        uint32 `json:"dmz_link_bandwidth,omitempty"`
			EbgpRecvDmz                             bool   `json:"ebgp_recv_dmz,omitempty"`
			EbgpSendDmzMode                         string `json:"ebgp_send_dmz_mode,omitempty"`
			EbgpTimeToLive                          uint32 `json:"ebgp_time_to_live,omitempty"`
			EgressPeerEngineeringEnabled            bool   `json:"egress_peer_engineering_enabled,omitempty"`
			ErrorNotifiesReceived                   uint32 `json:"error_notifies_received,omitempty"`
			ErrorNotifiesSent                       uint32 `json:"error_notifies_sent,omitempty"`
			FpbsnOffset                             uint32 `json:"fpbsn_offset,omitempty"`
			FssnOffset                              uint32 `json:"fssn_offset,omitempty"`
			GrRestartTime                           uint32 `json:"gr_restart_time,omitempty"`
			GrStalePathTime                         uint32 `json:"gr_stale_path_time,omitempty"`
			GracefulRestartEnabledNbr               bool   `json:"graceful_restart_enabled_nbr,omitempty"`
			HasInternalLink                         bool   `json:"has_internal_link,omitempty"`
			HoldTime                                uint32 `json:"hold_time,omitempty"`
			IgnoreConnected                         bool   `json:"ignore_connected,omitempty"`
			InternalVpnClient                       bool   `json:"internal_vpn_client,omitempty"`
			IoArmed                                 bool   `json:"io_armed,omitempty"`
			Is4ByteAsCapabilityReceived             bool   `json:"is4_byte_as_capability_received,omitempty"`
			Is4ByteAsCapabilitySent                 bool   `json:"is4_byte_as_capability_sent,omitempty"`
			IsAdministrativelyShutDown              bool   `json:"is_administratively_shut_down,omitempty"`
			IsCapabilityNegotiationPerformed        bool   `json:"is_capability_negotiation_performed,omitempty"`
			IsCapabilityNegotiationSuppressed       bool   `json:"is_capability_negotiation_suppressed,omitempty"`
			IsEbgpMultihopBgpmplsForwardingDisabled bool   `json:"is_ebgp_multihop_bgpmpls_forwarding_disabled,omitempty"`
			IsEbgpPeerAsLeague                      bool   `json:"is_ebgp_peer_as_league,omitempty"`
			IsEbgpPeerCommonAdmin                   bool   `json:"is_ebgp_peer_common_admin,omitempty"`
			IsExternalNeighborNotDirectlyConnected  bool   `json:"is_external_neighbor_not_directly_connected,omitempty"`
			IsGrAware                               bool   `json:"is_gr_aware,omitempty"`
			IsLocalAddressConfigured                bool   `json:"is_local_address_configured,omitempty"`
			IsNeighborMaxPrefixShutdown             bool   `json:"is_neighbor_max_prefix_shutdown,omitempty"`
			IsOutOfMemoryForcedUp                   bool   `json:"is_out_of_memory_forced_up,omitempty"`
			IsOutOfMemoryShutdown                   bool   `json:"is_out_of_memory_shutdown,omitempty"`
			IsPassiveClose                          bool   `json:"is_passive_close,omitempty"`
			IsReadDisabled                          bool   `json:"is_read_disabled,omitempty"`
			IsRouteRefreshCapabilityReceived        bool   `json:"is_route_refresh_capability_received,omitempty"`
			IsRouteRefreshOldCapabilityReceived     bool   `json:"is_route_refresh_old_capability_received,omitempty"`
			KeepAliveTime                           uint32 `json:"keep_alive_time,omitempty"`
			LastAckdSeqNo                           uint32 `json:"last_ackd_seq_no,omitempty"`
			LastKAerrorReset                        uint32 `json:"last_k_aerror_reset,omitempty"`
			LastKAexpiryReset                       uint32 `json:"last_k_aexpiry_reset,omitempty"`
			LastKAnotsentReset                      uint32 `json:"last_k_anotsent_reset,omitempty"`
			LastKAstartReset                        uint32 `json:"last_k_astart_reset,omitempty"`
			LastNotifyErrorCode                     uint32 `json:"last_notify_error_code,omitempty"`
			LastNotifyErrorSubcode                  uint32 `json:"last_notify_error_subcode,omitempty"`
			LastWriteEvent                          uint32 `json:"last_write_event,omitempty"`
			LocalAs                                 uint32 `json:"local_as,omitempty"`
			LocalAsDualAs                           bool   `json:"local_as_dual_as,omitempty"`
			LocalAsDualAsModeNative                 bool   `json:"local_as_dual_as_mode_native,omitempty"`
			LocalAsNoPrepend                        bool   `json:"local_as_no_prepend,omitempty"`
			LocalAsReplaceAs                        bool   `json:"local_as_replace_as,omitempty"`
			MessagesQueuedIn                        uint32 `json:"messages_queued_in,omitempty"`
			MessagesQueuedOut                       uint32 `json:"messages_queued_out,omitempty"`
			MessagesReceived                        uint32 `json:"messages_received,omitempty"`
			MessagesSent                            uint32 `json:"messages_sent,omitempty"`
			MinAdvertiseInterval                    uint32 `json:"min_advertise_interval,omitempty"`
			MinAdvertiseIntervalMsecs               uint32 `json:"min_advertise_interval_msecs,omitempty"`
			MinOriginationInterval                  uint32 `json:"min_origination_interval,omitempty"`
			MsgLogIn                                uint32 `json:"msg_log_in,omitempty"`
			MsgLogOut                               uint32 `json:"msg_log_out,omitempty"`
			MultiProtocolCapabilityReceived         bool   `json:"multi_protocol_capability_received,omitempty"`
			NbrClusterID                            uint32 `json:"nbr_cluster_id,omitempty"`
			NbrEnforceFirstAs                       bool   `json:"nbr_enforce_first_as,omitempty"`
			NbrInCluster                            uint32 `json:"nbr_in_cluster,omitempty"`
			NegotiatedProtocolVersion               uint32 `json:"negotiated_protocol_version,omitempty"`
			NeighborInterfaceHandle                 uint32 `json:"neighbor_interface_handle,omitempty"`
			NeighborLocalAs                         uint32 `json:"neighbor_local_as,omitempty"`
			NsrEnabled                              bool   `json:"nsr_enabled,omitempty"`
			NsrState                                string `json:"nsr_state,omitempty"`
			OpenCheckErrorCode                      string `json:"open_check_error_code,omitempty"`
			PeerErrorCode                           uint32 `json:"peer_error_code,omitempty"`
			PeerResetReason                         string `json:"peer_reset_reason,omitempty"`
			PerformanceStatistics                   struct {
				ActiveCollision             uint32 `json:"active_collision,omitempty"`
				Actives                     uint32 `json:"actives,omitempty"`
				ControlToReadThreadTrigger  uint32 `json:"control_to_read_thread_trigger,omitempty"`
				ControlToWriteThreadTrigger uint32 `json:"control_to_write_thread_trigger,omitempty"`
				DataBytesRead               uint32 `json:"data_bytes_read,omitempty"`
				DataBytesWritten            uint32 `json:"data_bytes_written,omitempty"`
				FailedPostActives           uint32 `json:"failed_post_actives,omitempty"`
				HighThrottledRead           uint32 `json:"high_throttled_read,omitempty"`
				InboundUpdateMessages       uint32 `json:"inbound_update_messages,omitempty"`
				InboundUpdateMessagesTime   uint32 `json:"inbound_update_messages_time,omitempty"`
				IoReadTime                  uint32 `json:"io_read_time,omitempty"`
				IoWriteTime                 uint32 `json:"io_write_time,omitempty"`
				LastNsrScopedSync           uint32 `json:"last_nsr_scoped_sync,omitempty"`
				LastSentSeqNo               uint32 `json:"last_sent_seq_no,omitempty"`
				LowThrottledRead            uint32 `json:"low_throttled_read,omitempty"`
				MaximumReadSize             uint32 `json:"maximum_read_size,omitempty"`
				NbrFd                       uint32 `json:"nbr_fd,omitempty"`
				NbrFlags                    uint32 `json:"nbr_flags,omitempty"`
				NetworkStatus               uint32 `json:"network_status,omitempty"`
				NsrOperDownCount            uint32 `json:"nsr_oper_down_count,omitempty"`
				PassiveCollision            uint32 `json:"passive_collision,omitempty"`
				Passives                    uint32 `json:"passives,omitempty"`
				ReadCallsCount              uint32 `json:"read_calls_count,omitempty"`
				ReadMessagesCount           uint32 `json:"read_messages_count,omitempty"`
				ReadThrottles               uint32 `json:"read_throttles,omitempty"`
				RejectedPassives            uint32 `json:"rejected_passives,omitempty"`
				ResetFlags                  uint32 `json:"reset_flags,omitempty"`
				ResetRetries                uint32 `json:"reset_retries,omitempty"`
				SubgroupListTime            uint32 `json:"subgroup_list_time,omitempty"`
				SyncFlags                   uint32 `json:"sync_flags,omitempty"`
				TimeSinceLastThrottledRead  uint32 `json:"time_since_last_throttled_read,omitempty"`
				WriteCallsCount             uint32 `json:"write_calls_count,omitempty"`
				WriteQueueCallsCount        uint32 `json:"write_queue_calls_count,omitempty"`
				WriteQueueMessagesCount     uint32 `json:"write_queue_messages_count,omitempty"`
				WriteQueueTime              uint32 `json:"write_queue_time,omitempty"`
				WriteSubgroupCallsCount     uint32 `json:"write_subgroup_calls_count,omitempty"`
				WriteSubgroupMessagesCount  uint32 `json:"write_subgroup_messages_count,omitempty"`
			} `json:"performance_statistics,omitempty"`
			PreviousConnectionState  uint32 `json:"previous_connection_state,omitempty"`
			ReadArmed                bool   `json:"read_armed,omitempty"`
			ReceivedNotificationInfo struct {
				NotificationErrorCode     uint32 `json:"notification_error_code,omitempty"`
				NotificationErrorSubcode  uint32 `json:"notification_error_subcode,omitempty"`
				TimeSinceLastNotification uint32 `json:"time_since_last_notification,omitempty"`
			} `json:"received_notification_info,omitempty"`
			RemoteAs                uint32 `json:"remote_as,omitempty"`
			RemoteAsNumber          uint32 `json:"remote_as_number,omitempty"`
			ResetNotificationSent   bool   `json:"reset_notification_sent,omitempty"`
			ResetReason             string `json:"reset_reason,omitempty"`
			RouterID                net.IP `json:"router_id,omitempty"`
			RpkiAllowInvalid        bool   `json:"rpki_allow_invalid,omitempty"`
			RpkiDisable             bool   `json:"rpki_disable,omitempty"`
			RpkiSignalIbgp          bool   `json:"rpki_signal_ibgp,omitempty"`
			RpkiUseValidity         bool   `json:"rpki_use_validity,omitempty"`
			SecondLastKAexpiryReset uint32 `json:"second_last_k_aexpiry_reset,omitempty"`
			SecondLastKAstartReset  uint32 `json:"second_last_k_astart_reset,omitempty"`
			SecondLastWriteEvent    uint32 `json:"second_last_write_event,omitempty"`
			SendNotificationInfo    struct {
				NotificationErrorCode     uint32 `json:"notification_error_code,omitempty"`
				NotificationErrorSubcode  uint32 `json:"notification_error_subcode,omitempty"`
				TimeSinceLastNotification uint32 `json:"time_since_last_notification,omitempty"`
			} `json:"send_notification_info,omitempty"`
			SocketReadBytes                uint32 `json:"socket_read_bytes,omitempty"`
			SpeakerID                      uint32 `json:"speaker_id,omitempty"`
			StandbyRp                      bool   `json:"standby_rp,omitempty"`
			Suppress4ByteAs                bool   `json:"suppress4_byte_as,omitempty"`
			TCPSessionOpenMode             string `json:"tcp_session_open_mode,omitempty"`
			Tcpmss                         uint32 `json:"tcpmss,omitempty"`
			TimeLastCb                     uint64 `json:"time_last_cb,omitempty"`
			TimeLastCbReset                uint32 `json:"time_last_cb_reset,omitempty"`
			TimeLastFb                     uint64 `json:"time_last_fb,omitempty"`
			TimeSinceConnectionLastDropped uint32 `json:"time_since_connection_last_dropped,omitempty"`
			TimeSinceLastRead              uint32 `json:"time_since_last_read,omitempty"`
			TimeSinceLastReadReset         uint32 `json:"time_since_last_read_reset,omitempty"`
			TimeSinceLastUpdate            uint32 `json:"time_since_last_update,omitempty"`
			TimeSinceLastWrite             uint32 `json:"time_since_last_write,omitempty"`
			TimeSinceLastWriteReset        uint32 `json:"time_since_last_write_reset,omitempty"`
			TimeSinceSecondLastWrite       uint32 `json:"time_since_second_last_write,omitempty"`
			TimeSinceSecondLastWriteReset  uint32 `json:"time_since_second_last_write_reset,omitempty"`
			TosType                        uint32 `json:"tos_type,omitempty"`
			TosValue                       uint32 `json:"tos_value,omitempty"`
			TTLSecurityEnabled             bool   `json:"ttl_security_enabled,omitempty"`
			UpdateBytesRead                uint32 `json:"update_bytes_read,omitempty"`
			UpdateMessagesIn               uint32 `json:"update_messages_in,omitempty"`
			UpdateMessagesOut              uint32 `json:"update_messages_out,omitempty"`
			VrfName                        string `json:"vrf_name,omitempty"`
			WriteArmed                     bool   `json:"write_armed,omitempty"`
		} `json:"Content,omitempty"`
		Keys struct {
			AfName          string `json:"af_name,omitempty"`
			InstanceName    string `json:"instance_name,omitempty"`
			NeighborAddress net.IP `json:"neighbor_address,omitempty"`
		} `json:"Keys,omitempty"`
		Timestamp uint64 `json:"Timestamp,omitempty"`
	} `json:"Rows,omitempty"`
	Source    string `json:"Source,omitempty"`
	Telemetry struct {
		CollectionEndTime   uint64 `json:"collection_end_time,omitempty"`
		CollectionID        uint64 `json:"collection_id,omitempty"`
		CollectionStartTime uint64 `json:"collection_start_time,omitempty"`
		EncodingPath        string `json:"encoding_path,omitempty"`
		MsgTimestamp        uint64 `json:"msg_timestamp,omitempty"`
		NodeIDStr           string `json:"node_id_str,omitempty"`
		SubscriptionIDStr   string `json:"subscription_id_str,omitempty"`
	} `json:"Telemetry,omitempty"`
}

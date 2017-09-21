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

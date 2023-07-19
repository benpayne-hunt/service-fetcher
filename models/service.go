package models

type Service struct {
	Id     string   `json:"_id,omitempty"`
	Name   string   `json:"name,omitempty" validate:"required"`
	Tags   []string `json:"tags,omitempty" validate:"required"`
	Active bool     `json:"active,omitempty"`
}

// string $_id
// string $id
// string $name
// array $rate_card
// array|null $auto_decline
// ?int $teamSize
// ?bool $availability_set_for_full_job_window
// ?array $trackingConfig
// ?array $driver_requirements
// ?string $jobTypeDescription
// ?array $first_generated_route_settings
// 'parent_id',
// 'active',
// 'name',
// 'service',
// 'tag',
// 'tags',
// 'driver_requirements',
// 'description',
// 'popup_enabled',
// 'popup_html',
// 'thumbnail_url',
// 'infinity_scroll',
// 'generateChildLegs',
// 'dates',
// 'availability_override',
// 'availability_option_availability_max_score',
// 'availability_option_availability_max_miles',
// 'availability_option_availability_days_to_generate',
// 'availability_option_h3_resolution',
// 'availability_option_h3_neighbour_distance',
// 'availability_option_enable_queue',
// 'timeslot_layout',
// 'time_templates',
// 'tracking',
// 'trackingConfig',
// 'time_windows',
// 'instructions',
// 'return_unavailabile_timeslots',
// 'on_demand_available_between',
// 'only_use_timeslot_label',
// 'only_use_timeslot_label_text',
// 'show_addresses',
// 'show_details_html',
// 'show_details_html_value',
// 'select_first_timeslot',
// 'confirmation_required',
// 'confirmation_html',
// 'price_template',
// 'logo_url',
// 'pricing_type',
// 'order',
// 'allocation_logic',
// 'allocation_logic_driver',
// 'allocation_logic_closest_limit',
// 'allocation_logic_options',
// 'selected_bins',
// 'Criteria',
// 'deposit_value',
// 'deposit_on',
// 'label_status_flow_config',
// 'ticket_status_flow_config',
// 'ticket_status_flow_config_v2',
// 'job_request_status_flow_config',
// 'visibility_time_range',
// 'use_integration_fleet_id',
// 'integration_fleet_id',
// 'integration_fleet_vehicle',
// 'unassigned_settings',
// 'check_availability_settings_enabled',
// 'check_availability_settings',
// 'enable_ticket_config',
// 'enable_job_offer_config',
// 'zone_based_max_score',
// 'auto_unserviceable',
// 'minimum_hours_before_timeslot_starts',
// 'minimum_hours_before_timeslot_ends',
// 'smart_routing',
// 'smart_routing_options',
// 'fees',
// 'rate_card',
// 'first_generated_route_settings',
// 'info',
// 'storage_bin_options',
// 'percentage_pick_for_me_price_increase',
// 'auto_decline',
// 'teamSize',
// 'availability_set_for_full_job_window',
// 'jobTypeDescription',

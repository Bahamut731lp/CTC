package main

func main() {
	config := map[string]interface{}{
		"cars": map[string]interface{}{
			"count":            150000,
			"arrival_time_min": 1,
			"arrival_time_max": 2,
		},
		"stations": map[string]interface{}{
			"gas": map[string]interface{}{
				"count":          2,
				"serve_time_min": 2,
				"serve_time_max": 5,
			},
			"diesel": map[string]interface{}{
				"count":          2,
				"serve_time_min": 3,
				"serve_time_max": 6,
			},
			"lpg": map[string]interface{}{
				"count":          1,
				"serve_time_min": 4,
				"serve_time_max": 7,
			},
			"electric": map[string]interface{}{
				"count":          1,
				"serve_time_min": 5,
				"serve_time_max": 10,
			},
		},
		"registers": map[string]interface{}{
			"count":           2,
			"handle_time_min": 1,
			"handle_time_max": 3,
		},
	}

	print(config)
}

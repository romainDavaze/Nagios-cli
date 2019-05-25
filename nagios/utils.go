package nagios

import "encoding/json"

// AddApplyConfigToJSON adds the applyconfig attribute to a json string
func AddApplyConfigToJSON(data []byte) ([]byte, error) {
	var d map[string]interface{}
	json.Unmarshal(data, &d)
	d["applyconfig"] = 0
	return json.Marshal(d)
}

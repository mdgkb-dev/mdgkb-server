package schema

type Schema struct {
	HumanSchema map[string]string `json:"humanSchema"`
}

func CreateSchema() Schema {
	return Schema{
		HumanSchema: createHumanSchema(),
	}
}

func createHumanSchema() map[string]string {
	return map[string]string{
		"tableName": "human",
		"dateBirth": "date_birth",
		"fullName":  "full_name",
		"isMale":    "is_male",
	}
}

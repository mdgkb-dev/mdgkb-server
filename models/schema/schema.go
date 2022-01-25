package schema

type Schema struct {
	HumanSchema    map[string]string `json:"humanSchema"`
	CommentsSchema map[string]string `json:"commentsSchema"`
}

func CreateSchema() Schema {
	return Schema{
		HumanSchema:    createHumanSchema(),
		CommentsSchema: createCommentsSchema(),
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

func createCommentsSchema() map[string]string {
	return map[string]string{
		"tableName":   "comment",
		"publishedOn": "publishedOn",
		"positive":    "positive",
	}
}

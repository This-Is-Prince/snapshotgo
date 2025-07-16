package snapshotgo

type Args struct {
	Query     string
	Variables map[string]any
}

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLResponse struct {
	Data   any            `json:"data"`
	Errors []GraphQLError `json:"errors"`
}

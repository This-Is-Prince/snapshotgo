package snapshotgo

type Args struct {
	Query     string
	Variables map[string]any
}

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLResponse[T any] struct {
	Data   T              `json:"data"`
	Errors []GraphQLError `json:"errors"`
}

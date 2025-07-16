package main

import (
	"fmt"

	"github.com/This-Is-Prince/snapshotgo"
)

const SPACES_QUERY = `
query Spaces($first: Int, $skip: Int) {
  spaces(
    first: $first,
    skip: $skip
  ) {
    id
    name
    about
    twitter
    github
  }
}
`

type SpacesResponse struct {
	Spaces []Space `json:"spaces"`
}

type Space struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	About   string `json:"about"`
	Twitter string `json:"twitter"`
	Github  string `json:"github"`
}

func main() {
	snapshotHub := snapshotgo.NewSnapshot()

	args := snapshotgo.Args{
		Query: SPACES_QUERY,
		Variables: map[string]interface{}{
			"first": 20,
			"skip":  0,
		},
	}

	var data SpacesResponse

	err := snapshotgo.Query(snapshotHub, args, &data)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	fmt.Printf("%v", data)
}

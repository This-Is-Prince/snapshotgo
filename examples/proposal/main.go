package main

import (
	"fmt"

	"github.com/This-Is-Prince/snapshotgo"
)

const PROPOSAL_QUERY = `
query Proposal($id: String!) {
  proposal(
    id: $id
  ) {
    id
    title
    body
    choices
    start
    end
    snapshot
    state
    author
    space {
      id
      name
    }
  }
}
`

const PROPOSAL_ID = "0x586de5bf366820c4369c041b0bbad2254d78fafe1dcc1528c1ed661bb4dfb671"

type ProposalResponse struct {
	Proposal Proposal `json:"proposal"`
}

type Proposal struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	snapshotHub := snapshotgo.NewSnapshot()

	args := snapshotgo.Args{
		Query: PROPOSAL_QUERY,
		Variables: map[string]interface{}{
			"id": PROPOSAL_ID,
		},
	}

	var data ProposalResponse

	err := snapshotgo.Query(snapshotHub, args, &data)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	fmt.Printf("%#v", data)
}

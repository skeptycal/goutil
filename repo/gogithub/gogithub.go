package gogithub

import (
	"context"

	"github.com/shurcooL/githubv4"
)

const nl = '\n'

var NL = []byte{nl}
var ctx = context.Background()

type Repo struct {
	org      string
	reponame string
	// repo     *git.Repository
}

type Org struct {
	name string
	id   int64
}

// func (o *Org) List() []*Repo {

// }

func Test() error {

	// orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func Test2() error {
	return nil
}

// query is a sample query
//  query {
//  	viewer {
//  		login
//  		createdAt
//  	}
//  }
var query struct {
	Viewer struct {
		Login     githubv4.String
		CreatedAt githubv4.DateTime
	}
}

func Query(ctx context.Context) (string, error) {
	err := client.Query(ctx, &query, nil)
	if err != nil {
		return "", err
	}
	sb := NewSBWriter()
	defer sb.Reset()

	sb.Println("    Login:", query.Viewer.Login)
	sb.Println("CreatedAt:", query.Viewer.CreatedAt)
	return sb.String(), nil
}

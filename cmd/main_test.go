package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

var jsonData = []byte(`{
	"zen": "Favor focus over features.",
	"hook_id": 1,
	"hook": {
	  "type": "Repository",
	  "id": 1,
	  "name": "web",
	  "active": true,
	  "events": [
		"push"
	  ],
	  "config": {
		"content_type": "json",
		"insecure_ssl": "0",
		"secret": "********",
		"url": "https://direktiv"
	  },
	  "updated_at": "2023-07-12T22:36:07Z",
	  "created_at": "2023-07-12T22:36:07Z",
	  "url": "https://api.github.com/repos/dummy/test-flow/hooks/423816598",
	  "test_url": "https://api.github.com/repos/dummy/test-flow/hooks/423816598/test",
	  "ping_url": "https://api.github.com/repos/dummy/test-flow/hooks/423816598/pings",
	  "deliveries_url": "https://api.github.com/repos/dummy/test-flow/hooks/423816598/deliveries",
	  "last_response": {
		"code": null,
		"status": "unused",
		"message": null
	  }
	},
	"repository": {
	  "id": 1,
	  "node_id": "1",
	  "name": "test-flow",
	  "full_name": "dummy/test-flow",
	  "private": false,
	  "owner": {
		"login": "dummy",
		"id": 1,
		"node_id": "1",
		"avatar_url": "https://avatars.githubusercontent.com/u/10879611?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/dummy",
		"html_url": "https://github.com/dummy",
		"followers_url": "https://api.github.com/users/dummy/followers",
		"following_url": "https://api.github.com/users/dummy/following{/other_user}",
		"gists_url": "https://api.github.com/users/dummy/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/dummy/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/dummy/subscriptions",
		"organizations_url": "https://api.github.com/users/dummy/orgs",
		"repos_url": "https://api.github.com/users/dummy/repos",
		"events_url": "https://api.github.com/users/dummy/events{/privacy}",
		"received_events_url": "https://api.github.com/users/dummy/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "html_url": "https://github.com/dummy/test-flow",
	  "description": null,
	  "fork": false,
	  "url": "https://api.github.com/repos/dummy/test-flow",
	  "forks_url": "https://api.github.com/repos/dummy/test-flow/forks",
	  "keys_url": "https://api.github.com/repos/dummy/test-flow/keys{/key_id}",
	  "collaborators_url": "https://api.github.com/repos/dummy/test-flow/collaborators{/collaborator}",
	  "teams_url": "https://api.github.com/repos/dummy/test-flow/teams",
	  "hooks_url": "https://api.github.com/repos/dummy/test-flow/hooks",
	  "issue_events_url": "https://api.github.com/repos/dummy/test-flow/issues/events{/number}",
	  "events_url": "https://api.github.com/repos/dummy/test-flow/events",
	  "assignees_url": "https://api.github.com/repos/dummy/test-flow/assignees{/user}",
	  "branches_url": "https://api.github.com/repos/dummy/test-flow/branches{/branch}",
	  "tags_url": "https://api.github.com/repos/dummy/test-flow/tags",
	  "blobs_url": "https://api.github.com/repos/dummy/test-flow/git/blobs{/sha}",
	  "git_tags_url": "https://api.github.com/repos/dummy/test-flow/git/tags{/sha}",
	  "git_refs_url": "https://api.github.com/repos/dummy/test-flow/git/refs{/sha}",
	  "trees_url": "https://api.github.com/repos/dummy/test-flow/git/trees{/sha}",
	  "statuses_url": "https://api.github.com/repos/dummy/test-flow/statuses/{sha}",
	  "languages_url": "https://api.github.com/repos/dummy/test-flow/languages",
	  "stargazers_url": "https://api.github.com/repos/dummy/test-flow/stargazers",
	  "contributors_url": "https://api.github.com/repos/dummy/test-flow/contributors",
	  "subscribers_url": "https://api.github.com/repos/dummy/test-flow/subscribers",
	  "subscription_url": "https://api.github.com/repos/dummy/test-flow/subscription",
	  "commits_url": "https://api.github.com/repos/dummy/test-flow/commits{/sha}",
	  "git_commits_url": "https://api.github.com/repos/dummy/test-flow/git/commits{/sha}",
	  "comments_url": "https://api.github.com/repos/dummy/test-flow/comments{/number}",
	  "issue_comment_url": "https://api.github.com/repos/dummy/test-flow/issues/comments{/number}",
	  "contents_url": "https://api.github.com/repos/dummy/test-flow/contents/{+path}",
	  "compare_url": "https://api.github.com/repos/dummy/test-flow/compare/{base}...{head}",
	  "merges_url": "https://api.github.com/repos/dummy/test-flow/merges",
	  "archive_url": "https://api.github.com/repos/dummy/test-flow/{archive_format}{/ref}",
	  "downloads_url": "https://api.github.com/repos/dummy/test-flow/downloads",
	  "issues_url": "https://api.github.com/repos/dummy/test-flow/issues{/number}",
	  "pulls_url": "https://api.github.com/repos/dummy/test-flow/pulls{/number}",
	  "milestones_url": "https://api.github.com/repos/dummy/test-flow/milestones{/number}",
	  "notifications_url": "https://api.github.com/repos/dummy/test-flow/notifications{?since,all,participating}",
	  "labels_url": "https://api.github.com/repos/dummy/test-flow/labels{/name}",
	  "releases_url": "https://api.github.com/repos/dummy/test-flow/releases{/id}",
	  "deployments_url": "https://api.github.com/repos/dummy/test-flow/deployments",
	  "created_at": "2023-04-10T11:56:22Z",
	  "updated_at": "2023-04-11T22:23:33Z",
	  "pushed_at": "2023-04-27T09:45:04Z",
	  "git_url": "git://github.com/dummy/test-flow.git",
	  "ssh_url": "git@github.com:dummy/test-flow.git",
	  "clone_url": "https://github.com/dummy/test-flow.git",
	  "svn_url": "https://github.com/dummy/test-flow",
	  "homepage": null,
	  "size": 13,
	  "stargazers_count": 1,
	  "watchers_count": 1,
	  "language": "Python",
	  "has_issues": true,
	  "has_projects": true,
	  "has_downloads": true,
	  "has_wiki": true,
	  "has_pages": false,
	  "has_discussions": false,
	  "forks_count": 0,
	  "mirror_url": null,
	  "archived": false,
	  "disabled": false,
	  "open_issues_count": 0,
	  "license": {
		"key": "apache-2.0",
		"name": "Apache License 2.0",
		"spdx_id": "Apache-2.0",
		"url": "https://api.github.com/licenses/apache-2.0",
		"node_id": "1"
	  },
	  "allow_forking": true,
	  "is_template": false,
	  "web_commit_signoff_required": false,
	  "topics": [],
	  "visibility": "public",
	  "forks": 0,
	  "open_issues": 0,
	  "watchers": 1,
	  "default_branch": "main"
	},
	"sender": {
	  "login": "dummy",
	  "id": 1,
	  "node_id": "1",
	  "avatar_url": "https://avatars.githubusercontent.com/u/10879611?v=4",
	  "gravatar_id": "",
	  "url": "https://api.github.com/users/dummy",
	  "html_url": "https://github.com/dummy",
	  "followers_url": "https://api.github.com/users/dummy/followers",
	  "following_url": "https://api.github.com/users/dummy/following{/other_user}",
	  "gists_url": "https://api.github.com/users/dummy/gists{/gist_id}",
	  "starred_url": "https://api.github.com/users/dummy/starred{/owner}{/repo}",
	  "subscriptions_url": "https://api.github.com/users/dummy/subscriptions",
	  "organizations_url": "https://api.github.com/users/dummy/orgs",
	  "repos_url": "https://api.github.com/users/dummy/repos",
	  "events_url": "https://api.github.com/users/dummy/events{/privacy}",
	  "received_events_url": "https://api.github.com/users/dummy/received_events",
	  "type": "User",
	  "site_admin": false
	}
  }`)

var receiver testServer

func init() {

	receiver = testServer{}
	receiver.prepareReceiver()

	go receiver.startReceiver()

	os.Setenv("DIREKTIV_GITHUB_ENDPOINT", fmt.Sprintf("http://%s", receiver.addr))

	go startServer()

	time.Sleep(1 * time.Second)

}

func TestSending(t *testing.T) {

	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/github", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Github-Event", "ping")
	request.Header.Set("X-GitHub-Delivery", "7e618490-2104-11ee-9afd-842cc89f2d83")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer response.Body.Close()

	for {
		if receiver.lastRequest != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	t.Log(receiver.lastRequest)

}

type testServer struct {
	addr        string
	hasError    bool
	lastRequest map[string]interface{}
	lastHeaders map[string]string
}

func (s *testServer) prepareReceiver() {

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic("can not get listener")
	}
	defer l.Close()

	s.addr = l.Addr().String()

}

func (s *testServer) startReceiver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		s.lastHeaders = make(map[string]string)

		for name, values := range r.Header {
			for _, value := range values {
				s.lastHeaders[name] = value
			}
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			s.hasError = true
			return
		}

		var resp map[string]interface{}

		err = json.Unmarshal(b, &resp)
		if err != nil {
			s.hasError = true
			return
		}

		s.lastRequest = resp

	})

	log.Fatal(http.ListenAndServe(s.addr, nil))

}

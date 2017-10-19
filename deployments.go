//
// Copyright 2017, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// DeploymentsService handles communication with the commit related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type DeploymentsService struct {
	client *Client
}

// ListCommitsOptions represents the available ListCommits() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#list-repository-commits
type ListDeploymentsOptions struct {
	ListOptions
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *string `url:"sort,omitempty" json:"sort,omitempty"`
	Search  *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListDeployments gets a list of repository commits in a project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#list-commits
func (s *DeploymentsService) ListDeployments(pid interface{}, opt ListDeploymentsOptions, options ...OptionFunc) ([]*Deployment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/deployments", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var d []*Deployment
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

type Deployment struct {
	CreatedAt  time.Time `json:"created_at"`
	Deployable struct {
		Commit struct {
			AuthorEmail string    `json:"author_email"`
			AuthorName  string    `json:"author_name"`
			CreatedAt   time.Time `json:"created_at"`
			ID          string    `json:"id"`
			Message     string    `json:"message"`
			ShortID     string    `json:"short_id"`
			Title       string    `json:"title"`
		} `json:"commit"`
		Coverage   interface{} `json:"coverage"`
		CreatedAt  time.Time   `json:"created_at"`
		FinishedAt time.Time   `json:"finished_at"`
		ID         int         `json:"id"`
		Name       string      `json:"name"`
		Ref        string      `json:"ref"`
		Runner     interface{} `json:"runner"`
		Stage      string      `json:"stage"`
		StartedAt  interface{} `json:"started_at"`
		Status     string      `json:"status"`
		Tag        bool        `json:"tag"`
		User       struct {
			AvatarURL  string      `json:"avatar_url"`
			Bio        interface{} `json:"bio"`
			CreatedAt  time.Time   `json:"created_at"`
			ID         int         `json:"id"`
			Linkedin   string      `json:"linkedin"`
			Location   interface{} `json:"location"`
			Name       string      `json:"name"`
			Skype      string      `json:"skype"`
			State      string      `json:"state"`
			Twitter    string      `json:"twitter"`
			Username   string      `json:"username"`
			WebURL     string      `json:"web_url"`
			WebsiteURL string      `json:"website_url"`
		} `json:"user"`
	} `json:"deployable"`
	Environment struct {
		ExternalURL string `json:"external_url"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"environment"`
	ID   int    `json:"id"`
	Iid  int    `json:"iid"`
	Ref  string `json:"ref"`
	Sha  string `json:"sha"`
	User struct {
		AvatarURL string `json:"avatar_url"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
		State     string `json:"state"`
		Username  string `json:"username"`
		WebURL    string `json:"web_url"`
	} `json:"user"`
}

package types

import "time"

type (
	GithubPRReviewData struct {
		Id     int    `json:"id"`
		NodeId string `json:"node_id"`
		User   struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Body              string `json:"body"`
		State             string `json:"state"`
		HtmlUrl           string `json:"html_url"`
		PullRequestUrl    string `json:"pull_request_url"`
		AuthorAssociation string `json:"author_association"`
		Links             struct {
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			PullRequest struct {
				Href string `json:"href"`
			} `json:"pull_request"`
		} `json:"_links"`
		SubmittedAt time.Time `json:"submitted_at"`
		CommitId    string    `json:"commit_id"`
	}

	GithubPRInfoData struct {
		PRUrl        string `json:"html_url"`
		Description  string `json:"body"`
		Merged       bool   `json:"merged"`
		Mergeable    bool   `json:"mergeable"`
		Commits      int    `json:"commits"`
		Additions    int    `json:"additions"`
		Deletions    int    `json:"deletions"`
		ChangedFiles int    `json:"changed_files"`
		Head         struct {
			Ref string `json:"ref"`
		} `json:"head"`
		Base struct {
			Ref string `json:"ref"`
		} `json:"base"`
	}

	MergePayload struct {
		CommitTitle string `json:"commit_title"`
		MergeMethod string `json:"merge_method"`
	}

	MergeResponseData struct {
		Sha     string `json:"sha"`
		Merged  bool   `json:"merged"`
		Message string `json:"message"`
	}

	FailedMergeResponse struct {
		Message string `json:"message"`
	}
)

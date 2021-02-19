package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	token	string
	client	http.Client
}

type RepoOwner struct {
	Login				string	`json:"login"`
	ID					int64	`json:"id"`
	NodeID				string	`json:"node_id"`
	AvatarURL			string	`json:"avatar_url"`
	GravatarID			string	`json:"gravatar_id"`
	URL					string	`json:"url"`
	HtmlURL				string	`json:"html_url"`
	FollowersURL		string	`json:"followers_url"`
	FollowingURL		string	`json:"following_url"`
	GistsURL			string	`json:"gists_url"`
	StarredURL			string	`json:"starred_url"`
	Subscriptions		string	`json:"subscriptions_url"`
	OrganizationsURL	string	`json:"organizations_url"`
	ReposURL			string	`json:"repos_url"`
	EventsURL			string	`json:"events_url"`
	ReceivedEvents		string	`json:"received_events_url"`
	Type				string	`json:"type"`
	SiteAdmin			bool	`json:"site_admin"`
}

type RepoPermissions struct {
	Admin	bool	`json:"admin"`
	Push	bool	`json:"push"`
	Pull	bool	`json:"pull"`
}

type RepoLicense struct {
	Key		string	`json:"key"`
	Name	string	`json:"name"`
	SpdxID	string	`json:"spdx_id"`
	URL		string	`json:"url"`
	NodeID	string	`json:"node_id"`
}

type Repo struct {
	ID					int64		`json:"id"`
	NodeID				string		`json:"node_id"`
	Name				string		`json:"name"`
	FullName			string		`json:"full_name"`
	Private				bool		`json:"private"`
	Owner				RepoOwner	`json:"owner"`
	Html				string		`json:"html_url"`
	Description			string		`json:"description"`
	Fork				bool		`json:"fork"`
	URL					string		`json:"url"`
	ForksURL			string		`json:"forks_url"`
	KeysURL				string		`json:"keys_url"`
	CollaboratorsURL	string		`json:"collaborators_url"`
	TeamsURL			string		`json:"teams_url"`
	HooksURL			string		`json:"hooks_url"`
	IssueEventsURL		string		`json:"issue_events_url"`
	EventsURL			string		`json:"events_url"`
	AssigneesURL		string		`json:"assignees_url"`
	BranchesURL			string		`json:"branches_url"`
	TagsURL				string		`json:"tags_url"`
	BlobsURL			string		`json:"blobs_url"`
	GitTagsURL			string		`json:"git_tags_url"`
	GitRefsURL			string		`json:"git_refs_url"`
	TreesURL			string		`json:"trees_url"`
	StatusesURL			string		`json:"statuses_url"`
	LanguagesURL		string		`json:"languages_url"`
	StargazersURL		string		`json:"stargazers_url"`
	ContributorsURL		string		`json:"contributors_url"`
	SubscribersURL		string		`json:"subscribers_url"`
	SubscriptionURL		string		`json:"subscription_url"`
	CommitsURL			string		`json:"commits_url"`
	GitCommitsURL		string		`json:"git_commits_url"`
	CommentsURL			string		`json:"comments_url"`
	IssueCommentURL		string		`json:"issue_comment_url"`
	ContentsURL			string		`json:"contents_url"`
	CompareURL			string		`json:"compare_url"`
	MergesURL			string		`json:"merges_url"`
	ArchiveURL			string		`json:"archive_url"`
	DownloadsURL		string		`json:"downloads_url"`
	IssuesURL			string		`json:"issues_url"`
	PullsURL			string		`json:"pulls_url"`
	MilestonesURL		string		`json:"milestones_url"`
	NotificationsURL	string		`json:"notifications_url"`
	LabelsURL			string		`json:"labels_url"`
	ReleasesURL			string		`json:"releases_url"`
	DeploymentsURL		string		`json:"deployments_url"`
	CreatedAt			time.Time	`json:"created_at"`
	UpdatedAt			time.Time	`json:"updated_at"`
	PushedAt			time.Time	`json:"pushed_at"`
	GitURL				string		`json:"git_url"`
	SshURL				string		`json:"ssh_url"`
	CloneURL			string		`json:"clone_url"`
	SvnURL				string		`json:"svn_url"`
	Homepage			string		`json:"homepage"`
	Size				int			`json:"size"`
	StargazersCount		int			`json:"stargazers_count"`
	WatchersCount		int			`json:"watchers_count"`
	Language			string		`json:"language"`
	HasIssues			bool		`json:"has_issues"`
	HasProjects			bool		`json:"has_projects"`
	HasDownloads		bool		`json:"has_downloads"`
	HasWiki				bool		`json:"has_wiki"`
	HasPages			bool		`json:"has_pages"`
	ForksCount			int			`json:"forks_count"`
	MirrorURL			string		`json:"mirror_url"`
	Archived			bool		`json:"archived"`
	Disabled			bool		`json:"disabled"`
	OpenIssuesCount		int			`json:"open_issues_count"`
	License				RepoLicense	`json:"license"`
	Forks				int			`json:"forks"`
	OpenIssues			int			`json:"open_issues"`
	Watchers			int			`json:"watchers"`
	DefaultBranch		string		`json:"default_branch"`
	Permissions			RepoPermissions	`json:"permissions"`
}

const GitHubAPIUrl = "https://api.github.com/"
const PerPage = 100

func NewClient(token string) *Client {
	return &Client{
		client: http.Client{},
		token: token,
	}
}

func (c *Client) GetRepoList() []Repo {

	recordsResponse := []Repo{}

	condition := true
	for ok := true; ok; ok = condition {

		currentResult := []Repo{}

		resp, err := c.Get(fmt.Sprintf("%s%d","user/repos?per_page=", PerPage))
		if err != nil {
			log.Fatal("Cannot get repo list")
		}

		err = c.unmarshalJSON(resp.Body, &currentResult)

		if err != nil {
			fmt.Println(err)
			log.Fatal("Cannot parse response")
		}

		recordsResponse = append(recordsResponse, currentResult...)

		condition = len(currentResult) == PerPage
	}

	return recordsResponse
}

func (c *Client) Get(url string) (resp *http.Response, err error) {
	return c.newRequest("GET", url, nil)
}
//
//func (c *Client) Post(session *SessionResponse, url string, body io.Reader) (resp *http.Response, err error) {
//	return c.newRequest(session, "POST", url, body)
//}

//func (c *Client) Login(username string, password string) (*SessionResponse, error) {
//
//	if c.BaseLoginURL == nil {
//		return nil, errors.New("Base Login URL is not set")
//	}
//
//	payload := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8" ?>
//<env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema"
//    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//    xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
//  <env:Body>
//    <n1:login xmlns:n1="urn:partner.soap.sforce.com">
//      <n1:username>%s</n1:username>
//      <n1:password>%s</n1:password>
//    </n1:login>
//  </env:Body>
//</env:Envelope>`, username, password))
//
//	req, err := http.NewRequest("POST", *c.BaseLoginURL+APIUrl, bytes.NewReader(payload))
//
//	if err != nil {
//		log.Fatal("Error on creating request object. ", err.Error())
//		return nil, err
//	}
//
//	req.Header.Set("Content-Type", "text/xml")
//	req.Header.Set("SOAPAction", "login")
//
//	res, err := c.client.Do(req)
//
//	if err != nil {
//		log.Fatal("Error on dispatching request. ", err.Error())
//		return nil, err
//	}
//
//	loginResponse := &XMLResponse{}
//	body, err := ioutil.ReadAll(res.Body)
//
//	if err != nil {
//		log.Fatal("Error ready body request. ", err.Error())
//		return nil, err
//	}
//
//	err = xml.Unmarshal(body, loginResponse)
//	if err != nil {
//		log.Fatal("Error unmarshal xml. ", err.Error())
//		return nil, err
//	}
//
//	fmt.Println(string(body))
//
//	if res.StatusCode != 200 {
//		return nil, errors.New(loginResponse.Body.Fault.FaultCode)
//	}
//
//	urlStr, _ := url.Parse(loginResponse.Body.LoginResponse.Result.ServerURL)
//
	//c.token = &loginResponse.Body.LoginResponse.Result.SessionID
	//c.instanceURL = urlStr.Scheme + "://" + urlStr.Host

	//sessionResponse := &SessionResponse{
	//	SessionID:   loginResponse.Body.LoginResponse.Result.SessionID,
	//	InstanceURL: urlStr.Scheme + "://" + urlStr.Host,
	//}

	//return sessionResponse, nil
//}

func (c *Client) newRequest(
	method string,
	url string,
	body io.Reader) (resp *http.Response, err error) {

	req, err := http.NewRequest(method, GitHubAPIUrl + url, body)

	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return nil, err
	}

	req.Header.Set("Authorization", "token " + c.token)

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.client.Do(req)

	return res, err
}

func (c *Client) unmarshalJSON(body io.ReadCloser, v interface{}) error {
	b, err := ioutil.ReadAll(body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}
package main

import ()

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
)

type V1Api struct {
	endpoint string
	method   HttpMethod
}

type user struct {
	SelectedEmail     string             `json:"selected_email"`
	AvatarUrl         string             `json:"avatar_url"`
	TrialEnd          string             `json:"trial_end"`
	Admin             bool               `json:"admin"`
	BasicEmailPrefs   string             `json:"basic_email_prefs"`
	SignInCount       int                `json:"sign_in_count"`
	GithubOAuthScopes []string           `json:"github_oauth_scopes"`
	Name              string             `json:"name"`
	GravatarId        int64              `json:"gravatar_id"`
	DaysLeftInTrial   int                `json:"days_left_in_trial"`
	Parallelism       int                `json:"parallelism"`
	GithubId          int64              `json:"github_id"`
	AllEmails         []string           `json:"all_emails"`
	CreatedAt         string             `json:"created_at"`
	Plan              string             `json:"plan"`
	HerokuApiKey      string             `json:"heroku_api_key"`
	Projects          map[string]project `json:"projects"`
	Login             string             `json:"login"`
	Containers        int                `json:"containers"`
}

type project struct {
	Emails      string `json:"emails"`
	OnDashboard bool   `json:"on_dashboard"`
}

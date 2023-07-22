package data

import "sedwards2009/llm-workbench/internal/data/responsestatus"

type SessionOverview struct {
	SessionSummaries []*SessionSummary `json:"sessionSummaries"`
}

type SessionSummary struct {
	ID                string `json:"id"`
	CreationTimestamp string `json:"creationTimestamp"`
	Title             string `json:"title"`
}

type Root struct {
	Sessions []Session `json:"sessions"`
}

type Session struct {
	ID                string         `json:"id"`
	CreationTimestamp string         `json:"creationTimestamp"`
	Title             string         `json:"title"`
	Prompt            string         `json:"prompt"`
	Responses         []*Response    `json:"responses"`
	ModelSettings     *ModelSettings `json:"modelSettings"`
}

type ModelSettings struct {
	ModelID string `json:"modelId"`
}

type ModelOverview struct {
	Models []*Model `json:"models"`
}

type Model struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Engine          string
	InternalModelID string
}

//-------------------------------------------------------------------------

type Response struct {
	ID                string                        `json:"id"`
	CreationTimestamp string                        `json:"creationTimestamp"`
	Status            responsestatus.ResponseStatus `json:"status"`
	Prompt            string                        `json:"prompt"`
	Text              string                        `json:"text"`
}

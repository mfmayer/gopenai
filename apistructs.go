package gopenai

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleFunction  Role = "function"
)

type FunctionParameter struct {
	// Type should be one of "string", "number", "integer", "boolean", "array", or "object"
	Type string `json:"type"`
	// Description is a human-readable description of the parameter's meaning
	Description string `json:"description,omitempty"`
	// Enum is an array of possible values for the parameter
	Enum []string `json:"enum,omitempty"`
}

type FunctionParameters struct {
	// Should be of "object"
	Type string `json:"type"`
	// Description is a human-readable description of the parameters
	Properties map[string]FunctionParameter `json:"properties"`
	// Required is an array of required parameter names
	Required []string `json:"required,omitempty"`
}

type Function struct {
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.
	Name string `json:"name"`
	// Description is a human-readable description of what function does
	Description string `json:"description,omitempty"`
	// Parameters that the function accepts
	Parameters FunctionParameters `json:"parameters,omitempty"`
}

// FunctionCall represents a function call object in a chat message.
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// Message represents a message object in a chat conversation.
// Role is the sender's role, either "system", "user", or "assistant".
// Content is the text content of the message.
type Message struct {
	// The role of the messages author. One of system, user, assistant, or function
	Role Role `json:"role"`
	// The name of the author of this message. name is required if role is function, and it should be the name of the function whose response is in the content. May contain a-z, A-Z, 0-9, and underscores, with a maximum length of 64 characters
	Name string `json:"name,omitempty"`
	// The contents of the message. content is required for all messages except assistant messages with function calls
	Content string `json:"content"`
	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

// ChatPromptConfig contains model and model configuration parameters.
// Model is the ID of the OpenAI model used for the chat completion.
type ChatPromptConfig struct {
	// Model name/id of the OpenAI model used for the chat completion (e.g. "gpt-3.5-turbo")
	Model string `json:"model"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. OpenAIO generally recommends altering this or top_p but not both.
	Temperature float32 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered. OpenAI generally recommend altering this or temperature but not both
	TopP float32 `json:"top_p,omitempty"`
	// The maximum number of tokens to generate in the chat completion. The total length of input tokens and generated tokens is limited by the model's context length.
	MaxTokens int `json:"max_tokens,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	PresencePenalty float32 `json:"presence_penalty,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`
}

// ChatPrompt represents a chat prompt to be sent to the API.
type ChatPrompt struct {
	*ChatPromptConfig
	// A list of messages describing the conversation so far
	Messages []*Message `json:"messages"`
	// A list of functions the model may generate JSON inputs for.
	Functions []*Function `json:"functions,omitempty"`
}

// Choice represents a single choice or response from the API.
type Choice struct {
	// Message is a Message object representing the response message.
	Message Message `json:"message"`
	// FinishReason is the reason the API ended the message (e.g., "stop", "max_tokens", or "temperature").
	FinishReason string `json:"finish_reason"`
	// Index is the index of the choice in the response.
	Index int `json:"index"`
}

// Usage represents the token usage for a chat completion.
type Usage struct {
	// PromptTokens is the number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`
	// CompletionTokens is the number of tokens in the completion.
	CompletionTokens int `json:"completion_tokens"`
	// TotalTokens is the total number of tokens used in the API call.
	TotalTokens int `json:"total_tokens"`
}

// Error details
type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code,omitempty"`
	Param   string `json:"param,omitempty"`
}

// ChatCompletion represents the API response for a chat completion.
type ChatCompletion struct {
	// Error in case an error occured or was detected
	Error *Error `json:"error,omitempty"`
	// ID is the unique identifier for the completion.
	ID string `json:"id"`
	// Object is the type of object returned by the API, usually "text_completion".
	Object string `json:"object"`
	// Created is a Unix timestamp of when the completion was created.
	Created int `json:"created"`
	// Model is the ID of the OpenAI model used for the completion.
	Model string `json:"model"`
	// Usage is a Usage object representing the token usage for the completion.
	Usage Usage `json:"usage"`
	// Choices is an array of Choice objects representing the response choices.
	Choices []Choice `json:"choices"`
}

// Model details
type Model struct {
	ID          string                 `json:"id"`
	Object      string                 `json:"object"`
	Created     int                    `json:"created"`
	OwnedBy     string                 `json:"owned_by"`
	Permissions map[string]interface{} `json:"permissions"`
	Root        string                 `json:"root"`
	Parent      string                 `json:"parent"`
}

// List of models
type ModelList struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

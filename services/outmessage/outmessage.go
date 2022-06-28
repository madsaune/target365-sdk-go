package outmessage

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Out Message
type Message struct {
	// Transaction id. Must be unique per message if used. This can be used for guarding against resending messages
	TransactionID *string `json:"transactionId,omitempty"`
	// Session id. This can be used as the clients to get all out-messages associated to a specific session
	SessionID *string `json:"sessionId,omitempty"`
	// Correlation id. This can be used as the clients correlation id for tracking messages and delivery reports
	CorrelationID *string `json:"correlationId,omitempty"`
	// Keyword id associated with message. Can be null
	KeywordID *string `json:"keywordId,omitempty"`
	// Sender. Can be an alphanumeric string, a phone number or a short number
	Sender string `json:"sender"`
	// Recipient phone number
	Recipient string `json:"recipient"`
	// Content. The actual text message content
	Content string `json:"content"`
	// Send time, in UTC. If omitted the send time is set to ASAP
	SendTime *time.Time `json:"sendTime,omitempty"`
	// Message Time-To-Live (TTL) in minutes. Must be between 5 and 1440. Default value is 120
	TimeToLive *int `json:"timeToLive,omitempty"`
	// Priority. Can be 'Low', 'Normal' or 'High'. If omitted, default value is 'Normal'
	Priority Priority `json:"priority,omitempty"`
	// Message delivery mode. Can be either 'AtLeastOnce' or 'AtMostOnce'. If omitted, default value is 'AtMostOnce'
	DeliveryMode DeliveryMode `json:"deliveryMode,omitempty"`
	// Delivery report url
	DeliveryReportURL *string `json:"deliveryReportUrl,omitempty"`
	// True to allow unicode SMS, false to fail if content is unicode, null to replace unicode chars to '?'
	AllowUnicode *bool `json:"allowUnicode,omitempty"`
	// External SMSC transaction id
	SmscTransactionID *string `json:"smscTransactionId,omitempty"`
	// SMSC message parts
	SmscMessageParts *int `json:"smscMessageParts,omitempty"`
	// Last modified time
	LastModified *time.Time `json:"lastModified,omitempty"`
	// Created time
	Created *time.Time `json:"created,omitempty"`
	// Delivery status code. Can be 'Queued', 'Sent', 'Failed', 'Ok' or 'Reversed'
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// Whether message was delivered. Null if status is unknown
	Delivered *bool `json:"delivered,omitempty"`
	// Operator id (from delivery report)
	OperatorID *string `json:"operatorId,omitempty"`
	// Whether billing was performed. Null if status is unknown
	Billed *bool `json:"billed,omitempty"`
	// Custom properties associated with message
	Properties *map[string]string `json:"properties,omitempty"`
	// Tags associated with message. Can be used for statistics and grouping
	Tags *[]string `json:"tags,omitempty"`
}

func (m Message) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Sender, validation.Required),
		validation.Field(&m.Recipient, validation.Required),
		validation.Field(&m.Content, validation.Required),
	)
}

type Response struct {
	// Transaction id. Must be unique per message if used. This can be used for guarding against resending messages
	TransactionID string
	// Resource uri of created out-message.
	Location string
}

type Priority string

const (
	PriorityLow    Priority = "Low"
	PriorityNormal Priority = "Normal"
	PriorityHigh   Priority = "High"
)

type DeliveryMode string

const (
	DeliveryModeAtLeastOnce DeliveryMode = "AtLeastOnce"
	DeliveryModeAtMostOnce  DeliveryMode = "AtMostOnce"
)

type StatusCode string

const (
	StatusCodeQueued   StatusCode = "Queued"
	StatusCodeSent     StatusCode = "Sent"
	StatusCodeFailed   StatusCode = "Failed"
	StatusCodeOk       StatusCode = "Ok"
	StatusCodeReversed StatusCode = "Reversed"
)

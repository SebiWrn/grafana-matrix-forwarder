package formatter

import (
	"golang.org/x/exp/maps"
	"grafana-matrix-forwarder/matrix"
	"grafana-matrix-forwarder/model"
	"log"
	"slices"
)

type alertMessageData struct {
	MetricRounding int
	StateStr       string
	StateEmoji     string
	Payload        model.AlertData
}

func GenerateMessage(alert model.AlertData, metricRounding int) (matrix.FormattedMessage, error) {
	var messageData = alertMessageData{
		StateStr:       "UNKNOWN",
		StateEmoji:     "❓",
		MetricRounding: metricRounding,
		Payload:        alert,
	}
	switch alert.State {
	case model.AlertStateAlerting:
		messageData.StateStr = "ALERT"
		messageData.StateEmoji = "💔"
	case model.AlertStateResolved:
		messageData.StateStr = "RESOLVED"
		messageData.StateEmoji = "💚"
	case model.AlertStateNoData:
		messageData.StateStr = "NO DATA"
		messageData.StateEmoji = "❓"
	default:
		log.Printf("alert received with unknown state: %s", alert.State)
	}
	html, err := executeHtmlTemplate(alertMessageTemplate, messageData)
	if slices.Contains(maps.Keys(alert.Labels), "err_type") && alert.Labels["err_type"] == "disk_full" {
		html, err = executeHtmlTemplate(alertMessageDiskFullTemplate, messageData)
	}
	if err != nil {
		return matrix.FormattedMessage{}, err
	}
	text := htmlMessageToTextMessage(html)
	return matrix.FormattedMessage{
		TextBody: text,
		HtmlBody: html,
	}, err
}

func GenerateReply(originalHtmlMessage string, alert model.AlertData) (matrix.FormattedMessage, error) {
	if alert.State == model.AlertStateResolved {
		html, err := executeTextTemplate(resolveReplyTemplate, originalHtmlMessage)
		if err != nil {
			return matrix.FormattedMessage{}, err
		}
		text := resolveReplyPlainStr
		return matrix.FormattedMessage{
			TextBody: text,
			HtmlBody: html,
		}, err
	}
	return matrix.FormattedMessage{}, nil
}

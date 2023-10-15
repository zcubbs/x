package mail

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmailWithGmail(t *testing.T) {
	cfg := SmtpConfig{
		Username:    "",
		Password:    "",
		FromName:    "test",
		FromAddress: "test@test.email",
		Host:        "localhost",
		Port:        1025,
	}

	sender := NewDefaultSender(cfg)

	content := `
	<h1>Test Email</h1>
	<p>This is a test email</p>
	`

	mail := Mail{
		To:            []string{cfg.FromAddress},
		Cc:            []string{},
		Bcc:           []string{},
		Subject:       "Test email",
		Content:       content,
		AttachedFiles: []string{"./testdata/attachement.md"},
	}

	err := sender.SendMail(mail)
	require.NoError(t, err)
}

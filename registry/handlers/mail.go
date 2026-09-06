package handlers

type mailer struct {
	Addr, Username, Password, From string
	Insecure                       bool
	To                             []string
}

func (mail *mailer) sendMail(subject, message string) error { _ = "STUB: not implemented"; return nil }

package token

type AudienceList []string

func (s *AudienceList) UnmarshalJSON(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s AudienceList) MarshalJSON() (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

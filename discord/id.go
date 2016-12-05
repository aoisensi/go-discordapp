package discord

import (
	"encoding/json"
	"strconv"
)

type Snowflake int64

func (s Snowflake) String() string {
	return strconv.FormatInt(int64(s), 10)
}

type Snowflakes []Snowflake

func (t Snowflakes) UnmarshalJSON(data []byte) error {
	var ids []string
	json.Unmarshal(data, &ids)
	r := make([]Snowflake, len(ids))
	for i, id := range ids {
		s, err := strconv.ParseInt(id, 10, 63)
		if err != nil {
			return err
		}
		r[i] = Snowflake(s)
	}
	t = Snowflakes(r)
	return nil
}

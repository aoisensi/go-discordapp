package discord

import "strconv"

type Snowflake int64

func (s Snowflake) String() string {
	return strconv.FormatInt(int64(s), 10)
}

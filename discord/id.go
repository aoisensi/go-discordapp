package discord

import "strconv"
import "fmt"

type Snowflake int64

func (s Snowflake) String() string {
	return strconv.FormatInt(int64(s), 10)
}

func (s Snowflake) UnmarshalJSON(d []byte) error {
	var id int64
	if _, err := fmt.Scanf("\"%d\"", &id); err != nil {
		return err
	}
	s = Snowflake(id)
	return nil
}

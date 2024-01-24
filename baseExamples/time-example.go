package main

import (
	"fmt"
	"time"
)

func main() {
	prnt := fmt.Println

	now := time.Now()
	prnt(now)

	//year, month, day, hour, min, second, nanosecond?, utc
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	prnt(then)

	prnt("Year", then.Year())
	prnt("Month", then.Month())
	prnt("Day", then.Day())
	prnt("Hour", then.Hour())
	prnt("Minute", then.Minute())
	prnt("Second", then.Second())
	prnt("Nanosecond", then.Nanosecond())
	prnt("Location", then.Location())

	prnt("Weekday", then.Weekday())

	prnt("Before now", then.Before(now))
	prnt("After now", then.After(now))
	prnt("Equal to now", then.Equal(now))

	diff := now.Sub(then)
	prnt("diff", diff)

	prnt("diff hours", diff.Hours())
	prnt("diff minutes", diff.Minutes())
	prnt("diff seconds", diff.Seconds())
	prnt("diff nanoseconds", diff.Nanoseconds())

	prnt("Add onto time", then.Add(diff))
	prnt("negative diff to move backwards", then.Add(-diff))
}

package date

import (
	"database/sql/driver"
	"time"
)

type CustomDate struct {
	time.Time
}

// type CustomDate time.Time

const Layout = "2006-01-02"

func (t *CustomDate) UnmarshalJSON(data []byte) (err error) {
	parsedTime, err := time.Parse(`"`+Layout+`"`, string(data))
	t.Time = parsedTime

	return
}

func (t CustomDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(Layout)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, Layout)
	b = append(b, '"')

	return b, nil
}

func (t CustomDate) Value() (driver.Value, error) {
	tTime := t.Time
	return tTime.Format("2006/01/02 15:04:05"), nil
}

func (t *CustomDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		t.Time = vt
	case string:
		tTime, _ := time.Parse("2006/01/02 15:04:05", vt)
		t.Time = tTime
	}
	return nil
}
func (d CustomDate) String() string {
	return d.Format(Layout)
}

// // GetBSON implements bson.Getter.
// func (d CustomDate) GetBSON() (interface{}, error) {
// 	return struct {
// 		time.Time `json:"time" bson:"time"`
// 	}{d.Time}, nil
// }

// // SetBSON implements bson.Setter.
// func (d *CustomDate) SetBSON(raw bson.Raw) error {

// 	decoded := new(struct {
// 		time.Time `json:"time" bson:"time"`
// 	})

// 	if bsonErr == nil {
// 		d.Time = decoded.Time
// 		return nil
// 	} else {
// 		return bsonErr
// 	}
// }

func Today() CustomDate {
	return CustomDate{time.Now()}
}

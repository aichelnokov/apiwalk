package types

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

type PointBlob struct {
	X, Y float64
}

// func (p PointBlob) GormDataType() string {
//   return "point"
// }

func (p PointBlob) Value() (driver.Value, error) {
	out := []byte{'('}
	out = strconv.AppendFloat(out, p.X, 'f', -1, 64)
	out = append(out, ',')
	out = strconv.AppendFloat(out, p.Y, 'f', -1, 64)
	out = append(out, ')')
	return out, nil
}

// func (p PointBlob) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
//   return clause.Expr{
//     SQL:  "ST_PointFromText(?)",
//     Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", p.X, p.Y)},
//   }
// }

func (p *PointBlob) Scan(src interface{}) (err error) {
	var data []byte

	switch src := src.(type) {
		case []byte:
			data = src
		case string:
			data = []byte(src)
		case nil:
			return nil
		default:
			return errors.New("(*PointBlob).Scan: unsupported data type")
	}

	if len(data) == 0 {
		return nil
	}

	data = data[1 : len(data)-1] // drop the surrounding parentheses
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			if p.X, err = strconv.ParseFloat(string(data[:i]), 64); err != nil {
				return err
			}
			if p.Y, err = strconv.ParseFloat(string(data[i+1:]), 64); err != nil {
				return err
			}
			break
		}
	}
	return nil
}
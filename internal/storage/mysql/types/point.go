package types

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Point struct {
  X, Y float64
}

func (p Point) GormDataType() string {
  return "point"
}

func (p Point) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", p.X, p.Y)},
  }
}

// Scan implements the sql.Scanner interface
func (p *Point) Scan(src interface{}) (err error) {
  var data []byte
  
  switch src := src.(type) {
    case []byte:
      data = src
    case string:
      data = []byte(src)
    case nil:
      return nil
    default:
      return errors.New("(*Point).Scan: unsupported data type")
  }
  
  if len(data) == 0 {
    return nil
	}
  
  var pos int = 0
  for i := 0; i < len(data); i++ {
    if data[i] == '(' {
      pos = i + 1;
    }
    if data[i] == ',' {
      if p.X, err = strconv.ParseFloat(string(data[pos:i]), 64); err != nil {
				return err
			}
      if p.Y, err = strconv.ParseFloat(string(data[i+1:len(data)-1]), 64); err != nil {
        return err
      }
      break
    }
  }

  return nil;
}
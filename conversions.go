package coinpaprika

import "strconv"

func convertStrToFloatPtr(value string) (converted *float64, err error) {
	if value != "" {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		converted = &f
	}
	return converted, nil
}

func convertStrToIntPtr(value string) (converted *int64, err error) {
	if value != "" {
		f, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		converted = &f
	}
	return converted, nil
}

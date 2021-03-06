package database

import "encoding/json"

func isSuccessCode(statusCode int) bool {
	return statusCode < 300
}

func couchGetReasonFromBody(body []byte) (string, error) {
	bodyMap := map[string]interface{}{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		return "", err
	}

	reasonIn, exists := bodyMap["reason"]
	if !exists {
		log.Sugar().Warnf("couchGetReasonFromBody: No reason in body, returning empty string.")
		return "", nil
	}

	reason, ok := reasonIn.(string)
	if !ok {
		log.Sugar().Warnf("couchGetReasonFromBody: Reason is not a string, returning empty string.")
		return "", nil
	}
	return reason, nil
}

func couchDecodeViewResult(body []byte) (int, int, []interface{}, error) {
	bodyMap := map[string]interface{}{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		log.Sugar().Errorf("couchDecodeViewResult: Failed to unmarshal CouchDB view result into map %s", err.Error())
		return 0, 0, nil, err
	}

	offsetIn, offsetExists := bodyMap["offset"]
	totalRowsIn, totalRowsExists := bodyMap["total_rows"]
	rowsIn, rowsExists := bodyMap["rows"]

	var offset, totalRows float64
	var rows []interface{}

	if offsetExists {
		offset, _ = offsetIn.(float64)
	} else {
		log.Sugar().Warnf("couchDecodeViewResult: 'offset' key absent. Default value 0 will be returned.")
	}
	if totalRowsExists {
		totalRows = totalRowsIn.(float64)
	} else {
		log.Sugar().Warnf("couchDecodeViewResult: 'total_rows' key absent. Default value 0 will be returned.")
	}
	if rowsExists {
		rows = rowsIn.([]interface{})
	} else {
		log.Sugar().Warnf("couchDecodeViewResult: 'rows' key absent. Default value nil will be returned.")
	}

	return int(offset), int(totalRows), rows, nil
}

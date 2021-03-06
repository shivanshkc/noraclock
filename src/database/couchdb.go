package database

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"noraclock/src/configs"
	"noraclock/src/constants"
	"noraclock/src/exception"
	"noraclock/src/logger"
)

var conf = configs.Get()
var log = logger.General()

// CouchDB : Implements the CouchDB methods.
var CouchDB = &couchdb{}

type couchdb struct{}

// CreateDoc : Creates a doc in the database.
func (c *couchdb) CreateDoc(database string, docID string, doc []byte) error {
	endpoint := fmt.Sprintf("%s/%s/%s", conf.CouchDB.Address, database, docID)

	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(doc))
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDoc: Failed to create HTTP request: %s", err.Error())
		return err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDoc: Failed to make HTTP request: %s", err.Error())
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if isSuccessCode(resp.StatusCode) {
		return nil
	}
	log.Sugar().Infof("CouchDB.CreateDoc: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDoc: Failed to CouchDB response body: %s", err.Error())
		return exception.Unexpected()
	}

	reason, err := couchGetReasonFromBody(bodyBytes)
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDoc: Failed to obtain CouchDB error reason from body: %s", err.Error())
		return exception.Unexpected()
	}

	return errors.New(reason)
}

// UpdateDoc : Updates a doc in the database.
func (c *couchdb) UpdateDoc(database string, docID string, doc []byte) error {
	rev, err := c.GetDocRev(database, docID)
	if err != nil {
		log.Sugar().Errorf("CouchDB.UpdateDoc: Failed to get document revision for updation: %s", err.Error())
		return err
	}
	return c.UpdateDocWithRev(database, docID, rev, doc)
}

// UpdateDocWithRev : Updates a doc using the provided revision in the database.
func (c *couchdb) UpdateDocWithRev(database string, docID string, rev string, doc []byte) error {
	endpoint := fmt.Sprintf("%s/%s/%s", conf.CouchDB.Address, database, docID)

	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(doc))
	if err != nil {
		log.Sugar().Errorf("CouchDB.UpdateDocWithRev: Failed to create HTTP request: %s", err.Error())
		return err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)
	req.Header.Set("if-match", rev)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.UpdateDocWithRev: Failed to make HTTP request: %s", err.Error())
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if isSuccessCode(resp.StatusCode) {
		return nil
	}
	log.Sugar().Infof("CouchDB.UpdateDocWithRev: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar().Errorf("CouchDB.UpdateDocWithRev: Failed to CouchDB response body: %s", err.Error())
		return exception.Unexpected()
	}

	reason, err := couchGetReasonFromBody(bodyBytes)
	if err != nil {
		log.Sugar().Errorf("CouchDB.UpdateDocWithRev: Failed to obtain CouchDB error reason from body: %s", err.Error())
		return exception.Unexpected()
	}

	return errors.New(reason)
}

// GetDoc : Gets a doc from the database.
func (c *couchdb) GetDoc(database string, docID string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/%s/%s", conf.CouchDB.Address, database, docID)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDoc: Failed to create HTTP request: %s", err.Error())
		return nil, err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDoc: Failed to make HTTP request: %s", err.Error())
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDoc: Failed to CouchDB response body: %s", err.Error())
		return nil, exception.Unexpected()
	}

	if isSuccessCode(resp.StatusCode) {
		return bodyBytes, nil
	}
	log.Sugar().Infof("CouchDB.GetDoc: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)

	reason, err := couchGetReasonFromBody(bodyBytes)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDoc: Failed to obtain CouchDB error reason from body: %s", err.Error())
		return nil, exception.Unexpected()
	}

	return nil, errors.New(reason)
}

// HeadDoc : Returns a header map containing minimal information about the document.
func (c *couchdb) HeadDoc(database string, docID string) (http.Header, error) {
	endpoint := fmt.Sprintf("%s/%s/%s", conf.CouchDB.Address, database, docID)

	req, err := http.NewRequest(http.MethodHead, endpoint, nil)
	if err != nil {
		log.Sugar().Errorf("CouchDB.HeadDoc: Failed to create HTTP request: %s", err.Error())
		return nil, err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.HeadDoc: Failed to make HTTP request: %s", err.Error())
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if isSuccessCode(resp.StatusCode) {
		return resp.Header, nil
	}
	log.Sugar().Infof("CouchDB.HeadDoc: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New(constants.CouchMissingReason)
	default:
		return nil, exception.Unexpected()
	}
}

// GetDocRev : Returns the current rev string for a document.
func (c *couchdb) GetDocRev(database string, docID string) (string, error) {
	headers, err := c.HeadDoc(database, docID)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDocRev: Failed to HEAD doc: %s", err.Error())
		return "", err
	}

	etag := headers.Get("etag")
	if etag == "" {
		log.Sugar().Errorf("CouchDB.GetDocRev: Failed to obtain document revision for updation.")
		return "", exception.Unexpected()
	}

	// Removing double quotes from rev.
	return etag[1 : len(etag)-1], nil
}

// GetDocsByView : Gets docs using a view.
func (c *couchdb) GetDocsByView(database string, designName string, viewName string, query url.Values) (int, int, []interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/_design/%s/_view/%s", conf.CouchDB.Address, database, designName, viewName)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDocsByView: Failed to create HTTP request: %s", err.Error())
		return 0, 0, nil, err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)
	req.URL.RawQuery = query.Encode()

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDocsByView: Failed to make HTTP request: %s", err.Error())
		return 0, 0, nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar().Errorf("CouchDB.GetDocsByView: Failed to CouchDB response body: %s", err.Error())
		return 0, 0, nil, exception.Unexpected()
	}

	if !isSuccessCode(resp.StatusCode) {
		log.Sugar().Infof("CouchDB.GetDocsByView: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)
		reason, err := couchGetReasonFromBody(bodyBytes)
		if err != nil {
			log.Sugar().Errorf("CouchDB.GetDocsByView: Failed to obtain CouchDB error reason from body: %s", err.Error())
			return 0, 0, nil, exception.Unexpected()
		}
		return 0, 0, nil, errors.New(reason)
	}

	return couchDecodeViewResult(bodyBytes)
}

// DeleteDoc : Deletes the doc with the given docID.
func (c *couchdb) DeleteDoc(database string, docID string) error {
	rev, err := c.GetDocRev(database, docID)
	if err != nil {
		log.Sugar().Errorf("CouchDB.DeleteDoc: Failed to get document revision for deletion: %s", err.Error())
		return err
	}
	return c.DeleteDocWithRev(database, docID, rev)
}

// DeleteDocWithRev : Deletes the doc with the given docID using the provided rev.
func (c *couchdb) DeleteDocWithRev(database string, docID string, rev string) error {
	endpoint := fmt.Sprintf("%s/%s/%s", conf.CouchDB.Address, database, docID)

	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		log.Sugar().Errorf("CouchDB.DeleteDocWithRev: Failed to create HTTP request: %s", err.Error())
		return err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)
	req.Header.Set("if-match", rev)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.DeleteDocWithRev: Failed to make HTTP request: %s", err.Error())
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if isSuccessCode(resp.StatusCode) {
		return nil
	}
	log.Sugar().Infof("CouchDB.DeleteDocWithRev: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar().Errorf("CouchDB.DeleteDocWithRev: Failed to CouchDB response body: %s", err.Error())
		return exception.Unexpected()
	}

	reason, err := couchGetReasonFromBody(bodyBytes)
	if err != nil {
		log.Sugar().Errorf("CouchDB.DeleteDocWithRev: Failed to obtain CouchDB error reason from body: %s", err.Error())
		return exception.Unexpected()
	}

	return errors.New(reason)
}

// CreateDesign : Creates a design doc in the database.
func (c *couchdb) CreateDesign(database string, designName string, doc []byte) error {
	endpoint := fmt.Sprintf("%s/%s/_design/%s", conf.CouchDB.Address, database, designName)

	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(doc))
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDesign: Failed to create HTTP request: %s", err.Error())
		return err
	}
	req.SetBasicAuth(conf.CouchDB.Username, conf.CouchDB.Password)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDesign: Failed to make HTTP request: %s", err.Error())
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if isSuccessCode(resp.StatusCode) {
		return nil
	}
	log.Sugar().Infof("CouchDB.CreateDesign: CouchDB API returned unsuccessful status code: %d", resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDesign: Failed to CouchDB response body: %s", err.Error())
		return exception.Unexpected()
	}

	reason, err := couchGetReasonFromBody(bodyBytes)
	if err != nil {
		log.Sugar().Errorf("CouchDB.CreateDesign: Failed to obtain CouchDB error reason from body: %s", err.Error())
		return exception.Unexpected()
	}

	return errors.New(reason)
}

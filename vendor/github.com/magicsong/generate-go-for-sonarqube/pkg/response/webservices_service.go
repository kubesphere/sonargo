// Get information on the web api supported on this instance.
package response

type WebservicesService struct {
	client *Client
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type WebservicesListResp struct{}
type WebservicesListOption struct {
	IncludeInternals string `url:"include_internals,omitempty"` // Description:"Include web services that are implemented for internal use only. Their forward-compatibility is not assured",ExampleValue:""
}

// List List web services
func (s *WebservicesService) List(opt *WebservicesListOption) (Resp *WebservicesListResp, err error) {
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/list", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ResponseExample> func manually and complete the fields of this struct
type WebservicesResponseExampleResp struct {
	Format  string `json:"format,omitempty"`
	Example string `json:"example,omitempty"`
	Name    string `json:"name,omitempty"`
}
type WebservicesResponseExampleOption struct {
	Action     string `url:"action,omitempty"`     // Description:"Action of the web service",ExampleValue:"search"
	Controller string `url:"controller,omitempty"` // Description:"Controller of the web service",ExampleValue:"api/issues"
}

// Response_example Display web service response example
func (s *WebservicesService) ResponseExample(opt *WebservicesResponseExampleOption) (resp *WebservicesResponseExampleResp, err error) {
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/response_example", opt)
	if err != nil {
		return
	}
	resp = new(WebservicesResponseExampleResp)
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	resp.Name = opt.Action
	return
}

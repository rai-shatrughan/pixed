package agm

type AgentRequest struct {
	EntityId        string `json:"entityId"`
	Name            string `json:"name"`
	SecurityProfile string `json:"securityProfile"`
}

type AgentResponse struct {
	Id   string `json:"id"`
	ETag string `json:"eTag"`
	AgentRequest
}

func CreateAgent() {

}

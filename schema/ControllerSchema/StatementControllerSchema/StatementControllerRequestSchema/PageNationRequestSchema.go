package StatementControllerRequestSchema

type PageNationColumn struct {
	PageSize int64 `form:"page_size"`
	Current  int64 `form:"current"`
}

func (PageNationRequest *PageNationColumn) InitDefaultRequest() {
	PageNationRequest.PageSize = 20
	PageNationRequest.Current = 1
}

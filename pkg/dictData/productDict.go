package dictData

type ProductDict struct {
	Accident       []string //意外
	SeriousIllness []string //重疾
	Annuity        []string //年金
	LifeInsurance  []string //寿险
}

func ProductObj() *ProductDict {
	var productDictObj ProductDict

	productDictObj.SeriousIllness = []string{"长城安康重大疾病保险", "百多臻爱倍至", "弘康哆啦A保重大疾病保险(2.0版)", "长城吉泰人生重大疾病保险"}
	productDictObj.Accident = []string{"弘康弘益安顺两全保险", "百年如意畅行综合意外保障计划"}
	productDictObj.Annuity = []string{"百年百福终身年金保险", "弘康共庆余年年金保险"}

	return &productDictObj
}

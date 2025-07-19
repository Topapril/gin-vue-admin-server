package utils

var (
	IdVerify               = Rules{"ID": []string{NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}}
	AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
	// 打印参数校验
	PrintVerify = Rules{"BusinessName": {NotEmpty()}, "GoodsName": {NotEmpty()}, "GoodsQuantity": {NotEmpty()}, "MealPeriod": {NotEmpty()}, "UserAddress": {NotEmpty()}, "DeliveryDate": {NotEmpty()}, "PrintType": {NotEmpty()}}

	// 消费者餐次数校验
	ConsumerMealCountVerify = Rules{"TotalMealCount": {NotEmpty()}}

	// 消费者状态校验
	ConsumerStatusVerify = Rules{"ConsumerStatus": {NotEmpty()}}

	// 订单参数校验
	OrderVerify = Rules{"ConsumerId": {NotEmpty()}, "ConsumerName": {NotEmpty()}, "ConsumerPhone": {NotEmpty()}, "ConsumerAddress": {NotEmpty()}, "GoodsName": {NotEmpty()}, "GoodsQuantity": {NotEmpty()}, "MealPeriod": {NotEmpty()}, "OrderType": {NotEmpty()}, "DeliveryDate": {NotEmpty()}}

	// 餐品参数校验
	MealVerify = Rules{"LunchName": {NotEmpty()}, "DinnerName": {NotEmpty()}, "businessDate": {NotEmpty()}}

	// 消费者ID校验
	ConsumerIdVerify = Rules{"ConsumerId": {NotEmpty()}}
)

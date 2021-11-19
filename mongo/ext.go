package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

// GetCompanyCollName 获取企业分库表名
func GetCompanyCollName(companyID, coll string) string {
	return companyID + MKCollMiddle + coll
}

const (
	// MKCollMiddle 中间连接名
	MKCollMiddle = "_mk_"
)

// GenStringObjectID 生成mongo objectid
func GenStringObjectID() string {
	return primitive.NewObjectID().Hex()
}

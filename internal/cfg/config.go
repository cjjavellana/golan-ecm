package cfg

type StoreType string

const (
	// StoreTypeMongoDB stores the document metadata and blob in MongoDB
	StoreTypeMongoDB StoreType = "mongodb"

	// StoreTypeMysql uses MySQL as the backing store and stores the documents as a blob
	StoreTypeMysql = "mysql"

	// StoreTypeAWS uses MySQL to store the metadata information and
	// AWS S3 to store the document content
	StoreTypeAWS = "aws"
)

type Config struct {
	StoreType StoreType
	MongoUri  string
}

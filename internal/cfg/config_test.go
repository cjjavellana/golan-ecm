package cfg

import (
	"reflect"
	"testing"
)

func TestParseConfigFromYamlString(t *testing.T) {
	type args struct {
		yamlString string
	}

	testArgs := args{
		yamlString: `
grpcport: 9000
storetype: aws
storeconfig:
  dynamodburi: dynamodburi
  dynamodbuser: dynamodbuser
  dynamodbpassword: dynamodbpassword
  elasticsearchuri: elasticsearchuri
  elasticsearchuser: elasticsearchuser
  elasticsearchpassword: elasticsearchpassword
`,
	}

	expected := AppConfig{
		GrpcPort: 9000,
		StoreType: StoreTypeAWS,
		StoreConfig: map[interface{}]interface{}{
			"dynamodburi":           "dynamodburi",
			"dynamodbuser":          "dynamodbuser",
			"dynamodbpassword":      "dynamodbpassword",
			"elasticsearchuri":      "elasticsearchuri",
			"elasticsearchuser":     "elasticsearchuser",
			"elasticsearchpassword": "elasticsearchpassword",
		},
	}

	tests := []struct {
		name string
		args args
		want AppConfig
	}{
		{
			name: "itShouldReturnAnAppConfig",
			args: testArgs,
			want: expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseConfigFromYamlString(tt.args.yamlString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfigFromYamlFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

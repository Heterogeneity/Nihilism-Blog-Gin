package flag

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"os"
	"server/global"
	"server/model/elasticsearch"
	"server/model/other"
	"server/service"
)

// ElasticsearchImport 从指定的 JSON 文件导入数据到 ES
func ElasticsearchImport(jsonPath string) (int, error) {
	// 读取指定路径的 JSON 文件
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		return 0, err
	}
	// 反序列化 JSON 数据到 ESIndexResponse 结构体
	var response other.ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		return 0, err
	}
	// 创建 Elasticsearch 索引
	esService := service.ServiceGroupApp.EsService
	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return 0, err
	}
	if indexExists {
		if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
			return 0, err
		}
		err = esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
		if err != nil {
			return 0, err
		}

	}
	// 构建批量请求数据
	var request bulk.Request
	for _, data := range response.Data {
		request = append(request, types.OperationContainer{Index: &types.IndexOperation{Id_: data.ID}})
		request = append(request, data.Doc)
	}
	// 使用 Elasticsearch 客户端执行批量操作
	_, err = global.ESClient.Bulk().Request(&request).Index(elasticsearch.ArticleIndex()).Refresh(refresh.True).Do(context.TODO())
	if err != nil {
		return 0, err
	}
	total := len(response.Data)
	// 返回导入的数据总条数
	return total, nil
}

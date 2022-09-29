/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

type AutoscaleOption struct {

	// 오토스케일 가능여부
	Enabled *bool `json:"enabled"`

	// 오토스케일 가능 최대 노드 수
	Max *int32 `json:"max"`

	// 오토스케일 가능 최소 노드 수
	Min *int32 `json:"min"`
}
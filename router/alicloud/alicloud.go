package alicloud

/* Disable Alicloud due to compile issues
var (
	provider = alicloud.Provider().(*schema.Provider)
)

// GetAliCloudFields returns a list of elements the terraform provider for alicloud supports
func GetAliCloudFields() []model.DynamicElement {
	var result []model.DynamicElement

	for rscName, rscAttr := range provider.ResourcesMap {
		fields := libs.ExtractFields(rscAttr)
		if len(fields) != 0 {
			result = append(result, model.DynamicElement{
				Name:   rscName,
				Fields: fields,
			})
		}
	}

	return result
}
*/

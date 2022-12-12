package libvirt

/* Disable libvirt provider, due to issues with open_terraform
var (
	provider = libvirtProvider.Provider().(*schema.Provider)
)

// GetLibvirtFields returns a list of elements the terraform provider for libvirt supports
func GetLibvirtFields() []model.DynamicElement {
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

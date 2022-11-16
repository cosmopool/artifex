package configuration

const (
	FILE_EXTENSION = "artifex.file.extension"

	// DOMAIN_PATH         = "artifex.path.domain"
	// EXTERNAL_PATH       = "artifex.path.external"
	// INFRASTRUCTURE_PATH = "artifex.path.infrastructure"

	USECASE_IMPLEMENTATION_PATH    = "artifex.path.implementation.usecase"
	REPOSITORY_IMPLEMENTATION_PATH = "artifex.path.implementation.repository"
	DATASOURCE_IMPLEMENTATION_PATH = "artifex.path.implementation.datasource"

	USECASE_INTERFACE_PATH    = "artifex.path.interface.usecase"
	REPOSITORY_INTERFACE_PATH = "artifex.path.interface.repository"
	DATASOURCE_INTERFACE_PATH = "artifex.path.interface.datasource"

	USECASE_IMPLEMENTATION    = "artifex.template.implementation.usecase"
	REPOSITORY_IMPLEMENTATION = "artifex.template.implementation.repository"
	DATASOURCE_IMPLEMENTATION = "artifex.template.implementation.datasource"

	USECASE_INTERFACE    = "artifex.template.interface.usecase"
	REPOSITORY_INTERFACE = "artifex.template.interface.repository"
	DATASOURCE_INTERFACE = "artifex.template.interface.datasource"
)

func ConfigList() []string {
	return []string{
		// DOMAIN_PATH,
		// EXTERNAL_PATH,
		// INFRASTRUCTURE_PATH,
		USECASE_IMPLEMENTATION_PATH,
		REPOSITORY_IMPLEMENTATION_PATH,
		DATASOURCE_IMPLEMENTATION_PATH,
		USECASE_IMPLEMENTATION,
		REPOSITORY_IMPLEMENTATION,
		DATASOURCE_IMPLEMENTATION,
		REPOSITORY_INTERFACE,
		USECASE_INTERFACE,
		DATASOURCE_INTERFACE,
		USECASE_INTERFACE_PATH,
		REPOSITORY_INTERFACE_PATH,
		DATASOURCE_INTERFACE_PATH,
	}
}

package core

const (
	// DOMAIN_PATH         = "artifex.config.path.domain"
	// EXTERNAL_PATH       = "artifex.config.path.external"
	// INFRASTRUCTURE_PATH = "artifex.config.path.infrastructure"
	// USECASE_PATH    = "artifex.config.path.usecase"
	// REPOSITORY_PATH = "artifex.config.path.repository"
	// DATASOURCE_PATH = "artifex.config.path.datasource"

	FILE_EXTENSION = "artifex.config.file.extension"

	USECASE_IMPLEMENTATION_PATH    = "artifex.config.implementation.path.usecase"
	REPOSITORY_IMPLEMENTATION_PATH = "artifex.config.implementation.path.repository"
	DATASOURCE_IMPLEMENTATION_PATH = "artifex.config.implementation.path.datasource"

	USECASE_INTERFACE_PATH    = "artifex.config.interface.path.usecase"
	REPOSITORY_INTERFACE_PATH = "artifex.config.interface.path.repository"
	DATASOURCE_INTERFACE_PATH = "artifex.config.interface.path.datasource"

	USECASE_IMPLEMENTATION    = "artifex.config.implementation.filename.usecase"
	REPOSITORY_IMPLEMENTATION = "artifex.config.implementation.filename.repository"
	DATASOURCE_IMPLEMENTATION = "artifex.config.implementation.filename.datasource"

	USECASE_INTERFACE    = "artifex.config.interface.filename.usecase"
	REPOSITORY_INTERFACE = "artifex.config.interface.filename.repository"
	DATASOURCE_INTERFACE = "artifex.config.interface.filename.datasource"
)

func ConfigList() []string {
	return []string{
		// DOMAIN_PATH,
		// EXTERNAL_PATH,
		// INFRASTRUCTURE_PATH,
		// USECASE_PATH,
		// REPOSITORY_PATH,
		// DATASOURCE_PATH,
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

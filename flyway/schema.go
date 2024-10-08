package flyway

type Schema struct {
	Version     int
	Description string
	Script      string
	Sql         string
}

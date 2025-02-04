package cluster

var (
	boltDBShipperSchemaConfigTemplate = `
schema_config:
  configs:
    - from: {{.curPeriodStart}}
      store: boltdb-shipper
      object_store: filesystem
      schema: v11
      index:
        prefix: index_
        period: 24h
`
	additionalBoltDBShipperSchemaConfigTemplate = `
schema_config:
  configs:
    - from: {{.additionalPeriodStart}}
      store: boltdb-shipper
      object_store: store-1
      schema: v11
      index:
        prefix: index_
        period: 24h
`

	tsdbShipperSchemaConfigTemplate = `
schema_config:
  configs:
    - from: {{.curPeriodStart}}
      store: tsdb
      object_store: filesystem
      schema: v11
      index:
        prefix: index_
        period: 24h
`
	additionalTSDBShipperSchemaConfigTemplate = `
schema_config:
  configs:
    - from: {{.additionalPeriodStart}}
      store: tsdb
      object_store: store-1
      schema: v11
      index:
        prefix: index_tsdb_
        period: 24h
`
)

func SchemaWithTSDB(c *Cluster) {
	c.periodCfgs = append(c.periodCfgs, additionalTSDBShipperSchemaConfigTemplate)
}

func SchemaWithBoltDBAndBoltDB(c *Cluster) {
	c.periodCfgs = append(c.periodCfgs, additionalBoltDBShipperSchemaConfigTemplate, boltDBShipperSchemaConfigTemplate)
}

func SchemaWithTSDBAndTSDB(c *Cluster) {
	c.periodCfgs = append(c.periodCfgs, additionalTSDBShipperSchemaConfigTemplate, tsdbShipperSchemaConfigTemplate)
}

func SchemaWithBoltDBAndTSDB(c *Cluster) {
	c.periodCfgs = append(c.periodCfgs, additionalBoltDBShipperSchemaConfigTemplate, tsdbShipperSchemaConfigTemplate)
}

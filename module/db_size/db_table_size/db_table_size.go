package db_table_size

import (
	"database/sql"
	"fmt"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/metricbeat/mb"

	_ "github.com/lib/pq" //switch postgresql driver
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("db_size", "db_table_size", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	oid         int
	tableSchema string
	tableName   string
	rowEstimate int
	totalBytes  int
	indexBytes  int
	toastBytes  int
	tableBytes  int
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The db_size db_table_size metricset is beta.")

	config := struct {
		oid         int    `config:"oid"`
		tableSchema string `config:"tableSchema"`
		tableName   string `config:"tableName"`
		rowEstimate int    `config:"rowEstimate"`
		totalBytes  int    `config:"totalBytes"`
		indexBytes  int    `config:"indexBytes"`
		toastBytes  int    `config:"toastBytes"`
		tableBytes  int    `config:"tableBytes"`
	}{
		oid:         0,
		tableSchema: "",
		tableName:   "",
		rowEstimate: 0,
		totalBytes:  0,
		indexBytes:  0,
		toastBytes:  0,
		tableBytes:  0,
	}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		oid:           config.oid,
		tableSchema:   config.tableSchema,
		tableName:     config.tableName,
		rowEstimate:   config.rowEstimate,
		totalBytes:    config.totalBytes,
		indexBytes:    config.indexBytes,
		toastBytes:    config.toastBytes,
		tableBytes:    config.tableBytes,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	report.Event(mb.Event{
		MetricSetFields: common.MapStr{
			"oid":         m.oid,
			"tableSchema": m.tableSchema,
			"tableName":   m.tableName,
			"rowEstimate": m.rowEstimate,
			"totalBytes":  m.totalBytes,
			"indexBytes":  m.indexBytes,
			"toastBytes":  m.toastBytes,
			"tableBytes":  m.tableBytes,
		},
	})

	const connectionStr string = "user=postgres dbname=california sslmode=disable" //change connectionStr!!!!!!!!!
	db, err := sql.Open("postgres", connectionStr)                                 //string
	if err != nil {
		panic(err)
	}
	defer db.Close()

	const queryStr string = `select
								*,
								total_bytes-index_bytes-coalesce(toast_bytes,0) as table_bytes
							from (
								select
									c.oid,
									nspname as table_schema,
									relname as table_name,
									c.reltuples as row_estimate,
									pg_total_relation_size(c.oid) as total_bytes,
									pg_indexes_size(c.oid) as index_bytes,
									pg_total_relation_size(reltoastrelid) as toast_bytes
								from
									pg_class c
								left join pg_namespace n on n.oid = c.relnamespace
								where relkind = 'r'
								) a
							order by total_bytes desc;`
	rows, err := db.Query(queryStr)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	MetricSets := []MetricSet{}

	for rows.Next() {
		m := MetricSet{}
		rows.Scan(
			&m.oid,
			&m.tableSchema,
			&m.tableName,
			&m.rowEstimate,
			&m.totalBytes,
			&m.indexBytes,
			&m.toastBytes,
			&m.tableBytes,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		MetricSets = append(MetricSets, m)
	}
	return nil
}

// package db_table_size

// import (
// 	"github.com/elastic/beats/libbeat/common"
// 	"github.com/elastic/beats/libbeat/common/cfgwarn"
// 	"github.com/elastic/beats/metricbeat/mb"
// )

// // init registers the MetricSet with the central registry as soon as the program
// // starts. The New function will be called later to instantiate an instance of
// // the MetricSet for each host defined in the module's configuration. After the
// // MetricSet has been created then Fetch will begin to be called periodically.
// func init() {
// 	mb.Registry.MustAddMetricSet("db_size", "db_table_size", New)
// }

// // MetricSet holds any configuration or state information. It must implement
// // the mb.MetricSet interface. And this is best achieved by embedding
// // mb.BaseMetricSet because it implements all of the required mb.MetricSet
// // interface methods except for Fetch.
// type MetricSet struct {
// 	mb.BaseMetricSet
// 	counter int
// }

// // New creates a new instance of the MetricSet. New is responsible for unpacking
// // any MetricSet specific configuration options if there are any.
// func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
// 	cfgwarn.Beta("The db_size db_table_size metricset is beta.")

// 	config := struct{}{}
// 	if err := base.Module().UnpackConfig(&config); err != nil {
// 		return nil, err
// 	}

// 	return &MetricSet{
// 		BaseMetricSet: base,
// 		counter:       1,
// 	}, nil
// }

// // Fetch methods implements the data gathering and data conversion to the right
// // format. It publishes the event which is then forwarded to the output. In case
// // of an error set the Error field of mb.Event or simply call report.Error().
// func (m *MetricSet) Fetch(report mb.ReporterV2) error {
// 	report.Event(mb.Event{
// 		MetricSetFields: common.MapStr{
// 			"counter": m.counter,
// 		},
// 	})
// 	m.counter++

// 	return nil
// }

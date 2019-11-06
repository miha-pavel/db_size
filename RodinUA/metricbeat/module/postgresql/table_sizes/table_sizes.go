package table_sizes

import (
	"context"

	"github.com/pkg/errors"

	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/postgresql"

	// Register postgresql database/sql driver
	_ "github.com/lib/pq"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("postgresql", "db_sizes", New,
	    mb.WithHostParser(postgresql.ParseURL),
    	mb.DefaultMetricSet(),
    )
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	*postgresql.MetricSet
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	ms, err := postgresql.NewMetricSet(base)
	if err != nil {
		return nil, err
	}
	return &MetricSet{MetricSet: ms}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().

func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	ctx := context.Background()
	const queryStr string = `
        SELECT
            d.datname AS Name,
            pg_catalog.pg_get_userbyid(d.datdba) AS Owner,
            CASE WHEN pg_catalog.has_database_privilege(d.datname, 'CONNECT')
                    THEN pg_catalog.pg_database_size(d.datname)
                    ELSE null
            END AS SIZE
        FROM pg_catalog.pg_database d;`
	results, err := m.QueryStats(ctx, queryStr)
	if err != nil {
		return errors.Wrap(err, "error in QueryStats")
	}

	for _, result := range results {
		data, _ := schema.Apply(result)
		reporter.Event(mb.Event{
			MetricSetFields: data,
		})
	}

	return nil
}

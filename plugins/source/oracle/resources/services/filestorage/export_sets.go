package filestorage

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/filestorage"
)

func ExportSets() *schema.Table {
	return &schema.Table{
		Name:      "oracle_filestorage_export_sets",
		Resolver:  fetchExportSets,
		Multiplex: client.AvailibilityDomainCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&filestorage.ExportSetSummary{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:       "region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOracleRegion,
				PrimaryKey: true,
			},
			{
				Name:       "compartment_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveCompartmentId,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}

package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DiskEncryptionSets() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_disk_encryption_sets",
		Resolver:    fetchDiskEncryptionSets,
		Description: "https://learn.microsoft.com/en-us/rest/api/compute/disk-encryption-sets/list?tabs=HTTP#diskencryptionset",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_compute_disk_encryption_sets", client.Namespacemicrosoft_compute),
		Transform:   transformers.TransformWithStruct(&armcompute.DiskEncryptionSet{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchDiskEncryptionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewDiskEncryptionSetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

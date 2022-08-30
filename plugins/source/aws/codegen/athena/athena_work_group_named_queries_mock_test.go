// Code generated by codegen; DO NOT EDIT.

package athena

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

func buildAthenaWorkGroupNamedQueries(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAthenaClient(ctrl)

	var item string
	if err := faker.FakeData(&item); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListNamedQueries(
		gomock.Any(),
		&athena.ListNamedQueriesInput{},
		gomock.Any(),
	).Return(
		&athena.ListNamedQueriesOutput{
			NamedQueryIds: []string{item},
		},
		nil,
	)

	var detail types.NamedQuery
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}

	detail.NamedQueryId = &item

	mock.EXPECT().GetNamedQuery(
		gomock.Any(),
		&athena.GetNamedQueryInput{

			NamedQueryId: &item,
		},
		gomock.Any(),
	).Return(
		&athena.GetNamedQueryOutput{
			NamedQuery: &detail,
		},
		nil,
	)

	return client.Services{
		Athena: mock,
	}
}

func TestAthenaWorkGroupNamedQueries(t *testing.T) {
	client.AwsMockTestHelper(t, AthenaWorkGroupNamedQueries(), buildAthenaWorkGroupNamedQueries, client.TestOptions{})
}

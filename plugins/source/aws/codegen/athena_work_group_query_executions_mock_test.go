// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

func buildAthenaWorkGroupQueryExecutions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAthenaClient(ctrl)

	var item string

	if err := faker.FakeData(&item); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListQueryExecutions(
		gomock.Any(),
		&athena.ListQueryExecutionsInput{},
		gomock.Any(),
	).Return(
		&athena.ListQueryExecutionsOutput{

			QueryExecutionIds: []string{item},
		},
		nil,
	)

	var detail types.QueryExecution

	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}

	detail.QueryExecutionId = &item

	mock.EXPECT().GetQueryExecution(
		gomock.Any(),
		&athena.GetQueryExecutionInput{

			QueryExecutionId: &item,
		},
		gomock.Any(),
	).Return(
		&athena.GetQueryExecutionOutput{
			QueryExecution: &detail,
		},
		nil,
	)

	return client.Services{
		Athena: mock,
	}
}

func TestAthenaWorkGroupQueryExecutions(t *testing.T) {
	client.AwsMockTestHelper(t, AthenaWorkGroupQueryExecutions(), buildAthenaWorkGroupQueryExecutions, client.TestOptions{})
}

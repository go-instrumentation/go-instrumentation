package db

import (
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/go-instrumentation/go-instrumentation/test/test_data"
	"github.com/go-instrumentation/go-instrumentation/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListSelfBuild(t *testing.T) {
	assert.NoError(t, test_db.InitSelfBuild(DB))
	selfBuildList, err := ListSelfBuild()
	assert.NoError(t, err)
	assert.Equal(t, test_data.SelfBuilds, selfBuildList)
}

func TestCreateSelfBuild(t *testing.T) {
	assert.NoError(t, test_db.MakeTableEmpty(DB, &model.SelfBuild{}))
	assert.NoError(t, CreateSelfBuild(test_data.SelfBuild1))
	selfBuildList, err := ListSelfBuild()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(selfBuildList))
	assert.Equal(t, test_data.SelfBuild1.ID, selfBuildList[0].ID)
}

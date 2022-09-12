package models

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/hexcraft-biz/model"
	"github.com/jmoiron/sqlx"
)

//================================================================
// Data Struct
//================================================================
type EntityEndpointTopicRel struct {
	*model.Prototype `dive:""`
	EndpointId       uuid.UUID `db:"endpoint_id"`
	TopicId          uuid.UUID `db:"topic_id"`
}

//================================================================
// View Table: endpoint_topics
//================================================================
type EndpointTopicRel struct {
	ID           uuid.UUID `db:"id" json:"id"`
	EndpointID   uuid.UUID `db:"endpoint_id" json:"endpointId"`
	EndpointPath string    `db:"endpoint_path" json:"endpointPath"`
	TopicID      uuid.UUID `db:"topic_id" json:"topicId"`
	TopicName    string    `db:"topic_name" json:"topicName"`
	CreatedAt    string    `db:"ctime" json:"createdAt"`
	UpdatedAt    string    `db:"mtime" json:"updatedAt"`
}

//================================================================
// Engine
//================================================================
type EndpointTopicRelsTableEngine struct {
	*model.Engine
	viewName string
}

func NewEndpointTopicRelsTableEngine(db *sqlx.DB) *EndpointTopicRelsTableEngine {
	return &EndpointTopicRelsTableEngine{
		Engine:   model.NewEngine(db, "endpoint_topic_rels"),
		viewName: "view_endpoint_topic_rels",
	}
}

func (e *EndpointTopicRelsTableEngine) Insert(EndpointId, TopicId string) (*EndpointTopicRel, error) {
	eid, _ := uuid.Parse(EndpointId)
	tid, _ := uuid.Parse(TopicId)

	etr := &EntityEndpointTopicRel{
		Prototype:  model.NewPrototype(),
		EndpointId: eid,
		TopicId:    tid,
	}

	if _, err := e.Engine.Insert(etr); err != nil {
		return nil, err
	}

	return e.GetByID(etr.ID.String())
}

func (e *EndpointTopicRelsTableEngine) GetByID(id string) (*EndpointTopicRel, error) {
	row := EndpointTopicRel{}
	q := `SELECT * FROM ` + e.viewName + ` WHERE id = UUID_TO_BIN(?);`
	if err := e.Engine.Get(&row, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &row, nil
}

type EtrListQuery struct {
	EndpointID   string
	EndpointPath string
	TopicID      string
	TopicName    string
}

func (e *EndpointTopicRelsTableEngine) List(listQuery EtrListQuery, pg *model.Pagination) ([]*EndpointTopicRel, error) {
	rows := []*EndpointTopicRel{}

	if pg == nil {
		pg = model.NewDefaultPagination()
	}

	args := []interface{}{}
	andSqlSlice := []string{}

	if listQuery.EndpointID != "" {
		andSqlSlice = append(andSqlSlice, "`endpoint_id` = UUID_TO_BIN(?)")
		args = append(args, listQuery.EndpointID)
	}
	if listQuery.EndpointPath != "" {
		andSqlSlice = append(andSqlSlice, "`endpoint_path` = ?")
		args = append(args, listQuery.EndpointPath)
	}
	if listQuery.TopicID != "" {
		andSqlSlice = append(andSqlSlice, "`topic_id` = UUID_TO_BIN(?)")
		args = append(args, listQuery.TopicID)
	}
	if listQuery.TopicName != "" {
		andSqlSlice = append(andSqlSlice, "`topic_name` = ?")
		args = append(args, listQuery.TopicName)
	}

	andSQL := "1"
	if len(andSqlSlice) >= 1 {
		andSQL = strings.Join(andSqlSlice[:], " AND ")
	}

	q := `SELECT * FROM ` + e.viewName + ` WHERE ` + andSQL + pg.ToString() + `;`

	if err := e.Engine.Select(&rows, q, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return rows, nil
}

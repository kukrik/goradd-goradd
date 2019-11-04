package node

// Code generated by goradd. DO NOT EDIT.

import (
	"encoding/gob"

	"github.com/goradd/goradd/pkg/orm/query"
)

type projectNode struct {
	query.ReferenceNodeI
}

func Project() *projectNode {
	n := projectNode{
		query.NewTableNode("goradd", "project", "Project"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

func (n *projectNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Num())
	nodes = append(nodes, n.ProjectStatusTypeID())
	nodes = append(nodes, n.ManagerID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.Description())
	nodes = append(nodes, n.StartDate())
	nodes = append(nodes, n.EndDate())
	nodes = append(nodes, n.Budget())
	nodes = append(nodes, n.Spent())
	return nodes
}
func (n *projectNode) PrimaryKeyNode_() *query.ColumnNode {
	return n.ID()
}
func (n *projectNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}
func (n *projectNode) Copy_() query.NodeI {
	return &projectNode{query.CopyNode(n.ReferenceNodeI)}
}

// ID represents the id column in the database.
func (n *projectNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Num represents the num column in the database.
func (n *projectNode) Num() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"num",
		"Num",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ProjectStatusTypeID represents the project_status_type_id column in the database.
func (n *projectNode) ProjectStatusTypeID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"project_status_type_id",
		"ProjectStatusTypeID",
		query.ColTypeUnsigned,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ProjectStatusType represents the the link to the ProjectStatusType object.
func (n *projectNode) ProjectStatusType() *projectStatusTypeNode {
	cn := &projectStatusTypeNode{
		query.NewReferenceNode(
			"goradd",
			"project",
			"project_status_type_id",
			"ProjectStatusTypeID",
			"ProjectStatusType",
			"project_status_type",
			"id",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// ManagerID represents the manager_id column in the database.
func (n *projectNode) ManagerID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"manager_id",
		"ManagerID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Manager represents the the link to the Manager object.
func (n *projectNode) Manager() *personNode {
	cn := &personNode{
		query.NewReferenceNode(
			"goradd",
			"project",
			"manager_id",
			"ManagerID",
			"Manager",
			"person",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *projectNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Description represents the description column in the database.
func (n *projectNode) Description() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"description",
		"Description",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// StartDate represents the start_date column in the database.
func (n *projectNode) StartDate() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"start_date",
		"StartDate",
		query.ColTypeDateTime,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// EndDate represents the end_date column in the database.
func (n *projectNode) EndDate() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"end_date",
		"EndDate",
		query.ColTypeDateTime,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Budget represents the budget column in the database.
func (n *projectNode) Budget() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"budget",
		"Budget",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Spent represents the spent column in the database.
func (n *projectNode) Spent() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"spent",
		"Spent",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ChildrenAsParent represents the many-to-many relationship formed by the related_project_assn table.
func (n *projectNode) ChildrenAsParent() *projectNode {
	cn := &projectNode{
		query.NewManyManyNode(
			"goradd",
			"related_project_assn",
			"parent_id",
			"ChildrenAsParent",
			"project",
			"child_id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// ParentsAsChild represents the many-to-many relationship formed by the related_project_assn table.
func (n *projectNode) ParentsAsChild() *projectNode {
	cn := &projectNode{
		query.NewManyManyNode(
			"goradd",
			"related_project_assn",
			"child_id",
			"ParentsAsChild",
			"project",
			"parent_id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// TeamMembers represents the many-to-many relationship formed by the team_member_project_assn table.
func (n *projectNode) TeamMembers() *personNode {
	cn := &personNode{
		query.NewManyManyNode(
			"goradd",
			"team_member_project_assn",
			"project_id",
			"TeamMembers",
			"person",
			"team_member_id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

// Milestones represents the many-to-one relationship formed by the reverse reference from the
// id column in the project table.
func (n *projectNode) Milestones() *milestoneNode {

	cn := &milestoneNode{
		query.NewReverseReferenceNode(
			"goradd",
			"project",
			"id",
			"Milestones",
			"milestone",
			"project_id",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}

func init() {
	gob.RegisterName("projectNode2", &projectNode{})
}
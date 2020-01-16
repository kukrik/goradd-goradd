package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/goradd/goradd/pkg/orm/broadcast"
	"github.com/goradd/goradd/pkg/orm/db"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"
	"github.com/goradd/goradd/pkg/stringmap"
	"github.com/goradd/goradd/web/examples/gen/goradd/model/node"

	//"./node"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// milestoneBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type milestoneBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	projectID        string
	projectIDIsValid bool
	projectIDIsDirty bool
	oProject         *Project

	name        string
	nameIsValid bool
	nameIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool
}

const (
	MilestoneIDDefault        = ""
	MilestoneProjectIDDefault = ""
	MilestoneNameDefault      = ""
)

const (
	Milestone_ID        = `ID`
	Milestone_ProjectID = `ProjectID`
	Milestone_Project   = `Project`
	Milestone_Name      = `Name`
)

// Initialize or re-initialize a Milestone database object to default values.
func (o *milestoneBase) Initialize() {

	o.id = ""
	o.idIsValid = false
	o.idIsDirty = false

	o.projectID = ""
	o.projectIDIsValid = false
	o.projectIDIsDirty = false

	o.name = ""
	o.nameIsValid = false
	o.nameIsDirty = false

	o._restored = false
}

func (o *milestoneBase) PrimaryKey() string {
	return o.id
}

// ID returns the loaded value of ID.
func (o *milestoneBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

func (o *milestoneBase) ProjectID() string {
	if o._restored && !o.projectIDIsValid {
		panic("projectID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.projectID
}

// ProjectIDIsValid returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) ProjectIDIsValid() bool {
	return o.projectIDIsValid
}

// Project returns the current value of the loaded Project, and nil if its not loaded.
func (o *milestoneBase) Project() *Project {
	return o.oProject
}

// LoadProject returns the related Project. If it is not already loaded,
// it will attempt to load it first.
func (o *milestoneBase) LoadProject(ctx context.Context) *Project {
	if !o.projectIDIsValid {
		return nil
	}

	if o.oProject == nil {
		// Load and cache
		o.oProject = LoadProject(ctx, o.ProjectID())
	}
	return o.oProject
}

// SetProjectID sets the value of ProjectID in the object, to be saved later using the Save() function.
func (o *milestoneBase) SetProjectID(v string) {
	o.projectIDIsValid = true
	if o.projectID != v || !o._restored {
		o.projectID = v
		o.projectIDIsDirty = true
		o.oProject = nil
	}

}

// SetProject sets the value of Project in the object, to be saved later using the Save() function.
func (o *milestoneBase) SetProject(v *Project) {
	if v == nil {
		panic("Cannot set Project to a null value.")
	} else {
		o.oProject = v
		o.projectIDIsValid = true
		if o.projectID != v.PrimaryKey() {
			o.projectID = v.PrimaryKey()
			o.projectIDIsDirty = true
		}
	}
}

func (o *milestoneBase) Name() string {
	if o._restored && !o.nameIsValid {
		panic("name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsValid returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) NameIsValid() bool {
	return o.nameIsValid
}

// SetName sets the value of Name in the object, to be saved later using the Save() function.
func (o *milestoneBase) SetName(v string) {
	o.nameIsValid = true
	if o.name != v || !o._restored {
		o.name = v
		o.nameIsDirty = true
	}

}

// GetAlias returns the alias for the given key.
func (o *milestoneBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// Load returns a Milestone from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
func LoadMilestone(ctx context.Context, primaryKey string, joinOrSelectNodes ...query.NodeI) *Milestone {
	return queryMilestones(ctx).Where(Equal(node.Milestone().ID(), primaryKey)).joinOrSelect(joinOrSelectNodes...).Get(ctx)
}

// The MilestonesBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type MilestonesBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newMilestoneBuilder() *MilestonesBuilder {
	b := &MilestonesBuilder{
		base: db.GetDatabase("goradd").
			NewBuilder(),
	}
	return b.Join(node.Milestone())
}

// Load terminates the query builder, performs the query, and returns a slice of Milestone objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *MilestonesBuilder) Load(ctx context.Context) (milestoneSlice []*Milestone) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Milestone)
		o.load(item, o, nil, "")
		milestoneSlice = append(milestoneSlice, o)
	}
	return milestoneSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *MilestonesBuilder) LoadI(ctx context.Context) (milestoneSlice []interface{}) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Milestone)
		o.load(item, o, nil, "")
		milestoneSlice = append(milestoneSlice, o)
	}
	return milestoneSlice
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *MilestonesBuilder) Get(ctx context.Context) *Milestone {
	results := b.Load(ctx)
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *MilestonesBuilder) Expand(n query.NodeI) *MilestonesBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *MilestonesBuilder) Join(n query.NodeI, conditions ...query.NodeI) *MilestonesBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// Where adds a condition to filter what gets selected.
func (b *MilestonesBuilder) Where(c query.NodeI) *MilestonesBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *MilestonesBuilder) OrderBy(nodes ...query.NodeI) *MilestonesBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *MilestonesBuilder) Limit(maxRowCount int, offset int) *MilestonesBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *MilestonesBuilder) Select(nodes ...query.NodeI) *MilestonesBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *MilestonesBuilder) Alias(name string, n query.NodeI) *MilestonesBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *MilestonesBuilder) Distinct() *MilestonesBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *MilestonesBuilder) GroupBy(nodes ...query.NodeI) *MilestonesBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *MilestonesBuilder) Having(node query.NodeI) *MilestonesBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
func (b *MilestonesBuilder) Count(ctx context.Context, distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(ctx, distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *MilestonesBuilder) Delete(ctx context.Context) {
	b.base.Delete(ctx)
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *MilestonesBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *MilestonesBuilder) joinOrSelect(nodes ...query.NodeI) *MilestonesBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.base.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

func CountMilestoneByID(ctx context.Context, id string) uint {
	return queryMilestones(ctx).Where(Equal(node.Milestone().ID(), id)).Count(ctx, false)
}

func CountMilestoneByProjectID(ctx context.Context, projectID string) uint {
	return queryMilestones(ctx).Where(Equal(node.Milestone().ProjectID(), projectID)).Count(ctx, false)
}

func CountMilestoneByName(ctx context.Context, name string) uint {
	return queryMilestones(ctx).Where(Equal(node.Milestone().Name(), name)).Count(ctx, false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *milestoneBase) load(m map[string]interface{}, objThis *Milestone, objParent interface{}, parentKey string) {
	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true
			o.idIsDirty = false
		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["project_id"]; ok && v != nil {
		if o.projectID, ok = v.(string); ok {
			o.projectIDIsValid = true
			o.projectIDIsDirty = false
		} else {
			panic("Wrong type found for project_id.")
		}
	} else {
		o.projectIDIsValid = false
		o.projectID = ""
	}

	if v, ok := m["Project"]; ok {
		if oProject, ok2 := v.(map[string]interface{}); ok2 {
			o.oProject = new(Project)
			o.oProject.load(oProject, o.oProject, objThis, "Milestones")
			o.projectIDIsValid = true
			o.projectIDIsDirty = false
		} else {
			panic("Wrong type found for oProject object.")
		}
	} else {
		o.oProject = nil
	}

	if v, ok := m["name"]; ok && v != nil {
		if o.name, ok = v.(string); ok {
			o.nameIsValid = true
			o.nameIsDirty = false
		} else {
			panic("Wrong type found for name.")
		}
	} else {
		o.nameIsValid = false
		o.name = ""
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *milestoneBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *milestoneBase) update(ctx context.Context) {
	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if o.oProject != nil {
			o.oProject.Save(ctx)
			id := o.oProject.PrimaryKey()
			o.SetProjectID(id)
		}

		if !o._restored {
			panic("Cannot update a record that was not originally read from the database.")
		}

		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "milestone", modifiedFields, "id", fmt.Sprint(o.id))
		}

	}) // transaction
	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "milestone", fmt.Sprint(o.id), stringmap.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the item into the database. Related items will be saved.
func (o *milestoneBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {
		if o.oProject != nil {
			o.oProject.Save(ctx)
			o.SetProject(o.oProject)
		}

		if !o.projectIDIsValid {
			panic("a value for ProjectID is required, and there is no default value. Call SetProjectID() before inserting the record.")
		}

		if !o.nameIsValid {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}
		m := o.getValidFields()

		id := d.Insert(ctx, "milestone", m)
		o.id = id

	}) // transaction
	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "milestone", fmt.Sprint(o.id))
}

func (o *milestoneBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.projectIDIsDirty {
		fields["project_id"] = o.projectID
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	return
}

func (o *milestoneBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.projectIDIsValid {
		fields["project_id"] = o.projectID
	}
	if o.nameIsValid {
		fields["name"] = o.name
	}
	return
}

// Delete deletes the associated record from the database.
func (o *milestoneBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "milestone", "id", o.id)
	broadcast.Delete(ctx, "goradd", "milestone", fmt.Sprint(o.id))
}

// deleteMilestone deletes the associated record from the database.
func deleteMilestone(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "milestone", "id", pk)
	broadcast.Delete(ctx, "goradd", "milestone", pk)
}

func (o *milestoneBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.projectIDIsDirty = false
	o.nameIsDirty = false

}

func (o *milestoneBase) IsDirty() bool {
	return o.idIsDirty ||
		o.projectIDIsDirty || (o.oProject != nil && o.oProject.IsDirty()) ||
		o.nameIsDirty
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *milestoneBase) Get(key string) interface{} {

	switch key {
	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "ProjectID":
		if !o.projectIDIsValid {
			return nil
		}
		return o.projectID

	case "Project":
		return o.Project()

	case "Name":
		if !o.nameIsValid {
			return nil
		}
		return o.name

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *milestoneBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.projectID); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.projectIDIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.projectIDIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.oProject); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.name); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.nameIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.nameIsDirty); err != nil {
		return nil, err
	}

	if o._aliases == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o._aliases); err != nil {
			return nil, err
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o *milestoneBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&o.id); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.projectID); err != nil {
		return
	}
	if err = dec.Decode(&o.projectIDIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.projectIDIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.oProject); err != nil {
		return
	}
	if err = dec.Decode(&o.name); err != nil {
		return
	}
	if err = dec.Decode(&o.nameIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return
	}

	var hasAliases bool
	if err = dec.Decode(&hasAliases); err != nil {
		return
	}
	if hasAliases {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object.
func (o *milestoneBase) MarshalJSON() (data []byte, err error) {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.projectIDIsValid {
		v["projectID"] = o.projectID
	}

	if val := o.Project(); val != nil {
		v["project"] = val
	}
	if o.nameIsValid {
		v["name"] = o.name
	}

	return json.Marshal(v)
}

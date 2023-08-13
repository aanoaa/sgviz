// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSgroups(t *testing.T) {
	t.Parallel()

	query := Sgroups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSgroupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSgroupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sgroups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSgroupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SgroupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSgroupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SgroupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Sgroup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SgroupExists to return true, but got false.")
	}
}

func testSgroupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	sgroupFound, err := FindSgroup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if sgroupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSgroupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Sgroups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSgroupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Sgroups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSgroupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	sgroupOne := &Sgroup{}
	sgroupTwo := &Sgroup{}
	if err = randomize.Struct(seed, sgroupOne, sgroupDBTypes, false, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}
	if err = randomize.Struct(seed, sgroupTwo, sgroupDBTypes, false, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sgroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sgroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sgroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSgroupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	sgroupOne := &Sgroup{}
	sgroupTwo := &Sgroup{}
	if err = randomize.Struct(seed, sgroupOne, sgroupDBTypes, false, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}
	if err = randomize.Struct(seed, sgroupTwo, sgroupDBTypes, false, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sgroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sgroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func sgroupBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func sgroupAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sgroup) error {
	*o = Sgroup{}
	return nil
}

func testSgroupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Sgroup{}
	o := &Sgroup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, sgroupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Sgroup object: %s", err)
	}

	AddSgroupHook(boil.BeforeInsertHook, sgroupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	sgroupBeforeInsertHooks = []SgroupHook{}

	AddSgroupHook(boil.AfterInsertHook, sgroupAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	sgroupAfterInsertHooks = []SgroupHook{}

	AddSgroupHook(boil.AfterSelectHook, sgroupAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	sgroupAfterSelectHooks = []SgroupHook{}

	AddSgroupHook(boil.BeforeUpdateHook, sgroupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	sgroupBeforeUpdateHooks = []SgroupHook{}

	AddSgroupHook(boil.AfterUpdateHook, sgroupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	sgroupAfterUpdateHooks = []SgroupHook{}

	AddSgroupHook(boil.BeforeDeleteHook, sgroupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	sgroupBeforeDeleteHooks = []SgroupHook{}

	AddSgroupHook(boil.AfterDeleteHook, sgroupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	sgroupAfterDeleteHooks = []SgroupHook{}

	AddSgroupHook(boil.BeforeUpsertHook, sgroupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	sgroupBeforeUpsertHooks = []SgroupHook{}

	AddSgroupHook(boil.AfterUpsertHook, sgroupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	sgroupAfterUpsertHooks = []SgroupHook{}
}

func testSgroupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSgroupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(sgroupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSgroupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSgroupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SgroupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSgroupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sgroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	sgroupDBTypes = map[string]string{`ID`: `INTEGER`, `Name`: `TEXT`, `Desc`: `TEXT`}
	_             = bytes.MinRead
)

func testSgroupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(sgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(sgroupAllColumns) == len(sgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSgroupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(sgroupAllColumns) == len(sgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sgroup{}
	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sgroupDBTypes, true, sgroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(sgroupAllColumns, sgroupPrimaryKeyColumns) {
		fields = sgroupAllColumns
	} else {
		fields = strmangle.SetComplement(
			sgroupAllColumns,
			sgroupPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, sgroupGeneratedColumns)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SgroupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSgroupsUpsert(t *testing.T) {
	t.Parallel()
	if len(sgroupAllColumns) == len(sgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Sgroup{}
	if err = randomize.Struct(seed, &o, sgroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sgroup: %s", err)
	}

	count, err := Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, sgroupDBTypes, false, sgroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sgroup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sgroup: %s", err)
	}

	count, err = Sgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"time"

	"github.com/kolide/fleet/server/kolide"
)

var _ kolide.LabelStore = (*LabelStore)(nil)

type NewLabelFunc func(Label *kolide.Label, opts ...kolide.OptionalArg) (*kolide.Label, error)

type DeleteLabelFunc func(lid uint) error

type LabelFunc func(lid uint) (*kolide.Label, error)

type ListLabelsFunc func(opt kolide.ListOptions) ([]*kolide.Label, error)

type LabelQueriesForHostFunc func(host *kolide.Host, cutoff time.Time) (map[string]string, error)

type RecordLabelQueryExecutionsFunc func(host *kolide.Host, results map[uint]bool, t time.Time) error

type ListLabelsForHostFunc func(hid uint) ([]kolide.Label, error)

type ListHostsInLabelFunc func(lid uint) ([]kolide.Host, error)

type ListUniqueHostsInLabelsFunc func(labels []uint) ([]kolide.Host, error)

type SearchLabelsFunc func(query string, omit ...uint) ([]kolide.Label, error)

type SaveLabelFunc func(label *kolide.Label) (*kolide.Label, error)

type LabelStore struct {
	NewLabelFunc        NewLabelFunc
	NewLabelFuncInvoked bool

	DeleteLabelFunc        DeleteLabelFunc
	DeleteLabelFuncInvoked bool

	LabelFunc        LabelFunc
	LabelFuncInvoked bool

	ListLabelsFunc        ListLabelsFunc
	ListLabelsFuncInvoked bool

	LabelQueriesForHostFunc        LabelQueriesForHostFunc
	LabelQueriesForHostFuncInvoked bool

	RecordLabelQueryExecutionsFunc        RecordLabelQueryExecutionsFunc
	RecordLabelQueryExecutionsFuncInvoked bool

	ListLabelsForHostFunc        ListLabelsForHostFunc
	ListLabelsForHostFuncInvoked bool

	ListHostsInLabelFunc        ListHostsInLabelFunc
	ListHostsInLabelFuncInvoked bool

	ListUniqueHostsInLabelsFunc        ListUniqueHostsInLabelsFunc
	ListUniqueHostsInLabelsFuncInvoked bool

	SearchLabelsFunc        SearchLabelsFunc
	SearchLabelsFuncInvoked bool

	SaveLabelFunc        SaveLabelFunc
	SaveLabelFuncInvoked bool
}

func (s *LabelStore) NewLabel(Label *kolide.Label, opts ...kolide.OptionalArg) (*kolide.Label, error) {
	s.NewLabelFuncInvoked = true
	return s.NewLabelFunc(Label, opts...)
}

func (s *LabelStore) DeleteLabel(lid uint) error {
	s.DeleteLabelFuncInvoked = true
	return s.DeleteLabelFunc(lid)
}

func (s *LabelStore) Label(lid uint) (*kolide.Label, error) {
	s.LabelFuncInvoked = true
	return s.LabelFunc(lid)
}

func (s *LabelStore) ListLabels(opt kolide.ListOptions) ([]*kolide.Label, error) {
	s.ListLabelsFuncInvoked = true
	return s.ListLabelsFunc(opt)
}

func (s *LabelStore) LabelQueriesForHost(host *kolide.Host, cutoff time.Time) (map[string]string, error) {
	s.LabelQueriesForHostFuncInvoked = true
	return s.LabelQueriesForHostFunc(host, cutoff)
}

func (s *LabelStore) RecordLabelQueryExecutions(host *kolide.Host, results map[uint]bool, t time.Time) error {
	s.RecordLabelQueryExecutionsFuncInvoked = true
	return s.RecordLabelQueryExecutionsFunc(host, results, t)
}

func (s *LabelStore) ListLabelsForHost(hid uint) ([]kolide.Label, error) {
	s.ListLabelsForHostFuncInvoked = true
	return s.ListLabelsForHostFunc(hid)
}

func (s *LabelStore) ListHostsInLabel(lid uint) ([]kolide.Host, error) {
	s.ListHostsInLabelFuncInvoked = true
	return s.ListHostsInLabelFunc(lid)
}

func (s *LabelStore) ListUniqueHostsInLabels(labels []uint) ([]kolide.Host, error) {
	s.ListUniqueHostsInLabelsFuncInvoked = true
	return s.ListUniqueHostsInLabelsFunc(labels)
}

func (s *LabelStore) SearchLabels(query string, omit ...uint) ([]kolide.Label, error) {
	s.SearchLabelsFuncInvoked = true
	return s.SearchLabelsFunc(query, omit...)
}

func (s *LabelStore) SaveLabel(label *kolide.Label) (*kolide.Label, error) {
	s.SaveLabelFuncInvoked = true
	return s.SaveLabelFunc(label)
}
